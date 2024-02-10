package service

import (
	"mindsculpt/api"
	"mindsculpt/domain"
	"mindsculpt/repository/cache"
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
