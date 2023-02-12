package video_tag

/* Unscoped *************************************** 满足常用的共有方法,根据mysql索引生成 **********************************************/
/**************************************** all 不是小表不要用 **************************************************/

func Save(f *VideoTag) *VideoTag {
	NewVideoTagQuery().save(f)
	return f
}

func UpdateVideoTagAll(p *VideoTag) int64 {
	build := NewVideoTagQuery()
	return build.update(p)
}

func FetchVideoTagAll() VideoTagCollect {
	build := NewVideoTagQuery()
	return build.Get()
}

func FetchVideoTagAllWithPageSize(page int, pageSize int) VideoTagCollect {
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * pageSize

	build := NewVideoTagQuery()
	return build.Skip(offset).Limit(pageSize).Get()
}

func CountVideoTagAll() int64 {
	build := NewVideoTagQuery()
	return build.Count()
}

/**************************************** 根据 Index 生成的方法 **************************************************/

func CheckExistByVideoIdAndTagId(videoId int64, tagId int64) bool {
	build := NewVideoTagQuery()
	build.kWheVideoId(videoId)
	build.kWheTagId(tagId)
	cnt := build.Count()
	return cnt > 0
}

func GetFirstByVideoIdAndTagId(videoId int64, tagId int64) *VideoTag {
	build := NewVideoTagQuery()
	build.kWheVideoId(videoId)
	build.kWheTagId(tagId)
	return build.First()
}

func DeleteByVideoIdAndTagId(videoId int64, tagId int64) int64 {
	build := NewVideoTagQuery()
	build.kWheVideoId(videoId)
	build.kWheTagId(tagId)
	return build.Delete()
}

func GetOneByVideoIdAndTagId(videoId int64, tagId int64) *VideoTag {
	build := NewVideoTagQuery()
	build.kWheVideoId(videoId)
	build.kWheTagId(tagId)
	return build.GetOne()
}

func UpdateVideoTagByVideoIdAndTagId(videoId int64, tagId int64, p *VideoTag) int64 {
	build := NewVideoTagQuery()
	build.kWheVideoId(videoId)
	build.kWheTagId(tagId)
	return build.update(p)
}

func CheckExistByVideoId(videoId int64) bool {
	build := NewVideoTagQuery()
	build.kWheVideoId(videoId)
	cnt := build.Count()
	return cnt > 0
}

func GetFirstByVideoId(videoId int64) *VideoTag {
	build := NewVideoTagQuery()
	build.kWheVideoId(videoId)
	return build.First()
}

func DeleteAllByVideoId(videoId int64) int64 {
	build := NewVideoTagQuery()
	build.kWheVideoId(videoId)
	return build.Delete()
}

func GetOneByVideoId(videoId int64) *VideoTag {
	build := NewVideoTagQuery()
	build.kWheVideoId(videoId)
	return build.GetOne()
}

func UpdateVideoTagByVideoId(videoId int64, p *VideoTag) int64 {
	build := NewVideoTagQuery()
	build.kWheVideoId(videoId)
	return build.update(p)
}

func FetchByVideoId(videoId int64) VideoTagCollect {
	build := NewVideoTagQuery()
	build.kWheVideoId(videoId)
	return build.Get()
}

func FetchByVideoIdWithPageSize(videoId int64, page int, pageSize int) VideoTagCollect {
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * pageSize

	build := NewVideoTagQuery()
	return build.Skip(offset).Limit(pageSize).Get()
}

func FetchByVideoIds(videoId []int64) VideoTagCollect {
	build := NewVideoTagQuery()

	if len(videoId) == 0 {
		return make(VideoTagCollect, 0)
	}

	if len(videoId) == 1 {
		build.kWheVideoId(videoId[0])
	} else {
		build.kWheVideoIdIn(videoId)
	}

	return build.Get()
}

func UpdateVideoTagByVideoIds(videoId []int64, p *VideoTag) int64 {
	build := NewVideoTagQuery()

	if len(videoId) == 0 {
		return 0
	}

	if len(videoId) == 1 {
		build.kWheVideoId(videoId[0])
	} else {
		build.kWheVideoIdIn(videoId)
	}

	return build.update(p)
}

func FetchByVideoIdsWithPageSize(videoId []int64, page int, pageSize int) VideoTagCollect {
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * pageSize

	build := NewVideoTagQuery()

	if len(videoId) == 0 {
		return make(VideoTagCollect, 0)
	}

	if len(videoId) == 1 {
		build.kWheVideoId(videoId[0])
	} else {
		build.kWheVideoIdIn(videoId)
	}

	return build.Skip(offset).Limit(pageSize).Get()
}

func UpdateVideoTagByVideoIdsWhatEver(videoId []int64, p *VideoTag) int64 {
	build := NewVideoTagQuery()

	if len(videoId) == 1 {
		build.kWheVideoId(videoId[0])
	} else {
		build.kWheVideoIdIn(videoId)
	}

	return build.update(p)
}

func CountVideoTagByVideoIds(videoId []int64) int64 {
	if len(videoId) == 0 {
		return 0
	}
	build := NewVideoTagQuery()
	if len(videoId) == 1 {
		build.kWheVideoId(videoId[0])
	} else {
		build.kWheVideoIdIn(videoId)
	}

	return build.Count()
}
func UpdateVideoTagByIds(id []int64, p *VideoTag) int64 {
	build := NewVideoTagQuery()

	if len(id) == 0 {
		return 0
	}

	if len(id) == 1 {
		build.kWheId(id[0])
	} else {
		build.kWheIdIn(id)
	}

	return build.update(p)
}

func UpdateVideoTagByIdsWhatEver(id []int64, p *VideoTag) int64 {
	build := NewVideoTagQuery()

	if len(id) == 1 {
		build.kWheId(id[0])
	} else {
		build.kWheIdIn(id)
	}

	return build.update(p)
}

func FetchByIdsWithPageSize(id []int64, page int, pageSize int) VideoTagCollect {
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * pageSize

	build := NewVideoTagQuery()

	if len(id) == 0 {
		return make(VideoTagCollect, 0)
	}

	if len(id) == 1 {
		build.kWheId(id[0])
	} else {
		build.kWheIdIn(id)
	}

	return build.Skip(offset).Limit(pageSize).Get()
}

func CheckExistById(id int64) bool {
	build := NewVideoTagQuery()
	build.kWheId(id)
	cnt := build.Count()
	return cnt > 0
}

func GetFirstById(id int64) *VideoTag {
	build := NewVideoTagQuery()
	build.kWheId(id)
	return build.First()
}

func DeleteById(id int64) int64 {
	build := NewVideoTagQuery()
	build.kWheId(id)
	return build.Delete()
}

func GetOneById(id int64) *VideoTag {
	build := NewVideoTagQuery()
	build.kWheId(id)
	return build.GetOne()
}

func UpdateVideoTagById(id int64, p *VideoTag) int64 {
	build := NewVideoTagQuery()
	build.kWheId(id)
	return build.update(p)
}

func CheckExistByTagId(tagId int64) bool {
	build := NewVideoTagQuery()
	build.kWheTagId(tagId)
	cnt := build.Count()
	return cnt > 0
}

func GetFirstByTagId(tagId int64) *VideoTag {
	build := NewVideoTagQuery()
	build.kWheTagId(tagId)
	return build.First()
}

func DeleteAllByTagId(tagId int64) int64 {
	build := NewVideoTagQuery()
	build.kWheTagId(tagId)
	return build.Delete()
}

func GetOneByTagId(tagId int64) *VideoTag {
	build := NewVideoTagQuery()
	build.kWheTagId(tagId)
	return build.GetOne()
}

func UpdateVideoTagByTagId(tagId int64, p *VideoTag) int64 {
	build := NewVideoTagQuery()
	build.kWheTagId(tagId)
	return build.update(p)
}

func FetchByTagId(tagId int64) VideoTagCollect {
	build := NewVideoTagQuery()
	build.kWheTagId(tagId)
	return build.Get()
}

func FetchByTagIdWithPageSize(tagId int64, page int, pageSize int) VideoTagCollect {
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * pageSize

	build := NewVideoTagQuery()
	return build.Skip(offset).Limit(pageSize).Get()
}

func FetchByTagIds(tagId []int64) VideoTagCollect {
	build := NewVideoTagQuery()

	if len(tagId) == 0 {
		return make(VideoTagCollect, 0)
	}

	if len(tagId) == 1 {
		build.kWheTagId(tagId[0])
	} else {
		build.kWheTagIdIn(tagId)
	}

	return build.Get()
}

func UpdateVideoTagByTagIds(tagId []int64, p *VideoTag) int64 {
	build := NewVideoTagQuery()

	if len(tagId) == 0 {
		return 0
	}

	if len(tagId) == 1 {
		build.kWheTagId(tagId[0])
	} else {
		build.kWheTagIdIn(tagId)
	}

	return build.update(p)
}

func FetchByTagIdsWithPageSize(tagId []int64, page int, pageSize int) VideoTagCollect {
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * pageSize

	build := NewVideoTagQuery()

	if len(tagId) == 0 {
		return make(VideoTagCollect, 0)
	}

	if len(tagId) == 1 {
		build.kWheTagId(tagId[0])
	} else {
		build.kWheTagIdIn(tagId)
	}

	return build.Skip(offset).Limit(pageSize).Get()
}

func UpdateVideoTagByTagIdsWhatEver(tagId []int64, p *VideoTag) int64 {
	build := NewVideoTagQuery()

	if len(tagId) == 1 {
		build.kWheTagId(tagId[0])
	} else {
		build.kWheTagIdIn(tagId)
	}

	return build.update(p)
}

func CountVideoTagByTagIds(tagId []int64) int64 {
	if len(tagId) == 0 {
		return 0
	}
	build := NewVideoTagQuery()
	if len(tagId) == 1 {
		build.kWheTagId(tagId[0])
	} else {
		build.kWheTagIdIn(tagId)
	}

	return build.Count()
}
