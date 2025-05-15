package util

import (
	"fmt"
	"go-speed/global"
	"strings"
)

const (
	Connector = "-hshs-"
)

func GetUserV2rayConfigEmail(email string) string {
	appName := strings.TrimSpace(global.Config.System.AppName)
	if appName != "" {
		return fmt.Sprintf("%s%s%s", email, Connector, appName)
	} else {
		return email
	}
}

func GetUserV2rayConfigUUID(uuid string) string {
	appName := strings.TrimSpace(global.Config.System.AppName)
	if appName != "" {
		return fmt.Sprintf("%s%s%s", uuid, Connector, appName)
	} else {
		return uuid
	}
}
