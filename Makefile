# Variables
CXX = clang++
CXXFLAGS = -dynamiclib -fPIC -Iengine
LDFLAGS = -Llib
LIB_NAME = libengine.dylib
ENGINE_SRC = engine/engine.cpp
API_SERVER = api/cmd/server/main.go

# Standard build
all: lib bridge-test server

# Build the C++ shared library
lib:
	mkdir -p lib
	$(CXX) $(CXXFLAGS) -install_name @rpath/$(LIB_NAME) -o lib/$(LIB_NAME) $(ENGINE_SRC)

# Build the Go server
server: lib
	cd api && go build -o ../server cmd/server/main.go

# Verify the bridge with a test
test: lib
	cd api && DYLD_LIBRARY_PATH=$(PWD)/lib go test -v ./internal/engine/...

# Clean build artifacts
clean:
	rm -rf lib server
