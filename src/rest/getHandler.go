package rest

import (
	"fmt"
	"net/http"

	. "../infrastructure"
	. "../model"
	. "../utils"

	"github.com/gin-gonic/gin"
)

// SayHelloHandler godoc
// @Summary says hello
// @Tags /
// @Description says hello
// @Produce json
// @Success 200 {object} model.Message string "hello message"
// @Router / [get]
func SayHelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello! I am the go app running for you :-)"})
}

// GetAllHandler godoc
// @Summary get all user
// @Tags users
// @Description get all user
// @Produce json
// @Success 200 {array} model.User "user found"
// @Failure 404 {object} model.Message string "no results"
// @Router /users/all [get]
func GetAllHandler(c *gin.Context) {
	results := GetAll()
	if len(results) == 0 {
		msg := Message{Message: "no results"}
		c.JSON(http.StatusOK, msg)
	} else {
		c.JSON(http.StatusOK, results)
	}
}

// GetGenericAttributeHandler godoc
// @Summary get all users that matches attribute and value
// @Tags users
// @Produce json
// @Param attribute query string true "attribute of a certain json"
// @Param value query string true "the value of the parameter attribute"
// @Success 200 {array} model.User "user found"
// @Failure 404 {object} model.Message string "user not found"
// @Failure 503 {object} model.Message string "server error while getting user"
// @Router /users [get]
func GetGenericAttributeHandler(c *gin.Context) {
	method := "GetGenericAttributeHandler"
	attribute := c.Query("attribute")
	value := c.Query("value")
	LogIt(INFO, method, fmt.Sprintf("attribute=%s, value=%s", attribute, value))
	results, err := FetchByGenericAttr(attribute, value)
	if err != nil {
		message := fmt.Sprintf("Error while getting user: '%+v'", err)
		LogIt(ERROR, method, message)
		c.JSON(http.StatusServiceUnavailable, gin.H{"message": message})
	} else if len(results) == 0 {
		message := fmt.Sprintf("No results found for attribute '%s' and value '%s'", attribute, value)
		LogIt(ERROR, method, message)
		c.JSON(http.StatusNotFound, gin.H{"message": message})
	} else {
		c.JSON(http.StatusOK, results)
	}
}
