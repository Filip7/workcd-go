# workcd

Directory changer that supports opening your editor instantly within the directory.  
Uses `fzf` for interactive selection of directory when your query is ambiguous.

## Features

- **Fast fuzzy finding**: Uses `fzf` for interactive directory selection
- **Configurable base directory**: Set your workspace root in a YAML config file
- **Editor integration**: Optionally open your editor after changing directories
- **Shell integration**: currently uses a helper shell script

## Installation

### Prerequisites

- Go 1.25.1 or later
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

- Also supports `XDG_CONFIG_HOME`

```yaml
base_dir: ~/Workspace # Your workspace root directory (default is $HOME/Workspace)
editor: nvim # Your preferred editor (default is vi)
preview_viewer: glow # Tool for previewing README files (default is less)
```

If `base_dir` is not set, it defaults to `~/Workspace`.

## Shell Integration

Add this function to your `~/.bashrc` or `~/.zshrc`:

```bash
eval "$(workcd-go --eval zsh)"
```

This will make workcd-go available as `wd` in your shell.  
Additional shell support when it happens, feel free to open a PR for support.

Option `--binary-path` is available if you store the binary elsewhere.  
Option `--function-name` is available if you wish for the function to be named differently.
Default is `wd`, but if you pass this option with `eval`, the function will be
available as whatever you passed.

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
# cd to the base directory (no argument)
wd

# interactive selection of base directory, pass "/"
wd /

# Search for directories (under base_dir) containing "project"
wd project

# Change directory and open editor in it
wd -e workcd-go

# Specify editor for this session
wd --editor nvim workcd-go
wd --editor idea java-project

# Specify base directory for this session
wd --base-dir ~/Projects proj1

# define the tool to preview readme files of the directories
wd --preview-viewer glow /

## Helper commands

# print the config file to verify it's contents
wd --print-config
```

`--base-dir` enables you to do a neat trick of creating aliases that search different directories  
For example:

```bash
alias px="wd --base-dir ~/Workspace/projectx"
alias docs="wd --base-dir ~/Documents"
alias jcd="wd --editor idea --base-dir ~/Java-Projects"
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

## Why?

I became a DevOps for certain customer (yeah, DevOps is not necessarily a role, but you can imagine the kind of tasks one has)
and with that came a lot of projects I had to manage, cding into each one became cumbersome, and every time typing `nvim .` (or `v .`).
So I decided to create a small shell script that would do that for me(asked AI to generate it, because I dislike writing too much bash scripts).  
With time I wanted to also be able to use `fzf` to enable better interactive selection. And with time bash script became hard to manage.
So asked AI to port it to go, and then took over the development myself.
