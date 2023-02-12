package video_tag

import (
	"sort"
)

func (s VideoTagCollect) Filter(f func(item VideoTag) bool) VideoTagCollect {
	m := make(VideoTagCollect, 0)
	for _, v := range s {
		if f(v) {
			m = append(m, v)
		}
	}
	return m
}
func (s VideoTagCollect) GroupByUptDatetime() map[string]VideoTagCollect {
	m := make(map[string]VideoTagCollect)
	for _, v := range s {
		if _, ok := m[v.UptDatetime]; !ok {
			m[v.UptDatetime] = make(VideoTagCollect, 0)
		}
		m[v.UptDatetime] = append(m[v.UptDatetime], v)
	}
	return m
}

func (s VideoTagCollect) GroupByVideoId() map[int64]VideoTagCollect {
	m := make(map[int64]VideoTagCollect)
	for _, v := range s {
		if _, ok := m[v.VideoId]; !ok {
			m[v.VideoId] = make(VideoTagCollect, 0)
		}
		m[v.VideoId] = append(m[v.VideoId], v)
	}
	return m
}

func (s VideoTagCollect) GroupByTagId() map[int64]VideoTagCollect {
	m := make(map[int64]VideoTagCollect)
	for _, v := range s {
		if _, ok := m[v.TagId]; !ok {
			m[v.TagId] = make(VideoTagCollect, 0)
		}
		m[v.TagId] = append(m[v.TagId], v)
	}
	return m
}

func (s VideoTagCollect) GroupById() map[int64]VideoTagCollect {
	m := make(map[int64]VideoTagCollect)
	for _, v := range s {
		if _, ok := m[v.Id]; !ok {
			m[v.Id] = make(VideoTagCollect, 0)
		}
		m[v.Id] = append(m[v.Id], v)
	}
	return m
}

func (s VideoTagCollect) GroupByAddDatetime() map[string]VideoTagCollect {
	m := make(map[string]VideoTagCollect)
	for _, v := range s {
		if _, ok := m[v.AddDatetime]; !ok {
			m[v.AddDatetime] = make(VideoTagCollect, 0)
		}
		m[v.AddDatetime] = append(m[v.AddDatetime], v)
	}
	return m
}

func (s VideoTagCollect) KeyByUptDatetime() map[string]VideoTag {
	m := make(map[string]VideoTag)
	for _, v := range s {
		m[v.UptDatetime] = v
	}
	return m
}
func (s VideoTagCollect) KeyByVideoId() map[int64]VideoTag {
	m := make(map[int64]VideoTag)
	for _, v := range s {
		m[v.VideoId] = v
	}
	return m
}
func (s VideoTagCollect) KeyByTagId() map[int64]VideoTag {
	m := make(map[int64]VideoTag)
	for _, v := range s {
		m[v.TagId] = v
	}
	return m
}
func (s VideoTagCollect) KeyById() map[int64]VideoTag {
	m := make(map[int64]VideoTag)
	for _, v := range s {
		m[v.Id] = v
	}
	return m
}
func (s VideoTagCollect) KeyByAddDatetime() map[string]VideoTag {
	m := make(map[string]VideoTag)
	for _, v := range s {
		m[v.AddDatetime] = v
	}
	return m
}
func (s VideoTagCollect) PluckUptDatetime() []string {
	list := make([]string, len(s))
	for i, v := range s {
		list[i] = v.UptDatetime
	}
	return list
}
func (s VideoTagCollect) PluckVideoId() []int64 {
	list := make([]int64, len(s))
	for i, v := range s {
		list[i] = v.VideoId
	}
	return list
}
func (s VideoTagCollect) PluckTagId() []int64 {
	list := make([]int64, len(s))
	for i, v := range s {
		list[i] = v.TagId
	}
	return list
}
func (s VideoTagCollect) PluckId() []int64 {
	list := make([]int64, len(s))
	for i, v := range s {
		list[i] = v.Id
	}
	return list
}
func (s VideoTagCollect) PluckAddDatetime() []string {
	list := make([]string, len(s))
	for i, v := range s {
		list[i] = v.AddDatetime
	}
	return list
}

func (s VideoTagCollect) PluckUniUptDatetime() []string {
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

func (s VideoTagCollect) PluckUniVideoId() []int64 {
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

func (s VideoTagCollect) PluckUniTagId() []int64 {
	uniMap := make(map[int64]bool)
	list := make([]int64, 0)
	for _, v := range s {
		_, ok := uniMap[v.TagId]
		if !ok {
			uniMap[v.TagId] = true
			list = append(list, v.TagId)
		}
	}
	return list
}

func (s VideoTagCollect) PluckUniId() []int64 {
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

func (s VideoTagCollect) PluckUniAddDatetime() []string {
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

func (s VideoTagCollect) SortByFunc(f func(i, j int) bool) VideoTagCollect {
	sort.SliceStable(s, f)
	return s
}
