package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	myconfig "github.com/livegoplayer/go_video_store/config"
	"github.com/livegoplayer/go_video_store/routers"
	"github.com/spf13/viper"

	c "github.com/livegoplayer/go_gin_helper"
	"github.com/livegoplayer/go_helper/config"
)

func main() {
	// 初始化一个http服务对象
	//默认有写入控制台的日志
	// 把这两个处理器替换
	r := gin.New()
	err := r.SetTrustedProxies(nil)
	if err != nil {
		panic("SetTrustedProxiesNil error")
	}
	r.NoMethod(c.HandleNotFound)
	r.NoRoute(c.HandleNotFound)

	//加载.env文件
	config.LoadEnv()

	//设置gin的运行模式
	switch viper.GetString("ENV") {
	case config.PRODUCTION_ENV:
		gin.SetMode(gin.ReleaseMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	myconfig.InitLog()
	r = myconfig.InitAccessLog(r)

	//解决跨域问题的中间件
	r.Use(c.Cors(viper.GetStringSlice("client_list")))

	//初始化数据库连接池
	myconfig.InitDb()

	routers.InitAppRouter(r)

	myconfig.InitValidate()

	err = r.Run(":" + viper.GetString("app_port"))
	if err != nil {
		fmt.Printf("server start error : " + err.Error())
		return
	}
}
