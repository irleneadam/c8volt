package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	flagRunTenantId string
)

var runCmd = &cobra.Command{
	Use:     "run",
	Short:   "Run resources",
	Aliases: []string{"r"},
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
	SuggestFor: []string{"rum", "runn"},
}

func init() {
	rootCmd.AddCommand(runCmd)

	addBackoffFlagsAndBindings(runCmd, viper.GetViper())

	runCmd.PersistentFlags().StringVarP(&flagRunTenantId, "tenant-id", "t", "", "tenant id for the run")
}
