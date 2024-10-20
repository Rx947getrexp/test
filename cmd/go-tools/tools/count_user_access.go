package tools

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os/exec"
	"strings"
	"time"
)

var date string

func init() {
	countUserAccessV2ray.Flags().StringVarP(&ip, "ip", "i", "", "节点IP(required)")
	countUserAccessV2ray.Flags().StringVarP(&date, "date", "d", "", "统计日期，格式为: (optional)")
	AddConfigCmd.AddCommand(countUserAccessV2ray)
}

var countUserAccessV2ray = &cobra.Command{
	Use:   "countUserAccess",
	Short: "count user access",
	Run: func(cmd *cobra.Command, args []string) {
		doCountUserAccessV2ray(ip)
	},
}

func doCountUserAccessV2ray(ip string) {
	if ip == "" {
		log.Fatalf("ip is required")
	}
	//ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	//defer ctx.Done()

	fmt.Printf("ip: %s\n", ip)

	ips := getNodeIps(ip)

	// 获取当前时间
	now := time.Now()

	// 计算昨天的日期
	yesterday := now.AddDate(0, 0, -1)

	// 获取年、月和日
	//year, month, day := yesterday.Date()

	for _, ip := range ips {
		// 定义命令和参数
		//cmd := exec.Command("scp",
		//	fmt.Sprintf("root@%s:/usr/local/etc/v2ray/config.json", ip),
		//	fmt.Sprintf("/wwwroot/go/go-api/config/v2ray/users-%s.json", ip),
		//)

		cmdName := "ssh"
		cmdArgs := []string{
			fmt.Sprintf("root@%s", ip),
			fmt.Sprintf("grep -a '%s' /var/log/v2ray/access.log | grep email | grep accepted | awk '{print $7}' | sort | uniq", yesterday.Format("2006/01/02")),
		}
		fmt.Println(cmdName, cmdArgs)

		// 创建Cmd结构体
		cmd := exec.Command(cmdName, cmdArgs...)

		// 创建一个缓冲区来存储命令的输出
		var out bytes.Buffer
		cmd.Stdout = &out

		// 运行命令
		err := cmd.Run()
		if err != nil {
			log.Fatalf("cmd.Run() failed with %s\n", err)
		}

		// 打印命令的输出
		fmt.Printf("Output: %q\n", out.String())

		// 将输出的字符串按照换行符分割，然后将结果存储到一个[]string变量中
		lines := strings.Split(out.String(), "\n")
		var emails []string
		for _, line := range lines {
			emails = append(emails, line)
		}
		fmt.Printf("emails: %q\n", emails)
	}
}

/*
+----------------+
| ip             |
+----------------+
| 110.42.42.229  |
| 147.45.178.51  |
| 154.93.104.115 |
| 185.39.207.104 |
| 185.39.207.20  |
| 193.124.41.88  |
| 193.233.48.70  |
| 207.90.237.91  |
| 212.18.104.23  |
| 213.159.68.106 |
| 38.55.136.95   |
| 45.147.200.112 |
| 45.147.201.21  |
| 46.17.41.7     |
| 46.17.44.132   |
| 62.133.60.81   |
| 62.133.63.237  |
| 91.149.218.194 |
| 92.118.112.133 |
| 92.118.112.89  |
+----------------+

authorized_keys


ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQCb5Om4AiG92MQtCoOPFtc2AK4oXBgygRZcRN794bmOZyfAjA0VZeXk5lX7O9njMBy2smACj8EfXhNezEXlpCkEy6CbPdfic9ecvEEpuh1LNrFDNkt1hcnK4azErbAK9gOYI1DFU7+EAQWo+72VTsF944JS/BfpfETcczDhky5UyoVU++46OSTIJOmCQBB72TK2yh1NoSyhPHFoK7K5Ky82NQlfXsQs+nLhfAnKvsACCZlzWAw9Z1HtFoV4hwEJKbnD8yMqwbDv9GabmQHUt/GfRHHTkG33ecFpGzWjhvpMCyZk0OmQ9IWStAvc0ogCgGPvKN0lq4cIUb0qkS2qJ+kQFCcFl2ZBIdyPBEfD1LdaJoSo+y4WqOMwREZZhLEowpN7lXzTV7HBhFClHBumF5IHxPjObPDjkfO2yReogFGu5+dBISX6KDsX9eS+rUsJpi0BFvDpuWK2FQ93us4MOkABO4MySskdrNKxTN0LogQ4wFQ0hEIEMe/bZIPv+p9C00U= root@vm857751


*/
