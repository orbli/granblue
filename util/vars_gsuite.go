package util

import (
	"cloud.google.com/go/firestore"
	"context"
	"log"
)

var (
	projectID        string = "granblue-247222"
	FirestoreClient  *firestore.Client
	FirestoreContext context.Context
)

func init() {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	FirestoreClient = client
	FirestoreContext = ctx
}
