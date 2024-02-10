package cache

import (
	"encoding/json"
	"mindsculpt/domain"
	log "mindsculpt/logger"
	"time"

	"github.com/go-redis/redis"
)

type ModelCache struct {
	redis *redis.Client
	ttl   time.Duration
}

func NewModelCache(redis *redis.Client, ttl time.Duration) *ModelCache {
	return &ModelCache{
		redis: redis,
		ttl:   ttl,
	}
}

func (c *ModelCache) getKey() string {
	return "models"
}

func (c *ModelCache) Set(data *domain.APIGetModelsResponse) error {
	val, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if err := c.redis.Set(c.getKey(), val, c.ttl).Err(); err != nil {
		return err
	}

	log.Info("cache created", "key", c.getKey(), "ttl", c.ttl)

	return nil
}

func (c *ModelCache) Get() (*domain.APIGetModelsResponse, error) {
	response := new(domain.APIGetModelsResponse)

	val, err := c.redis.Get(c.getKey()).Bytes()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(val, &response)
	if err != nil {
		return nil, err
	}

	log.Info("load data from cache")

	return response, nil
}
