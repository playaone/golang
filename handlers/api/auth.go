package api

import (
	"github.com/gin-gonic/gin"
	"github.com/playaone/golang/handlers/auth"
)

func AuthRoutes(r *gin.Engine, version string) {
	g := r.Group(version + "/auth")

	g.GET("/login", auth.LoginHandler)
	g.GET("/logout", auth.LogoutHandler)
}
