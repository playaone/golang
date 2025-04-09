package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	fmt.Println("You visited the login endpoint")
	c.JSON(http.StatusOK, "You visited the login endpoint")
}
