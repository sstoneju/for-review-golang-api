package service

import (
	"context"
	"edu-bridge/edu-app-api/configs"
	"edu-bridge/edu-app-api/models"
	"fmt"
)

type QuestionService struct {
	fs FirestoreService // 포인터 형 변수로 사용한다.
}

// QuestionService의 포인터형 변수를 return 한다.
func NewQuestionService(fs FirestoreService) *QuestionService {
	// QuestionService의 주소를 return -> return이 포인터형 변수이기 때문
	return &QuestionService{
		fs: fs,
	}
}

func (s QuestionService) Add(ctx context.Context, m interface{}) (interface{}, error) {
	model, ok := m.(models.Question)
	if !ok {
		// q가 models.Question 타입이 아닐 때의 에러 처리
		return models.Question{}, fmt.Errorf("type assertion failed: q is not a models.Question")
	}

	result, err := s.fs.Add(ctx, model)

	return result, err
}

func (s QuestionService) Get(ctx context.Context, id string) (interface{}, error) {
	// Query for "student" type users
	q := s.fs.collection.Doc(id)

	// Execute the query
	doc, err := q.Get(ctx)
	if err != nil {
		return models.Question{}, fmt.Errorf("questionService: func Get failure")
	}

	// Process the results
	var question models.Question
	if err := doc.DataTo(&question); err != nil {
		return models.Question{}, fmt.Errorf("questionService: result transfer failure")
	}
	return question, nil
}

// func (s QuestionService) List(models.Question) models.Question {
// 	return models.Question{}
// }

// func (s QuestionService) Update(models.Question) models.Question {
// 	return models.Question{}
// }

//	func (s QuestionService) Remove(models.Question) models.Question {
//		return models.Question{}
//	}

var questionfs FirestoreService = NewFirestoreService(configs.Client, "question")
var QuestionSvc StoreService = NewQuestionService(questionfs)
