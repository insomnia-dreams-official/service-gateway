package model

import catalog "github.com/insomnia-dreams-official/service-catalog/pkg/protobuf"

// GrpcToNavigationItem converts pb structure to graphql structure
func GrpcToNavigationItem(i *catalog.NavigationItem) *NavigationItem {
	// Create item
	ni := NavigationItem{
		ID:    i.Id,
		Name:  i.Name,
		Link:  i.Link,
		Items: []*NavigationItem{},
	}
	// Transform nested items
	for _, si := range i.Items {
		ni.Items = append(ni.Items, &NavigationItem{
			ID:    si.Id,
			Name:  si.Name,
			Link:  si.Link,
			Items: []*NavigationItem{},
		})
	}
	return &ni
}
