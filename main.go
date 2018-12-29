package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

// Run server: go run -race main.go
// Try requests: curl http://127.0.0.1:8080
func main() {
	r := setupRouter()
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", home)

	return r
}

func home(c *gin.Context) {
	c.String(200, "Hello to DevHouse")
}
