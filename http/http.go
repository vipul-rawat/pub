package http

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"encoding/json"
)

type Handler struct {
	app *gofr.Gofr
}

func New(app *gofr.Gofr) *Handler {
	return &Handler{
		app: app,
	}
}

type NewCustomerEvent struct {
	CustomerID     int    `json:"customerId"`
	OrganizationID string `json:"organizationId"`
	ReferralCode   string `json:"referralCode"`
}

func (h *Handler) PublishEvent(ctx *gofr.Context) (interface{}, error) {
	var ev NewCustomerEvent
	err := ctx.Bind(&ev)
	if err != nil {
		return nil, err
	}

	marshal, _ := json.Marshal(&ev)
	if err != nil {
		return nil, err
	}

	err = h.app.PubSub.PublishEvent("test", marshal, map[string]string{})
	if err != nil {
		h.app.Logger.Errorf("Error publishing event")
	}

	return nil, nil
}
