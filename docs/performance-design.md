# Performance Design

## Performance Parameters

1. Queue Size
2. Worker Batch Size
3. Worker Batch Timeout
4. Number of Workers

## Performance Metrics

1. Latency (P50, P95, P 99)
2. Throughput (req/sec)
3. Queue Wait Time
4. Batch Processing Time
5. Worker Utilization
6. Request Drop Rate
7. Batch Processed Batch Size (P50, P95, P99)

## Performance Scenarios

1. Optimize for Latency
2. Optimize for Throughput

## Performance Testing

Run combination of following parameters and measure performance metrics:

1. Queue Size: 100, 1000, 10000
2. Worker Batch Size: 1, 4, 8, 16, 32, 64
3. Worker Batch Timeout: 10ms, 50ms, 100ms, 500ms, 1000ms
4. Number of Workers: 1, 2, 4, 8, 16, 32, 64

Create performance scenario to test with each possible combination. For e.g., 

1. Queue Size: 100, Worker Batch Size: 1, Worker Batch Timeout: 10ms, Number of Workers: 1
2. Queue Size: 100, Worker Batch Size: 1, Worker Batch Timeout: 10ms, Number of Workers: 2
.
.
.
N. Queue Size: 10000, Worker Batch Size: 64, Worker Batch Timeout: 1000ms, Number of Workers: 64

## Performance Result

Document performance analyasis for:
1. Latency Optimized
2. Throughput Optimized
3. Balanced

Create documents under docs/performance-analysis