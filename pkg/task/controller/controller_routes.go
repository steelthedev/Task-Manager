package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := router.Group("task")

	routes.GET("/tasks", h.GetTasks)
	routes.POST("/create-task", h.CreateTask)
}
