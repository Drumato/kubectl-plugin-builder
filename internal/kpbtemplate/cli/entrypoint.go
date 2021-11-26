package cli

import (
	"os"
	"text/template"

	"github.com/Drumato/kubectl-plugin-builder/internal/kpbtemplate"
)

type EntrypointBuilder struct {
	data     *EntrypointData
	filePath string
	tmpl     *template.Template
}

var _ kpbtemplate.Builder = &EntrypointBuilder{}

func NewEntrypointBuilder(filePath string, data *EntrypointData) *EntrypointBuilder {
	return &EntrypointBuilder{
		filePath: filePath,
		data:     data,
	}
}

func (eb *EntrypointBuilder) Build() error {
	tmpl, err := template.ParseFS(kpbtemplate.GlobalTemplates, "templates/programs/main.go")
	if err != nil {
		return err
	}

	eb.tmpl = tmpl
	return nil
}

func (eb *EntrypointBuilder) Execute() error {
	f, err := os.Create(eb.filePath)
	if err != nil {
		return err
	}

	return eb.tmpl.Execute(f, eb.data)
}

type EntrypointData struct {
	PackageName     string
	PluginName      string
	RootCommandName string
}
