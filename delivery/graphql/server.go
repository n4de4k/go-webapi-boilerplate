package graphql

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/graph-gophers/graphql-go"
	"github.com/n4de4k/go-webapi-boilerplate/app/provider"
	"github.com/n4de4k/go-webapi-boilerplate/delivery/graphql/lib"
	"github.com/n4de4k/go-webapi-boilerplate/delivery/graphql/resolver"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func playgroundHandler() gin.HandlerFunc {
	playground := lib.Playground{}
	return func(c *gin.Context) {
		playground.ServeHTTP(c.Writer, c.Request)
	}
}

func grapqhlHandler(services provider.ServiceComponent) gin.HandlerFunc {
	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	schema := NewSchema()
	gqlSchema := graphql.MustParseSchema(
		*schema,
		&resolver.Resolver{Services: services},
		opts...,
	)

	gql := lib.GraphQL{Schema: gqlSchema}

	return func(c *gin.Context) {
		gql.ServeHTTP(c.Writer, c.Request)
	}
}

func InitGraphQLServer(services provider.ServiceComponent) {
	env := os.Getenv("APP_ENV")
	port := os.Getenv("APP_PORT")
	if env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	r.GET("/graphql", playgroundHandler())
	r.POST("/graphql", grapqhlHandler(services))

	log.Println(">>>> Running Application on Port", port)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: r,
	}

	// Listening signal
	stopSignal := make(chan os.Signal, 1)
	signal.Notify(stopSignal, os.Interrupt)

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			if err != http.ErrServerClosed {
				log.Fatal(err)
			}
		}
	}()

	<-stopSignal
	log.Println("Receiving signal to stop, shutting down service in 2sec")

	time.Sleep(2 * time.Second)

	// Timeout shutdown for 2 secs
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("Service exited, bye bye...")
}
