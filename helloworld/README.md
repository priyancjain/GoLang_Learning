# Go Hello World Program

This is a simple Go program that prints "hello world" to the console. It serves as a basic example of Go programming language syntax and structure.

## Project Structure

```
helloworld/
├── helloworld.go    # Main program file
└── README.md        # This documentation file
```

## Code Explanation

The program consists of three main parts:

1. **Package Declaration**
   ```go
   package main
   ```
   - Every Go program starts with a package declaration
   - `main` is a special package name that indicates this is an executable program

2. **Import Statement**
   ```go
   import "fmt"
   ```
   - Imports the `fmt` package which contains functions for formatted I/O
   - `fmt` is part of Go's standard library

3. **Main Function**
   ```go
   func main() {
       fmt.Println("hello world")
   }
   ```
   - `main()` is the entry point of the program
   - `fmt.Println()` is used to print text to the console
   - The program prints "hello world" when executed

## How to Run

1. Make sure you have Go installed on your system
   - You can check by running `go version` in your terminal
   - If not installed, download from [golang.org](https://golang.org)

2. Navigate to the project directory:
   ```bash
   cd helloworld
   ```

3. Run the program:
   ```bash
   go run helloworld.go
   ```

   Or build and run:
   ```bash
   go build helloworld.go
   ./helloworld
   ```

## Expected Output

When you run the program, you should see:
```
hello world
```

## Git Repository Management

### Initializing Git Repository
To initialize a Git repository in your project:
```bash
git init
```

### Removing Git Repository
If you want to remove the Git repository from your project:
```bash
# Remove the .git directory
rm -rf .git
```

This will completely remove Git version control from your project. Make sure you want to do this as it will delete all Git history and configuration.
