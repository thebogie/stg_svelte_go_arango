package registry

import (
	"back/pkg/adapter/controller"
	"github.com/arangodb/go-driver"
)

type registry struct {
	db driver.Database
}

// Registry is an interface of registry
type Registry interface {
	NewController() controller.Controller
}

// New registers entire controller with dependencies
func New(db driver.Database) Registry {

	return &registry{
		db: db,
	}
}

// NewController generates controllers
func (r *registry) NewController() controller.Controller {
	return controller.Controller{
		User:    r.NewUserController(),
		Game:    r.NewGameController(),
		Contest: r.NewContestController(),
	}
}
