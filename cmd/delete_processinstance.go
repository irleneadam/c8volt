package cmd

import (
	"fmt"

	"github.com/grafvonb/c8volt/c8volt/ferrors"
	"github.com/spf13/cobra"
)

var (
	flagDeletePIKey       string
	flagDeletePIWithForce bool
)

var deleteProcessInstanceCmd = &cobra.Command{
	Use:     "process-instance",
	Short:   "Delete a process instance by its key",
	Aliases: []string{"pi"},
	Run: func(cmd *cobra.Command, args []string) {
		cli, log, cfg, err := NewCli(cmd)
		if err != nil {
			ferrors.HandleAndExit(log, cfg.App.NoErrCodes, err)
		}
		_, err = cli.DeleteProcessInstance(cmd.Context(), flagDeletePIKey, collectOptions()...)
		if err != nil {
			ferrors.HandleAndExit(log, cfg.App.NoErrCodes, fmt.Errorf("deleting process instance with key %s: %w", flagDeletePIKey, err))
		}

	},
}

func init() {
	deleteCmd.AddCommand(deleteProcessInstanceCmd)

	deleteProcessInstanceCmd.Flags().StringVarP(&flagDeletePIKey, "key", "k", "", "process instance key to delete")
	_ = deleteProcessInstanceCmd.MarkFlagRequired("key")
	deleteProcessInstanceCmd.Flags().BoolVar(&flagDeletePIWithForce, "force", false, "force cancellation of the process instance, prior to deletion")
}
