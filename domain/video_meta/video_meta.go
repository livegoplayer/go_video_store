package video_meta

import (
	"github.com/livegoplayer/go_helper/utils"
	"github.com/livegoplayer/go_video_store/repository_template_model/video"
	"github.com/livegoplayer/go_video_store/repository_template_model/video_detail"
	"github.com/livegoplayer/go_video_store/repository_template_model/video_tag"
)

func GetVideoListByVideoIds(videoIds []int64, page int, pageSize int) []map[string]interface{} {
	videoList := video.FetchByVideoIdsWithPageSize(videoIds, page, pageSize)
	videoDetails := video_detail.FetchByVideoIds(videoIds)
	return utils.NewCollect(videoList, "json").WithKeyBy("details", utils.ToMapArray(videoDetails), "video_id", "json").ToMapArray()
}

func FetchByNameAndTagIds(search string, tagIds []int64, page int, pageSize int) []map[string]interface{} {
	videoList := video.VideoCollect{}
	var videoIds []int64
	if len(tagIds) > 0 {
		videoIds = append(videoIds, video_tag.FetchByTagIds(tagIds).PluckVideoId()...)
	}
	if len(search) > 0 {
		videoIds = append(videoIds, video.FetchByName(search).PluckVideoId()...)
	}
	videoIds = utils.Unique(videoIds)
	videoList = video.FetchByVideoIdsWithPageSize(videoIds, page, pageSize)
	videoDetails := video_detail.FetchByVideoIds(videoList.PluckVideoId())
	return utils.NewCollect(videoList, "json").WithGroupBy("details", utils.ToMapArray(videoDetails), "video_id", "json").ToMapArray()
}
