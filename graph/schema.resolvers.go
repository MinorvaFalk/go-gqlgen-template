package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"go-gqlgen-template/graph/generated"
	"go-gqlgen-template/graph/model"
	"go-gqlgen-template/utils"
)

// Photos is the resolver for the photos field.
func (r *albumResolver) Photos(ctx context.Context, obj *model.Album) ([]*model.Photo, error) {
	panic(fmt.Errorf("not implemented"))
}

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input model.CreatePostInput) (*model.MutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

// UpdatePost is the resolver for the updatePost field.
func (r *mutationResolver) UpdatePost(ctx context.Context, id int, input model.UpdatePostInput) (*model.MutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

// DeletePost is the resolver for the deletePost field.
func (r *mutationResolver) DeletePost(ctx context.Context, id int) (*model.MutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

// CreateComment is the resolver for the createComment field.
func (r *mutationResolver) CreateComment(ctx context.Context, input model.CreateCommentInput) (*model.MutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

// UpdateComment is the resolver for the updateComment field.
func (r *mutationResolver) UpdateComment(ctx context.Context, id int, input model.UpdateCommentInput) (*model.MutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

// DeleteComment is the resolver for the deleteComment field.
func (r *mutationResolver) DeleteComment(ctx context.Context, id int) (*model.MutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

// CreateAlbum is the resolver for the createAlbum field.
func (r *mutationResolver) CreateAlbum(ctx context.Context, input model.CreateAlbumInput) (*model.MutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

// UpdateAlbum is the resolver for the updateAlbum field.
func (r *mutationResolver) UpdateAlbum(ctx context.Context, id int, input model.UpdateAlbumInput) (*model.MutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

// DeleteAlbum is the resolver for the deleteAlbum field.
func (r *mutationResolver) DeleteAlbum(ctx context.Context, id int) (*model.MutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

// CreatePhoto is the resolver for the createPhoto field.
func (r *mutationResolver) CreatePhoto(ctx context.Context, input model.CreatePhotoInput) (*model.MutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

// UpdatePhoto is the resolver for the updatePhoto field.
func (r *mutationResolver) UpdatePhoto(ctx context.Context, id int, input model.UpdatePhotoInput) (*model.MutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

// DeletePhoto is the resolver for the deletePhoto field.
func (r *mutationResolver) DeletePhoto(ctx context.Context, id int) (*model.MutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.CreateTodoInput) (*model.MutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

// UpdateTodo is the resolver for the updateTodo field.
func (r *mutationResolver) UpdateTodo(ctx context.Context, id int, input model.UpdateTodoInput) (*model.MutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

// DeleteTodo is the resolver for the deleteTodo field.
func (r *mutationResolver) DeleteTodo(ctx context.Context, id int) (*model.MutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.MutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id int, input model.UpdateUserInput) (*model.MutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id int) (*model.MutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

// Comments is the resolver for the comments field.
func (r *postResolver) Comments(ctx context.Context, obj *model.Post) ([]*model.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

// Status is the resolver for the status field.
func (r *queryResolver) Status(ctx context.Context) (*model.StatusResponse, error) {
	return &model.StatusResponse{
		Response: utils.ServerStatus,
	}, nil
}

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

// Todos is the resolver for the todos field.
func (r *userResolver) Todos(ctx context.Context, obj *model.User) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

// Albums is the resolver for the albums field.
func (r *userResolver) Albums(ctx context.Context, obj *model.User) ([]*model.Album, error) {
	panic(fmt.Errorf("not implemented"))
}

// Album returns generated.AlbumResolver implementation.
func (r *Resolver) Album() generated.AlbumResolver { return &albumResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Post returns generated.PostResolver implementation.
func (r *Resolver) Post() generated.PostResolver { return &postResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type albumResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type postResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
