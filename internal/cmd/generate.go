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

	data := &command.CommandData{
		Short:       cliCmd.Description.Short,
		Long:        cliCmd.Description.Long,
		CommandName: cliCmd.Name,
		PackageName: gco.expected.PackageName,
		Children:    make([]command.CommandDataChildren, 0),
	}

	for _, cliChild := range cliCmd.Children {
		data.Children = append(data.Children, command.CommandDataChildren{
			Name:    cliChild.Name,
			DefPath: cliChild.DefPath,
		})
	}
	builder := command.NewCommandBuilder(path, data)

	if err := builder.Build(); err != nil {
		return err
	}

	return builder.Execute()
}

func (gco *generateCommandOptions) generateHandlerGo(command cli.CLIYamlCommand) error {
	path := fmt.Sprintf("%s/handler.go", command.DefPath)
	if _, err := os.Stat(path); err == nil {
		return nil
	}

	data := &handler.HandlerData{
		PackageName: gco.expected.PackageName,
		CommandName: command.Name,
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
