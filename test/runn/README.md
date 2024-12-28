## Run
note: the current directory structure requires the `--scopes read:parent` option.

```bash
$ USER=blck-snwmn runn run --scopes read:parent books/example-github.yaml
```

```bash
$ runn run --scopes read:parent books/example-ownserver.yaml
```

```bash
$ runn run --scopes read:parent books/example-o*.yaml --concurrent on
```

## Test
```bash
$ go test -v
```
