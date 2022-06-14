package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/tankip/go-social/api/users"
	"github.com/tankip/go-social/graph/generated"
	"github.com/tankip/go-social/graph/model"
)

func (r *queryResolver) Users(ctx context.Context, filter *model.UserFilter) ([]*model.User, error) {

	var resultUsers []*model.User
	var allUsers []users.User = users.GetUsers(filter)
	for _, user := range allUsers {
		resultUsers = append(resultUsers, &model.User{
			ID:   user.ID,
			Name: user.Name,
			Year: int(user.Year),
		})
	}
	return resultUsers, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
