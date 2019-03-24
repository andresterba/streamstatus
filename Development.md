streamstatus
==============================

To test new commands add this snippet to `go.mod`:

```
replace (
	github.com/andresterba/streamstatus/cmd => ./cmd
	github.com/andresterba/streamstatus/internal => ./internal
)
```