package common

import (
	"context"
	"go-speed/dao"
	"go-speed/model/do"
	"go-speed/model/entity"
	"log"
)

func GetNodeIps(ip string) (ips []string) {
	if ip != "" {
		return []string{ip}
	}

	var nodes []entity.TNode
	err := dao.TNode.Ctx(context.Background()).Where(do.TNode{Status: 1}).Scan(&nodes)
	if err != nil {
		log.Fatalf("get TNode failed, err: %s", err.Error())
	}

	for _, node := range nodes {
		ips = append(ips, node.Ip)
	}
	return
}
