package executor

import (
	"context"
	"go-speed/global"
	"go-speed/model/request"
	"strings"

	"google.golang.org/grpc"

	"github.com/v2fly/v2ray-core/v5/app/proxyman/command"
	"github.com/v2fly/v2ray-core/v5/common/protocol"
	"github.com/v2fly/v2ray-core/v5/common/serial"
	"github.com/v2fly/v2ray-core/v5/proxy/trojan"
)

const (
	V2rayProtocolTrojan = "trojan"
)

func addUser(req *request.NodeAddSubRequest) error {
	conf, err := ReadV2rayConfig(global.Config.System.V2rayConfigPath)
	if err != nil {
		global.Logger.Err(err).Msg("read v2ray config failed, err: " + err.Error())
		return err
	}

	// TODO：level写死为0
	err = addUserOperation(conf, req.Email, req.Uuid, 0)
	if err != nil {
		global.Logger.Err(err).Msg("add user by v2ray api failed, err: " + err.Error())
		return err
	}

	return addUserToV2rayConfig(conf, req.Email, req.Uuid)
}

func addUserToV2rayConfig(conf V2rayConfig, email, password string) error {
	err := conf.AddClientsByProtocol(V2rayProtocolTrojan, email, password) // TODO：写死 trojan
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

// trojan
func addUserOperation(conf V2rayConfig, email, password string, level uint32) error {
	// 在添加用户之前，检查该用户是否已经存在
	if conf.IsClientExist(V2rayProtocolTrojan, email) {
		global.Logger.Warn().Msgf("User %s already exists, skipping AddUser operation.", email)
		return nil // 如果用户已存在，直接返回，不再添加
	}
	// 准备新客户端的信息
	req := command.AlterInboundRequest{
		Tag: conf.GetTagByProtocol(V2rayProtocolTrojan), // TODO：写死 trojan
		Operation: serial.ToTypedMessage(
			&command.AddUserOperation{
				User: &protocol.User{
					Level: level,
					Email: email,
					Account: serial.ToTypedMessage(&trojan.Account{
						Password: password,
					}),
				},
			}),
	}
	// "127.0.0.1:10088"
	conn, err := grpc.Dial(global.Config.System.V2rayApiAddress, grpc.WithInsecure())
	if err != nil {
		global.Logger.Err(err).Msg("grpc.Dial failed, err: " + err.Error())
		return err
	}
	defer conn.Close()
	global.Logger.Info().Msgf("addUser req: Tag: %s, Operation: %s", req.Tag, req.Operation.String())
	resp, err := command.NewHandlerServiceClient(conn).AlterInbound(context.Background(), &req)
	if err != nil {
		global.Logger.Err(err).Msg("AlterInbound failed, err: " + err.Error())
		if strings.Contains(err.Error(), " already exists") {
			global.Logger.Warn().Msg(">>>>>>>>> already exists")
			return nil
		}
		return err
	}
	global.Logger.Info().Msgf("AlterInbound.Resp: %s", resp.String())
	return nil
}
