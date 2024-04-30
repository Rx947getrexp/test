package config

import (
	"github.com/gin-gonic/gin"
)

const (
	GeoIpFilePath     = "/wwwroot/go/go-api/geo/geo.ip"
	GeoDomainFilePath = "/wwwroot/go/go-api/geo/geo.domain"
)

func GenV2rayConfig(ctx *gin.Context, servers []Server, country string, withoutRules bool) V2rayConfig {
	return V2rayConfig{
		Routing: Routing{
			Rules:          genRules(ctx, country, withoutRules),
			DomainMatcher:  "hybrid",
			DomainStrategy: "AsIs",
			Balancers:      []string{},
		},
		Log: Log{
			Loglevel: "warning",
			DnsLog:   false,
		},
		Outbounds: genOutbounds(servers),
	}
}

func genRules(ctx *gin.Context, country string, withoutRules bool) []Rule {
	rules := make([]Rule, 0)
	rules = append(rules, Rule{
		Type:        "field",
		OutboundTag: "direct",
		Domain:      GenRuleDomain(ctx, country, withoutRules),
	})
	rules = append(rules, Rule{
		Type:        "field",
		OutboundTag: "direct",
		Ip:          GenRuleIp(ctx, country, withoutRules),
	})
	return rules
}

func GenRuleIp(ctx *gin.Context, country string, withoutRules bool) (ips []string) {
	// private
	ips = []string{
		//"0.0.0.0/8",
		//"10.0.0.0/8",
		//"100.64.0.0/10",
		//"127.0.0.0/8",
		//"169.254.0.0/16",
		//"172.16.0.0/12",
		//"192.0.0.0/24",
		//"192.0.2.0/24",
		//"192.88.99.0/24",
		//"192.168.0.0/16",
		//"198.18.0.0/15",
		//"198.51.100.0/24",
		//"203.0.113.0/24",
		//"224.0.0.0/3",
		//"::/127",
		//"fc00::/7",
		//"fe80::/10",
		//"ff00::/8",
		"geoip:private",
	}
	if withoutRules {
		return ips
	}
	var (
		err   error
		lines []string
	)

	lines, err = readFileLines(ctx, GeoIpFilePath, country)
	if err != nil {
		return
	}
	ips = append(ips, lines...)
	return
}

func GenRuleDomain(ctx *gin.Context, country string, withoutRules bool) (domains []string) {
	// private
	domains = []string{
		"icloud",
		"apple",
		"geosite:private",
	}
	if withoutRules {
		return domains
	}
	var (
		err   error
		lines []string
	)

	lines, err = readFileLines(ctx, GeoDomainFilePath, country)
	if err != nil {
		return
	}
	domains = append(domains, lines...)
	return
}

func genOutbounds(servers []Server) []Outbound {
	items := make([]Outbound, 0)
	items = append(items, Outbound{
		Tag: "proxy",
		Mux: &Mux{
			Enabled:     false,
			Concurrency: 50,
		},
		Protocol: "trojan",
		StreamSettings: &StreamSettings{
			WsSettings: WsSettings{
				Path: "/work",
				Headers: Headers{
					Host: "",
				},
			},
			TlsSettings: TlsSettings{
				Alpn:          []string{"http/1.1"},
				AllowInsecure: true,
				Fingerprint:   "",
			},
			Security: "tls",
			Network:  "ws",
		},
		Settings: &Settings{
			Servers: servers,
		},
	})
	items = append(items, Outbound{
		Tag:      "direct",
		Protocol: "freedom",
	})
	items = append(items, Outbound{
		Tag:      "reject",
		Protocol: "blackhole",
	})
	return items
}

type V2rayConfig struct {
	Routing   Routing    `json:"routing"`
	Log       Log        `json:"log"`
	Outbounds []Outbound `json:"outbounds"`
}

type Routing struct {
	Rules          []Rule   `json:"rules"`
	DomainMatcher  string   `json:"domainMatcher"`
	DomainStrategy string   `json:"domainStrategy"`
	Balancers      []string `json:"balancers"`
}

type Rule struct {
	Type        string   `json:"type"`
	OutboundTag string   `json:"outboundTag"`
	Domain      []string `json:"domain,omitempty"`
	Ip          []string `json:"ip,omitempty"`
}

type Log struct {
	Loglevel string `json:"loglevel"`
	DnsLog   bool   `json:"dnsLog"`
}

type Outbound struct {
	Tag            string          `json:"tag"`
	Mux            *Mux            `json:"mux,omitempty"`
	Protocol       string          `json:"protocol"`
	StreamSettings *StreamSettings `json:"streamSettings,omitempty"`
	Settings       *Settings       `json:"settings,omitempty"`
}

type Mux struct {
	Enabled     bool `json:"enabled"`
	Concurrency int  `json:"concurrency"`
}

type StreamSettings struct {
	WsSettings  WsSettings  `json:"wsSettings"`
	TlsSettings TlsSettings `json:"tlsSettings"`
	Security    string      `json:"security"`
	Network     string      `json:"network"`
}

type WsSettings struct {
	Path    string  `json:"path"`
	Headers Headers `json:"headers"`
}

type Headers struct {
	Host string `json:"host"`
}

type TlsSettings struct {
	Alpn          []string `json:"alpn"`
	AllowInsecure bool     `json:"allowInsecure"`
	Fingerprint   string   `json:"fingerprint"`
}

type Settings struct {
	Servers []Server `json:"servers"`
}

type Server struct {
	Password string `json:"password"`
	Port     int    `json:"port"`
	Email    string `json:"email"`
	Level    int    `json:"level"`
	Flow     string `json:"flow"`
	Address  string `json:"address"`
}
