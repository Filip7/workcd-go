package main

import (
	"flag"
)

type cliFlags struct {
	Execute       bool
	Editor        string
	BaseDir       string
	PreviewViewer string
}

var CmdFlags cliFlags

func setupFlags() {
	// Parse flags
	flag.BoolVar(&CmdFlags.Execute, "e", false, "Open editor after changing directory")
	flag.StringVar(&CmdFlags.Editor, "editor", "", "Editor to use (overrides config and $EDITOR)")
	flag.StringVar(&CmdFlags.BaseDir, "base-dir", "", "Base directory for workcd-go")
	flag.StringVar(&CmdFlags.PreviewViewer, "preview-viewer", "", "Tool to use for preview of markdown files")
	flag.Parse()
}
