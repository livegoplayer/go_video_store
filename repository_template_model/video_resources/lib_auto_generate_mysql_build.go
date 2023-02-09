package video_resources

// Unscoped 该文件中的所有方法都不应该在model层以外调用，部分共有方法是为了方便model的连表操作;
// must update with github.com.livegoplayer.go_db_helper.mysql 匹配其中的func (m *{name}Query) {funcName} *{name}Query 由于项目初始化之前不一定引入该包，所以采用手写的方式
// 复制当前的内容出一个新的build对象
/****************************************************** 代理方法 start ************************************************************/

func (m *VideoResourcesQuery) Unscoped() *VideoResourcesQuery {
	m.GetBuild().Unscoped()
	return m
}

func (m *VideoResourcesQuery) Select(columns ...interface{}) *VideoResourcesQuery {
	m.GetBuild().Select(columns...)
	return m
}

func (m *VideoResourcesQuery) where(column string, args ...interface{}) *VideoResourcesQuery {
	m.GetBuild().Where(column, args...)
	return m
}

func (m *VideoResourcesQuery) whereMap(datas map[string]interface{}) *VideoResourcesQuery {
	m.GetBuild().WhereMap(datas)
	return m
}

func (m *VideoResourcesQuery) orWhere(column string, args ...interface{}) *VideoResourcesQuery {
	m.GetBuild().OrWhere(column, args...)
	return m
}

func (m *VideoResourcesQuery) orWhereCb(cb WhereCb) *VideoResourcesQuery {
	m.GetBuild().OrWhereCb(cb)
	return m
}

func (m *VideoResourcesQuery) whereCb(cb WhereCb) *VideoResourcesQuery {
	m.GetBuild().WhereCb(cb)
	return m
}

func (m *VideoResourcesQuery) whereRaw(sql string, bindings ...interface{}) *VideoResourcesQuery {
	m.GetBuild().WhereRaw(sql, bindings...)
	return m
}

func (m *VideoResourcesQuery) whereNull(column string) *VideoResourcesQuery {
	m.GetBuild().WhereNull(column)
	return m
}

func (m *VideoResourcesQuery) whereNotNull(column string) *VideoResourcesQuery {
	m.GetBuild().WhereNotNull(column)
	return m
}

func (m *VideoResourcesQuery) orWhereNull(column string) *VideoResourcesQuery {
	m.GetBuild().OrWhereNull(column)
	return m
}

func (m *VideoResourcesQuery) orWhereNotNull(column string) *VideoResourcesQuery {
	m.GetBuild().OrWhereNotNull(column)
	return m
}

func (m *VideoResourcesQuery) whereIn(column string, values interface{}) *VideoResourcesQuery {
	m.GetBuild().WhereIn(column, values)
	return m
}

func (m *VideoResourcesQuery) whereNotIn(column string, values interface{}) *VideoResourcesQuery {
	m.GetBuild().WhereNotIn(column, values)
	return m
}

func (m *VideoResourcesQuery) LeftJoin(table string, one string, args ...interface{}) *VideoResourcesQuery {
	m.GetBuild().LeftJoin(table, one, args...)
	return m
}

func (m *VideoResourcesQuery) RightJoin(table string, one string, args ...interface{}) *VideoResourcesQuery {
	m.GetBuild().RightJoin(table, one, args...)
	return m
}

func (m *VideoResourcesQuery) InnerJoin(table string, one string, args ...interface{}) *VideoResourcesQuery {
	m.GetBuild().InnerJoin(table, one, args...)
	return m
}

func (m *VideoResourcesQuery) OrderBy(column string, direction string) *VideoResourcesQuery {
	m.GetBuild().OrderBy(column, direction)
	return m
}

func (m *VideoResourcesQuery) OrderDescBy(column string) *VideoResourcesQuery {
	m.GetBuild().OrderDescBy(column)
	return m
}

func (m *VideoResourcesQuery) OrderAscBy(column string) *VideoResourcesQuery {
	m.GetBuild().OrderAscBy(column)
	return m
}

func (m *VideoResourcesQuery) Group(column string) *VideoResourcesQuery {
	m.GetBuild().Group(column)
	return m
}

func (m *VideoResourcesQuery) Inc(col string, num int) *VideoResourcesQuery {
	m.GetBuild().Inc(col, num)
	return m
}

func (m *VideoResourcesQuery) Set(col string, val interface{}) *VideoResourcesQuery {
	m.GetBuild().Set(col, val)
	return m
}

func (m *VideoResourcesQuery) Skip(lines int) *VideoResourcesQuery {
	m.GetBuild().Skip(lines)
	return m
}

func (m *VideoResourcesQuery) Offset(lines int) *VideoResourcesQuery {
	m.GetBuild().Offset(lines)
	return m
}

func (m *VideoResourcesQuery) Limit(lines int) *VideoResourcesQuery {
	m.GetBuild().Limit(lines)
	return m
}

func (m *VideoResourcesQuery) Take(lines int) *VideoResourcesQuery {
	m.GetBuild().Take(lines)
	return m
}

func (m *VideoResourcesQuery) ForUpdate() *VideoResourcesQuery {
	m.GetBuild().ForUpdate()
	return m
}

func (m *VideoResourcesQuery) SplitBy(val int64) *VideoResourcesQuery {
	m.GetBuild().SplitBy(val)
	return m
}

func (m *VideoResourcesQuery) ModelType(v interface{}) *VideoResourcesQuery {
	m.GetBuild().ModelType(v)
	return m
}

func (m *VideoResourcesQuery) DisablePanic() *VideoResourcesQuery {
	m.GetBuild().DisablePanic()
	return m
}

func (m *VideoResourcesQuery) OnWrite() *VideoResourcesQuery {
	m.GetBuild().OnWrite()
	return m
}

/****************************************************** 代理方法 end ************************************************************/
/********************************************** 私有基于字段的build方法 start*****************************************************/

func (m *VideoResourcesQuery) kWheLastUpdateTime(args ...interface{}) *VideoResourcesQuery {
	return m.where("last_update_time", args...)
}

func (m *VideoResourcesQuery) kSetLastUpdateTime(val interface{}) *VideoResourcesQuery {
	return m.Set("last_update_time", val)
}

func (m *VideoResourcesQuery) kIncLastUpdateTime(num int) *VideoResourcesQuery {
	return m.Inc("last_update_time", num)
}

func (m *VideoResourcesQuery) kWheLastUpdateTimeIn(values interface{}) *VideoResourcesQuery {
	return m.whereIn("last_update_time", values)
}

func (m *VideoResourcesQuery) kWheLastUpdateTimeNotIn(values interface{}) *VideoResourcesQuery {
	return m.whereNotIn("last_update_time", values)
}

func (m *VideoResourcesQuery) kWheVideoUrl(args ...interface{}) *VideoResourcesQuery {
	return m.where("video_url", args...)
}

func (m *VideoResourcesQuery) kSetVideoUrl(val interface{}) *VideoResourcesQuery {
	return m.Set("video_url", val)
}

func (m *VideoResourcesQuery) kWheVideoUrlIn(values interface{}) *VideoResourcesQuery {
	return m.whereIn("video_url", values)
}

func (m *VideoResourcesQuery) kWheVideoUrlNotIn(values interface{}) *VideoResourcesQuery {
	return m.whereNotIn("video_url", values)
}

func (m *VideoResourcesQuery) kWheData(args ...interface{}) *VideoResourcesQuery {
	return m.where("data", args...)
}

func (m *VideoResourcesQuery) kSetData(val interface{}) *VideoResourcesQuery {
	return m.Set("data", val)
}

func (m *VideoResourcesQuery) kWheDataIn(values interface{}) *VideoResourcesQuery {
	return m.whereIn("data", values)
}

func (m *VideoResourcesQuery) kWheDataNotIn(values interface{}) *VideoResourcesQuery {
	return m.whereNotIn("data", values)
}

func (m *VideoResourcesQuery) kWheVideoReferDomain(args ...interface{}) *VideoResourcesQuery {
	return m.where("video_refer_domain", args...)
}

func (m *VideoResourcesQuery) kSetVideoReferDomain(val interface{}) *VideoResourcesQuery {
	return m.Set("video_refer_domain", val)
}

func (m *VideoResourcesQuery) kWheVideoReferDomainIn(values interface{}) *VideoResourcesQuery {
	return m.whereIn("video_refer_domain", values)
}

func (m *VideoResourcesQuery) kWheVideoReferDomainNotIn(values interface{}) *VideoResourcesQuery {
	return m.whereNotIn("video_refer_domain", values)
}

func (m *VideoResourcesQuery) kWheUptDatetime(args ...interface{}) *VideoResourcesQuery {
	return m.where("upt_datetime", args...)
}

func (m *VideoResourcesQuery) kSetUptDatetime(val interface{}) *VideoResourcesQuery {
	return m.Set("upt_datetime", val)
}

func (m *VideoResourcesQuery) kIncUptDatetime(num int) *VideoResourcesQuery {
	return m.Inc("upt_datetime", num)
}

func (m *VideoResourcesQuery) kWheUptDatetimeIn(values interface{}) *VideoResourcesQuery {
	return m.whereIn("upt_datetime", values)
}

func (m *VideoResourcesQuery) kWheUptDatetimeNotIn(values interface{}) *VideoResourcesQuery {
	return m.whereNotIn("upt_datetime", values)
}

func (m *VideoResourcesQuery) kWheName(args ...interface{}) *VideoResourcesQuery {
	return m.where("name", args...)
}

func (m *VideoResourcesQuery) kSetName(val interface{}) *VideoResourcesQuery {
	return m.Set("name", val)
}

func (m *VideoResourcesQuery) kWheNameIn(values interface{}) *VideoResourcesQuery {
	return m.whereIn("name", values)
}

func (m *VideoResourcesQuery) kWheNameNotIn(values interface{}) *VideoResourcesQuery {
	return m.whereNotIn("name", values)
}

func (m *VideoResourcesQuery) kWheSeason(args ...interface{}) *VideoResourcesQuery {
	return m.where("season", args...)
}

func (m *VideoResourcesQuery) kSetSeason(val interface{}) *VideoResourcesQuery {
	return m.Set("season", val)
}

func (m *VideoResourcesQuery) kIncSeason(num int) *VideoResourcesQuery {
	return m.Inc("season", num)
}

func (m *VideoResourcesQuery) kWheSeasonIn(values interface{}) *VideoResourcesQuery {
	return m.whereIn("season", values)
}

func (m *VideoResourcesQuery) kWheSeasonNotIn(values interface{}) *VideoResourcesQuery {
	return m.whereNotIn("season", values)
}

func (m *VideoResourcesQuery) kWheEpisode(args ...interface{}) *VideoResourcesQuery {
	return m.where("episode", args...)
}

func (m *VideoResourcesQuery) kSetEpisode(val interface{}) *VideoResourcesQuery {
	return m.Set("episode", val)
}

func (m *VideoResourcesQuery) kIncEpisode(num int) *VideoResourcesQuery {
	return m.Inc("episode", num)
}

func (m *VideoResourcesQuery) kWheEpisodeIn(values interface{}) *VideoResourcesQuery {
	return m.whereIn("episode", values)
}

func (m *VideoResourcesQuery) kWheEpisodeNotIn(values interface{}) *VideoResourcesQuery {
	return m.whereNotIn("episode", values)
}

func (m *VideoResourcesQuery) kWheId(args ...interface{}) *VideoResourcesQuery {
	return m.where("id", args...)
}

func (m *VideoResourcesQuery) kSetId(val interface{}) *VideoResourcesQuery {
	return m.Set("id", val)
}

func (m *VideoResourcesQuery) kIncId(num int) *VideoResourcesQuery {
	return m.Inc("id", num)
}

func (m *VideoResourcesQuery) kWheIdIn(values interface{}) *VideoResourcesQuery {
	return m.whereIn("id", values)
}

func (m *VideoResourcesQuery) kWheIdNotIn(values interface{}) *VideoResourcesQuery {
	return m.whereNotIn("id", values)
}

func (m *VideoResourcesQuery) kWheVideoRefer(args ...interface{}) *VideoResourcesQuery {
	return m.where("video_refer", args...)
}

func (m *VideoResourcesQuery) kSetVideoRefer(val interface{}) *VideoResourcesQuery {
	return m.Set("video_refer", val)
}

func (m *VideoResourcesQuery) kIncVideoRefer(num int) *VideoResourcesQuery {
	return m.Inc("video_refer", num)
}

func (m *VideoResourcesQuery) kWheVideoReferIn(values interface{}) *VideoResourcesQuery {
	return m.whereIn("video_refer", values)
}

func (m *VideoResourcesQuery) kWheVideoReferNotIn(values interface{}) *VideoResourcesQuery {
	return m.whereNotIn("video_refer", values)
}

func (m *VideoResourcesQuery) kWheAddDatetime(args ...interface{}) *VideoResourcesQuery {
	return m.where("add_datetime", args...)
}

func (m *VideoResourcesQuery) kSetAddDatetime(val interface{}) *VideoResourcesQuery {
	return m.Set("add_datetime", val)
}

func (m *VideoResourcesQuery) kIncAddDatetime(num int) *VideoResourcesQuery {
	return m.Inc("add_datetime", num)
}

func (m *VideoResourcesQuery) kWheAddDatetimeIn(values interface{}) *VideoResourcesQuery {
	return m.whereIn("add_datetime", values)
}

func (m *VideoResourcesQuery) kWheAddDatetimeNotIn(values interface{}) *VideoResourcesQuery {
	return m.whereNotIn("add_datetime", values)
}

func (m *VideoResourcesQuery) kWheVideoId(args ...interface{}) *VideoResourcesQuery {
	return m.where("video_id", args...)
}

func (m *VideoResourcesQuery) kSetVideoId(val interface{}) *VideoResourcesQuery {
	return m.Set("video_id", val)
}

func (m *VideoResourcesQuery) kIncVideoId(num int) *VideoResourcesQuery {
	return m.Inc("video_id", num)
}

func (m *VideoResourcesQuery) kWheVideoIdIn(values interface{}) *VideoResourcesQuery {
	return m.whereIn("video_id", values)
}

func (m *VideoResourcesQuery) kWheVideoIdNotIn(values interface{}) *VideoResourcesQuery {
	return m.whereNotIn("video_id", values)
}

/******************************************* 私有基于字段的build方法 end *****************************************************/
