package cmd

import (
	"fmt"

	"github.com/grafvonb/c8volt/c8volt/ferrors"
	"github.com/grafvonb/c8volt/c8volt/process"
	"github.com/spf13/cobra"
)

var (
	flagRunPIProcessDefinitionBpmnProcessIds []string
	flagRunPIProcessDefinitionSpecificId     []string
	flagRunPIProcessDefinitionVersion        int32
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
		var datas []process.ProcessInstanceData
		switch {
		case len(flagRunPIProcessDefinitionSpecificId) > 0:
			if len(flagRunPIProcessDefinitionBpmnProcessIds) > 0 {
				ferrors.HandleAndExit(log, cfg.App.NoErrCodes, fmt.Errorf("flags --pd-id and --bpmn-process-id are mutually exclusive"))
			}
			datas = make([]process.ProcessInstanceData, 0, len(flagRunPIProcessDefinitionSpecificId))
			for _, pdId := range flagRunPIProcessDefinitionSpecificId {
				data := process.ProcessInstanceData{
					ProcessDefinitionSpecificId: pdId,
					TenantId:                    cfg.App.Tenant,
				}
				datas = append(datas, data)
			}
		case len(flagRunPIProcessDefinitionBpmnProcessIds) > 0:
			if flagRunPIProcessDefinitionVersion > 0 && len(flagRunPIProcessDefinitionBpmnProcessIds) > 1 {
				ferrors.HandleAndExit(log, cfg.App.NoErrCodes, fmt.Errorf("flag --pd-version is not supported when running process instances for multiple BPMN process IDs"))
			}
			datas = make([]process.ProcessInstanceData, 0, len(flagRunPIProcessDefinitionBpmnProcessIds))
			for _, bpmnProcessId := range flagRunPIProcessDefinitionBpmnProcessIds {
				data := process.ProcessInstanceData{
					BpmnProcessId:            bpmnProcessId,
					TenantId:                 cfg.App.Tenant,
					ProcessDefinitionVersion: flagRunPIProcessDefinitionVersion,
				}
				datas = append(datas, data)
			}
		default:
			ferrors.HandleAndExit(log, cfg.App.NoErrCodes, fmt.Errorf("either --pd-id or --bpmn-process-id must be provided to run a process instance(s)"))
		}
		_, err = cli.CreateProcessInstances(cmd.Context(), datas, collectOptions()...)
		if err != nil {
			ferrors.HandleAndExit(log, cfg.App.NoErrCodes, fmt.Errorf("running process instance(s) for BPMN process ID(s) %v: %w", flagRunPIProcessDefinitionBpmnProcessIds, err))
		}
	},
}

func init() {
	runCmd.AddCommand(runProcessInstanceCmd)

	runProcessInstanceCmd.Flags().StringSliceVarP(&flagRunPIProcessDefinitionBpmnProcessIds, "bpmn-process-id", "b", nil, "BPMN process ID(s) to run process instance for (mutually exclusive with --pd-id). Runs latest version unless --pd-version is specified")
	runProcessInstanceCmd.Flags().Int32Var(&flagRunPIProcessDefinitionVersion, "pd-version", 0, "Specific version of the process definition to use when running by BPMN process ID (supported only with --bpmn-process-id)")
	runProcessInstanceCmd.Flags().StringSliceVar(&flagRunPIProcessDefinitionSpecificId, "pd-id", nil, "Specific process definition ID(s) to run process instance for (mutually exclusive with --bpmn-process-id)")
}
