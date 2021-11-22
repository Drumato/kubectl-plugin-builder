package module

import (
	"os"
	"text/template"

	"github.com/Drumato/kubectl-plugin-builder/internal/kpbtemplate"
)

const (
	defaultGoVersion         = "1.17"
	defaultCobraVersion      = "1.2.1"
	defaultCLIRuntimeVersion = "0.22.3"
)

type GoModuleBuilder struct {
	data     *GoModuleTemplateData
	filePath string
	tmpl     *template.Template
}

var _ kpbtemplate.Builder = &GoModuleBuilder{}

func NewGoModuleBuilder(filePath string, data *GoModuleTemplateData) *GoModuleBuilder {
	return &GoModuleBuilder{
		filePath: filePath,
		data:     data,
	}
}

func (gmb *GoModuleBuilder) Build() error {
	tmpl, err := template.ParseFS(kpbtemplate.GlobalTemplates, "templates/modules/go.template.mod")
	if err != nil {
		return err
	}

	gmb.tmpl = tmpl
	return nil
}

func (gmb *GoModuleBuilder) Execute() error {
	f, err := os.Create(gmb.filePath)
	if err != nil {
		return err
	}

	return gmb.tmpl.Execute(f, gmb.data)
}

type GoModuleTemplateData struct {
	PackageName string
	GoVersion   string
	Requires    []GoModuleRequire
}

type GoModuleRequire struct {
	Name    string
	Version string
}

func NewGoModuleTemplateData(pkgName string) *GoModuleTemplateData {
	return &GoModuleTemplateData{
		PackageName: pkgName,
		GoVersion:   defaultGoVersion,
		Requires: []GoModuleRequire{
			{Name: "github.com/spf13/cobra", Version: defaultCobraVersion},
			{Name: "k8s.io/cli-runtime", Version: defaultCLIRuntimeVersion},
		},
	}
}
