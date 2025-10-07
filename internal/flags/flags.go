// Package flags does flag setup and returns CmdFlags struct containg values
package flags

import (
	"flag"
)

type CmdFlags struct {
	Execute       bool
	Editor        string
	BaseDir       string
	PreviewViewer string
	PrintConfig   bool
}

func SetupFlags() CmdFlags {
	var cmdFlags CmdFlags
	// Parse flags
	flag.BoolVar(&cmdFlags.Execute, "e", false, "Open editor after changing directory")
	flag.StringVar(&cmdFlags.Editor, "editor", "", "Editor to use (overrides config and $EDITOR)")
	flag.StringVar(&cmdFlags.BaseDir, "base-dir", "", "Base directory for workcd-go")
	flag.StringVar(&cmdFlags.PreviewViewer, "preview-viewer", "", "Tool to use for preview of markdown files")
	flag.BoolVar(&cmdFlags.PrintConfig, "print-config", false, "Print config currently in use")
	flag.Parse()

	return cmdFlags
}

func GetCmdInput() string {
	if flag.NArg() > 0 {
		return flag.Arg(0)
	}

	return ""
}
