package video

func FetchByName(search string) VideoCollect {
	build := NewVideoQuery()
	build.kWheMetaName("like", "%"+search+"%")
	return build.Get()
}
