package admin

import (
	"rygo/app/controller/base"
	"rygo/app/model"

	"rygo/app/response"
	"rygo/app/service"

	"rygo/app/utils/gconv"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	base.BaseController
}

//列表页
func (c *PostController) List(ctx *gin.Context) {
	response.BuildTpl(ctx, "system/post/list").WriteTpl()
}

//列表分页数据
func (c *PostController) ListAjax(ctx *gin.Context) {
	var req *model.PostSelectPageReq
	//获取参数
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorResp(ctx).SetMsg(err.Error()).Log("岗位管理", req).WriteJsonExit()
		return
	}
	rows := make([]model.SysPost, 0)
	result, page, err := service.PostService.SelectListByPage(req)

	if err == nil && len(result) > 0 {
		rows = result
	}

	response.BuildTable(ctx, page.Total, rows).WriteJsonExit()
}

//新增页面
func (c *PostController) Add(ctx *gin.Context) {
	response.BuildTpl(ctx, "system/post/add").WriteTpl()
}

//新增页面保存
func (c *PostController) AddSave(ctx *gin.Context) {
	var req *model.PostAddReq
	//获取参数
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Add).SetMsg(err.Error()).Log("岗位管理", req).WriteJsonExit()
		return
	}

	if service.PostService.CheckPostNameUniqueAll(req.PostName) == "1" {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Add).SetMsg("岗位名称已存在").Log("岗位管理", req).WriteJsonExit()
		return
	}

	if service.PostService.CheckPostCodeUniqueAll(req.PostCode) == "1" {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Add).SetMsg("岗位编码已存在").Log("岗位管理", req).WriteJsonExit()
		return
	}

	pid, err := service.PostService.AddSave(req, ctx)

	if err != nil || pid <= 0 {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Add).Log("岗位管理", req).WriteJsonExit()
		return
	}
	response.ErrorResp(ctx).SetData(pid).SetBtype(model.Buniss_Add).Log("岗位管理", req).WriteJsonExit()
}

//修改页面
func (c *PostController) Edit(ctx *gin.Context) {
	id := gconv.Int64(ctx.Query("id"))
	if id <= 0 {
		response.BuildTpl(ctx, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "参数错误",
		})
		return
	}

	post, err := service.PostService.SelectRecordById(id)

	if err != nil || post == nil {
		response.BuildTpl(ctx, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "岗位不存在",
		})
		return
	}

	response.BuildTpl(ctx, "system/post/edit").WriteTpl(gin.H{
		"post": post,
	})
}

//修改页面保存
func (c *PostController) EditSave(ctx *gin.Context) {
	var req *model.PostEditReq
	//获取参数
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Edit).SetMsg(err.Error()).Log("岗位管理", req).WriteJsonExit()
		return
	}

	if service.PostService.CheckPostNameUnique(req.PostName, req.PostId) == "1" {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Edit).SetMsg("岗位名称已存在").Log("岗位管理", req).WriteJsonExit()
		return
	}

	if service.PostService.CheckPostCodeUnique(req.PostCode, req.PostId) == "1" {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Edit).SetMsg("岗位编码已存在").Log("岗位管理", req).WriteJsonExit()
		return
	}

	rs, err := service.PostService.EditSave(req, ctx)

	if err != nil || rs <= 0 {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Edit).Log("岗位管理", req).WriteJsonExit()
		return
	}
	response.SucessResp(ctx).SetData(rs).SetBtype(model.Buniss_Edit).Log("岗位管理", req).WriteJsonExit()
}

//删除数据
func (c *PostController) Remove(ctx *gin.Context) {
	var req *model.RemoveReq
	//获取参数
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorResp(ctx).SetMsg(err.Error()).SetBtype(model.Buniss_Del).Log("岗位管理", req).WriteJsonExit()
		return
	}

	rs := service.PostService.DeleteRecordByIds(req.Ids)

	if rs > 0 {
		response.SucessResp(ctx).SetBtype(model.Buniss_Del).Log("岗位管理", req).WriteJsonExit()
	} else {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Del).Log("岗位管理", req).WriteJsonExit()
	}
}

//导出
func (c *PostController) Export(ctx *gin.Context) {
	var req *model.PostSelectPageReq
	//获取参数
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorResp(ctx).SetMsg(err.Error()).Log("岗位管理", req).WriteJsonExit()
		return
	}
	url, err := service.PostService.Export(req)

	if err != nil {
		response.ErrorResp(ctx).SetMsg(err.Error()).Log("岗位管理", req).WriteJsonExit()
		return
	}
	response.SucessResp(ctx).SetMsg(url).SetBtype(model.Buniss_Del).Log("岗位管理", req).WriteJsonExit()
}

//检查岗位名称是否已经存在不包括本岗位
func (c *PostController) CheckPostNameUnique(ctx *gin.Context) {
	var req *model.PostCheckPostNameReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.Writer.WriteString("1")
		return
	}

	result := service.PostService.CheckPostNameUnique(req.PostName, req.PostId)

	ctx.Writer.WriteString(result)
}

//检查岗位名称是否已经存在
func (c *PostController) CheckPostNameUniqueAll(ctx *gin.Context) {
	var req *model.PostCheckPostNameALLReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.Writer.WriteString("1")
		return
	}

	result := service.PostService.CheckPostNameUniqueAll(req.PostName)

	ctx.Writer.WriteString(result)
}

//检查岗位编码是否已经存在不包括本岗位
func (c *PostController) CheckPostCodeUnique(ctx *gin.Context) {
	var req *model.PostCheckPostCodeReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.Writer.WriteString("1")
		return
	}

	result := service.PostService.CheckPostCodeUnique(req.PostCode, req.PostId)

	ctx.Writer.WriteString(result)
}

//检查岗位编码是否已经存在
func (c *PostController) CheckPostCodeUniqueAll(ctx *gin.Context) {
	var req *model.PostCheckPostCodeALLReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.Writer.WriteString("1")
		return
	}

	result := service.PostService.CheckPostCodeUniqueAll(req.PostCode)

	ctx.Writer.WriteString(result)
}
