package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/Drumato/kubectl-plugin-builder/internal/cli"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var (
	addAuthorFlag string
	addShortFlag  string
	addLongFlag   string
)

type addCommandOptions struct {
	args          []string
	cmdName       string
	cmdParentName string
	yamlTree      cli.CLIYaml
}

var _ Options = &addCommandOptions{}

func newAddCommand() *cobra.Command {
	c := NewCommand(&addCommandOptions{})
	c.Flags().StringVarP(&addAuthorFlag, "author", "a", "you", "new command's author")
	c.Flags().StringVarP(&addShortFlag, "short", "s", "short description", "new command's short description")
	c.Flags().StringVarP(&addLongFlag, "long", "l", "this is a long description", "new command's long description")
	return c
}

func (aco *addCommandOptions) Complete(cmd *cobra.Command, args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("usage: kubectl-plugin-builder CMD_NAME CMD_PARENT")
	}
	aco.args = args
	aco.cmdName = aco.args[0]
	aco.cmdParentName = aco.args[1]

	f, err := os.Open("cli.yaml")
	if err != nil {
		return err
	}
	yamlTree := cli.CLIYaml{}
	if err := yaml.NewDecoder(f).Decode(&yamlTree); err != nil {
		return err
	}

	aco.yamlTree = yamlTree

	return nil
}

func (aco *addCommandOptions) Validate() error {
	if !aco.commandNameExistsOnTree(aco.cmdParentName, aco.yamlTree.Root) {
		return fmt.Errorf("CMD_PARENT must exist on the cli.yaml")
	}

	return nil
}

func (aco *addCommandOptions) Run() error {
	if err := aco.hangNewCommandOnParent(); err != nil {
		return err
	}

	f, err := os.Create("cli.yaml")
	if err != nil {
		return err
	}

	if err := yaml.NewEncoder(f).Encode(&aco.yamlTree); err != nil {
		return err
	}

	return nil
}

func (aco *addCommandOptions) hangNewCommandOnParent() error {
	parentCmd := aco.getParentCmdReferenceFromTree()

	for _, child := range parentCmd.Children {
		if child.Name == aco.cmdName {
			return fmt.Errorf("the name '%s' already used in a child of %s cmd", aco.cmdName, aco.cmdParentName)
		}
	}

	newCmd := cli.CLIYamlCommand{
		Name:    aco.cmdName,
		Year:    uint(time.Now().Year()),
		Author:  addAuthorFlag,
		DefPath: fmt.Sprintf("%s/%s", parentCmd.DefPath, aco.cmdName),
		Description: cli.CLIYamlDescription{
			Short: addShortFlag,
			Long:  addLongFlag,
		},
	}
	parentCmd.Children = append(parentCmd.Children, newCmd)
	return nil
}

// this function assumes the parent command must exist on the tree.
// see `aco.Validate()`.
func (aco *addCommandOptions) getParentCmdReferenceFromTree() *cli.CLIYamlCommand {
	cmdQueue := []*cli.CLIYamlCommand{&aco.yamlTree.Root}

	for {
		curCmd := cmdQueue[0]
		cmdQueue = cmdQueue[1:]
		if curCmd.Name == aco.cmdParentName {
			return curCmd
		}

		for _, child := range curCmd.Children {
			cmdQueue = append(cmdQueue, &child)
		}
	}
}
func (aco *addCommandOptions) commandNameExistsOnTree(name string, cmd cli.CLIYamlCommand) bool {
	if cmd.Name == name {
		return true
	}

	for _, child := range cmd.Children {
		if aco.commandNameExistsOnTree(name, child) {
			return true
		}
	}

	return false
}

func (aco *addCommandOptions) CommandName() string {
	return "add"
}
