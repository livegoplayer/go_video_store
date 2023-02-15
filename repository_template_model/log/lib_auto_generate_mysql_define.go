package log

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

// LogCollect /** Collection
type LogCollect []Log

func NewLogQuery() *LogQuery {
	s := LogQuery{}
	s.SetBuild(mysql.NewBuild(&s))
	i, ok := reflect.ValueOf(&s).Interface().(BeforeHook)
	if ok {
		i.Before()
	}
	return &s
}

type _LogQueryColsStruct struct {
	Data        string
	Level       string
	Cat         string
	Id          string
	Message     string
	AddDatetime string
}

func GetLogQueryCols() *_LogQueryColsStruct {
	return &_LogQueryColsStruct{
		Data:        "string",
		Level:       "int64",
		Cat:         "int64",
		Id:          "int64",
		Message:     "string",
		AddDatetime: "string",
	}
}

func (cols *_LogQueryColsStruct) GetColsTags() []string {
	return []string{
		"data",
		"level",
		"cat",
		"id",
		"message",
		"add_datetime",
	}
}

func (cols *_LogQueryColsStruct) Has(colName string) bool {
	for _, clo := range cols.GetColsTags() {
		if clo == colName {
			return true
		}
	}
	return false
}

func (cols *_LogQueryColsStruct) GetPrimaryColsStr() string {
	primaryIndexList := []string{
		"id",
	}
	return strings.Join(primaryIndexList, ",")
}

func (m *LogQuery) First() *Log {
	s := make([]Log, 0)
	i, ok := reflect.ValueOf(m).Interface().(MustHook)
	if ok {
		i.Must()
	}
	m.GetBuild().ModelType(&s).Limit(1).Get()
	if len(s) > 0 {
		return &s[0]
	}
	return &Log{}
}

func (m *LogQuery) GetOne() *Log {
	s := make([]Log, 0)
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

func (m *LogQuery) Get() LogCollect {
	i, ok := reflect.ValueOf(m).Interface().(MustHook)
	if ok {
		i.Must()
	}
	s := make([]Log, 0)
	m.GetBuild().ModelType(&s).Get()
	return s
}

func (m *LogQuery) clone() *LogQuery {
	nm := NewLogQuery()
	nm.SetBuild(m.GetBuild().Clone())
	return nm
}

func (m *LogQuery) Count() int64 {
	i, ok := reflect.ValueOf(m).Interface().(MustHook)
	if ok {
		i.Must()
	}
	s := Log{}
	return m.GetBuild().ModelType(&s).Count()
}

func (m *LogQuery) sum(col string) float64 {
	s := Log{}
	return m.GetBuild().ModelType(&s).Sum(col)
}

func (m *LogQuery) max(col string) float64 {
	s := Log{}
	return m.GetBuild().ModelType(&s).Max(col)
}

func (m *LogQuery) DoneOperate() int64 {
	s := Log{}
	return m.GetBuild().ModelType(&s).DoneOperate()
}

func (m *LogQuery) update(h *Log) int64 {
	s := Log{}
	return m.GetBuild().ModelType(&s).Update(h)
}

func (m *LogQuery) Delete() int64 {
	s := Log{}
	return m.GetBuild().ModelType(&s).Delete()
}

func (m *LogQuery) save(x *Log) {
	m.GetBuild().ModelType(x).Save()
}

func (m *LogQuery) Error() error {
	return m.GetBuild().ModelType(m).Error()
}

// 支持分表
func (m *LogQuery) insert(argu interface{}) {
	s := Log{}
	m.GetBuild().ModelType(&s).Insert(argu)
}

func (m *LogQuery) increment(column string, amount int) int64 {
	i, ok := reflect.ValueOf(m).Interface().(MustHook)
	if ok {
		i.Must()
	}
	s := Log{}
	return m.GetBuild().ModelType(&s).Increment(column, amount)
}
