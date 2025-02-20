## Measure coverage

### Run

```bash
$ go test -cover ./...
```

## Incorporate coverage measurement into binary

### Build

```bash
$ go build --cover ./cmd/main.go
```

### Run

```bash
$ GOCOVERDIR=dir ./main
```

### Test

```bash
$ LOCAL=true go test ./... --count=1
```
