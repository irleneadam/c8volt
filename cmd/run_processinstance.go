package cmd

import (
	"fmt"

	"github.com/grafvonb/c8volt/c8volt/ferrors"
	"github.com/grafvonb/c8volt/c8volt/process"
	"github.com/grafvonb/c8volt/toolx"
	"github.com/spf13/cobra"
)

var (
	flagRunPIProcessDefinitionBpmnProcessIds []string
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
		l := len(flagRunPIProcessDefinitionBpmnProcessIds)
		if l == 0 {
			ferrors.HandleAndExit(log, cfg.App.NoErrCodes, fmt.Errorf("at least one BPMN process ID must be provided to run a process instance(s)"))
		}
		datas := make([]process.ProcessInstanceData, 0, l)
		for _, bpmnProcessId := range flagRunPIProcessDefinitionBpmnProcessIds {
			data := process.ProcessInstanceData{
				BpmnProcessId: bpmnProcessId,
				TenantId:      cfg.App.Tenant,
			}
			datas = append(datas, data)
		}
		pi, err := cli.CreateProcessInstances(cmd.Context(), datas, collectOptions()...)
		log.Debug(toolx.ToJSONString(pi))
		if err != nil {
			ferrors.HandleAndExit(log, cfg.App.NoErrCodes, fmt.Errorf("running process instance(s) for BPMN process ID(s) %v: %w", flagRunPIProcessDefinitionBpmnProcessIds, err))
		}
	},
}

func init() {
	runCmd.AddCommand(runProcessInstanceCmd)

	runProcessInstanceCmd.Flags().StringSliceVarP(&flagRunPIProcessDefinitionBpmnProcessIds, "bpmn-process-id", "b", nil, "BPMN process ID(s) to run process instance for")
}
