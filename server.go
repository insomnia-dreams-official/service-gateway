package main

import (
	"context"
	catalog "github.com/insomnia-dreams-official/service-catalog/pkg/protobuf"
	"github.com/rs/cors"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
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

	router := chi.NewRouter()
	router.Use(
		cors.New(cors.Options{
			AllowedOrigins:   []string{"http://localhost:3000"},
			AllowCredentials: true,
			Debug:            true,
		}).Handler,
	)

	// Inject grpc clients in resolvers
	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{
					CatalogClient: catalog.NewCatalogClient(conn),
				},
			},
		),
	)

	// Register http handlers
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	// Run http server
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func cors2(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		h(w, r)
	}
}
