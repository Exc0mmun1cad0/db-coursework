package chitaigorod

import (
	"net/http"
)

type Client struct {
	token      string // chitai gorod auth token
	httpClient *http.Client
}

func NewClient(authToken string) *Client {
	return &Client{
		token:      authToken,
		httpClient: &http.Client{},
	}
}
