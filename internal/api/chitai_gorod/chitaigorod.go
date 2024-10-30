package chitaigorod

import (
	"fmt"
	"net/http"
	"strings"
)

var (
	chitaiGorodURL = "https://chitai-gorod.ru"
)

// Online store API needs token of Bearer format to have access to store information
func AuthToken() (string, error) {
	const op = "chitaigorod.AuthToken"

	resp, err := http.Get(chitaiGorodURL)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	var headersList []string
	for key, list := range resp.Header {
		if key == "Set-Cookie" {
			headersList = list
		}
	}

	var rawToken string
	for _, val := range headersList {
		if strings.Contains(val, "access-token") {
			rawToken = val
		}
	}

	accessToken := strings.Split(rawToken, " ")[0]
	authToken := accessToken[strings.Index(accessToken, "Bearer"):]

	return authToken, nil
}
