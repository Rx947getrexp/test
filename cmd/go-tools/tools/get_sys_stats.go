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
	getSysStatsCmdDefine.Flags().StringVarP(&ip, "ip", "i", "", "节点IP(required)")
	AddConfigCmd.AddCommand(getSysStatsCmdDefine)
}

var getSysStatsCmdDefine = &cobra.Command{
	Use:   "getsysstats",
	Short: "get system stats",
	//Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		printSysStats(ip)
	},
}

func printSysStats(ip string) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	defer ctx.Done()

	ips := []string{ip}
	if ip == "" {
		resp, err := speed_api.DescribeNodeList(ctx)
		if err != nil {
			log.Fatalf("DescribeNodeList failed with %s\n", err.Error())
		}
		for _, item := range resp.Items {
			ips = append(ips, item.Ip)
		}
	}

	for _, ip := range ips {
		fmt.Printf("ip: %s\n", ip)
		_, _ = service.GetSysStatsByIp(ctx, ip)
		fmt.Print("\n\n")
	}
}
