package eventhandlers

import (
	"gofr.dev/pkg/datastore/pubsub"
	"gofr.dev/pkg/gofr"
)

type Sub struct {
	app        *gofr.Gofr
	subscriber pubsub.PublisherSubscriber
}

func New(app *gofr.Gofr, subscriber pubsub.PublisherSubscriber) *Sub {
	return &Sub{
		app:        app,
		subscriber: subscriber,
	}
}

func (s *Sub) Start() {
	for {
		s.subscribe()
	}
}

func (s *Sub) subscribe() {
	//defer func() {
	//	if rec := recover(); rec != nil {
	//		s.app.Logger.Errorf("recovered from panic, panic: %v", rec)
	//	}
	//}()

	// Continuously listen to event hub and subscribe to new events
	_, err := s.subscriber.SubscribeWithCommit(s.process)
	if err != nil {
		s.app.Logger.Errorf("Error in subscribing to Event Hub, Error - %v", err.Error())
		return
	}
}

func (s *Sub) process(msg *pubsub.Message) (isConsumed, isContinue bool) {
	s.app.Logger.Infof(msg.Value)
	return false, true
}
