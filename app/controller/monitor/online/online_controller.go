package online

import (
	"rygo/app/ginframe/response"
	"rygo/app/global"
	onlineModel "rygo/app/model/monitor/online"
	onlineService "rygo/app/service/monitor/online"
	userService "rygo/app/service/system/user"
	"strings"

	"github.com/gin-gonic/gin"
)

//列表页
func List(c *gin.Context) {
	sessinIdArr := make([]string, 0)

	global.SessionList.Range(func(k, v interface{}) bool {
		return true
	})
	if len(sessinIdArr) > 0 {
		onlineService.DeleteRecordNotInIds(sessinIdArr)
	}

	response.BuildTpl(c, "monitor/online/list").WriteTpl()
}

//列表分页数据
func ListAjax(c *gin.Context) {
	var req *onlineModel.SelectPageReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).WriteJsonExit()
		return
	}
	rows := make([]onlineModel.UserOnline, 0)
	result, page, err := onlineService.SelectListByPage(req)

	if err == nil && len(result) > 0 {
		rows = result
	}

	response.BuildTable(c, page.Total, rows).WriteJsonExit()
}

//用户强退
func ForceLogout(c *gin.Context) {
	sessionId := c.PostForm("sessionId")
	if sessionId == "" {
		response.ErrorResp(c).SetMsg("参数错误").Log("用户强退", gin.H{"sessionId": sessionId}).WriteJsonExit()
		return
	}

	err := userService.ForceLogout(sessionId)
	if err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).Log("用户强退", gin.H{"sessionId": sessionId}).WriteJsonExit()
		return
	}
	response.SucessResp(c).Log("用户强退", gin.H{"sessionId": sessionId}).WriteJsonExit()
}

//批量强退
func BatchForceLogout(c *gin.Context) {
	ids := c.Query("ids")
	if ids == "" {
		response.ErrorResp(c).SetMsg("参数错误").Log("批量强退", gin.H{"ids": ids}).WriteJsonExit()
		return
	}
	ids = strings.ReplaceAll(ids, "[", "")
	ids = strings.ReplaceAll(ids, "]", "")
	ids = strings.ReplaceAll(ids, `"`, "")
	idarr := strings.Split(ids, ",")
	if len(idarr) > 0 {
		for _, sessionId := range idarr {
			if sessionId != "" {
				userService.ForceLogout(sessionId)
			}
		}
	}
	response.SucessResp(c).Log("批量强退", gin.H{"ids": ids}).WriteJsonExit()
}
