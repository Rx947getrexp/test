// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"go-speed/dao/internal"
)

// internalTPayOrderDao is internal type for wrapping internal DAO implements.
type internalTPayOrderDao = *internal.TPayOrderDao

// tPayOrderDao is the data access object for table t_pay_order.
// You can define custom methods on it to extend its functionality as you wish.
type tPayOrderDao struct {
	internalTPayOrderDao
}

var (
	// TPayOrder is globally public accessible object for table t_pay_order operations.
	TPayOrder = tPayOrderDao{
		internal.NewTPayOrderDao(),
	}
)

// Fill with you ideas below.
