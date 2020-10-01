package main // make sure to change package name
import (
	"github.com/joho/godotenv"
	"github.com/n4de4k/web-api-boilerplate/delivery/rest"
	"github.com/n4de4k/web-api-boilerplate/internal/component"
)

func main() {
	godotenv.Load()
	services := component.NewServiceComponentImpl()

	// Use Rest API
	restHandler := rest.NewHandler(services)
	rest.Init(restHandler)
}
