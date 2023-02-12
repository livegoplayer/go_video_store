package video

// Unscoped 该文件中的所有方法都不应该在model层以外调用，部分共有方法是为了方便model的连表操作;
// must update with github.com.livegoplayer.go_db_helper.mysql 匹配其中的func (m *{name}Query) {funcName} *{name}Query 由于项目初始化之前不一定引入该包，所以采用手写的方式
// 复制当前的内容出一个新的build对象
/****************************************************** 代理方法 start ************************************************************/

func (m *VideoQuery) Unscoped() *VideoQuery {
	m.GetBuild().Unscoped()
	return m
}

func (m *VideoQuery) Select(columns ...interface{}) *VideoQuery {
	m.GetBuild().Select(columns...)
	return m
}

func (m *VideoQuery) where(column string, args ...interface{}) *VideoQuery {
	m.GetBuild().Where(column, args...)
	return m
}

func (m *VideoQuery) whereMap(datas map[string]interface{}) *VideoQuery {
	m.GetBuild().WhereMap(datas)
	return m
}

func (m *VideoQuery) orWhere(column string, args ...interface{}) *VideoQuery {
	m.GetBuild().OrWhere(column, args...)
	return m
}

func (m *VideoQuery) orWhereCb(cb WhereCb) *VideoQuery {
	m.GetBuild().OrWhereCb(cb)
	return m
}

func (m *VideoQuery) whereCb(cb WhereCb) *VideoQuery {
	m.GetBuild().WhereCb(cb)
	return m
}

func (m *VideoQuery) whereRaw(sql string, bindings ...interface{}) *VideoQuery {
	m.GetBuild().WhereRaw(sql, bindings...)
	return m
}

func (m *VideoQuery) whereNull(column string) *VideoQuery {
	m.GetBuild().WhereNull(column)
	return m
}

func (m *VideoQuery) whereNotNull(column string) *VideoQuery {
	m.GetBuild().WhereNotNull(column)
	return m
}

func (m *VideoQuery) orWhereNull(column string) *VideoQuery {
	m.GetBuild().OrWhereNull(column)
	return m
}

func (m *VideoQuery) orWhereNotNull(column string) *VideoQuery {
	m.GetBuild().OrWhereNotNull(column)
	return m
}

func (m *VideoQuery) whereIn(column string, values interface{}) *VideoQuery {
	m.GetBuild().WhereIn(column, values)
	return m
}

func (m *VideoQuery) whereNotIn(column string, values interface{}) *VideoQuery {
	m.GetBuild().WhereNotIn(column, values)
	return m
}

func (m *VideoQuery) LeftJoin(table string, one string, args ...interface{}) *VideoQuery {
	m.GetBuild().LeftJoin(table, one, args...)
	return m
}

func (m *VideoQuery) RightJoin(table string, one string, args ...interface{}) *VideoQuery {
	m.GetBuild().RightJoin(table, one, args...)
	return m
}

func (m *VideoQuery) InnerJoin(table string, one string, args ...interface{}) *VideoQuery {
	m.GetBuild().InnerJoin(table, one, args...)
	return m
}

func (m *VideoQuery) OrderBy(column string, direction string) *VideoQuery {
	m.GetBuild().OrderBy(column, direction)
	return m
}

func (m *VideoQuery) OrderDescBy(column string) *VideoQuery {
	m.GetBuild().OrderDescBy(column)
	return m
}

func (m *VideoQuery) OrderAscBy(column string) *VideoQuery {
	m.GetBuild().OrderAscBy(column)
	return m
}

func (m *VideoQuery) Group(column string) *VideoQuery {
	m.GetBuild().Group(column)
	return m
}

func (m *VideoQuery) Inc(col string, num int) *VideoQuery {
	m.GetBuild().Inc(col, num)
	return m
}

func (m *VideoQuery) Set(col string, val interface{}) *VideoQuery {
	m.GetBuild().Set(col, val)
	return m
}

func (m *VideoQuery) Skip(lines int) *VideoQuery {
	m.GetBuild().Skip(lines)
	return m
}

func (m *VideoQuery) Offset(lines int) *VideoQuery {
	m.GetBuild().Offset(lines)
	return m
}

func (m *VideoQuery) Limit(lines int) *VideoQuery {
	m.GetBuild().Limit(lines)
	return m
}

func (m *VideoQuery) Take(lines int) *VideoQuery {
	m.GetBuild().Take(lines)
	return m
}

func (m *VideoQuery) ForUpdate() *VideoQuery {
	m.GetBuild().ForUpdate()
	return m
}

func (m *VideoQuery) SplitBy(val int64) *VideoQuery {
	m.GetBuild().SplitBy(val)
	return m
}

func (m *VideoQuery) ModelType(v interface{}) *VideoQuery {
	m.GetBuild().ModelType(v)
	return m
}

func (m *VideoQuery) DisablePanic() *VideoQuery {
	m.GetBuild().DisablePanic()
	return m
}

func (m *VideoQuery) OnWrite() *VideoQuery {
	m.GetBuild().OnWrite()
	return m
}

/****************************************************** 代理方法 end ************************************************************/
/********************************************** 私有基于字段的build方法 start*****************************************************/

func (m *VideoQuery) kWhePrettyName(args ...interface{}) *VideoQuery {
	return m.where("pretty_name", args...)
}

func (m *VideoQuery) kSetPrettyName(val interface{}) *VideoQuery {
	return m.Set("pretty_name", val)
}

func (m *VideoQuery) kWhePrettyNameIn(values interface{}) *VideoQuery {
	return m.whereIn("pretty_name", values)
}

func (m *VideoQuery) kWhePrettyNameNotIn(values interface{}) *VideoQuery {
	return m.whereNotIn("pretty_name", values)
}

func (m *VideoQuery) kWheMetaName(args ...interface{}) *VideoQuery {
	return m.where("meta_name", args...)
}

func (m *VideoQuery) kSetMetaName(val interface{}) *VideoQuery {
	return m.Set("meta_name", val)
}

func (m *VideoQuery) kWheMetaNameIn(values interface{}) *VideoQuery {
	return m.whereIn("meta_name", values)
}

func (m *VideoQuery) kWheMetaNameNotIn(values interface{}) *VideoQuery {
	return m.whereNotIn("meta_name", values)
}

func (m *VideoQuery) kWheUptDatetime(args ...interface{}) *VideoQuery {
	return m.where("upt_datetime", args...)
}

func (m *VideoQuery) kSetUptDatetime(val interface{}) *VideoQuery {
	return m.Set("upt_datetime", val)
}

func (m *VideoQuery) kWheUptDatetimeIn(values interface{}) *VideoQuery {
	return m.whereIn("upt_datetime", values)
}

func (m *VideoQuery) kWheUptDatetimeNotIn(values interface{}) *VideoQuery {
	return m.whereNotIn("upt_datetime", values)
}

func (m *VideoQuery) kWheAddDatetime(args ...interface{}) *VideoQuery {
	return m.where("add_datetime", args...)
}

func (m *VideoQuery) kSetAddDatetime(val interface{}) *VideoQuery {
	return m.Set("add_datetime", val)
}

func (m *VideoQuery) kWheAddDatetimeIn(values interface{}) *VideoQuery {
	return m.whereIn("add_datetime", values)
}

func (m *VideoQuery) kWheAddDatetimeNotIn(values interface{}) *VideoQuery {
	return m.whereNotIn("add_datetime", values)
}

func (m *VideoQuery) kWheVideoId(args ...interface{}) *VideoQuery {
	return m.where("video_id", args...)
}

func (m *VideoQuery) kSetVideoId(val interface{}) *VideoQuery {
	return m.Set("video_id", val)
}

func (m *VideoQuery) kIncVideoId(num int) *VideoQuery {
	return m.Inc("video_id", num)
}

func (m *VideoQuery) kWheVideoIdIn(values interface{}) *VideoQuery {
	return m.whereIn("video_id", values)
}

func (m *VideoQuery) kWheVideoIdNotIn(values interface{}) *VideoQuery {
	return m.whereNotIn("video_id", values)
}

/******************************************* 私有基于字段的build方法 end *****************************************************/
