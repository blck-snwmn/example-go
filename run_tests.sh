#!/bin/bash

# Get the repository root directory
REPO_ROOT=$(git rev-parse --show-toplevel)
WORK_FILE="$REPO_ROOT/go.work"

if [ ! -f "$WORK_FILE" ]; then
    echo "go.work not found."
    exit 1
fi

echo "Extracting module directories from go.work..."

# Extract directories from the use() block
DIRS=$(awk '/use \(/,/\)/{
    if ($1 ~ /^\.\//) print $1
}' "$WORK_FILE")

# Run go test in each directory
for dir in $DIRS; do
    echo "Running tests in directory: $dir..."
    if [ "$dir" = "./test/ginkgo" ]; then
        (cd "$REPO_ROOT/$dir" && go tool ginkgo -p)
    else
        (cd "$REPO_ROOT/$dir" && go test ./...)
    fi
    if [ $? -ne 0 ]; then
         echo "Tests failed in $dir."
         exit 1
    fi
done

echo "All tests ran successfully."
