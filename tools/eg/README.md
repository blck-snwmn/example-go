## Run
```bash
$ eg -t template/template.go before/sample.go  > after/sample.go
```

rewrite
```bash
$ eg -t template/template.go -w before/sample.go
```

## Note
The import of `eg` refers to the import of the file to be refactored.
It is not the import of a template.
