package controller

import (
	"github.com/livegoplayer/go_video_store/repository_template_model/log"
	"github.com/livegoplayer/go_video_store/server"
)

func TestHandler(c *server.Context) {
	res := log.FetchLogAll("id desc")
	c.JSON(200, res)
}
