package config

import (
	"content-system/internal/api/operate"
	"content-system/internal/middleware"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/opentracing/opentracing-go"
	zipkinot "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
	reporter "github.com/openzipkin/zipkin-go/reporter/http"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	clientv3 "go.etcd.io/etcd/client/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormOpenTracing "gorm.io/plugin/opentracing"
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

type FlowServiceClientConfig struct {
	Host     string
	Port     int
	FlowName string
}

type AppClientConfig struct {
	Host string
	Port int
}

type EtcdClientConfig struct {
	Host string
	Port int
}

type ClientConfig struct {
	MySQL             *MysqlConfig
	Redis             *RedisConfig
	FlowServiceClient *FlowServiceClientConfig
	AppClient         *AppClientConfig
	EtcdClient        *EtcdClientConfig
}

var (
	once      sync.Once
	ClientCfg *ClientConfig
	Logger    = middleware.GetLogger()
	// 容器内部的配置文件挂载点
	containerConfig = "/app/config"
)

func LoadDBConfig() {

	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(containerConfig)

		if err := viper.ReadInConfig(); err != nil {
			Logger.Error("error reading db config file, %s", err)
			panic(err)

		}
		if err := viper.Unmarshal(&ClientCfg); err != nil {
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
	gormTracer := NewGormTracer()
	pluginErr := mysqlDB.Use(gormTracer)
	if pluginErr != nil {
		panic(pluginErr)
	}
	return mysqlDB
}

func NewGormTracer() gorm.Plugin {
	// 创建上报节点
	report := reporter.NewReporter("http://container-zipkin:9411/api/v2/spans")
	// 创建本地节点
	endpoint, err := zipkin.NewEndpoint("content-system-gorm", "container-system:8080")

	if err != nil {
		panic(err)
	}
	// 创建zipkin tracer

	tracer, err := zipkin.NewTracer(report,
		zipkin.WithLocalEndpoint(endpoint), // 设置本地节点
		zipkin.WithTraceID128Bit(true))     // 设置ID 128位
	if err != nil {
		panic(err)
	}

	zipTracer := zipkinot.Wrap(tracer)

	opentracing.SetGlobalTracer(zipTracer)

	return gormOpenTracing.New(gormOpenTracing.WithTracer(zipTracer))

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

func NewAppClient(cfg *EtcdClientConfig) operate.AppClient {
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	client, err := clientv3.New(clientv3.Config{
		Endpoints: []string{addr},
	})
	if err != nil {
		panic(err)
	}
	dis := etcd.New(client)
	// 依赖的manage grpc服务
	endPoint := "discovery:///Content-Manage-MicroService"

	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(endPoint),
		grpc.WithDiscovery(dis),
		grpc.WithTimeout(time.Second*1000),
	)

	if err != nil {
		panic(err)
	}
	appClient := operate.NewAppClient(conn)
	return appClient
}
