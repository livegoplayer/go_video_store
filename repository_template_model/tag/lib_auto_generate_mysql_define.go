package tag

import (
	"github.com/livegoplayer/go_db_helper/mysql"
	"reflect"
)

type WhereCb = mysql.WhereCb

type BeforeHook interface {
	Before()
}

type MustHook interface {
	Must()
}

var AllModels []interface{} // 自动化代码生成器中赋值，用于给外界动态监测model的能力

// TagCollect /** Collection
type TagCollect []Tag

func NewTagQuery() *TagQuery {
	s := TagQuery{}
	s.SetBuild(mysql.NewBuild(&s))
	i, ok := reflect.ValueOf(&s).Interface().(BeforeHook)
	if ok {
		i.Before()
	}
	return &s
}

type _TagQueryColsStruct struct {
	UptDatetime string
	TagName     string
	TagId       string
	Id          string
	TagUrl      string
	AddDatetime string
}

func GetTagQueryCols() *_TagQueryColsStruct {
	return &_TagQueryColsStruct{
		UptDatetime: "int64",
		TagName:     "string",
		TagId:       "int64",
		Id:          "int64",
		TagUrl:      "string",
		AddDatetime: "int64",
	}
}

func (m *TagQuery) First() *Tag {
	s := make([]Tag, 0)
	i, ok := reflect.ValueOf(m).Interface().(MustHook)
	if ok {
		i.Must()
	}
	m.GetBuild().ModelType(&s).Limit(1).Get()
	if len(s) > 0 {
		return &s[0]
	}
	return &Tag{}
}

func (m *TagQuery) GetOne() *Tag {
	s := make([]Tag, 0)
	i, ok := reflect.ValueOf(m).Interface().(MustHook)
	if ok {
		i.Must()
	}
	m.GetBuild().ModelType(&s).Limit(1).Get()
	if len(s) > 0 {
		return &s[0]
	}
	return nil
}

func (m *TagQuery) Get() TagCollect {
	i, ok := reflect.ValueOf(m).Interface().(MustHook)
	if ok {
		i.Must()
	}
	s := make([]Tag, 0)
	m.GetBuild().ModelType(&s).Get()
	return s
}

func (m *TagQuery) clone() *TagQuery {
	nm := NewTagQuery()
	nm.SetBuild(m.GetBuild().Clone())
	return nm
}

func (m *TagQuery) Count() int64 {
	i, ok := reflect.ValueOf(m).Interface().(MustHook)
	if ok {
		i.Must()
	}
	s := Tag{}
	return m.GetBuild().ModelType(&s).Count()
}

func (m *TagQuery) sum(col string) float64 {
	s := Tag{}
	return m.GetBuild().ModelType(&s).Sum(col)
}

func (m *TagQuery) max(col string) float64 {
	s := Tag{}
	return m.GetBuild().ModelType(&s).Max(col)
}

func (m *TagQuery) DoneOperate() int64 {
	s := Tag{}
	return m.GetBuild().ModelType(&s).DoneOperate()
}

func (m *TagQuery) update(h *Tag) int64 {
	s := Tag{}
	return m.GetBuild().ModelType(&s).Update(h)
}

func (m *TagQuery) Delete() int64 {
	s := Tag{}
	return m.GetBuild().ModelType(&s).Delete()
}

func (m *TagQuery) save(x *Tag) {
	m.GetBuild().ModelType(x).Save()
}

func (m *TagQuery) Error() error {
	return m.GetBuild().ModelType(m).Error()
}

// 支持分表
func (m *TagQuery) insert(argu interface{}) {
	s := Tag{}
	m.GetBuild().ModelType(&s).Insert(argu)
}

func (m *TagQuery) increment(column string, amount int) int64 {
	i, ok := reflect.ValueOf(m).Interface().(MustHook)
	if ok {
		i.Must()
	}
	s := Tag{}
	return m.GetBuild().ModelType(&s).Increment(column, amount)
}
