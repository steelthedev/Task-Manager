package utils

import "github.com/gin-gonic/gin"

type AppError struct {
	Message string
	Error   string
}

type ApiError struct {
	Message    string
	StatusCode int
	Error      string
}

func ApiErrorFunction(ctx *gin.Context, code int, message string, error string) *gin.Context {
	ctx.AbortWithStatusJSON(code, gin.H{
		"message": message,
		"error":   error,
	})
	return nil
}
