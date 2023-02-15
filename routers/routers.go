package routers

import (
	"github.com/gin-gonic/gin"
	ginHelper "github.com/livegoplayer/go_gin_helper"
	. "github.com/livegoplayer/go_video_store/controller"
	"github.com/livegoplayer/go_video_store/server"
)

func InitAppRouter(r *gin.Engine) {

	r.GET("/health", server.NewCtl(TestHandler))

	r.Use(ginHelper.ParseParams(&server.MustParams{}))
	videoGroup := r.Group("/api/video")
	{
		videoGroup.GET("/search", server.NewCtl(VideoHandler))
	}

	tagGroup := r.Group("/api/tag")
	{
		tagGroup.GET("/all", server.NewCtl(TagHandler))
	}
}
