// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"go-speed/dao/internal"
)

// internalTPaymentChannelDao is internal type for wrapping internal DAO implements.
type internalTPaymentChannelDao = *internal.TPaymentChannelDao

// tPaymentChannelDao is the data access object for table t_payment_channel.
// You can define custom methods on it to extend its functionality as you wish.
type tPaymentChannelDao struct {
	internalTPaymentChannelDao
}

var (
	// TPaymentChannel is globally public accessible object for table t_payment_channel operations.
	TPaymentChannel = tPaymentChannelDao{
		internal.NewTPaymentChannelDao(),
	}
)

// Fill with you ideas below.
