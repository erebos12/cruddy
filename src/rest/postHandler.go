package rest

import (
	"fmt"
	"net/http"

	. "../infrastructure"
	. "../model"
	. "../utils"

	"github.com/gin-gonic/gin"
)

func PostHandler(c *gin.Context) {
	method := "PostHandler"
	var user User
	err := c.BindJSON(&user)
	if err != nil {
		message := fmt.Sprintf("Error while unmarshalling json: '%+v'", err)
		LogIt(ERROR, method, message)
		c.JSON(http.StatusBadRequest, gin.H{"message": message})
	} else {
		Insert(user)
		c.JSON(http.StatusOK, user)
	}
}
