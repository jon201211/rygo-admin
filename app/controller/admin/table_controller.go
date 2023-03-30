package admin

import (
	"net/http"
	"rygo/app/controller/base"
	"rygo/app/model"
	"rygo/app/response"

	"github.com/gin-gonic/gin"
)


type TableController struct {
	base.BaseController
}

func (c *TableController) Button(ctx *gin.Context) {
	response.BuildTpl(ctx, "demo/table/button").WriteTpl()
}

func (c *TableController) Child(ctx *gin.Context) {
	response.BuildTpl(ctx, "demo/table/child").WriteTpl()
}

func (c *TableController) Curd(ctx *gin.Context) {
	response.BuildTpl(ctx, "demo/table/curd").WriteTpl()
}

func (c *TableController) Detail(ctx *gin.Context) {
	response.BuildTpl(ctx, "demo/table/detail").WriteTpl()
}

func (c *TableController) Editable(ctx *gin.Context) {
	response.BuildTpl(ctx, "demo/table/editable").WriteTpl()
}

func (c *TableController) Event(ctx *gin.Context) {
	response.BuildTpl(ctx, "demo/table/event").WriteTpl()
}

func (c *TableController) Export(ctx *gin.Context) {
	response.BuildTpl(ctx, "demo/table/export").WriteTpl()
}

func (c *TableController) FixedColumns(ctx *gin.Context) {
	response.BuildTpl(ctx, "demo/table/fixedColumns").WriteTpl()
}

func (c *TableController) Footer(ctx *gin.Context) {
	response.BuildTpl(ctx, "demo/table/footer").WriteTpl()
}

func (c *TableController) GroupHeader(ctx *gin.Context) {
	response.BuildTpl(ctx, "demo/table/groupHeader").WriteTpl()
}

func (c *TableController) Image(ctx *gin.Context) {
	response.BuildTpl(ctx, "demo/table/image").WriteTpl()
}

func (c *TableController) Multi(ctx *gin.Context) {
	response.BuildTpl(ctx, "demo/table/multi").WriteTpl()
}

func (c *TableController) Other(ctx *gin.Context) {
	response.BuildTpl(ctx, "demo/table/other").WriteTpl()
}

func (c *TableController) PageGo(ctx *gin.Context) {
	response.BuildTpl(ctx, "demo/table/pageGo").WriteTpl()
}

func (c *TableController) Params(ctx *gin.Context) {
	response.BuildTpl(ctx, "demo/table/params").WriteTpl()
}

func (c *TableController) Remember(ctx *gin.Context) {
	response.BuildTpl(ctx, "demo/table/remember").WriteTpl()
}

func (c *TableController) Recorder(ctx *gin.Context) {
	response.BuildTpl(ctx, "demo/table/recorder").WriteTpl()
}

func (c *TableController) Search(ctx *gin.Context) {
	response.BuildTpl(ctx, "demo/table/search").WriteTpl()
}

func (c *TableController) List(ctx *gin.Context) {
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
