package cmd

import (
	"github.com/grafvonb/c8volt/c8volt/ferrors"
	"github.com/spf13/cobra"
)

var (
	flagRunPIProcessDefinitionKeys []string
)

var runProcessInstanceCmd = &cobra.Command{
	Use:     "process-instance",
	Short:   "Run a process instance by process definition key(s)",
	Aliases: []string{"pi"},
	Run: func(cmd *cobra.Command, args []string) {
		cli, log, cfg, err := NewCli(cmd)
		if err != nil {
			ferrors.HandleAndExit(log, cfg.App.NoErrCodes, err)
		}
		_ = cli
		/*
			_, err = cli.RunProcessInstance(cmd.Context(), flagRunPIProcessDefinitionKeys, collectOptions()...)
			if err != nil {
				ferrors.HandleAndExit(log, cfg.App.NoErrCodes, fmt.Errorf("running process instance for process definition key %s: %w", flagRunPIProcessDefinitionKeys, err))
			}

		*/

	},
}

func init() {
	runCmd.AddCommand(runProcessInstanceCmd)

	runProcessInstanceCmd.Flags().StringSliceVarP(&flagRunPIProcessDefinitionKeys, "keys", "k", nil, "process definition key(s) to run process instance(s) for")
	_ = runProcessInstanceCmd.MarkFlagRequired("keys")
}
