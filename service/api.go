package service

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"mindsculpt/api"
	"mindsculpt/config"
	"mindsculpt/domain"
	"mindsculpt/repository/cache"

	firebase "firebase.google.com/go"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/api/option"
)

type APIService struct {
	modelCache *cache.ModelCache
}

func NewAPIService(modelCache *cache.ModelCache) *APIService {
	return &APIService{
		modelCache: modelCache,
	}
}

func (s *APIService) GetGenerationModels() (*domain.APIGetModelsResponse, error) {
	data, _ := s.modelCache.Get()
	if data != nil {
		return data, nil
	}

	resp, err := api.GetGenerationModels()
	if err != nil {
		return nil, err
	}

	if err = s.modelCache.Set(resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *APIService) GenerateImage(payload domain.APIGenerateImageRequest) (*domain.APIGenerateImageResponse, error) {
	resp, err := api.GenerateImage(payload)
	if err != nil {
		return nil, err
	}

	var response *domain.APIGenerateImageResponse

	if err := mapstructure.Decode(resp, &response); err != nil {
		return nil, err
	}

	opt := option.WithCredentialsFile(config.GetConfig().SecretKey)

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}

	client, err := app.Storage(context.TODO())
	if err != nil {
		return nil, err
	}

	bucketHandle, err := client.Bucket(config.GetConfig().BucketName)
	if err != nil {
		return nil, err
	}

	imageData, err := base64.StdEncoding.DecodeString(resp.Images[0])
	if err != nil {
		return nil, err
	}

	objKey := fmt.Sprintf("%s.%s", uuid.New().String(), api.IMAGE_EXTENSION)

	objectHandle := bucketHandle.Object(objKey)

	token := uuid.New().String()

	writer := objectHandle.NewWriter(context.Background())
	writer.ObjectAttrs.Metadata = map[string]string{api.FIREBASE_TOKEN_METADATA: token}

	defer writer.Close()

	if _, err := io.Copy(writer, bytes.NewReader(imageData)); err != nil {
		return nil, err
	}

	response.ImageUrl = fmt.Sprintf(api.FIREBASE_IMAGE_URL, config.GetConfig().BucketName, objKey, token)

	return response, nil
}
