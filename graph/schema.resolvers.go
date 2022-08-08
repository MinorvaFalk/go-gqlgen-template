package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"go-gqlgen-template/graph/generated"
	"go-gqlgen-template/graph/model"
)

// GetPosts is the resolver for the getPosts field.
func (r *queryResolver) GetPosts(ctx context.Context, id *int) ([]*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

// GetComments is the resolver for the getComments field.
func (r *queryResolver) GetComments(ctx context.Context, id *int) ([]*model.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

// GetAlbums is the resolver for the getAlbums field.
func (r *queryResolver) GetAlbums(ctx context.Context, id *int) ([]*model.Album, error) {
	panic(fmt.Errorf("not implemented"))
}

// GetPhotos is the resolver for the getPhotos field.
func (r *queryResolver) GetPhotos(ctx context.Context, id *int) ([]*model.Photo, error) {
	panic(fmt.Errorf("not implemented"))
}

// GetTodos is the resolver for the getTodos field.
func (r *queryResolver) GetTodos(ctx context.Context, id *int) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

// GetUsers is the resolver for the getUsers field.
func (r *queryResolver) GetUsers(ctx context.Context, id *int) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
