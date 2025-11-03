package task

type API interface{}

var _ API = (*client)(nil)
