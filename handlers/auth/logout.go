package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LogoutHandler(c *gin.Context) {
	fmt.Println("You visited the logout endpoint")
	c.JSON(http.StatusOK, "You visited the logout endpoint")
}
