package task

import (
	"github.com/gin-gonic/gin"
	"github.com/steelthedev/task-go/pkg/task/controller"
)

func RegisterRoutes(router *gin.Engine) {

	h := controller.ControllerHandler

	routes := router.Group("task")
	routes.POST("/create-task", h.CreateTask)
}
