package admin

import (
	"rygo/app/controller/base"
	"rygo/app/response"

	"github.com/gin-gonic/gin"
)

type ErrorController struct {
	base.BaseController
}

func (c *ErrorController) Unauth(ctx *gin.Context) {
	response.BuildTpl(ctx, "error/unauth").WriteTpl()
}

func (c *ErrorController) Error(ctx *gin.Context) {
	response.BuildTpl(ctx, "error/500").WriteTpl()
}

func (c *ErrorController) NotFound(ctx *gin.Context) {
	response.BuildTpl(ctx, "error/404").WriteTpl()
}
