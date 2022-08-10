package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	tools "go-gqlgen-template"
	"go-gqlgen-template/ent"
	"go-gqlgen-template/graph/generated"
	"strconv"
)

// AddUser is the resolver for the addUser field.
func (r *mutationResolver) AddUser(ctx context.Context, input ent.CreateUserInput) (*tools.MutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input ent.UpdateUserInput) (*tools.MutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

// GetUser is the resolver for the getUser field.
func (r *queryResolver) GetUser(ctx context.Context, id string) (*ent.User, error) {
	i, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	u, err := r.controller.User.Get(ctx, &i)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// AllUser is the resolver for the allUser field.
func (r *queryResolver) AllUser(ctx context.Context) ([]*ent.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// CreateAt is the resolver for the createAt field.
func (r *userResolver) CreateAt(ctx context.Context, obj *ent.User) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// UpdateAt is the resolver for the updateAt field.
func (r *userResolver) UpdateAt(ctx context.Context, obj *ent.User) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Todo is the resolver for the todo field.
func (r *userResolver) Todo(ctx context.Context, obj *ent.User) ([]*ent.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

// ID is the resolver for the id field.
func (r *createUserInputResolver) ID(ctx context.Context, obj *ent.CreateUserInput, data string) error {
	panic(fmt.Errorf("not implemented"))
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

// CreateUserInput returns generated.CreateUserInputResolver implementation.
func (r *Resolver) CreateUserInput() generated.CreateUserInputResolver {
	return &createUserInputResolver{r}
}

type userResolver struct{ *Resolver }
type createUserInputResolver struct{ *Resolver }
