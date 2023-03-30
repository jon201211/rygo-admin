package response

import (
	"net/http"
	"rygo/app/model"

	"github.com/gin-gonic/gin"
)

// 通用api响应
type TableResp struct {
	t   *model.TableDataInfo
	ctx *gin.Context
}

//返回一个成功的消息体
func BuildTable(ctx *gin.Context, total int, rows interface{}) *TableResp {
	msg := model.TableDataInfo{
		Code:  0,
		Msg:   "操作成功",
		Total: total,
		Rows:  rows,
	}
	a := TableResp{
		t: &msg,
		ctx: ctx,
	}
	return &a
}

//输出json到客户端
func (resp *TableResp) WriteJsonExit() {
	resp.ctx.JSON(http.StatusOK, resp.t)
	resp.ctx.Abort()
}
