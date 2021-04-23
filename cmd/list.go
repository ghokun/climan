package cmd

// TODO find local installed version
// TODO cache list all function
// TODO load from cache-rotate cache
// TODO fit results to screen
// TODO show no version found for empty list

import (
	"os"
	"sort"
	"strings"

	"github.com/ghokun/climan/cmd/tools"
	"github.com/ghokun/termdim"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var (
	forceRemote       bool
	includeUnreleased bool
)

func init() {
	examples := `
# List all tools
climan list

# List versions of kubectl
climan list kubectl

# List versions of oc
climan list oc

# List versions of helm but force remote call
climan list helm --force-remote

# List versions of kubectl including alpha, beta and release-candidates
climan list kubectl -u`

	validTools := make([]string, len(tools.Tools))
	i := 0
	for tool := range tools.Tools {
		validTools[i] = tool
		i++
	}

	listCmd := &cobra.Command{
		Aliases:   []string{"l", "li", "lis"},
		Args:      cobra.RangeArgs(0, 1),
		Example:   examples,
		Use:       "list",
		Short:     "Lists all tools or versions of a tool",
		Long:      "Lists all tools or versions of a tool",
		RunE:      list,
		ValidArgs: validTools,
	}

	listCmd.Flags().BoolVarP(&includeUnreleased, "include-unreleased", "u", false, "Includes alpha, beta and release candidates.")
	listCmd.Flags().BoolVarP(&forceRemote, "force-remote", "f", false, "Force remote checking of versions instead of local cache")
	rootCmd.AddCommand(listCmd)
}

func list(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return listAllTools()
	}
	err := cobra.OnlyValidArgs(cmd, args)
	if err != nil {
		return err
	}
	return listTool(args[0])
}

func listAllTools() error {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Tool", "Latest", "Current"})
	for _, tool := range tools.Tools {
		latest, err := tool.GetLatest()
		if err != nil {
			latest = tools.ToolVersion{Version: ""}
		}
		table.Append([]string{tool.Name, latest.Version, " "})
	}
	table.Render()
	return nil
}

func listTool(name string) error {
	tool := tools.Tools[name]
	all, err := tool.GetAll()
	if err != nil {
		return err
	}
	table := tablewriter.NewWriter(os.Stdout)
	if len(all) < 1 {
		table.SetHeader([]string{name})
		table.Append([]string{"No version found"})
	} else {
		width, height, err := termdim.GetSize(int(os.Stdout.Fd()))
		if err != nil {
			width = 80
			height = 24
		}
		sort.Sort(all)
		maxLength := len(name)
		var filtered []string
		for _, version := range all {
			if includeUnreleased || (!strings.Contains(version.Version, "alpha") && !strings.Contains(version.Version, "beta") && !strings.Contains(version.Version, "rc")) {
				if maxLength < len(version.Version) {
					maxLength = len(version.Version)
				}
				filtered = append(filtered, version.Version)
			}
		}
		if height > len(filtered) {
			table.SetHeader([]string{name})
			for _, row := range filtered {
				table.Append([]string{row})
			}
		} else {
			colNum := (width - 1) / (maxLength + 4)
			header := make([]string, colNum)
			header[0] = name
			for i := 1; i < colNum; i++ {
				header[i] = ""
			}
			table.SetHeader(header)
			rowNum := len(filtered) / colNum
			if rowNum*colNum < len(filtered) {
				rowNum += 1
			}
			rows := make([][]string, rowNum)
			for r := 0; r < rowNum; r++ {
				rows[r] = make([]string, colNum)
				for c := 0; c < colNum; c++ {
					if c*rowNum+r < len(filtered) {
						rows[r][c] = filtered[c*rowNum+r]
					}
				}
			}
			table.AppendBulk(rows)
		}
	}

	table.Render()
	return nil
}
