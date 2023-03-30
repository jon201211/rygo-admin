package controller

import (
	"rygo/app/controller/admin"
	"rygo/app/middleware/auth"
	"rygo/app/middleware/jwt"
	"rygo/app/router"
)

//加载路由
func init() {
	// 参数路由
	tenant := &admin.TenantController{}
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
