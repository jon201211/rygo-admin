package admin

import (
	"rygo/app/controller/base"
	"rygo/app/model"

	"rygo/app/response"
	"rygo/app/service"

	"rygo/app/utils/gconv"

	"github.com/gin-gonic/gin"
)

type TenantController struct {
	base.BaseController
}

//列表页
func (c *TenantController) List(ctx *gin.Context) {
	response.BuildTpl(ctx, "module/tenant/list.html").WriteTpl()
}

//列表分页数据
func (c *TenantController) ListAjax(ctx *gin.Context) {
	req := new(model.TenantSelectPageReq)
	//获取参数
	if err := ctx.ShouldBind(req); err != nil {
		response.ErrorResp(ctx).SetMsg(err.Error()).Log("tenant管理", req).WriteJsonExit()
		return
	}
	rows := make([]model.SysTenant, 0)
	result, page, err := service.TenantService.SelectListByPage(req)

	if err == nil && len(result) > 0 {
		rows = result
	}

	response.BuildTable(ctx, page.Total, rows).WriteJsonExit()
}

//新增页面
func (c *TenantController) Add(ctx *gin.Context) {
	response.BuildTpl(ctx, "module/tenant/add.html").WriteTpl()
}

//新增页面保存
func (c *TenantController) AddSave(ctx *gin.Context) {
	req := new(model.TenantAddReq)
	//获取参数
	if err := ctx.ShouldBind(req); err != nil {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Add).SetMsg(err.Error()).Log("商户信息新增数据", req).WriteJsonExit()
		return
	}

	id, err := service.TenantService.AddSave(req, ctx)

	if err != nil || id <= 0 {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Add).Log("商户信息新增数据", req).WriteJsonExit()
		return
	}
	response.SucessResp(ctx).SetData(id).Log("商户信息新增数据", req).WriteJsonExit()
}

//修改页面
func (c *TenantController) Edit(ctx *gin.Context) {
	id := gconv.Int64(ctx.Query("id"))

	if id <= 0 {
		response.ErrorTpl(ctx).WriteTpl(gin.H{
			"desc": "参数错误",
		})
		return
	}

	entity, err := service.TenantService.SelectRecordById(id)

	if err != nil || entity == nil {
		response.ErrorTpl(ctx).WriteTpl(gin.H{
			"desc": "数据不存在",
		})
		return
	}

	response.BuildTpl(ctx, "module/tenant/edit.html").WriteTpl(gin.H{
		"tenant": entity,
	})
}

//修改页面保存
func (c *TenantController) EditSave(ctx *gin.Context) {
	var req = new(model.TenantEditReq)
	//获取参数
	if err := ctx.ShouldBind(req); err != nil {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Edit).SetMsg(err.Error()).Log("商户信息修改数据", req).WriteJsonExit()
		return
	}

	rs, err := service.TenantService.EditSave(req, ctx)
	if err != nil || rs <= 0 {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Edit).Log("商户信息修改数据", req).WriteJsonExit()
		return
	}
	response.SucessResp(ctx).SetBtype(model.Buniss_Edit).Log("商户信息修改数据", req).WriteJsonExit()
}

//删除数据
func (c *TenantController) Remove(ctx *gin.Context) {
	req := new(model.RemoveReq)
	//获取参数
	if err := ctx.ShouldBind(req); err != nil {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Del).SetMsg(err.Error()).Log("商户信息删除数据", req).WriteJsonExit()
		return
	}

	rs := service.TenantService.DeleteRecordByIds(req.Ids)

	if rs > 0 {
		response.SucessResp(ctx).SetBtype(model.Buniss_Del).Log("商户信息删除数据", req).WriteJsonExit()
	} else {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Del).Log("商户信息删除数据", req).WriteJsonExit()
	}
}

//导出
func (c *TenantController) Export(ctx *gin.Context) {
	req := new(model.TenantSelectPageReq)
	//获取参数
	if err := ctx.ShouldBind(req); err != nil {
		response.ErrorResp(ctx).Log("商户信息导出数据", req).WriteJsonExit()
		return
	}
	url, err := service.TenantService.Export(req)

	if err != nil {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Other).Log("商户信息导出数据", req).WriteJsonExit()
		return
	}
	response.SucessResp(ctx).SetBtype(model.Buniss_Other).SetMsg(url).WriteJsonExit()
}
