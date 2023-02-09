package video_detail

/* Unscoped *************************************** 满足常用的共有方法,根据mysql索引生成 **********************************************/
/**************************************** all 不是小表不要用 **************************************************/

func Save(f *VideoDetail) *VideoDetail {
	NewVideoDetailQuery().save(f)
	return f
}

func UpdateVideoDetailAll(p *VideoDetail) int64 {
	build := NewVideoDetailQuery()
	return build.update(p)
}

func FetchVideoDetailAll() VideoDetailCollect {
	build := NewVideoDetailQuery()
	return build.Get()
}

func FetchVideoDetailAllWithPageSize(page int, pageSize int) VideoDetailCollect {
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * pageSize

	build := NewVideoDetailQuery()
	return build.Skip(offset).Limit(pageSize).Get()
}

func CountVideoDetailAll() int64 {
	build := NewVideoDetailQuery()
	return build.Count()
}

/**************************************** 根据 Index 生成的方法 **************************************************/

func UpdateVideoDetailByIds(id []int64, p *VideoDetail) int64 {
	build := NewVideoDetailQuery()

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

func UpdateVideoDetailByIdsWhatEver(id []int64, p *VideoDetail) int64 {
	build := NewVideoDetailQuery()

	if len(id) == 1 {
		build.kWheId(id[0])
	} else {
		build.kWheIdIn(id)
	}

	return build.update(p)
}

func FetchByIdsWithPageSize(id []int64, page int, pageSize int) VideoDetailCollect {
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * pageSize

	build := NewVideoDetailQuery()

	if len(id) == 0 {
		return make(VideoDetailCollect, 0)
	}

	if len(id) == 1 {
		build.kWheId(id[0])
	} else {
		build.kWheIdIn(id)
	}

	return build.Skip(offset).Limit(pageSize).Get()
}

func CheckExistById(id int64) bool {
	build := NewVideoDetailQuery()
	build.kWheId(id)
	cnt := build.Count()
	return cnt > 0
}

func GetFirstById(id int64) *VideoDetail {
	build := NewVideoDetailQuery()
	build.kWheId(id)
	return build.First()
}

func DeleteById(id int64) int64 {
	build := NewVideoDetailQuery()
	build.kWheId(id)
	return build.Delete()
}

func GetOneById(id int64) *VideoDetail {
	build := NewVideoDetailQuery()
	build.kWheId(id)
	return build.GetOne()
}

func UpdateVideoDetailById(id int64, p *VideoDetail) int64 {
	build := NewVideoDetailQuery()
	build.kWheId(id)
	return build.update(p)
}

func CheckExistByVideoIdAndSeason(videoId int64, season int64) bool {
	build := NewVideoDetailQuery()
	build.kWheVideoId(videoId)
	build.kWheSeason(season)
	cnt := build.Count()
	return cnt > 0
}

func GetFirstByVideoIdAndSeason(videoId int64, season int64) *VideoDetail {
	build := NewVideoDetailQuery()
	build.kWheVideoId(videoId)
	build.kWheSeason(season)
	return build.First()
}

func DeleteAllByVideoIdAndSeason(videoId int64, season int64) int64 {
	build := NewVideoDetailQuery()
	build.kWheVideoId(videoId)
	build.kWheSeason(season)
	return build.Delete()
}

func GetOneByVideoIdAndSeason(videoId int64, season int64) *VideoDetail {
	build := NewVideoDetailQuery()
	build.kWheVideoId(videoId)
	build.kWheSeason(season)
	return build.GetOne()
}

func UpdateVideoDetailByVideoIdAndSeason(videoId int64, season int64, p *VideoDetail) int64 {
	build := NewVideoDetailQuery()
	build.kWheVideoId(videoId)
	build.kWheSeason(season)
	return build.update(p)
}

func FetchByVideoIdAndSeason(videoId int64, season int64) VideoDetailCollect {
	build := NewVideoDetailQuery()
	build.kWheVideoId(videoId)
	build.kWheSeason(season)
	return build.Get()
}

func FetchByVideoIdAndSeasonWithPageSize(videoId int64, season int64, page int, pageSize int) VideoDetailCollect {
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * pageSize

	build := NewVideoDetailQuery()
	return build.Skip(offset).Limit(pageSize).Get()
}

func CheckExistByVideoId(videoId int64) bool {
	build := NewVideoDetailQuery()
	build.kWheVideoId(videoId)
	cnt := build.Count()
	return cnt > 0
}

func GetFirstByVideoId(videoId int64) *VideoDetail {
	build := NewVideoDetailQuery()
	build.kWheVideoId(videoId)
	return build.First()
}

func DeleteAllByVideoId(videoId int64) int64 {
	build := NewVideoDetailQuery()
	build.kWheVideoId(videoId)
	return build.Delete()
}

func GetOneByVideoId(videoId int64) *VideoDetail {
	build := NewVideoDetailQuery()
	build.kWheVideoId(videoId)
	return build.GetOne()
}

func UpdateVideoDetailByVideoId(videoId int64, p *VideoDetail) int64 {
	build := NewVideoDetailQuery()
	build.kWheVideoId(videoId)
	return build.update(p)
}

func FetchByVideoId(videoId int64) VideoDetailCollect {
	build := NewVideoDetailQuery()
	build.kWheVideoId(videoId)
	return build.Get()
}

func FetchByVideoIdWithPageSize(videoId int64, page int, pageSize int) VideoDetailCollect {
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * pageSize

	build := NewVideoDetailQuery()
	return build.Skip(offset).Limit(pageSize).Get()
}

func FetchByVideoIds(videoId []int64) VideoDetailCollect {
	build := NewVideoDetailQuery()

	if len(videoId) == 0 {
		return make(VideoDetailCollect, 0)
	}

	if len(videoId) == 1 {
		build.kWheVideoId(videoId[0])
	} else {
		build.kWheVideoIdIn(videoId)
	}

	return build.Get()
}

func UpdateVideoDetailByVideoIds(videoId []int64, p *VideoDetail) int64 {
	build := NewVideoDetailQuery()

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

func FetchByVideoIdsWithPageSize(videoId []int64, page int, pageSize int) VideoDetailCollect {
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * pageSize

	build := NewVideoDetailQuery()

	if len(videoId) == 0 {
		return make(VideoDetailCollect, 0)
	}

	if len(videoId) == 1 {
		build.kWheVideoId(videoId[0])
	} else {
		build.kWheVideoIdIn(videoId)
	}

	return build.Skip(offset).Limit(pageSize).Get()
}

func UpdateVideoDetailByVideoIdsWhatEver(videoId []int64, p *VideoDetail) int64 {
	build := NewVideoDetailQuery()

	if len(videoId) == 1 {
		build.kWheVideoId(videoId[0])
	} else {
		build.kWheVideoIdIn(videoId)
	}

	return build.update(p)
}

func CountVideoDetailByVideoIds(videoId []int64) int64 {
	if len(videoId) == 0 {
		return 0
	}
	build := NewVideoDetailQuery()
	if len(videoId) == 1 {
		build.kWheVideoId(videoId[0])
	} else {
		build.kWheVideoIdIn(videoId)
	}

	return build.Count()
}
