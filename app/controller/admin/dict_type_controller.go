package admin

import (
	"net/http"
	"rygo/app/controller/base"
	"rygo/app/model"

	"rygo/app/response"
	"rygo/app/service"

	"rygo/app/utils/gconv"

	"github.com/gin-gonic/gin"
)

type DictTypeController struct {
	base.BaseController
}

//列表页
func (c *DictTypeController) List(ctx *gin.Context) {
	response.BuildTpl(ctx, "system/dict/type/list").WriteTpl()
}

//列表分页数据
func (c *DictTypeController) ListAjax(ctx *gin.Context) {
	var req *model.DictTypeSelectPageReq
	//获取参数
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorResp(ctx).SetMsg(err.Error()).Log("字典类型管理", req).WriteJsonExit()
		return
	}
	rows := make([]model.SysDictType, 0)
	result, page, err := service.DictTypeService.SelectListByPage(req)

	if err == nil && len(result) > 0 {
		rows = result
	}

	response.BuildTable(ctx, page.Total, rows).WriteJsonExit()
}

//新增页面
func (c *DictTypeController) Add(ctx *gin.Context) {
	response.BuildTpl(ctx, "system/dict/type/add").WriteTpl()
}

//新增页面保存
func (c *DictTypeController) AddSave(ctx *gin.Context) {
	var req *model.DictTypeAddReq
	//获取参数
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Add).SetMsg(err.Error()).Log("字典管理", req).WriteJsonExit()
		return
	}

	if service.DictTypeService.CheckDictTypeUniqueAll(req.DictType) == "1" {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Add).SetMsg("字典类型已存在").Log("字典管理", req).WriteJsonExit()
		return
	}

	rid, err := service.DictTypeService.AddSave(req, ctx)

	if err != nil || rid <= 0 {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Add).Log("字典管理", req).WriteJsonExit()
		return
	}
	response.SucessResp(ctx).SetData(rid).Log("字典管理", req).WriteJsonExit()
}

//修改页面
func (c *DictTypeController) Edit(ctx *gin.Context) {
	id := gconv.Int64(ctx.Query("id"))
	if id <= 0 {
		response.BuildTpl(ctx, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "字典类型错误",
		})
		return
	}

	entity, err := service.DictTypeService.SelectRecordById(id)

	if err != nil || entity == nil {
		response.BuildTpl(ctx, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "字典类型不存在",
		})
		return
	}

	response.BuildTpl(ctx, "system/dict/type/edit").WriteTpl(gin.H{
		"dict": entity,
	})
}

//修改页面保存
func (c *DictTypeController) EditSave(ctx *gin.Context) {
	var req *model.DictTypeEditReq
	//获取参数
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Edit).SetMsg(err.Error()).Log("字典类型管理", req).WriteJsonExit()
		return
	}

	if service.DictTypeService.CheckDictTypeUnique(req.DictType, req.DictId) == "1" {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Edit).SetMsg("字典类型已存在").Log("字典类型管理", req).WriteJsonExit()
		return
	}

	rs, err := service.DictTypeService.EditSave(req, ctx)

	if err != nil || rs <= 0 {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Edit).Log("字典类型管理", req).WriteJsonExit()
		return
	}
	response.SucessResp(ctx).Log("字典类型管理", req).WriteJsonExit()
}

//删除数据
func (c *DictTypeController) Remove(ctx *gin.Context) {
	var req *model.RemoveReq
	//获取参数
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Del).SetMsg(err.Error()).Log("字典管理", req).WriteJsonExit()
		return
	}

	rs := service.DictTypeService.DeleteRecordByIds(req.Ids)

	if rs > 0 {
		response.SucessResp(ctx).SetBtype(model.Buniss_Del).Log("字典管理", req).WriteJsonExit()
	} else {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Del).Log("字典管理", req).WriteJsonExit()
	}
}

//数据详情
func (c *DictTypeController) Detail(ctx *gin.Context) {
	dictId := gconv.Int64(ctx.Query("dictId"))
	if dictId <= 0 {
		response.BuildTpl(ctx, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "参数错误",
		})
		return
	}
	dict, _ := service.DictTypeService.SelectRecordById(dictId)

	if dict == nil {
		response.BuildTpl(ctx, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "字典类别不存在",
		})
		return
	}

	dictList, _ := service.DictTypeService.SelectListAll(nil)
	if dictList == nil {
		response.BuildTpl(ctx, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "参数错误2",
		})
		return
	}

	response.BuildTpl(ctx, "system/dict/data/list").WriteTpl(gin.H{
		"dict":     dict,
		"dictList": dictList,
	})
}

//选择字典树
func (c *DictTypeController) SelectDictTree(ctx *gin.Context) {
	columnId := gconv.Int64(ctx.Query("columnId"))
	dictType := ctx.Query("dictType")
	if columnId <= 0 || dictType == "" {
		response.BuildTpl(ctx, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "参数错误",
		})

		return
	}

	if dictType == "-" {
		dictType = "-"
	}

	var dict model.SysDictType
	rs := service.DictTypeService.SelectDictTypeByType(dictType)
	if rs != nil {
		dict = *rs
	}

	response.BuildTpl(ctx, "system/dict/type/tree").WriteTpl(gin.H{
		"columnId": columnId,
		"dict":     dict,
	})
}

//导出
func (c *DictTypeController) Export(ctx *gin.Context) {
	var req *model.DictTypeSelectPageReq
	//获取参数
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorResp(ctx).SetMsg(err.Error()).Log("字典管理", req).WriteJsonExit()
		return
	}
	url, err := service.DictTypeService.Export(req)

	if err != nil {
		response.ErrorResp(ctx).SetMsg(err.Error()).Log("字典管理", req).WriteJsonExit()
		return
	}

	response.SucessResp(ctx).SetMsg(url).Log("导出Excel", req).WriteJsonExit()
}

//检查字典类型是否唯一不包括本参数
func (c *DictTypeController) CheckDictTypeUnique(ctx *gin.Context) {
	var req *model.CheckDictTypeReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.Writer.WriteString("1")
		return
	}

	result := service.DictTypeService.CheckDictTypeUnique(req.DictType, req.DictId)
	ctx.Writer.WriteString(result)
}

//检查字典类型是否唯一
func (c *DictTypeController) CheckDictTypeUniqueAll(ctx *gin.Context) {
	var req *model.CheckDictTypeALLReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.Writer.WriteString("1")
		return
	}

	result := service.DictTypeService.CheckDictTypeUniqueAll(req.DictType)

	ctx.Writer.WriteString(result)
}

//加载部门列表树结构的数据
func (c *DictTypeController) TreeData(ctx *gin.Context) {
	result := service.DictTypeService.SelectDictTree(nil)
	ctx.JSON(http.StatusOK, result)
}
