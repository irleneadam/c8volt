package cmd

import (
	"fmt"

	"github.com/grafvonb/c8volt/c8volt/ferrors"
	"github.com/grafvonb/c8volt/c8volt/process"
	"github.com/grafvonb/c8volt/toolx"
	"github.com/spf13/cobra"
)

var (
	flagRunPIProcessDefinitionBpmnProcessId string
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
		data := process.ProcessInstanceData{
			BpmnProcessId: flagRunPIProcessDefinitionBpmnProcessId,
		}
		pi, err := cli.CreateProcessInstance(cmd.Context(), data, collectOptions()...)
		log.Debug(toolx.ToJSONString(pi))
		if err != nil {
			ferrors.HandleAndExit(log, cfg.App.NoErrCodes, fmt.Errorf("running process instance for process definition bpmn id %s: %w", flagRunPIProcessDefinitionBpmnProcessId, err))
		}
	},
}

func init() {
	runCmd.AddCommand(runProcessInstanceCmd)

	runProcessInstanceCmd.Flags().StringVarP(&flagRunPIProcessDefinitionBpmnProcessId, "bpmn-process-id", "b", "", "BPMN process ID to run process instance for")
}
