# Go 1.14

Welcome to the Go 1.14 Release Party!

https://www.crowdcast.io/e/pdx-go-march-12-2020

## Loops

Go 1.14 comes with a change in how the runtime does preemption of Goroutines.

- Release notes are [here](https://golang.org/doc/go1.14#runtime)
- Code samples are [here](./01-loop_preemption.go)

## Tests

When we're writing tests, we often have to clean up after ourselves after we're done. For example, we might need to delete a database table or close a connection. Before 1.14, we had to come up with our own ways to reliably clean up. Now, Go has a native way to cleanup. The [`*testing.T`](https://golang.org/pkg/testing/#T.Cleanup) and [`*testing.B`](https://golang.org/pkg/testing/#B.Cleanup) types both have a `Cleanup(func())` method on them.

- Release notes are [here](https://golang.org/doc/go1.14#testing)
- Code samples are [here](./02-test-cleanup/main_test.go)
