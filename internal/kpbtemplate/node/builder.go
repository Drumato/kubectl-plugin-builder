package node

import (
	"os"
	"text/template"

	"github.com/Drumato/kubectl-plugin-builder/internal/kpbtemplate"
)

type NodeBuilder struct {
	data     *NodeData
	filePath string
	tmpl     *template.Template
}

var _ kpbtemplate.Builder = &NodeBuilder{}

func NewNodeBuilder(filePath string, data *NodeData) *NodeBuilder {
	return &NodeBuilder{
		filePath: filePath,
		data:     data,
	}
}

func (nb *NodeBuilder) Build() error {
	tmpl, err := template.ParseFS(kpbtemplate.GlobalTemplates, "templates/programs/node.go")
	if err != nil {
		return err
	}

	nb.tmpl = tmpl
	return nil
}

func (nb *NodeBuilder) Execute() error {
	f, err := os.Create(nb.filePath)
	if err != nil {
		return err
	}

	return nb.tmpl.Execute(f, nb.data)
}

type NodeData struct {
	SourceHeaderLicense string
}
