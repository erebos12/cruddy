package rest

import (
	"fmt"
	"net/http"

	. "../infrastructure"
	. "../model"
	. "../utils"

	"github.com/gin-gonic/gin"
)

func DeleteGenericAttributeHandler(c *gin.Context) {
	method := "DeleteGenericAttributeHandler"
	attribute := c.Query("attribute")
	value := c.Query("value")
	parameter := fmt.Sprintf("attribute=%s, value=%s", attribute, value)
	LogIt(INFO, method, fmt.Sprintf("DeleteGenericAttributeHandler %s", parameter))
	err := DeleteByGenericAttr(attribute, value)
	if err != nil {
		msg := fmt.Sprintf("Cannot delete entity with %s", parameter)
		LogIt(ERROR, method, msg)
		c.JSON(http.StatusNotFound, Message{Message: msg})
	} else {
		msg := fmt.Sprintf("Successfully deleted entity with %s", parameter)
		LogIt(INFO, method, msg)
		c.JSON(http.StatusOK, Message{Message: msg})
	}
}

func DeleteAllHandler(c *gin.Context) {
	method := "DeleteAllHandler"
	info, err := DeleteAll()
	if err != nil {
		msg := fmt.Sprintf("Cannot delete all entries")
		LogIt(ERROR, method, msg)
		c.JSON(http.StatusNotFound, Message{Message: msg})
	} else {
		msg := fmt.Sprintf("Successfully deleted all entities")
		LogIt(INFO, method, msg)
		c.JSON(http.StatusOK, Message{Message: info})
	}
}
