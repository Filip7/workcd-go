package main

import (
	"fmt"
	"os"

	"github.com/Filip7/workcd-go/internal/command"
	"github.com/Filip7/workcd-go/internal/config"
	"github.com/Filip7/workcd-go/internal/flags"
)

func main() {
	cmdFlags := flags.SetupFlags()
	if cmdFlags.PrintConfig {
		config.PrintConfig(cmdFlags)
		os.Exit(0)
	}

	command := command.PrepareCommand(cmdFlags)
	fmt.Println(command)
}
