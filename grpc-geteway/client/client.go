package main

import (
	"context"
	"fmt"
	v1 "gateway/service/v1"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("enter")

	conn, err := grpc.Dial("127.0.0.1:8088", grpc.WithInsecure())
	if err != nil {
		log.Panic(err)
	}
	client := v1.NewYourServiceClient(conn)
	ctx := context.Background()
	param := v1.StringMessage{Value: "12"}

	res, err := client.Echo(ctx, &param)

	fmt.Println(res.Value)
	fmt.Println("end")

}
