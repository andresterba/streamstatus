streamstatus
==============================

To test new commands add this snippet to `go.mod`:

```
replace (
	github.com/andresterba/streamstatus/cmd => ./cmd
	github.com/andresterba/streamstatus/internal => ./internal
)
```

# Prepare new release

1. update version in `cmd/version.go`
2. update cmd module
3. add tag and release
