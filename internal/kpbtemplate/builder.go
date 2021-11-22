package kpbtemplate

import (
	"embed"
)

//go:embed templates
var GlobalTemplates embed.FS

type Builder interface {
	Build() error
	Execute() error
}
