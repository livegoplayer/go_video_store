package controller

import (
	"github.com/livegoplayer/go_helper/pagination"
	"github.com/livegoplayer/go_helper/utils"
	"github.com/livegoplayer/go_video_store/controller/dto"
	"github.com/livegoplayer/go_video_store/server"
	"github.com/livegoplayer/go_video_store/service/tags"
	"github.com/livegoplayer/go_video_store/service/video_search"
)

func VideoHandler(c *server.Context) {
	videoRequest := &dto.VideoRequest{}
	c.CheckError(c.Bind(videoRequest))

	VideoIds := utils.StrToNumSlice(videoRequest.VideoId)
	BaseTagList := utils.StrToNumSlice(videoRequest.BaseTagList)
	page := pagination.NewNumPage(int(videoRequest.Page), int(videoRequest.PageSize))

	res := video_search.SearchByCondition(VideoIds, videoRequest.VideoNameSearch, BaseTagList, page.PageNum, page.Size)

	data := map[string]interface{}{
		"list": res,
	}
	c.SuccessResp(data)
	return
}

func TagHandler(c *server.Context) {
	res := tags.TagList()
	c.SuccessResp(res)
}
