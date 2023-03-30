package admin

import (
	"rygo/app/controller/base"
	"rygo/app/global"
	"rygo/app/model"

	"rygo/app/response"
	"rygo/app/service"

	"strings"

	"github.com/gin-gonic/gin"
)

type OnlineController struct {
	base.BaseController
}

//列表页
func (c *OnlineController) List(ctx *gin.Context) {
	sessinIdArr := make([]string, 0)

	global.SessionList.Range(func(k, v interface{}) bool {
		return true
	})
	if len(sessinIdArr) > 0 {
		service.OnlineService.DeleteRecordNotInIds(sessinIdArr)
	}

	response.BuildTpl(ctx, "monitor/online/list").WriteTpl()
}

//列表分页数据
func (c *OnlineController) ListAjax(ctx *gin.Context) {
	var req *model.UserOnlineSelectPageReq
	//获取参数
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorResp(ctx).SetMsg(err.Error()).WriteJsonExit()
		return
	}
	rows := make([]model.UserOnline, 0)
	result, page, err := service.OnlineService.SelectListByPage(req)

	if err == nil && len(result) > 0 {
		rows = result
	}

	response.BuildTable(ctx, page.Total, rows).WriteJsonExit()
}

//用户强退
func (c *OnlineController) ForceLogout(ctx *gin.Context) {
	sessionId := ctx.PostForm("sessionId")
	if sessionId == "" {
		response.ErrorResp(ctx).SetMsg("参数错误").Log("用户强退", gin.H{"sessionId": sessionId}).WriteJsonExit()
		return
	}

	err := service.UserService.ForceLogout(sessionId)
	if err != nil {
		response.ErrorResp(ctx).SetMsg(err.Error()).Log("用户强退", gin.H{"sessionId": sessionId}).WriteJsonExit()
		return
	}
	response.SucessResp(ctx).Log("用户强退", gin.H{"sessionId": sessionId}).WriteJsonExit()
}

//批量强退
func (c *OnlineController) BatchForceLogout(ctx *gin.Context) {
	ids := ctx.Query("ids")
	if ids == "" {
		response.ErrorResp(ctx).SetMsg("参数错误").Log("批量强退", gin.H{"ids": ids}).WriteJsonExit()
		return
	}
	ids = strings.ReplaceAll(ids, "[", "")
	ids = strings.ReplaceAll(ids, "]", "")
	ids = strings.ReplaceAll(ids, `"`, "")
	idarr := strings.Split(ids, ",")
	if len(idarr) > 0 {
		for _, sessionId := range idarr {
			if sessionId != "" {
				service.UserService.ForceLogout(sessionId)
			}
		}
	}
	response.SucessResp(ctx).Log("批量强退", gin.H{"ids": ids}).WriteJsonExit()
}
