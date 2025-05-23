// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"go-speed/dao/internal"
)

// internalTAppDnsDao is internal type for wrapping internal DAO implements.
type internalTAppDnsDao = *internal.TAppDnsDao

// tAppDnsDao is the data access object for table t_app_dns.
// You can define custom methods on it to extend its functionality as you wish.
type tAppDnsDao struct {
	internalTAppDnsDao
}

var (
	// TAppDns is globally public accessible object for table t_app_dns operations.
	TAppDns = tAppDnsDao{
		internal.NewTAppDnsDao(),
	}
)

// Fill with you ideas below.
