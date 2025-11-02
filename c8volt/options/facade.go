package options

import "github.com/grafvonb/c8volt/internal/services"

func WithNoStateCheck() FacadeOption { return func(c *FacadeCfg) { c.NoStateCheck = true } }
func WithForce() FacadeOption        { return func(c *FacadeCfg) { c.Force = true } }
func WithNoWait() FacadeOption       { return func(c *FacadeCfg) { c.NoWait = true } }

type FacadeOption func(*FacadeCfg)

type FacadeCfg struct {
	NoStateCheck bool
	Force        bool
	NoWait       bool
}

func ApplyFacadeOptions(opts []FacadeOption) *FacadeCfg {
	c := &FacadeCfg{}
	for _, o := range opts {
		o(c)
	}
	return c
}

func MapFacadeOptionsToCallOptions(opts []FacadeOption) []services.CallOption {
	c := ApplyFacadeOptions(opts)
	var out []services.CallOption
	if c.NoStateCheck {
		out = append(out, services.WithNoStateCheck())
	}
	if c.Force {
		out = append(out, services.WithForce())
	}
	if c.NoWait {
		out = append(out, services.WithNoWait())
	}
	return out
}
