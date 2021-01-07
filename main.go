package main // make sure to change package name

import (
	"github.com/joho/godotenv"
	impl "github.com/n4de4k/go-webapi-boilerplate/app/provider/v1"
	gqlDelivery "github.com/n4de4k/go-webapi-boilerplate/delivery/graphql"
	//"github.com/n4de4k/go-webapi-boilerplate/delivery/rest"
)

func main() {
	godotenv.Load()
	services := impl.NewServiceComponentImpl()

	// Use Rest API
	//restHandler := rest.NewHandler(services)
	//rest.Init(restHandler)

	// Use GraphQl
	gqlDelivery.InitGraphQLServer(services)
}
