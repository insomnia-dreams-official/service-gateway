package model

import catalog "github.com/insomnia-dreams-official/service-catalog/pkg/protobuf"

// GrpcToNavigationItem converts pb structure to graphql structure
func GrpcToNavigationItem(pni *catalog.NavigationItem) *NavigationItem {
	// Create item
	ni := NavigationItem{
		ID:    pni.Id,
		Name:  pni.Name,
		Link:  pni.Link,
		Items: []*NavigationItem{},
	}
	// Transform nested items
	for _, si := range pni.Items {
		ni.Items = append(ni.Items, &NavigationItem{
			ID:    si.Id,
			Name:  si.Name,
			Link:  si.Link,
			Items: []*NavigationItem{},
		})
	}
	return &ni
}

func GrpcToCategory(pc *catalog.Category) *Category {
	return &Category{
		Articul:  pc.Articul,
		Name:     pc.Name,
		Path:     pc.Path,
		Link:     pc.Link,
		FullLink: pc.FullLink,
	}
}
