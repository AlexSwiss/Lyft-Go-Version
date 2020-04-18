package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/AlexSwiss/lyft/graph/generated"
	"github.com/AlexSwiss/lyft/graph/model"
)

func (r *mutationResolver) CreateMovie(ctx context.Context, input model.NewMovie) (*model.Movie, error) {
	movie := &model.Movie{
		ID:   fmt.Sprintf("T%d", rand.Int()),
		Name: input.Name,
		Year: input.Year,
		Actor: &model.Actor{
			ID:        fmt.Sprintf("T%d", rand.Int()),
			Firstname: input.ActorFirstname,
			Lastname:  input.ActorLastname},
	}
	r.movies = append(r.movies, movie)
	return movie, nil
}

func (r *queryResolver) Movies(ctx context.Context) ([]*model.Movie, error) {
	return r.movies, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
