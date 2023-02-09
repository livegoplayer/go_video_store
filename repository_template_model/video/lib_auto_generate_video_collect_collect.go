package video

import (
	"sort"
)

func (s VideoCollect) Filter(f func(item Video) bool) VideoCollect {
	m := make(VideoCollect, 0)
	for _, v := range s {
		if f(v) {
			m = append(m, v)
		}
	}
	return m
}
func (s VideoCollect) GroupByUptDatetime() map[int64]VideoCollect {
	m := make(map[int64]VideoCollect)
	for _, v := range s {
		if _, ok := m[v.UptDatetime]; !ok {
			m[v.UptDatetime] = make(VideoCollect, 0)
		}
		m[v.UptDatetime] = append(m[v.UptDatetime], v)
	}
	return m
}

func (s VideoCollect) GroupByVideoId() map[int64]VideoCollect {
	m := make(map[int64]VideoCollect)
	for _, v := range s {
		if _, ok := m[v.VideoId]; !ok {
			m[v.VideoId] = make(VideoCollect, 0)
		}
		m[v.VideoId] = append(m[v.VideoId], v)
	}
	return m
}

func (s VideoCollect) GroupByPrettyName() map[string]VideoCollect {
	m := make(map[string]VideoCollect)
	for _, v := range s {
		if _, ok := m[v.PrettyName]; !ok {
			m[v.PrettyName] = make(VideoCollect, 0)
		}
		m[v.PrettyName] = append(m[v.PrettyName], v)
	}
	return m
}

func (s VideoCollect) GroupByMetaName() map[string]VideoCollect {
	m := make(map[string]VideoCollect)
	for _, v := range s {
		if _, ok := m[v.MetaName]; !ok {
			m[v.MetaName] = make(VideoCollect, 0)
		}
		m[v.MetaName] = append(m[v.MetaName], v)
	}
	return m
}

func (s VideoCollect) GroupByAddDatetime() map[int64]VideoCollect {
	m := make(map[int64]VideoCollect)
	for _, v := range s {
		if _, ok := m[v.AddDatetime]; !ok {
			m[v.AddDatetime] = make(VideoCollect, 0)
		}
		m[v.AddDatetime] = append(m[v.AddDatetime], v)
	}
	return m
}

func (s VideoCollect) KeyByUptDatetime() map[int64]Video {
	m := make(map[int64]Video)
	for _, v := range s {
		m[v.UptDatetime] = v
	}
	return m
}
func (s VideoCollect) KeyByVideoId() map[int64]Video {
	m := make(map[int64]Video)
	for _, v := range s {
		m[v.VideoId] = v
	}
	return m
}
func (s VideoCollect) KeyByPrettyName() map[string]Video {
	m := make(map[string]Video)
	for _, v := range s {
		m[v.PrettyName] = v
	}
	return m
}
func (s VideoCollect) KeyByMetaName() map[string]Video {
	m := make(map[string]Video)
	for _, v := range s {
		m[v.MetaName] = v
	}
	return m
}
func (s VideoCollect) KeyByAddDatetime() map[int64]Video {
	m := make(map[int64]Video)
	for _, v := range s {
		m[v.AddDatetime] = v
	}
	return m
}
func (s VideoCollect) PluckUptDatetime() []int64 {
	list := make([]int64, len(s))
	for i, v := range s {
		list[i] = v.UptDatetime
	}
	return list
}
func (s VideoCollect) PluckVideoId() []int64 {
	list := make([]int64, len(s))
	for i, v := range s {
		list[i] = v.VideoId
	}
	return list
}
func (s VideoCollect) PluckPrettyName() []string {
	list := make([]string, len(s))
	for i, v := range s {
		list[i] = v.PrettyName
	}
	return list
}
func (s VideoCollect) PluckMetaName() []string {
	list := make([]string, len(s))
	for i, v := range s {
		list[i] = v.MetaName
	}
	return list
}
func (s VideoCollect) PluckAddDatetime() []int64 {
	list := make([]int64, len(s))
	for i, v := range s {
		list[i] = v.AddDatetime
	}
	return list
}

func (s VideoCollect) PluckUniUptDatetime() []int64 {
	uniMap := make(map[int64]bool)
	list := make([]int64, 0)
	for _, v := range s {
		_, ok := uniMap[v.UptDatetime]
		if !ok {
			uniMap[v.UptDatetime] = true
			list = append(list, v.UptDatetime)
		}
	}
	return list
}

func (s VideoCollect) PluckUniVideoId() []int64 {
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

func (s VideoCollect) PluckUniPrettyName() []string {
	uniMap := make(map[string]bool)
	list := make([]string, 0)
	for _, v := range s {
		_, ok := uniMap[v.PrettyName]
		if !ok {
			uniMap[v.PrettyName] = true
			list = append(list, v.PrettyName)
		}
	}
	return list
}

func (s VideoCollect) PluckUniMetaName() []string {
	uniMap := make(map[string]bool)
	list := make([]string, 0)
	for _, v := range s {
		_, ok := uniMap[v.MetaName]
		if !ok {
			uniMap[v.MetaName] = true
			list = append(list, v.MetaName)
		}
	}
	return list
}

func (s VideoCollect) PluckUniAddDatetime() []int64 {
	uniMap := make(map[int64]bool)
	list := make([]int64, 0)
	for _, v := range s {
		_, ok := uniMap[v.AddDatetime]
		if !ok {
			uniMap[v.AddDatetime] = true
			list = append(list, v.AddDatetime)
		}
	}
	return list
}

func (s VideoCollect) SortByFunc(f func(i, j int) bool) VideoCollect {
	sort.SliceStable(s, f)
	return s
}
