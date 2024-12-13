package tools

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go-speed/api/executor"
	"go-speed/service/api/speed_api"
	"log"
	"net/http/httptest"
	"os/exec"
	"strings"
)

func init() {
	loadNodeConfig.Flags().StringVarP(&ip, "ip", "i", "", "节点IP(required)")
	AddConfigCmd.AddCommand(loadNodeConfig)
}

var loadNodeConfig = &cobra.Command{
	Use:   "loadnodeconfig",
	Short: "load node config",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
		defer ctx.Done()

		var ips []string
		if ip == "" {
			resp, err := speed_api.DescribeNodeList(ctx)
			if err != nil {
				log.Fatalf("DescribeNodeList failed with %s\n", err.Error())
			}
			for _, item := range resp.Items {
				ips = append(ips, item.Ip)
			}
		} else {
			ips = append(ips, ip)
		}
		for _, i := range ips {
			LoadNodeConfig(i)
		}
	},
}

func LoadNodeConfig(ip string) (emails []string, err error) {
	cmdName := "ssh"
	cmdArgs := []string{
		fmt.Sprintf("root@%s", ip),
		fmt.Sprintf("cat /usr/local/etc/v2ray/config.json"),
	}
	fmt.Printf("ssh cmdArgs: %+v\n", cmdArgs)

	// 创建Cmd结构体
	cmd := exec.Command(cmdName, cmdArgs...)

	// 创建一个缓冲区来存储命令的输出
	var out bytes.Buffer
	cmd.Stdout = &out

	// 运行命令
	err = cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed, err: %s", err.Error())
	}
	// 打印命令的输出
	//global.Logger.Debug().Msgf("Output: %+v", out.String())

	var config executor.V2rayConfig
	err = json.Unmarshal(out.Bytes(), &config)
	if err != nil {
		return
	}

	m := make(map[string]executor.InboundSettingsClient)
	for _, v := range config.GetClients() {
		em := strings.ToLower(v.Email)
		if _, ok := m[em]; !ok {
			m[em] = v
		} else {
			fmt.Printf("duplicate: %s\t%s\n", v.Email, v.Password)
		}
	}
	fmt.Printf("%s len(emails)=%d\n\n\n", ip, len(m))
	return
}
