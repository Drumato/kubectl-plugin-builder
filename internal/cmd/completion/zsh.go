package completion

import (
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

func newZshCommand(streams *genericclioptions.IOStreams) *cobra.Command {
	c := &cobra.Command{
		Use:   "zsh",
		Short: "generate zsh completion script",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Root().GenZshCompletion(streams.Out)
		},
	}
	return c
}
