package chitaigorod

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/pkg/errors"
)

const (
	chitaiGorodAPIHost = "web-gate.chitai-gorod.ru"
	chitaigorodPath    = "/api/v2/products"
	batch              = 20 //perPage
)

// URL query params
// perPage also can be used as batch size in GetBooks() function
var (
	customerCityIdParam = 213 // customerCityIdParam=...
	// products[page]=...
	// products[per-page]=...
	sortPresetParam = "relevance"                                                                                   //sortPreset=...
	includeParam    = "productTexts,publisher,publisherBrand,publisherSeries,dates,literatureWorkCycle,rating,tags" // include=...
)

func (c *Client) GetBooks(count int) ([]Data, error) {
	v := make(url.Values)
	v.Set("customerCityId", strconv.Itoa(customerCityIdParam))
	v.Set("sortPreset", sortPresetParam)
	v.Set("include", includeParam)

	receivedBooks := make([]Data, 0, count)
	for i := 0; i < count; i += batch {
		perPage := 20
		if i+batch > count {
			perPage = count - i
		}

		newBooks, err := c.getBooks(v, i/20+1, perPage)
		if err != nil {
			return nil, errors.Wrap(err, "error during receiving new books batch")
		}

		receivedBooks = append(receivedBooks, newBooks...)
	}

	return receivedBooks, nil
}

func (c *Client) getBooks(v url.Values, page, perPage int) ([]Data, error) {
	v.Set("products[page]", strconv.Itoa(page))
	v.Set("products[perPage]", strconv.Itoa(perPage))

	u := url.URL{
		Scheme:   "https",
		Path:     chitaigorodPath,
		Host:     chitaiGorodAPIHost,
		RawQuery: v.Encode(),
	}
	perform := u.String()

	req, err := http.NewRequest(http.MethodGet, perform, nil)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create new http request")
	}
	req.Header.Set("Authorization", c.token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "cannot make http request to chitaigorod api")
	}
	defer resp.Body.Close()

	rawData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "cannot read response body")
	}

	var response Response

	err = json.Unmarshal(rawData, &response)
	if err != nil {
		return nil, errors.Wrap(err, "cannot unmarshal response body")
	}

	// if we need less than min perPage value (20)
	return response.Data[:perPage], nil
}
