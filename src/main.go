package main

import (
	. "./configuration"
	. "./rest"
	"github.com/gin-gonic/gin"
)

// @title Cruddy Swagger UI
// @version 1.0
// @description This is CRUDDY

// @host localhost:8080
// @BasePath /
func GetRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", SayHelloHandler)
	router.GET("/users/all", GetAllHandler)
	router.GET("/users", GetGenericAttributeHandler)
	router.POST("/users", PostHandler)
	router.DELETE("/users", DeleteGenericAttributeHandler)
	router.DELETE("/users/all", DeleteAllHandler)
	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}

func main() {
	r := GetRouter()
	r.Run(":" + GetPort())
}
