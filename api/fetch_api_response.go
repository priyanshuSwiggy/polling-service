package api

import (
	"encoding/json"
	"net/http"
	"polling-service/util"
)

type Response struct {
	Rates map[string]float64 `json:"rates"`
}

func FetchRates() (map[string]float64, error) {
	resp, err := http.Get(util.AppConfig.API.Url + "?app_id=" + util.AppConfig.API.Key)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var apiResponse Response
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, err
	}
	return apiResponse.Rates, nil
}
