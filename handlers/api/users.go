package api

import (
	"github.com/gin-gonic/gin"
	"github.com/playaone/golang/handlers/users"
)

func UsersRoutes(r *gin.Engine, version string) {
	g := r.Group(version + "/users")
	g.POST("/create", users.CreateUserHandler)
	g.DELETE("/delete", users.DeleteUserHandler)
}
