package russpay

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"testing"
)

func TestCreateOrder(*testing.T) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	resp, err := CreateOrder(ctx, CreateOrderReq{
		Amount:      "1",
		OrderNumber: "202412140001",
		CompanyPage: "https://123.com/api/notify",
	})
	if err != nil {
		fmt.Printf("CreateOrder failed, err: %s\n", err.Error())
		return
	}

	fmt.Printf("CreateOrder success, resp: %#v\n", *resp)
}

func TestQueryOrder(*testing.T) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	resp, err := QueryOrder(ctx, QueryOrderReq{BillingNumber: "5f053f9e9850453fb449d8327e06c842"})
	if err != nil {
		fmt.Printf("QueryOrder failed, err: %s\n", err.Error())
		return
	}

	fmt.Printf("QueryOrder success, resp: %#v\n", *resp)
}

func TestDecode(t *testing.T) {
	content := "eyJwYXlVcmwiOiJodHRwczovL2FwaXdlYi55b2xlc2RrLmNvbS8jL3Rlc3QvcGF5bWVudFJlc3VsdHMvOWI4ZDYxN2EwZDE5NDk0Yjk5ZDExNzFmMjBmN2NiZTUiLCJ1cmwiOiJodHRwczovL2FwaXdlYi55b2xlc2RrLmNvbS8jL3Rlc3QvcGF5bWVudFJlc3VsdHMvOWI4ZDYxN2EwZDE5NDk0Yjk5ZDExNzFmMjBmN2NiZTUiLCJiaWxsaW5nTnVtYmVyIjoiOWI4ZDYxN2EwZDE5NDk0Yjk5ZDExNzFmMjBmN2NiZTUiLCJvcmRlck51bWJlciI6IjIwMjQxMjE0MDAwMSIsImludm9pY2VJZCI6IiJ9"

	// 使用base64.StdEncoding.DecodeString函数进行解码
	decodedBytes, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		fmt.Println("解码错误: ", err)
		return
	}

	fmt.Println("decodedBytes:", string(decodedBytes))
}
