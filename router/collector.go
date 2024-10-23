package router

import (
	"github.com/gin-gonic/gin"
	administrator "go-speed/api/admin/report"
	"go-speed/api/api/node"
	"go-speed/api/api/report"
)

func CollectorRoute(group *gin.RouterGroup) {
	group.GET("get_user_op_log_list", administrator.GetUserOpLogList)
	group.GET("list_node_for_report", node.ListNodeForReport)          //获取节点ip列表，上报ping结果
	group.GET("get_client_ip", report.GetClientIP)                     //获取客户端IP
	group.POST("report_node_ping_result", report.ReportNodePingResult) //上报ping结果
	group.POST("report_user_op_log", report.ReportUserOpLog)           // 连接代理
}
