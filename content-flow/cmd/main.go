package main

import (
	"content-flow/internal/config"
	"content-flow/internal/middleware"
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
}
