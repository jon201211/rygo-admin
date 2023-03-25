// ==========================================================================
// RYGO自动生成路由代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2021-06-29 22:21:21 +0800 CST
// 生成路径: app/controller/module/tenant_router.go
// 生成人：rygo
// ==========================================================================
package module

import (
	"rygo/app/controller/module/tenant"
	"rygo/app/ginframe/router"
	"rygo/app/service/middleware/auth"
	"rygo/app/service/middleware/jwt"
)

//加载路由
func init() {
	// 参数路由
	g1 := router.New("admin", "/module/tenant", jwt.JWTAuthMiddleware(), auth.Auth)
	g1.GET("/", "module:tenant:view", tenant.List)
	g1.POST("/list", "module:tenant:list", tenant.ListAjax)
	g1.GET("/add", "module:tenant:add", tenant.Add)
	g1.POST("/add", "module:tenant:add", tenant.AddSave)
	g1.POST("/remove", "module:tenant:remove", tenant.Remove)
	g1.GET("/edit", "module:tenant:edit", tenant.Edit)
	g1.POST("/edit", "module:tenant:edit", tenant.EditSave)
	g1.POST("/export", "module:tenant:export", tenant.Export)
}
