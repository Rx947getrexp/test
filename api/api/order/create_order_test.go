package order

import (
	"fmt"
	"testing"
)

func TestGenerateOrderID(t *testing.T) {
	var orderId string
	orderId = generateOrderID()
	fmt.Println(orderId)

	orderId = generateOrderID()
	fmt.Println(orderId)
	orderId = generateOrderID()
	fmt.Println(orderId)
	orderId = generateOrderID()
	fmt.Println(orderId)
	orderId = generateOrderID()
	fmt.Println(orderId)
}
