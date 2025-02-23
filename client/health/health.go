package health

import (
	"net/http"

	"go.uber.org/fx"
)

type ClientParams struct {
	fx.In
}

type Client struct {
	client *http.Client
}

func NewClient(params ClientParams) *Client {
	return &Client{
		client: &http.Client{},
	}
}
