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
	Shell         string
	BinaryPath    string
	FunctionName  string
}

func SetupFlags() CmdFlags {
	var cmdFlags CmdFlags
	// Parse flags
	flag.BoolVar(&cmdFlags.Execute, "e", false, "Open editor after changing directory")
	flag.StringVar(&cmdFlags.Editor, "editor", "", "Editor to use (overrides config and $EDITOR)")
	flag.StringVar(&cmdFlags.BaseDir, "base-dir", "", "Base directory for workcd-go")
	flag.StringVar(&cmdFlags.PreviewViewer, "preview-viewer", "", "Tool to use for preview of markdown files")
	flag.BoolVar(&cmdFlags.PrintConfig, "print-config", false, "Print config currently in use")
	flag.StringVar(&cmdFlags.Shell, "eval", "", "Get function used for shell integration, pass the shell that you use")
	flag.StringVar(&cmdFlags.BinaryPath, "binary-path", "/bin/workcd-go", "Define path to the binary of the tool, default to /bin/workcd-go")
	flag.StringVar(&cmdFlags.FunctionName, "function-name", "wd", "Define name of the function that will be passed to the shell, default is `wd`")
	flag.Parse()

	return cmdFlags
}

func GetCmdInput() string {
	if flag.NArg() > 0 {
		return flag.Arg(0)
	}

	return ""
}
