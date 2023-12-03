package service

import (
	"context"
	"edu-bridge/edu-app-api/models"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
)

type FirestoreService struct {
	collection *firestore.CollectionRef
}

type StoreService interface {
	Add(ctx context.Context, m interface{}) (interface{}, error)
	Get(ctx context.Context, id string) (interface{}, error)
	// List(ctx context.Context, m interface{}) (interface{}, error)
	// Update(ctx context.Context, id string) (interface{}, error)
	// Remove(ctx context.Context, id string) (interface{}, error)
}

func NewFirestoreService(client *firestore.Client, collectionName string) FirestoreService {
	log.Printf("Connected to FireStore(%s)", collectionName)
	return FirestoreService{
		collection: client.Collection(collectionName),
	}
}

func (s FirestoreService) Add(ctx context.Context, m interface{}) (interface{}, error) {
	model, ok := m.(models.Base)
	if !ok {
		return models.Base{}, fmt.Errorf("type assertion failed: q is not a models.Base")
	}

	newDoc := s.collection.NewDoc()
	model.Id = newDoc.ID

	_, err := newDoc.Create(ctx, model)
	if err != nil {
		log.Println("error: ", err)
		return models.Base{}, fmt.Errorf("firestoreService: func Add failure")
	}
	return model, nil
}

func (s FirestoreService) Get(ctx context.Context, id string) (interface{}, error) {
	// Query for "student" type users
	q := s.collection.Doc(id)

	// Execute the query
	doc, err := q.Get(ctx)
	if err != nil {
		return models.Base{}, fmt.Errorf("questionService: func Get failure")
	}

	// Process the results
	var question models.Base
	if err := doc.DataTo(&question); err != nil {
		return models.Base{}, fmt.Errorf("questionService: result transfer failure")
	}
	return question, nil
}

// func (s FirestoreService) List(ctx context.Context, m interface{}) (interface{}, error) {
// 	return (interface{}, error)
// }

// func (s FirestoreService) Update(ctx context.Context, m interface{}) (interface{}, error) {
// 	return (interface{}, error)
// }

// func (s FirestoreService) Remove(ctx context.Context, m interface{}) (interface{}, error) {
// 	return (interface{}, error)
// }
