package cluster

import (
	"context"

	"github.com/grafvonb/c8volt/c8volt/options"
)

type API interface {
	GetClusterTopology(ctx context.Context, opts ...options.FacadeOption) (Topology, error)
}

var _ API = (*client)(nil)
