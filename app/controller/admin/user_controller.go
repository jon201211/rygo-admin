package admin

import (
	"rygo/app/common/session"
	"rygo/app/controller/base"
	"rygo/app/model"

	"rygo/app/response"
	"rygo/app/service"

	"rygo/app/utils/gconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	base.BaseController
}

//用户列表页
func (c *UserController) List(ctx *gin.Context) {
	response.BuildTpl(ctx, "system/user/list").WriteTpl()
}

//用户列表分页数据
func (c *UserController) ListAjax(ctx *gin.Context) {
	var req *model.UserSelectPageReq
	//获取参数
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorResp(ctx).SetMsg(err.Error()).Log("用户管理", req).WriteJsonExit()
		return
	}
	tenantId := session.GetProfile(ctx).TenantId
	req.TenantId = tenantId
	rows := make([]model.UserListEntity, 0)
	result, page, err := service.UserService.SelectRecordList(req)

	if err == nil && len(result) > 0 {
		rows = result
	}
	response.BuildTable(ctx, page.Total, rows).WriteJsonExit()
}

//用户新增页面
func (c *UserController) Add(ctx *gin.Context) {
	var paramsRole *model.RoleSelectPageReq
	var paramsPost *model.PostSelectPageReq

	roles := make([]model.RoleEntityFlag, 0)
	posts := make([]model.PostEntityFlag, 0)

	rolesP, _ := service.RoleService.SelectRecordAll(paramsRole)

	if rolesP != nil {
		roles = rolesP
	}

	postP, _ := service.PostService.SelectListAll(paramsPost)

	if postP != nil {
		posts = postP
	}
	response.BuildTpl(ctx, "system/user/add").WriteTpl(gin.H{
		"roles": roles,
		"posts": posts,
	})
}

//保存新增用户数据
func (c *UserController) AddSave(ctx *gin.Context) {
	var req *model.UserAddReq
	//获取参数
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Add).SetMsg(err.Error()).Log("新增用户", req).WriteJsonExit()
		return
	}

	//判断登陆名是否已注册
	isHadName := service.UserService.CheckLoginName(req.LoginName)
	if isHadName {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Add).SetMsg("登陆名已经存在").Log("新增用户", req).WriteJsonExit()
		return
	}

	//判断手机号码是否已注册
	isHadPhone := service.UserService.CheckPhoneUniqueAll(req.Phonenumber)
	if isHadPhone {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Add).SetMsg("手机号码已经存在").Log("新增用户", req).WriteJsonExit()
		return
	}

	//判断邮箱是否已注册
	isHadEmail := service.UserService.CheckEmailUniqueAll(req.Email)
	if isHadEmail {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Add).SetMsg("邮箱已经存在").Log("新增用户", req).WriteJsonExit()
		return
	}

	uid, err := service.UserService.AddSave(req, ctx)

	if err != nil || uid <= 0 {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Add).Log("新增用户", req).WriteJsonExit()
		return
	}
	response.SucessResp(ctx).SetData(uid).SetBtype(model.Buniss_Add).Log("新增用户", req).WriteJsonExit()
}

//用户修改页面
func (c *UserController) Edit(ctx *gin.Context) {
	id := gconv.Int64(ctx.Query("id"))

	if id <= 0 {
		response.BuildTpl(ctx, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "参数错误",
		})
		return
	}

	user, err := service.UserService.SelectRecordById(id)

	if err != nil || user == nil {
		response.BuildTpl(ctx, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "用户不存在",
		})
		return
	}

	//获取部门信息
	deptName := ""
	if user.DeptId > 0 {
		dept := service.DeptService.SelectDeptById(user.DeptId)
		if dept != nil {
			deptName = dept.DeptName
		}
	}

	roles := make([]model.RoleEntityFlag, 0)
	posts := make([]model.PostEntityFlag, 0)

	rolesP, _ := service.RoleService.SelectRoleContactVo(id)

	if rolesP != nil {
		roles = rolesP
	}

	postP, _ := service.PostService.SelectPostsByUserId(id)

	if postP != nil {
		posts = postP
	}

	response.BuildTpl(ctx, "system/user/edit").WriteTpl(gin.H{
		"user":     user,
		"deptName": deptName,
		"roles":    roles,
		"posts":    posts,
	})
}

//重置密码
func (c *UserController) ResetPwd(ctx *gin.Context) {
	id := gconv.Int64(ctx.Query("userId"))
	if id <= 0 {
		response.BuildTpl(ctx, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "参数错误",
		})
		return
	}

	user, err := service.UserService.SelectRecordById(id)

	if err != nil || user == nil {
		response.BuildTpl(ctx, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "用户不存在",
		})
		return
	}
	response.BuildTpl(ctx, "system/user/resetPwd").WriteTpl(gin.H{
		"user": user,
	})
}

//重置密码保存
func (c *UserController) ResetPwdSave(ctx *gin.Context) {
	var req *model.UserResetPwdReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Edit).SetMsg(err.Error()).Log("重置密码", req).WriteJsonExit()
	}

	result, err := service.UserService.ResetPassword(req)

	if err != nil || !result {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Edit).SetMsg(err.Error()).Log("重置密码", req).WriteJsonExit()
	} else {
		response.SucessResp(ctx).SetBtype(model.Buniss_Edit).Log("重置密码", req).WriteJsonExit()
	}
}

//保存修改用户数据
func (c *UserController) EditSave(ctx *gin.Context) {
	var req *model.UserEditReq
	//获取参数
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Edit).SetMsg(err.Error()).Log("修改用户", req).WriteJsonExit()
		return
	}

	//判断手机号码是否已注册
	isHadPhone := service.UserService.CheckPhoneUnique(req.UserId, req.Phonenumber)
	if isHadPhone {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Edit).SetMsg("手机号码已经存在").Log("修改用户", req).WriteJsonExit()
		return
	}

	//判断邮箱是否已注册
	isHadEmail := service.UserService.CheckEmailUnique(req.UserId, req.Email)
	if isHadEmail {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Edit).SetMsg("邮箱已经存在").Log("修改用户", req).WriteJsonExit()
		return
	}

	uid, err := service.UserService.EditSave(req, ctx)

	if err != nil || uid <= 0 {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Edit).Log("修改用户", req).WriteJsonExit()
		return
	}

	response.SucessResp(ctx).SetData(uid).SetBtype(model.Buniss_Edit).Log("修改用户", req).WriteJsonExit()
}

//删除数据
func (c *UserController) Remove(ctx *gin.Context) {
	var req *model.RemoveReq
	//获取参数
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Del).SetMsg(err.Error()).Log("删除用户", req).WriteJsonExit()
	}

	rs := service.UserService.DeleteRecordByIds(req.Ids)

	if rs > 0 {
		response.SucessResp(ctx).SetData(rs).SetBtype(model.Buniss_Del).Log("删除用户", req).WriteJsonExit()
	} else {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Del).Log("删除用户", req).WriteJsonExit()
	}
}

//导出
func (c *UserController) Export(ctx *gin.Context) {
	var req *model.UserSelectPageReq
	//获取参数
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorResp(ctx).SetMsg(err.Error()).Log("导出Excel", req).WriteJsonExit()
	}
	url, err := service.UserService.Export(req)

	if err != nil {
		response.ErrorResp(ctx).SetMsg(err.Error()).Log("导出Excel", req).WriteJsonExit()
		return
	}
	response.SucessResp(ctx).SetMsg(url).Log("导出Excel", req).WriteJsonExit()
}
