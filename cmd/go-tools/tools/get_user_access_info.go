package tools

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go-speed/service"
	"go-speed/service/api/speed_api"
	"log"
	"net/http/httptest"
)

func init() {
	getUserAccessInfoCmdDefine.Flags().StringVarP(&email, "email", "e", "", "用户邮箱(required)")
	getUserAccessInfoCmdDefine.Flags().StringVarP(&ip, "ip", "i", "", "节点IP")
	AddConfigCmd.AddCommand(getUserAccessInfoCmdDefine)
}

var getUserAccessInfoCmdDefine = &cobra.Command{
	Use:   "getuseraccessinfo",
	Short: "查询用户节点流量信息",
	//Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		getUserAccessInfo(email, ip)
	},
}

func getUserAccessInfo(email, ip string) {
	if email == "" {
		log.Fatalf("email不能为空")
	}

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	defer ctx.Done()

	ips := []string{ip}
	if ip == "" {
		ips = []string{}
		resp, err := speed_api.DescribeNodeList(ctx)
		if err != nil {
			log.Fatalf("DescribeNodeList failed with %s\n", err.Error())
		}
		for _, item := range resp.Items {
			ips = append(ips, item.Ip)
		}
	}
	fmt.Println("查询结果：")
	for _, ip := range ips {
		//fmt.Printf("ip: %s\n", ip)
		resp, err := service.GetUserTraffic(ctx, email, ip)
		if err != nil {
			log.Fatalf("\n\nGetUserTraffic failed from ip: %s", ip)
		}

		fmt.Printf("Email: %s, IP: %s, Traffic: %+v \n", email, ip, resp)
	}
}

/*

alt.j9-co173v7i@yopmail.com
kakyoin2929@gmail.com
cry69gry@mail.ru
vlad.ku4erov2015@gmail.com
sergey-gamzik11@mail.ru
slavamihailov222@gmail.com
Kapshanova14@gmail.com

./speedctl  v2ray getuseraccessinfo -e Kapshanova14@gmail.com -i 185.39.207.104
./speedctl  v2ray getuseraccessinfo -e slavamihailov222@gmail.com -i 185.39.207.20
./speedctl  v2ray getuseraccessinfo -e sergey-gamzik11@mail.ru -i 185.39.207.20

*/
