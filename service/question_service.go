package service

import (
	"context"
	"edu-bridge/edu-app-api/configs"
	"edu-bridge/edu-app-api/models"
	"fmt"
	"log"
)

type QuestionService struct {
	fs FirestoreService // 포인터 형 변수로 사용한다.
}

type StoreService interface {
	Add(ctx context.Context, m models.Question) (models.Question, error)
	Get(ctx context.Context, id string) (models.Question, error)
}

// QuestionService의 포인터형 변수를 return 한다.
func NewQuestionService(fs FirestoreService) *QuestionService {
	// QuestionService의 주소를 return -> return이 포인터형 변수이기 때문
	return &QuestionService{
		fs: fs,
	}
}

func (s QuestionService) Add(ctx context.Context, m models.Question) (models.Question, error) {
	newDoc := s.fs.collection.NewDoc()
	m.Id = newDoc.ID

	_, err := newDoc.Create(ctx, m)
	if err != nil {
		log.Println("error: ", err)
		return models.Question{}, fmt.Errorf("firestoreService: func Add failure")
	}
	return m, nil
}

func (s QuestionService) Get(ctx context.Context, id string) (models.Question, error) {
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

var questionfs FirestoreService = NewFirestoreService(configs.Client, "question")
var QuestionSvc StoreService = NewQuestionService(questionfs)
