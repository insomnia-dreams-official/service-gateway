package main

import (
	"context"
	catalog "github.com/insomnia-dreams-official/service-catalog/pkg/protobuf"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/insomnia-dreams-official/service-gateway/graph"
	"github.com/insomnia-dreams-official/service-gateway/graph/generated"
)

// Load environment variables from a config file corresponding to stack (dev/test/prod)
func init() {
	// Define config file's path
	if dir, err := os.Getwd(); err == nil {
		viper.AddConfigPath(path.Join(dir, "config"))
	} else {
		log.Fatalf("can't get process's working directory for loading viper config file")
	}

	// Define config file's name
	switch os.Getenv("ENVIRONMENT") {
	case "dev":
		viper.SetConfigName("dev.config")
	case "test":
		viper.SetConfigName("test.config")
	case "prod":
		viper.SetConfigName("prod.config")
	default:
		log.Fatalf("define environment variable \"ENVIRONMENT\" with one of values: \"dev\", \"test\", \"prod\"; then restart service")
	}

	// Load config
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("can't load viper config file, reason: %v", err)
	}
}

func main() {
	// Configure targets and credentials
	port := viper.GetString(`gateway.port`)
	target := viper.GetString(`catalog.host`) + ":" + viper.GetString(`catalog.port`)

	// Create context with timeout to stop gateway if grpc services are dead
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Connect to grpc services
	conn, err := grpc.DialContext(ctx, target, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Inject grpc clients in resolvers
	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{
					CatalogClient: catalog.NewCatalogClient(conn),
				}}))

	// Register http handlers
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	// Run http server
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
