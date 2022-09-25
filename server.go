package main

import (
	controller "ark_backend_go/controllers"
	"ark_backend_go/database"
	"ark_backend_go/graph"
	"ark_backend_go/graph/generated"
	"ark_backend_go/helper"
	repository "ark_backend_go/repositories"
	service "ark_backend_go/services"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/gorilla/websocket"
)

const defaultPort = "8080"

var (
	validate = validator.New()
	db       = database.SetupGetConnection()

	paymentRepository = repository.NewPaymentRepository()
	paymentService    = service.NewPaymentService(paymentRepository, db, validate)
	paymentController = controller.NewPaymentController(paymentService)
)

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	resol := graph.NewResolver(paymentService)

	return func(c *gin.Context) {
		srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resol}))
		srv.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	r := gin.Default()
	r.Static("assets", "./assets")
	r.POST("/query", graphqlHandler())
	r.GET("/graphql", playgroundHandler())

	v1 := r.Group("/api/v1")
	{
		auth := v1.Group("/payment")
		{
			auth.GET("/method", paymentController.GetPaymentMethod)
		}
	}

	r.GET("/websoc", func(ctx *gin.Context) {
		conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			log.Print("upgrade:", err)
		}

		defer conn.Close()

		for i := 0; i < 1000; i++ {
			time.Sleep(time.Second * 5)
			paymentResponse := paymentService.GetPaymentMethod()
			conn.WriteJSON(helper.DefaultResponse{
				Code:    http.StatusOK,
				Message: "Payment method has been listed",
				Data:    paymentResponse,
				Status:  true})
		}
	})

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	r.Run(":" + port)
}