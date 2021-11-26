package cli

import (
	"os"
	"text/template"

	"github.com/Drumato/kubectl-plugin-builder/internal/kpbtemplate"
)

type CLIYamlBuilder struct {
	data     *CLIYamlData
	filePath string
	tmpl     *template.Template
}

var _ kpbtemplate.Builder = &CLIYamlBuilder{}

func NewCLIYamlBuilder(filePath string, data *CLIYamlData) *CLIYamlBuilder {
	return &CLIYamlBuilder{
		filePath: filePath,
		data:     data,
	}
}

func (cyb *CLIYamlBuilder) Build() error {
	tmpl, err := template.ParseFS(kpbtemplate.GlobalTemplates, "templates/cli.yaml")
	if err != nil {
		return err
	}

	cyb.tmpl = tmpl
	return nil
}

func (cyb *CLIYamlBuilder) Execute() error {
	f, err := os.Create(cyb.filePath)
	if err != nil {
		return err
	}

	return cyb.tmpl.Execute(f, cyb.data)
}

type CLIYamlData struct {
	RootCommandName string
	Author          string
	Year            uint
	License         string
	PackageName     string
}
