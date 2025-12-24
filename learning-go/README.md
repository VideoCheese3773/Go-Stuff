# Learning Go Programming

Welcome to your Go learning journey! This is a starter project to help you get started with Go.

## Installation

### Windows

1. **Download Go:**
   - Visit https://go.dev/dl/
   - Download the Windows installer (`.msi` file)
   - Run the installer and follow the instructions

2. **Verify Installation:**
   ```powershell
   go version
   ```

3. **Set up your workspace:**
   - Go will be installed to `C:\Program Files\Go` by default
   - The installer should automatically add Go to your PATH

## Getting Started

Once Go is installed:

1. **Initialize the Go module:**
   ```powershell
   go mod init learning-go
   ```

2. **Run your first program:**
   ```powershell
   go run main.go
   ```

3. **Build an executable:**
   ```powershell
   go build main.go
   ./main.exe
   ```

## Learning Path

### Basics (Start Here) ✅

The basics are now interactive! Run lessons to see examples and learn by doing.

**Available Lessons:**
1. **Variables and Types** - Learn how to declare and use variables in Go
2. **Functions** - Understand function syntax, parameters, and return values
3. **Control Structures** - Master if/else, for loops, and switch statements
4. **Arrays and Slices** - Work with fixed arrays and dynamic slices
5. **Maps** - Store and retrieve key-value pairs
6. **Structs** - Create custom data types and methods

**How to Use:**
```powershell
# See the menu
go run .

# Run a specific lesson (1-6)
go run . 1    # Variables and Types
go run . 2    # Functions
go run . 3    # Control Structures
go run . 4    # Arrays and Slices
go run . 5    # Maps
go run . 6    # Structs

# Run all lessons at once
go run . all
```

**Learning Tips:**
- Run each lesson and read the output carefully
- Try modifying the code in `basics.go` to experiment
- Uncomment and modify examples in `examples.go`
- Practice by creating your own variations of the examples

### Intermediate
- Pointers
- Methods and interfaces
- Error handling
- Goroutines and channels (concurrency)
- Packages and modules

### Advanced
- Reflection
- Testing
- Web development
- Database connectivity
- Building APIs

## Resources

- **Official Tutorial:** https://go.dev/doc/tutorial/
- **Tour of Go:** https://go.dev/tour/
- **Effective Go:** https://go.dev/doc/effective_go
- **Go by Example:** https://gobyexample.com/

## Next Steps

1. Install Go using the instructions above
2. Run `go mod init learning-go` to initialize the module
3. Run `go run main.go` to see your first program
4. Start modifying `main.go` to experiment with Go features!

