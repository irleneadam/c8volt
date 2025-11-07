package process

import (
	"context"

	"github.com/grafvonb/c8volt/c8volt/ferrors"
	"github.com/grafvonb/c8volt/c8volt/foptions"
	"github.com/grafvonb/c8volt/c8volt/fpool"
	"github.com/grafvonb/c8volt/toolx"
)

func (c *client) CreateNProcessInstances(ctx context.Context, data ProcessInstanceData, n, parallel int, opts ...foptions.FacadeOption) ([]ProcessInstance, error) {
	cCfg := foptions.ApplyFacadeOptions(opts)

	return fpool.ExecuteNTimes[ProcessInstance](ctx, n, parallel, cCfg.FailFast, func(ctx context.Context, _ int) (ProcessInstance, error) {
		pic, err := c.piApi.CreateProcessInstance(ctx, toProcessInstanceData(data), foptions.MapFacadeOptionsToCallOptions(opts)...)
		if err != nil {
			return ProcessInstance{}, ferrors.FromDomain(err)
		}
		return fromDomainProcessInstanceCreation(pic), nil
	})
}

func (c *client) CancelProcessInstances(ctx context.Context, keys []string, parallel int, failFast bool, opts ...foptions.FacadeOption) ([]CancelResponse, error) {
	ukeys := toolx.UniqueSlice(keys)
	return fpool.ExecuteSlice[string, CancelResponse](ctx, ukeys, parallel, failFast, func(ctx context.Context, key string, _ int) (CancelResponse, error) {
		return c.CancelProcessInstance(ctx, key, opts...)
	})
}
