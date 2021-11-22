package completion

import (
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

func newPowershellCommand(streams *genericclioptions.IOStreams) *cobra.Command {
	c := &cobra.Command{
		Use:   "powershell",
		Short: "generate powershell completion script",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Root().GenPowerShellCompletion(streams.Out)
		},
	}
	return c
}
