package video_tag

import (
	"github.com/livegoplayer/go_db_helper/mysql"
)

const PREFIX = "vs_"

type VideoTag struct {
	UptDatetime string `gorm:"upt_datetime" json:"upt_datetime"`
	TagId       int64  `gorm:"tag_id" json:"tag_id"`
	Id          int64  `gorm:"id" json:"id"`
	AddDatetime string `gorm:"add_datetime" json:"add_datetime"`
	VideoId     int64  `gorm:"video_id" json:"video_id"`
}

func (VideoTag) TableName() string {
	return PREFIX + "video_tag"
}

type VideoTagQuery struct {
	mysql.Query
}

func (VideoTag) Connect() string {
	return "video_store"
}
