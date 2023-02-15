package video_tag

import (
	"github.com/livegoplayer/go_video_store/repository_template_model/tag"
)

func GetAllTags() map[string]interface{} {
	actorList := tag.FetchActorList().KeyByTagId()
	cateList := tag.FetchCateList().KeyByTagId()
	countryList := tag.FetchCountryList().KeyByTagId()
	yearList := tag.FetchYearList().KeyByTagId()
	screenwriterList := tag.FetchScreenwriterList().KeyByTagId()
	directorList := tag.FetchDirectorList().KeyByTagId()
	languageList := tag.FetchLanguageList().KeyByTagId()
	videoCateList := tag.FetchVideoCateList().KeyByTagId()
	m := map[string]interface{}{
		"action_list":       actorList,
		"video_cate_list":   videoCateList,
		"cate_list":         cateList,
		"country_list":      countryList,
		"year_list":         yearList,
		"screenwriter_list": screenwriterList,
		"director_list":     directorList,
		"language_list":     languageList,
	}
	return m
}
