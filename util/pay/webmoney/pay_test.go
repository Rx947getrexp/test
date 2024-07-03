package webmoney

import (
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"testing"
)

func TestQueryDeal(*testing.T) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//resp, err := QueryDeal(ctx, "100629092041783")
	//if err != nil {
	//	global.MyLogger(ctx).Err(err).Msgf("QueryDeal failed")
	//	return
	//}
	//if resp == nil {
	//	fmt.Println("resp is nil")
	//} else {
	//	fmt.Println(*resp)
	//}

	//QueryOrder(ctx, "100630133031068")
	QueryOrder(ctx, "100701092254247")

	//QueryDeal(ctx, "0")
}
