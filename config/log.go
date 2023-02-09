package config

import (
	"github.com/gin-gonic/gin"
	c "github.com/livegoplayer/go_gin_helper"
	myLogger "github.com/livegoplayer/go_logger/logger"
	"github.com/livegoplayer/go_logger/logger/loggers"
	"github.com/spf13/viper"
	"time"
)

func InitLog() {
	//初始化 app_log， 以后使用mylogger.Info() 打印log
	//如果是debug模式的话，直接打印到控制台
	logType := viper.GetString("log.log_type")
	if logType == "file" {
		logFilePath := viper.GetString("log.file.path")
		cleanTimeHour := viper.GetInt64("log.file.clean_time_hour")
		splitTimeHour := viper.GetInt64("log.file.split_time_hour")
		myLogger.InitBaseFileLogByPath(logFilePath, time.Duration(cleanTimeHour*int64(time.Hour)), time.Duration(splitTimeHour*int64(time.Hour)))
	} else if logType == "mysql" {
		username := viper.GetString("log.mysql.username")
		password := viper.GetString("log.mysql.password")
		host := viper.GetString("log.mysql.host")
		port := viper.GetString("log.mysql.port")
		dbname := viper.GetString("log.mysql.dbname")
		tableName := viper.GetString("log.mysql.table_name")
		myLogger.InitBaseMysqlLogByConfig(host, port, dbname, tableName, username, password)
	} else if logType == "rabbitmq" {
		url := viper.GetString("log.rabbitmq.url")
		routingKey := viper.GetString("log.rabbitmq.routingKey")
		exchange := viper.GetString("log.rabbitmq.exchange")
		appType := viper.GetInt("log.rabbitmq.app_type")
		myLogger.InitBaseRabbitmqLogByConfig(url, routingKey, exchange, appType)
	}

	if gin.IsDebugging() {
		myLogger.AddDebugLogHook()
	}
}

func InitAccessLog(r *gin.Engine) *gin.Engine {
	//gin的格式化参数
	//初始化access_log
	r.Use(loggers.GetGinAccessFileLogger(viper.GetString("log.access_log_file_path"), ""))
	//如果是debug模式的话，使用logger另外打印一份输出到控制台的logger
	if gin.IsDebugging() {
		// 额外在控制台输出一些accesslog
		r.Use(loggers.GetDebugLogger())
	}
	r.Use(c.ErrHandler())
	return r
}
