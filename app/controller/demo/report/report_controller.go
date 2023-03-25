package report

import (
	"rygo/app/ginframe/response"

	"github.com/gin-gonic/gin"
)

func Echarts(c *gin.Context) {
	response.BuildTpl(c, "demo/report/echarts").WriteTpl()
}

func Metrics(c *gin.Context) {
	response.BuildTpl(c, "demo/report/metrics").WriteTpl()
}

func Peity(c *gin.Context) {
	response.BuildTpl(c, "demo/report/peity").WriteTpl()
}

func Sparkline(c *gin.Context) {
	response.BuildTpl(c, "demo/report/sparkline").WriteTpl()
}
