package video_meta

import (
	"github.com/livegoplayer/video_store/repository_template_model/video"
	"github.com/livegoplayer/video_store/repository_template_model/video_tag"
)

func GetVideoListByVideoIds(videoIds []int64, page int64, pageSize int64) video.VideoCollect {
	list := video.VideoCollect{}
	if len(videoIds) == 1 {
		list = append(list, *getVideoInfoByVideoId(videoIds[0]))
		return list
	}
	return video.FetchByVideoIdsWithPageSize(videoIds, int(page), int(pageSize))
}

func getVideoInfoByVideoId(videoId int64) *video.Video {
	return video.GetFirstByVideoId(videoId)
}

func FetchByNameAndTagIds(search string, tagIds []int64, page int, pageSize int) video.VideoCollect {
	videoList := video.VideoCollect{}
	var videoIds []int64
	if len(tagIds) > 0 {
		videoIds = append(videoIds, video_tag.FetchByTagIds(tagIds).PluckVideoId()...)
	}
	if len(search) > 0 {
		videoIds = append(videoIds, video.FetchByName(search).PluckVideoId()...)
	}
	videoIds = Unique(videoIds)
	if len(videoIds) > 0 {
		videoList = video.FetchByVideoIdsWithPageSize(videoIds, page, pageSize)
	}
	return videoList
}

// 定义一个新的类型：element
type element interface {
	// element 支持如下类型
	string | int8 | int16 | int32 | int64 | int | float32 | float64 | uint | uint8 | uint16 | uint32 | uint64
}

// Unique 数组或slice边去重
func Unique[T element](arr []T) []T {
	tmp := make(map[T]struct{})
	l := len(arr)
	if l == 0 {
		return arr
	}

	rel := make([]T, 0, l)
	for _, item := range arr {
		_, ok := tmp[item]
		if ok {
			continue
		}
		tmp[item] = struct{}{}
		rel = append(rel, item)
	}

	return rel[:len(tmp)]
}
