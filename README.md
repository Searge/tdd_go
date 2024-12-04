# Learn Go with Tests

[![Learn Go with Tests](assets/tdd_go.webp)](https://quii.gitbook.io/learn-go-with-tests)

[![Visits Badge](https://badges.pufler.dev/visits/Searge/tdd_go?style=flat-square&logo=refinedgithub&logoColor=white)](https://badges.pufler.dev)
![Top language](https://img.shields.io/github/languages/top/Searge/tdd_go?style=flat-square)
![Repo size](https://img.shields.io/github/repo-size/Searge/tdd_go?style=flat-square)
![Created](https://img.shields.io/date/1732961863.svg?style=flat-square&logo=github&label=created)
[![Updates](https://img.shields.io/github/last-commit/Searge/tdd_go.svg?style=flat-square&logo=git&logoColor=white&color=blue)](https://github.com/Searge/tdd_go/commits/main/)
[![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/Searge/tdd_go/go.yml?branch=main&style=flat-square&logo=githubactions&logoColor=white&label=test-n-build)](https://github.com/Searge/tdd_go/actions/workflows/go.yml)

This is a collection of [Go](https://golang.org) code examples that show how to write [TDD](https://en.wikipedia.org/wiki/Test-driven_development) tests.

## Links from the book

- [New module changes in Go 1.16](https://go.dev/blog/go116-module-changes)
  - [Using Go Modules](https://go.dev/blog/using-go-modules)
- [go.mod file reference](https://go.dev/doc/modules/gomod-ref)
- [fmt package](https://pkg.go.dev/fmt#hdr-Printing)
- [Five suggestions for setting up a Go project](https://dave.cheney.net/2014/12/01/five-suggestions-for-setting-up-a-go-project)
- [Testable Examples in Go](https://go.dev/blog/examples)
- [Testing | Benchmarking](https://pkg.go.dev/testing#hdr-Benchmarks)
- [The cover story](https://go.dev/blog/cover)

## Progress

1. **Helloworld**: Completed implementation of the [Hello](helloworld/hello.go#L15-L21) function with tests for different languages (English, Spanish, French) and default behavior for empty input.
2. **Integers**: Implemented the [Add](integers/adder.go#L6-L13) function with tests for basic arithmetic operations and logging of results.
   1. **Examples**: Completed implementation of the `ByAge` struct and its methods for sorting a slice of [Person](examples/sort_test.go#L8-L15) structs by age.
3. **Iteration**: Implemented the [Repeat](iteration/repeat.go#L3-L9) function with tests for repeating a character a specified number of times.
