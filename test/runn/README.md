## Run
note: the current directory structure requires the `--scopes read:parent` option.

```bash
$ USER=blck-snwmn go tool runn run --scopes read:parent books/example-github.yaml
```

```bash
$ go tool runn run --scopes read:parent books/example-ownserver.yaml
```

```bash
$ go tool runn run --scopes read:parent books/example-o*.yaml --concurrent on
```

## Test
```bash
$ go test -v
```
