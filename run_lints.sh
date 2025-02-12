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
    echo "Running lint in directory: $dir"
    (cd "$REPO_ROOT/$dir" && go tool golangci-lint run --enable=gosec)
    if [ $? -ne 0 ]; then
         echo "Lint failed in $dir."
         exit 1
    fi
done

echo "All lint ran successfully."
