package geolite

import (
	"net"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/oschwald/geoip2-golang"
)

func QueryIPCountry(ip string) (country string, err error) {
	if ip == "" {
		return
	}
	var (
		db         *geoip2.Reader
		geoCountry *geoip2.Country
	)
	db, err = geoip2.Open("/wwwroot/go/go-api/geo/GeoLite2-Country.mmdb")
	if err != nil {
		err = gerror.Wrap(err, `geoip2.Open() failed`)
		return
	}
	defer db.Close()

	netip := net.ParseIP(ip)
	geoCountry, err = db.Country(netip)
	if err != nil {
		err = gerror.Wrap(err, `geoip2.Country() failed`)
		return
	}
	return geoCountry.Country.IsoCode, nil
}
