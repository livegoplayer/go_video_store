package controller

import (
	"fmt"
	"github.com/livegoplayer/go_user_rpc/user"
	userpb "github.com/livegoplayer/go_user_rpc/user/user_grpc"
	"github.com/livegoplayer/video_store/server"
)

func TestHandler(c *server.Context) {
	//test
	fmt.Printf("test")
	userClient := user.GetUserClient()
	res, err := userClient.AddUser(c, &userpb.AddUserRequest{
		UserName: "123",
		Password: "123456",
	})

	c.CheckError(err, "新建用户失败")

	data := res.GetData()
	fmt.Printf(string(data.GetUid()))

	c.SuccessResp(data)
}

func HealthHandler(c *server.Context) {
	c.SuccessResp()
}

// 根据内容搜索
func Search(c *server.Context) {

}
