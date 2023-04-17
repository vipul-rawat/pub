// Sample pubsub-quickstart creates a Google Cloud Pub/Sub topic.
package main

import (
	"context"
	"pub/http"

	"cloud.google.com/go/pubsub"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

func main() {
	app := gofr.New()
	app.Server.ValidateHeaders = false

	h := http.New(app)

	app.POST("/events", h.PublishEvent)

	app.Start()

}

func create(app *gofr.Gofr, projectID, topicID string) error {
	// projectID := "my-project-id"
	// topicID := "my-topic"
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		app.Logger.Errorf("pubsub.NewClient: %v", err)
		return err
	}
	defer client.Close()

	t, err := client.CreateTopic(ctx, topicID)
	if err != nil {
		app.Logger.Errorf("CreateTopic: %v", err)
		return err
	}
	app.Logger.Infof("Topic created: %v\n", t)
	return nil
}
