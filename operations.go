package osbserver

import (
	"github.com/arschles/osbserver/request"
	"github.com/arschles/osbserver/response"
)

type Operations interface {
	Catalog() (*response.Catalog, error)
	Provision(*request.Provision) (*response.Provision, error)
	Deprovision(string) (*response.Deprovision, error)
	UpdateInstance(*request.UpdateServiceInstance) (*response.UpdateServiceInstance, error)
	Bind(*request.Bind) (*response.Bind, error)
	Unbind(string, string) (*response.Unbind, error)
}
