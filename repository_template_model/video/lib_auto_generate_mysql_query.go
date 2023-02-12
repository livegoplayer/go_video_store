package video

/* Unscoped *************************************** 满足常用的共有方法,根据mysql索引生成 **********************************************/
/**************************************** all 不是小表不要用 **************************************************/

func Save(f *Video) *Video {
	NewVideoQuery().save(f)
	return f
}

func UpdateVideoAll(p *Video) int64 {
	build := NewVideoQuery()
	return build.update(p)
}

func FetchVideoAll() VideoCollect {
	build := NewVideoQuery()
	return build.Get()
}

func FetchVideoAllWithPageSize(page int, pageSize int) VideoCollect {
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * pageSize

	build := NewVideoQuery()
	return build.Skip(offset).Limit(pageSize).Get()
}

func CountVideoAll() int64 {
	build := NewVideoQuery()
	return build.Count()
}

/**************************************** 根据 Index 生成的方法 **************************************************/

func CheckExistByPrettyName(prettyName string) bool {
	build := NewVideoQuery()
	build.kWhePrettyName(prettyName)
	cnt := build.Count()
	return cnt > 0
}

func GetFirstByPrettyName(prettyName string) *Video {
	build := NewVideoQuery()
	build.kWhePrettyName(prettyName)
	return build.First()
}

func DeleteAllByPrettyName(prettyName string) int64 {
	build := NewVideoQuery()
	build.kWhePrettyName(prettyName)
	return build.Delete()
}

func GetOneByPrettyName(prettyName string) *Video {
	build := NewVideoQuery()
	build.kWhePrettyName(prettyName)
	return build.GetOne()
}

func UpdateVideoByPrettyName(prettyName string, p *Video) int64 {
	build := NewVideoQuery()
	build.kWhePrettyName(prettyName)
	return build.update(p)
}

func FetchByPrettyName(prettyName string) VideoCollect {
	build := NewVideoQuery()
	build.kWhePrettyName(prettyName)
	return build.Get()
}

func FetchByPrettyNameWithPageSize(prettyName string, page int, pageSize int) VideoCollect {
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * pageSize

	build := NewVideoQuery()
	return build.Skip(offset).Limit(pageSize).Get()
}

func FetchByPrettyNames(prettyName []string) VideoCollect {
	build := NewVideoQuery()

	if len(prettyName) == 0 {
		return make(VideoCollect, 0)
	}

	if len(prettyName) == 1 {
		build.kWhePrettyName(prettyName[0])
	} else {
		build.kWhePrettyNameIn(prettyName)
	}

	return build.Get()
}

func UpdateVideoByPrettyNames(prettyName []string, p *Video) int64 {
	build := NewVideoQuery()

	if len(prettyName) == 0 {
		return 0
	}

	if len(prettyName) == 1 {
		build.kWhePrettyName(prettyName[0])
	} else {
		build.kWhePrettyNameIn(prettyName)
	}

	return build.update(p)
}

func FetchByPrettyNamesWithPageSize(prettyName []string, page int, pageSize int) VideoCollect {
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * pageSize

	build := NewVideoQuery()

	if len(prettyName) == 0 {
		return make(VideoCollect, 0)
	}

	if len(prettyName) == 1 {
		build.kWhePrettyName(prettyName[0])
	} else {
		build.kWhePrettyNameIn(prettyName)
	}

	return build.Skip(offset).Limit(pageSize).Get()
}

func UpdateVideoByPrettyNamesWhatEver(prettyName []string, p *Video) int64 {
	build := NewVideoQuery()

	if len(prettyName) == 1 {
		build.kWhePrettyName(prettyName[0])
	} else {
		build.kWhePrettyNameIn(prettyName)
	}

	return build.update(p)
}

func CountVideoByPrettyNames(prettyName []string) int64 {
	if len(prettyName) == 0 {
		return 0
	}
	build := NewVideoQuery()
	if len(prettyName) == 1 {
		build.kWhePrettyName(prettyName[0])
	} else {
		build.kWhePrettyNameIn(prettyName)
	}

	return build.Count()
}
func UpdateVideoByVideoIds(videoId []int64, p *Video) int64 {
	build := NewVideoQuery()

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

func UpdateVideoByVideoIdsWhatEver(videoId []int64, p *Video) int64 {
	build := NewVideoQuery()

	if len(videoId) == 1 {
		build.kWheVideoId(videoId[0])
	} else {
		build.kWheVideoIdIn(videoId)
	}

	return build.update(p)
}

func FetchByVideoIdsWithPageSize(videoId []int64, page int, pageSize int) VideoCollect {
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * pageSize

	build := NewVideoQuery()

	if len(videoId) == 0 {
		return make(VideoCollect, 0)
	}

	if len(videoId) == 1 {
		build.kWheVideoId(videoId[0])
	} else {
		build.kWheVideoIdIn(videoId)
	}

	return build.Skip(offset).Limit(pageSize).Get()
}

func CheckExistByVideoId(videoId int64) bool {
	build := NewVideoQuery()
	build.kWheVideoId(videoId)
	cnt := build.Count()
	return cnt > 0
}

func GetFirstByVideoId(videoId int64) *Video {
	build := NewVideoQuery()
	build.kWheVideoId(videoId)
	return build.First()
}

func DeleteByVideoId(videoId int64) int64 {
	build := NewVideoQuery()
	build.kWheVideoId(videoId)
	return build.Delete()
}

func GetOneByVideoId(videoId int64) *Video {
	build := NewVideoQuery()
	build.kWheVideoId(videoId)
	return build.GetOne()
}

func UpdateVideoByVideoId(videoId int64, p *Video) int64 {
	build := NewVideoQuery()
	build.kWheVideoId(videoId)
	return build.update(p)
}
