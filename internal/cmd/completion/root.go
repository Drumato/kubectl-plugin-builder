package completion

import (
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

func NewRootCommand(streams *genericclioptions.IOStreams) *cobra.Command {
	c := &cobra.Command{}

	c.AddCommand(newBashCommand(streams))
	c.AddCommand(newZshCommand(streams))
	c.AddCommand(newFishCommand(streams))
	c.AddCommand(newPowershellCommand(streams))
	return c
}
