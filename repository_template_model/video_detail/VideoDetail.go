package video_detail

import (
	"github.com/livegoplayer/go_db_helper/mysql"
)

const PREFIX = "vs_"

type VideoDetail struct {
	OfficialWebsite  string  `gorm:"official_website" json:"official_website"`   //视频官网地址
	DescriptionUrl   string  `gorm:"description_url" json:"description_url"`     //豆瓣介绍页面或者其他
	Year             int64   `gorm:"year" json:"year"`                           //年份，tag库存储1+个年份，第一个是该年份，其他年份取自release_date
	ReleasedDateStr  string  `gorm:"released_date_str" json:"released_date_str"` //上映时间 str tag库中的release_date是去除地区的，这里做个缓存
	VideoPosterUrl   string  `gorm:"video_poster_url" json:"video_poster_url"`   //视频的海报
	TimeStr          string  `gorm:"time_str" json:"time_str"`                   //时间str
	PosterCache      string  `gorm:"poster_cache" json:"poster_cache"`           //豆瓣缓存封面图片
	VideoDescription string  `gorm:"video_description" json:"video_description"` //视频简介
	Score            float64 `gorm:"score" json:"score"`                         //豆瓣评分，仅爬取的时候的分数
	PosterListUrl    string  `gorm:"poster_list_url" json:"poster_list_url"`     //豆瓣电影剧照页面
	UptDatetime      string  `gorm:"upt_datetime" json:"upt_datetime"`
	PerUpdateTime    string  `gorm:"per_update_time" json:"per_update_time"` //每次更新时间 ddys: 周几(深夜/凌晨)
	Season           int64   `gorm:"season" json:"season"`                   //季 * 100 如果要得出正确的季/100 就行了 \r\n有编号的是特别篇
	Id               int64   `gorm:"id" json:"id"`
	SetNum           int64   `gorm:"set_num" json:"set_num"` //本季集数，仅第一次爬取时候的集数，仅豆瓣有的时候显示为最大， 为0的话表示没有该信息，前端可以显示为当前集数
	AddDatetime      string  `gorm:"add_datetime" json:"add_datetime"`
	VideoId          int64   `gorm:"video_id" json:"video_id"`
}

func (VideoDetail) TableName() string {
	return PREFIX + "video_detail"
}

type VideoDetailQuery struct {
	mysql.Query
}

func (VideoDetail) Connect() string {
	return "video_store"
}
