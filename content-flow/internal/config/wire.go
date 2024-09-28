package config

import (
	"content-flow/internal/middleware"
	"content-flow/internal/process"
	goflow "github.com/s8sg/goflow/v1"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"sync"
)

type FlowServiceConfig struct {
	Port              int
	RedisURL          string
	WorkerConcurrency int
}

var (
	once     sync.Once
	DBConfig *FlowServiceConfig
	Logger   = middleware.GetLogger()
)

func LoadFlowConfig() {

	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("internal/config")

		if err := viper.ReadInConfig(); err != nil {
			Logger.Error("error reading db config file, %s", err)
			panic(err)

		}
		if err := viper.Unmarshal(&DBConfig); err != nil {
			Logger.Error("unable to decode into struct, %v", err)
			panic(err)
		}
	})
}

func NewFlowService(cfg *FlowServiceConfig, db *gorm.DB) *goflow.FlowService {
	fs := goflow.FlowService{
		Port:              cfg.Port,
		RedisURL:          cfg.RedisURL,
		WorkerConcurrency: cfg.WorkerConcurrency,
	}
	contentFlow := process.NewContentFlow(db)
	err := fs.Register("content-flow", contentFlow.ContentFlowHandle)
	if err != nil {
		panic(err)
	}

	return &fs
}
