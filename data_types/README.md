# Go Data Types Examples

This directory contains examples of basic data types in Go programming language. Each file demonstrates different data types and their usage.

## Files Overview

### 1. strings.go
This file demonstrates string operations in Go:
- Declares a string variable `s` with the value "I am string"
- Uses the `strings` package to perform string operations
- Shows how to use `strings.HasSuffix()` to check if a string ends with a specific suffix
- Demonstrates string formatting using `fmt.Printf()`

### 2. numbers.go
This file shows examples of numeric data types in Go:
- Demonstrates integer (`int`) variables:
  - Uses short declaration (`:=`) to create variable `i` with value 10
  - Shows explicit type declaration with `var j int = 1234`
- Shows floating-point number (`float64`) usage:
  - Declares variable `f` with value 1.5
- Demonstrates basic arithmetic operations (addition)
- Uses `fmt.Printf()` for formatted output

### 3. boolean.go
This file illustrates boolean data type usage in Go:
- Declares two boolean variables:
  - `a` set to `true`
  - `b` set to `false`
- Demonstrates boolean comparison using `==` operator
- Shows how to print boolean values using `fmt.Printf()`

## Running the Examples
To run any of these examples, use the following command:
```bash
go run filename.go
```

For example:
```bash
go run strings.go
```

Each file is a complete, runnable program that demonstrates basic data type operations in Go. 