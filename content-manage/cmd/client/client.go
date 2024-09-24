package main

import (
	"content-manage/api/content/operate"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

func main() {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("localhost:9000"),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := operate.NewAppClient(conn)

	reply, err := client.CreateContent(context.Background(), &operate.CreateContentReq{
		Content: &operate.Content{
			Title:       "test content_manage with kratos",
			VideoUrl:    "https://baidu.com",
			Author:      "sky-we",
			Description: "test kratos",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("[grpc reply +%v]", reply)
}
