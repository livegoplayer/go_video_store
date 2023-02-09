package controller

import (
	"github.com/livegoplayer/go_video_store/controller/dto"
	"github.com/livegoplayer/go_video_store/server"
	"github.com/livegoplayer/go_video_store/service/video_search"
)

func VideoHandler(c *server.Context) {
	videoRequest := &dto.VideoRequest{}
	c.CheckError(c.Bind(videoRequest))

	res := video_search.SearchByCondition(videoRequest.VideoIds, videoRequest.VideoNameSearch, videoRequest.BaseTagList, videoRequest.Page, videoRequest.PageSize)
	c.SuccessResp(200, res)
	return
}
