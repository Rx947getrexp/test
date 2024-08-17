package utils

import (
	"context"
	"go-speed/dao"
	"go-speed/model/do"
	"go-speed/model/entity"
)

func GetUser(ctx context.Context, email string) *entity.TUser {
	if email == "" {
		panic("input 'email' can not be empty")
	}
	var user *entity.TUser
	if err := dao.TUser.Ctx(ctx).Where(do.TUser{Email: email}).Scan(&user); err != nil {
		panic(err)
	} else {
		if user == nil {
			panic("input 'email' can not find user info")
		}
		return user
	}
}

func GetNode(ctx context.Context, ip string) *entity.TNode {
	if ip == "" {
		panic("input 'ip' can not be empty")
	}
	var node *entity.TNode
	if err := dao.TNode.Ctx(ctx).Where(do.TNode{Ip: ip}).Scan(&node); err != nil {
		panic(err)
	} else {
		if node == nil {
			panic("input 'ip' can not find node info")
		}
		return node
	}

}
