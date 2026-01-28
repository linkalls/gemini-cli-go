#!/bin/bash

# Installation script for Gemini CLI (Go version)

set -e

echo "Installing Gemini CLI (Go version)..."

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "Error: Go is not installed. Please install Go 1.21 or higher first."
    echo "Visit: https://go.dev/doc/install"
    exit 1
fi

# Build the binary
echo "Building binary..."
go build -o gemini main.go

# Move to /usr/local/bin (requires sudo on Linux/macOS)
if [ -w /usr/local/bin ]; then
    echo "Installing to /usr/local/bin..."
    mv gemini /usr/local/bin/
else
    echo "Installing to /usr/local/bin (requires sudo)..."
    sudo mv gemini /usr/local/bin/
fi

echo ""
echo "✅ Gemini CLI installed successfully!"
echo ""
echo "To get started:"
echo "  1. Set your API key: export GEMINI_API_KEY=your-api-key"
echo "  2. Run: gemini"
echo ""
echo "For help: gemini --help"
