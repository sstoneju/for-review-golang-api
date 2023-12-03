package routes

import (
	"edu-bridge/edu-app-api/controllers"

	"github.com/gin-gonic/gin"
)

func QuestionRoute(router *gin.RouterGroup) {
	router.POST("/question", controllers.CreateQuestion())         //add this
	router.GET("/question/:questionId", controllers.GetQuestion()) //add this
}
