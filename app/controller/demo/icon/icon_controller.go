package icon

import (
	"rygo/app/ginframe/response"

	"github.com/gin-gonic/gin"
)

func Fontawesome(c *gin.Context) {
	response.BuildTpl(c, "demo/icon/fontawesome").WriteTpl()
}

func Glyphicons(c *gin.Context) {
	response.BuildTpl(c, "demo/icon/glyphicons").WriteTpl()
}
