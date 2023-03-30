package admin

import (
	"rygo/app/controller/base"
	"rygo/app/response"

	"github.com/gin-gonic/gin"
)

type ModalController struct {
	base.BaseController
}

func (c *ModalController) Dialog(ctx *gin.Context) {
	response.BuildTpl(ctx, "demo/modal/dialog").WriteTpl()
}

func (c *ModalController) Form(ctx *gin.Context) {
	response.BuildTpl(ctx, "demo/modal/form").WriteTpl()
}

func (c *ModalController) Layer(ctx *gin.Context) {
	response.BuildTpl(ctx, "demo/modal/layer").WriteTpl()
}

func (c *ModalController) Table(ctx *gin.Context) {
	response.BuildTpl(ctx, "demo/modal/table").WriteTpl()
}

func (c *ModalController) Check(ctx *gin.Context) {
	response.BuildTpl(ctx, "demo/modal/table/check").WriteTpl()
}

func (c *ModalController) Parent(ctx *gin.Context) {
	response.BuildTpl(ctx, "demo/modal/table/parent").WriteTpl()
}

func (c *ModalController) Radio(ctx *gin.Context) {
	response.BuildTpl(ctx, "demo/modal/table/radio").WriteTpl()
}
