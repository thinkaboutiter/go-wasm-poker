#!/bin/bash

# Build script for Go WebAssembly Texas Hold'em Poker

# Ensure GOROOT and GOPATH are set
if [ -z "$GOROOT" ]; then
  export GOROOT=/usr/local/go
fi

if [ -z "$GOPATH" ]; then
  export GOPATH=$HOME/go
fi

# Add Go binary directory to PATH
export PATH=$PATH:$GOROOT/bin

# Set GOOS and GOARCH for WebAssembly
export GOOS=js
export GOARCH=wasm

echo "Building Go WebAssembly Texas Hold'em Poker..."

# Create web directory if it doesn't exist
mkdir -p web

# Copy WebAssembly execution JavaScript support file
cp $GOROOT/misc/wasm/wasm_exec.js web/

# Build the WebAssembly binary
echo "Building WebAssembly binary..."
go build -o web/poker.wasm cmd/poker/main.go

# Reset environment variables
unset GOOS
unset GOARCH

# Build the server
echo "Building server..."
go build -o server cmd/server/main.go

echo "Build complete!"
echo "Run './server' to start the server, then visit http://localhost:8080"
