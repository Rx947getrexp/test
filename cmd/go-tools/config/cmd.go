package config

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go-speed/cmd/go-tools/common"
	"go-speed/dao"
	"go-speed/model/entity"
	"go-speed/service"
	"log"
	"net/http/httptest"
	"strings"
)

const (
	CmdTypeCheck            = "check"
	CmdTypeCheckConfigAndDB = "check-config-and-db"
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
	command.Flags().StringVarP(&cmdAction, "action", "a", "check", "操作类型(check|check-config-and-db)")
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
		case CmdTypeCheckConfigAndDB:
			checkV2rayConfigAndDB(ctx, nodeIp)
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
	log.Printf("nodeIp: %s check finished. len(items): %d\n", nodeIp, len(resp.Items))
}

func checkV2rayConfigAndDB(ctx *gin.Context, nodeIp string) {
	resp, err := service.GetUserListFromNode(ctx, nodeIp)
	if err != nil {
		log.Fatalf("%s GetUserListFromNode failed, err: %s", nodeIp, err.Error())
	}

	if nodeIp == "154.93.104.115" {
		out, _ := json.Marshal(resp)
		log.Printf("resp: %s\n", string(out))
	}

	m := make(map[string]struct{})
	configUser := make(map[string]string)
	var emails []string
	for _, item := range resp.Items {

		email := strings.ToLower(item.Email)
		if _, ok := m[email]; ok {
			log.Printf("%s %s, dumplicate\n", nodeIp, email)
		} else {
			m[strings.ToLower(item.Email)] = struct{}{}
		}

		if _, ok := configUser[item.Email]; ok {
			log.Printf("%s %s, dumplicate configUser\n", nodeIp, email)
		} else {
			configUser[item.Email] = item.Password
			emails = append(emails, item.Email)
		}
	}
	var users []entity.TUser
	err = dao.TUser.Ctx(ctx).WhereIn(dao.TUser.Columns().Email, emails).Scan(&users)
	if err != nil {
		log.Printf("%s query user from db failed, err: %s", nodeIp, err.Error())
	} else {
		dbUser := make(map[string]string)
		for _, user := range users {
			if _, ok := dbUser[user.Email]; ok {
				log.Printf("%s, dumplicate db user info\n", user.Email)
			} else {
				dbUser[user.Email] = user.V2RayUuid
			}
		}
		for email, uuid := range configUser {
			if dbUuid, ok := dbUser[email]; ok {
				if uuid != dbUuid {
					log.Printf("%s %s, config uuid is diff to db user info\n", nodeIp, email)
				}
			} else {
				//log.Printf("%s %s, config user can not find in db user info\n", nodeIp, email)
			}
		}
		if nodeIp == "154.93.104.115" {
			out1, _ := json.Marshal(configUser)
			out2, _ := json.Marshal(dbUser)
			log.Printf("configUser: %s\n", string(out1))
			log.Printf("dbUser: %s\n", string(out2))
		}
	}
	log.Printf("nodeIp: %s check finished. len(items): %d\n", nodeIp, len(resp.Items))
}
