package completion

import (
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

func newFishCommand(streams *genericclioptions.IOStreams) *cobra.Command {
	c := &cobra.Command{
		Use:   "fish",
		Short: "generate fish completion script",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Root().GenFishCompletion(streams.Out, true)
		},
	}
	return c
}
