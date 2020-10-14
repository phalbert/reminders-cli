## Parsing commands using flag sets

This is a much [better way](https://golang.org/pkg/flag/) of working with cmd args

```golang
flag.NewFlagSet(name string, errorHandling ErrorHandling) *FlagSet
```