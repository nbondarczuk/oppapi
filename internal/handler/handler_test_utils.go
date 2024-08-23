package handler

import "github.com/gin-gonic/gin"

func testRequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func testResponseLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
