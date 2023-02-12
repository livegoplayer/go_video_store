package video_resources

import (
	"sort"
)

func (s VideoResourcesCollect) Filter(f func(item VideoResources) bool) VideoResourcesCollect {
	m := make(VideoResourcesCollect, 0)
	for _, v := range s {
		if f(v) {
			m = append(m, v)
		}
	}
	return m
}
func (s VideoResourcesCollect) GroupByVideoRefer() map[int64]VideoResourcesCollect {
	m := make(map[int64]VideoResourcesCollect)
	for _, v := range s {
		if _, ok := m[v.VideoRefer]; !ok {
			m[v.VideoRefer] = make(VideoResourcesCollect, 0)
		}
		m[v.VideoRefer] = append(m[v.VideoRefer], v)
	}
	return m
}

func (s VideoResourcesCollect) GroupByEpisode() map[int64]VideoResourcesCollect {
	m := make(map[int64]VideoResourcesCollect)
	for _, v := range s {
		if _, ok := m[v.Episode]; !ok {
			m[v.Episode] = make(VideoResourcesCollect, 0)
		}
		m[v.Episode] = append(m[v.Episode], v)
	}
	return m
}

func (s VideoResourcesCollect) GroupByUptDatetime() map[string]VideoResourcesCollect {
	m := make(map[string]VideoResourcesCollect)
	for _, v := range s {
		if _, ok := m[v.UptDatetime]; !ok {
			m[v.UptDatetime] = make(VideoResourcesCollect, 0)
		}
		m[v.UptDatetime] = append(m[v.UptDatetime], v)
	}
	return m
}

func (s VideoResourcesCollect) GroupByVideoId() map[int64]VideoResourcesCollect {
	m := make(map[int64]VideoResourcesCollect)
	for _, v := range s {
		if _, ok := m[v.VideoId]; !ok {
			m[v.VideoId] = make(VideoResourcesCollect, 0)
		}
		m[v.VideoId] = append(m[v.VideoId], v)
	}
	return m
}

func (s VideoResourcesCollect) GroupByLastUpdateTime() map[int64]VideoResourcesCollect {
	m := make(map[int64]VideoResourcesCollect)
	for _, v := range s {
		if _, ok := m[v.LastUpdateTime]; !ok {
			m[v.LastUpdateTime] = make(VideoResourcesCollect, 0)
		}
		m[v.LastUpdateTime] = append(m[v.LastUpdateTime], v)
	}
	return m
}

func (s VideoResourcesCollect) GroupByVideoReferDomain() map[string]VideoResourcesCollect {
	m := make(map[string]VideoResourcesCollect)
	for _, v := range s {
		if _, ok := m[v.VideoReferDomain]; !ok {
			m[v.VideoReferDomain] = make(VideoResourcesCollect, 0)
		}
		m[v.VideoReferDomain] = append(m[v.VideoReferDomain], v)
	}
	return m
}

func (s VideoResourcesCollect) GroupByData() map[string]VideoResourcesCollect {
	m := make(map[string]VideoResourcesCollect)
	for _, v := range s {
		if _, ok := m[v.Data]; !ok {
			m[v.Data] = make(VideoResourcesCollect, 0)
		}
		m[v.Data] = append(m[v.Data], v)
	}
	return m
}

func (s VideoResourcesCollect) GroupById() map[int64]VideoResourcesCollect {
	m := make(map[int64]VideoResourcesCollect)
	for _, v := range s {
		if _, ok := m[v.Id]; !ok {
			m[v.Id] = make(VideoResourcesCollect, 0)
		}
		m[v.Id] = append(m[v.Id], v)
	}
	return m
}

func (s VideoResourcesCollect) GroupBySeason() map[int64]VideoResourcesCollect {
	m := make(map[int64]VideoResourcesCollect)
	for _, v := range s {
		if _, ok := m[v.Season]; !ok {
			m[v.Season] = make(VideoResourcesCollect, 0)
		}
		m[v.Season] = append(m[v.Season], v)
	}
	return m
}

func (s VideoResourcesCollect) GroupByAddDatetime() map[string]VideoResourcesCollect {
	m := make(map[string]VideoResourcesCollect)
	for _, v := range s {
		if _, ok := m[v.AddDatetime]; !ok {
			m[v.AddDatetime] = make(VideoResourcesCollect, 0)
		}
		m[v.AddDatetime] = append(m[v.AddDatetime], v)
	}
	return m
}

func (s VideoResourcesCollect) GroupByVideoUrl() map[string]VideoResourcesCollect {
	m := make(map[string]VideoResourcesCollect)
	for _, v := range s {
		if _, ok := m[v.VideoUrl]; !ok {
			m[v.VideoUrl] = make(VideoResourcesCollect, 0)
		}
		m[v.VideoUrl] = append(m[v.VideoUrl], v)
	}
	return m
}

func (s VideoResourcesCollect) GroupByName() map[string]VideoResourcesCollect {
	m := make(map[string]VideoResourcesCollect)
	for _, v := range s {
		if _, ok := m[v.Name]; !ok {
			m[v.Name] = make(VideoResourcesCollect, 0)
		}
		m[v.Name] = append(m[v.Name], v)
	}
	return m
}

func (s VideoResourcesCollect) KeyByVideoRefer() map[int64]VideoResources {
	m := make(map[int64]VideoResources)
	for _, v := range s {
		m[v.VideoRefer] = v
	}
	return m
}
func (s VideoResourcesCollect) KeyByEpisode() map[int64]VideoResources {
	m := make(map[int64]VideoResources)
	for _, v := range s {
		m[v.Episode] = v
	}
	return m
}
func (s VideoResourcesCollect) KeyByUptDatetime() map[string]VideoResources {
	m := make(map[string]VideoResources)
	for _, v := range s {
		m[v.UptDatetime] = v
	}
	return m
}
func (s VideoResourcesCollect) KeyByVideoId() map[int64]VideoResources {
	m := make(map[int64]VideoResources)
	for _, v := range s {
		m[v.VideoId] = v
	}
	return m
}
func (s VideoResourcesCollect) KeyByLastUpdateTime() map[int64]VideoResources {
	m := make(map[int64]VideoResources)
	for _, v := range s {
		m[v.LastUpdateTime] = v
	}
	return m
}
func (s VideoResourcesCollect) KeyByVideoReferDomain() map[string]VideoResources {
	m := make(map[string]VideoResources)
	for _, v := range s {
		m[v.VideoReferDomain] = v
	}
	return m
}
func (s VideoResourcesCollect) KeyByData() map[string]VideoResources {
	m := make(map[string]VideoResources)
	for _, v := range s {
		m[v.Data] = v
	}
	return m
}
func (s VideoResourcesCollect) KeyById() map[int64]VideoResources {
	m := make(map[int64]VideoResources)
	for _, v := range s {
		m[v.Id] = v
	}
	return m
}
func (s VideoResourcesCollect) KeyBySeason() map[int64]VideoResources {
	m := make(map[int64]VideoResources)
	for _, v := range s {
		m[v.Season] = v
	}
	return m
}
func (s VideoResourcesCollect) KeyByAddDatetime() map[string]VideoResources {
	m := make(map[string]VideoResources)
	for _, v := range s {
		m[v.AddDatetime] = v
	}
	return m
}
func (s VideoResourcesCollect) KeyByVideoUrl() map[string]VideoResources {
	m := make(map[string]VideoResources)
	for _, v := range s {
		m[v.VideoUrl] = v
	}
	return m
}
func (s VideoResourcesCollect) KeyByName() map[string]VideoResources {
	m := make(map[string]VideoResources)
	for _, v := range s {
		m[v.Name] = v
	}
	return m
}
func (s VideoResourcesCollect) PluckVideoRefer() []int64 {
	list := make([]int64, len(s))
	for i, v := range s {
		list[i] = v.VideoRefer
	}
	return list
}
func (s VideoResourcesCollect) PluckEpisode() []int64 {
	list := make([]int64, len(s))
	for i, v := range s {
		list[i] = v.Episode
	}
	return list
}
func (s VideoResourcesCollect) PluckUptDatetime() []string {
	list := make([]string, len(s))
	for i, v := range s {
		list[i] = v.UptDatetime
	}
	return list
}
func (s VideoResourcesCollect) PluckVideoId() []int64 {
	list := make([]int64, len(s))
	for i, v := range s {
		list[i] = v.VideoId
	}
	return list
}
func (s VideoResourcesCollect) PluckLastUpdateTime() []int64 {
	list := make([]int64, len(s))
	for i, v := range s {
		list[i] = v.LastUpdateTime
	}
	return list
}
func (s VideoResourcesCollect) PluckVideoReferDomain() []string {
	list := make([]string, len(s))
	for i, v := range s {
		list[i] = v.VideoReferDomain
	}
	return list
}
func (s VideoResourcesCollect) PluckData() []string {
	list := make([]string, len(s))
	for i, v := range s {
		list[i] = v.Data
	}
	return list
}
func (s VideoResourcesCollect) PluckId() []int64 {
	list := make([]int64, len(s))
	for i, v := range s {
		list[i] = v.Id
	}
	return list
}
func (s VideoResourcesCollect) PluckSeason() []int64 {
	list := make([]int64, len(s))
	for i, v := range s {
		list[i] = v.Season
	}
	return list
}
func (s VideoResourcesCollect) PluckAddDatetime() []string {
	list := make([]string, len(s))
	for i, v := range s {
		list[i] = v.AddDatetime
	}
	return list
}
func (s VideoResourcesCollect) PluckVideoUrl() []string {
	list := make([]string, len(s))
	for i, v := range s {
		list[i] = v.VideoUrl
	}
	return list
}
func (s VideoResourcesCollect) PluckName() []string {
	list := make([]string, len(s))
	for i, v := range s {
		list[i] = v.Name
	}
	return list
}

func (s VideoResourcesCollect) PluckUniVideoRefer() []int64 {
	uniMap := make(map[int64]bool)
	list := make([]int64, 0)
	for _, v := range s {
		_, ok := uniMap[v.VideoRefer]
		if !ok {
			uniMap[v.VideoRefer] = true
			list = append(list, v.VideoRefer)
		}
	}
	return list
}

func (s VideoResourcesCollect) PluckUniEpisode() []int64 {
	uniMap := make(map[int64]bool)
	list := make([]int64, 0)
	for _, v := range s {
		_, ok := uniMap[v.Episode]
		if !ok {
			uniMap[v.Episode] = true
			list = append(list, v.Episode)
		}
	}
	return list
}

func (s VideoResourcesCollect) PluckUniUptDatetime() []string {
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

func (s VideoResourcesCollect) PluckUniVideoId() []int64 {
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

func (s VideoResourcesCollect) PluckUniLastUpdateTime() []int64 {
	uniMap := make(map[int64]bool)
	list := make([]int64, 0)
	for _, v := range s {
		_, ok := uniMap[v.LastUpdateTime]
		if !ok {
			uniMap[v.LastUpdateTime] = true
			list = append(list, v.LastUpdateTime)
		}
	}
	return list
}

func (s VideoResourcesCollect) PluckUniVideoReferDomain() []string {
	uniMap := make(map[string]bool)
	list := make([]string, 0)
	for _, v := range s {
		_, ok := uniMap[v.VideoReferDomain]
		if !ok {
			uniMap[v.VideoReferDomain] = true
			list = append(list, v.VideoReferDomain)
		}
	}
	return list
}

func (s VideoResourcesCollect) PluckUniData() []string {
	uniMap := make(map[string]bool)
	list := make([]string, 0)
	for _, v := range s {
		_, ok := uniMap[v.Data]
		if !ok {
			uniMap[v.Data] = true
			list = append(list, v.Data)
		}
	}
	return list
}

func (s VideoResourcesCollect) PluckUniId() []int64 {
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

func (s VideoResourcesCollect) PluckUniSeason() []int64 {
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

func (s VideoResourcesCollect) PluckUniAddDatetime() []string {
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

func (s VideoResourcesCollect) PluckUniVideoUrl() []string {
	uniMap := make(map[string]bool)
	list := make([]string, 0)
	for _, v := range s {
		_, ok := uniMap[v.VideoUrl]
		if !ok {
			uniMap[v.VideoUrl] = true
			list = append(list, v.VideoUrl)
		}
	}
	return list
}

func (s VideoResourcesCollect) PluckUniName() []string {
	uniMap := make(map[string]bool)
	list := make([]string, 0)
	for _, v := range s {
		_, ok := uniMap[v.Name]
		if !ok {
			uniMap[v.Name] = true
			list = append(list, v.Name)
		}
	}
	return list
}

func (s VideoResourcesCollect) SortByFunc(f func(i, j int) bool) VideoResourcesCollect {
	sort.SliceStable(s, f)
	return s
}
