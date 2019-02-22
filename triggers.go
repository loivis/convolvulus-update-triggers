package convvls

import (
	"context"
	"log"
	"os"

	pubsub "cloud.google.com/go/pubsub/apiv1"
)

var projectID string
var pubClient *pubsub.PublisherClient
var topic string = "booksToUpdate"

func init() {
	projectID = os.Getenv("GCP_PROJECT")

	var err error
	ctx := context.Background()
	pubClient, err = pubsub.NewPublisherClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
