package responses

import "edu-bridge/edu-app-api/models"

type QuestionResponse struct {
	Status  int               `json:"status"`
	Message string            `json:"message"`
	Data    []models.Question `json:"data"`
}
