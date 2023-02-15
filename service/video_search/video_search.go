package video_search

import (
	"github.com/livegoplayer/go_video_store/domain/video_meta"
)

func SearchByCondition(videoIds []int64, searchStr string, tagIds []int64, page int, pageSize int) []map[string]interface{} {
	if len(videoIds) > 0 {
		return video_meta.GetVideoListByVideoIds(videoIds, page, pageSize)
	}

	return video_meta.FetchByNameAndTagIds(searchStr, tagIds, page, pageSize)
}
