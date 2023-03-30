package admin

import (
	"net/http"
	"rygo/app/controller/base"
	"rygo/app/model"
	"rygo/app/response"

	"github.com/gin-gonic/gin"
)

type OperateController struct {
	base.BaseController
}

func (c *OperateController) Add(ctx *gin.Context) {
	response.BuildTpl(ctx, "demo/operate/add").WriteTpl()
}

func (c *OperateController) Detail(ctx *gin.Context) {
	var tmp model.UserInfo
	tmp.UserId = 1
	tmp.UserName = "测试1"
	tmp.Status = "0"
	tmp.CreateTime = "2020-01-12 02:02:02"
	tmp.UserBalance = 100
	tmp.UserCode = "1000001"
	tmp.UserSex = "0"
	tmp.UserPhone = "15888888888"
	tmp.UserEmail = "111@qq.com"
	response.BuildTpl(ctx, "demo/operate/detail").WriteTpl(gin.H{"user": tmp})
}

func (c *OperateController) EditSave(ctx *gin.Context) {
	var tmp model.UserInfo
	tmp.UserId = 1
	tmp.UserName = "测试1"
	tmp.Status = "0"
	tmp.CreateTime = "2020-01-12 02:02:02"
	tmp.UserBalance = 100
	tmp.UserCode = "1000001"
	tmp.UserSex = "0"
	tmp.UserPhone = "15888888888"
	tmp.UserEmail = "111@qq.com"
	response.SucessResp(ctx).SetData(tmp).Log("demo演示", gin.H{"UserId": 1}).WriteJsonExit()
}

func (c *OperateController) Edit(ctx *gin.Context) {
	var tmp model.UserInfo
	tmp.UserId = 1
	tmp.UserName = "测试1"
	tmp.Status = "0"
	tmp.CreateTime = "2020-01-12 02:02:02"
	tmp.UserBalance = 100
	tmp.UserCode = "1000001"
	tmp.UserSex = "0"
	tmp.UserPhone = "15888888888"
	tmp.UserEmail = "111@qq.com"
	response.BuildTpl(ctx, "demo/operate/edit").WriteTpl(gin.H{"user": tmp})
}

func (c *OperateController) Other(ctx *gin.Context) {
	response.BuildTpl(ctx, "demo/operate/other").WriteTpl()
}

func (c *OperateController) Table(ctx *gin.Context) {
	response.BuildTpl(ctx, "demo/operate/table").WriteTpl()
}

func (c *OperateController) List(ctx *gin.Context) {
	var rows = make([]model.UserInfo, 0)
	for i := 1; i <= 10; i++ {
		var tmp model.UserInfo
		tmp.UserId = int64(i)
		tmp.UserName = "测试" + string(i)
		tmp.Status = "0"
		tmp.CreateTime = "2020-01-12 02:02:02"
		tmp.UserBalance = 100
		tmp.UserCode = "100000" + string(i)
		tmp.UserSex = "0"
		tmp.UserPhone = "15888888888"
		tmp.UserEmail = "111@qq.com"
		rows = append(rows, tmp)
	}
	ctx.JSON(http.StatusOK, model.TableDataInfo{
		Code:  0,
		Msg:   "操作成功",
		Total: len(rows),
		Rows:  rows,
	})
}
