# --- Build Stage ---
FROM golang:1.23-bookworm AS builder

# Install C++ build tools
RUN apt-get update && apt-get install -y \
    g++ \
    make \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copy the monorepo structure
COPY Makefile .
COPY engine/ ./engine/
COPY api/ ./api/

# Create the lib directory for our static archive
RUN mkdir -p lib

# Compile C++ engine into a static library (.a)
# We avoid dynamic linking (.so) inside the container to make the binary truly portable
RUN g++ -c -O3 -fPIC -Iengine engine/engine.cpp -o engine/engine.o && \
    ar rcs lib/libengine.a engine/engine.o

# Build the Go application, linking against our new static library
# We use the -tags netgo,osusergo for extra portability
WORKDIR /app/api
RUN CGO_ENABLED=1 \
    CC=gcc \
    CXX=g++ \
    CGO_LDFLAGS="-L/app/lib -lengine -lstdc++" \
    go build -tags netgo,osusergo -ldflags="-linkmode external -extldflags '-static'" -o /app/inferx cmd/server/main.go

# --- Final Runtime Stage ---
FROM debian:bookworm-slim

WORKDIR /app

# Copy the compiled binary from the builder
COPY --from=builder /app/inferx .

# Copy configurations
COPY api/configs/ ./configs/

# Expose the server port
EXPOSE 8080

# Run the server
ENTRYPOINT ["./inferx"]
