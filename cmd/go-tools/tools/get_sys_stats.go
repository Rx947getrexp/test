package tools

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go-speed/service"
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
		fmt.Printf("ip: %s\n", ip)

		ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
		defer ctx.Done()
		_, _ = service.GetSysStatsByIp(ctx, ip)
	},
}
