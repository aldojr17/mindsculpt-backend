package api

import (
	"encoding/json"
	"fmt"
	"mindsculpt/config"
	"mindsculpt/domain"
	"net/http"
)

func GetGenerationModels() (*domain.APIGetModelsResponse, error) {
	url := fmt.Sprintf("%s%s", config.GetConfig().API.URL, PATH_GET_MODELS)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add(KEY_API_HEADER_X_KEY, config.GetConfig().API.GetHeaderKey())
	req.Header.Add(KEY_API_HEADER_X_SECRET, config.GetConfig().API.GetHeaderSecret())

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var response []*domain.APIGetModelsResponse

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}

	return response[0], nil
}
