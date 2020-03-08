# Go 1.14

Welcome to the Go 1.14 Release Party!

https://www.crowdcast.io/e/pdx-go-march-12-2020

## Loops

Go 1.14 comes with a change in how the runtime does preemption of Goroutines.

- Release notes are [here](https://golang.org/doc/go1.14#runtime)
- Code samples are [here](./preemption/main.go)

## Tests

Go 1.14 comes with two nice additions for testing. Previously, we had to implement this stuff ourselves and it was hard.

## Cleaning up

When we're writing tests, we often have to clean up after ourselves after we're done. For example, we might need to delete a database table or close a connection. Before 1.14, we had to come up with our own ways to reliably clean up. Now, Go has a native way to cleanup. The [`*testing.T`](https://golang.org/pkg/testing/#T.Cleanup) and [`*testing.B`](https://golang.org/pkg/testing/#B.Cleanup) types both have a `Cleanup(func())` method on them.

- Release notes are [here](https://golang.org/doc/go1.14#testing)
- Code samples are [here](./tests/cleanup_test.go)
  - Run them with `go test -v cleanup_test.go`

## Log streaming

When we ran `go test -v`, Go waited until the end of the test to print all of the strings passed to `t.Log`. 

Sometimes it's more helpful to have the output streamed as it happens. Now, if you run `go test -v`, each `t.Log` call gets flushed to `STDOUT` as it happens.

- Release notes are [here](https://golang.org/doc/go1.14#go-command) (see the "testing" header at the bottom of the section)
- Code samples are [here](./tests/stream_test.go)
  - Run them with `go test -v stream_test.go`

## Overlapping interfaces

Before 1.14, you couldn't do this:

```go
type Interface1 {
    Hello()
}

type Interface2 {
    Hello()
}

type Interface3 {
    Interface1
    Interface2
}
```

But now you can!
