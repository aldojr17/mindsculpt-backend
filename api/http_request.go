package api

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"mindsculpt/config"
	"mindsculpt/domain"
	log "mindsculpt/logger"
	"net/http"
	"net/textproto"
)

func newRequest(method string, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add(KEY_API_HEADER_X_KEY, config.GetConfig().API.GetHeaderKey())
	req.Header.Add(KEY_API_HEADER_X_SECRET, config.GetConfig().API.GetHeaderSecret())

	log.Infof("Method: %s; URL: %s", req.Method, req.URL)

	return req, nil
}

func NewGetRequest(url string) (*http.Request, error) {
	return newRequest(http.MethodGet, url, nil)
}

func NewPostRequest(url string, body bytes.Buffer) (*http.Request, error) {
	return newRequest(http.MethodPost, url, &body)
}

func DoRequest(req *http.Request) (*http.Response, error) {
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	log.Infof("Status: %s", res.Status)

	return res, nil
}

func GenerateHeader(payload domain.APIGenerateImageRequest) (bytes.Buffer, *multipart.Writer, error) {
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	paramsPartHeader := textproto.MIMEHeader{}
	paramsPartHeader.Set("Content-Disposition", `form-data; name="params"`)
	paramsPartHeader.Set("Content-Type", "application/json")

	paramsWriter, err := writer.CreatePart(paramsPartHeader)
	if err != nil {
		return body, writer, err
	}

	params := fmt.Sprintf(GENERATE_PARAMS, payload.Width, payload.Height, payload.NegativePromptUnclip, payload.Query)

	_, err = paramsWriter.Write([]byte(params))
	if err != nil {
		return body, writer, err
	}

	_ = writer.WriteField("model_id", "4")

	writer.Close()

	return body, writer, nil
}
