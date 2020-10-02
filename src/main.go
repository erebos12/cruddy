package main

import (
	. "./configuration"
	. "./rest"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", SayHelloHandler)
	router.GET("/users/all", GetAllHandler)
	router.GET("/users", GetGenericAttributeHandler)
	router.POST("/users", PostHandler)
	router.DELETE("/users", DeleteGenericAttributeHandler)
	router.DELETE("/users/all", DeleteAllHandler)
	return router
}

func main() {
	r := GetRouter()
	r.Run(":" + GetPort())
}
