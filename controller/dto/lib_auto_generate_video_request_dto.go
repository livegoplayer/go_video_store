package dto

type unknownSlice []string
type BaseTagList []int64
type VideoId []int64

type VideoRequest struct {
	VideoNameSearch string      `json:"video_name_search" form:"video_name_search"`
	BaseTagList     BaseTagList `json:"base_tag_list" form:"base_tag_list"`
	Page            int64       `json:"page" form:"page"`
	VideoIds        VideoId     `json:"video_ids" form:"video_ids"`
	PageSize        int64       `json:"page_size" form:"page_size"`
}

var VideoRequestMock = VideoRequest{
	VideoNameSearch: "武则天",
	VideoIds: VideoId{
		1,
		2,
		3,
		4,
		5,
	},
	PageSize: 10,
	BaseTagList: BaseTagList{
		1,
		2,
		3,
		4,
		5,
	},
	Page: 1,
}
