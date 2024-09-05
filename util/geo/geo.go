package geo

import (
	"fmt"
	"go-speed/global"
	"go-speed/util/geo/geolite"
	"strings"

	"github.com/gin-gonic/gin"
)

func IsNeedDisablePaymentFeature(ctx *gin.Context, email string) bool {
	var (
		err         error
		country     string
		countryList = global.Config.PayConfig.DisablePaymentCountryList
		emailList   = global.Config.PayConfig.DisablePaymentEmailList
	)
	global.MyLogger(ctx).Info().Msgf("current user clientIP(%s) Email(%s). config disabled payment countryList(%s), emailList(%+v)", ctx.ClientIP(), email, countryList, emailList)

	country, err = geolite.QueryIPCountry(ctx.ClientIP())
	global.MyLogger(ctx).Info().Msgf("QueryIPCountry(%s) -> Country(%s), err(%+v)", ctx.ClientIP(), country, err)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("IP(%s) -> Country(%s) failed, err: %s, so disable payment", ctx.ClientIP(), country, err.Error())
		return true
	}

	if country == "" {
		err = fmt.Errorf("QueryIPCountry is empty, can not judge the country")
		global.MyLogger(ctx).Err(err).Msgf("IP(%s) -> Country(%s) is nil, err: %s, so disable payment", ctx.ClientIP(), country, err.Error())
		return true
	}

	var (
		emails    = strings.Split(emailList, ",")
		countries = strings.Split(countryList, ",")
	)

	if IsInArrayIgnoreCase(country, countries) {
		global.MyLogger(ctx).Info().Msgf("IP(%s) -> Country(%s) is in(%s), disable payment", ctx.ClientIP(), country, countryList)
		return true
	}
	if IsInArray(email, emails) || emailList == "all" {
		global.MyLogger(ctx).Info().Msgf("email(%s) is in(%s), disable payment", email, emails)
		return true
	}
	global.MyLogger(ctx).Info().Msgf("email(%s) IP(%s) -> Country(%s), is a payment user", email, ctx.ClientIP(), country)
	return false
}

func IsInArrayIgnoreCase(a string, arrs []string) bool {
	a = strings.ToLower(a)
	for _, v := range arrs {
		if a == strings.ToLower(v) {
			return true
		}
	}
	return false
}

func IsInArray(a string, arrs []string) bool {
	for _, v := range arrs {
		if a == v {
			return true
		}
	}
	return false
}
