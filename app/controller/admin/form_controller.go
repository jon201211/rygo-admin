package admin

import (
	"rygo/app/controller/base"
	"rygo/app/response"

	"github.com/gin-gonic/gin"
)

type FormController struct {
	base.BaseController
}

func (c *FormController) Autocomplete(ctx *gin.Context) {

	response.BuildTpl(ctx, "demo/form/autocomplete").WriteTpl()
}

func (c *FormController) Basic(ctx *gin.Context) {
	response.BuildTpl(ctx, "demo/form/basic").WriteTpl()
}

func (c *FormController) Button(ctx *gin.Context) {
	response.BuildTpl(ctx, "demo/form/button").WriteTpl()
}

func (c *FormController) Cards(ctx *gin.Context) {
	response.BuildTpl(ctx, "demo/form/cards").WriteTpl()
}

func (c *FormController) Datetime(ctx *gin.Context) {
	response.BuildTpl(ctx, "demo/form/datetime").WriteTpl()
}

func (c *FormController) Duallistbox(ctx *gin.Context) {
	response.BuildTpl(ctx, "demo/form/duallistbox").WriteTpl()
}

func (c *FormController) Grid(ctx *gin.Context) {
	response.BuildTpl(ctx, "demo/form/grid").WriteTpl()
}

func (c *FormController) Jasny(ctx *gin.Context) {
	response.BuildTpl(ctx, "demo/form/jasny").WriteTpl()
}

func (c *FormController) Select(ctx *gin.Context) {
	response.BuildTpl(ctx, "demo/form/select").WriteTpl()
}

func (c *FormController) Sortable(ctx *gin.Context) {
	response.BuildTpl(ctx, "demo/form/sortable").WriteTpl()
}

func (c *FormController) Summernote(ctx *gin.Context) {
	response.BuildTpl(ctx, "demo/form/summernote").WriteTpl()
}

func (c *FormController) Tabs_panels(ctx *gin.Context) {
	response.BuildTpl(ctx, "demo/form/tabs_panels").WriteTpl()
}

func (c *FormController) Timeline(ctx *gin.Context) {
	response.BuildTpl(ctx, "demo/form/timeline").WriteTpl()
}

func (c *FormController) Upload(ctx *gin.Context) {
	response.BuildTpl(ctx, "demo/form/upload").WriteTpl()
}

func (c *FormController) Validate(ctx *gin.Context) {
	response.BuildTpl(ctx, "demo/form/validate").WriteTpl()
}

func (c *FormController) Wizard(ctx *gin.Context) {
	response.BuildTpl(ctx, "demo/form/wizard").WriteTpl()
}
