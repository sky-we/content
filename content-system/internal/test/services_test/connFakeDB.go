package services

import (
	"content-system/internal/api/operate"
	"content-system/internal/process"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	goflow "github.com/s8sg/goflow/v1"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var FakeDbCfg *FakeDBConfig

type FakeMysqlConfig struct {
	Host        string
	Port        int
	Username    string
	Password    string
	DBName      string
	ChartSet    string
	ParseTime   string
	Loc         string
	MaxOpenConn int
	MaxIdleConn int
}

type FlowServiceConfig struct {
	RedisURL          string
	Port              int
	WorkerConcurrency int
}
type AppClientConfig struct {
	Host string
	Port int
}
type FakeDBConfig struct {
	MySQL       *FakeMysqlConfig
	FlowService *FlowServiceConfig
	AppClient   *AppClientConfig
}

func LoadFakeDBConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../base/fakeConfig")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("viper ReadInConfig /base/fakeConfig panic")
		panic(err)

	}
	if err := viper.Unmarshal(&FakeDbCfg); err != nil {
		panic(err)
	}

}

func NewFakeMySqlDB(cfg *FakeMysqlConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%s&loc=%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
		cfg.ChartSet,
		cfg.ParseTime,
		cfg.Loc,
	)
	mysqlDB, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		fmt.Println("connect mysql error:", err)
		panic(err)
	}
	db, err := mysqlDB.DB()
	if err != nil {
		fmt.Println("get mysql instance error:", err)
		panic(err)
	}
	db.SetMaxOpenConns(cfg.MaxOpenConn)
	db.SetMaxIdleConns(cfg.MaxIdleConn)
	return mysqlDB
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

func NewAppClient(cfg *AppClientConfig) operate.AppClient {
	endPoint := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(endPoint),
		grpc.WithMiddleware(
			recovery.Recovery(),
		))
	if err != nil {
		panic(err)
	}
	client := operate.NewAppClient(conn)
	return client
}
