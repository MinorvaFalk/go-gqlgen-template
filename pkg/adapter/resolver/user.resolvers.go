package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"go-gqlgen-template/ent"
	"go-gqlgen-template/ent/schema/ulid"
	"go-gqlgen-template/graph/generated"
	"go-gqlgen-template/pkg/adapter/handler"
	"go-gqlgen-template/pkg/entity/model"
	"go-gqlgen-template/utils"
)

// AddUser is the resolver for the addUser field.
func (r *mutationResolver) AddUser(ctx context.Context, input ent.CreateUserInput) (*model.UserMutationResponse, error) {
	u, err := r.controller.User.Create(ctx, input)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}

	return &model.UserMutationResponse{
		Code:    &utils.SuccessCode,
		Success: &utils.SuccessStatus,
		Message: &utils.SuccessMessage,
		Data:    u,
	}, nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input ent.UpdateUserInput) (*model.UserMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id ulid.ID) (*ent.User, error) {
	return r.controller.User.Get(ctx, &id)
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.UserWhereInput) (*ent.UserConnection, error) {
	return r.controller.User.List(ctx, after, first, before, last, where)
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
	t, err := r.controller.User.GetTodo(ctx, &obj.ID)
	if err != nil {
		return nil, err
	}

	return t, nil
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
