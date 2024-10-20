package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go-speed/dao"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/service"
	"io/ioutil"
	"log"
	"net/http/httptest"
)

func init() {
	getV2rayConfig.Flags().StringVarP(&ip, "ip", "i", "", "节点IP(optional)")
	AddConfigCmd.AddCommand(getV2rayConfig)
}

var getV2rayConfig = &cobra.Command{
	Use:   "getv2rayconfig",
	Short: "get v2ray config",
	//Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("ip: %s\n", ip)
		doGetV2rayConfig(ip)
	},
}

func doGetV2rayConfig(ip string) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	defer ctx.Done()

	ips := getNodeIps(ip)

	for _, ip := range ips {
		resp, err := service.GetUserListFromNode(ctx, ip)
		if err != nil {
			log.Fatalf("GetUserListFromNode failed, err: %s", err.Error())
		}

		m := make(map[string]string)
		for _, item := range resp.Items {
			m[item.Email] = item.Password
		}

		jsonData, err := json.MarshalIndent(m, "", "  ")
		if err != nil {
			log.Fatalf("Marshal resp failed, err: %s", err.Error())
		}

		fileName := fmt.Sprintf("/wwwroot/go/go-api/config/v2ray/users-%s.json", ip)
		// 将JSON数据写入到文件中
		err = ioutil.WriteFile(fileName, jsonData, 0644)
		if err != nil {
			log.Fatalf("WriteFile failed: %s", err)
		}
	}

	//cmd := exec.Command("scp",
	//	fmt.Sprintf("root@%s:/usr/local/etc/v2ray/config.json", ip),
	//	fmt.Sprintf("/wwwroot/go/go-api/config/v2ray/users-%s.json", ip),
	//)
	//
	//err := cmd.Run()
	//if err != nil {
	//	log.Fatalf("cmd.Run() failed with %s\n", err)
	//}
}

func getNodeIps(ip string) (ips []string) {
	if ip != "" {
		return []string{ip}
	}

	var nodes []entity.TNode
	err := dao.TNode.Ctx(context.Background()).Where(do.TNode{Status: 1}).Scan(&nodes)
	if err != nil {
		log.Fatalf("get TNode failed, err: %s", err.Error())
	}

	for _, node := range nodes {
		ips = append(ips, node.Ip)
	}
	return
}
