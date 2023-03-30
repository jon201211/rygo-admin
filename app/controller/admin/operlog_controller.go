package admin

import (
	"html/template"
	"net/http"
	"rygo/app/controller/base"
	"rygo/app/model"
	"rygo/app/response"
	"rygo/app/service"

	"rygo/app/utils/gconv"

	"github.com/gin-gonic/gin"
)

type OperlogController struct {
	base.BaseController
}

//用户列表页
func (c *OperlogController) List(ctx *gin.Context) {
	response.BuildTpl(ctx, "monitor/operlog/list").WriteTpl()
}

//用户列表分页数据
func (c *OperlogController) ListAjax(ctx *gin.Context) {
	var req *model.OperLogSelectPageReq
	//获取参数
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, model.CommonRes{
			Code: 500,
			Msg:  err.Error(),
		})
	}

	rows := make([]model.OperLogEntity, 0)

	result, page, err := service.OperlogService.SelectPageList(req)

	if err == nil && len(*result) > 0 {
		rows = *result
	}

	response.BuildTable(ctx, page.Total, rows).WriteJsonExit()
}

//清空记录
func (c *OperlogController) Clean(ctx *gin.Context) {

	rs, _ := service.OperlogService.DeleteRecordAll()

	if rs > 0 {
		response.SucessResp(ctx).SetBtype(model.Buniss_Del).SetData(rs).Log("操作日志管理", "all").WriteJsonExit()
	} else {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Del).Log("操作日志管理", "all").WriteJsonExit()
	}
}

//删除数据
func (c *OperlogController) Remove(ctx *gin.Context) {
	var req *model.RemoveReq
	//获取参数
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Del).SetMsg(err.Error()).Log("操作日志管理", req).WriteJsonExit()
		return
	}

	rs := service.OperlogService.DeleteRecordByIds(req.Ids)

	if rs > 0 {
		response.SucessResp(ctx).SetBtype(model.Buniss_Del).SetData(rs).Log("操作日志管理", req).WriteJsonExit()
	} else {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Del).Log("操作日志管理", req).WriteJsonExit()
	}
}

//记录详情
func (c *OperlogController) Detail(ctx *gin.Context) {
	id := gconv.Int64(ctx.Query("id"))
	if id <= 0 {
		response.BuildTpl(ctx, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "参数错误",
		})
		return
	}

	operLog, err := service.OperlogService.SelectRecordById(id)

	if err != nil {
		response.BuildTpl(ctx, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "数据不存在",
		})
		return
	}

	jsonResult := template.HTML(operLog.JsonResult)
	operParam := template.HTML(operLog.OperParam)
	response.BuildTpl(ctx, "monitor/operlog/detail").WriteTpl(gin.H{
		"operLog":    operLog,
		"jsonResult": jsonResult,
		"operParam":  operParam,
	})
}

//导出
func (c *OperlogController) Export(ctx *gin.Context) {
	var req *model.OperLogSelectPageReq
	//获取参数
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorResp(ctx).SetMsg(err.Error()).Log("导出操作日志", req).WriteJsonExit()
		return
	}
	url, err := service.OperlogService.Export(req)

	if err != nil {
		response.ErrorResp(ctx).SetMsg(err.Error()).Log("导出操作日志", req).WriteJsonExit()
	} else {
		response.SucessResp(ctx).SetMsg(url).Log("导出操作日志", req).WriteJsonExit()
	}
}
