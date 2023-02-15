package video_tag

import (
	"github.com/livegoplayer/go_db_helper/mysql"
	"reflect"
	"strings"
)

type WhereCb = mysql.WhereCb

type BeforeHook interface {
	Before()
}

type MustHook interface {
	Must()
}

var AllModels []interface{} // 自动化代码生成器中赋值，用于给外界动态监测model的能力

// VideoTagCollect /** Collection
type VideoTagCollect []VideoTag

func NewVideoTagQuery() *VideoTagQuery {
	s := VideoTagQuery{}
	s.SetBuild(mysql.NewBuild(&s))
	i, ok := reflect.ValueOf(&s).Interface().(BeforeHook)
	if ok {
		i.Before()
	}
	return &s
}

type _VideoTagQueryColsStruct struct {
	UptDatetime string
	TagId       string
	Id          string
	AddDatetime string
	VideoId     string
}

func GetVideoTagQueryCols() *_VideoTagQueryColsStruct {
	return &_VideoTagQueryColsStruct{
		UptDatetime: "string",
		TagId:       "int64",
		Id:          "int64",
		AddDatetime: "string",
		VideoId:     "int64",
	}
}

func (cols *_VideoTagQueryColsStruct) GetColsTags() []string {
	return []string{
		"upt_datetime",
		"tag_id",
		"id",
		"add_datetime",
		"video_id",
	}
}

func (cols *_VideoTagQueryColsStruct) Has(colName string) bool {
	for _, clo := range cols.GetColsTags() {
		if clo == colName {
			return true
		}
	}
	return false
}

func (cols *_VideoTagQueryColsStruct) GetPrimaryColsStr() string {
	primaryIndexList := []string{
		"id",
	}
	return strings.Join(primaryIndexList, ",")
}

func (m *VideoTagQuery) First() *VideoTag {
	s := make([]VideoTag, 0)
	i, ok := reflect.ValueOf(m).Interface().(MustHook)
	if ok {
		i.Must()
	}
	m.GetBuild().ModelType(&s).Limit(1).Get()
	if len(s) > 0 {
		return &s[0]
	}
	return &VideoTag{}
}

func (m *VideoTagQuery) GetOne() *VideoTag {
	s := make([]VideoTag, 0)
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

func (m *VideoTagQuery) Get() VideoTagCollect {
	i, ok := reflect.ValueOf(m).Interface().(MustHook)
	if ok {
		i.Must()
	}
	s := make([]VideoTag, 0)
	m.GetBuild().ModelType(&s).Get()
	return s
}

func (m *VideoTagQuery) clone() *VideoTagQuery {
	nm := NewVideoTagQuery()
	nm.SetBuild(m.GetBuild().Clone())
	return nm
}

func (m *VideoTagQuery) Count() int64 {
	i, ok := reflect.ValueOf(m).Interface().(MustHook)
	if ok {
		i.Must()
	}
	s := VideoTag{}
	return m.GetBuild().ModelType(&s).Count()
}

func (m *VideoTagQuery) sum(col string) float64 {
	s := VideoTag{}
	return m.GetBuild().ModelType(&s).Sum(col)
}

func (m *VideoTagQuery) max(col string) float64 {
	s := VideoTag{}
	return m.GetBuild().ModelType(&s).Max(col)
}

func (m *VideoTagQuery) DoneOperate() int64 {
	s := VideoTag{}
	return m.GetBuild().ModelType(&s).DoneOperate()
}

func (m *VideoTagQuery) update(h *VideoTag) int64 {
	s := VideoTag{}
	return m.GetBuild().ModelType(&s).Update(h)
}

func (m *VideoTagQuery) Delete() int64 {
	s := VideoTag{}
	return m.GetBuild().ModelType(&s).Delete()
}

func (m *VideoTagQuery) save(x *VideoTag) {
	m.GetBuild().ModelType(x).Save()
}

func (m *VideoTagQuery) Error() error {
	return m.GetBuild().ModelType(m).Error()
}

// 支持分表
func (m *VideoTagQuery) insert(argu interface{}) {
	s := VideoTag{}
	m.GetBuild().ModelType(&s).Insert(argu)
}

func (m *VideoTagQuery) increment(column string, amount int) int64 {
	i, ok := reflect.ValueOf(m).Interface().(MustHook)
	if ok {
		i.Must()
	}
	s := VideoTag{}
	return m.GetBuild().ModelType(&s).Increment(column, amount)
}
