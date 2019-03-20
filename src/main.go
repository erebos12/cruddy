package main

import (
	"configuration"
	"rest"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "./docs" // docs is generated by Swag CLI, you have to import it.
)

// @title Cruddy Swagger UI
// @version 1.0
// @description This is CRUDDY

// @host localhost:8080
// @BasePath /
func GetRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", rest.SayHelloHandler)
	router.GET("/users/all", rest.GetAllHandler)
	router.GET("/users", rest.GetGenericAttributeHandler)
	router.POST("/users", rest.PostHandler)
	router.DELETE("/users", rest.DeleteGenericAttributeHandler)
	router.DELETE("/users/all", rest.DeleteAllHandler)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}

func main() {
	r := GetRouter()
	r.Run(":" + configuration.GetPort())
}