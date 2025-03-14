# Ginkgo Sample

This directory contains examples of tests using the Ginkgo testing framework.

## Overview

Ginkgo is a testing framework for Go that allows you to write tests in a BDD (Behavior-Driven Development) style. When combined with the Gomega matcher library, you can write more expressive tests.

This sample demonstrates the following features:

- Basic Ginkgo test structure (Describe, Context, It)
- Setup and cleanup using BeforeEach/AfterEach
- Testing with mocks
- Testing asynchronous operations (Eventually, Consistently)
- Table-driven tests (DescribeTable, Entry)

## File Structure

- `calculator.go` - A package providing basic calculation functions
- `calculator_test.go` - Tests for Calculator
- `user.go` - A package handling user information
- `user_test.go` - Tests for UserService (including mocks and asynchronous operations)
- `table_test.go` - Examples of table-driven tests

## How to Run Tests

You can run all tests in this directory with the following command:

```bash
go test ./...
```

Or you can use the Ginkgo command:

```bash
go tool ginkgo
```

If you want to run only specific tests, you can use focus:

```bash
go tool ginkgo --focus="Calculator"
```

For parallel execution:

```bash
go tool ginkgo -p
```

## Main Features of Ginkgo

### Test Structure

- `Describe` - Defines a group of tests
- `Context` - Groups tests under specific conditions
- `It` - Defines individual test cases

### Setup and Cleanup

- `BeforeEach` - Code executed before each test
- `AfterEach` - Code executed after each test
- `BeforeSuite` - Code executed once before the entire test suite
- `AfterSuite` - Code executed once after the entire test suite

### Asynchronous Testing

- `Eventually` - Verifies that a condition is eventually met
- `Consistently` - Verifies that a condition is met continuously for a period of time

### Table-Driven Tests

- `DescribeTable` - Defines a table-driven test
- `Entry` - Defines test cases

## Main Gomega Matchers

- `Equal` - Verifies that values are equal
- `BeNil` - Verifies that a value is nil
- `HaveOccurred` - Verifies that an error occurred
- `Succeed` - Verifies that no error occurred
- `BeTrue`/`BeFalse` - Verifies boolean values
- `ContainElement` - Verifies that a slice contains a specific element
- `Receive` - Verifies that a value can be received from a channel
- `BeClosed` - Verifies that a channel is closed

## Reference Links

- [Ginkgo Official Documentation](https://onsi.github.io/ginkgo/)
- [Gomega Official Documentation](https://onsi.github.io/gomega/) 