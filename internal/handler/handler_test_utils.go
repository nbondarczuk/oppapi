package handler

import "github.com/gin-gonic/gin"

func TestRequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func TestResponseLogger() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
