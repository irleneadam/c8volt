package v88

import (
	"context"

	camundav88 "github.com/grafvonb/c8volt/internal/clients/camunda/v88/camunda"
	operatev88 "github.com/grafvonb/c8volt/internal/clients/camunda/v88/operate"
)

type GenClusterClientCamunda interface {
	CancelProcessInstanceWithResponse(ctx context.Context, processInstanceKey string, body camundav88.CancelProcessInstanceJSONRequestBody, reqEditors ...camundav88.RequestEditorFn) (*camundav88.CancelProcessInstanceResponse, error)
	GetProcessInstanceWithResponse(ctx context.Context, processInstanceKey string, reqEditors ...camundav88.RequestEditorFn) (*camundav88.GetProcessInstanceResponse, error)
}

type GenClusterClientOperate interface {
	SearchProcessInstancesWithResponse(ctx context.Context, body operatev88.SearchProcessInstancesJSONRequestBody, reqEditors ...operatev88.RequestEditorFn) (*operatev88.SearchProcessInstancesResponse, error)
	DeleteProcessInstanceAndAllDependantDataByKeyWithResponse(ctx context.Context, key int64, reqEditors ...operatev88.RequestEditorFn) (*operatev88.DeleteProcessInstanceAndAllDependantDataByKeyResponse, error)
}

var _ GenClusterClientCamunda = (*camundav88.ClientWithResponses)(nil)
var _ GenClusterClientOperate = (*operatev88.ClientWithResponses)(nil)
