package cmd

import (
	"github.com/Drumato/kubectl-plugin-builder/internal/cmd/completion"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

func NewRootCommand(streams *genericclioptions.IOStreams) *cobra.Command {
	return NewCommand(
		&rootCommandOptions{},
		completion.NewRootCommand(streams),
		newNewCommand(),
		newAddCommand(),
		newGenerateCommand(),
	)
}

type rootCommandOptions struct {
	cmd *cobra.Command
}

func (rco *rootCommandOptions) Complete(cmd *cobra.Command, args []string) error {
	rco.cmd = cmd
	return nil
}

func (rco *rootCommandOptions) Validate() error {
	return nil
}

func (rco *rootCommandOptions) Run() error {
	return rco.cmd.Help()
}

func (rco *rootCommandOptions) CommandName() string {
	return "kubectl-plugin-builder"
}
