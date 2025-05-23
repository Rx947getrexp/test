// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"go-speed/dao/internal"
)

// internalTUserTrafficLogDao is internal type for wrapping internal DAO implements.
type internalTUserTrafficLogDao = *internal.TUserTrafficLogDao

// tUserTrafficLogDao is the data access object for table t_user_traffic_log.
// You can define custom methods on it to extend its functionality as you wish.
type tUserTrafficLogDao struct {
	internalTUserTrafficLogDao
}

var (
	// TUserTrafficLog is globally public accessible object for table t_user_traffic_log operations.
	TUserTrafficLog = tUserTrafficLogDao{
		internal.NewTUserTrafficLogDao(),
	}
)

// Fill with you ideas below.
