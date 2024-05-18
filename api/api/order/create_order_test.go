package order

import (
	"fmt"
	"testing"
	"time"
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

func TestGenUPayAmountDecimalPartValue(t *testing.T) {
	amount := genUPayAmountDecimalPartValue()
	fmt.Println(amount)

	amount = genUPayAmountDecimalPartValue()
	fmt.Println(amount)

	amount = genUPayAmountDecimalPartValue()
	fmt.Println(amount)

	amount = genUPayAmountDecimalPartValue()
	fmt.Println(amount)

	amount = genUPayAmountDecimalPartValue()
	fmt.Println(amount)

	amount = genUPayAmountDecimalPartValue()
	fmt.Println(amount)
}

func TestGenPayAmount(t *testing.T) {
	//fmt.Println(genPayAmount(10, 1.1111, constant.PayChannelUPay))
	//fmt.Println(genPayAmount(10, 1.1111, ""))
	//fmt.Println(genPayAmount(10, 10.111101998, constant.PayChannelUPay))

	fmt.Println(getNDurationAgoTime(time.Minute * 30))
}
