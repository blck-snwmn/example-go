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
go get --tool github.com/golangci/golangci-lint/cmd/golangci-lint@v1.62.2
go work use .

# Obtain repository root and file paths
REPO_ROOT=$(git rev-parse --show-toplevel)
WORK_FILE="$REPO_ROOT/go.work"
DEPENDABOT_FILE="$REPO_ROOT/.github/dependabot.yml"

# Get the latest list of directories from the use() block in go.work
NEW_DIRS=$(awk '/use \(/,/\)/{
    if ($1 ~ /^\.\//) print $1
}' "$WORK_FILE" | sed 's/^/      - "/; s/$/"/')

# Update the directories list in the gomod block of dependabot.yml
awk -v new_dirs="$NEW_DIRS" '
    BEGIN { inGomod=0; inDirs=0 }
    /- package-ecosystem: "gomod"/ { inGomod=1 }
    {
      if(inGomod && $0 ~ /^[[:space:]]*directories:/) {
         print "    directories:";
         print new_dirs;
         inDirs=1;
         inGomod=0; next
      }
      if(inDirs && $0 ~ /^[[:space:]]*[^-[:space:]]/) { inDirs=0 }
      if(!inDirs) { print }
    }
' "$DEPENDABOT_FILE" > "$DEPENDABOT_FILE.tmp" && mv "$DEPENDABOT_FILE.tmp" "$DEPENDABOT_FILE"

# Return to the original working directory
cd "$ORIG_DIR"
