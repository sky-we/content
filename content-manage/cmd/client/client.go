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
			VideoURL:    "https://baidu.com",
			Author:      "sky-we",
			Description: "test kratos",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("[grpc create content reply +%v\n]", reply)

	reply2, err := client.UpdateContent(context.Background(), &operate.UpdateContentReq{
		Content: &operate.Content{
			ID:          37,
			Title:       "test content_manage with kratos",
			VideoURL:    "https://baidu.com",
			Author:      "sky-we",
			Description: "test kratos",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("[grpc update content reply +%v\n]", reply2)

	reply3, err := client.DeleteContent(context.Background(), &operate.DeleteContentReq{
		Id: 77,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("[grpc delete content reply3 +%v\n]", reply3)

	reply4, err := client.FindContent(context.Background(), &operate.FindContentReq{
		Title:    "test content_manage with kratos",
		Page:     int64(1),
		PageSize: int64(5),
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("[grpc delete content reply4 [%v]]", reply4)
}
