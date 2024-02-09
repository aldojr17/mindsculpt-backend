package initialize

import (
	"mindsculpt/config"
	log "mindsculpt/logger"

	"github.com/go-redis/redis"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Application struct {
	Database *gorm.DB
	Redis    *redis.Client
}

func InitApp() *Application {
	cfg := config.GetConfig()

	app := &Application{
		Database: initializeDB(cfg),
		Redis:    initializeRedis(cfg),
	}

	return app
}

func initializeDB(cfg *config.Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(cfg.Database.Config()), &gorm.Config{})

	if err != nil {
		log.Errorf("error initialize database : %s", err.Error())
		panic("error initialize database")
	}

	log.Info("Database:\n",
		"config", cfg.Database.ConfigInfo(),
	)

	return db
}

func initializeRedis(cfg *config.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: cfg.Redis.Addr(),
		DB:   cfg.Redis.Index(),
	})

	log.Info("Redis:\n",
		"config", cfg.Redis.ConfigInfo(),
	)

	return client
}
