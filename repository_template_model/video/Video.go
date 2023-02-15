package video

import (
	"github.com/livegoplayer/go_db_helper/mysql"
)

const PREFIX = "vs_"

type Video struct {
	PrettyName  string `gorm:"pretty_name" json:"pretty_name"` //用来显示的name
	MetaName    string `gorm:"meta_name" json:"meta_name"`     //分国别，用电影原来的名称，确保唯一
	UptDatetime string `gorm:"upt_datetime" json:"upt_datetime"`
	AddDatetime string `gorm:"add_datetime" json:"add_datetime"`
	VideoId     int64  `gorm:"video_id;PRIMARY_KEY" json:"video_id"`
}

func (Video) TableName() string {
	return PREFIX + "video"
}

type VideoQuery struct {
	mysql.Query
}

func (Video) Connect() string {
	return "video_store"
}
