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

type RoleController struct {
	base.BaseController
}

//列表页
func (c *RoleController) List(ctx *gin.Context) {
	response.BuildTpl(ctx, "system/role/list").WriteTpl()
}

//列表分页数据
func (c *RoleController) ListAjax(ctx *gin.Context) {
	var req *model.RoleSelectPageReq
	//获取参数
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorResp(ctx).SetMsg(err.Error()).Log("角色管理", req).WriteJsonExit()
		return
	}
	rows := make([]model.RoleEntity, 0)
	result, page, err := service.RoleService.SelectRecordPage(req)

	if err == nil && len(result) > 0 {
		rows = result
	}
	response.BuildTable(ctx, page.Total, rows).WriteJsonExit()
}

//新增页面
func (c *RoleController) Add(ctx *gin.Context) {
	response.BuildTpl(ctx, "system/role/add").WriteTpl()
}

//新增页面保存
func (c *RoleController) AddSave(ctx *gin.Context) {
	var req *model.RoleAddReq
	//获取参数
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Add).SetMsg(err.Error()).Log("角色管理", req).WriteJsonExit()
		return
	}

	if service.RoleService.CheckRoleNameUniqueAll(req.RoleName) == "1" {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Add).SetMsg("角色名称已存在").Log("角色管理", req).WriteJsonExit()
		return
	}

	if service.RoleService.CheckRoleKeyUniqueAll(req.RoleKey) == "1" {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Add).SetMsg("角色权限已存在").Log("角色管理", req).WriteJsonExit()
		return
	}

	rid, err := service.RoleService.AddSave(req, ctx)

	if err != nil || rid <= 0 {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Add).Log("角色管理", req).WriteJsonExit()
		return
	}
	response.SucessResp(ctx).SetData(rid).SetBtype(model.Buniss_Add).Log("角色管理", req).WriteJsonExit()
}

//修改页面
func (c *RoleController) Edit(ctx *gin.Context) {
	id := gconv.Int64(ctx.Query("id"))
	if id <= 0 {
		response.BuildTpl(ctx, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "参数错误",
		})
		return
	}

	role, err := service.RoleService.SelectRecordById(id)

	if err != nil || role == nil {
		response.BuildTpl(ctx, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "角色不存在",
		})
		return
	}

	response.BuildTpl(ctx, "system/role/edit").WriteTpl(gin.H{
		"role": role,
	})
}

//修改页面保存
func (c *RoleController) EditSave(ctx *gin.Context) {
	var req *model.RoleEditReq
	//获取参数
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Edit).SetMsg(err.Error()).Log("角色管理", req).WriteJsonExit()
		return
	}

	if service.RoleService.CheckRoleNameUnique(req.RoleName, req.RoleId) == "1" {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Edit).SetMsg("角色名称已存在").Log("角色管理", req).WriteJsonExit()
		return
	}

	if service.RoleService.CheckRoleKeyUnique(req.RoleKey, req.RoleId) == "1" {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Edit).SetMsg("角色权限已存在").Log("角色管理", req).WriteJsonExit()
		return
	}

	rs, err := service.RoleService.EditSave(req, ctx)

	if err != nil || rs <= 0 {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Edit).Log("角色管理", req).WriteJsonExit()
		return
	}
	response.SucessResp(ctx).SetBtype(model.Buniss_Edit).SetData(rs).Log("角色管理", req).WriteJsonExit()
}

//分配用户添加
func (c *RoleController) SelectUser(ctx *gin.Context) {
	id := gconv.Int64(ctx.Query("id"))
	if id <= 0 {
		response.BuildTpl(ctx, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "参数错误",
		})
		return
	}

	role, err := service.RoleService.SelectRecordById(id)

	if err != nil {
		response.BuildTpl(ctx, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "角色不存在",
		})
	} else {
		response.BuildTpl(ctx, "system/role/selectUser").WriteTpl(gin.H{
			"role": role,
		})
	}
}

//获取用户列表
func (c *RoleController) UnallocatedList(ctx *gin.Context) {
	roleId := gconv.Int64(ctx.PostForm("roleId"))
	loginName := ctx.PostForm("loginName")
	phonenumber := ctx.PostForm("phonenumber")
	var rows []model.SysUser
	userList, err := service.UserService.SelectUnallocatedList(roleId, loginName, phonenumber)

	if err == nil && userList != nil {
		rows = userList
	}

	ctx.JSON(http.StatusOK, model.TableDataInfo{
		Code:  0,
		Msg:   "操作成功",
		Total: len(rows),
		Rows:  rows,
	})
}

//删除数据
func (c *RoleController) Remove(ctx *gin.Context) {
	var req *model.RemoveReq
	//获取参数
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Del).SetMsg(err.Error()).Log("角色管理", req).WriteJsonExit()
		return
	}

	rs := service.RoleService.DeleteRecordByIds(req.Ids)

	if rs > 0 {
		response.SucessResp(ctx).SetBtype(model.Buniss_Del).SetData(rs).Log("角色管理", req).WriteJsonExit()
	} else {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Del).Log("角色管理", req).WriteJsonExit()
	}
}

//导出
func (c *RoleController) Export(ctx *gin.Context) {
	var req *model.RoleSelectPageReq
	//获取参数
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorResp(ctx).SetMsg(err.Error()).Log("角色管理", req).WriteJsonExit()
		return
	}
	url, err := service.RoleService.Export(req)

	if err != nil {
		response.ErrorResp(ctx).SetMsg(err.Error()).Log("角色管理", req).WriteJsonExit()
		return
	}
	response.SucessResp(ctx).SetMsg(url).Log("角色管理", req).WriteJsonExit()
}

//数据权限
func (c *RoleController) AuthDataScope(ctx *gin.Context) {
	roleId := gconv.Int64(ctx.Query("id"))
	role, err := service.RoleService.SelectRecordById(roleId)
	if err != nil {
		response.BuildTpl(ctx, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "角色不存在",
		})
	} else {
		response.BuildTpl(ctx, "system/role/dataScope").WriteTpl(gin.H{
			"role": role,
		})
	}
}

//数据权限保存
func (c *RoleController) AuthDataScopeSave(ctx *gin.Context) {
	var req *model.RoleDataScopeReq
	//获取参数
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorResp(ctx).SetMsg(err.Error()).Log("角色管理", req).WriteJsonExit()
		return
	}
	if !service.RoleService.CheckRoleAllowed(req.RoleId) {
		response.ErrorResp(ctx).SetMsg("不允许操作超级管理员角色").Log("角色管理", req).WriteJsonExit()
		return
	}

	rs, err := service.RoleService.AuthDataScope(req, ctx)
	if err != nil || rs <= 0 {
		response.ErrorResp(ctx).SetMsg("保存数据失败").SetMsg(err.Error()).Log("角色管理", req).WriteJsonExit()
	} else {
		response.SucessResp(ctx).Log("角色管理", req).WriteJsonExit()
	}
}

//分配用户
func (c *RoleController) AuthUser(ctx *gin.Context) {
	roleId := gconv.Int64(ctx.Query("id"))
	role, err := service.RoleService.SelectRecordById(roleId)
	if err != nil {
		response.BuildTpl(ctx, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "角色不存在",
		})
	} else {
		response.BuildTpl(ctx, "system/role/authUser").WriteTpl(gin.H{
			"role": role,
		})
	}
}

//查询已分配用户角色列表
func (c *RoleController) AllocatedList(ctx *gin.Context) {
	roleId := gconv.Int64(ctx.PostForm("roleId"))
	loginName := ctx.PostForm("loginName")
	phonenumber := ctx.PostForm("phonenumber")
	var rows []model.SysUser
	userList, err := service.UserService.SelectAllocatedList(roleId, loginName, phonenumber)

	if err == nil && userList != nil {
		rows = userList
	}

	ctx.JSON(http.StatusOK, model.TableDataInfo{
		Code:  0,
		Msg:   "操作成功",
		Total: len(rows),
		Rows:  rows,
	})
}

//保存角色选择
func (c *RoleController) SelectAll(ctx *gin.Context) {
	roleId := gconv.Int64(ctx.PostForm("roleId"))
	userIds := ctx.PostForm("userIds")

	if roleId <= 0 {
		response.ErrorResp(ctx).SetMsg("参数错误1").SetBtype(model.Buniss_Add).Log("角色管理", gin.H{
			"roleId":  roleId,
			"userIds": userIds,
		}).WriteJsonExit()
		return
	}
	if userIds == "" {
		response.ErrorResp(ctx).SetMsg("参数错误2").SetBtype(model.Buniss_Add).Log("角色管理", gin.H{
			"roleId":  roleId,
			"userIds": userIds,
		}).WriteJsonExit()
		return
	}

	rs := service.RoleService.InsertAuthUsers(roleId, userIds)
	if rs > 0 {
		response.SucessResp(ctx).SetBtype(model.Buniss_Add).Log("角色管理", gin.H{
			"roleId":  roleId,
			"userIds": userIds,
		}).WriteJsonExit()
	} else {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Add).Log("角色管理", gin.H{
			"roleId":  roleId,
			"userIds": userIds,
		}).WriteJsonExit()
	}

}

//取消用户角色授权
func (c *RoleController) CancelAll(ctx *gin.Context) {
	roleId := gconv.Int64(ctx.PostForm("roleId"))
	userIds := ctx.PostForm("userIds")
	if roleId > 0 && userIds != "" {
		service.RoleService.DeleteUserRoleInfos(roleId, userIds)
		response.SucessResp(ctx).SetBtype(model.Buniss_Del).Log("角色管理", gin.H{
			"roleId":  roleId,
			"userIds": userIds,
		}).WriteJsonExit()
	} else {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Del).SetMsg("参数错误").Log("角色管理", gin.H{
			"roleId":  roleId,
			"userIds": userIds,
		}).WriteJsonExit()
	}
}

//批量取消用户角色授权
func (c *RoleController) Cancel(ctx *gin.Context) {
	roleId := gconv.Int64(ctx.PostForm("roleId"))
	userId := gconv.Int64(ctx.PostForm("userId"))
	if roleId > 0 && userId > 0 {
		service.RoleService.DeleteUserRoleInfo(userId, roleId)
		response.SucessResp(ctx).SetBtype(model.Buniss_Del).Log("角色管理", gin.H{
			"roleId": roleId,
			"userId": userId,
		}).WriteJsonExit()
	} else {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Del).SetMsg("参数错误").Log("角色管理", gin.H{
			"roleId": roleId,
			"userId": userId,
		}).WriteJsonExit()
	}
}

//检查角色是否已经存在不包括本角色
func (c *RoleController) CheckRoleNameUnique(ctx *gin.Context) {
	var req *model.RoleCheckRoleNameReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.Writer.WriteString("1")
		return
	}

	result := service.RoleService.CheckRoleNameUnique(req.RoleName, req.RoleId)

	ctx.Writer.WriteString(result)
}

//检查角色是否已经存在
func (c *RoleController) CheckRoleNameUniqueAll(ctx *gin.Context) {
	var req *model.RoleCheckRoleNameALLReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.Writer.WriteString("1")
		return
	}

	result := service.RoleService.CheckRoleNameUniqueAll(req.RoleName)

	ctx.Writer.WriteString(result)
}

//检查角色是否已经存在不包括本角色
func (c *RoleController) CheckRoleKeyUnique(ctx *gin.Context) {
	var req *model.RoleCheckRoleKeyReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.Writer.WriteString("1")
		return
	}

	result := service.RoleService.CheckRoleKeyUnique(req.RoleKey, req.RoleId)

	ctx.Writer.WriteString(result)
}

//检查角色是否已经存在
func (c *RoleController) CheckRoleKeyUniqueAll(ctx *gin.Context) {
	var req *model.RoleCheckRoleKeyALLReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.Writer.WriteString("1")
		return
	}

	result := service.RoleService.CheckRoleKeyUniqueAll(req.RoleKey)

	ctx.Writer.WriteString(result)
}
