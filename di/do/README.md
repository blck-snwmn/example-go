# samber/do Dependency Injection Example

This example demonstrates how to use [samber/do](https://github.com/samber/do) for dependency injection in Go.

## Features

- Simple HTTP server with user management API
- In-memory database implementation
- Dependency injection using samber/do v2
- Clean separation of concerns

## samber/do Key Features

- **Type-safe**: Uses Go generics for type safety
- **Simple API**: Easy to use `do.Provide` and `do.Invoke`
- **No reflection**: Compile-time safety
- **Lightweight**: Minimal overhead

## Project Structure

```
.
├── main.go         # Application entry point
├── container.go    # DI container setup using samber/do
├── server.go       # HTTP server implementation
├── service.go      # Business logic layer
└── database.go     # Data access layer
```

## Dependencies

- Database Layer: Provides data storage interface
- Service Layer: Uses Database for business logic
- Server Layer: Uses Service for HTTP handlers
- App: Coordinates Server with configuration

## Running the Example

```bash
go run .
```

The server will start on port 8080 with the following endpoints:

- `GET /users` - List all users
- `POST /users` - Create a new user (JSON: `{"name": "username"}`)

## samber/do vs Other DI Libraries

### Advantages:
- Type-safe with generics
- Simple, intuitive API
- No code generation needed
- Good performance (no reflection)

### Key Differences from uber/dig:
- Uses generics instead of reflection
- More explicit dependency resolution
- Simpler error handling
- Less magic, more predictable behavior

## Example Usage

```bash
# List users
curl http://localhost:8080/users

# Create a user
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name": "new_user"}'
```