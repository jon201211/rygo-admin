package controller

import (
	"rygo/app/controller/admin"
	"rygo/app/middleware/auth"
	"rygo/app/middleware/jwt"
	"rygo/app/router"
)

//加载路由
func init() {
	// 服务监控
	gen := &admin.GenController{}
	tool := &admin.ToolController{}
	g1 := router.New("admin", "/tool", jwt.JWTAuthMiddleware(), auth.Auth)
	g1.GET("/build", "tool:build:view", tool.Build)
	g1.GET("/swagger", "tool:swagger:view", tool.Swagger)
	g1.GET("/gen", "tool:gen:view", gen.Gen)
	g1.POST("/gen/list", "tool:gen:list", gen.GenList)
	g1.POST("/gen/remove", "tool:gen:remove", gen.Remove)
	g1.GET("/gen/importTable", "tool:gen:list", gen.ImportTable)
	g1.POST("/gen/db/list", "tool:gen:list", gen.DataList)
	g1.POST("/gen/importTable", "tool:gen:list", gen.ImportTableSave)
	g1.GET("/gen/edit", "tool:gen:edit", gen.Edit)
	g1.POST("/gen/edit", "tool:gen:edit", gen.EditSave)
	g1.POST("/gen/column/list", "tool:gen:list", gen.ColumnList)
	g1.GET("/gen/preview", "tool:gen:preview", gen.Preview)
	g1.GET("/gen/genCode", "tool:gen:code", gen.GenCode)
}
