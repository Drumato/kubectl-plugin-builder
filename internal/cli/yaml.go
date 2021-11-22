package cli

type CLIYaml struct {
	Root        CLIYamlCommand `yaml:"root"`
	License     string         `yaml:"license"`
	PackageName string         `yaml:"packageName"`
}

type CLIYamlCommand struct {
	Name        string                   `yaml:"name"`
	Year        uint                     `yaml:"year"`
	Author      string                   `yaml:"author"`
	Description CLIYamlDescription       `yaml:"description"`
	Aliases     []string                 `yaml:"aliases"`
	Arguments   []CLIYamlCommandArgument `yaml:"args"`
	Flags       []CLIYamlCommandFlag     `yaml:"flags"`
	Children    []CLIYamlCommand         `yaml:"children"`
	DefPath     string                   `yaml:"defPath"`
}

type CLIYamlDescription struct {
	Short string `yaml:"short"`
	Long  string `yaml:"long"`
}

type CLIYamlCommandArgument struct {
	Name  string `yaml:"name"`
	Type  string `yaml:"type"`
	Index uint   `yaml:"index"`
}

type CLIYamlCommandFlag struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
}
