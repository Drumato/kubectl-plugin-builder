// Code generated by kubectl-plugin-builder; DO NOT EDIT.
{{ .SourceHeaderLicense }}
package {{ .CommandName }}

import (
        "github.com/spf13/cobra"

        "k8s.io/cli-runtime/pkg/genericclioptions"
        {{- range .Children }}
        "{{ $.PackageName }}/{{ .DefPath }}"
        {{- end }}
)

// WARNING: don't rename this function.
func NewCommand(streams *genericclioptions.IOStreams) *cobra.Command {
        c := &cobra.Command{
                Use:  "{{ .CommandName }}",
                {{- if ne .Short ""}}
                Short: "{{ .Short }}",
                {{- end }}
                {{- if ne .Long ""}}
                Long: "{{ .Long }}",
                {{- end }}
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

        {{- range .Children }}
        c.AddCommand({{ .Name }}.NewCommand(streams))
        {{- end }}

        return c
}