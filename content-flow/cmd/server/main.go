package main

import (
	"content-flow/internal/config"
	"content-flow/internal/middleware"
	"os"
	"os/signal"
)

func init() {
	config.LoadFlowCfg()
}

func main() {
	middleware.InitLogger()
	Logger := middleware.GetLogger()
	fs := config.NewFlowService(config.ClientCfg.FlowService)
	if err := fs.Start(); err != nil {
		Logger.Error("go-flow service start error")
		panic(err)
	}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	Logger.Info("shutting down go-flow server")
}
