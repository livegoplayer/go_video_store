package video_resources

import (
	"github.com/livegoplayer/go_db_helper/mysql"
)

const PREFIX = "vs_"

type VideoResources struct {
	LastUpdateTime   int64  `gorm:"last_update_time" json:"last_update_time"`     //上次更新时间，一般获取某个video_id的最后一集的更新时间就是该视频的更新时间，其他的没用，这个是用来判断是否爬取最新信息的依据之一，其他还有剧集的集数等等， 网站的最后更新时间解析出来都是一样的，所以直接用大小判断即可
	VideoUrl         string `gorm:"video_url" json:"video_url"`                   //视频原来的播放地址
	Data             string `gorm:"data" json:"data"`                             //数据，一般存储的是有益于获取视频播放源的数据    \r\nddys.art:  有 id
	VideoReferDomain string `gorm:"video_refer_domain" json:"video_refer_domain"` //视频原来的网站地址
	UptDatetime      string `gorm:"upt_datetime" json:"upt_datetime"`
	Name             string `gorm:"name" json:"name"`       //源网站对该集或者视频的播放item描述
	Season           int64  `gorm:"season" json:"season"`   //季 * 100 如果要得出正确的季/100 就行了 \r\n有编号的是特别篇
	Episode          int64  `gorm:"episode" json:"episode"` //集，分属季的集
	Id               int64  `gorm:"id;PRIMARY_KEY" json:"id"`
	VideoRefer       int64  `gorm:"video_refer" json:"video_refer"` //1： ddys.art
	AddDatetime      string `gorm:"add_datetime" json:"add_datetime"`
	VideoId          int64  `gorm:"video_id" json:"video_id"`
}

func (VideoResources) TableName() string {
	return PREFIX + "video_resources"
}

type VideoResourcesQuery struct {
	mysql.Query
}

func (VideoResources) Connect() string {
	return "video_store"
}
