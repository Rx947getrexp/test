package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mssola/user_agent"
	"go-speed/service"
	"go-speed/service/email"
	"go-speed/util"
	"math/rand"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNetworkDelay(t *testing.T) {
	url := "http://10.10.10.222:13001"
	fmt.Println(service.CheckUrlDelay(url), "ms")
}

func TestUserAgent(t *testing.T) {
	ua := user_agent.New("Mozilla/5.0 (Linux; U; Android 2.3.7; en-us; Nexus One Build/FRF91) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1")

	fmt.Printf("%v\n", ua.Mobile())  // => true
	fmt.Printf("%v\n", ua.Bot())     // => false
	fmt.Printf("%v\n", ua.Mozilla()) // => "5.0"

	fmt.Printf("%v\n", ua.Platform()) // => "Linux"
	fmt.Printf("%v\n", ua.OS())       // => "Android 2.3.7"

	name, version := ua.Engine()
	fmt.Printf("%v\n", name)    // => "AppleWebKit"
	fmt.Printf("%v\n", version) // => "533.1"

	name, version = ua.Browser()
	fmt.Printf("%v\n", name)    // => "Android"
	fmt.Printf("%v\n", version) // => "4.0"

	// Let's see an example with a bot.
	ua.Parse("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")

	fmt.Printf("%v\n", ua.Bot()) // => true

	name, version = ua.Browser()
	fmt.Printf("%v\n", name)    // => Googlebot
	fmt.Printf("%v\n", version) // => 2.1
}

func TestEmail(t *testing.T) {
	address := []string{"xiaomingchuan1990@gmail.com", "446117327@qq.com", "xiaomingchuan1990@gmail.com", "hthff96@mail.ru", "pikabuke@yandex.ru", "23344@internet.ru", "togxucuse@gmail.com"}
	//subject := "Speed密码找回"
	//body :=
	//	`<br>hello!</br>
	//<br>this is a test email, pls ignore it.</br>` + fmt.Sprintf(constant.SmsMsg, util.EncodeToString(6))
	////email.SendMail(address, subject, body)
	//subject := constant.ForgetSubject
	//body := fmt.Sprintf(constant.ForgetBody, util.EncodeToString(6))
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	defer c.Done()

	var (
		emailSubject = `VPN Короля сброс пароля`
		emailContent = "<br>【VPN Короля】</br>Код подтверждения: <font color='red'>%s</font>, действителен в течение 5 минут. Для обеспечения безопасности вашего аккаунта, пожалуйста, не раскрывайте эту информацию другим людям."
	)
	err := email.SendEmail(c, emailSubject, fmt.Sprintf(emailContent, util.EncodeToString(6)), address)

	//err := email.SendEmail(c, subject, body, address)
	fmt.Println(err)
}

func TestRand(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		println(rand.Intn(1))
	}
	//3-6范围，左闭右开
	for i := 0; i < 30; i++ {
		println(service.GenerateRangeNum(3, 6))
		time.Sleep(time.Second)
	}
}

func TestUuid(t *testing.T) {
	nonce, err := uuid.NewUUID()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(nonce.String())

	for i := 0; i <= 50; i++ {
		rnd := rand.New(rand.NewSource(int64(i)))
		uuid.SetRand(rnd)
		nonce2, _ := uuid.NewRandomFromReader(rnd)
		fmt.Println("nonce2:", nonce2.String())
	}

}

func TestPwd(t *testing.T) {
	pwd := util.AesEncrypt("654321")
	fmt.Println(pwd)
}
