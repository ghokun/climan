package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/ghokun/climan/pkg/client"
	"github.com/ghokun/climan/pkg/platform"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

func init() {
	examples := `
# List all tools
climan list

# List versions of kubectl
climan list kubectl
`
	listCmd := &cobra.Command{
		Aliases: []string{"l", "li", "lis"},
		Args:    cobra.RangeArgs(0, 0),
		Example: examples,
		Use:     "list",
		Short:   "Lists all tools or versions of a tool",
		Long:    "Lists all tools or versions of a tool",
		RunE:    list,
	}

	rootCmd.AddCommand(listCmd)
}

func list(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return listAllTools()
	}
	return errors.Unwrap(fmt.Errorf("list command does not accept any arguments"))
}

func listAllTools() (err error) {
	tools, err := client.GetTools()
	if err != nil {
		return nil
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Latest", "Description"})
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding(" ")
	table.SetNoWhiteSpace(true)
	for _, value := range tools {
		if ok, _ := platform.CheckToolSupport(value.Supports); ok {
			table.Append([]string{value.Name, value.Latest, value.Description})
		}
	}
	table.Render()
	return nil
}
