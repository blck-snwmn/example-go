#!/bin/bash

# Record current working directory
ORIG_DIR=$(pwd)

# Exit if no argument is provided
if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <directory_name>"
    exit 1
fi

# Set variable from the specified directory name
DIR_NAME=$1

# Create directory
mkdir -p "$DIR_NAME"

# Change to the created directory
cd "$DIR_NAME"

# Get current directory and extract the path until 'github.com'
CURRENT_PATH=$(pwd)
GITHUB_PATH=$(echo $CURRENT_PATH | sed -n 's/.*\(github\.com.*\)/\1/p')

# Execute go mod init command
go mod init $GITHUB_PATH

# Add new directory to dependabot.yml using yq
REPO_ROOT=$(git rev-parse --show-toplevel)
DEPENDABOT_FILE="$REPO_ROOT/.github/dependabot.yml"
NEW_DIR="./$DIR_NAME"

# Check if the directory is already in dependabot.yml
if ! yq eval '.updates[0].directories[] | select(. == "'$NEW_DIR'")' "$DEPENDABOT_FILE" | grep -q "$NEW_DIR"; then
    # Add the new directory to the gomod package-ecosystem directories array
    yq eval '.updates[0].directories += ["'$NEW_DIR'"]' -i "$DEPENDABOT_FILE"
    # Sort the directories array to maintain order
    yq eval '.updates[0].directories |= sort' -i "$DEPENDABOT_FILE"
    echo "Added $DIR_NAME to dependabot.yml"
else
    echo "$DIR_NAME already exists in dependabot.yml"
fi

# Return to the original working directory
cd "$ORIG_DIR"
