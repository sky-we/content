package config

import (
	"content-system/internal/api/operate"
	"content-system/internal/middleware"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"time"
)

type MysqlConfig struct {
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

type RedisConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	DB       int
}

type RedisWinConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
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

type DataBaseConfig struct {
	MySQL       *MysqlConfig
	Redis       *RedisConfig
	RedisWin    *RedisWinConfig
	FlowService *FlowServiceConfig
	AppClient   *AppClientConfig
}

var (
	once     sync.Once
	DBConfig *DataBaseConfig
	Logger   = middleware.GetLogger()
)

func LoadDBConfig() {

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

func NewMySqlDB(cfg *MysqlConfig) *gorm.DB {
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
	Logger.Info("mysql connect dsn:", dsn)
	mysqlDB, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	db, err := mysqlDB.DB()
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(cfg.MaxOpenConn)
	db.SetMaxIdleConns(cfg.MaxIdleConn)
	return mysqlDB
}

func NewRdb(cfg *RedisConfig) *redis.Client {
	option := redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	}
	Logger.Info("redis connect option:", option)
	rdb := redis.NewClient(&option)
	return rdb
}

func NewAppClient(cfg *AppClientConfig) operate.AppClient {
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
