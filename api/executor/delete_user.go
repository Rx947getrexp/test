package executor

import (
	"context"
	"strings"

	"go-speed/global"
	"go-speed/model/request"

	"github.com/xtls/xray-core/app/proxyman/command"
	"github.com/xtls/xray-core/common/serial"
	"google.golang.org/grpc"
)

func delUser(req *request.NodeAddSubRequest) error {
	conf, err := ReadV2rayConfig(global.Config.System.V2rayConfigPath)
	if err != nil {
		global.Logger.Err(err).Msg("read v2ray config failed, err: " + err.Error())
		return err
	}

	// TODO：level写死为0
	err = deleteUserToV2rayConfig(conf, req.Email)
	if err != nil {
		global.Logger.Err(err).Msg("delete user from config failed, err: " + err.Error())
		return err
	}
	return removeUserOperation(conf, req.Email)
}

// trojan
func removeUserOperation(conf V2rayConfig, email string) error {
	// 准备新客户端的信息
	req := command.AlterInboundRequest{
		Tag: conf.GetTagByProtocol(V2rayProtocolTrojan), // TODO：写死 trojan
		Operation: serial.ToTypedMessage(
			&command.RemoveUserOperation{
				Email: email,
			}),
	}

	// "127.0.0.1:10088"
	conn, err := grpc.Dial(global.Config.System.V2rayApiAddress, grpc.WithInsecure())
	if err != nil {
		global.Logger.Err(err).Msg("grpc.Dial failed, err: " + err.Error())
		return err
	}
	defer conn.Close()
	global.Logger.Info().Msgf("remove user req: Tag: %s, Operation: %s", req.Tag, req.Operation.String())

	resp, err := command.NewHandlerServiceClient(conn).AlterInbound(context.Background(), &req)
	if err != nil {
		global.Logger.Err(err).Msg("RemoveUserOperation failed, err: " + err.Error())
		if strings.Contains(err.Error(), " not found") {
			global.Logger.Warn().Msg(">>>>>>>>> not found")
			return nil
		}
		return err
	}
	global.Logger.Info().Msgf("RemoveUserOperation.Resp: %s", resp.String())
	return nil
}

func deleteUserToV2rayConfig(conf V2rayConfig, email string) error {
	err := conf.DeleteClientsByProtocol(V2rayProtocolTrojan, email) // TODO：写死 trojan
	if err != nil {
		global.Logger.Err(err).Msg("add clients to config failed")
		return err
	}

	err = conf.PersistToConfigFile(global.Config.System.V2rayConfigPath)
	if err != nil {
		global.Logger.Err(err).Msg("persist to config file failed")
		return err
	}
	return nil
}
