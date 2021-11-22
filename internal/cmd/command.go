package cmd

import "github.com/spf13/cobra"

func NewCommand(opts Options, children ...*cobra.Command) *cobra.Command {
	c := &cobra.Command{
		Use: opts.CommandName(),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.Complete(cmd, args); err != nil {
				return err
			}

			if err := opts.Validate(); err != nil {
				return err
			}

			return opts.Run()
		},
	}

	for _, child := range children {
		c.AddCommand(child)
	}

	return c
}
