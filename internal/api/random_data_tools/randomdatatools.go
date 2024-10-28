package randomdatatools

import (
	"db-coursework/internal/models"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/pkg/errors"
)

var (
	apiParams = "LastName,FirstName,FatherName,Gender,DateOfBirth,Phone,Email,Address"
	apiURL    = "https://api.randomdatatools.ru"
)

func GetCustomers(count int) ([]models.Customer, error) {
	result := make([]models.Customer, count)

	params := url.Values{}
	params.Add("count", strconv.Itoa(count))
	params.Add("params", apiParams)

	url := apiURL + "?" + params.Encode()

	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "error during request to random data tools api")
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "error during reading request body")
	}

	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, errors.Wrap(err, "error during unmarshalling request")
	}

	return result, nil
}
