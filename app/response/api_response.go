package response

import (
	"encoding/json"
	"net/http"
	"rygo/app/model"
	"rygo/app/service"

	"rygo/app/utils/gconv"

	"github.com/gin-gonic/gin"
)

// 通用api响应
type ApiResp struct {
	r   *model.CommonRes
	ctx *gin.Context
}

//返回一个成功的消息体
func SucessResp(ctx *gin.Context) *ApiResp {
	msg := model.CommonRes{
		Code:  0,
		Btype: model.Buniss_Other,
		Msg:   "操作成功",
	}
	var a = ApiResp{
		r:   &msg,
		ctx: ctx,
	}
	return &a
}

//返回一个错误的消息体
func ErrorResp(ctx *gin.Context) *ApiResp {
	msg := model.CommonRes{
		Code:  500,
		Btype: model.Buniss_Other,
		Msg:   "操作失败",
	}
	var a = ApiResp{
		r:   &msg,
		ctx: ctx,
	}
	return &a
}

//返回一个拒绝访问的消息体
func ForbiddenResp(ctx *gin.Context) *ApiResp {
	msg := model.CommonRes{
		Code:  403,
		Btype: model.Buniss_Other,
		Msg:   "无操作权限",
	}
	var a = ApiResp{
		r:   &msg,
		ctx: ctx,
	}
	return &a
}

//设置消息体的内容
func (resp *ApiResp) SetMsg(msg string) *ApiResp {
	resp.r.Msg = msg
	return resp
}

//设置消息体的编码
func (resp *ApiResp) SetCode(code int) *ApiResp {
	resp.r.Code = code
	return resp
}

//设置消息体的数据
func (resp *ApiResp) SetData(data interface{}) *ApiResp {
	resp.r.Data = data
	return resp
}

//设置消息体的业务类型
func (resp *ApiResp) SetBtype(btype model.BunissType) *ApiResp {
	resp.r.Btype = btype
	return resp
}

//记录操作日志到数据库
func (resp *ApiResp) Log(title string, inParam interface{}) *ApiResp {
	var inContentStr string
	switch inParam.(type) {
	case string, []byte:
		inContentStr = gconv.String(inParam)
	}
	// Else use json.Marshal function to encode the parameter.
	if b, err := json.Marshal(inParam); err != nil {
		inContentStr = ""
	} else {
		inContentStr = string(b)
	}
	service.OperlogService.Add(resp.ctx, title, inContentStr, resp.r)
	return resp
}

//输出json到客户端
func (resp *ApiResp) WriteJsonExit() {
	resp.ctx.JSON(http.StatusOK, resp.r)
	resp.ctx.Abort()
}
