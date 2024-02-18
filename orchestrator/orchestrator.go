package orchestrator

import (
	mdl "distr-calc/model"
	"time"
)

var resources = map[string][]mdl.ComputingResource{
	"Resources": {
		{Status: "Work", Name: "computing server", LastPing: time.Now()},
		{Status: "Reconnect", Name: "computing server", LastPing: time.Now().Add(time.Minute * (-5))},
		{Status: "Lost", Name: "computing server", LastPing: time.UnixMicro(0)},
	},
}

func GetResources() map[string][]mdl.ComputingResource {
	return resources
}
