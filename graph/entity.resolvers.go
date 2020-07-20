package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	generated1 "github.com/josemarluedke/gqlgen-fed-service-type/graph/generated"
	"github.com/josemarluedke/gqlgen-fed-service-type/graph/model"
)

func (r *entityResolver) FindServiceByID(ctx context.Context, id string) (*model.Service, error) {
	return &model.Service{
		ID: id,
	}, nil
}

// Entity returns generated1.EntityResolver implementation.
func (r *Resolver) Entity() generated1.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
