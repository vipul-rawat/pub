// Sample pubsub-quickstart creates a Google Cloud Pub/Sub topic.
package main

import (
	"context"
	"pub/eventhandlers"
	"pub/http"

	"cloud.google.com/go/pubsub"
	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()
	app.Server.ValidateHeaders = false

	//create(app, "", "")
	h := http.New(app)

	app.POST("/events", h.PublishEvent)
	app.POST("/any", h.Publish)

	// create(app, app.Config.Get("PUBSUB_PROJECT_ID"), app.Config.Get("TOPIC_NAME"))
	sub := eventhandlers.New(app, app.PubSub)

	go sub.Start()

	app.Start()

}

func create(app *gofr.Gofr, projectID, topicID string) error {
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
