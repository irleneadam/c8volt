package v88

import (
	camundav88 "github.com/grafvonb/c8volt/internal/clients/camunda/v88/camunda"
	d "github.com/grafvonb/c8volt/internal/domain"
	"github.com/grafvonb/c8volt/toolx"
)

//nolint:unused
func fromProcessDefinitionResponse(r camundav88.ProcessDefinitionResult) d.ProcessDefinition {
	return d.ProcessDefinition{
		BpmnProcessId:     toolx.Deref(r.ProcessDefinitionId, ""),
		Key:               toolx.Deref(r.ProcessDefinitionKey, ""),
		Name:              toolx.Deref(r.Name, ""),
		TenantId:          toolx.Deref(r.TenantId, ""),
		ProcessVersion:    toolx.Deref(r.Version, int32(0)),
		ProcessVersionTag: toolx.Deref(r.VersionTag, ""),
	}
}

func fromProcessDefinitionResult(r camundav88.ProcessDefinitionResult) d.ProcessDefinition {
	return d.ProcessDefinition{
		BpmnProcessId:     toolx.Deref(r.ProcessDefinitionId, ""),
		Key:               toolx.Deref(r.ProcessDefinitionKey, ""),
		Name:              toolx.Deref(r.Name, ""),
		TenantId:          toolx.Deref(r.TenantId, ""),
		ProcessVersion:    toolx.Deref(r.Version, int32(0)),
		ProcessVersionTag: toolx.Deref(r.VersionTag, ""),
	}
}
