package chitaigorod

import (
	"net/http"

	"github.com/pkg/errors"
)

type Client struct {
	token      string // chitai gorod auth token
	httpClient *http.Client
}

func NewClient() (*Client, error) {
	authToken, err := newAuthToken()
	if err != nil {
		return nil, errors.Wrap(err, "cannot receive auth token for chitai gorod API")
	}
	return &Client{
		token:      authToken,
		httpClient: &http.Client{},
	}, nil
}
