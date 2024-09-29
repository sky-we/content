package main

import (
	"content-flow/internal/config"
	"content-flow/internal/middleware"
	"fmt"
)

func init() {
	config.LoadFlowCfg()
}

func main() {
	middleware.InitLogger()
	fs := config.NewFlowService(config.WireCfg.FlowService)
	if err := fs.Start(); err != nil {
		fmt.Println("go-flow service start error")
		panic(err)
	}
}
