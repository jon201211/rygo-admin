package admin

import (
	"rygo/app/controller/base"
	"rygo/app/model"

	"rygo/app/response"
	"rygo/app/service"

	"rygo/app/utils/gconv"

	"github.com/gin-gonic/gin"
)

type DictDataController struct {
	base.BaseController
}

//列表分页数据
func (c *DictDataController) ListAjax(ctx *gin.Context) {
	var req *model.DictDataSelectPageReq
	//获取参数
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorResp(ctx).SetMsg(err.Error()).Log("字典数据管理", req).WriteJsonExit()
		return
	}
	rows := make([]model.SysDictData, 0)
	result, page, err := service.DictDataService.SelectListByPage(req)

	if err == nil && len(*result) > 0 {
		rows = *result
	}

	response.BuildTable(ctx, page.Total, rows).WriteJsonExit()
}

//新增页面
func (c *DictDataController) Add(ctx *gin.Context) {
	dictType := ctx.Query("dictType")
	response.BuildTpl(ctx, "system/dict/data/add").WriteTpl(gin.H{"dictType": dictType})
}

//新增页面保存
func (c *DictDataController) AddSave(ctx *gin.Context) {
	var req *model.DictDataAddReq
	//获取参数
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Add).SetMsg(err.Error()).Log("字典数据管理", req).WriteJsonExit()
		return
	}

	rid, err := service.DictDataService.AddSave(req, ctx)

	if err != nil || rid <= 0 {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Add).Log("字典数据管理", req).WriteJsonExit()
		return
	}
	response.SucessResp(ctx).SetData(rid).SetBtype(model.Buniss_Add).Log("字典数据管理", req).WriteJsonExit()
}

//修改页面
func (c *DictDataController) Edit(ctx *gin.Context) {
	id := gconv.Int64(ctx.Query("id"))
	if id <= 0 {
		response.BuildTpl(ctx, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "字典数据错误",
		})
		return
	}

	entity, err := service.DictDataService.SelectRecordById(id)

	if err != nil || entity == nil {
		response.BuildTpl(ctx, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "字典数据不存在",
		})
		return
	}

	response.BuildTpl(ctx, "system/dict/data/edit").WriteTpl(gin.H{
		"dict": entity,
	})
}

//修改页面保存
func (c *DictDataController) EditSave(ctx *gin.Context) {
	var req *model.DictDataEditReq
	//获取参数
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Edit).SetMsg(err.Error()).Log("字典数据管理", req).WriteJsonExit()
		return
	}

	rs, err := service.DictDataService.EditSave(req, ctx)

	if err != nil || rs <= 0 {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Edit).Log("字典数据管理", req).WriteJsonExit()
		return
	}
	response.SucessResp(ctx).SetBtype(model.Buniss_Edit).SetData(rs).Log("字典数据管理", req).WriteJsonExit()
}

//删除数据
func (c *DictDataController) Remove(ctx *gin.Context) {
	var req *model.RemoveReq
	//获取参数
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Del).SetMsg(err.Error()).Log("字典数据管理", req).WriteJsonExit()
		return
	}

	rs := service.DictDataService.DeleteRecordByIds(req.Ids)

	if rs > 0 {
		response.SucessResp(ctx).SetBtype(model.Buniss_Del).SetData(rs).Log("字典数据管理", req).WriteJsonExit()
	} else {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Del).Log("字典数据管理", req).WriteJsonExit()
	}
}

//导出
func (c *DictDataController) Export(ctx *gin.Context) {
	var req *model.DictDataSelectPageReq
	//获取参数
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorResp(ctx).SetMsg(err.Error()).Log("字典数据导出", req).WriteJsonExit()
		return
	}
	url, err := service.DictDataService.Export(req)

	if err != nil {
		response.ErrorResp(ctx).SetMsg(err.Error()).Log("字典数据导出", req).WriteJsonExit()
		return
	}
	response.SucessResp(ctx).SetMsg(url).Log("导出Excel", req).WriteJsonExit()
}
