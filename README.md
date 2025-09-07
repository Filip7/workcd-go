# workcd-g

A fast directory changer for your terminal using fuzzy finding. Quickly navigate to any subdirectory within your workspace using `fzf` for interactive selection.

This started as a zsh script that was vibe coded. But I wanted more features and easier editing of the code.  
So this is a vibe coded Go port, but it's human verified.

> [!NOTE]
> Plan is to rewrite as many parts as possible in version 2.0, because I'm not happy with many vibe coded parts, and I prefer writing code myself.
> Work in progress...

## Features

- **Fast fuzzy finding**: Uses `fzf` for interactive directory selection
- **Configurable base directory**: Set your workspace root in a YAML config file
- **Editor integration**: Optionally open your editor after changing directories
- **Shell integration**: Seamlessly integrates with bash/zsh through a simple function

## Installation

### Prerequisites

- Go 1.24.6 or later
- `fzf` (fuzzy finder) - Install from [junegun/fzf](https://github.com/junegunn/fzf)

### Build from source

```bash
git clone https://github.com/Filip7/workcd-go.git
cd workcd-go
go build -o workcd-go
```

### Install to PATH

```bash
# Copy the binary to a directory in your PATH
sudo cp workcd-go /usr/local/bin/
```

**Important**: Shell integration is required for this program to work. The binary alone cannot change your shell's current directory - you must set up the shell function as described in the Shell Integration section below.

## Configuration

Create a configuration file at `~/.config/workcd/config.yaml`:

```yaml
base_dir: ~/Workspace # Your workspace root directory
editor: nvim # Your preferred editor (optional)
preview_viewer: glow # Tool for previewing README files
```

If `base_dir` is not set, it defaults to `~/Workspace`.

## Shell Integration

Add this function to your `~/.bashrc` or `~/.zshrc`:

```bash
function wd() {
    # Temporary file for capturing stderr
    local stderr_file=$(mktemp)

    # Run the command, capturing stdout normally and redirecting stderr to the temp file
    local stdout
    stdout=$(workcd-go "$@" 2>"$stderr_file")

    # Check if there was any error output
    if [[ -s "$stderr_file" ]]; then
        local stderr_content
        stderr_content=$(cat "$stderr_file")
        echo "Error: $stderr_content" >&2
        rm -f "$stderr_file"
        return 1
    fi

    # Clean up the temp file
    rm -f "$stderr_file"

    # Process stdout if no errors
    if [[ -n "$stdout" ]]; then
        eval "$stdout"
    fi
}
```

Then reload your shell configuration:

```bash
# For bash
source ~/.bashrc

# For zsh
source ~/.zshrc
```

## Usage

### Basic usage

```bash
# Change to the base directory (no argument)
wd

# interactive selection of base directory, pass "/"
wd /

# Search for directories containing "project"
wd project

# Change directory and open editor
wd -e workcd-go

# Specify editor for this session
wd --editor nvim workcd-go
wd --editor idea java-project

# Specify base directory for this session
wd --base-dir ~/Projects proj1

# define the tool to preview readme files of the directories
wd --preview-viewer glow /
```

`--base-dir` enables you to do a neat trick of creating aliases that search different directories  
For example:

```bash
alias px="wd --base-dir ~/Workspace/projectx"
alias docs="wd --base-dir ~/Documents"
```

### Command line options

- `-e`: Open editor after changing directory
- `--editor <editor>`: Specify editor (overrides config and $EDITOR)
- `--base-dir <path>`: Specify base directory for workcd-go (overrides config)
- `--preview-viewer <viewer>`: Specify viewer for looking at README files in the directories

## How it works

1. Reads your base directory from config (defaults to `~/Workspace`)
2. Lists all subdirectories in your workspace
3. Uses `fzf` for interactive fuzzy finding
4. Outputs a `cd` command to change to the selected directory
5. Optionally opens your editor in the new directory
