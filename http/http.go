package http

import (
	"encoding/json"

	"gofr.dev/pkg/gofr"
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
	CustomerID     int         `json:"customerId"`
	OrganizationID interface{} `json:"organizationId"`
	ReferralCode   string      `json:"referralCode"`
}

type OrderStatus struct {
	ReferenceNumber int    `json:"referenceNumber"`
	OrganizationId  int    `json:"organizationId"`
	OldStatus       string `json:"oldStatus"`
	NewStatus       string `json:"newStatus"`
	StoreId         int    `json:"storeId"`
	Communication   bool   `json:"communication"`
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

func (h *Handler) Publish(ctx *gofr.Context) (interface{}, error) {
	var order OrderStatus
	err := ctx.Bind(&order)
	if err != nil {
		return nil, err
	}

	marshal, _ := json.Marshal(&order)
	if err != nil {
		return nil, err
	}

	err = h.app.PubSub.PublishEvent("test", marshal, map[string]string{})
	if err != nil {
		h.app.Logger.Errorf("Error publishing event")
	}

	return nil, nil
}
