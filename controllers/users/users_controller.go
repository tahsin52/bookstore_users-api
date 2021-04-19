package users

import (
	"github.com/gin-gonic/gin"
	"github.com/tahsin52/bookstore_users-api/domain/users"
	"github.com/tahsin52/bookstore_users-api/services"
	"github.com/tahsin52/bookstore_users-api/utils/errors"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context)  {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("Invalid User ID Should Be A Number")
		c.JSON(err.Status, err)
		return
	}
	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		// Handle User Creation Error
		return
	}
	c.JSON(http.StatusOK, user)}

func CreateUser(c *gin.Context)  {
	var user  users.User
	if err := c.ShouldBindJSON(&user); err != nil{
		// Handle User BadRequest to the caller.
		restErr := errors.NewBadRequestError("Invalid Json Body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		// Handle User Creation Error
		return
	}
	c.JSON(http.StatusCreated, result)
}

func SearchUser(c *gin.Context)  {
	c.String(http.StatusNotImplemented, "implement me!")

}