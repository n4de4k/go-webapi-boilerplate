package rest

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	component2 "github.com/n4de4k/go-webapi-boilerplate/app/provider"
	"github.com/n4de4k/go-webapi-boilerplate/delivery/rest/internal"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Handler struct {
	serviceComponent component2.ServiceComponent
}

func NewHandler(serviceComponent component2.ServiceComponent) *Handler {
	return &Handler{serviceComponent}
}

func Init(handler *Handler) {
	env := os.Getenv("APP_ENV")
	port := os.Getenv("APP_PORT")
	if env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()


	r.Use(gin.Recovery())
	r.Use(internal.JSONAppErrorReporter())

	// Enable logger middleware with custom format
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


	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "You service is on",
		})
	})

	// --> Register all routes here
	r.POST("/login", handler.SignIn)

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