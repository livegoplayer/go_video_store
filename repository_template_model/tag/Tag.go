package tag

import (
	"github.com/livegoplayer/go_db_helper/mysql"
)

const PREFIX = "vs_"

type Tag struct {
	UptDatetime string `gorm:"upt_datetime" json:"upt_datetime"`
	TagName     string `gorm:"tag_name" json:"tag_name"` //名称唯一
	TagId       int64  `gorm:"tag_id" json:"tag_id"`     //0-100 视频大类 \r\n100-200 电影分类 \r\n200-300 电影国家\r\n300-400 电影语言\r\n1930-3000 电影年份 特殊，预留，有一个电影插入一个\r\n\r\n3000000000-4294967295 电影演员列表\r\n1000000000-200000000 电影导演列表\r\n2000000000-300000000 电影编剧列表\r\n
	Id          int64  `gorm:"id" json:"id"`
	TagUrl      string `gorm:"tag_url" json:"tag_url"` //有些tag是有外链的
	AddDatetime string `gorm:"add_datetime" json:"add_datetime"`
}

func (Tag) TableName() string {
	return PREFIX + "tag"
}

type TagQuery struct {
	mysql.Query
}

func (Tag) Connect() string {
	return "video_store"
}
