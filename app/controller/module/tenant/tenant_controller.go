// ==========================================================================
// RYGO自动生成控制器相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2021-06-29 22:21:21 +0800 CST
// 生成路径: app/controller/module/tenant/tenant_controller.go
// 生成人：rygo
// ==========================================================================
package tenant

import (
	"rygo/app/ginframe/response"
	"rygo/app/ginframe/utils/gconv"
	"rygo/app/model"
	tenantModel "rygo/app/model/module/tenant"
	tenantService "rygo/app/service/module/tenant"

	"github.com/gin-gonic/gin"
)

//列表页
func List(c *gin.Context) {
	response.BuildTpl(c, "module/tenant/list.html").WriteTpl()
}

//列表分页数据
func ListAjax(c *gin.Context) {
	req := new(tenantModel.SelectPageReq)
	//获取参数
	if err := c.ShouldBind(req); err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).Log("tenant管理", req).WriteJsonExit()
		return
	}
	rows := make([]tenantModel.SysTenant, 0)
	result, page, err := tenantService.SelectListByPage(req)

	if err == nil && len(result) > 0 {
		rows = result
	}

	response.BuildTable(c, page.Total, rows).WriteJsonExit()
}

//新增页面
func Add(c *gin.Context) {
	response.BuildTpl(c, "module/tenant/add.html").WriteTpl()
}

//新增页面保存
func AddSave(c *gin.Context) {
	req := new(tenantModel.AddReq)
	//获取参数
	if err := c.ShouldBind(req); err != nil {
		response.ErrorResp(c).SetBtype(model.Buniss_Add).SetMsg(err.Error()).Log("商户信息新增数据", req).WriteJsonExit()
		return
	}

	id, err := tenantService.AddSave(req, c)

	if err != nil || id <= 0 {
		response.ErrorResp(c).SetBtype(model.Buniss_Add).Log("商户信息新增数据", req).WriteJsonExit()
		return
	}
	response.SucessResp(c).SetData(id).Log("商户信息新增数据", req).WriteJsonExit()
}

//修改页面
func Edit(c *gin.Context) {
	id := gconv.Int64(c.Query("id"))

	if id <= 0 {
		response.ErrorTpl(c).WriteTpl(gin.H{
			"desc": "参数错误",
		})
		return
	}

	entity, err := tenantService.SelectRecordById(id)

	if err != nil || entity == nil {
		response.ErrorTpl(c).WriteTpl(gin.H{
			"desc": "数据不存在",
		})
		return
	}

	response.BuildTpl(c, "module/tenant/edit.html").WriteTpl(gin.H{
		"tenant": entity,
	})
}

//修改页面保存
func EditSave(c *gin.Context) {
	var req = new(tenantModel.EditReq)
	//获取参数
	if err := c.ShouldBind(req); err != nil {
		response.ErrorResp(c).SetBtype(model.Buniss_Edit).SetMsg(err.Error()).Log("商户信息修改数据", req).WriteJsonExit()
		return
	}

	rs, err := tenantService.EditSave(req, c)
	if err != nil || rs <= 0 {
		response.ErrorResp(c).SetBtype(model.Buniss_Edit).Log("商户信息修改数据", req).WriteJsonExit()
		return
	}
	response.SucessResp(c).SetBtype(model.Buniss_Edit).Log("商户信息修改数据", req).WriteJsonExit()
}

//删除数据
func Remove(c *gin.Context) {
	req := new(model.RemoveReq)
	//获取参数
	if err := c.ShouldBind(req); err != nil {
		response.ErrorResp(c).SetBtype(model.Buniss_Del).SetMsg(err.Error()).Log("商户信息删除数据", req).WriteJsonExit()
		return
	}

	rs := tenantService.DeleteRecordByIds(req.Ids)

	if rs > 0 {
		response.SucessResp(c).SetBtype(model.Buniss_Del).Log("商户信息删除数据", req).WriteJsonExit()
	} else {
		response.ErrorResp(c).SetBtype(model.Buniss_Del).Log("商户信息删除数据", req).WriteJsonExit()
	}
}

//导出
func Export(c *gin.Context) {
	req := new(tenantModel.SelectPageReq)
	//获取参数
	if err := c.ShouldBind(req); err != nil {
		response.ErrorResp(c).Log("商户信息导出数据", req).WriteJsonExit()
		return
	}
	url, err := tenantService.Export(req)

	if err != nil {
		response.ErrorResp(c).SetBtype(model.Buniss_Other).Log("商户信息导出数据", req).WriteJsonExit()
		return
	}
	response.SucessResp(c).SetBtype(model.Buniss_Other).SetMsg(url).WriteJsonExit()
}
