package cmd

import "github.com/spf13/cobra"

type Options interface {
	CommandName() string
	Complete(cmd *cobra.Command, args []string) error
	Validate() error
	Run() error
}
