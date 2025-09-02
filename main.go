package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// Config represents the configuration structure
type Config struct {
	BaseDir string `yaml:"base_dir"`
	Editor  string `yaml:"editor"`
}

func main() {
	// Parse flags
	executeFlag := flag.Bool("e", false, "Open editor after changing directory")
	editorFlag := flag.String("editor", "", "Editor to use (overrides config and $EDITOR)")
	flag.Parse()

	// Determine the fzf query
	fzfQuery := "/"
	if flag.NArg() > 0 {
		fzfQuery = flag.Arg(0)
	}

	// Determine the base directory from config
	baseDir, err := getBaseDirFromConfig()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
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
		os.Exit(0)
	}

	// Prepare fzf command
	fzfCmd := exec.Command("fzf", "--query", fzfQuery, "--select-1", "--exit-0")
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
		return
	}

	// Output the shell command to stdout
	cmd := fmt.Sprintf("cd %q", selectedDir)
	if *executeFlag {
		editor := getEditor(*editorFlag)
		cmd += fmt.Sprintf(" && %s .", editor)
	}
	fmt.Println(cmd)
}

func getEditor(editorFlag string) string {
	if editorFlag != "" {
		return editorFlag
	}

	// Try to read editor from config
	config, err := readConfig(getConfigPathOrDefault())
	if err == nil && config.Editor != "" {
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

func getBaseDirFromConfig() (string, error) {
	config, err := readConfig(getConfigPathOrDefault())
	if err != nil && !os.IsNotExist(err) {
		return "", err
	}

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

func getConfigPathOrDefault() string {
	configHome := os.Getenv("XDG_CONFIG_HOME")
	if configHome == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting user home directory: %v", err)
			os.Exit(1)
		}
		configHome = filepath.Join(home, ".config")
	}

	configDir := filepath.Join(configHome, "workcd-go")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating config directory: %v", err)
		os.Exit(1)
	}

	return filepath.Join(configDir, "config.yaml")
}

func readConfig(path string) (*Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return &Config{}, err
	}

	var config Config
	if err := yaml.Unmarshal(file, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
