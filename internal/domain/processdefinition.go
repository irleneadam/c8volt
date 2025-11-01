package domain

import "slices"

type ProcessDefinition struct {
	BpmnProcessId string
	Key           string
	Name          string
	TenantId      string
	Version       int32
	VersionTag    string
}

type ProcessDefinitionSearchFilterOpts struct {
	Key           string
	BpmnProcessId string
	Version       int32
	VersionTag    string
}

func SortByVersionDesc(pds []ProcessDefinition) {
	slices.SortFunc(pds, func(a, b ProcessDefinition) int {
		switch {
		case a.Version > b.Version:
			return -1 // a before b
		case a.Version < b.Version:
			return 1 // b before a
		default:
			return 0
		}
	})
}

func SortByBpmnProcessIdAscThenByVersionDesc(pds []ProcessDefinition) {
	slices.SortFunc(pds, func(a, b ProcessDefinition) int {
		if a.BpmnProcessId < b.BpmnProcessId {
			return -1
		}
		if a.BpmnProcessId > b.BpmnProcessId {
			return 1
		}
		switch {
		case a.Version > b.Version:
			return -1
		case a.Version < b.Version:
			return 1
		default:
			return 0
		}
	})
}
