package command

import (
	"os"
	"text/template"

	"github.com/Drumato/kubectl-plugin-builder/internal/kpbtemplate"
)

type CommandBuilder struct {
	data     *CommandData
	filePath string
	tmpl     *template.Template
}

var _ kpbtemplate.Builder = &CommandBuilder{}

func NewCommandBuilder(filePath string, data *CommandData) *CommandBuilder {
	return &CommandBuilder{
		filePath: filePath,
		data:     data,
	}
}

func (cb *CommandBuilder) Build() error {
	tmpl, err := template.ParseFS(kpbtemplate.GlobalTemplates, "templates/programs/command.go")
	if err != nil {
		return err
	}

	cb.tmpl = tmpl
	return nil
}

func (cb *CommandBuilder) Execute() error {
	f, err := os.Create(cb.filePath)
	if err != nil {
		return err
	}

	return cb.tmpl.Execute(f, cb.data)
}

type CommandData struct {
	SourceHeaderLicense string
	CommandName         string
	PackageName         string
	Short               string
	Long                string
	Children            []CommandDataChildren
}

type CommandDataChildren struct {
	Name    string
	DefPath string
}
