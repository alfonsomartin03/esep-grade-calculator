# ESEP Grade Calculator

This repository contains a small Go library for calculating a course's final
letter grade from assignment, exam, and essay grades.

## Tech stack

- Go 1.24.4 (see `grade-calculator/go.mod`)
- Go's standard `testing` package for unit tests

## Repository files

- `README.MD` describes the project and how to run its tests.
- `grade-calculator/go.mod` declares the Go module path (`esep/grade-calculator`)
  and the Go language version.
- `grade-calculator/grade_calculator.go` defines grade types, the
  `GradeCalculator` API, and the grade calculation helpers.
- `grade-calculator/grade_calculator_test.go` contains unit tests for the
  calculator's letter-grade results.

## Run the tests

Install Go 1.24.4 or later, then run these commands from the repository root:

```sh
cd grade-calculator
go test ./...
```

To run the tests and report statement coverage:

```sh
cd grade-calculator
go test -coverprofile=coverage.txt ./...
go tool cover -func=coverage.txt
```

