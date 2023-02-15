package log

/* Unscoped *************************************** 满足常用的共有方法,根据mysql索引生成 **********************************************/
/**************************************** all 不是小表不要用 **************************************************/

func Save(f *Log) *Log {
	NewLogQuery().save(f)
	return f
}

func UpdateLogAll(p *Log) int64 {
	build := NewLogQuery()
	return build.update(p)
}

func FetchLogAll(order ...string) LogCollect {
	build := NewLogQuery()
	build.OrderBy(parseOrderString(order))
	return build.Get()
}

func FetchLogAllWithPageSize(page int, pageSize int, order ...string) LogCollect {
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * pageSize

	build := NewLogQuery()
	build.OrderBy(parseOrderString(order))
	return build.Skip(offset).Limit(pageSize).Get()
}

func CountLogAll() int64 {
	build := NewLogQuery()
	return build.Count()
}

/**************************************** 根据 Index 生成的方法 **************************************************/

func CheckExistByLevel(level int64) bool {
	build := NewLogQuery()
	build.kWheLevel(level)
	cnt := build.Count()
	return cnt > 0
}

func GetFirstByLevel(level int64) *Log {
	build := NewLogQuery()
	build.kWheLevel(level)
	return build.First()
}

func DeleteAllByLevel(level int64) int64 {
	build := NewLogQuery()
	build.kWheLevel(level)
	return build.Delete()
}

func GetOneByLevel(level int64) *Log {
	build := NewLogQuery()
	build.kWheLevel(level)
	return build.GetOne()
}

func UpdateLogByLevel(level int64, p *Log) int64 {
	build := NewLogQuery()
	build.kWheLevel(level)
	return build.update(p)
}

func FetchByLevel(level int64, order ...string) LogCollect {
	build := NewLogQuery()
	build.kWheLevel(level)
	build.OrderBy(parseOrderString(order))
	return build.Get()
}

func FetchByLevelWithPageSize(level int64, page int, pageSize int, order ...string) LogCollect {
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * pageSize

	build := NewLogQuery()
	build.kWheLevel(level)
	build.OrderBy(parseOrderString(order))
	return build.Skip(offset).Limit(pageSize).Get()
}

func FetchByLevels(level []int64, order ...string) LogCollect {
	build := NewLogQuery()

	if len(level) == 0 {
		return make(LogCollect, 0)
	}

	if len(level) == 1 {
		build.kWheLevel(level[0])
	} else {
		build.kWheLevelIn(level)
	}

	build.OrderBy(parseOrderString(order))
	return build.Get()
}

func UpdateLogByLevels(level []int64, p *Log) int64 {
	build := NewLogQuery()

	if len(level) == 0 {
		return 0
	}

	if len(level) == 1 {
		build.kWheLevel(level[0])
	} else {
		build.kWheLevelIn(level)
	}

	return build.update(p)
}

func FetchByLevelsWithPageSize(level []int64, page int, pageSize int, order ...string) LogCollect {
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * pageSize

	build := NewLogQuery()

	if len(level) == 0 {
		return make(LogCollect, 0)
	}

	if len(level) == 1 {
		build.kWheLevel(level[0])
	} else {
		build.kWheLevelIn(level)
	}

	build.OrderBy(parseOrderString(order))
	return build.Skip(offset).Limit(pageSize).Get()
}

func UpdateLogByLevelsWhatEver(level []int64, p *Log) int64 {
	build := NewLogQuery()

	if len(level) == 1 {
		build.kWheLevel(level[0])
	} else {
		build.kWheLevelIn(level)
	}

	return build.update(p)
}

func CountLogByLevels(level []int64) int64 {
	if len(level) == 0 {
		return 0
	}
	build := NewLogQuery()
	if len(level) == 1 {
		build.kWheLevel(level[0])
	} else {
		build.kWheLevelIn(level)
	}

	return build.Count()
}
func UpdateLogByIds(id []int64, p *Log) int64 {
	build := NewLogQuery()

	if len(id) == 0 {
		return 0
	}

	if len(id) == 1 {
		build.kWheId(id[0])
	} else {
		build.kWheIdIn(id)
	}

	return build.update(p)
}

func UpdateLogByIdsWhatEver(id []int64, p *Log) int64 {
	build := NewLogQuery()

	if len(id) == 1 {
		build.kWheId(id[0])
	} else {
		build.kWheIdIn(id)
	}

	return build.update(p)
}

func FetchByIdsWithPageSize(id []int64, page int, pageSize int, order ...string) LogCollect {
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * pageSize

	build := NewLogQuery()

	if len(id) == 0 {
		return make(LogCollect, 0)
	}

	if len(id) == 1 {
		build.kWheId(id[0])
	} else {
		build.kWheIdIn(id)
	}
	build.OrderBy(parseOrderString(order))
	return build.Skip(offset).Limit(pageSize).Get()
}

func CheckExistById(id int64) bool {
	build := NewLogQuery()
	build.kWheId(id)
	cnt := build.Count()
	return cnt > 0
}

func GetFirstById(id int64) *Log {
	build := NewLogQuery()
	build.kWheId(id)
	return build.First()
}

func DeleteById(id int64) int64 {
	build := NewLogQuery()
	build.kWheId(id)
	return build.Delete()
}

func GetOneById(id int64) *Log {
	build := NewLogQuery()
	build.kWheId(id)
	return build.GetOne()
}

func UpdateLogById(id int64, p *Log) int64 {
	build := NewLogQuery()
	build.kWheId(id)
	return build.update(p)
}
