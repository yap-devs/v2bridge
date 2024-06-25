package main

import (
	"fmt"
	"github.com/v2fly/v2ray-core/v5/app/proxyman/command"
	"github.com/yap-devs/v2bridge/cmd"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	server := "dmm-tv:10085"
	conn, err := grpc.Dial(server, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("error connecting to server: ", err)
	}
	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)

	client := command.NewHandlerServiceClient(conn)

	res, err := cmd.AddV2rayVmessUser(client, "main-inbound", "233@gmail.com", "12345678-7890-1234-4321-123456789012")

	if err != nil {
		fmt.Println("error adding user: ", err)
	}

	fmt.Println(res)
}
