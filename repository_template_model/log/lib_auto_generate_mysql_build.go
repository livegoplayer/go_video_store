package log

// Unscoped 该文件中的所有方法都不应该在model层以外调用，部分共有方法是为了方便model的连表操作;
// must update with github.com.livegoplayer.go_db_helper.mysql 匹配其中的func (m *{name}Query) {funcName} *{name}Query 由于项目初始化之前不一定引入该包，所以采用手写的方式
// 复制当前的内容出一个新的build对象
/****************************************************** 代理方法 start ************************************************************/

func (m *LogQuery) Unscoped() *LogQuery {
	m.GetBuild().Unscoped()
	return m
}

func (m *LogQuery) Select(columns ...interface{}) *LogQuery {
	m.GetBuild().Select(columns...)
	return m
}

func (m *LogQuery) where(column string, args ...interface{}) *LogQuery {
	m.GetBuild().Where(column, args...)
	return m
}

func (m *LogQuery) whereMap(datas map[string]interface{}) *LogQuery {
	m.GetBuild().WhereMap(datas)
	return m
}

func (m *LogQuery) orWhere(column string, args ...interface{}) *LogQuery {
	m.GetBuild().OrWhere(column, args...)
	return m
}

func (m *LogQuery) orWhereCb(cb WhereCb) *LogQuery {
	m.GetBuild().OrWhereCb(cb)
	return m
}

func (m *LogQuery) whereCb(cb WhereCb) *LogQuery {
	m.GetBuild().WhereCb(cb)
	return m
}

func (m *LogQuery) whereRaw(sql string, bindings ...interface{}) *LogQuery {
	m.GetBuild().WhereRaw(sql, bindings...)
	return m
}

func (m *LogQuery) whereNull(column string) *LogQuery {
	m.GetBuild().WhereNull(column)
	return m
}

func (m *LogQuery) whereNotNull(column string) *LogQuery {
	m.GetBuild().WhereNotNull(column)
	return m
}

func (m *LogQuery) orWhereNull(column string) *LogQuery {
	m.GetBuild().OrWhereNull(column)
	return m
}

func (m *LogQuery) orWhereNotNull(column string) *LogQuery {
	m.GetBuild().OrWhereNotNull(column)
	return m
}

func (m *LogQuery) whereIn(column string, values interface{}) *LogQuery {
	m.GetBuild().WhereIn(column, values)
	return m
}

func (m *LogQuery) whereNotIn(column string, values interface{}) *LogQuery {
	m.GetBuild().WhereNotIn(column, values)
	return m
}

func (m *LogQuery) LeftJoin(table string, one string, args ...interface{}) *LogQuery {
	m.GetBuild().LeftJoin(table, one, args...)
	return m
}

func (m *LogQuery) RightJoin(table string, one string, args ...interface{}) *LogQuery {
	m.GetBuild().RightJoin(table, one, args...)
	return m
}

func (m *LogQuery) InnerJoin(table string, one string, args ...interface{}) *LogQuery {
	m.GetBuild().InnerJoin(table, one, args...)
	return m
}

func (m *LogQuery) OrderBy(column string, direction string) *LogQuery {
	m.GetBuild().OrderBy(column, direction)
	return m
}

func (m *LogQuery) OrderDescBy(column string) *LogQuery {
	m.GetBuild().OrderDescBy(column)
	return m
}

func (m *LogQuery) OrderAscBy(column string) *LogQuery {
	m.GetBuild().OrderAscBy(column)
	return m
}

func (m *LogQuery) Group(column string) *LogQuery {
	m.GetBuild().Group(column)
	return m
}

func (m *LogQuery) Inc(col string, num int) *LogQuery {
	m.GetBuild().Inc(col, num)
	return m
}

func (m *LogQuery) Set(col string, val interface{}) *LogQuery {
	m.GetBuild().Set(col, val)
	return m
}

func (m *LogQuery) Skip(lines int) *LogQuery {
	m.GetBuild().Skip(lines)
	return m
}

func (m *LogQuery) Offset(lines int) *LogQuery {
	m.GetBuild().Offset(lines)
	return m
}

func (m *LogQuery) Limit(lines int) *LogQuery {
	m.GetBuild().Limit(lines)
	return m
}

func (m *LogQuery) Take(lines int) *LogQuery {
	m.GetBuild().Take(lines)
	return m
}

func (m *LogQuery) ForUpdate() *LogQuery {
	m.GetBuild().ForUpdate()
	return m
}

func (m *LogQuery) SplitBy(val int64) *LogQuery {
	m.GetBuild().SplitBy(val)
	return m
}

func (m *LogQuery) ModelType(v interface{}) *LogQuery {
	m.GetBuild().ModelType(v)
	return m
}

func (m *LogQuery) DisablePanic() *LogQuery {
	m.GetBuild().DisablePanic()
	return m
}

func (m *LogQuery) OnWrite() *LogQuery {
	m.GetBuild().OnWrite()
	return m
}

/****************************************************** 代理方法 end ************************************************************/
/********************************************** 私有基于字段的build方法 start*****************************************************/

func (m *LogQuery) kWheData(args ...interface{}) *LogQuery {
	return m.where("data", args...)
}

func (m *LogQuery) kSetData(val interface{}) *LogQuery {
	return m.Set("data", val)
}

func (m *LogQuery) kWheDataIn(values interface{}) *LogQuery {
	return m.whereIn("data", values)
}

func (m *LogQuery) kWheDataNotIn(values interface{}) *LogQuery {
	return m.whereNotIn("data", values)
}

func (m *LogQuery) kWheLevel(args ...interface{}) *LogQuery {
	return m.where("level", args...)
}

func (m *LogQuery) kSetLevel(val interface{}) *LogQuery {
	return m.Set("level", val)
}

func (m *LogQuery) kIncLevel(num int) *LogQuery {
	return m.Inc("level", num)
}

func (m *LogQuery) kWheLevelIn(values interface{}) *LogQuery {
	return m.whereIn("level", values)
}

func (m *LogQuery) kWheLevelNotIn(values interface{}) *LogQuery {
	return m.whereNotIn("level", values)
}

func (m *LogQuery) kWheCat(args ...interface{}) *LogQuery {
	return m.where("cat", args...)
}

func (m *LogQuery) kSetCat(val interface{}) *LogQuery {
	return m.Set("cat", val)
}

func (m *LogQuery) kIncCat(num int) *LogQuery {
	return m.Inc("cat", num)
}

func (m *LogQuery) kWheCatIn(values interface{}) *LogQuery {
	return m.whereIn("cat", values)
}

func (m *LogQuery) kWheCatNotIn(values interface{}) *LogQuery {
	return m.whereNotIn("cat", values)
}

func (m *LogQuery) kWheId(args ...interface{}) *LogQuery {
	return m.where("id", args...)
}

func (m *LogQuery) kSetId(val interface{}) *LogQuery {
	return m.Set("id", val)
}

func (m *LogQuery) kIncId(num int) *LogQuery {
	return m.Inc("id", num)
}

func (m *LogQuery) kWheIdIn(values interface{}) *LogQuery {
	return m.whereIn("id", values)
}

func (m *LogQuery) kWheIdNotIn(values interface{}) *LogQuery {
	return m.whereNotIn("id", values)
}

func (m *LogQuery) kWheMessage(args ...interface{}) *LogQuery {
	return m.where("message", args...)
}

func (m *LogQuery) kSetMessage(val interface{}) *LogQuery {
	return m.Set("message", val)
}

func (m *LogQuery) kWheMessageIn(values interface{}) *LogQuery {
	return m.whereIn("message", values)
}

func (m *LogQuery) kWheMessageNotIn(values interface{}) *LogQuery {
	return m.whereNotIn("message", values)
}

func (m *LogQuery) kWheAddDatetime(args ...interface{}) *LogQuery {
	return m.where("add_datetime", args...)
}

func (m *LogQuery) kSetAddDatetime(val interface{}) *LogQuery {
	return m.Set("add_datetime", val)
}

func (m *LogQuery) kWheAddDatetimeIn(values interface{}) *LogQuery {
	return m.whereIn("add_datetime", values)
}

func (m *LogQuery) kWheAddDatetimeNotIn(values interface{}) *LogQuery {
	return m.whereNotIn("add_datetime", values)
}

/******************************************* 私有基于字段的build方法 end *****************************************************/
