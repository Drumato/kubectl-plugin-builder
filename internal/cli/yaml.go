package cli

const (
	GoTypeString = "string"
	GoTypeInt    = "int"
	GoTypeUInt   = "uint"
)

type CLIYaml struct {
	PackageName string         `yaml:"packageName"`
	License     string         `yaml:"license"`
	Root        CLIYamlCommand `yaml:"root"`
}

type CLIYamlCommand struct {
	Name        string               `yaml:"name"`
	Year        uint                 `yaml:"year"`
	Author      string               `yaml:"author"`
	Description CLIYamlDescription   `yaml:"description"`
	Aliases     []string             `yaml:"aliases"`
	Flags       []CLIYamlCommandFlag `yaml:"flags"`
	Children    []CLIYamlCommand     `yaml:"children"`
	DefPath     string               `yaml:"defPath"`
}

type CLIYamlDescription struct {
	Short string `yaml:"short"`
	Long  string `yaml:"long"`
}

type CLIYamlCommandFlag struct {
	Name        string `yaml:"name"`
	Type        string `yaml:"type"`
	Description string `yaml:"description"`
}
