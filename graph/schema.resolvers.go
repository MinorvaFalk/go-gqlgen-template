package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"go-gqlgen-template/graph/generated"
)

// Example is the resolver for the example field.
func (r *queryResolver) Example(ctx context.Context) (*string, error) {
	response := "hello world"
	return &response, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
