package tag

// Unscoped 该文件中的所有方法都不应该在model层以外调用，部分共有方法是为了方便model的连表操作;
// must update with github.com.livegoplayer.go_db_helper.mysql 匹配其中的func (m *{name}Query) {funcName} *{name}Query 由于项目初始化之前不一定引入该包，所以采用手写的方式
// 复制当前的内容出一个新的build对象
/****************************************************** 代理方法 start ************************************************************/

func (m *TagQuery) Unscoped() *TagQuery {
	m.GetBuild().Unscoped()
	return m
}

func (m *TagQuery) Select(columns ...interface{}) *TagQuery {
	m.GetBuild().Select(columns...)
	return m
}

func (m *TagQuery) where(column string, args ...interface{}) *TagQuery {
	m.GetBuild().Where(column, args...)
	return m
}

func (m *TagQuery) whereMap(datas map[string]interface{}) *TagQuery {
	m.GetBuild().WhereMap(datas)
	return m
}

func (m *TagQuery) orWhere(column string, args ...interface{}) *TagQuery {
	m.GetBuild().OrWhere(column, args...)
	return m
}

func (m *TagQuery) orWhereCb(cb WhereCb) *TagQuery {
	m.GetBuild().OrWhereCb(cb)
	return m
}

func (m *TagQuery) whereCb(cb WhereCb) *TagQuery {
	m.GetBuild().WhereCb(cb)
	return m
}

func (m *TagQuery) whereRaw(sql string, bindings ...interface{}) *TagQuery {
	m.GetBuild().WhereRaw(sql, bindings...)
	return m
}

func (m *TagQuery) whereNull(column string) *TagQuery {
	m.GetBuild().WhereNull(column)
	return m
}

func (m *TagQuery) whereNotNull(column string) *TagQuery {
	m.GetBuild().WhereNotNull(column)
	return m
}

func (m *TagQuery) orWhereNull(column string) *TagQuery {
	m.GetBuild().OrWhereNull(column)
	return m
}

func (m *TagQuery) orWhereNotNull(column string) *TagQuery {
	m.GetBuild().OrWhereNotNull(column)
	return m
}

func (m *TagQuery) whereIn(column string, values interface{}) *TagQuery {
	m.GetBuild().WhereIn(column, values)
	return m
}

func (m *TagQuery) whereNotIn(column string, values interface{}) *TagQuery {
	m.GetBuild().WhereNotIn(column, values)
	return m
}

func (m *TagQuery) LeftJoin(table string, one string, args ...interface{}) *TagQuery {
	m.GetBuild().LeftJoin(table, one, args...)
	return m
}

func (m *TagQuery) RightJoin(table string, one string, args ...interface{}) *TagQuery {
	m.GetBuild().RightJoin(table, one, args...)
	return m
}

func (m *TagQuery) InnerJoin(table string, one string, args ...interface{}) *TagQuery {
	m.GetBuild().InnerJoin(table, one, args...)
	return m
}

func (m *TagQuery) OrderBy(column string, direction string) *TagQuery {
	m.GetBuild().OrderBy(column, direction)
	return m
}

func (m *TagQuery) OrderDescBy(column string) *TagQuery {
	m.GetBuild().OrderDescBy(column)
	return m
}

func (m *TagQuery) OrderAscBy(column string) *TagQuery {
	m.GetBuild().OrderAscBy(column)
	return m
}

func (m *TagQuery) Group(column string) *TagQuery {
	m.GetBuild().Group(column)
	return m
}

func (m *TagQuery) Inc(col string, num int) *TagQuery {
	m.GetBuild().Inc(col, num)
	return m
}

func (m *TagQuery) Set(col string, val interface{}) *TagQuery {
	m.GetBuild().Set(col, val)
	return m
}

func (m *TagQuery) Skip(lines int) *TagQuery {
	m.GetBuild().Skip(lines)
	return m
}

func (m *TagQuery) Offset(lines int) *TagQuery {
	m.GetBuild().Offset(lines)
	return m
}

func (m *TagQuery) Limit(lines int) *TagQuery {
	m.GetBuild().Limit(lines)
	return m
}

func (m *TagQuery) Take(lines int) *TagQuery {
	m.GetBuild().Take(lines)
	return m
}

func (m *TagQuery) ForUpdate() *TagQuery {
	m.GetBuild().ForUpdate()
	return m
}

func (m *TagQuery) SplitBy(val int64) *TagQuery {
	m.GetBuild().SplitBy(val)
	return m
}

func (m *TagQuery) ModelType(v interface{}) *TagQuery {
	m.GetBuild().ModelType(v)
	return m
}

func (m *TagQuery) DisablePanic() *TagQuery {
	m.GetBuild().DisablePanic()
	return m
}

func (m *TagQuery) OnWrite() *TagQuery {
	m.GetBuild().OnWrite()
	return m
}

/****************************************************** 代理方法 end ************************************************************/
/********************************************** 私有基于字段的build方法 start*****************************************************/

func (m *TagQuery) kWheUptDatetime(args ...interface{}) *TagQuery {
	return m.where("upt_datetime", args...)
}

func (m *TagQuery) kSetUptDatetime(val interface{}) *TagQuery {
	return m.Set("upt_datetime", val)
}

func (m *TagQuery) kIncUptDatetime(num int) *TagQuery {
	return m.Inc("upt_datetime", num)
}

func (m *TagQuery) kWheUptDatetimeIn(values interface{}) *TagQuery {
	return m.whereIn("upt_datetime", values)
}

func (m *TagQuery) kWheUptDatetimeNotIn(values interface{}) *TagQuery {
	return m.whereNotIn("upt_datetime", values)
}

func (m *TagQuery) kWheTagName(args ...interface{}) *TagQuery {
	return m.where("tag_name", args...)
}

func (m *TagQuery) kSetTagName(val interface{}) *TagQuery {
	return m.Set("tag_name", val)
}

func (m *TagQuery) kWheTagNameIn(values interface{}) *TagQuery {
	return m.whereIn("tag_name", values)
}

func (m *TagQuery) kWheTagNameNotIn(values interface{}) *TagQuery {
	return m.whereNotIn("tag_name", values)
}

func (m *TagQuery) kWheTagId(args ...interface{}) *TagQuery {
	return m.where("tag_id", args...)
}

func (m *TagQuery) kSetTagId(val interface{}) *TagQuery {
	return m.Set("tag_id", val)
}

func (m *TagQuery) kIncTagId(num int) *TagQuery {
	return m.Inc("tag_id", num)
}

func (m *TagQuery) kWheTagIdIn(values interface{}) *TagQuery {
	return m.whereIn("tag_id", values)
}

func (m *TagQuery) kWheTagIdNotIn(values interface{}) *TagQuery {
	return m.whereNotIn("tag_id", values)
}

func (m *TagQuery) kWheId(args ...interface{}) *TagQuery {
	return m.where("id", args...)
}

func (m *TagQuery) kSetId(val interface{}) *TagQuery {
	return m.Set("id", val)
}

func (m *TagQuery) kIncId(num int) *TagQuery {
	return m.Inc("id", num)
}

func (m *TagQuery) kWheIdIn(values interface{}) *TagQuery {
	return m.whereIn("id", values)
}

func (m *TagQuery) kWheIdNotIn(values interface{}) *TagQuery {
	return m.whereNotIn("id", values)
}

func (m *TagQuery) kWheTagUrl(args ...interface{}) *TagQuery {
	return m.where("tag_url", args...)
}

func (m *TagQuery) kSetTagUrl(val interface{}) *TagQuery {
	return m.Set("tag_url", val)
}

func (m *TagQuery) kWheTagUrlIn(values interface{}) *TagQuery {
	return m.whereIn("tag_url", values)
}

func (m *TagQuery) kWheTagUrlNotIn(values interface{}) *TagQuery {
	return m.whereNotIn("tag_url", values)
}

func (m *TagQuery) kWheAddDatetime(args ...interface{}) *TagQuery {
	return m.where("add_datetime", args...)
}

func (m *TagQuery) kSetAddDatetime(val interface{}) *TagQuery {
	return m.Set("add_datetime", val)
}

func (m *TagQuery) kIncAddDatetime(num int) *TagQuery {
	return m.Inc("add_datetime", num)
}

func (m *TagQuery) kWheAddDatetimeIn(values interface{}) *TagQuery {
	return m.whereIn("add_datetime", values)
}

func (m *TagQuery) kWheAddDatetimeNotIn(values interface{}) *TagQuery {
	return m.whereNotIn("add_datetime", values)
}

/******************************************* 私有基于字段的build方法 end *****************************************************/
