package tag

/* Unscoped *************************************** 满足常用的共有方法,根据mysql索引生成 **********************************************/
/**************************************** all 不是小表不要用 **************************************************/

func Save(f *Tag) *Tag {
	NewTagQuery().save(f)
	return f
}

func UpdateTagAll(p *Tag) int64 {
	build := NewTagQuery()
	return build.update(p)
}

func FetchTagAll() TagCollect {
	build := NewTagQuery()
	return build.Get()
}

func FetchTagAllWithPageSize(page int, pageSize int) TagCollect {
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * pageSize

	build := NewTagQuery()
	return build.Skip(offset).Limit(pageSize).Get()
}

func CountTagAll() int64 {
	build := NewTagQuery()
	return build.Count()
}

/**************************************** 根据 Index 生成的方法 **************************************************/

func CheckExistByTagId(tagId int64) bool {
	build := NewTagQuery()
	build.kWheTagId(tagId)
	cnt := build.Count()
	return cnt > 0
}

func GetFirstByTagId(tagId int64) *Tag {
	build := NewTagQuery()
	build.kWheTagId(tagId)
	return build.First()
}

func DeleteAllByTagId(tagId int64) int64 {
	build := NewTagQuery()
	build.kWheTagId(tagId)
	return build.Delete()
}

func GetOneByTagId(tagId int64) *Tag {
	build := NewTagQuery()
	build.kWheTagId(tagId)
	return build.GetOne()
}

func UpdateTagByTagId(tagId int64, p *Tag) int64 {
	build := NewTagQuery()
	build.kWheTagId(tagId)
	return build.update(p)
}

func FetchByTagId(tagId int64) TagCollect {
	build := NewTagQuery()
	build.kWheTagId(tagId)
	return build.Get()
}

func FetchByTagIdWithPageSize(tagId int64, page int, pageSize int) TagCollect {
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * pageSize

	build := NewTagQuery()
	return build.Skip(offset).Limit(pageSize).Get()
}

func FetchByTagIds(tagId []int64) TagCollect {
	build := NewTagQuery()

	if len(tagId) == 0 {
		return make(TagCollect, 0)
	}

	if len(tagId) == 1 {
		build.kWheTagId(tagId[0])
	} else {
		build.kWheTagIdIn(tagId)
	}

	return build.Get()
}

func UpdateTagByTagIds(tagId []int64, p *Tag) int64 {
	build := NewTagQuery()

	if len(tagId) == 0 {
		return 0
	}

	if len(tagId) == 1 {
		build.kWheTagId(tagId[0])
	} else {
		build.kWheTagIdIn(tagId)
	}

	return build.update(p)
}

func FetchByTagIdsWithPageSize(tagId []int64, page int, pageSize int) TagCollect {
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * pageSize

	build := NewTagQuery()

	if len(tagId) == 0 {
		return make(TagCollect, 0)
	}

	if len(tagId) == 1 {
		build.kWheTagId(tagId[0])
	} else {
		build.kWheTagIdIn(tagId)
	}

	return build.Skip(offset).Limit(pageSize).Get()
}

func UpdateTagByTagIdsWhatEver(tagId []int64, p *Tag) int64 {
	build := NewTagQuery()

	if len(tagId) == 1 {
		build.kWheTagId(tagId[0])
	} else {
		build.kWheTagIdIn(tagId)
	}

	return build.update(p)
}

func CountTagByTagIds(tagId []int64) int64 {
	if len(tagId) == 0 {
		return 0
	}
	build := NewTagQuery()
	if len(tagId) == 1 {
		build.kWheTagId(tagId[0])
	} else {
		build.kWheTagIdIn(tagId)
	}

	return build.Count()
}
func UpdateTagByIds(id []int64, p *Tag) int64 {
	build := NewTagQuery()

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

func UpdateTagByIdsWhatEver(id []int64, p *Tag) int64 {
	build := NewTagQuery()

	if len(id) == 1 {
		build.kWheId(id[0])
	} else {
		build.kWheIdIn(id)
	}

	return build.update(p)
}

func FetchByIdsWithPageSize(id []int64, page int, pageSize int) TagCollect {
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * pageSize

	build := NewTagQuery()

	if len(id) == 0 {
		return make(TagCollect, 0)
	}

	if len(id) == 1 {
		build.kWheId(id[0])
	} else {
		build.kWheIdIn(id)
	}

	return build.Skip(offset).Limit(pageSize).Get()
}

func CheckExistById(id int64) bool {
	build := NewTagQuery()
	build.kWheId(id)
	cnt := build.Count()
	return cnt > 0
}

func GetFirstById(id int64) *Tag {
	build := NewTagQuery()
	build.kWheId(id)
	return build.First()
}

func DeleteById(id int64) int64 {
	build := NewTagQuery()
	build.kWheId(id)
	return build.Delete()
}

func GetOneById(id int64) *Tag {
	build := NewTagQuery()
	build.kWheId(id)
	return build.GetOne()
}

func UpdateTagById(id int64, p *Tag) int64 {
	build := NewTagQuery()
	build.kWheId(id)
	return build.update(p)
}
