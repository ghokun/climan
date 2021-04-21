package cmd

import (
	"github.com/spf13/cobra"
)

var (
	checksum string
)

func init() {
	installCmd := &cobra.Command{
		Aliases: []string{"i", "ins"},
		Use:     "install",
		Short:   "Installs a tool",
		Long:    "Installs a tool",
		RunE:    install,
	}
	installCmd.Flags().StringVarP(&checksum, "checksum", "c", "", "Overrides remote checksum value")
	rootCmd.AddCommand(installCmd)
}

func install(cmd *cobra.Command, args []string) error {

	return nil
}
