package makefile

import (
	"os"
	"text/template"

	"github.com/Drumato/kubectl-plugin-builder/internal/kpbtemplate"
)

type MakefileBuilder struct {
	data     *MakefileData
	filePath string
	tmpl     *template.Template
}

var _ kpbtemplate.Builder = &MakefileBuilder{}

func NewMakefileBuilder(filePath string, data *MakefileData) *MakefileBuilder {
	return &MakefileBuilder{
		filePath: filePath,
		data:     data,
	}
}

func (mfb *MakefileBuilder) Build() error {
	tmpl, err := template.ParseFS(kpbtemplate.GlobalTemplates, "templates/Makefile")
	if err != nil {
		return err
	}

	mfb.tmpl = tmpl
	return nil
}

func (mfb *MakefileBuilder) Execute() error {
	f, err := os.Create(mfb.filePath)
	if err != nil {
		return err
	}

	return mfb.tmpl.Execute(f, mfb.data)
}

type MakefileData struct {
	PluginName string
}

func NewMakefileData(pluginName string) *MakefileData {
	return &MakefileData{
		PluginName: pluginName,
	}
}
