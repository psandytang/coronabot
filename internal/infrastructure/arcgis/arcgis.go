package arcgis

import (
	"encoding/json"
	"net/http"

	"github.com/disiqueira/coronabot/internal/domain/model"
)

type (
	CoranaREST struct {
		httpClient *http.Client
	}

	arcgisReponse []struct {
		CountryRegion string `json:"country"`
		Confirmed     int    `json:"cases"`
		NewCases      int    `json:"todayCases"`
		Deaths        int    `json:"deaths"`
		NewDeaths     int    `json:"todayDeaths"`
		Recovered     int    `json:"recovered"`
		ActiveCases   int    `json:"active"`
		Critical      int    `json:"critical"`
	}
)

const (
	arcgisURL = "https://coronavirus-19-api.herokuapp.com/countries"
)

func New(httpClient *http.Client) *CoranaREST {
	return &CoranaREST{
		httpClient: httpClient,
	}
}

func (c *CoranaREST) StatusPerCountry() ([]model.Status, error) {
	req, err := http.NewRequest("GET", arcgisURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	var response arcgisReponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	statusList := make([]model.Status, len(response))
	for i, feature := range response {
		statusList[i] = model.NewStatus(
			feature.CountryRegion,
			feature.Confirmed,
			feature.NewCases,
			feature.Deaths,
			feature.NewDeaths,
			feature.Recovered,
			feature.ActiveCases,
			feature.Critical,
		)
	}

	return statusList, nil
}
