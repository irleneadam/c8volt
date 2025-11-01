package cmd

import (
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:     "config",
	Short:   "Manage configuration settings",
	Aliases: []string{"cfg"},
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
	SuggestFor: []string{"confige", "exepct"},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
