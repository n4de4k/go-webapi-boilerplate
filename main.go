package main // make sure to change package name
import (
	"github.com/joho/godotenv"
	"github.com/n4de4k/web-api-boilerplate/component/impl"
	gqlDelivery "github.com/n4de4k/web-api-boilerplate/delivery/graphql"

	//"github.com/n4de4k/web-api-boilerplate/delivery/rest"
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
