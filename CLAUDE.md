# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Architecture

This is a Go monorepo containing multiple independent modules demonstrating various Go technologies and patterns. Each directory represents a separate Go module with its own go.mod file, organized around specific topics:

- **standard/**: Go standard library examples (slog, iter, unique, generics, HTTP)
- **test/**: Testing methodologies (standard testing, gotestsum, Ginkgo, coverage, runn)
- **openapi/**: API development with ogen and oapi-codegen
- **db/**: Database operations with sqlx and sqlc
- **cli/**: Command-line tool development
- **tools/**: Development utilities and code generation
- **di/**: Dependency injection examples

Each module is independent and managed separately. The repository scripts automatically discover modules by finding go.mod files.

## Essential Commands

### Module Management
```bash
# Create new module with automatic dependabot integration
./genmod.sh <directory_name>
```

### Testing
```bash
# Run all tests across all modules
./run_tests.sh

# Run tests in specific module
cd <module_directory> && go test ./...

# Run Ginkgo tests (for test/ginkgo module)
cd test/ginkgo && go tool ginkgo -p
```

### Linting
```bash
# Run linting across all modules (auto-detects golangci-lint in PATH)
./run_lints.sh
# If golangci-lint is installed: uses host command
# If not installed: uses go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.62.2

# Run linting in specific module
cd <module_directory> && go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.62.2 run --enable=gosec
```

### Basic Development
```bash
# Format code
go fmt ./...

# Install dependencies
go mod tidy

# Build all packages
go build ./...
```

## Development Workflow

1. **Adding new examples**: Use `./genmod.sh <name>` to create a new module, which automatically:
   - Creates the directory and initializes go.mod
   - Updates .github/dependabot.yml to include the new module

2. **Testing strategy**: The project uses different testing approaches:
   - Standard Go testing for basic unit tests
   - gotestsum for enhanced test output
   - Ginkgo for BDD-style testing with table tests
   - Coverage measurement tools

3. **Module structure**: Each module is self-contained with its own dependencies and can be developed independently

## Code Generation and Tools

- **OpenAPI**: Two approaches available - ogen (newer, more performant) and oapi-codegen (more mature)
- **Database**: sqlx for manual SQL, sqlc for code generation from SQL
- **Linting**: golangci-lint v1.62.2 with gosec security checking enabled (executed via go run)

## Important Notes

- golangci-lint is centrally managed and executed via `go run` (no per-module tool dependency)
- The repository root scripts handle cross-module operations automatically
- Each module can have different dependency versions as needed
- Ginkgo tests require special handling with `go tool ginkgo` command