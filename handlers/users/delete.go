package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteUserHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "You visited the delete user endpoint")
}
