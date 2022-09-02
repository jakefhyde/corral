package cmd_package

import (
	"fmt"
	"strings"

	pkgcmd "github.com/rancherlabs/corral/pkg/cmd"
	_package "github.com/rancherlabs/corral/pkg/package"
	"github.com/spf13/cobra"
)

var output = pkgcmd.OutputFormatTable

func NewCommandInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "info PACKAGE",
		Short: "Display details about the given package",
		Args:  cobra.ExactArgs(1),
		RunE:  info,
	}

	cmd.Flags().VarP(&output, "output", "o", "Output format. One of: table|json|yaml")

	return cmd
}

func info(_ *cobra.Command, args []string) error {
	pkg, err := _package.LoadPackage(args[0])
	if err != nil {
		return err
	}

	pkgVars := map[string]interface{}{}

	for k, v := range pkg.VariableSchemas {
		pkgVars[k] = v.Description
	}

	out, err := pkgcmd.Output(pkgVars, output, pkgcmd.OutputOptions{
		TitleFunc: func() string {
			return strings.Trim(fmt.Sprintf("%s\n%s", pkg.Name, pkg.Description), "\n")
		},
		HeaderFunc: func() []interface{} {
			return []interface{}{"VARIABLE", "DESCRIPTION"}
		},
	})
	if err != nil {
		return err
	}
	fmt.Println(out)
	return nil
}
