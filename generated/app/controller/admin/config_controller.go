/* ==========================================================================
 RYGO自动生成控制器相关代码，只生成一次，按需修改,再次生成不会覆盖.
 生成日期：2020-03-27 04:35:17 +0800 CST
 ==========================================================================*/
package config

import (
	"github.com/gin-gonic/gin"
	"yj-app/app/model"
	configModel "yj-app/app/model/config"
	"yj-app/app/response"
	"yj-app/app/utils/gconv"
)

type ConfigController struct {
	base.BaseController
}

//列表页
func (c* ConfigController) List(ctx *gin.Context) {
	response.BuildTpl(c, "config/list.html").WriteTpl()
}

//列表分页数据
func (c* ConfigController) ListAjax(ctx *gin.Context) {
	req := new(configModel.SelectPageReq)
	//获取参数
	if err := c.ShouldBind(req); err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).Log("config管理", req).WriteJsonExit()
		return
	}
	rows := make([]configModel.config, 0)
	result, page, err := service.configService.SelectListByPage(req)

	if err == nil && len(result) > 0 {
		rows = result
	}

	response.BuildTable(c, page.Total, rows).WriteJsonExit()
}

//新增页面
func (c* ConfigController) Add(ctx *gin.Context) {
	response.BuildTpl(c, "config/add.html").WriteTpl()
}

//新增页面保存
func (c* ConfigController) AddSave(ctx *gin.Context) {
	req := new(configModel.AddReq)
	//获取参数
	if err := c.ShouldBind(req); err != nil {
		response.ErrorResp(c).SetBtype(model.Buniss_Add).SetMsg(err.Error()).Log("参数配置新增数据", req).WriteJsonExit()
		return
	}

	id, err := service.configService.AddSave(req, c)

	if err != nil || id <= 0 {
		response.ErrorResp(c).SetBtype(model.Buniss_Add).Log("参数配置新增数据", req).WriteJsonExit()
		return
	}
	response.SucessResp(c).SetData(id).Log("参数配置新增数据", req).WriteJsonExit()
}

//修改页面
func (c* ConfigController) Edit(ctx *gin.Context) {
	id := gconv.Int64(c.Query("id"))

	if id <= 0 {
		response.ErrorTpl(c).WriteTpl(gin.H{
			"desc": "参数错误",
		})
		return
	}

	entity, err := service.configService.SelectRecordById(id)

	if err != nil || entity == nil {
		response.ErrorTpl(c).WriteTpl(gin.H{
			"desc": "数据不存在",
		})
		return
	}

	response.BuildTpl(c, "config/edit.html").WriteTpl(gin.H{
		"config": entity,
	})
}

//修改页面保存
func (c* ConfigController) EditSave(ctx *gin.Context) {
	req := new(configModel.EditReq)
	//获取参数
	if err := c.ShouldBind(req); err != nil {
		response.ErrorResp(c).SetBtype(model.Buniss_Edit).SetMsg(err.Error()).Log("参数配置修改数据", req).WriteJsonExit()
		return
	}

	rs, err := service.configService.EditSave(req, c)

	if err != nil || rs <= 0 {
		response.ErrorResp(c).SetBtype(model.Buniss_Edit).Log("参数配置修改数据", req).WriteJsonExit()
		return
	}
	response.SucessResp(c).SetBtype(model.Buniss_Edit).Log("参数配置修改数据", req).WriteJsonExit()
}

//删除数据
func (c* ConfigController) Remove(ctx *gin.Context) {
	req := new(model.RemoveReq)
	//获取参数
	if err := c.ShouldBind(req); err != nil {
		response.ErrorResp(c).SetBtype(model.Buniss_Del).SetMsg(err.Error()).Log("参数配置删除数据", req).WriteJsonExit()
		return
	}

	rs := service.configService.DeleteRecordByIds(req.Ids)

	if rs > 0 {
		response.SucessResp(c).SetBtype(model.Buniss_Del).Log("参数配置删除数据", req).WriteJsonExit()
	} else {
		response.ErrorResp(c).SetBtype(model.Buniss_Del).Log("参数配置删除数据", req).WriteJsonExit()
	}
}

//导出
func (c* ConfigController) Export(ctx *gin.Context) {
	req := new(configModel.SelectPageReq)
	//获取参数
	if err := c.ShouldBind(req); err != nil {
		response.ErrorResp(c).Log("参数配置导出数据", req).WriteJsonExit()
		return
	}
	url, err := service.configService.Export(req)

	if err != nil {
		response.ErrorResp(c).SetBtype(model.Buniss_Other).Log("参数配置导出数据", req).WriteJsonExit()
		return
	}
	response.SucessResp(c).SetBtype(model.Buniss_Other).SetMsg(url).WriteJsonExit()
}