// Package command is responsible for building the command passed to the shell
package command

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/Filip7/workcd-go/internal/config"
	"github.com/Filip7/workcd-go/internal/flags"
)

func PrepareCommand(cmdFlags flags.CmdFlags) string {
	// Determine the fzf query
	fzfQuery := flags.GetCmdInput()

	// Determine the base directory from config or flag
	var err error

	// Merge config and flags under one config
	config := config.MergeConfig(cmdFlags)
	baseDir := config.BaseDir

	if fzfQuery == "" {
		fmt.Printf("cd %q", baseDir)
		return ""
	}

	if strings.HasSuffix(fzfQuery, "/") {
		// If the input contains '/' then query inside that directory instead of the baseDir
		baseDir += "/" + fzfQuery
	} else if strings.Contains(fzfQuery, "/") {
		// If the input does not end with '/' but contains it, open the sub dir instead of further querying
		dirs := strings.Split(fzfQuery, "/")
		baseDir += "/" + strings.Join(dirs[:len(dirs)-1], "/")
		fzfQuery = dirs[len(dirs)-1]
	}

	// Check if base directory exists
	if _, err := os.Stat(baseDir); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Base directory not found: %s\n", baseDir)
		os.Exit(1)
	}

	// List subdirectories
	dirs, err := listDirs(&baseDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	if len(dirs) == 0 {
		fmt.Fprintln(os.Stderr, "No directories found in base directory")
		os.Exit(1)
	}

	// Prepare fzf command
	previewStr := "if [ -f {}/README.md ]; then " + config.PreviewViewer + " {}/README.md; elif [ -f {}/README ]; then " + config.PreviewViewer + " {}/README; else echo \"No README found\"; fi"
	fzfCmd := exec.Command("fzf", "--query", fzfQuery, "--select-1", "--exit-0", "--preview", previewStr)
	fzfCmd.Stdin = strings.NewReader(strings.Join(dirs, "\n"))
	output, err := fzfCmd.Output()
	if err != nil {
		// fzf exits with non-zero status on error (with --exit-0, no match is exit 0)
		if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() != 1 {
			// fzf exited with non-zero status (likely user pressed Ctrl-C or Esc)
			os.Exit(exitErr.ExitCode())
		}
		// Most likely that fzf could not find the directory
		fmt.Fprintf(os.Stderr, "Error: could not find the directory, %v\n", err)
		os.Exit(1)
	}
	selectedDir := strings.TrimSpace(string(output))
	if selectedDir == "" {
		fmt.Fprintln(os.Stderr, "No directory selected.")
		os.Exit(1)
	}

	// Output the shell command to stdout
	cmd := fmt.Sprintf("cd %q", selectedDir)
	if cmdFlags.Execute {
		editor := config.Editor
		cmd += fmt.Sprintf(" && %s .", editor)
	}

	return cmd
}

func listDirs(dir *string) ([]string, error) {
	var dirs []string
	entries, err := os.ReadDir(*dir)
	if err != nil {
		return nil, err
	}
	for _, entry := range entries {
		if entry.IsDir() {
			fullPath := filepath.Join(*dir, entry.Name())
			dirs = append(dirs, fullPath)
		}
	}
	return dirs, nil
}
