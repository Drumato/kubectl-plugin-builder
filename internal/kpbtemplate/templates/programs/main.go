// Code generated by kubectl-plugin-builder.
package main

import (
        "fmt"
        "os"
        "{{ .PackageName }}/internal/cmd/{{ .RootCommandName }}"

        "k8s.io/cli-runtime/pkg/genericclioptions"
)

func main() {
        streams := genericclioptions.IOStreams{
                In:     os.Stdin,
                Out:    os.Stdout,
                ErrOut: os.Stderr,
        }

        if err := {{ .RootCommandName }}.NewCommand(&streams).Execute(); err != nil {
                fmt.Fprintf(os.Stderr, "ERROR: %+v\n", err)
                os.Exit(1)
        }
}
