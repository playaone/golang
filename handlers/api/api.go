package api

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func New(version string) {
	r := gin.Default()

	AuthRoutes(r, version)
	UsersRoutes(r, version)

	err := r.Run(":8080")

	if err != nil {
		log.Fatalf("error starting server %v", err)
	}

	fmt.Println("Listening on port 8080...")
}
