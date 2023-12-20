package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"back/graph/model"
	"context"
	"fmt"
)

// Todos is the resolver for the todos field.
func (r *queryResolver) Contests(ctx context.Context) ([]*model.Contest, error) {
	return r.Contest.List(ctx)
}

// GetContestsPlayerTotalResults is the resolver for the GetContestsPlayerTotalResults field.
func (r *queryResolver) GetContestsPlayerTotalResults(ctx context.Context, player string) ([]*model.Contest, error) {
	fmt.Println("This is from contest resolver")
	var test, err = r.Contest.GetContestsPlayerTotalResults(ctx, player)

	fmt.Println("This is from contest rescolver return %+v\n", test)
	return test, err
}
