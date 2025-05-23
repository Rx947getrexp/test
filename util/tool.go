package util

import (
	"net"
	"reflect"
	"strings"

	"github.com/bwmarrin/snowflake"
	"github.com/google/uuid"
	"golang.org/x/net/publicsuffix"
)

func GenSnowflake(node int64) (int64, error) {
	snowFlakeNode, err := snowflake.NewNode(node)
	if err != nil {
		return 0, err
	}
	return snowFlakeNode.Generate().Int64(), nil
}

func GetUuid() string {
	nonce, err := uuid.NewUUID()
	if err != nil {
		return ""
	}
	tmp := strings.Split(nonce.String(), "-")
	var result string
	for _, item := range tmp {
		result += item
	}
	return result
}

func MapsWith(newContainer interface{}, oldContainer interface{}, args ...[]string) {
	oldType := reflect.TypeOf(oldContainer)
	oldValue := reflect.ValueOf(oldContainer)
	newModel := reflect.ValueOf(newContainer).Elem()

	// 过滤
	var filterArgs []string
	// 映射指定
	// conditionArgs
	if len(args) > 0 {
		for _, val := range args {
			filterArgs = val
			break
		}
	}

OutModelLoop:
	for i := 0; i < oldType.NumField(); i++ {
		mapOldIndex := oldType.Field(i)

		// 判断是否过滤
		if len(filterArgs) > 0 {
			for s := 0; s < len(filterArgs); s++ {
				if mapOldIndex.Name == filterArgs[s] {
					continue OutModelLoop
				}
			}
		}

		oldVal := oldValue.Field(i).Interface()
		if isTrue := newModel.FieldByName(mapOldIndex.Name).IsValid(); isTrue {
			newModel.FieldByName(mapOldIndex.Name).Set(reflect.ValueOf(oldVal))
		}
	}
}

// 获取主域名
func GetDomain(host string) string {
	if host == "" {
		return ""
	}
	// 去掉端口
	if h, _, err := net.SplitHostPort(host); err == nil {
		host = h
	}

	// 如果是 IPv4 或 IPv6，直接返回
	if ip := net.ParseIP(host); ip != nil {
		return ip.String()
	}

	host = strings.ToLower(host)
	host = strings.TrimPrefix(host, "www.")

	// 使用 publicsuffix 提取主域名
	domain, err := publicsuffix.EffectiveTLDPlusOne(host)
	if err != nil {
		return ""
	}
	return domain
}
