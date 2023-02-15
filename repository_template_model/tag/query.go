package tag

import "github.com/livegoplayer/go_helper/utils"

func GetLastRangeId(first, end int64) int64 {
	build := NewTagQuery().kWheTagId(">", first)
	if end > 0 {
		build.kWheTagId("<", end)
	}

	res := build.OrderBy("tagId", "DESC").First()

	return res.TagId
}

func FetchByRange(first, end int64) TagCollect {
	query := NewTagQuery().kWheTagId(">", first)
	if end > 0 && end > first {
		query.kWheTagId("<", end)
	}

	return query.Get()
}

// FetchCateList 查出电影类型
func FetchCateList() TagCollect {
	return FetchByRange(100, 200)
}

func GetLastCateId() int64 {
	return GetLastRangeId(100, 200)
}

// InsertCate 使用之前请先去重
func InsertCate(tagName string) int64 {
	lastId := GetLastCateId()

	item := &Tag{
		TagId:   lastId + 1,
		TagName: tagName,
	}

	Save(item)
	return lastId + 1
}

// FetchCountryList 查出电影国家
func FetchCountryList() TagCollect {
	return FetchByRange(200, 300)
}

func getLastCountryId() int64 {
	return GetLastRangeId(200, 300)
}

// InsertCountry 使用之前请先去重
func InsertCountry(tagName string) int64 {
	lastId := getLastCountryId()
	item := &Tag{
		TagId:   lastId + 1,
		TagName: tagName,
	}
	Save(item)
	return lastId + 1
}

// FetchLanguageList 查出电影语言
func FetchLanguageList() TagCollect {
	return FetchByRange(300, 400)
}

func getLastLanguageId() int64 {
	return GetLastRangeId(300, 400)
}

// InsertLanguage 使用之前请先去重
func InsertLanguage(tagName string) int64 {
	lastId := getLastLanguageId()
	item := &Tag{
		TagId:   lastId + 1,
		TagName: tagName,
	}
	Save(item)
	return lastId + 1
}

// FetchDirectorList 查出电影导演
func FetchDirectorList() TagCollect {
	return FetchByRange(1000000000, 2000000000)
}

func getLastDirectorId() int64 {
	return GetLastRangeId(1000000000, 2000000000)
}

// InsertDirector 使用之前请先去重
func InsertDirector(tagName, tagUrl string) int64 {
	lastId := getLastDirectorId()
	item := &Tag{
		TagId:   lastId + 1,
		TagName: tagName,
		TagUrl:  tagUrl,
	}
	Save(item)
	return lastId + 1
}

// FetchScreenwriterList 查出电影编剧
func FetchScreenwriterList() TagCollect {
	return FetchByRange(2000000000, 3000000000)
}

func getLastScreenwriterId() int64 {
	return GetLastRangeId(2000000000, 3000000000)
}

// InsertScreenwriter 使用之前请先去重
func InsertScreenwriter(tagName, tag_url string) int64 {
	lastId := getLastScreenwriterId()
	item := &Tag{
		TagId:   lastId + 1,
		TagName: tagName,
		TagUrl:  tag_url,
	}
	Save(item)
	return lastId + 1
}

// FetchActorList 查出电影演员
func FetchActorList() TagCollect {
	return FetchByRange(3000000000, 0)
}

// FetchVideoCateList 查出video类型
func FetchVideoCateList() TagCollect {
	return FetchByRange(0, 100)
}

func getLastActorId() int64 {
	return GetLastRangeId(3000000000, 0)
}

// InsertActor 使用之前请先去重
func InsertActor(tagName, tagUrl string) int64 {
	lastId := getLastActorId()
	item := &Tag{
		TagId:   lastId + 1,
		TagName: tagName,
		TagUrl:  tagUrl,
	}
	Save(item)
	return lastId + 1
}

// FetchYearList 查出电影上映年份
func FetchYearList() TagCollect {
	return FetchByRange(1887, 3001)
}

func GetLastYearId() int64 {
	return GetLastRangeId(1887, 3001)
}

// InsertYear 使用之前请先去重
func InsertYear(tagName string) string {
	item := &Tag{
		TagId:   utils.AsInt64(tagName),
		TagName: tagName,
	}
	Save(item)
	return tagName
}

func FetchByTagName(tagName string) TagCollect {
	build := NewTagQuery()
	build.kWheTagId(">", 1000000000)
	build.kWheTagName(tagName)
	return build.Get()
}
