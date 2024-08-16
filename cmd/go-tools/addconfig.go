package tools

import (
	"fmt"
	"github.com/gogf/gf/os/gctx"
	"github.com/spf13/cobra"
	"go-speed/cmd/go-tools/utils"
	"go-speed/constant"
	"go-speed/model/request"
	"go-speed/model/response"
	"go-speed/util"
	"time"
)

var (
	email string
	ip string
)

func init() {
	addConfigCmdDefine.Flags().StringVarP(&email, "email", "e", "", "用户邮箱(required)")
	addConfigCmdDefine.Flags().StringVarP(&ip, "ip", "i", "", "节点IP(required)")
	AddConfigCmd.AddCommand(addConfigCmdDefine)
}

var AddConfigCmd = &cobra.Command{
	Use:   "v2ray",
	Short: "add v2ray config",
}

var addConfigCmdDefine = &cobra.Command{
	Use:   "addconfig",
	Short: "Command add v2ray config",
	//Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("email: %s, ip: %s\n", email, ip)
		addConfig(email, ip)
	},
}

func addConfig(userName, nodeIP string) {
	// 获取参数值
	var (
		ctx  = gctx.New()
		user = utils.GetUser(ctx, userName)
		node = utils.GetNode(ctx, nodeIP)
	)

	url := fmt.Sprintf("http://%s:15003/node/add_sub", node.Ip)
	fmt.Println(url)

	nodeAddSubRequest := &request.NodeAddSubRequest{}
	nodeAddSubRequest.Tag = "1"
	nodeAddSubRequest.Uuid = user.V2RayUuid
	nodeAddSubRequest.Email = user.Email
	nodeAddSubRequest.Level = fmt.Sprintf("%d", user.Level)

	timestamp := fmt.Sprint(time.Now().Unix())
	headerParam := make(map[string]string)
	res := new(response.Response)
	headerParam["timestamp"] = timestamp
	headerParam["accessToken"] = util.MD5(fmt.Sprint(timestamp, constant.AccessTokenSalt))
	err := util.HttpClientPostV2(url, headerParam, nodeAddSubRequest, res)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	res.Print()
}

