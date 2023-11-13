package executor

import (
	"encoding/json"
	"fmt"
	"go-speed/global"
	"io/ioutil"
	"os"
)

/*
{
    "log":{
        "loglevel":"debug",
        "access":"/var/log/v2ray/access.log",
        "error":"/var/log/v2ray/error.log"
    },
    "stats":{

    },
    "api":{
        "tag":"api",
        "services":[
            "HandlerService",
            "LoggerService",
            "StatsService"
        ]
    },
    "policy":{
        "levels":{
            "0":Object{...}
        },
        "system":{
            "statsInboundUplink":true,
            "statsInboundDownlink":true,
            "statsOutboundUplink":true,
            "statsOutboundDownlink":true
        }
    },
    "inbounds":[
        {
            "port":10085,
            "listen":"0.0.0.0",
            "protocol":"trojan",
            "settings":{
                "decryption":"none",
                "clients":[

                ]
            },
            "streamSettings":{
                "network":"ws",
                "security":"none",
                "tlsSettings":{
                    "alpn":[
                        "http/1.1"
                    ],
                    "certificates":[
                        {
                            "certificateFile":"/usr/local/cert/cert.crt",
                            "keyFile":"/usr/local/cert/private.key"
                        }
                    ]
                },
                "wsSettings":{
                    "path":"/work",
                    "headers":{

                    }
                }
            },
            "tag":"tcp-ws",
            "sniffing":{
                "enabled":true,
                "destOverride":[
                    "http",
                    "tls"
                ]
            }
        },
        {
            "listen":"127.0.0.1",
            "port":10088,
            "protocol":"dokodemo-door",
            "settings":{
                "address":"127.0.0.1"
            },
            "tag":"api"
        }
    ],
    "outbounds":[
        {
            "protocol":"freedom",
            "settings":{

            }
        },
        {
            "protocol":"blackhole",
            "settings":{

            },
            "tag":"blocked"
        }
    ],
    "routing":{
        "rules":[
            {
                "inboundTag":[
                    "api"
                ],
                "ip":[
                    "geoip:private"
                ],
                "outboundTag":"api",
                "type":"field"
            },
            {
                "outboundTag":"blocked",
                "protocol":[
                    "bittorrent"
                ],
                "type":"field"
            }
        ]
    }
}
*/

type V2rayConfig struct {
	Log       ConfLog        `json:"log,omitempty"`
	Stats     ConfStats      `json:"stats,omitempty"`
	Api       ConfApi        `json:"api,omitempty"`
	Policy    ConfPolicy     `json:"policy,omitempty"`
	Inbounds  []ConfInbound  `json:"inbounds,omitempty"`
	Outbounds []ConfOutbound `json:"outbounds,omitempty"`
	Routing   ConfRouting    `json:"routing,omitempty"`
}

type ConfLog struct {
	Loglevel string `json:"loglevel,omitempty"`
	Access   string `json:"access,omitempty"`
	Error    string `json:"error,omitempty"`
}

type ConfStats struct {
}

type ConfApi struct {
	Tag      string   `json:"tag,omitempty"`
	Services []string `json:"services,omitempty"`
}

type ConfPolicy struct {
	Levels map[string]PolicyLevel `json:"levels,omitempty"`
	System PolicySystem           `json:"system,omitempty"`
}

type PolicyLevel struct {
	StatsUserUplink   bool `json:"statsUserUplink,omitempty"`
	StatsUserDownlink bool `json:"statsUserDownlink,omitempty"`
	Handshake         int  `json:"handshake,omitempty"`
	ConnIdle          int  `json:"connIdle,omitempty"`
	BufferSize        int  `json:"bufferSize,omitempty"`
}

type PolicySystem struct {
	StatsInboundUplink    bool `json:"statsInboundUplink,omitempty"`
	StatsInboundDownlink  bool `json:"statsInboundDownlink,omitempty"`
	StatsOutboundUplink   bool `json:"statsOutboundUplink,omitempty"`
	StatsOutboundDownlink bool `json:"statsOutboundDownlink,omitempty"`
}

type ConfInbound struct {
	Port           int                   `json:"port,omitempty"`
	Listen         string                `json:"listen,omitempty"`
	Protocol       string                `json:"protocol,omitempty"`
	Settings       InboundSettings       `json:"settings,omitempty"`
	StreamSettings InboundStreamSettings `json:"streamSettings,omitempty"`
	Tag            string                `json:"tag,omitempty"`
	Sniffing       Sniffing              `json:"sniffing,omitempty"`
}

type InboundSettings struct {
	Decryption string                  `json:"decryption,omitempty"`
	Clients    []InboundSettingsClient `json:"clients,omitempty"`
	Address    string                  `json:"address,omitempty"`
}

type InboundSettingsClient struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type InboundStreamSettings struct {
	Network     string     `json:"network,omitempty"`
	Security    string     `json:"security,omitempty"`
	TlsSettings TlsSetting `json:"tlsSettings,omitempty"`
	WsSettings  WsSetting  `json:"wsSettings,omitempty"`
}

type TlsSetting struct {
	Alpn         []string      `json:"alpn,omitempty"`
	Certificates []Certificate `json:"certificates,omitempty"`
}

type WsSetting struct {
	Path    string  `json:"path,omitempty"`
	Headers Headers `json:"headers,omitempty"`
}

type Certificate struct {
	CertificateFile string `json:"certificateFile,omitempty"`
	KeyFile         string `json:"keyFile,omitempty"`
}

type Sniffing struct {
	Enabled      bool     `json:"enabled,omitempty"`
	DestOverride []string `json:"destOverride,omitempty"`
}

type Headers struct {
}

type ConfOutbound struct {
	Protocol string          `json:"protocol,omitempty"`
	Settings OutboundSetting `json:"settings,omitempty"`
	Tag      string          `json:"tag,omitempty"`
}

type OutboundSetting struct {
}

type ConfRouting struct {
	Rules []Rule `json:"rules,omitempty"`
}

type Rule struct {
	InboundTag  []string `json:"inboundTag,omitempty"`
	Ip          []string `json:"ip,omitempty"`
	OutboundTag string   `json:"outboundTag,omitempty"`
	Type        string   `json:"type,omitempty"`
	Protocol    []string `json:"protocol,omitempty"`
}

func (c *V2rayConfig) GetTagByProtocol(protocol string) string {
	for _, i := range c.Inbounds {
		if i.Protocol == protocol {
			return i.Tag
		}
	}
	return ""
}

func (c *V2rayConfig) GetClientsByProtocol(protocol string) []InboundSettingsClient {
	for _, i := range c.Inbounds {
		if i.Protocol == protocol {
			return i.Settings.Clients
		}
	}
	return make([]InboundSettingsClient, 0)
}

func (c *V2rayConfig) AddClientsByProtocol(protocol, email, password string) error {
	if c.IsClientExist(protocol, email) {
		global.Logger.Warn().Msg("user already exits: " + email)
		return nil
	}

	clients := c.GetClientsByProtocol(protocol)
	clients = append(clients, InboundSettingsClient{
		Email:    email,
		Password: password,
	})

	for ind, i := range c.Inbounds {
		if i.Protocol == protocol {
			c.Inbounds[ind].Settings.Clients = clients
			return nil
		}
	}
	global.Logger.Error().Msg("can not find protocol: " + protocol)
	return fmt.Errorf("can not find protocol: " + protocol)
}

func (c *V2rayConfig) DeleteClientsByProtocol(protocol, email string) error {
	if !c.IsClientExist(protocol, email) {
		global.Logger.Warn().Msg("user not exits: " + email)
		return nil
	}

	clients := c.GetClientsByProtocol(protocol)
	clientsNew := make([]InboundSettingsClient, 0)
	for _, item := range clients {
		if item.Email != email {
			clientsNew = append(clientsNew, item)
		}
	}

	for ind, i := range c.Inbounds {
		if i.Protocol == protocol {
			c.Inbounds[ind].Settings.Clients = clientsNew
			return nil
		}
	}
	global.Logger.Error().Msg("can not find protocol: " + protocol)
	return fmt.Errorf("can not find protocol: " + protocol)
}

func (c *V2rayConfig) IsClientExist(protocol, email string) bool {
	for _, i := range c.Inbounds {
		if i.Protocol == protocol {
			for _, item := range i.Settings.Clients {
				if item.Email == email {
					return true
				}
			}
		}
	}
	return false
}

func (c *V2rayConfig) PersistToConfigFile(filePath string) error {
	jsonBytes, err := json.MarshalIndent(*c, "", "	")
	if err != nil {
		global.Logger.Error().Msg("MarshalIndent failed, err: " + err.Error())
		return err
	}
	global.Logger.Info().Msgf(">>>>>> config: \n%s", c.ToString())
	global.Logger.Info().Msg(">>>>>> MarshalIndent config: \n" + string(jsonBytes))
	err = ioutil.WriteFile(filePath, jsonBytes, 0644)
	if err != nil {
		global.Logger.Error().Msg("WriteFile failed, err: " + err.Error())
		return err
	}
	return nil
}

func (c *V2rayConfig) ToString() string {
	b, err := json.Marshal(c)
	if err != nil {
		global.Logger.Error().Msg("Marshal failed, err: " + err.Error())
		return ""
	}
	return string(b)
}

// 读取v2ray配置文件
func ReadV2rayConfig(filePath string) (c V2rayConfig, err error) {
	var content []byte
	content, err = os.ReadFile(filePath)
	if err != nil {
		return
	}
	err = json.Unmarshal(content, &c)
	if err != nil {
		return
	}
	global.Logger.Info().Msgf(">>>>>>>>> filePath: %s, V2rayConfig: %s", filePath, c.ToString())
	return
}
