package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get resources",
	Long: "Get resources such as process definitions or process instances.\n" +
		"It is a root command and requires a subcommand to specify the resource type to get.",
	Aliases: []string{"g", "read"},
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
	SuggestFor: []string{"gett", "getr"},
}

func init() {
	rootCmd.AddCommand(getCmd)

	addBackoffFlagsAndBindings(getCmd, viper.GetViper())
}
