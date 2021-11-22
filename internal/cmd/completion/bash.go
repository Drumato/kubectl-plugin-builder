package completion

import (
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

func newBashCommand(streams *genericclioptions.IOStreams) *cobra.Command {
	c := &cobra.Command{
		Use:   "bash",
		Short: "generate bash completion script",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Root().GenBashCompletion(streams.Out)
		},
	}
	return c
}
