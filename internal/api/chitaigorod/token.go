// Online store API needs token of Bearer format to have access to store information
package chitaigorod

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
)

var (
	chitaiGorodHost = "chitai-gorod.ru"
	// name for the file where auth token will be cached
	tokenFile = "token.txt"
)

// NewAuthToken checks whether jwt auth token lies in token.txt file
// If it doesn't exist, it receives token by making request to chitaigorod main site
// This is done in order to avoid frequent visits to the site and as a result not get banned
func newAuthToken() (string, error) {
	const op = "chitaigorod.AuthToken"

	oldTokenBytes, err := os.ReadFile(tokenFile)
	if errors.Is(err, os.ErrNotExist) {
		newToken, err := getTokenFromRequest()
		if err != nil {
			return "", errors.Wrap(err, op)
		}

		err = os.WriteFile(tokenFile, []byte(newToken), 0755)
		if err != nil {
			return "", errors.Wrap(err, op)
		}

		return formatToken(newToken), nil
	}

	newToken := string(oldTokenBytes)

	expTime, err := parseAuthToken(newToken)
	if err != nil {
		return "", errors.Wrap(err, op)
	}

	now := time.Now().Unix()
	if now > expTime+5*60 {
		newToken, err = getTokenFromRequest()
		if err != nil {
			return "", errors.Wrap(err, op)
		}
		err = os.WriteFile(tokenFile, []byte(newToken), 0755)
		if err != nil {
			return "", errors.Wrap(err, op)
		}
	}

	return formatToken(newToken), nil
}

// parseAuthToken parses jwt token and returns the timestamp after which it will expire
func parseAuthToken(rawToken string) (int64, error) {
	const op = "chitaigorod.ParseAuthToken"

	tokenString := strings.Replace(formatToken(rawToken), "Bearer ", "", 1)
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return 0, errors.Wrap(err, op)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return int64(claims["exp"].(float64)), nil
	}

	return 0, errors.Wrap(fmt.Errorf("no expiration date info in jwt token"), op)
}

func getTokenFromRequest() (string, error) {
	const op = "chitaigorod.getTokenFromRequest"

	u := url.URL{
		Scheme: "https",
		Host:   chitaiGorodHost,
	}

	resp, err := http.Get(u.String())
	if err != nil {
		return "", errors.Wrap(err, op)
	}

	var headersList []string
	for key, list := range resp.Header {
		if key == "Set-Cookie" {
			headersList = list
		}
	}

	var token string
	for _, val := range headersList {
		if strings.Contains(val, "access-token") {
			token = val
		}
	}

	return token, nil
}

func formatToken(rawToken string) string {
	tokenString := strings.Split(rawToken, " ")[0]
	tokenString = tokenString[strings.Index(tokenString, "Bearer"):]
	tokenString = strings.Replace(tokenString, "%20", " ", 1)
	tokenString = tokenString[:len(tokenString)-1]

	return tokenString
}
