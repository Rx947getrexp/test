package main

import (
	"fmt"
	"github.com/mssola/user_agent"
	"go-speed/service"
	"go-speed/service/email"
	"testing"
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
	address := []string{"aaa@qq.com", "aaa@163.com"}
	subject := "test mail"
	body :=
		`<br>hello!</br>
  <br>this is a test email, pls ignore it.</br>`
	email.SendMail(address, subject, body)
}
