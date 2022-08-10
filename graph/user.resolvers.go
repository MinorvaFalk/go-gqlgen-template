package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"go-gqlgen-template/ent"
	"go-gqlgen-template/ent/user"
	"go-gqlgen-template/graph/generated"
	"strconv"
)

// GetUser is the resolver for the getUser field.
func (r *queryResolver) GetUser(ctx context.Context, id *string) (*ent.User, error) {
	i, err := strconv.Atoi(*id)
	if err != nil {
		return nil, err
	}

	u, err := r.client.User.Query().Where(user.IDEQ(i)).Only(ctx)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// AllUser is the resolver for the allUser field.
func (r *queryResolver) AllUser(ctx context.Context) ([]*ent.User, error) {
	u, err := r.client.User.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// CreateAt is the resolver for the createAt field.
func (r *userResolver) CreateAt(ctx context.Context, obj *ent.User) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// UpdateAt is the resolver for the updateAt field.
func (r *userResolver) UpdateAt(ctx context.Context, obj *ent.User) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
