package main

import (
	"fmt"
	"go-speed/service"
	"testing"
)

func TestNetworkDelay(t *testing.T) {
	url := "http://10.10.10.222:13001"
	fmt.Println(service.CheckUrlDelay(url), "ms")
}
