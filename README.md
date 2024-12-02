# Learn Go with Tests

[![Learn Go with Tests](assets/tdd_go.jpg)](https://quii.gitbook.io/learn-go-with-tests)

This is a collection of [Go](https://golang.org) code examples that show how to write [TDD](https://en.wikipedia.org/wiki/Test-driven_development) tests.

## Links from the book

- [New module changes in Go 1.16](https://go.dev/blog/go116-module-changes)
  - [Using Go Modules](https://go.dev/blog/using-go-modules)
- [go.mod file reference](https://go.dev/doc/modules/gomod-ref)
- [fmt package](https://pkg.go.dev/fmt#hdr-Printing)
- [Five suggestions for setting up a Go project](https://dave.cheney.net/2014/12/01/five-suggestions-for-setting-up-a-go-project)
- [Testable Examples in Go](https://go.dev/blog/examples)

## Progress

1. **Helloworld**: Completed implementation of the [Hello](helloworld/hello.go#L15-L21) function with tests for different languages (English, Spanish, French) and default behavior for empty input.
2. **Integers**: Implemented the [Add](integers/adder.go#L6-L13) function with tests for basic arithmetic operations and logging of results.
   1. **Examples**: Completed implementation of the `ByAge` struct and its methods for sorting a slice of [Person](examples/sort_test.go#L8-L15) structs by age.
