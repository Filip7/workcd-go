package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	setupFlags()

	// Determine the fzf query
	var fzfQuery string
	if flag.NArg() > 0 {
		fzfQuery = flag.Arg(0)
	}

	// Determine the base directory from config or flag
	var baseDir string
	var err error

	config, err := readConfig(getConfigPathOrDefault())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading config: %v\n", err)
		os.Exit(1)
	}

	if CmdFlags.BaseDir != "" {
		baseDir = CmdFlags.BaseDir
	} else {
		baseDir, err = getBaseDirFromConfig(*config)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting base direcotry: %v\n", err)
			os.Exit(1)
		}
	}

	if fzfQuery == "" {
		fmt.Printf("cd %q\n", baseDir)
		return
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
	dirs, err := listDirs(baseDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	if len(dirs) == 0 {
		fmt.Fprintln(os.Stderr, "No directories found in base directory")
		os.Exit(1)
	}

	// Prepare the preview viewer
	var preview string
	if CmdFlags.PreviewViewer != "" {
		preview = CmdFlags.PreviewViewer
	} else {
		preview, err = getPreviewViewer(*config)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	// Prepare fzf command
	fzfCmd := exec.Command("fzf", "--query", fzfQuery, "--select-1", "--exit-0", "--preview", "if [ -f {}/README.md ]; then "+preview+" {}/README.md; else echo \"No README found\"; fi")
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
	if CmdFlags.Execute {
		editor := getEditor(*config)
		cmd += fmt.Sprintf(" && %s .", editor)
	}

	fmt.Println(cmd)
}

func getEditor(config Config) string {
	if CmdFlags.Editor != "" {
		return CmdFlags.Editor
	}

	if config.Editor != "" {
		return config.Editor
	}

	// Fall back to $EDITOR environment variable
	editor := os.Getenv("EDITOR")
	if editor != "" {
		return editor
	}

	// Final fallback
	return "vi"
}

func getBaseDirFromConfig(config Config) (string, error) {
	// If base_dir is not set in config, use the default
	if config.BaseDir == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(home, "Workspace"), nil
	}

	return config.BaseDir, nil
}

func getPreviewViewer(config Config) (string, error) {
	// If base_dir is not set in config, use the default
	if config.PreviewViewer == "" {
		return "less", nil
	}

	return config.PreviewViewer, nil
}

func listDirs(dir string) ([]string, error) {
	var dirs []string
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, entry := range entries {
		if entry.IsDir() {
			fullPath := filepath.Join(dir, entry.Name())
			dirs = append(dirs, fullPath)
		}
	}
	return dirs, nil
}
