package utils

import "github.com/gin-gonic/gin"

func ApiSuccessIndented(ctx *gin.Context, code int, message string, data interface{}) *gin.Context {
	ctx.IndentedJSON(code, gin.H{
		"message": message,
		"data":    data,
	})
	return nil
}
