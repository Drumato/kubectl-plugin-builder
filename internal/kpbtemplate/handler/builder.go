package handler

import (
	"os"
	"text/template"

	"github.com/Drumato/kubectl-plugin-builder/internal/kpbtemplate"
)

type HandlerBuilder struct {
	data     *HandlerData
	filePath string
	tmpl     *template.Template
}

var _ kpbtemplate.Builder = &HandlerBuilder{}

func NewHandlerBuilder(filePath string, data *HandlerData) *HandlerBuilder {
	return &HandlerBuilder{
		filePath: filePath,
		data:     data,
	}
}

func (hb *HandlerBuilder) Build() error {
	tmpl, err := template.ParseFS(kpbtemplate.GlobalTemplates, "templates/programs/handler.go")
	if err != nil {
		return err
	}

	hb.tmpl = tmpl
	return nil
}

func (hb *HandlerBuilder) Execute() error {
	f, err := os.Create(hb.filePath)
	if err != nil {
		return err
	}

	return hb.tmpl.Execute(f, hb.data)
}

type HandlerData struct {
	SourceHeaderLicense string
	CommandName         string
	PackageName         string
}
