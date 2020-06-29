package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/insomnia-dreams-official/service-gateway/graph/generated"
	"github.com/insomnia-dreams-official/service-gateway/graph/model"
)

// Grpc call to the category service
func (r *queryResolver) Navigation(ctx context.Context) ([]*model.NavigationItem, error) {
	item := model.NavigationItem{ID: "1", Name: "Test", Link: "test"}
	return []*model.NavigationItem{
		{
			ID:   "1",
			Name: "Test",
			Link: "test",
			Items: []*model.NavigationItem{
				&item,
			},
		},
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
