package admin

import (
	"net/http"
	"rygo/app/common/session"
	"rygo/app/controller/base"
	"rygo/app/model"

	"rygo/app/response"
	"rygo/app/service"

	"rygo/app/utils/gconv"

	"github.com/gin-gonic/gin"
)

type DeptController struct {
	base.BaseController
}

//列表页
func (c *DeptController) List(ctx *gin.Context) {
	response.BuildTpl(ctx, "system/dept/list").WriteTpl()
}

//列表分页数据
func (c *DeptController) ListAjax(ctx *gin.Context) {
	var req = model.DeptSelectPageReq{}
	//获取参数
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorResp(ctx).SetMsg(err.Error()).Log("部门管理", req).WriteJsonExit()
		return
	}
	profile := session.GetProfile(ctx)
	req.TenantId = profile.TenantId

	rows := make([]model.SysDept, 0)
	result, err := service.DeptService.SelectListAll(&req)

	if err == nil && len(result) > 0 {
		rows = result
	}

	ctx.JSON(http.StatusOK, rows)
}

//新增页面
func (c *DeptController) Add(ctx *gin.Context) {
	pid := gconv.Int64(ctx.Query("pid"))

	if pid == 0 {
		pid = 100
	}

	tmp := service.DeptService.SelectDeptById(pid)

	response.BuildTpl(ctx, "system/dept/add").WriteTpl(gin.H{"dept": tmp})
}

//新增页面保存
func (c *DeptController) AddSave(ctx *gin.Context) {
	var req *model.DeptAddReq
	//获取参数
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Add).SetMsg(err.Error()).Log("部门管理", req).WriteJsonExit()
		return
	}

	if service.DeptService.CheckDeptNameUniqueAll(req.DeptName, req.ParentId) == "1" {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Add).SetMsg("部门名称已存在").Log("部门管理", req).WriteJsonExit()
		return
	}
	user := session.GetProfile(ctx)
	isAdmin := session.IsAdminUser(user)
	if isAdmin == false { //非管理员，以当前用户所属租户为限
		req.TenantId = user.TenantId
	}
	rid, err := service.DeptService.AddSave(req, ctx)

	if err != nil || rid <= 0 {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Add).Log("部门管理", req).WriteJsonExit()
		return
	}
	response.SucessResp(ctx).SetBtype(model.Buniss_Add).Log("部门管理", req).WriteJsonExit()
}

//修改页面
func (c *DeptController) Edit(ctx *gin.Context) {
	id := gconv.Int64(ctx.Query("id"))
	if id <= 0 {
		response.BuildTpl(ctx, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "参数错误",
		})
		return
	}

	dept := service.DeptService.SelectDeptById(id)

	if dept == nil || dept.DeptId <= 0 {
		response.BuildTpl(ctx, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "部门不存在",
		})
		return
	}

	response.BuildTpl(ctx, "system/dept/edit").WriteTpl(gin.H{
		"dept": dept,
	})
}

//修改页面保存
func (c *DeptController) EditSave(ctx *gin.Context) {
	var req *model.DeptEditReq
	//获取参数
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Edit).SetMsg(err.Error()).Log("部门管理", req).WriteJsonExit()
		return
	}

	if service.DeptService.CheckDeptNameUnique(req.DeptName, req.DeptId, req.ParentId) == "1" {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Edit).SetMsg("部门名称已存在").Log("部门管理", req).WriteJsonExit()
		return
	}

	rs, err := service.DeptService.EditSave(req, ctx)

	if err != nil || rs <= 0 {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Edit).Log("部门管理", req).WriteJsonExit()
		return
	}
	response.SucessResp(ctx).SetData(rs).SetBtype(model.Buniss_Edit).Log("部门管理", req).WriteJsonExit()
}

//删除数据
func (c *DeptController) Remove(ctx *gin.Context) {
	id := gconv.Int64(ctx.Query("id"))
	rs := service.DeptService.DeleteDeptById(id)

	if rs > 0 {
		response.SucessResp(ctx).SetBtype(model.Buniss_Del).Log("部门管理", gin.H{"id": id}).WriteJsonExit()
	} else {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Del).Log("部门管理", gin.H{"id": id}).WriteJsonExit()
	}
}

//加载部门列表树结构的数据
func (c *DeptController) TreeData(ctx *gin.Context) {
	tenantId := session.GetTenantId(ctx)
	result, _ := service.DeptService.SelectDeptTree(0, "", "", tenantId)
	ctx.JSON(http.StatusOK, result)
}

//加载部门列表树选择页面
func (c *DeptController) SelectDeptTree(ctx *gin.Context) {
	deptId := gconv.Int64(ctx.Query("deptId"))
	deptPoint := service.DeptService.SelectDeptById(deptId)

	if deptPoint != nil {
		response.BuildTpl(ctx, "system/dept/tree").WriteTpl(gin.H{
			"dept": *deptPoint,
		})
	} else {
		response.BuildTpl(ctx, "system/dept/tree").WriteTpl()
	}
}

//加载角色部门（数据权限）列表树
func (c *DeptController) RoleDeptTreeData(ctx *gin.Context) {
	tenantId := session.GetTenantId(ctx)
	roleId := gconv.Int64(ctx.Query("roleId"))
	result, err := service.DeptService.RoleDeptTreeData(roleId, tenantId)

	if err != nil {
		response.ErrorResp(ctx).SetMsg(err.Error()).Log("菜单树", gin.H{"roleId": roleId})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

//检查部门名称是否已经存在
func (c *DeptController) CheckDeptNameUnique(ctx *gin.Context) {
	var req *model.CheckDeptNameReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.Writer.WriteString("1")
		return
	}

	result := service.DeptService.CheckDeptNameUnique(req.DeptName, req.DeptId, req.ParentId)

	ctx.Writer.WriteString(result)
}

//检查部门名称是否已经存在
func (c *DeptController) CheckDeptNameUniqueAll(ctx *gin.Context) {
	var req *model.CheckDeptNameALLReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.Writer.WriteString("1")
		return
	}

	result := service.DeptService.CheckDeptNameUniqueAll(req.DeptName, req.ParentId)

	ctx.Writer.WriteString(result)
}
