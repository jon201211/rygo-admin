package admin

import (
	"rygo/app/controller/base"
	"rygo/app/response"

	"github.com/gin-gonic/gin"
)

type IconController struct {
	base.BaseController
}

func (c *IconController) Fontawesome(ctx *gin.Context) {
	response.BuildTpl(ctx, "demo/icon/fontawesome").WriteTpl()
}

func (c *IconController) Glyphicons(ctx *gin.Context) {
	response.BuildTpl(ctx, "demo/icon/glyphicons").WriteTpl()
}
