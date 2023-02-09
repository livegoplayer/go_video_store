package log

import (
	"github.com/livegoplayer/go_db_helper/mysql"
)

const PREFIX = "vs_"

type Log struct {
	Data        string `gorm:"data" json:"data"`   //数据存根，使用json
	Level       int64  `gorm:"level" json:"level"` //存入level所代表的数字
	Cat         int64  `gorm:"cat" json:"cat"`     //0: 普通日志\r\n1：需要管理员介入的日志
	Id          int64  `gorm:"id;PRIMARY_KEY" json:"id"`
	Message     string `gorm:"message" json:"message"`           //日志记录
	AddDatetime string `gorm:"add_datetime" json:"add_datetime"` //操作时间
}

func (Log) TableName() string {
	return PREFIX + "log"
}

type LogQuery struct {
	mysql.Query
}

func (Log) Connect() string {
	return "video_store"
}
