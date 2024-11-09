package tools

import (
	"bufio"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go-speed/dao"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/service"
	"go-speed/service/api/speed_api"
	"log"
	"net/http/httptest"
	"os"
	"strings"
)

var (
	fileName string
)

func init() {
	checkUserTrafficCmdDefine.Flags().StringVarP(&fileName, "file", "f", "./email.txt", "邮箱列表文件")
	checkUserTrafficCmdDefine.Flags().StringVarP(&email, "email", "e", "", "用户邮箱")
	AddConfigCmd.AddCommand(checkUserTrafficCmdDefine)
}

var checkUserTrafficCmdDefine = &cobra.Command{
	Use:   "checkusertraffic",
	Short: "检查用户是否产生过流量",
	Run: func(cmd *cobra.Command, args []string) {
		checkUserTraffic()
	},
}

func readFile(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line) // 去除空格和换行
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

func checkUserTraffic() {
	if fileName == "" && email == "" {
		log.Printf("file and email both empty")
		return
	}
	var emails []string
	if email != "" {
		emails = []string{email}
	} else {
		emails = readFile(fileName)
	}
	log.Printf("emails: %+v", emails)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	defer ctx.Done()

	var ips []string
	resp, err := speed_api.DescribeNodeList(ctx)
	if err != nil {
		log.Fatalf("DescribeNodeList failed with %s\n", err.Error())
	}
	for _, item := range resp.Items {
		ips = append(ips, item.Ip)
	}
	if len(ips) == 0 {
		log.Fatalf("ips is empty %+v\n", ips)
	}

	for ind, e := range emails {
		if e == "" {
			continue
		}
		log.Printf("\n\n######## %d) %s 查询结果为：", ind, e)
		doCheck(ctx, e, ips)
	}
}

func doCheck(ctx *gin.Context, email string, ips []string) {
	var items []entity.TV2RayUserTraffic
	err := dao.TV2RayUserTraffic.Ctx(ctx).Where(do.TV2RayUserTraffic{Email: email}).Scan(&items)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range items {
		log.Printf("########[Traffic] %s:\t%d", item.Ip, item.Date)
	}

	for _, nodeIp := range ips {
		resp, err := service.GetUserTraffic(ctx, email, nodeIp)
		if err != nil {
			log.Fatalf("\n\nGetUserTraffic failed from ip: %s", nodeIp)
		}
		for _, node := range resp {
			log.Printf("########[Traffic] %s:\t%d|%d", nodeIp, node.UpLink, node.DownLink)
		}
	}
}
