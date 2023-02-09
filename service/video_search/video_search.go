package video_search

import (
	"github.com/livegoplayer/go_video_store/domain/video_meta"
	"github.com/livegoplayer/go_video_store/repository_template_model/video"
)

func SearchByCondition(videoIds []int64, searchStr string, tagIds []int64, page int64, pageSize int64) video.VideoCollect {
	if len(videoIds) > 0 {
		return video_meta.GetVideoListByVideoIds(videoIds, page, pageSize)
	}

	return video_meta.FetchByNameAndTagIds(searchStr, tagIds, int(page), int(pageSize))
}
