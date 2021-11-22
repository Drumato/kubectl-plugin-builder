package gitignore

import (
	"os"
	"text/template"

	"github.com/Drumato/kubectl-plugin-builder/internal/kpbtemplate"
)

type GitIgnoreBuilder struct {
	data     *GitIgnoreData
	filePath string
	tmpl     *template.Template
}

var _ kpbtemplate.Builder = &GitIgnoreBuilder{}

func NewGitIgnoreBuilder(filePath string, data *GitIgnoreData) *GitIgnoreBuilder {
	return &GitIgnoreBuilder{
		filePath: filePath,
		data:     data,
	}
}

func (gib *GitIgnoreBuilder) Build() error {
	tmpl, err := template.ParseFS(kpbtemplate.GlobalTemplates, "templates/gitignore")
	if err != nil {
		return err
	}

	gib.tmpl = tmpl
	return nil
}

func (gib *GitIgnoreBuilder) Execute() error {
	f, err := os.Create(gib.filePath)
	if err != nil {
		return err
	}

	return gib.tmpl.Execute(f, gib.data)
}

type GitIgnoreData struct {
	ExecutableName string
}

func NewGitIgnoreData(exeName string) *GitIgnoreData {
	return &GitIgnoreData{
		ExecutableName: exeName,
	}
}
