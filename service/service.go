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
