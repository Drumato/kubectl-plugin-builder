package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Drumato/kubectl-plugin-builder/internal/kpbtemplate"
	"github.com/Drumato/kubectl-plugin-builder/internal/kpbtemplate/cli"
	"github.com/Drumato/kubectl-plugin-builder/internal/kpbtemplate/command"
	"github.com/Drumato/kubectl-plugin-builder/internal/kpbtemplate/gitignore"
	"github.com/Drumato/kubectl-plugin-builder/internal/kpbtemplate/handler"
	"github.com/Drumato/kubectl-plugin-builder/internal/kpbtemplate/license"
	"github.com/Drumato/kubectl-plugin-builder/internal/kpbtemplate/makefile"
	"github.com/Drumato/kubectl-plugin-builder/internal/kpbtemplate/module"
	"github.com/Drumato/kubectl-plugin-builder/internal/kpbtemplate/node"
	"github.com/spf13/cobra"
)

const (
	licenseMode uint = iota
	goModuleMode
	gitIgnoreMode
	makeFileMode
	cliYamlMode
	cliNodeOptionsMode
	rootCommandPkgMode
	rootCommandHandlerPkgMode
	mainPackageMode
)

var (
	newLicenseFlag string
	newAuthorFlag  string
	newYearFlag    uint
)

type newCommandOptions struct {
	args            []string
	license         license.LicenseBuilder
	pkgName         string
	pluginName      string
	rootCommandName string
}

var _ Options = &newCommandOptions{}

func newNewCommand() *cobra.Command {
	c := NewCommand(&newCommandOptions{})
	c.Flags().StringVarP(&newLicenseFlag, "license", "l", license.MITLicense, "project license")
	c.Flags().StringVarP(&newAuthorFlag, "author", "a", "you", "project author")
	c.Flags().UintVar(&newYearFlag, "year", uint(time.Now().Year()), "copyright year")
	return c
}

func (nco *newCommandOptions) Complete(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("you must specify a package name")
	}
	nco.args = args
	nco.pkgName = nco.args[0]

	// `github.com/Drumato/kubectl-sample` -> `kubectl-sample`
	pluginName := filepath.Base(nco.pkgName)
	nco.pluginName = pluginName
	nco.rootCommandName = strings.TrimPrefix(pluginName, "kubectl-")

	return nil
}

func (nco *newCommandOptions) Validate() error {
	if nco.hasInvalidLisence(newLicenseFlag) {
		return fmt.Errorf("unsupported license '%s' detect", newLicenseFlag)
	}

	return nil
}

func (nco *newCommandOptions) Run() error {
	if nco.ProjectAlreadyExists() {
		fmt.Printf("Project %s alredy exists.", nco.pkgName)
		return nil
	}

	// nco.license will be used after generating LICENSE file.
	// so the assignment was placed here. (not generateLicense())
	l, err := license.ChooseLicenseBuilder(newLicenseFlag, newYearFlag, newAuthorFlag)
	if err != nil {
		return nil
	}
	nco.license = l

	if err := nco.makePackageDirs(); err != nil {
		return err
	}

	if err := nco.generateFiles(); err != nil {
		return err
	}

	fmt.Println("Initialization Complete!")
	fmt.Println("Run `go mod tidy` to install third-party modules.")
	return nil
}

func (nco *newCommandOptions) makePackageDirs() error {
	if err := os.MkdirAll("internal/cmd", 0o755); err != nil {
		return err
	}

	rootCmdDir := fmt.Sprintf("internal/cmd/%s", nco.rootCommandName)
	if err := os.MkdirAll(rootCmdDir, 0o755); err != nil {
		return err
	}

	mainDir := fmt.Sprintf("cmd/%s", nco.pluginName)
	if err := os.MkdirAll(mainDir, 0o755); err != nil {
		return err
	}

	return nil
}

func (nco *newCommandOptions) generateFiles() error {
	modes := []uint{
		licenseMode,
		goModuleMode,
		gitIgnoreMode,
		makeFileMode,
		cliYamlMode,
		cliNodeOptionsMode,
		rootCommandPkgMode,
		rootCommandHandlerPkgMode,
		cliNodeOptionsMode,
		mainPackageMode,
	}

	for _, mode := range modes {
		builder := nco.chooseBuilder(mode)
		if err := builder.Build(); err != nil {
			return err
		}

		if err := builder.Execute(); err != nil {
			return err
		}
	}

	return nil
}

func (nco *newCommandOptions) ProjectAlreadyExists() bool {
	files := []string{
		"go.mod", "Makefile", ".gitignore", "cli.yaml",
		"internal", "cmd",
	}

	for _, file := range files {
		if _, err := os.Stat(file); err == nil {
			return true
		}
	}

	return false
}

func (nco *newCommandOptions) hasInvalidLisence(licenseName string) bool {
	switch licenseName {
	case license.MITLicense:
		return false
	default:
		return true
	}
}

func (nco *newCommandOptions) CommandName() string {
	return "new"
}

func (nco *newCommandOptions) chooseBuilder(mode uint) kpbtemplate.Builder {
	switch mode {
	case licenseMode:
		return nco.license
	case goModuleMode:
		data := &module.GoModuleTemplateData{
			PackageName: nco.pkgName,
			GoVersion:   module.DefaultGoVersion,
			Requires:    module.DefaultRequires,
		}
		return module.NewGoModuleBuilder("go.mod", data)
	case gitIgnoreMode:
		data := &gitignore.GitIgnoreData{
			ExecutableName: nco.pluginName,
		}
		return gitignore.NewGitIgnoreBuilder(".gitignore", data)
	case makeFileMode:
		data := &makefile.MakefileData{
			PluginName: filepath.Base(nco.pkgName),
		}
		return makefile.NewMakefileBuilder("Makefile", data)
	case rootCommandPkgMode:
		data := &command.CommandData{
			SourceHeaderLicense: nco.license.SourceFileHeader(),
			CommandName:         nco.rootCommandName,
			PackageName:         nco.pkgName,
			Short:               "short description",
			Long:                "long description",
			Children:            []command.CommandDataChildren{},
		}

		path := fmt.Sprintf("internal/cmd/%s/command.go", nco.rootCommandName)
		return command.NewCommandBuilder(path, data)
	case rootCommandHandlerPkgMode:
		data := &handler.HandlerData{
			SourceHeaderLicense: nco.license.SourceFileHeader(),
			CommandName:         nco.rootCommandName,
			PackageName:         nco.pkgName,
		}
		path := fmt.Sprintf("internal/cmd/%s/handler.go", nco.rootCommandName)
		return handler.NewHandlerBuilder(path, data)
	case cliNodeOptionsMode:
		data := &node.NodeData{
			SourceHeaderLicense: nco.license.SourceFileHeader(),
		}
		path := "internal/cmd/node.go"
		return node.NewNodeBuilder(path, data)
	case cliYamlMode:
		data := &cli.CLIYamlData{
			RootCommandName: nco.rootCommandName,
			Author:          newAuthorFlag,
			Year:            newYearFlag,
			License:         newLicenseFlag,
			PackageName:     nco.pkgName,
		}
		return cli.NewCLIYamlBuilder("cli.yaml", data)
	case mainPackageMode:
		data := &cli.EntrypointData{
			PluginName:      nco.pluginName,
			PackageName:     nco.pkgName,
			RootCommandName: nco.rootCommandName,
		}
		path := fmt.Sprintf("cmd/%s/main.go", nco.pluginName)
		return cli.NewEntrypointBuilder(path, data)
	default:
		panic("unreachable")
	}
}
