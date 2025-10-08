package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Filip7/workcd-go/internal/flags"
	"go.yaml.in/yaml/v4"
)

// MergeConfig merges config file and flags passed by the user
// Idea is to have only one config as source of truth
// In case no config is set, nor flags passed, following defaults are set:
// editor: vi
// preview viewer: less
// default search dir: $HOME/Workspace
func MergeConfig(cmdFlags flags.CmdFlags) *Config {
	// First read the config file
	config, err := ReadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading config %v\n", err)
		os.Exit(1)
	}

	// Get the base dir
	if cmdFlags.BaseDir != "" {
		config.BaseDir = cmdFlags.BaseDir
	} else {
		config.BaseDir, err = getBaseDirFromConfig(config)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting base direcotry: %v\n", err)
			os.Exit(1)
		}
	}

	// Prepare the preview viewer
	if cmdFlags.PreviewViewer != "" {
		config.PreviewViewer = cmdFlags.PreviewViewer
	} else {
		config.PreviewViewer, err = getPreviewViewer(config)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	config.Editor = getEditor(config, &cmdFlags)

	return config
}

func PrintConfig(cmdFlags flags.CmdFlags) {
	config := MergeConfig(cmdFlags)

	yamlConfig, _ := yaml.Marshal(config)

	fmt.Println(string(yamlConfig))
}

func getBaseDirFromConfig(config *Config) (string, error) {
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

func getPreviewViewer(config *Config) (string, error) {
	// If base_dir is not set in config, use the default
	if config.PreviewViewer == "" {
		return "less", nil
	}

	return config.PreviewViewer, nil
}

func getEditor(config *Config, cmdFlags *flags.CmdFlags) string {
	if cmdFlags.Editor != "" {
		return cmdFlags.Editor
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
