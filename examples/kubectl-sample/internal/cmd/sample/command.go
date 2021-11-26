// Code generated by kubectl-plugin-builder; DO NOT EDIT.

package sample

import (
        "github.com/spf13/cobra"

        "k8s.io/cli-runtime/pkg/genericclioptions"
        "github.com/Drumato/kubectl-sample/internal/cmd/sample/subcmd"
)

var (
)

// WARNING: don't rename this function.
func NewCommand(streams *genericclioptions.IOStreams) *cobra.Command {
        c := &cobra.Command{
                Use:  "sample",
                RunE: func(cmd *cobra.Command, args []string) error {
                        o := &options{streams: streams}
                        if err := o.Complete(cmd, args); err != nil {
                                return err
                        }

                        if err := o.Validate(); err != nil {
                                return err
                        }

                        return o.Run()
                },
        }
        c.AddCommand(subcmd.NewCommand(streams))

        return c
}