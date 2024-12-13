package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.57

import (
	"context"
	"time"

	"github.com/TrinityKnights/Backend/internal/delivery/graph"
	"github.com/TrinityKnights/Backend/internal/domain/model"
)

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateRequest) (*model.UserResponse, error) {
	user, err := r.UserService.Update(ctx, &input)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Profile is the resolver for the profile field.
func (r *queryResolver) Profile(ctx context.Context) (*model.UserResponse, error) {
	user, err := r.UserService.Profile(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// CreatedAt is the resolver for the createdAt field.
func (r *userResponseResolver) CreatedAt(ctx context.Context, obj *model.UserResponse) (*time.Time, error) {
	if obj.CreatedAt == "" {
		return nil, nil
	}
	t, err := time.Parse("2006-01-02 15:04:05.999999 -0700 MST", obj.CreatedAt)
	if err != nil {
		t, err = time.Parse("2006-01-02 15:04:05 -0700 MST", obj.CreatedAt)
		if err != nil {
			return nil, err
		}
	}
	return &t, nil
}

// UpdatedAt is the resolver for the updatedAt field.
func (r *userResponseResolver) UpdatedAt(ctx context.Context, obj *model.UserResponse) (*time.Time, error) {
	if obj.UpdatedAt == "" {
		return nil, nil
	}
	t, err := time.Parse("2006-01-02 15:04:05.999999 -0700 MST", obj.UpdatedAt)
	if err != nil {
		t, err = time.Parse("2006-01-02 15:04:05 -0700 MST", obj.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}
	return &t, nil
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

// UserResponse returns graph.UserResponseResolver implementation.
func (r *Resolver) UserResponse() graph.UserResponseResolver { return &userResponseResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResponseResolver struct{ *Resolver }