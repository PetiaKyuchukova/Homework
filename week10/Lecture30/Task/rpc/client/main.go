package main

import (
	"context"
	"log"
	"topstories/rpc/chat"

	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewHNServiceClient(conn)

	response, err := c.GetTopTenArticles(context.Background(), &chat.MessageRequest{MessReq: "Hello From Client!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Stories)

}
