package v88

import (
	"errors"
	"time"

	camundav88 "github.com/grafvonb/c8volt/internal/clients/camunda/v88/camunda"
	operatev88 "github.com/grafvonb/c8volt/internal/clients/camunda/v88/operate"
	d "github.com/grafvonb/c8volt/internal/domain"
	"github.com/grafvonb/c8volt/toolx"
)

func fromProcessInstanceResponse(r operatev88.ProcessInstance) d.ProcessInstance {
	return d.ProcessInstance{
		BpmnProcessId:             toolx.Deref(r.BpmnProcessId, ""),
		EndDate:                   toolx.Deref(r.EndDate, ""),
		Incident:                  toolx.Deref(r.Incident, false),
		Key:                       toolx.Int64PtrToString(r.Key),
		ParentFlowNodeInstanceKey: toolx.Int64PtrToString(r.ParentFlowNodeInstanceKey),
		ParentKey:                 toolx.Int64PtrToString(r.ParentKey),
		ProcessDefinitionKey:      toolx.Int64PtrToString(r.ProcessDefinitionKey),
		ProcessVersion:            toolx.Deref(r.ProcessVersion, int32(0)),
		ProcessVersionTag:         toolx.Deref(r.ProcessVersionTag, ""),
		StartDate:                 toolx.Deref(r.StartDate, ""),
		State:                     d.State(*r.State),
		TenantId:                  toolx.Deref(r.TenantId, ""),
	}
}

func fromProcessInstanceResult(r camundav88.ProcessInstanceResult) d.ProcessInstance {
	return d.ProcessInstance{
		BpmnProcessId:             r.ProcessDefinitionId,
		EndDate:                   formatTimePtr(r.EndDate),
		Incident:                  r.HasIncident,
		Key:                       r.ProcessInstanceKey,
		ParentFlowNodeInstanceKey: toolx.Deref(r.ParentElementInstanceKey, ""),
		ParentKey:                 toolx.Deref(r.ParentProcessInstanceKey, ""),
		ProcessDefinitionKey:      r.ProcessDefinitionKey,
		ProcessVersion:            r.ProcessDefinitionVersion,
		ProcessVersionTag:         toolx.Deref(r.ProcessDefinitionVersionTag, ""),
		StartDate:                 formatTime(r.StartDate),
		State:                     d.State(r.State),
		TenantId:                  r.TenantId,
	}
}

func formatTime(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.UTC().Format(time.RFC3339Nano)
}

func formatTimePtr(p *time.Time) string {
	if p == nil {
		return ""
	}
	return formatTime(*p)
}

func toProcessInstanceCreationInstruction(d d.ProcessInstanceData) (camundav88.ProcessInstanceCreationInstruction, error) {
	var instr camundav88.ProcessInstanceCreationInstruction

	switch {
	// BPMN Process ID
	case d.BpmnProcessId != "":
		err := instr.FromProcessInstanceCreationInstructionById(
			camundav88.ProcessInstanceCreationInstructionById{
				ProcessDefinitionId:      d.BpmnProcessId,
				ProcessDefinitionVersion: normalizeVersion(d.ProcessDefinitionVersion), // -1 => latest
				Variables:                toolx.PtrCopyMap(d.Variables),
				TenantId:                 toolx.PtrIf(d.TenantId, ""),
			},
		)
		return instr, err
	// Process Definition "Key", actually internal unique ID
	case d.ProcessDefinitionSpecificId != "":
		err := instr.FromProcessInstanceCreationInstructionByKey(
			camundav88.ProcessInstanceCreationInstructionByKey{
				ProcessDefinitionKey: d.ProcessDefinitionSpecificId,
				Variables:            toolx.PtrCopyMap(d.Variables),
				TenantId:             toolx.PtrIf(d.TenantId, ""),
			},
		)
		return instr, err
	default:
		return instr, errors.New("provide ProcessDefinitionId or ProcessDefinitionKey")
	}
}

func normalizeVersion(v int32) *int32 {
	// Camunda latest sentinel is -1
	switch {
	case v == -1:
		return toolx.Ptr(int32(-1))
	case v > 0:
		return toolx.Ptr(v)
	default:
		// 0 or unset. Default to latest = -1
		return toolx.Ptr(int32(-1))
	}
}

func fromPostProcessInstancesResponse(r camundav88.CreateProcessInstanceResult) d.ProcessInstanceCreation {
	return d.ProcessInstanceCreation{
		Key:                      r.ProcessInstanceKey,
		BpmnProcessId:            r.ProcessDefinitionId,
		ProcessDefinitionKey:     r.ProcessDefinitionKey,
		ProcessDefinitionVersion: r.ProcessDefinitionVersion,
		TenantId:                 r.TenantId,
		Variables:                toolx.CopyMap(r.Variables),
		StartConfirmedAt:         "<not available>",
	}
}
