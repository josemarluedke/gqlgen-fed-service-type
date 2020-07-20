package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/josemarluedke/gqlgen-fed-service-type/graph/generated"
	"github.com/josemarluedke/gqlgen-fed-service-type/graph/model"
)

func (r *queryResolver) Service(ctx context.Context) (*model.Service, error) {
	return &model.Service{
		ID: "1234",
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
