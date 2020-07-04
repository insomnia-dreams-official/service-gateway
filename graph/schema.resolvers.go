package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	catalogPB "github.com/insomnia-dreams-official/service-catalog/pkg/protobuf"
	"log"
	"time"

	"github.com/insomnia-dreams-official/service-gateway/graph/generated"
	"github.com/insomnia-dreams-official/service-gateway/graph/model"
)

func (r *queryResolver) Navigation(ctx context.Context) ([]*model.NavigationItem, error) {
	// Resolver for getting site navigation by grpc from service-catalog.
	// ******************************************************************

	// Create cancellation context
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	// Get navigation items from service-catalog
	resp, err := r.CatalogClient.GetNavigationItems(ctx, &catalogPB.GetNavigationItemsRequest{})
	if err != nil {
		log.Printf("can't get navigation from service-catalog: %v", err)
		return nil, err
	}

	// Transform navigation items to graphql type
	var navigationItems []*model.NavigationItem
	for _, i := range resp.NavigationItems {
		navigationItems = append(navigationItems, model.GrpcToNavigationItem(i))
	}

	return navigationItems, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
