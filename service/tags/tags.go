package tags

import "github.com/livegoplayer/go_video_store/domain/video_tag"

func TagList() map[string]interface{} {
	return video_tag.GetAllTags()
}
