package executor

import (
	"context"
	"fmt"
	statscmd "github.com/xtls/xray-core/app/stats/command"
	"go-speed/global"
	"google.golang.org/grpc"
)

// trojan
func GetUserTrafficByEmail(conn *grpc.ClientConn, email string) (int64, int64, error) {
	global.Logger.Info().Msgf("getUserTraffic email: %s", email)
	var err error
	if conn == nil {
		// "127.0.0.1:10088"
		conn, err = grpc.Dial(global.Config.System.V2rayApiAddress, grpc.WithInsecure())
		if err != nil {
			global.Logger.Err(err).Msg("grpc.Dial failed, err: " + err.Error())
			return 0, 0, err
		}
		defer conn.Close()
	}
	resp1, err := statscmd.NewStatsServiceClient(conn).GetStats(context.Background(), &statscmd.GetStatsRequest{
		Name:   fmt.Sprintf("user>>>%s>>>traffic>>>uplink", email),
		Reset_: false,
	})
	if err != nil {
		global.Logger.Err(err).Msg("uplink GetStats failed, err: " + err.Error())
		return 0, 0, err
	}
	global.Logger.Info().Msgf("uplink GetStats.Resp: %s", resp1.String())

	resp2, err := statscmd.NewStatsServiceClient(conn).GetStats(context.Background(), &statscmd.GetStatsRequest{
		Name:   fmt.Sprintf("user>>>%s>>>traffic>>>downlink", email),
		Reset_: false,
	})
	if err != nil {
		global.Logger.Err(err).Msg("downlink GetStats failed, err: " + err.Error())
		return 0, 0, err
	}
	global.Logger.Info().Msgf("downlink GetStats.Resp: %s", resp2.String())
	return resp1.GetStat().GetValue(), resp2.GetStat().GetValue(), nil
}

func GetSysTraffic(conn *grpc.ClientConn) error {
	var err error
	if conn == nil {
		// "127.0.0.1:10088"
		conn, err = grpc.Dial(global.Config.System.V2rayApiAddress, grpc.WithInsecure())
		if err != nil {
			global.Logger.Err(err).Msg("grpc.Dial failed, err: " + err.Error())
			return err
		}
		defer conn.Close()
	}
	resp, err := statscmd.NewStatsServiceClient(conn).GetSysStats(context.Background(), &statscmd.SysStatsRequest{})
	if err != nil {
		global.Logger.Err(err).Msg("GetSysStats failed, err: " + err.Error())
		return err
	}
	global.Logger.Info().Msgf("GetSysStats.Resp: %s", resp.String())

	return nil
}

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

func QueryUserStats(conn *grpc.ClientConn, patterns []string, reset bool) ([]*statscmd.Stat, error) {
	var err error
	if conn == nil {
		// "127.0.0.1:10088"
		conn, err = grpc.Dial(global.Config.System.V2rayApiAddress, grpc.WithInsecure())
		if err != nil {
			global.Logger.Err(err).Msg("grpc.Dial failed, err: " + err.Error())
			return nil, err
		}
		defer conn.Close()
	}
	items := make([]*statscmd.Stat, 0)
	for _, pattern := range patterns {
		resp, err := statscmd.NewStatsServiceClient(conn).QueryStats(context.Background(), &statscmd.QueryStatsRequest{
			Pattern: pattern,
			Reset_:  reset,
		})
		if err != nil {
			global.Logger.Err(err).Msg("QueryStats failed, err: " + err.Error())
			return nil, err
		}
		if resp != nil && len(resp.Stat) > 0 {
			items = append(items, resp.Stat...)
		}
	}

	global.Logger.Info().Msgf("QueryStats.Resp: %#v", items)
	return items, nil
}
