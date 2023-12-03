package main

import (
	"edu-bridge/edu-app-api/routes"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/basic/docs"
)

func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

func main() {
	router := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"

	// example routes
	eg := router.Group("/example")
	{
		eg.GET("/helloworld", Helloworld)
	}

	// routes
	v1_group := router.Group("/api/v1")
	routes.QuestionRoute(v1_group) //add this

	router.Run("localhost:8080")
}
