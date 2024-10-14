package model

import "time"

type TPayOrder struct {
	Id                 int64
	UserId             uint64
	Email              string
	OrderNo            string
	PaymentChannelId   string
	GoodsId            int
	OrderAmount        string
	Currency           string
	PayTypeCode        string
	Status             string
	ReturnStatus       string
	StatusMes          string
	OrderData          string
	ResultStatus       string
	OrderRealityAmount string
	PaymentProof       string
	PaymentChannelErr  int
	CreatedAt          time.Time
	UpdatedAt          time.Time
	Version            int
	Commission         float64
	DeviceType         string
}
