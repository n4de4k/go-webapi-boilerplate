package main // make sure to change package name
import (
	"github.com/joho/godotenv"
	"github.com/n4de4k/web-api-boilerplate/api/rest"
	"github.com/n4de4k/web-api-boilerplate/app/impl/component"
)

func main() {
	godotenv.Load()
	services := component.NewServiceComponentImpl()

	// Use Rest API
	restHandler := rest.NewHandler(services)
	rest.Init(restHandler)
}
