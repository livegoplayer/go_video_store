package log

import (
	"sort"
)

func (s LogCollect) Filter(f func(item Log) bool) LogCollect {
	m := make(LogCollect, 0)
	for _, v := range s {
		if f(v) {
			m = append(m, v)
		}
	}
	return m
}
func (s LogCollect) GroupByMessage() map[string]LogCollect {
	m := make(map[string]LogCollect)
	for _, v := range s {
		if _, ok := m[v.Message]; !ok {
			m[v.Message] = make(LogCollect, 0)
		}
		m[v.Message] = append(m[v.Message], v)
	}
	return m
}

func (s LogCollect) GroupByCat() map[int64]LogCollect {
	m := make(map[int64]LogCollect)
	for _, v := range s {
		if _, ok := m[v.Cat]; !ok {
			m[v.Cat] = make(LogCollect, 0)
		}
		m[v.Cat] = append(m[v.Cat], v)
	}
	return m
}

func (s LogCollect) GroupByData() map[string]LogCollect {
	m := make(map[string]LogCollect)
	for _, v := range s {
		if _, ok := m[v.Data]; !ok {
			m[v.Data] = make(LogCollect, 0)
		}
		m[v.Data] = append(m[v.Data], v)
	}
	return m
}

func (s LogCollect) GroupByLevel() map[int64]LogCollect {
	m := make(map[int64]LogCollect)
	for _, v := range s {
		if _, ok := m[v.Level]; !ok {
			m[v.Level] = make(LogCollect, 0)
		}
		m[v.Level] = append(m[v.Level], v)
	}
	return m
}

func (s LogCollect) GroupById() map[int64]LogCollect {
	m := make(map[int64]LogCollect)
	for _, v := range s {
		if _, ok := m[v.Id]; !ok {
			m[v.Id] = make(LogCollect, 0)
		}
		m[v.Id] = append(m[v.Id], v)
	}
	return m
}

func (s LogCollect) GroupByAddDatetime() map[string]LogCollect {
	m := make(map[string]LogCollect)
	for _, v := range s {
		if _, ok := m[v.AddDatetime]; !ok {
			m[v.AddDatetime] = make(LogCollect, 0)
		}
		m[v.AddDatetime] = append(m[v.AddDatetime], v)
	}
	return m
}

func (s LogCollect) KeyByMessage() map[string]Log {
	m := make(map[string]Log)
	for _, v := range s {
		m[v.Message] = v
	}
	return m
}
func (s LogCollect) KeyByCat() map[int64]Log {
	m := make(map[int64]Log)
	for _, v := range s {
		m[v.Cat] = v
	}
	return m
}
func (s LogCollect) KeyByData() map[string]Log {
	m := make(map[string]Log)
	for _, v := range s {
		m[v.Data] = v
	}
	return m
}
func (s LogCollect) KeyByLevel() map[int64]Log {
	m := make(map[int64]Log)
	for _, v := range s {
		m[v.Level] = v
	}
	return m
}
func (s LogCollect) KeyById() map[int64]Log {
	m := make(map[int64]Log)
	for _, v := range s {
		m[v.Id] = v
	}
	return m
}
func (s LogCollect) KeyByAddDatetime() map[string]Log {
	m := make(map[string]Log)
	for _, v := range s {
		m[v.AddDatetime] = v
	}
	return m
}
func (s LogCollect) PluckMessage() []string {
	list := make([]string, len(s))
	for i, v := range s {
		list[i] = v.Message
	}
	return list
}
func (s LogCollect) PluckCat() []int64 {
	list := make([]int64, len(s))
	for i, v := range s {
		list[i] = v.Cat
	}
	return list
}
func (s LogCollect) PluckData() []string {
	list := make([]string, len(s))
	for i, v := range s {
		list[i] = v.Data
	}
	return list
}
func (s LogCollect) PluckLevel() []int64 {
	list := make([]int64, len(s))
	for i, v := range s {
		list[i] = v.Level
	}
	return list
}
func (s LogCollect) PluckId() []int64 {
	list := make([]int64, len(s))
	for i, v := range s {
		list[i] = v.Id
	}
	return list
}
func (s LogCollect) PluckAddDatetime() []string {
	list := make([]string, len(s))
	for i, v := range s {
		list[i] = v.AddDatetime
	}
	return list
}

func (s LogCollect) PluckUniMessage() []string {
	uniMap := make(map[string]bool)
	list := make([]string, 0)
	for _, v := range s {
		_, ok := uniMap[v.Message]
		if !ok {
			uniMap[v.Message] = true
			list = append(list, v.Message)
		}
	}
	return list
}

func (s LogCollect) PluckUniCat() []int64 {
	uniMap := make(map[int64]bool)
	list := make([]int64, 0)
	for _, v := range s {
		_, ok := uniMap[v.Cat]
		if !ok {
			uniMap[v.Cat] = true
			list = append(list, v.Cat)
		}
	}
	return list
}

func (s LogCollect) PluckUniData() []string {
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

func (s LogCollect) PluckUniLevel() []int64 {
	uniMap := make(map[int64]bool)
	list := make([]int64, 0)
	for _, v := range s {
		_, ok := uniMap[v.Level]
		if !ok {
			uniMap[v.Level] = true
			list = append(list, v.Level)
		}
	}
	return list
}

func (s LogCollect) PluckUniId() []int64 {
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

func (s LogCollect) PluckUniAddDatetime() []string {
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

func (s LogCollect) SortByFunc(f func(i, j int) bool) LogCollect {
	sort.SliceStable(s, f)
	return s
}
