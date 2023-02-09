package video_detail

// Unscoped 该文件中的所有方法都不应该在model层以外调用，部分共有方法是为了方便model的连表操作;
// must update with github.com.livegoplayer.go_db_helper.mysql 匹配其中的func (m *{name}Query) {funcName} *{name}Query 由于项目初始化之前不一定引入该包，所以采用手写的方式
// 复制当前的内容出一个新的build对象
/****************************************************** 代理方法 start ************************************************************/

func (m *VideoDetailQuery) Unscoped() *VideoDetailQuery {
	m.GetBuild().Unscoped()
	return m
}

func (m *VideoDetailQuery) Select(columns ...interface{}) *VideoDetailQuery {
	m.GetBuild().Select(columns...)
	return m
}

func (m *VideoDetailQuery) where(column string, args ...interface{}) *VideoDetailQuery {
	m.GetBuild().Where(column, args...)
	return m
}

func (m *VideoDetailQuery) whereMap(datas map[string]interface{}) *VideoDetailQuery {
	m.GetBuild().WhereMap(datas)
	return m
}

func (m *VideoDetailQuery) orWhere(column string, args ...interface{}) *VideoDetailQuery {
	m.GetBuild().OrWhere(column, args...)
	return m
}

func (m *VideoDetailQuery) orWhereCb(cb WhereCb) *VideoDetailQuery {
	m.GetBuild().OrWhereCb(cb)
	return m
}

func (m *VideoDetailQuery) whereCb(cb WhereCb) *VideoDetailQuery {
	m.GetBuild().WhereCb(cb)
	return m
}

func (m *VideoDetailQuery) whereRaw(sql string, bindings ...interface{}) *VideoDetailQuery {
	m.GetBuild().WhereRaw(sql, bindings...)
	return m
}

func (m *VideoDetailQuery) whereNull(column string) *VideoDetailQuery {
	m.GetBuild().WhereNull(column)
	return m
}

func (m *VideoDetailQuery) whereNotNull(column string) *VideoDetailQuery {
	m.GetBuild().WhereNotNull(column)
	return m
}

func (m *VideoDetailQuery) orWhereNull(column string) *VideoDetailQuery {
	m.GetBuild().OrWhereNull(column)
	return m
}

func (m *VideoDetailQuery) orWhereNotNull(column string) *VideoDetailQuery {
	m.GetBuild().OrWhereNotNull(column)
	return m
}

func (m *VideoDetailQuery) whereIn(column string, values interface{}) *VideoDetailQuery {
	m.GetBuild().WhereIn(column, values)
	return m
}

func (m *VideoDetailQuery) whereNotIn(column string, values interface{}) *VideoDetailQuery {
	m.GetBuild().WhereNotIn(column, values)
	return m
}

func (m *VideoDetailQuery) LeftJoin(table string, one string, args ...interface{}) *VideoDetailQuery {
	m.GetBuild().LeftJoin(table, one, args...)
	return m
}

func (m *VideoDetailQuery) RightJoin(table string, one string, args ...interface{}) *VideoDetailQuery {
	m.GetBuild().RightJoin(table, one, args...)
	return m
}

func (m *VideoDetailQuery) InnerJoin(table string, one string, args ...interface{}) *VideoDetailQuery {
	m.GetBuild().InnerJoin(table, one, args...)
	return m
}

func (m *VideoDetailQuery) OrderBy(column string, direction string) *VideoDetailQuery {
	m.GetBuild().OrderBy(column, direction)
	return m
}

func (m *VideoDetailQuery) OrderDescBy(column string) *VideoDetailQuery {
	m.GetBuild().OrderDescBy(column)
	return m
}

func (m *VideoDetailQuery) OrderAscBy(column string) *VideoDetailQuery {
	m.GetBuild().OrderAscBy(column)
	return m
}

func (m *VideoDetailQuery) Group(column string) *VideoDetailQuery {
	m.GetBuild().Group(column)
	return m
}

func (m *VideoDetailQuery) Inc(col string, num int) *VideoDetailQuery {
	m.GetBuild().Inc(col, num)
	return m
}

func (m *VideoDetailQuery) Set(col string, val interface{}) *VideoDetailQuery {
	m.GetBuild().Set(col, val)
	return m
}

func (m *VideoDetailQuery) Skip(lines int) *VideoDetailQuery {
	m.GetBuild().Skip(lines)
	return m
}

func (m *VideoDetailQuery) Offset(lines int) *VideoDetailQuery {
	m.GetBuild().Offset(lines)
	return m
}

func (m *VideoDetailQuery) Limit(lines int) *VideoDetailQuery {
	m.GetBuild().Limit(lines)
	return m
}

func (m *VideoDetailQuery) Take(lines int) *VideoDetailQuery {
	m.GetBuild().Take(lines)
	return m
}

func (m *VideoDetailQuery) ForUpdate() *VideoDetailQuery {
	m.GetBuild().ForUpdate()
	return m
}

func (m *VideoDetailQuery) SplitBy(val int64) *VideoDetailQuery {
	m.GetBuild().SplitBy(val)
	return m
}

func (m *VideoDetailQuery) ModelType(v interface{}) *VideoDetailQuery {
	m.GetBuild().ModelType(v)
	return m
}

func (m *VideoDetailQuery) DisablePanic() *VideoDetailQuery {
	m.GetBuild().DisablePanic()
	return m
}

func (m *VideoDetailQuery) OnWrite() *VideoDetailQuery {
	m.GetBuild().OnWrite()
	return m
}

/****************************************************** 代理方法 end ************************************************************/
/********************************************** 私有基于字段的build方法 start*****************************************************/

func (m *VideoDetailQuery) kWheOfficialWebsite(args ...interface{}) *VideoDetailQuery {
	return m.where("official_website", args...)
}

func (m *VideoDetailQuery) kSetOfficialWebsite(val interface{}) *VideoDetailQuery {
	return m.Set("official_website", val)
}

func (m *VideoDetailQuery) kWheOfficialWebsiteIn(values interface{}) *VideoDetailQuery {
	return m.whereIn("official_website", values)
}

func (m *VideoDetailQuery) kWheOfficialWebsiteNotIn(values interface{}) *VideoDetailQuery {
	return m.whereNotIn("official_website", values)
}

func (m *VideoDetailQuery) kWheDescriptionUrl(args ...interface{}) *VideoDetailQuery {
	return m.where("description_url", args...)
}

func (m *VideoDetailQuery) kSetDescriptionUrl(val interface{}) *VideoDetailQuery {
	return m.Set("description_url", val)
}

func (m *VideoDetailQuery) kWheDescriptionUrlIn(values interface{}) *VideoDetailQuery {
	return m.whereIn("description_url", values)
}

func (m *VideoDetailQuery) kWheDescriptionUrlNotIn(values interface{}) *VideoDetailQuery {
	return m.whereNotIn("description_url", values)
}

func (m *VideoDetailQuery) kWheYear(args ...interface{}) *VideoDetailQuery {
	return m.where("year", args...)
}

func (m *VideoDetailQuery) kSetYear(val interface{}) *VideoDetailQuery {
	return m.Set("year", val)
}

func (m *VideoDetailQuery) kIncYear(num int) *VideoDetailQuery {
	return m.Inc("year", num)
}

func (m *VideoDetailQuery) kWheYearIn(values interface{}) *VideoDetailQuery {
	return m.whereIn("year", values)
}

func (m *VideoDetailQuery) kWheYearNotIn(values interface{}) *VideoDetailQuery {
	return m.whereNotIn("year", values)
}

func (m *VideoDetailQuery) kWheReleasedDateStr(args ...interface{}) *VideoDetailQuery {
	return m.where("released_date_str", args...)
}

func (m *VideoDetailQuery) kSetReleasedDateStr(val interface{}) *VideoDetailQuery {
	return m.Set("released_date_str", val)
}

func (m *VideoDetailQuery) kWheReleasedDateStrIn(values interface{}) *VideoDetailQuery {
	return m.whereIn("released_date_str", values)
}

func (m *VideoDetailQuery) kWheReleasedDateStrNotIn(values interface{}) *VideoDetailQuery {
	return m.whereNotIn("released_date_str", values)
}

func (m *VideoDetailQuery) kWheVideoPosterUrl(args ...interface{}) *VideoDetailQuery {
	return m.where("video_poster_url", args...)
}

func (m *VideoDetailQuery) kSetVideoPosterUrl(val interface{}) *VideoDetailQuery {
	return m.Set("video_poster_url", val)
}

func (m *VideoDetailQuery) kWheVideoPosterUrlIn(values interface{}) *VideoDetailQuery {
	return m.whereIn("video_poster_url", values)
}

func (m *VideoDetailQuery) kWheVideoPosterUrlNotIn(values interface{}) *VideoDetailQuery {
	return m.whereNotIn("video_poster_url", values)
}

func (m *VideoDetailQuery) kWheTimeStr(args ...interface{}) *VideoDetailQuery {
	return m.where("time_str", args...)
}

func (m *VideoDetailQuery) kSetTimeStr(val interface{}) *VideoDetailQuery {
	return m.Set("time_str", val)
}

func (m *VideoDetailQuery) kWheTimeStrIn(values interface{}) *VideoDetailQuery {
	return m.whereIn("time_str", values)
}

func (m *VideoDetailQuery) kWheTimeStrNotIn(values interface{}) *VideoDetailQuery {
	return m.whereNotIn("time_str", values)
}

func (m *VideoDetailQuery) kWhePosterCache(args ...interface{}) *VideoDetailQuery {
	return m.where("poster_cache", args...)
}

func (m *VideoDetailQuery) kSetPosterCache(val interface{}) *VideoDetailQuery {
	return m.Set("poster_cache", val)
}

func (m *VideoDetailQuery) kWhePosterCacheIn(values interface{}) *VideoDetailQuery {
	return m.whereIn("poster_cache", values)
}

func (m *VideoDetailQuery) kWhePosterCacheNotIn(values interface{}) *VideoDetailQuery {
	return m.whereNotIn("poster_cache", values)
}

func (m *VideoDetailQuery) kWheVideoDescription(args ...interface{}) *VideoDetailQuery {
	return m.where("video_description", args...)
}

func (m *VideoDetailQuery) kSetVideoDescription(val interface{}) *VideoDetailQuery {
	return m.Set("video_description", val)
}

func (m *VideoDetailQuery) kWheVideoDescriptionIn(values interface{}) *VideoDetailQuery {
	return m.whereIn("video_description", values)
}

func (m *VideoDetailQuery) kWheVideoDescriptionNotIn(values interface{}) *VideoDetailQuery {
	return m.whereNotIn("video_description", values)
}

func (m *VideoDetailQuery) kWheScore(args ...interface{}) *VideoDetailQuery {
	return m.where("score", args...)
}

func (m *VideoDetailQuery) kSetScore(val interface{}) *VideoDetailQuery {
	return m.Set("score", val)
}

func (m *VideoDetailQuery) kWheScoreIn(values interface{}) *VideoDetailQuery {
	return m.whereIn("score", values)
}

func (m *VideoDetailQuery) kWheScoreNotIn(values interface{}) *VideoDetailQuery {
	return m.whereNotIn("score", values)
}

func (m *VideoDetailQuery) kWhePosterListUrl(args ...interface{}) *VideoDetailQuery {
	return m.where("poster_list_url", args...)
}

func (m *VideoDetailQuery) kSetPosterListUrl(val interface{}) *VideoDetailQuery {
	return m.Set("poster_list_url", val)
}

func (m *VideoDetailQuery) kWhePosterListUrlIn(values interface{}) *VideoDetailQuery {
	return m.whereIn("poster_list_url", values)
}

func (m *VideoDetailQuery) kWhePosterListUrlNotIn(values interface{}) *VideoDetailQuery {
	return m.whereNotIn("poster_list_url", values)
}

func (m *VideoDetailQuery) kWheUptDatetime(args ...interface{}) *VideoDetailQuery {
	return m.where("upt_datetime", args...)
}

func (m *VideoDetailQuery) kSetUptDatetime(val interface{}) *VideoDetailQuery {
	return m.Set("upt_datetime", val)
}

func (m *VideoDetailQuery) kIncUptDatetime(num int) *VideoDetailQuery {
	return m.Inc("upt_datetime", num)
}

func (m *VideoDetailQuery) kWheUptDatetimeIn(values interface{}) *VideoDetailQuery {
	return m.whereIn("upt_datetime", values)
}

func (m *VideoDetailQuery) kWheUptDatetimeNotIn(values interface{}) *VideoDetailQuery {
	return m.whereNotIn("upt_datetime", values)
}

func (m *VideoDetailQuery) kWhePerUpdateTime(args ...interface{}) *VideoDetailQuery {
	return m.where("per_update_time", args...)
}

func (m *VideoDetailQuery) kSetPerUpdateTime(val interface{}) *VideoDetailQuery {
	return m.Set("per_update_time", val)
}

func (m *VideoDetailQuery) kWhePerUpdateTimeIn(values interface{}) *VideoDetailQuery {
	return m.whereIn("per_update_time", values)
}

func (m *VideoDetailQuery) kWhePerUpdateTimeNotIn(values interface{}) *VideoDetailQuery {
	return m.whereNotIn("per_update_time", values)
}

func (m *VideoDetailQuery) kWheSeason(args ...interface{}) *VideoDetailQuery {
	return m.where("season", args...)
}

func (m *VideoDetailQuery) kSetSeason(val interface{}) *VideoDetailQuery {
	return m.Set("season", val)
}

func (m *VideoDetailQuery) kIncSeason(num int) *VideoDetailQuery {
	return m.Inc("season", num)
}

func (m *VideoDetailQuery) kWheSeasonIn(values interface{}) *VideoDetailQuery {
	return m.whereIn("season", values)
}

func (m *VideoDetailQuery) kWheSeasonNotIn(values interface{}) *VideoDetailQuery {
	return m.whereNotIn("season", values)
}

func (m *VideoDetailQuery) kWheId(args ...interface{}) *VideoDetailQuery {
	return m.where("id", args...)
}

func (m *VideoDetailQuery) kSetId(val interface{}) *VideoDetailQuery {
	return m.Set("id", val)
}

func (m *VideoDetailQuery) kIncId(num int) *VideoDetailQuery {
	return m.Inc("id", num)
}

func (m *VideoDetailQuery) kWheIdIn(values interface{}) *VideoDetailQuery {
	return m.whereIn("id", values)
}

func (m *VideoDetailQuery) kWheIdNotIn(values interface{}) *VideoDetailQuery {
	return m.whereNotIn("id", values)
}

func (m *VideoDetailQuery) kWheSetNum(args ...interface{}) *VideoDetailQuery {
	return m.where("set_num", args...)
}

func (m *VideoDetailQuery) kSetSetNum(val interface{}) *VideoDetailQuery {
	return m.Set("set_num", val)
}

func (m *VideoDetailQuery) kIncSetNum(num int) *VideoDetailQuery {
	return m.Inc("set_num", num)
}

func (m *VideoDetailQuery) kWheSetNumIn(values interface{}) *VideoDetailQuery {
	return m.whereIn("set_num", values)
}

func (m *VideoDetailQuery) kWheSetNumNotIn(values interface{}) *VideoDetailQuery {
	return m.whereNotIn("set_num", values)
}

func (m *VideoDetailQuery) kWheAddDatetime(args ...interface{}) *VideoDetailQuery {
	return m.where("add_datetime", args...)
}

func (m *VideoDetailQuery) kSetAddDatetime(val interface{}) *VideoDetailQuery {
	return m.Set("add_datetime", val)
}

func (m *VideoDetailQuery) kIncAddDatetime(num int) *VideoDetailQuery {
	return m.Inc("add_datetime", num)
}

func (m *VideoDetailQuery) kWheAddDatetimeIn(values interface{}) *VideoDetailQuery {
	return m.whereIn("add_datetime", values)
}

func (m *VideoDetailQuery) kWheAddDatetimeNotIn(values interface{}) *VideoDetailQuery {
	return m.whereNotIn("add_datetime", values)
}

func (m *VideoDetailQuery) kWheVideoId(args ...interface{}) *VideoDetailQuery {
	return m.where("video_id", args...)
}

func (m *VideoDetailQuery) kSetVideoId(val interface{}) *VideoDetailQuery {
	return m.Set("video_id", val)
}

func (m *VideoDetailQuery) kIncVideoId(num int) *VideoDetailQuery {
	return m.Inc("video_id", num)
}

func (m *VideoDetailQuery) kWheVideoIdIn(values interface{}) *VideoDetailQuery {
	return m.whereIn("video_id", values)
}

func (m *VideoDetailQuery) kWheVideoIdNotIn(values interface{}) *VideoDetailQuery {
	return m.whereNotIn("video_id", values)
}

/******************************************* 私有基于字段的build方法 end *****************************************************/
