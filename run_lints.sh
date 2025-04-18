#!/bin/bash

# Get the repository root directory
REPO_ROOT=$(git rev-parse --show-toplevel)

echo "Extracting module directories using go list..."

# Extract module directories using go list
DIRS=$(go list -m -json | jq -s '.' | jq -c '.[].Dir' | tr -d '"')

# Run lint in each directory
for dir in $DIRS; do
    echo "Running lint in directory: $dir"
    (cd "$dir" && go tool golangci-lint run --enable=gosec)
    if [ $? -ne 0 ]; then
         echo "Lint failed in $dir."
         exit 1
    fi
done

echo "All lint ran successfully."
