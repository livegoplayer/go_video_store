package tag

import (
	"sort"
)

func (s TagCollect) Filter(f func(item Tag) bool) TagCollect {
	m := make(TagCollect, 0)
	for _, v := range s {
		if f(v) {
			m = append(m, v)
		}
	}
	return m
}
func (s TagCollect) GroupByUptDatetime() map[int64]TagCollect {
	m := make(map[int64]TagCollect)
	for _, v := range s {
		if _, ok := m[v.UptDatetime]; !ok {
			m[v.UptDatetime] = make(TagCollect, 0)
		}
		m[v.UptDatetime] = append(m[v.UptDatetime], v)
	}
	return m
}

func (s TagCollect) GroupByTagName() map[string]TagCollect {
	m := make(map[string]TagCollect)
	for _, v := range s {
		if _, ok := m[v.TagName]; !ok {
			m[v.TagName] = make(TagCollect, 0)
		}
		m[v.TagName] = append(m[v.TagName], v)
	}
	return m
}

func (s TagCollect) GroupByTagId() map[int64]TagCollect {
	m := make(map[int64]TagCollect)
	for _, v := range s {
		if _, ok := m[v.TagId]; !ok {
			m[v.TagId] = make(TagCollect, 0)
		}
		m[v.TagId] = append(m[v.TagId], v)
	}
	return m
}

func (s TagCollect) GroupById() map[int64]TagCollect {
	m := make(map[int64]TagCollect)
	for _, v := range s {
		if _, ok := m[v.Id]; !ok {
			m[v.Id] = make(TagCollect, 0)
		}
		m[v.Id] = append(m[v.Id], v)
	}
	return m
}

func (s TagCollect) GroupByTagUrl() map[string]TagCollect {
	m := make(map[string]TagCollect)
	for _, v := range s {
		if _, ok := m[v.TagUrl]; !ok {
			m[v.TagUrl] = make(TagCollect, 0)
		}
		m[v.TagUrl] = append(m[v.TagUrl], v)
	}
	return m
}

func (s TagCollect) GroupByAddDatetime() map[int64]TagCollect {
	m := make(map[int64]TagCollect)
	for _, v := range s {
		if _, ok := m[v.AddDatetime]; !ok {
			m[v.AddDatetime] = make(TagCollect, 0)
		}
		m[v.AddDatetime] = append(m[v.AddDatetime], v)
	}
	return m
}

func (s TagCollect) KeyByUptDatetime() map[int64]Tag {
	m := make(map[int64]Tag)
	for _, v := range s {
		m[v.UptDatetime] = v
	}
	return m
}
func (s TagCollect) KeyByTagName() map[string]Tag {
	m := make(map[string]Tag)
	for _, v := range s {
		m[v.TagName] = v
	}
	return m
}
func (s TagCollect) KeyByTagId() map[int64]Tag {
	m := make(map[int64]Tag)
	for _, v := range s {
		m[v.TagId] = v
	}
	return m
}
func (s TagCollect) KeyById() map[int64]Tag {
	m := make(map[int64]Tag)
	for _, v := range s {
		m[v.Id] = v
	}
	return m
}
func (s TagCollect) KeyByTagUrl() map[string]Tag {
	m := make(map[string]Tag)
	for _, v := range s {
		m[v.TagUrl] = v
	}
	return m
}
func (s TagCollect) KeyByAddDatetime() map[int64]Tag {
	m := make(map[int64]Tag)
	for _, v := range s {
		m[v.AddDatetime] = v
	}
	return m
}
func (s TagCollect) PluckUptDatetime() []int64 {
	list := make([]int64, len(s))
	for i, v := range s {
		list[i] = v.UptDatetime
	}
	return list
}
func (s TagCollect) PluckTagName() []string {
	list := make([]string, len(s))
	for i, v := range s {
		list[i] = v.TagName
	}
	return list
}
func (s TagCollect) PluckTagId() []int64 {
	list := make([]int64, len(s))
	for i, v := range s {
		list[i] = v.TagId
	}
	return list
}
func (s TagCollect) PluckId() []int64 {
	list := make([]int64, len(s))
	for i, v := range s {
		list[i] = v.Id
	}
	return list
}
func (s TagCollect) PluckTagUrl() []string {
	list := make([]string, len(s))
	for i, v := range s {
		list[i] = v.TagUrl
	}
	return list
}
func (s TagCollect) PluckAddDatetime() []int64 {
	list := make([]int64, len(s))
	for i, v := range s {
		list[i] = v.AddDatetime
	}
	return list
}

func (s TagCollect) PluckUniUptDatetime() []int64 {
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

func (s TagCollect) PluckUniTagName() []string {
	uniMap := make(map[string]bool)
	list := make([]string, 0)
	for _, v := range s {
		_, ok := uniMap[v.TagName]
		if !ok {
			uniMap[v.TagName] = true
			list = append(list, v.TagName)
		}
	}
	return list
}

func (s TagCollect) PluckUniTagId() []int64 {
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

func (s TagCollect) PluckUniId() []int64 {
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

func (s TagCollect) PluckUniTagUrl() []string {
	uniMap := make(map[string]bool)
	list := make([]string, 0)
	for _, v := range s {
		_, ok := uniMap[v.TagUrl]
		if !ok {
			uniMap[v.TagUrl] = true
			list = append(list, v.TagUrl)
		}
	}
	return list
}

func (s TagCollect) PluckUniAddDatetime() []int64 {
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

func (s TagCollect) SortByFunc(f func(i, j int) bool) TagCollect {
	sort.SliceStable(s, f)
	return s
}
