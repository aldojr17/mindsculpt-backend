package api

import (
	"encoding/json"
	"fmt"
	"time"

	"mindsculpt/domain"
	"mindsculpt/utils"
)

func GetGenerationModels() (*domain.APIGetModelsResponse, error) {
	url := utils.GetUrl(PATH_GET_MODELS)

	req, err := NewGetRequest(url)
	if err != nil {
		return nil, err
	}

	res, err := DoRequest(req)
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

func GenerateImage(payload domain.APIGenerateImageRequest) (*domain.APIGetGenerationRawResponse, error) {
	url := utils.GetUrl(PATH_CREATE_GENERATION)

	body, writer, err := GenerateHeader(payload)
	if err != nil {
		return nil, err
	}

	req, err := NewPostRequest(url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	res, err := DoRequest(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var raw *domain.APIGenerateImageRawResponse

	if res.StatusCode != 201 {
		return &domain.APIGetGenerationRawResponse{}, fmt.Errorf("error")
	}

	if err := json.NewDecoder(res.Body).Decode(&raw); err != nil {
		return nil, err
	}

	return GetGeneration(*raw)
}

func GetGeneration(payload domain.APIGenerateImageRawResponse) (*domain.APIGetGenerationRawResponse, error) {
	formattedUrl := fmt.Sprintf(PATH_GET_GENERATION, payload.UUID)
	url := utils.GetUrl(formattedUrl)

	req, err := NewGetRequest(url)
	if err != nil {
		return nil, err
	}

	for {
		res, err := DoRequest(req)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		var raw *domain.APIGetGenerationRawResponse

		if err := json.NewDecoder(res.Body).Decode(&raw); err != nil {
			return nil, err
		}

		if raw.Status == STATUS_DONE {
			return raw, nil
		}

		time.Sleep(TIMEOUT * time.Second)
	}
}
