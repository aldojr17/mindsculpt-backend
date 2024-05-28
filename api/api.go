package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"time"

	"mindsculpt/config"
	"mindsculpt/domain"
	"net/http"
	"net/textproto"
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

func GenerateImage(payload domain.APIGenerateImageRequest) (*domain.APIGetGenerationRawResponse, error) {
	url := fmt.Sprintf("%s%s", config.GetConfig().API.URL, PATH_CREATE_GENERATION)

	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	paramsPartHeader := textproto.MIMEHeader{}
	paramsPartHeader.Set("Content-Disposition", `form-data; name="params"`)
	paramsPartHeader.Set("Content-Type", "application/json")

	paramsWriter, err := w.CreatePart(paramsPartHeader)
	if err != nil {
		return nil, err
	}

	paramsRaw := `{
		"type": "GENERATE",
		"width": %d,
		"height": %d,
		"num_images": 1,
		"negativePromptUnclip": "%s",
		"generateParams": {
			"query": "%s"
		}
	}`

	params := fmt.Sprintf(paramsRaw, payload.Width, payload.Height, payload.NegativePromptUnclip, payload.Query)

	_, err = paramsWriter.Write([]byte(params))
	if err != nil {
		return nil, err
	}

	_ = w.WriteField("model_id", "4")

	w.Close()

	req, err := http.NewRequest(http.MethodPost, url, &b)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", w.FormDataContentType())
	req.Header.Add(KEY_API_HEADER_X_KEY, config.GetConfig().API.GetHeaderKey())
	req.Header.Add(KEY_API_HEADER_X_SECRET, config.GetConfig().API.GetHeaderSecret())

	res, err := http.DefaultClient.Do(req)
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
	url := fmt.Sprintf("%s%s", config.GetConfig().API.URL, formattedUrl)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add(KEY_API_HEADER_X_KEY, config.GetConfig().API.GetHeaderKey())
	req.Header.Add(KEY_API_HEADER_X_SECRET, config.GetConfig().API.GetHeaderSecret())

	for {
		res, err := http.DefaultClient.Do(req)
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
