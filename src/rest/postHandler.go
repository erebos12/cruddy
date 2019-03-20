package rest

import (
	"fmt"
	"net/http"

	"model"

	"github.com/gin-gonic/gin"

	"infrastructure"
	. "utils"
)

// PostHandler godoc
// @Summary add a user
// @Description add by json user
// @Tags users
// @Accept json
// @Produce json
// @Param user body model.User true "Add user"
// @Success 200 {object} model.User string "user successfully created"
// @Failure 400 {object} model.Message string "error reading message"
// @Router /users [post]
func PostHandler(c *gin.Context) {
	method := "PostHandler"
	var user model.User
	err := c.BindJSON(&user)
	if err != nil {
		message := fmt.Sprintf("Error while unmarshalling json: '%+v'", err)
		LogIt(ERROR, method, message)
		c.JSON(http.StatusBadRequest, gin.H{"message": message})
	} else {
		infrastructure.Insert(user)
		c.JSON(http.StatusOK, user)
	}
}
