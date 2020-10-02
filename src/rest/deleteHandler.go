package rest

import (
	"fmt"
	"net/http"

	. "../infrastructure"
	. "../model"
	. "../utils"

	"github.com/gin-gonic/gin"
)

// DeleteGenericAttributeHandler godoc
// @Summary delete all users which match attribute and value
// @Description delete all users which match attribute and value
// @Tags users
// @Produce json
// @Param attribute query string true "attribute of a certain json"
// @Param value query string true "the value of the parameter attribute"
// @Success 200 {object} model.Message string "successfully deleted"
// @Failure 404 {object} model.Message string "could not delete user"
// @Router /users [delete]
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

// DeleteAllHandler godoc
// @Summary delete all users
// @Description delete all users
// @Tags users
// @Produce json
// @Success 200 {object} model.Message string "successfully deleted all users"
// @Failure 404 {object} model.Message string "could not delete all users"
// @Router /users/all [delete]
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
