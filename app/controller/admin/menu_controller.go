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

type MenuController struct {
	base.BaseController
}

//列表页
func (c *MenuController) List(ctx *gin.Context) {
	response.BuildTpl(ctx, "system/menu/list").WriteTpl()
}

//列表分页数据
func (c *MenuController) ListAjax(ctx *gin.Context) {
	var req *model.MenuSelectPageReq
	//获取参数
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorResp(ctx).SetMsg(err.Error()).Log("菜单管理", req).WriteJsonExit()
		return
	}
	rows := make([]model.MenuEntity, 0)
	result, err := service.MenuService.SelectListAll(req)

	if err == nil && len(result) > 0 {
		rows = result
	}
	ctx.JSON(http.StatusOK, rows)
}

//新增页面
func (c *MenuController) Add(ctx *gin.Context) {
	pid := gconv.Int64(ctx.Query("pid"))
	var pmenu model.MenuEntityExtend
	pmenu.MenuId = 0
	pmenu.MenuName = "主目录"

	tmp, err := service.MenuService.SelectRecordById(pid)
	if err == nil && tmp != nil && tmp.MenuId > 0 {
		pmenu.MenuId = tmp.MenuId
		pmenu.MenuName = tmp.MenuName
	}
	response.BuildTpl(ctx, "system/menu/add").WriteTpl(gin.H{"menu": pmenu})
}

//新增页面保存
func (c *MenuController) AddSave(ctx *gin.Context) {
	var req *model.MenuAddReq
	//获取参数
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Add).SetMsg(err.Error()).Log("菜单管理", req).WriteJsonExit()
		return
	}

	if service.MenuService.CheckMenuNameUniqueAll(req.MenuName, req.ParentId) == "1" {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Add).SetMsg("菜单名称已存在").Log("菜单管理", req).WriteJsonExit()
		return
	}

	id, err := service.MenuService.AddSave(req, ctx)

	if err != nil || id <= 0 {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Add).SetMsg(err.Error()).Log("菜单管理", req).WriteJsonExit()
		return
	}
	response.SucessResp(ctx).SetBtype(model.Buniss_Add).SetData(id).Log("菜单管理", req).WriteJsonExit()
}

//修改页面
func (c *MenuController) Edit(ctx *gin.Context) {
	id := gconv.Int64(ctx.Query("id"))
	if id <= 0 {
		response.BuildTpl(ctx, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "参数错误",
		})
		return
	}

	menu, err := service.MenuService.SelectRecordById(id)

	if err != nil || menu == nil {
		response.BuildTpl(ctx, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "菜单不存在",
		})
		return
	}

	response.BuildTpl(ctx, "system/menu/edit").WriteTpl(gin.H{
		"menu": menu,
	})
}

//修改页面保存
func (c *MenuController) EditSave(ctx *gin.Context) {
	var req *model.MenuEditReq
	//获取参数
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Edit).SetMsg(err.Error()).Log("菜单管理", req).WriteJsonExit()
		return
	}

	if service.MenuService.CheckMenuNameUnique(req.MenuName, req.MenuId, req.ParentId) == "1" {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Edit).SetMsg("菜单名称已存在").Log("菜单管理", req).WriteJsonExit()
		return
	}

	rs, err := service.MenuService.EditSave(req, ctx)

	if err != nil || rs <= 0 {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Edit).Log("菜单管理", req).WriteJsonExit()
		return
	}
	response.SucessResp(ctx).SetBtype(model.Buniss_Edit).SetData(rs).Log("菜单管理", req).WriteJsonExit()
}

//删除数据
func (c *MenuController) Remove(ctx *gin.Context) {
	id := gconv.Int64(ctx.Query("id"))
	rs := service.MenuService.DeleteRecordById(id)

	if rs {
		response.SucessResp(ctx).SetBtype(model.Buniss_Del).Log("菜单管理", gin.H{"id": id}).WriteJsonExit()
	} else {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Del).Log("菜单管理", gin.H{"id": id}).WriteJsonExit()
	}
}

//选择菜单树
func (c *MenuController) SelectMenuTree(ctx *gin.Context) {
	menuId := gconv.Int64(ctx.Query("menuId"))
	menu, err := service.MenuService.SelectRecordById(menuId)
	if err != nil {
		response.BuildTpl(ctx, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "菜单不存在",
		})
		return
	}
	response.BuildTpl(ctx, "system/menu/tree").WriteTpl(gin.H{
		"menu": menu,
	})
}

//加载所有菜单列表树
func (c *MenuController) MenuTreeData(ctx *gin.Context) {
	user := service.UserService.GetProfile(ctx)
	if user == nil {
		response.ErrorResp(ctx).SetMsg("登陆超时").Log("菜单管理", gin.H{"userId": user.UserId}).WriteJsonExit()
		return
	}
	ztrees, err := service.MenuService.MenuTreeData(user.UserId)
	if err != nil {
		response.ErrorResp(ctx).SetMsg(err.Error()).Log("菜单管理", gin.H{"userId": user.UserId}).WriteJsonExit()
		return
	}
	ctx.JSON(http.StatusOK, ztrees)
}

//选择图标
func (c *MenuController) Icon(ctx *gin.Context) {
	response.BuildTpl(ctx, "system/menu/icon").WriteTpl()
}

//加载角色菜单列表树
func (c *MenuController) RoleMenuTreeData(ctx *gin.Context) {
	roleId := gconv.Int64(ctx.Query("roleId"))
	user := service.UserService.GetProfile(ctx)
	if user == nil || user.UserId <= 0 {
		response.ErrorResp(ctx).SetMsg("登陆超时").Log("菜单管理", gin.H{"roleId": roleId}).WriteJsonExit()
		return
	}

	result, err := service.MenuService.RoleMenuTreeData(roleId, user.UserId)

	if err != nil {
		response.ErrorResp(ctx).SetMsg(err.Error()).Log("菜单管理", gin.H{"roleId": roleId}).WriteJsonExit()
		return
	}

	ctx.JSON(http.StatusOK, result)
}

//检查菜单名是否已经存在不包括自身
func (c *MenuController) CheckMenuNameUnique(ctx *gin.Context) {
	var req *model.MenuCheckMenuNameReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.Writer.WriteString("1")
		return
	}

	result := service.MenuService.CheckMenuNameUnique(req.MenuName, req.MenuId, req.ParentId)

	ctx.Writer.WriteString(result)
}

//检查菜单名是否已经存在
func (c *MenuController) CheckMenuNameUniqueAll(ctx *gin.Context) {
	var req *model.MenuCheckMenuNameALLReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.Writer.WriteString("1")
		return
	}

	result := service.MenuService.CheckMenuNameUniqueAll(req.MenuName, req.ParentId)

	ctx.Writer.WriteString(result)
}
