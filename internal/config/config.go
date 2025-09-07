// Package config handles reading config from XDG_CONFIG_HOME or $HOME/.config
package config

import (
	"fmt"
	"os"
	"path/filepath"

	"go.yaml.in/yaml/v4"
)

// Config represents the configuration structure
type Config struct {
	BaseDir       string `yaml:"base_dir"`
	Editor        string `yaml:"editor"`
	PreviewViewer string `yaml:"preview_viewer"`
}

func ReadConfig() (*Config, error) {
	path := getConfigPathOrDefault()
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

	configDir := filepath.Join(configHome, "workcd")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating config directory: %v", err)
		os.Exit(1)
	}

	return filepath.Join(configDir, "config.yaml")
}
