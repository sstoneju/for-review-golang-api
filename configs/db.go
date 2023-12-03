package configs

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

func fsClient() *firestore.Client {
	// Load enviroment variable
	LoadEnvVariableFromDot()

	ctx := context.Background()

	// Get 환경 별로 firestore 인증서를 다르게 가지고 옴.
	sa := option.WithCredentialsFile(fmt.Sprintf("configs/review-db-%s.json", os.Getenv("ENV")))
	client, err := firestore.NewClient(ctx, os.Getenv("firestore_project_id"), sa)

	if err != nil {
		log.Fatalf("Failed to create Firestore client: %v", err)
	}
	return client
}

var Client *firestore.Client = fsClient()
