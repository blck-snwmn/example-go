# Go 1.27 examples

Small examples of language features introduced in Go 1.27:

- generic methods, demonstrated by a lazy, functional-style `Seq[T]` with
  `Map`, `Filter`, and `FlatMap`
- terminal sequence operations demonstrated by `Fold` and `Collect`
- nested field selectors in struct literals
- generic function type inference from an assignment context
- stricter decoding defaults in `encoding/json/v2`
- UUID v7 generation and parsing with the standard `uuid` package
- in-memory HTTP testing with `httptest.NewTestServer` and `testing/synctest`

## Run

```sh
go run .
go test ./...
```
