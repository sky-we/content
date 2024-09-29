package config

import (
	"content-flow/internal/api/operate"
	"content-flow/internal/middleware"
	"content-flow/internal/process"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	goflow "github.com/s8sg/goflow/v1"
	"github.com/spf13/viper"
	"sync"
	"time"
)

type FlowService struct {
	Port              int
	RedisURL          string
	WorkerConcurrency int
	FlowName          string
}

type AppClient struct {
	Host string
	Port int
}

type Required struct {
	FlowService *FlowService
	AppClient   *AppClient
}

var (
	once    sync.Once
	WireCfg Required
	Logger  = middleware.GetLogger()
)

func LoadFlowCfg() {

	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("internal/config")

		if err := viper.ReadInConfig(); err != nil {
			Logger.Error("error reading db config file, %s", err)
			panic(err)

		}
		if err := viper.Unmarshal(&WireCfg); err != nil {
			Logger.Error("unable to decode into struct, %v", err)
			panic(err)
		}
	})
}

func NewFlowService(cfg *FlowService) *goflow.FlowService {
	fs := goflow.FlowService{
		Port:              cfg.Port,
		RedisURL:          cfg.RedisURL,
		WorkerConcurrency: cfg.WorkerConcurrency,
	}
	client := NewAppClient(WireCfg.AppClient)
	contentFlow := process.NewContentFlow(client)
	err := fs.Register(cfg.FlowName, contentFlow.ContentFlowHandle)
	if err != nil {
		panic(err)
	}

	return &fs
}

func NewAppClient(cfg *AppClient) operate.AppClient {
	endPoint := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(endPoint),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
		grpc.WithTimeout(time.Second*1000),
	)
	if err != nil {
		panic(err)
	}
	client := operate.NewAppClient(conn)
	return client
}
