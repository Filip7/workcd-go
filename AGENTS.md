# AGENTS.md

## Bash Commands

- `go run main.go`: Runs the Go application
- `go test ./...`: Runs all tests in the project
- `go mod tidy`: Ensures the go.mod file is up to date
- `go fmt ./...`: Formats all Go code in the project
- `go vet ./...`: Runs static analysis on the code
- `go build`: Builds the Go application
- `go install`: Installs the Go application
- `go clean`: Removes object files from the package source directory
- `go doc`: Shows documentation for a package or symbol
- `go env`: Prints Go environment information
- `go list`: Lists packages or modules
- `go mod download`: Downloads dependencies
- `go mod edit`: Edits the go.mod file
- `go mod graph`: Prints the module dependency graph
- `go mod init`: Initializes a new module
- `go mod vendor`: Creates a vendor directory
- `go mod verify`: Verifies the dependencies
- `go mod why`: Explains why packages or modules are needed
- `go run`: Compiles and runs Go programs
- `go test`: Tests packages
- `go tool`: Runs specified Go tools
- `go version`: Prints Go version information

## Git Commands

- `git status`: Shows the status of the working directory and staging area
- `git add <file>`: Adds a file to the staging area
- `git commit -m "<message>"`: Commits changes with a message
- `git push`: Pushes changes to the remote repository
- `git pull`: Fetches and merges changes from the remote repository
- `git branch`: Lists all branches
- `git checkout <branch>`: Switches to a specified branch
- `git merge <branch>`: Merges a specified branch into the current branch
- `git log`: Shows the commit history
- `git diff`: Shows changes between commits, commit and working tree, etc.
- `git clone <repository>`: Clones a repository into a new directory
- `git remote -v`: Lists all remote repositories
- `git fetch`: Downloads objects and refs from another repository
- `git reset`: Resets the current HEAD to the specified state
- `git stash`: Stashes changes in a dirty working directory
- `git tag`: Creates, lists, deletes, or verifies a tag object signed with GPG

## GitHub Commands

- `gh pr create`: Creates a pull request
- `gh pr list`: Lists pull requests
- `gh pr checkout <number>`: Checks out a pull request
- `gh pr view <number>`: Views a pull request
- `gh pr merge <number>`: Merges a pull request
- `gh issue create`: Creates an issue
- `gh issue list`: Lists issues
- `gh issue view <number>`: Views an issue
- `gh issue close <number>`: Closes an issue
- `gh release create`: Creates a release
- `gh release list`: Lists releases
- `gh release view <tag>`: Views a release
- `gh release download <tag>`: Downloads a release

## Other Commands

- `make`: Executes commands from a Makefile
- `make test`: Runs tests using the Makefile
- `make build`: Builds the project using the Makefile
- `make clean`: Cleans the project using the Makefile
- `make install`: Installs the project using the Makefile
- `make run`: Runs the project using the Makefile
- `make lint`: Lints the project using the Makefile
- `make fmt`: Formats the project using the Makefile
- `make vet`: Runs static analysis on the project using the Makefile
- `make doc`: Generates documentation for the project using the Makefile
- `make env`: Prints environment information for the project using the Makefile
- `make list`: Lists packages or modules for the project using the Makefile
- `make download`: Downloads dependencies for the project using the Makefile
- `make edit`: Edits the project configuration using the Makefile
- `make graph`: Prints the dependency graph for the project using the Makefile
- `make init`: Initializes the project using the Makefile
- `make vendor`: Creates a vendor directory for the project using the Makefile
- `make verify`: Verifies the dependencies for the project using the Makefile
- `make why`: Explains why packages or modules are needed for the project using the Makefile
- `make tool`: Runs specified tools for the project using the Makefile
- `make version`: Prints version information for the project using the Makefile

## Tools

- `gofmt`: Formats Go source code
- `goimports`: Updates Go import lines
- `golint`: Lints Go source code
- `gometalinter`: Runs multiple linters on Go source code
- `goconst`: Finds repeated strings that could be replaced by constants
- `gocyclo`: Computes cyclomatic complexities of functions
- `go vet`: Examines Go source code and reports suspicious constructs
- `gosec`: Scans Go source code for security problems
- `gocover`: Runs coverage analysis on Go source code
- `gocritic`: Provides diagnostics for Go source code
- `godoctor`: Provides refactoring suggestions for Go source code
- `gopls`: Language server for Go
- `delve`: Debugger for Go
- `pprof`: Profiling tool for Go
- `godoc`: Documentation tool for Go
- `go mod`: Module maintenance tool for Go
- `go build`: Compiles Go source code
- `go install`: Compiles and installs Go packages
- `go test`: Tests Go packages
- `go run`: Compiles and runs Go programs
- `go tool`: Runs specified Go tools
- `go version`: Prints Go version information
- `go env`: Prints Go environment information
- `go list`: Lists packages or modules
- `go mod download`: Downloads dependencies
- `go mod edit`: Edits the go.mod file
- `go mod graph`: Prints the module dependency graph
- `go mod init`: Initializes a new module
- `go mod vendor`: Creates a vendor directory
- `go mod verify`: Verifies the dependencies
- `go mod why`: Explains why packages or modules are needed

## Agents

- `general`: General-purpose agent for researching complex questions, searching for code, and executing multi-step tasks. When you are searching for a keyword or file and are not confident that you will find the right match in the first few tries use this agent to perform the search for you.
