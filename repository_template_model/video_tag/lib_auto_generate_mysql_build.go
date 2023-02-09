package video_tag

// Unscoped 该文件中的所有方法都不应该在model层以外调用，部分共有方法是为了方便model的连表操作;
// must update with github.com.livegoplayer.go_db_helper.mysql 匹配其中的func (m *{name}Query) {funcName} *{name}Query 由于项目初始化之前不一定引入该包，所以采用手写的方式
// 复制当前的内容出一个新的build对象
/****************************************************** 代理方法 start ************************************************************/

func (m *VideoTagQuery) Unscoped() *VideoTagQuery {
	m.GetBuild().Unscoped()
	return m
}

func (m *VideoTagQuery) Select(columns ...interface{}) *VideoTagQuery {
	m.GetBuild().Select(columns...)
	return m
}

func (m *VideoTagQuery) where(column string, args ...interface{}) *VideoTagQuery {
	m.GetBuild().Where(column, args...)
	return m
}

func (m *VideoTagQuery) whereMap(datas map[string]interface{}) *VideoTagQuery {
	m.GetBuild().WhereMap(datas)
	return m
}

func (m *VideoTagQuery) orWhere(column string, args ...interface{}) *VideoTagQuery {
	m.GetBuild().OrWhere(column, args...)
	return m
}

func (m *VideoTagQuery) orWhereCb(cb WhereCb) *VideoTagQuery {
	m.GetBuild().OrWhereCb(cb)
	return m
}

func (m *VideoTagQuery) whereCb(cb WhereCb) *VideoTagQuery {
	m.GetBuild().WhereCb(cb)
	return m
}

func (m *VideoTagQuery) whereRaw(sql string, bindings ...interface{}) *VideoTagQuery {
	m.GetBuild().WhereRaw(sql, bindings...)
	return m
}

func (m *VideoTagQuery) whereNull(column string) *VideoTagQuery {
	m.GetBuild().WhereNull(column)
	return m
}

func (m *VideoTagQuery) whereNotNull(column string) *VideoTagQuery {
	m.GetBuild().WhereNotNull(column)
	return m
}

func (m *VideoTagQuery) orWhereNull(column string) *VideoTagQuery {
	m.GetBuild().OrWhereNull(column)
	return m
}

func (m *VideoTagQuery) orWhereNotNull(column string) *VideoTagQuery {
	m.GetBuild().OrWhereNotNull(column)
	return m
}

func (m *VideoTagQuery) whereIn(column string, values interface{}) *VideoTagQuery {
	m.GetBuild().WhereIn(column, values)
	return m
}

func (m *VideoTagQuery) whereNotIn(column string, values interface{}) *VideoTagQuery {
	m.GetBuild().WhereNotIn(column, values)
	return m
}

func (m *VideoTagQuery) LeftJoin(table string, one string, args ...interface{}) *VideoTagQuery {
	m.GetBuild().LeftJoin(table, one, args...)
	return m
}

func (m *VideoTagQuery) RightJoin(table string, one string, args ...interface{}) *VideoTagQuery {
	m.GetBuild().RightJoin(table, one, args...)
	return m
}

func (m *VideoTagQuery) InnerJoin(table string, one string, args ...interface{}) *VideoTagQuery {
	m.GetBuild().InnerJoin(table, one, args...)
	return m
}

func (m *VideoTagQuery) OrderBy(column string, direction string) *VideoTagQuery {
	m.GetBuild().OrderBy(column, direction)
	return m
}

func (m *VideoTagQuery) OrderDescBy(column string) *VideoTagQuery {
	m.GetBuild().OrderDescBy(column)
	return m
}

func (m *VideoTagQuery) OrderAscBy(column string) *VideoTagQuery {
	m.GetBuild().OrderAscBy(column)
	return m
}

func (m *VideoTagQuery) Group(column string) *VideoTagQuery {
	m.GetBuild().Group(column)
	return m
}

func (m *VideoTagQuery) Inc(col string, num int) *VideoTagQuery {
	m.GetBuild().Inc(col, num)
	return m
}

func (m *VideoTagQuery) Set(col string, val interface{}) *VideoTagQuery {
	m.GetBuild().Set(col, val)
	return m
}

func (m *VideoTagQuery) Skip(lines int) *VideoTagQuery {
	m.GetBuild().Skip(lines)
	return m
}

func (m *VideoTagQuery) Offset(lines int) *VideoTagQuery {
	m.GetBuild().Offset(lines)
	return m
}

func (m *VideoTagQuery) Limit(lines int) *VideoTagQuery {
	m.GetBuild().Limit(lines)
	return m
}

func (m *VideoTagQuery) Take(lines int) *VideoTagQuery {
	m.GetBuild().Take(lines)
	return m
}

func (m *VideoTagQuery) ForUpdate() *VideoTagQuery {
	m.GetBuild().ForUpdate()
	return m
}

func (m *VideoTagQuery) SplitBy(val int64) *VideoTagQuery {
	m.GetBuild().SplitBy(val)
	return m
}

func (m *VideoTagQuery) ModelType(v interface{}) *VideoTagQuery {
	m.GetBuild().ModelType(v)
	return m
}

func (m *VideoTagQuery) DisablePanic() *VideoTagQuery {
	m.GetBuild().DisablePanic()
	return m
}

func (m *VideoTagQuery) OnWrite() *VideoTagQuery {
	m.GetBuild().OnWrite()
	return m
}

/****************************************************** 代理方法 end ************************************************************/
/********************************************** 私有基于字段的build方法 start*****************************************************/

func (m *VideoTagQuery) kWheUptDatetime(args ...interface{}) *VideoTagQuery {
	return m.where("upt_datetime", args...)
}

func (m *VideoTagQuery) kSetUptDatetime(val interface{}) *VideoTagQuery {
	return m.Set("upt_datetime", val)
}

func (m *VideoTagQuery) kIncUptDatetime(num int) *VideoTagQuery {
	return m.Inc("upt_datetime", num)
}

func (m *VideoTagQuery) kWheUptDatetimeIn(values interface{}) *VideoTagQuery {
	return m.whereIn("upt_datetime", values)
}

func (m *VideoTagQuery) kWheUptDatetimeNotIn(values interface{}) *VideoTagQuery {
	return m.whereNotIn("upt_datetime", values)
}

func (m *VideoTagQuery) kWheTagId(args ...interface{}) *VideoTagQuery {
	return m.where("tag_id", args...)
}

func (m *VideoTagQuery) kSetTagId(val interface{}) *VideoTagQuery {
	return m.Set("tag_id", val)
}

func (m *VideoTagQuery) kIncTagId(num int) *VideoTagQuery {
	return m.Inc("tag_id", num)
}

func (m *VideoTagQuery) kWheTagIdIn(values interface{}) *VideoTagQuery {
	return m.whereIn("tag_id", values)
}

func (m *VideoTagQuery) kWheTagIdNotIn(values interface{}) *VideoTagQuery {
	return m.whereNotIn("tag_id", values)
}

func (m *VideoTagQuery) kWheId(args ...interface{}) *VideoTagQuery {
	return m.where("id", args...)
}

func (m *VideoTagQuery) kSetId(val interface{}) *VideoTagQuery {
	return m.Set("id", val)
}

func (m *VideoTagQuery) kIncId(num int) *VideoTagQuery {
	return m.Inc("id", num)
}

func (m *VideoTagQuery) kWheIdIn(values interface{}) *VideoTagQuery {
	return m.whereIn("id", values)
}

func (m *VideoTagQuery) kWheIdNotIn(values interface{}) *VideoTagQuery {
	return m.whereNotIn("id", values)
}

func (m *VideoTagQuery) kWheAddDatetime(args ...interface{}) *VideoTagQuery {
	return m.where("add_datetime", args...)
}

func (m *VideoTagQuery) kSetAddDatetime(val interface{}) *VideoTagQuery {
	return m.Set("add_datetime", val)
}

func (m *VideoTagQuery) kIncAddDatetime(num int) *VideoTagQuery {
	return m.Inc("add_datetime", num)
}

func (m *VideoTagQuery) kWheAddDatetimeIn(values interface{}) *VideoTagQuery {
	return m.whereIn("add_datetime", values)
}

func (m *VideoTagQuery) kWheAddDatetimeNotIn(values interface{}) *VideoTagQuery {
	return m.whereNotIn("add_datetime", values)
}

func (m *VideoTagQuery) kWheVideoId(args ...interface{}) *VideoTagQuery {
	return m.where("video_id", args...)
}

func (m *VideoTagQuery) kSetVideoId(val interface{}) *VideoTagQuery {
	return m.Set("video_id", val)
}

func (m *VideoTagQuery) kIncVideoId(num int) *VideoTagQuery {
	return m.Inc("video_id", num)
}

func (m *VideoTagQuery) kWheVideoIdIn(values interface{}) *VideoTagQuery {
	return m.whereIn("video_id", values)
}

func (m *VideoTagQuery) kWheVideoIdNotIn(values interface{}) *VideoTagQuery {
	return m.whereNotIn("video_id", values)
}

/******************************************* 私有基于字段的build方法 end *****************************************************/
