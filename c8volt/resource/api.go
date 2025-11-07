package resource

import (
	"context"

	"github.com/grafvonb/c8volt/c8volt/foptions"
)

type API interface {
	DeployProcessDefinition(ctx context.Context, tenantId string, units []DeploymentUnitData, opts ...foptions.FacadeOption) ([]ProcessDefinitionDeployment, error)

	DeleteProcessDefinitionByKey(ctx context.Context, key string, opts ...foptions.FacadeOption) error
	DeleteProcessDefinitionVersionsByBpmnProcessId(ctx context.Context, bpmnProcessId string, opts ...foptions.FacadeOption) error
}

var _ API = (*client)(nil)
