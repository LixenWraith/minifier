#!/bin/bash

# Save current GOOS and GOARCH
ORIGINAL_GOOS=$GOOS
ORIGINAL_GOARCH=$GOARCH
ORIGINAL_CGO_ENABLED=$CGO_ENABLED

# Set environment for FreeBSD 64-bit
export GOOS=freebsd
export GOARCH=amd64
export CGO_ENABLED=0

# Compile the program
echo "Compiling for FreeBSD 64-bit..."
go build -o bin/minifier_freebsd_amd64 main.go

# Check if compilation was successful
if [ $? -eq 0 ]; then
    echo "Compilation successful. FreeBSD 64-bit executable created at bin/minifier_freebsd_amd64"
else
    echo "Compilation failed."
fi

# Restore original GOOS and GOARCH
export GOOS=$ORIGINAL_GOOS
export GOARCH=$ORIGINAL_GOARCH
export CGO_ENABLED=$ORIGINAL_CGO_ENABLED

echo "Environment restored to original settings."
