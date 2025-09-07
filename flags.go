package main

import (
	"flag"
)

type CmdFlags struct {
	Execute       bool
	Editor        string
	BaseDir       string
	PreviewViewer string
}

func setupFlags() CmdFlags {
	var cmdFlags CmdFlags
	// Parse flags
	flag.BoolVar(&cmdFlags.Execute, "e", false, "Open editor after changing directory")
	flag.StringVar(&cmdFlags.Editor, "editor", "", "Editor to use (overrides config and $EDITOR)")
	flag.StringVar(&cmdFlags.BaseDir, "base-dir", "", "Base directory for workcd-go")
	flag.StringVar(&cmdFlags.PreviewViewer, "preview-viewer", "", "Tool to use for preview of markdown files")
	flag.Parse()

	return cmdFlags
}
