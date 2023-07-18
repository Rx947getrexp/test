package main

import (
	"context"
	"fmt"
	proxymanService "github.com/v2fly/v2ray-core/v5/app/proxyman/command"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"testing"

	// 统计服务
	statsService "github.com/v2fly/v2ray-core/v5/app/stats/command"
	"github.com/v2fly/v2ray-core/v5/common/protocol"
	"github.com/v2fly/v2ray-core/v5/common/serial"
	"github.com/v2fly/v2ray-core/v5/proxy/vmess"
)

var conn *grpc.ClientConn

func init() {
	var err error
	// 连接grpc服务
	conn, err = grpc.Dial("104.233.171.69:443", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
}

func close() {
	if err := conn.Close(); err != nil {
		fmt.Println("关闭连接失败")
	}
}

func TestAdd(t *testing.T) {
	defer close()
	err := addUser(conn)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestRemove(t *testing.T) {
	defer close()
	err := removeUser(conn)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestQuery(t *testing.T) {
	defer close()
	var err error
	err = queryUserTrafficV2(conn)
	err = queryUserTraffic(conn)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// 添加用户
func addUser(conn *grpc.ClientConn) error {
	client := proxymanService.NewHandlerServiceClient(conn)
	userUUID := "07084eac-892e-4a45-8845-420a13c9d7b3" //uuid.NewV4().String() //使用UUID库生成一个UUID
	resp, err := client.AlterInbound(context.Background(), &proxymanService.AlterInboundRequest{
		Tag: "tcp-ws", // 要添加用户的tag，目前只支持vmess协议
		Operation: serial.ToTypedMessage(&proxymanService.AddUserOperation{
			User: &protocol.User{
				Level: 0,            // 用户等级
				Email: "ccc@qq.com", // 用户邮箱，删除和统计要用到
				Account: serial.ToTypedMessage(&vmess.Account{
					Id:               userUUID,                                                   //用户UUID
					AlterId:          64,                                                         // 额外ID
					SecuritySettings: &protocol.SecurityConfig{Type: protocol.SecurityType_AUTO}, // 安全设置，auto为自动判断加密协议
				}),
			},
		}),
	})
	if err != nil {
		return err
	}
	fmt.Println(resp)
	fmt.Println(userUUID)
	return nil
}

// 删除用户
func removeUser(conn *grpc.ClientConn) error {
	client := proxymanService.NewHandlerServiceClient(conn)
	resp, err := client.AlterInbound(context.Background(), &proxymanService.AlterInboundRequest{
		Tag: "tcp-ws",
		Operation: serial.ToTypedMessage(&proxymanService.RemoveUserOperation{
			Email: "ccc@qq.com", // 用户邮箱地址
		}),
	})
	if err != nil {
		return err
	}
	fmt.Println(resp)
	return nil
}

// 获取用户流量
func queryUserTraffic(conn *grpc.ClientConn) error {
	client := statsService.NewStatsServiceClient(conn)
	resp, err := client.QueryStats(context.Background(), &statsService.QueryStatsRequest{
		/*
		   user>>>[email]>>>traffic>>>uplink
		   特定用户的上行流量，单位字节。
		   user>>>[email]>>>traffic>>>downlink
		   特定用户的下行流量，单位字节。
		   inbound>>>[tag]>>>traffic>>>uplink
		   特定入站代理的上行流量，单位字节。
		   inbound>>>[tag]>>>traffic>>>downlink
		   特定入站代理的下行流量，单位字节。
		*/
		Patterns: []string{"user>>>ccc@qq.com", "user>>>aaa@qq.com", "user"}, // 筛选用户表达式
		Reset_:   false,                                                      // 查询完成后是否重置流量
	})
	if err != nil {
		return err
	}
	// 获取返回值中的流量信息
	stat := resp.GetStat()
	// 返回的是一个数组，对其进行遍历输出
	for _, e := range stat {
		fmt.Println(e.Name, e.Value)
	}
	return nil
}

// 获取用户流量
func queryUserTrafficV2(conn *grpc.ClientConn) error {
	client := statsService.NewStatsServiceClient(conn)
	resp, err := client.QueryStats(context.Background(), &statsService.QueryStatsRequest{
		/*
		   user>>>[email]>>>traffic>>>uplink
		   特定用户的上行流量，单位字节。
		   user>>>[email]>>>traffic>>>downlink
		   特定用户的下行流量，单位字节。
		   inbound>>>[tag]>>>traffic>>>uplink
		   特定入站代理的上行流量，单位字节。
		   inbound>>>[tag]>>>traffic>>>downlink
		   特定入站代理的下行流量，单位字节。
		*/
		Patterns: []string{"user>>>ccc@qq.com"}, // 筛选用户表达式
		Reset_:   true,                          // 查询完成后是否重置流量
	})
	if err != nil {
		return err
	}
	// 获取返回值中的流量信息
	stat := resp.GetStat()
	// 返回的是一个数组，对其进行遍历输出
	for _, e := range stat {
		fmt.Println(e.Name, e.Value)
	}
	return nil
}
