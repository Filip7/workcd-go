package integrations

import (
	"embed"
)

//go:embed templates/*
var TemplateScripts embed.FS
