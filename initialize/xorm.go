package initialize

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"go-speed/config"
	"go-speed/global"
	log_adaptor "go-speed/initialize/log-adaptor"
	"go-speed/util"
	"time"
	"xorm.io/xorm"
)

func initPgsqlDb(dbConfig config.Db) *xorm.Engine {
	if len(dbConfig.Host) == 0 {
		global.Logger.Warn().Msg("数据库参数未配置")
		return nil
	}
	password := dbConfig.Password
	dataSourceName := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		dbConfig.User, password, dbConfig.DbName, dbConfig.Host, dbConfig.Port)
	engine, err := xorm.NewEngine("postgres", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = engine.Ping(); err != nil {
		panic(err)
	}
	writer := util.GetLogWriter(dbConfig.Log)
	xormLogger := log_adaptor.NewXormLogger(writer)
	xormLogger.SetLevel2(dbConfig.Log.Level)
	if dbConfig.MaxIdleConn > 0 {
		engine.SetMaxIdleConns(dbConfig.MaxIdleConn)
	}
	engine.SetLogger(xormLogger)
	engine.EnableSessionID(global.Config.System.ShowSql)
	engine.ShowSQL(global.Config.System.ShowSql)
	engine.SetConnMaxLifetime(30 * time.Second)
	return engine
}

func initMysqlDb(dbConfig config.Db) *xorm.Engine {
	if len(dbConfig.Host) == 0 {
		global.Logger.Warn().Msg("数据库参数未配置")
		return nil
	}
	var password string
	if global.Config.System.Env == "release" {
		password = util.AesDecryptV2(dbConfig.Password)
	} else {
		password = dbConfig.Password
	}

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		dbConfig.User, password, dbConfig.Host, dbConfig.Port, dbConfig.DbName)
	engine, err := xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = engine.Ping(); err != nil {
		panic(err)
	}
	writer := util.GetLogWriter(dbConfig.Log)
	xormLogger := log_adaptor.NewXormLogger(writer)
	xormLogger.SetLevel2(dbConfig.Log.Level)
	if dbConfig.MaxIdleConn > 0 {
		engine.SetMaxIdleConns(dbConfig.MaxIdleConn)
	}
	engine.SetLogger(xormLogger)
	engine.EnableSessionID(global.Config.System.ShowSql)
	engine.ShowSQL(global.Config.System.ShowSql)
	engine.SetConnMaxLifetime(60 * time.Second)
	engine.SetMaxOpenConns(20)
	engine.SetMaxIdleConns(10)
	return engine
}
