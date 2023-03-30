package admin

import (
	"rygo/app/controller/base"
	"rygo/app/model"

	"rygo/app/response"
	"rygo/app/service"

	"github.com/gin-gonic/gin"
)

type LogininforController struct {
	base.BaseController
}

//用户列表页
func (c *LogininforController) List(ctx *gin.Context) {
	response.BuildTpl(ctx, "monitor/logininfor/list").WriteTpl()
}

//用户列表分页数据
func (c *LogininforController) ListAjax(ctx *gin.Context) {
	var req *model.LogininforSelectPageReq
	//获取参数
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorResp(ctx).SetMsg(err.Error()).WriteJsonExit()
		return
	}

	rows := make([]model.LogininforEntity, 0)

	result, page, err := service.LogininforService.SelectPageList(req)

	if err == nil && len(*result) > 0 {
		rows = *result
	}
	response.BuildTable(ctx, page.Total, rows).WriteJsonExit()
}

//删除数据
func (c *LogininforController) Remove(ctx *gin.Context) {
	var req *model.RemoveReq
	//获取参数
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Del).SetMsg(err.Error()).Log("登陆日志管理", req).WriteJsonExit()
		return
	}

	rs := service.LogininforService.DeleteRecordByIds(req.Ids)

	if rs > 0 {
		response.SucessResp(ctx).SetBtype(model.Buniss_Del).SetData(rs).Log("登陆日志管理", req).WriteJsonExit()
	} else {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Del).Log("登陆日志管理", req).WriteJsonExit()
	}
}

//清空记录
func (c *LogininforController) Clean(ctx *gin.Context) {

	rs, _ := service.LogininforService.DeleteRecordAll()

	if rs > 0 {
		response.SucessResp(ctx).SetBtype(model.Buniss_Del).SetData(rs).Log("登陆日志管理", "all").WriteJsonExit()
	} else {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Del).Log("登陆日志管理", "all").WriteJsonExit()
	}
}

//导出
func (c *LogininforController) Export(ctx *gin.Context) {
	var req *model.LogininforSelectPageReq
	//获取参数
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorResp(ctx).SetMsg(err.Error()).Log("导出登陆日志", req).WriteJsonExit()
		return
	}

	url, err := service.LogininforService.Export(req)

	if err != nil {
		response.ErrorResp(ctx).SetMsg(err.Error()).Log("导出登陆日志", req).WriteJsonExit()
	} else {
		response.SucessResp(ctx).SetMsg(url).Log("导出登陆日志", req).WriteJsonExit()
	}
}

//解锁账号
func (c *LogininforController) Unlock(ctx *gin.Context) {
	loginName := ctx.Query("loginName")
	if loginName == "" {
		response.ErrorResp(ctx).SetMsg("参数错误").Log("解锁账号", "loginName="+loginName).WriteJsonExit()
	} else {
		service.LogininforService.RemovePasswordCounts(loginName)
		service.LogininforService.Unlock(loginName)
		response.SucessResp(ctx).Log("解锁账号", "loginName="+loginName).WriteJsonExit()
	}

}
