package video_resources

/* Unscoped *************************************** 满足常用的共有方法,根据mysql索引生成 **********************************************/
/**************************************** all 不是小表不要用 **************************************************/

func Save(f *VideoResources) *VideoResources {
	NewVideoResourcesQuery().save(f)
	return f
}

func UpdateVideoResourcesAll(p *VideoResources) int64 {
	build := NewVideoResourcesQuery()
	return build.update(p)
}

func FetchVideoResourcesAll(order ...string) VideoResourcesCollect {
	build := NewVideoResourcesQuery()
	build.OrderBy(parseOrderString(order))
	return build.Get()
}

func FetchVideoResourcesAllWithPageSize(page int, pageSize int, order ...string) VideoResourcesCollect {
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * pageSize

	build := NewVideoResourcesQuery()
	build.OrderBy(parseOrderString(order))
	return build.Skip(offset).Limit(pageSize).Get()
}

func CountVideoResourcesAll() int64 {
	build := NewVideoResourcesQuery()
	return build.Count()
}

/**************************************** 根据 Index 生成的方法 **************************************************/

func CheckExistByVideoId(videoId int64) bool {
	build := NewVideoResourcesQuery()
	build.kWheVideoId(videoId)
	cnt := build.Count()
	return cnt > 0
}

func GetFirstByVideoId(videoId int64) *VideoResources {
	build := NewVideoResourcesQuery()
	build.kWheVideoId(videoId)
	return build.First()
}

func DeleteAllByVideoId(videoId int64) int64 {
	build := NewVideoResourcesQuery()
	build.kWheVideoId(videoId)
	return build.Delete()
}

func GetOneByVideoId(videoId int64) *VideoResources {
	build := NewVideoResourcesQuery()
	build.kWheVideoId(videoId)
	return build.GetOne()
}

func UpdateVideoResourcesByVideoId(videoId int64, p *VideoResources) int64 {
	build := NewVideoResourcesQuery()
	build.kWheVideoId(videoId)
	return build.update(p)
}

func FetchByVideoId(videoId int64, order ...string) VideoResourcesCollect {
	build := NewVideoResourcesQuery()
	build.kWheVideoId(videoId)
	build.OrderBy(parseOrderString(order))
	return build.Get()
}

func FetchByVideoIdWithPageSize(videoId int64, page int, pageSize int, order ...string) VideoResourcesCollect {
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * pageSize

	build := NewVideoResourcesQuery()
	build.kWheVideoId(videoId)
	build.OrderBy(parseOrderString(order))
	return build.Skip(offset).Limit(pageSize).Get()
}

func FetchByVideoIds(videoId []int64, order ...string) VideoResourcesCollect {
	build := NewVideoResourcesQuery()

	if len(videoId) == 0 {
		return make(VideoResourcesCollect, 0)
	}

	if len(videoId) == 1 {
		build.kWheVideoId(videoId[0])
	} else {
		build.kWheVideoIdIn(videoId)
	}

	build.OrderBy(parseOrderString(order))
	return build.Get()
}

func UpdateVideoResourcesByVideoIds(videoId []int64, p *VideoResources) int64 {
	build := NewVideoResourcesQuery()

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

func FetchByVideoIdsWithPageSize(videoId []int64, page int, pageSize int, order ...string) VideoResourcesCollect {
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * pageSize

	build := NewVideoResourcesQuery()

	if len(videoId) == 0 {
		return make(VideoResourcesCollect, 0)
	}

	if len(videoId) == 1 {
		build.kWheVideoId(videoId[0])
	} else {
		build.kWheVideoIdIn(videoId)
	}

	build.OrderBy(parseOrderString(order))
	return build.Skip(offset).Limit(pageSize).Get()
}

func UpdateVideoResourcesByVideoIdsWhatEver(videoId []int64, p *VideoResources) int64 {
	build := NewVideoResourcesQuery()

	if len(videoId) == 1 {
		build.kWheVideoId(videoId[0])
	} else {
		build.kWheVideoIdIn(videoId)
	}

	return build.update(p)
}

func CountVideoResourcesByVideoIds(videoId []int64) int64 {
	if len(videoId) == 0 {
		return 0
	}
	build := NewVideoResourcesQuery()
	if len(videoId) == 1 {
		build.kWheVideoId(videoId[0])
	} else {
		build.kWheVideoIdIn(videoId)
	}

	return build.Count()
}

func CheckExistByVideoIdAndSeason(videoId int64, season int64) bool {
	build := NewVideoResourcesQuery()
	build.kWheVideoId(videoId)
	build.kWheSeason(season)
	cnt := build.Count()
	return cnt > 0
}

func GetFirstByVideoIdAndSeason(videoId int64, season int64) *VideoResources {
	build := NewVideoResourcesQuery()
	build.kWheVideoId(videoId)
	build.kWheSeason(season)
	return build.First()
}

func DeleteAllByVideoIdAndSeason(videoId int64, season int64) int64 {
	build := NewVideoResourcesQuery()
	build.kWheVideoId(videoId)
	build.kWheSeason(season)
	return build.Delete()
}

func GetOneByVideoIdAndSeason(videoId int64, season int64) *VideoResources {
	build := NewVideoResourcesQuery()
	build.kWheVideoId(videoId)
	build.kWheSeason(season)
	return build.GetOne()
}

func UpdateVideoResourcesByVideoIdAndSeason(videoId int64, season int64, p *VideoResources) int64 {
	build := NewVideoResourcesQuery()
	build.kWheVideoId(videoId)
	build.kWheSeason(season)
	return build.update(p)
}

func FetchByVideoIdAndSeason(videoId int64, season int64, order ...string) VideoResourcesCollect {
	build := NewVideoResourcesQuery()
	build.kWheVideoId(videoId)
	build.kWheSeason(season)
	build.OrderBy(parseOrderString(order))
	return build.Get()
}

func FetchByVideoIdAndSeasonWithPageSize(videoId int64, season int64, page int, pageSize int, order ...string) VideoResourcesCollect {
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * pageSize

	build := NewVideoResourcesQuery()
	build.kWheVideoId(videoId)
	build.kWheSeason(season)
	build.OrderBy(parseOrderString(order))
	return build.Skip(offset).Limit(pageSize).Get()
}

func CheckExistByVideoIdAndSeasonAndEpisode(videoId int64, season int64, episode int64) bool {
	build := NewVideoResourcesQuery()
	build.kWheVideoId(videoId)
	build.kWheSeason(season)
	build.kWheEpisode(episode)
	cnt := build.Count()
	return cnt > 0
}

func GetFirstByVideoIdAndSeasonAndEpisode(videoId int64, season int64, episode int64) *VideoResources {
	build := NewVideoResourcesQuery()
	build.kWheVideoId(videoId)
	build.kWheSeason(season)
	build.kWheEpisode(episode)
	return build.First()
}

func DeleteAllByVideoIdAndSeasonAndEpisode(videoId int64, season int64, episode int64) int64 {
	build := NewVideoResourcesQuery()
	build.kWheVideoId(videoId)
	build.kWheSeason(season)
	build.kWheEpisode(episode)
	return build.Delete()
}

func GetOneByVideoIdAndSeasonAndEpisode(videoId int64, season int64, episode int64) *VideoResources {
	build := NewVideoResourcesQuery()
	build.kWheVideoId(videoId)
	build.kWheSeason(season)
	build.kWheEpisode(episode)
	return build.GetOne()
}

func UpdateVideoResourcesByVideoIdAndSeasonAndEpisode(videoId int64, season int64, episode int64, p *VideoResources) int64 {
	build := NewVideoResourcesQuery()
	build.kWheVideoId(videoId)
	build.kWheSeason(season)
	build.kWheEpisode(episode)
	return build.update(p)
}

func FetchByVideoIdAndSeasonAndEpisode(videoId int64, season int64, episode int64, order ...string) VideoResourcesCollect {
	build := NewVideoResourcesQuery()
	build.kWheVideoId(videoId)
	build.kWheSeason(season)
	build.kWheEpisode(episode)
	build.OrderBy(parseOrderString(order))
	return build.Get()
}

func FetchByVideoIdAndSeasonAndEpisodeWithPageSize(videoId int64, season int64, episode int64, page int, pageSize int, order ...string) VideoResourcesCollect {
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * pageSize

	build := NewVideoResourcesQuery()
	build.kWheVideoId(videoId)
	build.kWheSeason(season)
	build.kWheEpisode(episode)
	build.OrderBy(parseOrderString(order))
	return build.Skip(offset).Limit(pageSize).Get()
}

func UpdateVideoResourcesByIds(id []int64, p *VideoResources) int64 {
	build := NewVideoResourcesQuery()

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

func UpdateVideoResourcesByIdsWhatEver(id []int64, p *VideoResources) int64 {
	build := NewVideoResourcesQuery()

	if len(id) == 1 {
		build.kWheId(id[0])
	} else {
		build.kWheIdIn(id)
	}

	return build.update(p)
}

func FetchByIdsWithPageSize(id []int64, page int, pageSize int, order ...string) VideoResourcesCollect {
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * pageSize

	build := NewVideoResourcesQuery()

	if len(id) == 0 {
		return make(VideoResourcesCollect, 0)
	}

	if len(id) == 1 {
		build.kWheId(id[0])
	} else {
		build.kWheIdIn(id)
	}
	build.OrderBy(parseOrderString(order))
	return build.Skip(offset).Limit(pageSize).Get()
}

func CheckExistById(id int64) bool {
	build := NewVideoResourcesQuery()
	build.kWheId(id)
	cnt := build.Count()
	return cnt > 0
}

func GetFirstById(id int64) *VideoResources {
	build := NewVideoResourcesQuery()
	build.kWheId(id)
	return build.First()
}

func DeleteById(id int64) int64 {
	build := NewVideoResourcesQuery()
	build.kWheId(id)
	return build.Delete()
}

func GetOneById(id int64) *VideoResources {
	build := NewVideoResourcesQuery()
	build.kWheId(id)
	return build.GetOne()
}

func UpdateVideoResourcesById(id int64, p *VideoResources) int64 {
	build := NewVideoResourcesQuery()
	build.kWheId(id)
	return build.update(p)
}
