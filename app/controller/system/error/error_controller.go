package error

import (
	"rygo/app/ginframe/response"

	"github.com/gin-gonic/gin"
)

func Unauth(c *gin.Context) {
	response.BuildTpl(c, "error/unauth").WriteTpl()
}

func Error(c *gin.Context) {
	response.BuildTpl(c, "error/500").WriteTpl()
}

func NotFound(c *gin.Context) {
	response.BuildTpl(c, "error/404").WriteTpl()
}
