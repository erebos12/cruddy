package rest

import (
	"fmt"
	"net/http"

	. "../infrastructure"
	. "../model"
	. "../utils"

	"github.com/gin-gonic/gin"
)

func SayHelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello! I am the go app running for you :-)"})
}

func GetAllHandler(c *gin.Context) {
	results := GetAll()
	if len(results) == 0 {
		msg := Message{Message: "no results"}
		c.JSON(http.StatusOK, msg)
	} else {
		c.JSON(http.StatusOK, results)
	}
}

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
