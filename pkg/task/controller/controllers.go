package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/steelthedev/task-go/pkg/task/services"
	"gorm.io/gorm"
)

var (
	ControllerHandler handler
)

type handler struct {
	db *gorm.DB
}

func (h handler) CreateTask(ctx *gin.Context) {
	_, err := services.CreateTask(ctx, h.db)
	if err != nil {
		ctx.AbortWithStatusJSON(err.StatusCode, gin.H{
			"message": err.Message,
		})
	}
	ctx.IndentedJSON(http.StatusCreated, gin.H{
		"message": "Task created successfully",
	})

}
