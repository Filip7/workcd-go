# AGENTS.md

## Build Commands
- `go build` - Build the binary
- `go run .` - Run the application directly
- `go build -o workcd-go` - Build with specific output name

## Test Commands
- `go test ./...` - Run all tests (none currently exist)
- `go test -v ./...` - Run tests with verbose output
- `go test -run TestName` - Run specific test (when tests are added)

## Lint/Format Commands
- `go fmt ./...` - Format Go code
- `go vet ./...` - Run static analysis
- `go mod tidy` - Clean up dependencies

## Code Style Guidelines

### Imports
- Group imports: standard library first, then blank line, then third-party packages
- Use `gopkg.in/yaml.v3` for YAML parsing

### Naming Conventions
- Functions: camelCase (e.g., `getBaseDirFromConfig`)
- Structs: PascalCase (e.g., `Config`)
- Struct fields: PascalCase for exported, camelCase for unexported

### Error Handling
- Check `err != nil` immediately after operations that can fail
- Use `log.Fatal()` for fatal errors that should terminate the program
- Use `fmt.Fprintf(os.Stderr, ...)` for user-facing error messages
- Use `os.Exit()` for controlled program termination

### Types and Structs
- Use struct tags for YAML serialization: `yaml:"field_name"`
- Define config structs at package level

### Formatting
- Follow standard Go formatting (enforced by `go fmt`)
- Use meaningful variable names
- Keep functions focused and under 50 lines when possible