package dto

type unknownSlice []string

type VideoRequest struct {
	VideoNameSearch string `json:"video_name_search" form:"video_name_search"`
	BaseTagList     string `json:"base_tag_list" form:"base_tag_list"`
	Page            int64  `json:"page" form:"page"`
	VideoId         string `json:"video_id" form:"video_id"`
	PageSize        int64  `json:"page_size" form:"page_size"`
}

var VideoRequestMock = VideoRequest{
	VideoNameSearch: "姝﹀垯澶?",
	VideoId:         "1,2,3,4,5",
	PageSize:        10,
	BaseTagList:     "1,2,3,4,5",
	Page:            1,
}
