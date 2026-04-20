#!/bin/bash

# Configuration
RESULTS_FILE="docs/performance-analysis/results.csv"
# Load test command for different phases
# Phase 1: Latency (Moderate load)
# Phase 2: Throughput (High load)
# Phase 3: Stress (Extreme load)

# Header for results
mkdir -p docs/performance-analysis
echo "Category,BatchSize,Workers,QueueSize,TimeoutMS,RPS,P50,P95,P99,Success,Rejected" > $RESULTS_FILE

# Determine which docker compose command to use
if docker compose version >/dev/null 2>&1; then
    DOCKER_COMPOSE="docker compose"
elif docker-compose version >/dev/null 2>&1; then
    DOCKER_COMPOSE="docker-compose"
else
    echo "Error: Neither 'docker compose' nor 'docker-compose' found."
    exit 1
fi

run_test() {
    local category=$1
    local batch=$2
    local workers=$3
    local queue=$4
    local timeout=$5
    local load_n=$6
    local load_c=$7

    echo "----------------------------------------------------------"
    echo "Exp: $category (B:$batch, W:$workers, Q:$queue, T:$timeout)"
    echo "----------------------------------------------------------"

    # Start server
    BATCH_SIZE=$batch NUM_WORKERS=$workers QUEUE_SIZE=$queue BATCH_TIMEOUT_MS=$timeout $DOCKER_COMPOSE up -d --build

    # Wait for server to be ready
    echo "Waiting for server to start..."
    for i in {1..30}; do
        if curl --output /dev/null --silent --head --fail http://localhost:8080/metrics; then
            echo " Server is UP!"
            break
        fi
        printf '.'
        sleep 1
        if [ $i -eq 30 ]; then
            echo " Server failed to start!"
            $DOCKER_COMPOSE logs
            $DOCKER_COMPOSE down
            return
        fi
    done

    # Run load test and capture output
    temp_log=$(mktemp)
    echo "Running load test: -n $load_n -c $load_c"
    (cd api && go run cmd/loadtest/main.go -n $load_n -c $load_c) > "$temp_log" 2>&1
    
    # Extract metrics using more robust grep/sed
    rps=$(grep "Requests/sec:" "$temp_log" | sed 's/.*Requests\/sec:[[:space:]]*//')
    p50=$(grep "P50 Latency:" "$temp_log" | sed 's/.*P50 Latency:[[:space:]]*//' | sed 's/ms//')
    p95=$(grep "P95 Latency:" "$temp_log" | sed 's/.*P95 Latency:[[:space:]]*//' | sed 's/ms//')
    p99=$(grep "P99 Latency:" "$temp_log" | sed 's/.*P99 Latency:[[:space:]]*//' | sed 's/ms//')
    success=$(grep "Successes (202):" "$temp_log" | sed 's/.*Successes (202):[[:space:]]*//')
    rejected=$(grep "Rejections (429):" "$temp_log" | sed 's/.*Rejections (429):[[:space:]]*//')

    echo "Results: RPS=$rps, P50=$p50, P95=$p95, P99=$p99"
    
    # Save to CSV (handle empty values with "N/A")
    echo "$category,$batch,$workers,$queue,$timeout,${rps:-N/A},${p50:-N/A},${p95:-N/A},${p99:-N/A},${success:-0},${rejected:-0}" >> $RESULTS_FILE

    # Cleanup temp log
    rm "$temp_log"

    # Stop server
    $DOCKER_COMPOSE down
}

# --- 1. LATENCY FOCUS SWEEP (~18 runs) ---
# Goal: Find lowest possible response time under moderate load
echo ">>> PHASE 1: LATENCY FOCUS"
for w in 4 8 16; do
    for b in 1 2 4; do
        for t in 10 20; do
            run_test "Latency" $b $w 100 $t 500 50
        done
    done
done

# --- 2. THROUGHPUT FOCUS SWEEP (~18 runs) ---
# Goal: Find maximum sustainable RPS with reasonable P99
echo ">>> PHASE 2: THROUGHPUT FOCUS"
for w in 4 8 16; do
    for b in 8 16 32; do
        for t in 100 200; do
            run_test "Throughput" $b $w 1000 $t 1000 100
        done
    done
done

# --- 3. QUEUE STABILITY FOCUS (~5 runs) ---
# Goal: See how queue size impacts drop rate vs latency
echo ">>> PHASE 3: QUEUE STABILITY"
for q in 50 100 500 1000 5000; do
    run_test "Stability" 16 8 $q 50 2000 200
done

echo "Intelligent Sweep complete. Results in $RESULTS_FILE"
