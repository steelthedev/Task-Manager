package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	task "github.com/steelthedev/task-go/pkg/task/routes"
	"go.uber.org/zap"
)

var (
	logger zap.Logger
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"message": "Welcome to Akin Task Manager v1",
			"status":  true,
		})
	})
	task.RegisterRoutes(router)

	router.Use(func(c *gin.Context) {
		cors.New(cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
			AllowedHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
		}).ServeHTTP(c.Writer, c.Request, func(w http.ResponseWriter, r *http.Request) {
		})
	})

	if err := router.Run(":8080"); err != nil {
		logger.Error(err.Error())
	}
}
