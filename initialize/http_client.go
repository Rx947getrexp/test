package initialize

import (
	"crypto/tls"
	"go-speed/global"
	"net/http"
	"net/url"
)

func initClient() *http.Client {
	client := &http.Client{}
	if len(global.Config.System.HttpProxy) > 0 {
		proxyUrl, err := url.Parse(global.Config.System.HttpProxy)
		if err == nil {
			tr := &http.Transport{TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			}}
			tr.Proxy = http.ProxyURL(proxyUrl)
			client.Transport = tr
		}
	}
	return client
}
