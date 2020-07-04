package graph

import catalog "github.com/insomnia-dreams-official/service-catalog/pkg/protobuf"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	CatalogClient catalog.CatalogClient
}
