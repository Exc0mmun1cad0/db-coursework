package chitaigorod

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
)

var (
	chitaiGorodURL   = "https://chitai-gorod.ru"
)

// Online store API needs token of Bearer format to have access to store information
func NewAuthToken() (string, error) {
	const op = "chitaigorod.AuthToken"

	resp, err := http.Get(chitaiGorodURL)
	if err != nil {
		return "", errors.Wrap(err, op)
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

	authToken := strings.Split(rawToken, " ")[0]
	authToken = authToken[strings.Index(authToken, "Bearer"):]

	authToken = strings.Replace(authToken, "%20", " ", 1)
	authToken = authToken[:len(authToken)-1]

	return authToken, nil
}

// ParseAuthToken parses jwt token and returns the timestamp after which it will expire
func ParseAuthToken(tokenString string) (int64, error) {
	const op = "chitaigorod.ParseAuthToken"

	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return 0, errors.Wrap(err, op)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return int64(claims["exp"].(float64)), nil
	}

	return 0, errors.Wrap(fmt.Errorf("no expiration date info in jwt token"), op)
}
