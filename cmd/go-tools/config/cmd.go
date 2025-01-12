package config

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go-speed/cmd/go-tools/common"
	"go-speed/service"
	"log"
	"net/http/httptest"
	"strings"
)

const (
	CmdTypeCheck = "check"
)

var (
	cmdAction string
	ip        string
)

func NewCommand() *cobra.Command {
	command := &cobra.Command{
		Use:     "config",
		Short:   "config",
		Example: `./speedctl config -a check`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runCommand()
		},
	}
	command.Flags().StringVarP(&cmdAction, "action", "a", "check", "操作类型")
	command.Flags().StringVarP(&ip, "ip", "i", "", "机器IP")
	return command
}

func runCommand() (err error) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	defer ctx.Done()

	ips := common.GetNodeIps(ip)
	log.Printf("ips: %v", ips)
	for _, nodeIp := range ips {
		switch cmdAction {
		case CmdTypeCheck:
			checkV2rayConfig(ctx, nodeIp)
		default:
			log.Fatalf("cmdAction %s invalid", cmdAction)
		}
	}

	return
}

func checkV2rayConfig(ctx *gin.Context, nodeIp string) {
	resp, err := service.GetUserListFromNode(ctx, nodeIp)
	if err != nil {
		log.Fatalf("GetUserListFromNode failed, err: %s", err.Error())
	}

	m := make(map[string]struct{})
	for _, item := range resp.Items {
		email := strings.ToLower(item.Email)
		if _, ok := m[email]; ok {
			log.Printf("%s, dumplicate\n", email)
		} else {
			m[strings.ToLower(item.Email)] = struct{}{}
		}
	}
	log.Printf("nodeIp: %s check finished\n", nodeIp)
}
