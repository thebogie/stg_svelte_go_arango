package resolver

import (
	"back/graph/generated"
	"back/pkg/adapter/controller"
	"github.com/99designs/gqlgen/graphql"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	controller.Controller
}

func NewSchema(controller controller.Controller) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{controller},
	})
}
