package main

import (
	"fmt"
	"os"

	"github.com/arisnacg/go-graphql-example/graph"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(cors.Default())
	router.POST("/todo", graph.TodoGraphRouter)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
	fmt.Println("Server is running: http://localhost:" + port)
}
