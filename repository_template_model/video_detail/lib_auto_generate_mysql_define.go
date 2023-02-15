package video_detail

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

// VideoDetailCollect /** Collection
type VideoDetailCollect []VideoDetail

func NewVideoDetailQuery() *VideoDetailQuery {
	s := VideoDetailQuery{}
	s.SetBuild(mysql.NewBuild(&s))
	i, ok := reflect.ValueOf(&s).Interface().(BeforeHook)
	if ok {
		i.Before()
	}
	return &s
}

type _VideoDetailQueryColsStruct struct {
	OfficialWebsite  string
	DescriptionUrl   string
	Year             string
	ReleasedDateStr  string
	VideoPosterUrl   string
	TimeStr          string
	PosterCache      string
	VideoDescription string
	Score            string
	PosterListUrl    string
	UptDatetime      string
	PerUpdateTime    string
	Season           string
	Id               string
	SetNum           string
	AddDatetime      string
	VideoId          string
}

func GetVideoDetailQueryCols() *_VideoDetailQueryColsStruct {
	return &_VideoDetailQueryColsStruct{
		OfficialWebsite:  "string",
		DescriptionUrl:   "string",
		Year:             "int64",
		ReleasedDateStr:  "string",
		VideoPosterUrl:   "string",
		TimeStr:          "string",
		PosterCache:      "string",
		VideoDescription: "string",
		Score:            "float64",
		PosterListUrl:    "string",
		UptDatetime:      "string",
		PerUpdateTime:    "string",
		Season:           "int64",
		Id:               "int64",
		SetNum:           "int64",
		AddDatetime:      "string",
		VideoId:          "int64",
	}
}

func (cols *_VideoDetailQueryColsStruct) GetColsTags() []string {
	return []string{
		"official_website",
		"description_url",
		"year",
		"released_date_str",
		"video_poster_url",
		"time_str",
		"poster_cache",
		"video_description",
		"score",
		"poster_list_url",
		"upt_datetime",
		"per_update_time",
		"season",
		"id",
		"set_num",
		"add_datetime",
		"video_id",
	}
}

func (cols *_VideoDetailQueryColsStruct) Has(colName string) bool {
	for _, clo := range cols.GetColsTags() {
		if clo == colName {
			return true
		}
	}
	return false
}

func (cols *_VideoDetailQueryColsStruct) GetPrimaryColsStr() string {
	primaryIndexList := []string{
		"id",
	}
	return strings.Join(primaryIndexList, ",")
}

func (m *VideoDetailQuery) First() *VideoDetail {
	s := make([]VideoDetail, 0)
	i, ok := reflect.ValueOf(m).Interface().(MustHook)
	if ok {
		i.Must()
	}
	m.GetBuild().ModelType(&s).Limit(1).Get()
	if len(s) > 0 {
		return &s[0]
	}
	return &VideoDetail{}
}

func (m *VideoDetailQuery) GetOne() *VideoDetail {
	s := make([]VideoDetail, 0)
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

func (m *VideoDetailQuery) Get() VideoDetailCollect {
	i, ok := reflect.ValueOf(m).Interface().(MustHook)
	if ok {
		i.Must()
	}
	s := make([]VideoDetail, 0)
	m.GetBuild().ModelType(&s).Get()
	return s
}

func (m *VideoDetailQuery) clone() *VideoDetailQuery {
	nm := NewVideoDetailQuery()
	nm.SetBuild(m.GetBuild().Clone())
	return nm
}

func (m *VideoDetailQuery) Count() int64 {
	i, ok := reflect.ValueOf(m).Interface().(MustHook)
	if ok {
		i.Must()
	}
	s := VideoDetail{}
	return m.GetBuild().ModelType(&s).Count()
}

func (m *VideoDetailQuery) sum(col string) float64 {
	s := VideoDetail{}
	return m.GetBuild().ModelType(&s).Sum(col)
}

func (m *VideoDetailQuery) max(col string) float64 {
	s := VideoDetail{}
	return m.GetBuild().ModelType(&s).Max(col)
}

func (m *VideoDetailQuery) DoneOperate() int64 {
	s := VideoDetail{}
	return m.GetBuild().ModelType(&s).DoneOperate()
}

func (m *VideoDetailQuery) update(h *VideoDetail) int64 {
	s := VideoDetail{}
	return m.GetBuild().ModelType(&s).Update(h)
}

func (m *VideoDetailQuery) Delete() int64 {
	s := VideoDetail{}
	return m.GetBuild().ModelType(&s).Delete()
}

func (m *VideoDetailQuery) save(x *VideoDetail) {
	m.GetBuild().ModelType(x).Save()
}

func (m *VideoDetailQuery) Error() error {
	return m.GetBuild().ModelType(m).Error()
}

// 支持分表
func (m *VideoDetailQuery) insert(argu interface{}) {
	s := VideoDetail{}
	m.GetBuild().ModelType(&s).Insert(argu)
}

func (m *VideoDetailQuery) increment(column string, amount int) int64 {
	i, ok := reflect.ValueOf(m).Interface().(MustHook)
	if ok {
		i.Must()
	}
	s := VideoDetail{}
	return m.GetBuild().ModelType(&s).Increment(column, amount)
}
