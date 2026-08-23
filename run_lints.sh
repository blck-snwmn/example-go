#!/bin/bash

echo "Finding directories with go.mod files..."

# Extract directories containing go.mod files using find
DIRS=$(find . -name "go.mod" -exec dirname {} \;)

# Initialize failure flag
FAILED=0

if ! command -v golangci-lint >/dev/null 2>&1; then
    echo "golangci-lint is not available. Install the aqua-managed tools first." >&2
    exit 1
fi

# Run lint in each directory
for dir in $DIRS; do
    echo "Running lint in directory: $dir"
    (cd "$dir" && golangci-lint run --allow-parallel-runners)
    if [ $? -ne 0 ]; then
         echo "Lint failed in $dir."
         FAILED=1
    fi
done

if [ $FAILED -eq 1 ]; then
    echo "Some lint checks failed. Please fix the issues above."
    exit 1
else
    echo "All lint ran successfully."
fi
