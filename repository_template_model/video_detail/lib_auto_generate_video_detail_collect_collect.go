package video_detail

import (
	"sort"
)

func (s VideoDetailCollect) Filter(f func(item VideoDetail) bool) VideoDetailCollect {
	m := make(VideoDetailCollect, 0)
	for _, v := range s {
		if f(v) {
			m = append(m, v)
		}
	}
	return m
}
func (s VideoDetailCollect) GroupByDescriptionUrl() map[string]VideoDetailCollect {
	m := make(map[string]VideoDetailCollect)
	for _, v := range s {
		if _, ok := m[v.DescriptionUrl]; !ok {
			m[v.DescriptionUrl] = make(VideoDetailCollect, 0)
		}
		m[v.DescriptionUrl] = append(m[v.DescriptionUrl], v)
	}
	return m
}

func (s VideoDetailCollect) GroupByTimeStr() map[string]VideoDetailCollect {
	m := make(map[string]VideoDetailCollect)
	for _, v := range s {
		if _, ok := m[v.TimeStr]; !ok {
			m[v.TimeStr] = make(VideoDetailCollect, 0)
		}
		m[v.TimeStr] = append(m[v.TimeStr], v)
	}
	return m
}

func (s VideoDetailCollect) GroupByUptDatetime() map[string]VideoDetailCollect {
	m := make(map[string]VideoDetailCollect)
	for _, v := range s {
		if _, ok := m[v.UptDatetime]; !ok {
			m[v.UptDatetime] = make(VideoDetailCollect, 0)
		}
		m[v.UptDatetime] = append(m[v.UptDatetime], v)
	}
	return m
}

func (s VideoDetailCollect) GroupBySetNum() map[int64]VideoDetailCollect {
	m := make(map[int64]VideoDetailCollect)
	for _, v := range s {
		if _, ok := m[v.SetNum]; !ok {
			m[v.SetNum] = make(VideoDetailCollect, 0)
		}
		m[v.SetNum] = append(m[v.SetNum], v)
	}
	return m
}

func (s VideoDetailCollect) GroupByVideoId() map[int64]VideoDetailCollect {
	m := make(map[int64]VideoDetailCollect)
	for _, v := range s {
		if _, ok := m[v.VideoId]; !ok {
			m[v.VideoId] = make(VideoDetailCollect, 0)
		}
		m[v.VideoId] = append(m[v.VideoId], v)
	}
	return m
}

func (s VideoDetailCollect) GroupByPosterCache() map[string]VideoDetailCollect {
	m := make(map[string]VideoDetailCollect)
	for _, v := range s {
		if _, ok := m[v.PosterCache]; !ok {
			m[v.PosterCache] = make(VideoDetailCollect, 0)
		}
		m[v.PosterCache] = append(m[v.PosterCache], v)
	}
	return m
}

func (s VideoDetailCollect) GroupByOfficialWebsite() map[string]VideoDetailCollect {
	m := make(map[string]VideoDetailCollect)
	for _, v := range s {
		if _, ok := m[v.OfficialWebsite]; !ok {
			m[v.OfficialWebsite] = make(VideoDetailCollect, 0)
		}
		m[v.OfficialWebsite] = append(m[v.OfficialWebsite], v)
	}
	return m
}

func (s VideoDetailCollect) GroupByReleasedDateStr() map[string]VideoDetailCollect {
	m := make(map[string]VideoDetailCollect)
	for _, v := range s {
		if _, ok := m[v.ReleasedDateStr]; !ok {
			m[v.ReleasedDateStr] = make(VideoDetailCollect, 0)
		}
		m[v.ReleasedDateStr] = append(m[v.ReleasedDateStr], v)
	}
	return m
}

func (s VideoDetailCollect) GroupByScore() map[float64]VideoDetailCollect {
	m := make(map[float64]VideoDetailCollect)
	for _, v := range s {
		if _, ok := m[v.Score]; !ok {
			m[v.Score] = make(VideoDetailCollect, 0)
		}
		m[v.Score] = append(m[v.Score], v)
	}
	return m
}

func (s VideoDetailCollect) GroupByYear() map[int64]VideoDetailCollect {
	m := make(map[int64]VideoDetailCollect)
	for _, v := range s {
		if _, ok := m[v.Year]; !ok {
			m[v.Year] = make(VideoDetailCollect, 0)
		}
		m[v.Year] = append(m[v.Year], v)
	}
	return m
}

func (s VideoDetailCollect) GroupByPosterListUrl() map[string]VideoDetailCollect {
	m := make(map[string]VideoDetailCollect)
	for _, v := range s {
		if _, ok := m[v.PosterListUrl]; !ok {
			m[v.PosterListUrl] = make(VideoDetailCollect, 0)
		}
		m[v.PosterListUrl] = append(m[v.PosterListUrl], v)
	}
	return m
}

func (s VideoDetailCollect) GroupByVideoDescription() map[string]VideoDetailCollect {
	m := make(map[string]VideoDetailCollect)
	for _, v := range s {
		if _, ok := m[v.VideoDescription]; !ok {
			m[v.VideoDescription] = make(VideoDetailCollect, 0)
		}
		m[v.VideoDescription] = append(m[v.VideoDescription], v)
	}
	return m
}

func (s VideoDetailCollect) GroupByVideoPosterUrl() map[string]VideoDetailCollect {
	m := make(map[string]VideoDetailCollect)
	for _, v := range s {
		if _, ok := m[v.VideoPosterUrl]; !ok {
			m[v.VideoPosterUrl] = make(VideoDetailCollect, 0)
		}
		m[v.VideoPosterUrl] = append(m[v.VideoPosterUrl], v)
	}
	return m
}

func (s VideoDetailCollect) GroupById() map[int64]VideoDetailCollect {
	m := make(map[int64]VideoDetailCollect)
	for _, v := range s {
		if _, ok := m[v.Id]; !ok {
			m[v.Id] = make(VideoDetailCollect, 0)
		}
		m[v.Id] = append(m[v.Id], v)
	}
	return m
}

func (s VideoDetailCollect) GroupBySeason() map[int64]VideoDetailCollect {
	m := make(map[int64]VideoDetailCollect)
	for _, v := range s {
		if _, ok := m[v.Season]; !ok {
			m[v.Season] = make(VideoDetailCollect, 0)
		}
		m[v.Season] = append(m[v.Season], v)
	}
	return m
}

func (s VideoDetailCollect) GroupByPerUpdateTime() map[string]VideoDetailCollect {
	m := make(map[string]VideoDetailCollect)
	for _, v := range s {
		if _, ok := m[v.PerUpdateTime]; !ok {
			m[v.PerUpdateTime] = make(VideoDetailCollect, 0)
		}
		m[v.PerUpdateTime] = append(m[v.PerUpdateTime], v)
	}
	return m
}

func (s VideoDetailCollect) GroupByAddDatetime() map[string]VideoDetailCollect {
	m := make(map[string]VideoDetailCollect)
	for _, v := range s {
		if _, ok := m[v.AddDatetime]; !ok {
			m[v.AddDatetime] = make(VideoDetailCollect, 0)
		}
		m[v.AddDatetime] = append(m[v.AddDatetime], v)
	}
	return m
}

func (s VideoDetailCollect) KeyByDescriptionUrl() map[string]VideoDetail {
	m := make(map[string]VideoDetail)
	for _, v := range s {
		m[v.DescriptionUrl] = v
	}
	return m
}
func (s VideoDetailCollect) KeyByTimeStr() map[string]VideoDetail {
	m := make(map[string]VideoDetail)
	for _, v := range s {
		m[v.TimeStr] = v
	}
	return m
}
func (s VideoDetailCollect) KeyByUptDatetime() map[string]VideoDetail {
	m := make(map[string]VideoDetail)
	for _, v := range s {
		m[v.UptDatetime] = v
	}
	return m
}
func (s VideoDetailCollect) KeyBySetNum() map[int64]VideoDetail {
	m := make(map[int64]VideoDetail)
	for _, v := range s {
		m[v.SetNum] = v
	}
	return m
}
func (s VideoDetailCollect) KeyByVideoId() map[int64]VideoDetail {
	m := make(map[int64]VideoDetail)
	for _, v := range s {
		m[v.VideoId] = v
	}
	return m
}
func (s VideoDetailCollect) KeyByPosterCache() map[string]VideoDetail {
	m := make(map[string]VideoDetail)
	for _, v := range s {
		m[v.PosterCache] = v
	}
	return m
}
func (s VideoDetailCollect) KeyByOfficialWebsite() map[string]VideoDetail {
	m := make(map[string]VideoDetail)
	for _, v := range s {
		m[v.OfficialWebsite] = v
	}
	return m
}
func (s VideoDetailCollect) KeyByReleasedDateStr() map[string]VideoDetail {
	m := make(map[string]VideoDetail)
	for _, v := range s {
		m[v.ReleasedDateStr] = v
	}
	return m
}
func (s VideoDetailCollect) KeyByScore() map[float64]VideoDetail {
	m := make(map[float64]VideoDetail)
	for _, v := range s {
		m[v.Score] = v
	}
	return m
}
func (s VideoDetailCollect) KeyByYear() map[int64]VideoDetail {
	m := make(map[int64]VideoDetail)
	for _, v := range s {
		m[v.Year] = v
	}
	return m
}
func (s VideoDetailCollect) KeyByPosterListUrl() map[string]VideoDetail {
	m := make(map[string]VideoDetail)
	for _, v := range s {
		m[v.PosterListUrl] = v
	}
	return m
}
func (s VideoDetailCollect) KeyByVideoDescription() map[string]VideoDetail {
	m := make(map[string]VideoDetail)
	for _, v := range s {
		m[v.VideoDescription] = v
	}
	return m
}
func (s VideoDetailCollect) KeyByVideoPosterUrl() map[string]VideoDetail {
	m := make(map[string]VideoDetail)
	for _, v := range s {
		m[v.VideoPosterUrl] = v
	}
	return m
}
func (s VideoDetailCollect) KeyById() map[int64]VideoDetail {
	m := make(map[int64]VideoDetail)
	for _, v := range s {
		m[v.Id] = v
	}
	return m
}
func (s VideoDetailCollect) KeyBySeason() map[int64]VideoDetail {
	m := make(map[int64]VideoDetail)
	for _, v := range s {
		m[v.Season] = v
	}
	return m
}
func (s VideoDetailCollect) KeyByPerUpdateTime() map[string]VideoDetail {
	m := make(map[string]VideoDetail)
	for _, v := range s {
		m[v.PerUpdateTime] = v
	}
	return m
}
func (s VideoDetailCollect) KeyByAddDatetime() map[string]VideoDetail {
	m := make(map[string]VideoDetail)
	for _, v := range s {
		m[v.AddDatetime] = v
	}
	return m
}
func (s VideoDetailCollect) PluckDescriptionUrl() []string {
	list := make([]string, len(s))
	for i, v := range s {
		list[i] = v.DescriptionUrl
	}
	return list
}
func (s VideoDetailCollect) PluckTimeStr() []string {
	list := make([]string, len(s))
	for i, v := range s {
		list[i] = v.TimeStr
	}
	return list
}
func (s VideoDetailCollect) PluckUptDatetime() []string {
	list := make([]string, len(s))
	for i, v := range s {
		list[i] = v.UptDatetime
	}
	return list
}
func (s VideoDetailCollect) PluckSetNum() []int64 {
	list := make([]int64, len(s))
	for i, v := range s {
		list[i] = v.SetNum
	}
	return list
}
func (s VideoDetailCollect) PluckVideoId() []int64 {
	list := make([]int64, len(s))
	for i, v := range s {
		list[i] = v.VideoId
	}
	return list
}
func (s VideoDetailCollect) PluckPosterCache() []string {
	list := make([]string, len(s))
	for i, v := range s {
		list[i] = v.PosterCache
	}
	return list
}
func (s VideoDetailCollect) PluckOfficialWebsite() []string {
	list := make([]string, len(s))
	for i, v := range s {
		list[i] = v.OfficialWebsite
	}
	return list
}
func (s VideoDetailCollect) PluckReleasedDateStr() []string {
	list := make([]string, len(s))
	for i, v := range s {
		list[i] = v.ReleasedDateStr
	}
	return list
}
func (s VideoDetailCollect) PluckScore() []float64 {
	list := make([]float64, len(s))
	for i, v := range s {
		list[i] = v.Score
	}
	return list
}
func (s VideoDetailCollect) PluckYear() []int64 {
	list := make([]int64, len(s))
	for i, v := range s {
		list[i] = v.Year
	}
	return list
}
func (s VideoDetailCollect) PluckPosterListUrl() []string {
	list := make([]string, len(s))
	for i, v := range s {
		list[i] = v.PosterListUrl
	}
	return list
}
func (s VideoDetailCollect) PluckVideoDescription() []string {
	list := make([]string, len(s))
	for i, v := range s {
		list[i] = v.VideoDescription
	}
	return list
}
func (s VideoDetailCollect) PluckVideoPosterUrl() []string {
	list := make([]string, len(s))
	for i, v := range s {
		list[i] = v.VideoPosterUrl
	}
	return list
}
func (s VideoDetailCollect) PluckId() []int64 {
	list := make([]int64, len(s))
	for i, v := range s {
		list[i] = v.Id
	}
	return list
}
func (s VideoDetailCollect) PluckSeason() []int64 {
	list := make([]int64, len(s))
	for i, v := range s {
		list[i] = v.Season
	}
	return list
}
func (s VideoDetailCollect) PluckPerUpdateTime() []string {
	list := make([]string, len(s))
	for i, v := range s {
		list[i] = v.PerUpdateTime
	}
	return list
}
func (s VideoDetailCollect) PluckAddDatetime() []string {
	list := make([]string, len(s))
	for i, v := range s {
		list[i] = v.AddDatetime
	}
	return list
}

func (s VideoDetailCollect) PluckUniDescriptionUrl() []string {
	uniMap := make(map[string]bool)
	list := make([]string, 0)
	for _, v := range s {
		_, ok := uniMap[v.DescriptionUrl]
		if !ok {
			uniMap[v.DescriptionUrl] = true
			list = append(list, v.DescriptionUrl)
		}
	}
	return list
}

func (s VideoDetailCollect) PluckUniTimeStr() []string {
	uniMap := make(map[string]bool)
	list := make([]string, 0)
	for _, v := range s {
		_, ok := uniMap[v.TimeStr]
		if !ok {
			uniMap[v.TimeStr] = true
			list = append(list, v.TimeStr)
		}
	}
	return list
}

func (s VideoDetailCollect) PluckUniUptDatetime() []string {
	uniMap := make(map[string]bool)
	list := make([]string, 0)
	for _, v := range s {
		_, ok := uniMap[v.UptDatetime]
		if !ok {
			uniMap[v.UptDatetime] = true
			list = append(list, v.UptDatetime)
		}
	}
	return list
}

func (s VideoDetailCollect) PluckUniSetNum() []int64 {
	uniMap := make(map[int64]bool)
	list := make([]int64, 0)
	for _, v := range s {
		_, ok := uniMap[v.SetNum]
		if !ok {
			uniMap[v.SetNum] = true
			list = append(list, v.SetNum)
		}
	}
	return list
}

func (s VideoDetailCollect) PluckUniVideoId() []int64 {
	uniMap := make(map[int64]bool)
	list := make([]int64, 0)
	for _, v := range s {
		_, ok := uniMap[v.VideoId]
		if !ok {
			uniMap[v.VideoId] = true
			list = append(list, v.VideoId)
		}
	}
	return list
}

func (s VideoDetailCollect) PluckUniPosterCache() []string {
	uniMap := make(map[string]bool)
	list := make([]string, 0)
	for _, v := range s {
		_, ok := uniMap[v.PosterCache]
		if !ok {
			uniMap[v.PosterCache] = true
			list = append(list, v.PosterCache)
		}
	}
	return list
}

func (s VideoDetailCollect) PluckUniOfficialWebsite() []string {
	uniMap := make(map[string]bool)
	list := make([]string, 0)
	for _, v := range s {
		_, ok := uniMap[v.OfficialWebsite]
		if !ok {
			uniMap[v.OfficialWebsite] = true
			list = append(list, v.OfficialWebsite)
		}
	}
	return list
}

func (s VideoDetailCollect) PluckUniReleasedDateStr() []string {
	uniMap := make(map[string]bool)
	list := make([]string, 0)
	for _, v := range s {
		_, ok := uniMap[v.ReleasedDateStr]
		if !ok {
			uniMap[v.ReleasedDateStr] = true
			list = append(list, v.ReleasedDateStr)
		}
	}
	return list
}

func (s VideoDetailCollect) PluckUniScore() []float64 {
	uniMap := make(map[float64]bool)
	list := make([]float64, 0)
	for _, v := range s {
		_, ok := uniMap[v.Score]
		if !ok {
			uniMap[v.Score] = true
			list = append(list, v.Score)
		}
	}
	return list
}

func (s VideoDetailCollect) PluckUniYear() []int64 {
	uniMap := make(map[int64]bool)
	list := make([]int64, 0)
	for _, v := range s {
		_, ok := uniMap[v.Year]
		if !ok {
			uniMap[v.Year] = true
			list = append(list, v.Year)
		}
	}
	return list
}

func (s VideoDetailCollect) PluckUniPosterListUrl() []string {
	uniMap := make(map[string]bool)
	list := make([]string, 0)
	for _, v := range s {
		_, ok := uniMap[v.PosterListUrl]
		if !ok {
			uniMap[v.PosterListUrl] = true
			list = append(list, v.PosterListUrl)
		}
	}
	return list
}

func (s VideoDetailCollect) PluckUniVideoDescription() []string {
	uniMap := make(map[string]bool)
	list := make([]string, 0)
	for _, v := range s {
		_, ok := uniMap[v.VideoDescription]
		if !ok {
			uniMap[v.VideoDescription] = true
			list = append(list, v.VideoDescription)
		}
	}
	return list
}

func (s VideoDetailCollect) PluckUniVideoPosterUrl() []string {
	uniMap := make(map[string]bool)
	list := make([]string, 0)
	for _, v := range s {
		_, ok := uniMap[v.VideoPosterUrl]
		if !ok {
			uniMap[v.VideoPosterUrl] = true
			list = append(list, v.VideoPosterUrl)
		}
	}
	return list
}

func (s VideoDetailCollect) PluckUniId() []int64 {
	uniMap := make(map[int64]bool)
	list := make([]int64, 0)
	for _, v := range s {
		_, ok := uniMap[v.Id]
		if !ok {
			uniMap[v.Id] = true
			list = append(list, v.Id)
		}
	}
	return list
}

func (s VideoDetailCollect) PluckUniSeason() []int64 {
	uniMap := make(map[int64]bool)
	list := make([]int64, 0)
	for _, v := range s {
		_, ok := uniMap[v.Season]
		if !ok {
			uniMap[v.Season] = true
			list = append(list, v.Season)
		}
	}
	return list
}

func (s VideoDetailCollect) PluckUniPerUpdateTime() []string {
	uniMap := make(map[string]bool)
	list := make([]string, 0)
	for _, v := range s {
		_, ok := uniMap[v.PerUpdateTime]
		if !ok {
			uniMap[v.PerUpdateTime] = true
			list = append(list, v.PerUpdateTime)
		}
	}
	return list
}

func (s VideoDetailCollect) PluckUniAddDatetime() []string {
	uniMap := make(map[string]bool)
	list := make([]string, 0)
	for _, v := range s {
		_, ok := uniMap[v.AddDatetime]
		if !ok {
			uniMap[v.AddDatetime] = true
			list = append(list, v.AddDatetime)
		}
	}
	return list
}

func (s VideoDetailCollect) SortByFunc(f func(i, j int) bool) VideoDetailCollect {
	sort.SliceStable(s, f)
	return s
}
