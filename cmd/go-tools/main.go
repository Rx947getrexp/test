package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/spf13/cobra"
	"go-speed/cmd/go-tools/tools"
)

var rootCmd *cobra.Command

func cmdInit() {
	// 创建根命令
	rootCmd = &cobra.Command{
		Use:   "speedctl",
		Short: "A command line tool with multiple functions and commands",
	}
	// 添加顶级命令

	rootCmd.AddCommand(tools.AddConfigCmd)
	rootCmd.Execute()
}

func AddCommand(cmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	cmdInit()
}
