#!/bin/bash

# Get the repository root directory
REPO_ROOT=$(git rev-parse --show-toplevel)

echo "Extracting module directories using go list..."

# Extract module directories using go list
DIRS=$(go list -m -json | jq -s '.' | jq -c '.[].Dir' | tr -d '"')

# Run go test in each directory
for dir in $DIRS; do
    echo "Running tests in directory: $dir..."
    if [[ "$dir" == */test/ginkgo ]]; then
        (cd "$dir" && go tool ginkgo -p)
    else
        (cd "$dir" && go test ./...)
    fi
    if [ $? -ne 0 ]; then
         echo "Tests failed in $dir."
         exit 1
    fi
done

echo "All tests ran successfully."
