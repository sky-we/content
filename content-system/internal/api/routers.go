package api

import (
	"content-system/internal/config"
	"content-system/internal/middleware"
	"content-system/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	rootPath    = "/api"
	outRootPath = "/out/api"
)

func CmsRouters(r *gin.Engine) {
	db := config.NewMySqlDB(config.ClientCfg.MySQL)
	rdb := config.NewRdb(config.ClientCfg.Redis)
	appClient := config.NewAppClient(config.ClientCfg.EtcdClient)

	// 依赖注入
	cmsApp := services.NewCmsApp(db, rdb, appClient)

	r.Use(middleware.Prometheus())
	r.Use(middleware.OpenTracing())

	// 鉴权中间件
	sessionMiddleware := &middleware.SessionAuth{Rdb: rdb}

	// 路由
	router := r.Group(rootPath)
	router.Use(sessionMiddleware.Auth)
	{
		// 服务探测
		router.GET("/cms/probe", cmsApp.Probe)

		// 内容创建
		router.POST("/cms/content/create", cmsApp.ContentCreate)

		// 内容更新
		router.POST("/cms/content/update", cmsApp.ContentUpdate)

		// 内容删除
		router.POST("/cms/content/delete", cmsApp.ContentDelete)

		// 内容查询
		router.POST("/cms/content/find", cmsApp.ContentFind)
	}

	outRoot := r.Group(outRootPath)
	{
		// 用户注册
		outRoot.POST("/cms/register", cmsApp.Register)

		// 用户登录
		outRoot.POST("/cms/login", cmsApp.Login)
	}
	// prometheus 采集指标
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

}
