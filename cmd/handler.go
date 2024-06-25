package cmd

import (
	"context"
	"github.com/v2fly/v2ray-core/v5/app/proxyman/command"
	"github.com/v2fly/v2ray-core/v5/common/protocol"
	"github.com/v2fly/v2ray-core/v5/common/serial"
	"github.com/v2fly/v2ray-core/v5/proxy/vmess"
)

func AddV2rayVmessUser(client command.HandlerServiceClient, inboundTag string, email string, id string) (*command.AlterInboundResponse, error) {
	return client.AlterInbound(context.Background(), &command.AlterInboundRequest{
		Tag: inboundTag,
		Operation: serial.ToTypedMessage(&command.AddUserOperation{
			User: &protocol.User{
				Email: email,
				Account: serial.ToTypedMessage(&vmess.Account{
					Id: id,
				}),
			},
		}),
	})
}

func RemoveV2rayUser(client command.HandlerServiceClient, inboundTag string, email string) error {
	_, err := client.AlterInbound(context.Background(), &command.AlterInboundRequest{
		Tag: inboundTag,
		Operation: serial.ToTypedMessage(&command.RemoveUserOperation{
			Email: email,
		}),
	})
	return err
}
