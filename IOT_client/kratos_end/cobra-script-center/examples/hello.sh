#!/bin/bash
# Simple hello world script for testing

echo "Hello from Cobra Script Center!"
echo "Current time: $(date)"
echo "Current user: $(whoami)"
echo "Working directory: $(pwd)"

# Check if any parameters were passed
if [ $# -gt 0 ]; then
    echo "Parameters received:"
    for param in "$@"; do
        echo "  - $param"
    done
fi

# Check environment variables
if [ ! -z "$NAME" ]; then
    echo "Hello, $NAME!"
fi

if [ ! -z "$MESSAGE" ]; then
    echo "Message: $MESSAGE"
fi

echo "Script execution completed successfully!"