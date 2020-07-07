module github.com/insomnia-dreams-official/service-gateway

go 1.14

replace github.com/insomnia-dreams-official/service-catalog => /home/max/new-store/service-catalog

require (
	github.com/99designs/gqlgen v0.11.3
	github.com/go-chi/chi v3.3.2+incompatible
	github.com/insomnia-dreams-official/service-catalog v0.0.0-20200702123949-140ea87614e2
	github.com/rs/cors v1.6.0
	github.com/spf13/viper v1.7.0
	github.com/vektah/gqlparser/v2 v2.0.1
	google.golang.org/grpc v1.30.0
)
