package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 通用tpl响应
type TplResp struct {
	ctx   *gin.Context
	tpl string
}

//返回一个tpl响应
func BuildTpl(ctx *gin.Context, tpl string) *TplResp {
	var t = TplResp{
		ctx:   ctx,
		tpl: tpl,
	}
	return &t
}

//返回一个错误的tpl响应
func ErrorTpl(ctx *gin.Context) *TplResp {
	var t = TplResp{
		ctx:   ctx,
		tpl: "error/error.html",
	}
	return &t
}

//返回一个无操作权限tpl响应
func ForbiddenTpl(ctx *gin.Context) *TplResp {
	var t = TplResp{
		ctx:   ctx,
		tpl: "error/unauth.html",
	}
	return &t
}

//输出页面模板
func (resp *TplResp) WriteTpl(params ...gin.H) {
	//session := sessions.Default(resp.ctx)
	//uid := session.Get(model.USER_ID)
	uid := 1
	if params == nil || len(params) == 0 {
		resp.ctx.HTML(http.StatusOK, resp.tpl, gin.H{"uid": uid})
	} else {
		params[0]["uid"] = uid
		resp.ctx.HTML(http.StatusOK, resp.tpl, params[0])
	}
	//resp.ctx.Abort()
}
