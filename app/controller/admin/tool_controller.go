package admin

import (
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"rygo/app/controller/base"
	"rygo/app/model"
	"rygo/app/response"

	"github.com/gin-gonic/gin"
)

const (
	swaggoRepoPath = "github.com/swaggo/swag/cmd/swag"
	// Separator for file system.
	Separator = string(filepath.Separator)
)

type ToolController struct {
	base.BaseController
}

//表单构建
func (c *ToolController) Build(ctx *gin.Context) {
	response.BuildTpl(ctx, "tool/build").WriteTpl()
}

//swagger文档
func (c *ToolController) Swagger(ctx *gin.Context) {
	a := ctx.Query("a")
	if a == "r" {
		//重新生成文档
		curDir, err := os.Getwd()
		if err != nil {
			response.BuildTpl(ctx, model.ERROR_PAGE).WriteTpl(gin.H{
				"desc": "参数错误",
			})
			ctx.Abort()
			return
		}

		genPath := curDir + "/public/swagger"
		err = c.generateSwaggerFiles(genPath)
		if err != nil {
			response.BuildTpl(ctx, model.ERROR_PAGE).WriteTpl(gin.H{
				"desc": "参数错误",
			})
			ctx.Abort()
			return
		}
	}
	ctx.Redirect(http.StatusFound, "/static/swagger/index.html")
}

//自动生成文档 swag init -o /Volumes/File/WorkSpaces/app-yjzx/public/swagger
func (c *ToolController) generateSwaggerFiles(output string) error {

	cmd := exec.Command("swag", "init -o "+output)
	// 保证关闭输出流
	if err := cmd.Start(); err != nil { // 运行命令
		return err
	}

	return nil
}
