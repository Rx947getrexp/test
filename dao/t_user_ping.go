// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"go-speed/dao/internal"
)

// internalTUserPingDao is internal type for wrapping internal DAO implements.
type internalTUserPingDao = *internal.TUserPingDao

// tUserPingDao is the data access object for table t_user_ping.
// You can define custom methods on it to extend its functionality as you wish.
type tUserPingDao struct {
	internalTUserPingDao
}

var (
	// TUserPing is globally public accessible object for table t_user_ping operations.
	TUserPing = tUserPingDao{
		internal.NewTUserPingDao(),
	}
)

// Fill with you ideas below.
