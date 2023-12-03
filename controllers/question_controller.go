package controllers

import (
	"context"
	"edu-bridge/edu-app-api/models"
	"edu-bridge/edu-app-api/responses"
	"edu-bridge/edu-app-api/service"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateQuestion() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("CreateQuestion ...")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var body models.Question

		// 요청 본문의 JSON을 newMember 구조체로 바인딩
		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 유효성 검증
		validate := validator.New()
		if err := validate.Struct(body); err != nil {
			// 유효성 검사 실패 처리
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Question add 작업
		result, err := service.QuestionSvc.Add(ctx, body)
		if err != nil {
			// 유효성 검사 실패 처리
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 타입 단언을 사용하여 result를 models.Question 타입으로 변환
		question, ok := result.(models.Question)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Type assertion failed"})
			return
		}

		c.JSON(http.StatusCreated, responses.QuestionResponse{Status: http.StatusCreated, Message: "success", Data: []models.Question{question}})
	}
}

func GetQuestion() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("GetQuestion ...")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		userId := c.Param("userId")

		result, err := service.QuestionSvc.Get(ctx, userId)
		if err != nil {
			// data featch is failure
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 타입 단언을 사용하여 result를 models.Question 타입으로 변환
		question, ok := result.(models.Question)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Type assertion failed"})
			return
		}

		c.JSON(http.StatusOK, responses.QuestionResponse{Status: http.StatusOK, Message: "success", Data: []models.Question{question}})
	}
}
