package cmd

import (
	"github.com/spf13/cobra"
)

var embedCmd = &cobra.Command{
	Use:     "embed",
	Short:   "Manage embedded resources",
	Aliases: []string{"em", "emb"},
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
	SuggestFor: []string{"embedd", "embd", "embedded", "embeded"},
}

func init() {
	rootCmd.AddCommand(embedCmd)
}
