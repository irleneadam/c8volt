package cmd

import (
	"github.com/spf13/cobra"
)

var embedCmd = &cobra.Command{
	Use:   "embed",
	Short: "Manage embedded resources",
	Long: "Manage embedded resources such as embedded BPMN process definitions.\n" +
		"It is a root command and requires a subcommand to specify the action to perform on embedded resources.",
	Aliases: []string{"em", "emb"},
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
	SuggestFor: []string{"embedd", "embd", "embedded", "embeded"},
}

func init() {
	rootCmd.AddCommand(embedCmd)
}
