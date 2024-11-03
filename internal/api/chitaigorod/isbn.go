// It is impossible to extract ISBN from chitaigorod API so it will be done via main site
package chitaigorod

import (
	"io"
	"net/http"
	"net/url"
	"regexp"

	"github.com/pkg/errors"
)

func GetISBN(bookPath string) (string, error) {
	u := url.URL{
		Scheme: "https",
		Host:   chitaiGorodHost,
		Path:   bookPath,
	}
	resp, err := http.Get(u.String())
	if err != nil {
		return "", errors.Wrap(err, "error during request for isbn")
	}
	defer resp.Body.Close()

	page, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrap(err, "error during reading isbn response body")
	}

	re := regexp.MustCompile(`<meta data-n-head="ssr" data-hid="og:isbn" name="og:isbn" content="([^"]+)">`)
	match := re.FindStringSubmatch(string(page))
	if len(match) <= 1 {
		return "", errors.Wrap(err, "isbn not found on this page")
	}

	return match[1], nil
}
