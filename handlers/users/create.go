package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUserHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "You visited the create user endpoint")
}
