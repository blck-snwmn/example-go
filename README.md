# example-go
Example Go

## Usage of genmod

genmod automatically creates a new Go module, registers it in go.work, and updates the directories list in dependabot configuration.

Usage:
  1. From any directory in the repository, run:
  
         $ ./genmod.sh <directory_name>
  
  2. The specified directory will be created and its go.mod initialized.
  3. The new module will be added to go.work and the gomod directories list in .github/dependabot.yml will be updated.

## Run Tests
You can also run all tests using the `run_tests.sh` script in the repository root.
