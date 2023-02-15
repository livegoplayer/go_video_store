package video_resources

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

// VideoResourcesCollect /** Collection
type VideoResourcesCollect []VideoResources

func NewVideoResourcesQuery() *VideoResourcesQuery {
	s := VideoResourcesQuery{}
	s.SetBuild(mysql.NewBuild(&s))
	i, ok := reflect.ValueOf(&s).Interface().(BeforeHook)
	if ok {
		i.Before()
	}
	return &s
}

type _VideoResourcesQueryColsStruct struct {
	LastUpdateTime   string
	VideoUrl         string
	Data             string
	VideoReferDomain string
	UptDatetime      string
	Name             string
	Season           string
	Episode          string
	Id               string
	VideoRefer       string
	AddDatetime      string
	VideoId          string
}

func GetVideoResourcesQueryCols() *_VideoResourcesQueryColsStruct {
	return &_VideoResourcesQueryColsStruct{
		LastUpdateTime:   "int64",
		VideoUrl:         "string",
		Data:             "string",
		VideoReferDomain: "string",
		UptDatetime:      "string",
		Name:             "string",
		Season:           "int64",
		Episode:          "int64",
		Id:               "int64",
		VideoRefer:       "int64",
		AddDatetime:      "string",
		VideoId:          "int64",
	}
}

func (cols *_VideoResourcesQueryColsStruct) GetColsTags() []string {
	return []string{
		"last_update_time",
		"video_url",
		"data",
		"video_refer_domain",
		"upt_datetime",
		"name",
		"season",
		"episode",
		"id",
		"video_refer",
		"add_datetime",
		"video_id",
	}
}

func (cols *_VideoResourcesQueryColsStruct) Has(colName string) bool {
	for _, clo := range cols.GetColsTags() {
		if clo == colName {
			return true
		}
	}
	return false
}

func (cols *_VideoResourcesQueryColsStruct) GetPrimaryColsStr() string {
	primaryIndexList := []string{
		"id",
	}
	return strings.Join(primaryIndexList, ",")
}

func (m *VideoResourcesQuery) First() *VideoResources {
	s := make([]VideoResources, 0)
	i, ok := reflect.ValueOf(m).Interface().(MustHook)
	if ok {
		i.Must()
	}
	m.GetBuild().ModelType(&s).Limit(1).Get()
	if len(s) > 0 {
		return &s[0]
	}
	return &VideoResources{}
}

func (m *VideoResourcesQuery) GetOne() *VideoResources {
	s := make([]VideoResources, 0)
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

func (m *VideoResourcesQuery) Get() VideoResourcesCollect {
	i, ok := reflect.ValueOf(m).Interface().(MustHook)
	if ok {
		i.Must()
	}
	s := make([]VideoResources, 0)
	m.GetBuild().ModelType(&s).Get()
	return s
}

func (m *VideoResourcesQuery) clone() *VideoResourcesQuery {
	nm := NewVideoResourcesQuery()
	nm.SetBuild(m.GetBuild().Clone())
	return nm
}

func (m *VideoResourcesQuery) Count() int64 {
	i, ok := reflect.ValueOf(m).Interface().(MustHook)
	if ok {
		i.Must()
	}
	s := VideoResources{}
	return m.GetBuild().ModelType(&s).Count()
}

func (m *VideoResourcesQuery) sum(col string) float64 {
	s := VideoResources{}
	return m.GetBuild().ModelType(&s).Sum(col)
}

func (m *VideoResourcesQuery) max(col string) float64 {
	s := VideoResources{}
	return m.GetBuild().ModelType(&s).Max(col)
}

func (m *VideoResourcesQuery) DoneOperate() int64 {
	s := VideoResources{}
	return m.GetBuild().ModelType(&s).DoneOperate()
}

func (m *VideoResourcesQuery) update(h *VideoResources) int64 {
	s := VideoResources{}
	return m.GetBuild().ModelType(&s).Update(h)
}

func (m *VideoResourcesQuery) Delete() int64 {
	s := VideoResources{}
	return m.GetBuild().ModelType(&s).Delete()
}

func (m *VideoResourcesQuery) save(x *VideoResources) {
	m.GetBuild().ModelType(x).Save()
}

func (m *VideoResourcesQuery) Error() error {
	return m.GetBuild().ModelType(m).Error()
}

// 支持分表
func (m *VideoResourcesQuery) insert(argu interface{}) {
	s := VideoResources{}
	m.GetBuild().ModelType(&s).Insert(argu)
}

func (m *VideoResourcesQuery) increment(column string, amount int) int64 {
	i, ok := reflect.ValueOf(m).Interface().(MustHook)
	if ok {
		i.Must()
	}
	s := VideoResources{}
	return m.GetBuild().ModelType(&s).Increment(column, amount)
}
