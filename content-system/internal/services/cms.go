package services

import (
	"content-system/internal/api/operate"
	"content-system/internal/middleware"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type CmsApp struct {
	db               *gorm.DB
	rdb              *redis.Client
	operateAppClient operate.AppClient
}

var Logger = middleware.GetLogger()

func NewCmsApp(
	db *gorm.DB,
	rdb *redis.Client,
	operateAppClient operate.AppClient,
) *CmsApp {
	app := &CmsApp{
		db:               db,
		rdb:              rdb,
		operateAppClient: operateAppClient,
	}
	return app
}
