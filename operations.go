package osbserver

import (
	"github.com/arschles/osbserver/request"
	"github.com/arschles/osbserver/response"
)

type Operations interface {
	Catalog() (*response.Catalog, *BrokerError)
	Provision(*request.Provision) (*response.Provision, *BrokerError)
	Deprovision(string) (*response.Deprovision, *BrokerError)
	PollLastOperation(string, *PollLastOperationParams) (*response.PollLastOperation, *BrokerError)
	UpdateInstance(*request.UpdateServiceInstance) (*response.UpdateServiceInstance, *BrokerError)
	Bind(*request.Bind) (*response.Bind, *BrokerError)
	Unbind(string, string) (*response.Unbind, *BrokerError)
}

type PollLastOperationParams struct {
	ServiceID string
	PlanID    string
	Operation string
}
