package video

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

// VideoCollect /** Collection
type VideoCollect []Video

func NewVideoQuery() *VideoQuery {
	s := VideoQuery{}
	s.SetBuild(mysql.NewBuild(&s))
	i, ok := reflect.ValueOf(&s).Interface().(BeforeHook)
	if ok {
		i.Before()
	}
	return &s
}

type _VideoQueryColsStruct struct {
	PrettyName  string
	MetaName    string
	UptDatetime string
	AddDatetime string
	VideoId     string
}

func GetVideoQueryCols() *_VideoQueryColsStruct {
	return &_VideoQueryColsStruct{
		PrettyName:  "string",
		MetaName:    "string",
		UptDatetime: "string",
		AddDatetime: "string",
		VideoId:     "int64",
	}
}

func (cols *_VideoQueryColsStruct) GetColsTags() []string {
	return []string{
		"pretty_name",
		"meta_name",
		"upt_datetime",
		"add_datetime",
		"video_id",
	}
}

func (cols *_VideoQueryColsStruct) Has(colName string) bool {
	for _, clo := range cols.GetColsTags() {
		if clo == colName {
			return true
		}
	}
	return false
}

func (cols *_VideoQueryColsStruct) GetPrimaryColsStr() string {
	primaryIndexList := []string{
		"video_id",
	}
	return strings.Join(primaryIndexList, ",")
}

func (m *VideoQuery) First() *Video {
	s := make([]Video, 0)
	i, ok := reflect.ValueOf(m).Interface().(MustHook)
	if ok {
		i.Must()
	}
	m.GetBuild().ModelType(&s).Limit(1).Get()
	if len(s) > 0 {
		return &s[0]
	}
	return &Video{}
}

func (m *VideoQuery) GetOne() *Video {
	s := make([]Video, 0)
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

func (m *VideoQuery) Get() VideoCollect {
	i, ok := reflect.ValueOf(m).Interface().(MustHook)
	if ok {
		i.Must()
	}
	s := make([]Video, 0)
	m.GetBuild().ModelType(&s).Get()
	return s
}

func (m *VideoQuery) clone() *VideoQuery {
	nm := NewVideoQuery()
	nm.SetBuild(m.GetBuild().Clone())
	return nm
}

func (m *VideoQuery) Count() int64 {
	i, ok := reflect.ValueOf(m).Interface().(MustHook)
	if ok {
		i.Must()
	}
	s := Video{}
	return m.GetBuild().ModelType(&s).Count()
}

func (m *VideoQuery) sum(col string) float64 {
	s := Video{}
	return m.GetBuild().ModelType(&s).Sum(col)
}

func (m *VideoQuery) max(col string) float64 {
	s := Video{}
	return m.GetBuild().ModelType(&s).Max(col)
}

func (m *VideoQuery) DoneOperate() int64 {
	s := Video{}
	return m.GetBuild().ModelType(&s).DoneOperate()
}

func (m *VideoQuery) update(h *Video) int64 {
	s := Video{}
	return m.GetBuild().ModelType(&s).Update(h)
}

func (m *VideoQuery) Delete() int64 {
	s := Video{}
	return m.GetBuild().ModelType(&s).Delete()
}

func (m *VideoQuery) save(x *Video) {
	m.GetBuild().ModelType(x).Save()
}

func (m *VideoQuery) Error() error {
	return m.GetBuild().ModelType(m).Error()
}

// 支持分表
func (m *VideoQuery) insert(argu interface{}) {
	s := Video{}
	m.GetBuild().ModelType(&s).Insert(argu)
}

func (m *VideoQuery) increment(column string, amount int) int64 {
	i, ok := reflect.ValueOf(m).Interface().(MustHook)
	if ok {
		i.Must()
	}
	s := Video{}
	return m.GetBuild().ModelType(&s).Increment(column, amount)
}
