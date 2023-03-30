package admin

import (
	"rygo/app/controller/base"
	"rygo/app/model"

	"rygo/app/response"
	"rygo/app/service"
	"rygo/app/utils/gconv"

	"github.com/gin-gonic/gin"
)

type ConfigController struct {
	base.BaseController
}

//列表页
func (c *ConfigController) List(ctx *gin.Context) {
	response.BuildTpl(ctx, "system/config/list").WriteTpl()
}

//列表分页数据
func (c *ConfigController) ListAjax(ctx *gin.Context) {
	req := new(model.SelectPageReq)
	//获取参数
	if err := ctx.ShouldBind(req); err != nil {
		response.ErrorResp(ctx).SetMsg(err.Error()).Log("参数管理", req).WriteJsonExit()
		return
	}
	rows := make([]model.SysConfig, 0)
	result, page, err := service.ConfigService.SelectListByPage(req)

	if err == nil && len(result) > 0 {
		rows = result
	}
	response.BuildTable(ctx, page.Total, rows).WriteJsonExit()
}

//新增页面
func (c *ConfigController) Add(ctx *gin.Context) {
	response.BuildTpl(ctx, "system/config/add").WriteTpl()
}

//新增页面保存
func (c *ConfigController) AddSave(ctx *gin.Context) {
	req := new(model.ConfigAddReq)
	//获取参数
	if err := ctx.ShouldBind(req); err != nil {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Add).SetMsg(err.Error()).Log("参数管理", req).WriteJsonExit()
		return
	}

	if service.ConfigService.CheckConfigKeyUniqueAll(req.ConfigKey) == "1" {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Add).SetMsg("参数键名已存在").Log("参数管理", req).WriteJsonExit()
		return
	}

	rid, err := service.ConfigService.AddSave(req, ctx)

	if err != nil || rid <= 0 {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Add).Log("参数管理", req).WriteJsonExit()
		return
	}
	response.SucessResp(ctx).SetData(rid).Log("参数管理", req).WriteJsonExit()
}

//修改页面
func (c *ConfigController) Edit(ctx *gin.Context) {
	id := gconv.Int64(ctx.Query("id"))
	if id <= 0 {
		response.BuildTpl(ctx, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "参数错误",
		})
		return
	}

	entity, err := service.ConfigService.SelectRecordById(id)

	if err != nil || entity == nil {
		response.BuildTpl(ctx, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "数据不存在",
		})
		return
	}

	response.BuildTpl(ctx, "system/config/edit").WriteTpl(gin.H{
		"config": entity,
	})
}

//修改页面保存
func (c *ConfigController) EditSave(ctx *gin.Context) {
	req := new(model.ConfigEditReq)
	//获取参数
	if err := ctx.ShouldBind(req); err != nil {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Edit).SetMsg(err.Error()).Log("参数管理", req).WriteJsonExit()
		return
	}

	if service.ConfigService.CheckConfigKeyUnique(req.ConfigKey, req.ConfigId) == "1" {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Edit).SetMsg("参数键名已存在").Log("参数管理", req).WriteJsonExit()
		return
	}

	rs, err := service.ConfigService.EditSave(req, ctx)

	if err != nil || rs <= 0 {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Edit).Log("参数管理", req).WriteJsonExit()
		return
	}
	response.SucessResp(ctx).SetBtype(model.Buniss_Edit).Log("参数管理", req).WriteJsonExit()
}

//删除数据
func (c *ConfigController) Remove(ctx *gin.Context) {
	req := new(model.RemoveReq)
	//获取参数
	if err := ctx.ShouldBind(req); err != nil {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Del).SetMsg(err.Error()).Log("参数管理", req).WriteJsonExit()
		return
	}

	rs := service.ConfigService.DeleteRecordByIds(req.Ids)

	if rs > 0 {
		response.SucessResp(ctx).SetBtype(model.Buniss_Del).Log("参数管理", req).WriteJsonExit()
	} else {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Del).Log("参数管理", req).WriteJsonExit()
	}
}

//导出
func (c *ConfigController) Export(ctx *gin.Context) {
	req := new(model.SelectPageReq)
	//获取参数
	if err := ctx.ShouldBind(req); err != nil {
		response.ErrorResp(ctx).Log("参数管理", req).WriteJsonExit()
		return
	}
	url, err := service.ConfigService.Export(req)

	if err != nil {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Other).Log("参数管理", req).WriteJsonExit()
		return
	}

	response.SucessResp(ctx).SetBtype(model.Buniss_Other).SetMsg(url).WriteJsonExit()
}

//检查参数键名是否已经存在不包括本参数
func (c *ConfigController) CheckConfigKeyUnique(ctx *gin.Context) {
	var req *model.CheckConfigKeyReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.Writer.WriteString("1")
		return
	}

	result := service.ConfigService.CheckConfigKeyUnique(req.ConfigKey, req.ConfigId)

	ctx.Writer.WriteString(result)
}

//检查参数键名是否已经存在
func (c *ConfigController) CheckConfigKeyUniqueAll(ctx *gin.Context) {
	var req *model.CheckPostCodeALLReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.Writer.WriteString("1")
		return
	}

	result := service.ConfigService.CheckConfigKeyUniqueAll(req.ConfigKey)

	ctx.Writer.WriteString(result)
}
