package cluster

import (
	"context"

	"github.com/grafvonb/c8volt/c8volt/foptions"
)

type API interface {
	GetClusterTopology(ctx context.Context, opts ...foptions.FacadeOption) (Topology, error)
}

var _ API = (*client)(nil)
