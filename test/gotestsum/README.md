## Run
```bash
$ go tool gotestsum
```

### Format
```bash
$ go tool gotestsum --format testname
```

```bash
$ go tool gotestsum --format testdox
```

### Pick slow test
```bash
$ go tool gotestsum --format testdox  --jsonfile tmp.json.log   --post-run-command "bash -c '
    echo; echo Slowest tests;
    go tool gotestsum tool slowest --num 10 --jsonfile tmp.json.log'"
```