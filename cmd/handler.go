package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/v2fly/v2ray-core/v5/app/proxyman/command"
	"github.com/v2fly/v2ray-core/v5/common/protocol"
	"github.com/v2fly/v2ray-core/v5/common/serial"
	"github.com/v2fly/v2ray-core/v5/proxy/vmess"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type jsonRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

var handlerCmd = &cobra.Command{
	Use:   "handler",
	Short: "HandlerService commands",
}

// addV2rayVmessUserCmd represents the addV2rayVmessUser command
var addV2rayVmessUserCmd = &cobra.Command{
	Use:   "addV2rayVmessUser",
	Short: "Add a V2ray VMess user to the inbound",
	Run: func(cmd *cobra.Command, args []string) {
		server := cmd.Flag("server").Value.String()
		inboundTag := cmd.Flag("tag").Value.String()
		email := cmd.Flag("email").Value.String()
		uuid := cmd.Flag("uuid").Value.String()

		client, err := NewClient(server)
		if err != nil {
			res, _ := json.Marshal(jsonRes{
				Code: 1,
				Msg:  err.Error(),
			})
			fmt.Println(string(res))
			return
		}

		alterInboundResponse, err := AddV2rayVmessUser(client, inboundTag, email, uuid)
		if err != nil {
			res, _ := json.Marshal(jsonRes{
				Code: 2,
				Msg:  err.Error(),
			})
			fmt.Println(string(res))
			return
		}

		res, _ := json.Marshal(jsonRes{
			Code: 0,
			Msg:  alterInboundResponse.String(),
		})
		fmt.Println(string(res))
	},
}

// removeV2rayUserCmd represents the removeV2rayUser command
var removeV2rayUserCmd = &cobra.Command{
	Use:   "removeV2rayUser",
	Short: "Remove a V2ray user from the inbound",
	Run: func(cmd *cobra.Command, args []string) {
		server := cmd.Flag("server").Value.String()
		inboundTag := cmd.Flag("tag").Value.String()
		email := cmd.Flag("email").Value.String()

		client, err := NewClient(server)
		if err != nil {
			res, _ := json.Marshal(jsonRes{
				Code: 1,
				Msg:  err.Error(),
			})
			fmt.Println(string(res))
			return
		}

		alterInboundResponse, err := RemoveV2rayUser(client, inboundTag, email)
		if err != nil {
			res, _ := json.Marshal(jsonRes{
				Code: 2,
				Msg:  err.Error(),
			})
			fmt.Println(string(res))
			return
		}

		res, _ := json.Marshal(jsonRes{
			Code: 0,
			Msg:  alterInboundResponse.String(),
		})
		fmt.Println(string(res))
	},
}

func NewClient(server string) (command.HandlerServiceClient, error) {
	conn, err := grpc.Dial(server, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("error connecting to server: %w", err)
	}

	return command.NewHandlerServiceClient(conn), nil
}

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

func RemoveV2rayUser(client command.HandlerServiceClient, inboundTag string, email string) (*command.AlterInboundResponse, error) {
	return client.AlterInbound(context.Background(), &command.AlterInboundRequest{
		Tag: inboundTag,
		Operation: serial.ToTypedMessage(&command.RemoveUserOperation{
			Email: email,
		}),
	})
}

func init() {
	rootCmd.AddCommand(handlerCmd)

	handlerCmd.AddCommand(addV2rayVmessUserCmd)
	addV2rayVmessUserCmd.Flags().StringP("tag", "t", "main-inbound", "Inbound tag")
	addV2rayVmessUserCmd.Flags().StringP("email", "e", "", "User email")
	addV2rayVmessUserCmd.Flags().StringP("uuid", "u", "", "User uuid")
	_ = addV2rayVmessUserCmd.MarkFlagRequired("email")
	_ = addV2rayVmessUserCmd.MarkFlagRequired("uuid")

	handlerCmd.AddCommand(removeV2rayUserCmd)
	removeV2rayUserCmd.Flags().StringP("tag", "t", "main-inbound", "Inbound tag")
	removeV2rayUserCmd.Flags().StringP("email", "e", "", "User email")
	_ = removeV2rayUserCmd.MarkFlagRequired("email")
}
