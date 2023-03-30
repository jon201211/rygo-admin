package admin

import (
	"rygo/app/controller/base"
	"rygo/app/response"

	"github.com/gin-gonic/gin"
)

type ReportController struct {
	base.BaseController
}

func (c *ReportController) Echarts(ctx *gin.Context) {
	response.BuildTpl(ctx, "demo/report/echarts").WriteTpl()
}

func (c *ReportController) Metrics(ctx *gin.Context) {
	response.BuildTpl(ctx, "demo/report/metrics").WriteTpl()
}

func (c *ReportController) Peity(ctx *gin.Context) {
	response.BuildTpl(ctx, "demo/report/peity").WriteTpl()
}

func (c *ReportController) Sparkline(ctx *gin.Context) {
	response.BuildTpl(ctx, "demo/report/sparkline").WriteTpl()
}
