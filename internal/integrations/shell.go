// Package integrations handles integration part of the program with various shells
package integrations

import (
	"bytes"
	"fmt"
	"os"
	"text/template"

	"github.com/Filip7/workcd-go/internal/flags"
)

func GetShellIntegration(shell string, cmdFlags flags.CmdFlags) string {
	var tmplPath string
	switch shell {
	case "zsh":
		tmplPath = "templates/script.zsh.tmpl"
	default:
		fmt.Fprintf(os.Stderr, "Shell %s is currently not supported!\n", shell)
		os.Exit(1)
	}

	tmpl, err := template.ParseFS(TemplateScripts, tmplPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error opening the template ", err)
		os.Exit(1)
	}
	var tpl bytes.Buffer
	err = tmpl.Execute(&tpl, struct{ BinaryPath string }{BinaryPath: cmdFlags.BinaryPath})
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error opening the template ", err)
		os.Exit(1)
	}

	return tpl.String()
}
