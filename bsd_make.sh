#!/bin/bash

# Save current GOOS, GOARCH, and CGO_ENABLED
ORIGINAL_GOOS=$GOOS
ORIGINAL_GOARCH=$GOARCH
ORIGINAL_CGO_ENABLED=$CGO_ENABLED

# Set environment for FreeBSD 64-bit and disable CGo
export GOOS=freebsd
export GOARCH=amd64
export CGO_ENABLED=0


# Compile the program
echo "Compiling for FreeBSD 64-bit with CGo disabled..."
go build -o bin/minifier_freebsd_amd64 main.go

# Check if compilation was successful
if [ $? -eq 0 ]; then
    echo "Compilation successful. FreeBSD 64-bit executable created at bin/minifier_freebsd_amd64"
else

    echo "Compilation failed."
fi

# Restore original GOOS, GOARCH, and CGO_ENABLED
export GOOS=$ORIGINAL_GOOS
export GOARCH=$ORIGINAL_GOARCH
export CGO_ENABLED=$ORIGINAL_CGO_ENABLED

echo "Environment restored to original settings."
