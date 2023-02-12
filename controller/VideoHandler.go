package controller

import (
	"github.com/livegoplayer/go_helper/utils"
	"github.com/livegoplayer/go_video_store/controller/dto"
	"github.com/livegoplayer/go_video_store/server"
	"github.com/livegoplayer/go_video_store/service/video_search"
	"strings"
)

func VideoHandler(c *server.Context) {
	videoRequest := &dto.VideoRequest{}
	c.CheckError(c.Bind(videoRequest))

	VideoIds := utils.AsInt64Slice(strings.Split(videoRequest.VideoIds, ","))
	BaseTagList := utils.AsInt64Slice(strings.Split(videoRequest.BaseTagList, ","))

	res := video_search.SearchByCondition(VideoIds, videoRequest.VideoNameSearch, BaseTagList, videoRequest.Page, videoRequest.PageSize)
	c.SuccessResp(res)
	return
}
