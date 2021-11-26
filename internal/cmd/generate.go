package cmd

import (
	"fmt"
	"os"

	"github.com/Drumato/kubectl-plugin-builder/internal/cli"
	"github.com/Drumato/kubectl-plugin-builder/internal/kpbtemplate/command"
	"github.com/Drumato/kubectl-plugin-builder/internal/kpbtemplate/handler"
	"github.com/Drumato/kubectl-plugin-builder/internal/kpbtemplate/license"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type generateCommandOptions struct {
	args     []string
	expected cli.CLIYaml
	license  license.LicenseBuilder
}

var _ Options = &generateCommandOptions{}

func newGenerateCommand() *cobra.Command {
	c := NewCommand(&generateCommandOptions{})
	return c
}

func (gco *generateCommandOptions) Complete(cmd *cobra.Command, args []string) error {
	gco.args = args

	f, err := os.Open("cli.yaml")
	if err != nil {
		return err
	}

	expected := cli.CLIYaml{}
	if err := yaml.NewDecoder(f).Decode(&expected); err != nil {
		return err
	}

	gco.expected = expected

	lb, err := license.ChooseLicenseBuilder(
		gco.expected.License,
		gco.expected.Root.Year,
		gco.expected.Root.Author)
	if err != nil {
		return err
	}
	gco.license = lb

	return nil
}

func (gco *generateCommandOptions) Validate() error {
	return nil
}

func (gco *generateCommandOptions) Run() error {
	rootCLICmd := gco.expected.Root
	gco.generateCommandPackage(rootCLICmd)

	return nil
}

func (gco *generateCommandOptions) generateCommandPackage(command cli.CLIYamlCommand) error {
	if err := os.MkdirAll(command.DefPath, 0o755); err != nil {
		return err
	}

	if err := gco.generateCommandGo(command); err != nil {
		return err
	}
	if err := gco.generateHandlerGo(command); err != nil {
		return err
	}

	for _, child := range command.Children {
		gco.generateCommandPackage(child)
	}

	return nil
}

func (gco *generateCommandOptions) generateCommandGo(cliCmd cli.CLIYamlCommand) error {
	path := fmt.Sprintf("%s/command.go", cliCmd.DefPath)
	if _, err := os.Stat(path); err == nil {
		if err := os.Remove(path); err != nil {
			return err
		}
	}

	children := gco.generateCommandChildrenFromYAML(cliCmd)
	flags := gco.generateCommandFlagsFromYAML(cliCmd)
	aliases := gco.generateCommandAliasesFromYAML(cliCmd)
	data := &command.CommandData{
		SourceHeaderLicense: gco.license.SourceFileHeader(),
		Short:               cliCmd.Description.Short,
		Long:                cliCmd.Description.Long,
		CommandName:         cliCmd.Name,
		PackageName:         gco.expected.PackageName,
		Children:            children,
		Flags:               flags,
		Aliases:             aliases,
	}
	builder := command.NewCommandBuilder(path, data)

	if err := builder.Build(); err != nil {
		return err
	}

	return builder.Execute()
}

func (gco *generateCommandOptions) generateHandlerGo(command cli.CLIYamlCommand) error {
	// we mustn't replace the existing handler.go
	// because the file may be overwritten by the user.
	path := fmt.Sprintf("%s/handler.go", command.DefPath)
	if _, err := os.Stat(path); err == nil {
		return nil
	}

	data := &handler.HandlerData{
		SourceHeaderLicense: gco.license.SourceFileHeader(),
		PackageName:         gco.expected.PackageName,
		CommandName:         command.Name,
	}
	builder := handler.NewHandlerBuilder(path, data)

	if err := builder.Build(); err != nil {
		return err
	}

	return builder.Execute()
}

func (gco *generateCommandOptions) CommandName() string {
	return "generate"
}

func (gco *generateCommandOptions) generateCommandAliasesFromYAML(cliCmd cli.CLIYamlCommand) []command.CommandAlias {
	aliases := make([]command.CommandAlias, len(cliCmd.Aliases))

	for i, cliAlias := range cliCmd.Aliases {
		aliases[i] = command.CommandAlias{
			Name: cliAlias,
		}
	}

	return aliases
}

func (gco *generateCommandOptions) generateCommandChildrenFromYAML(cliCmd cli.CLIYamlCommand) []command.CommandDataChildren {
	children := make([]command.CommandDataChildren, len(cliCmd.Children))

	for i, cliChild := range cliCmd.Children {
		children[i] = command.CommandDataChildren{
			Name:    cliChild.Name,
			DefPath: cliChild.DefPath,
		}
	}

	return children
}

func (gco *generateCommandOptions) generateCommandFlagsFromYAML(cliCmd cli.CLIYamlCommand) []command.CommandFlag {
	flags := make([]command.CommandFlag, len(cliCmd.Flags))

	for i, cliFlag := range cliCmd.Flags {
		upperType, defaultValue := gco.genTVFromGoTypeString(cliFlag.Type)
		flags[i] = command.CommandFlag{
			Name:         cliFlag.Name,
			UpperType:    upperType,
			Type:         cliFlag.Type,
			DefaultValue: defaultValue,
			Description:  cliFlag.Description,
		}
	}

	return flags
}

func (gco *generateCommandOptions) genTVFromGoTypeString(goType string) (string, string) {
	switch goType {
	case cli.GoTypeString:
		return "String", "\"\""
	case cli.GoTypeInt:
		return "Int", "0"
	case cli.GoTypeUInt:
		return "Uint", "0"
	default:
		panic("unreachable!")
	}
}
