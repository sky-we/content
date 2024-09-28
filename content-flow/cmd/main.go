package cmd

import "content-flow/internal/config"

func init() {
	config.LoadFlowConfig()
}

//func main() {
//	fs := config.NewFlowService(config.)
//	fs.start
//}
