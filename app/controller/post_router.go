package controller

import (
	"rygo/app/controller/admin"
	"rygo/app/middleware/auth"
	"rygo/app/middleware/jwt"
	"rygo/app/router"
)

//加载路由
func init() {
	// 岗位路由
	post := &admin.PostController{}
	g1 := router.New("admin", "/system/post", jwt.JWTAuthMiddleware(), auth.Auth)

	g1.GET("/", "system:post:view", post.List)
	g1.POST("/list", "system:post:list", post.ListAjax)
	g1.GET("/add", "system:post:add", post.Add)
	g1.POST("/add", "system:post:add", post.AddSave)
	g1.POST("/remove", "system:post:remove", post.Remove)
	g1.GET("/edit", "system:post:edit", post.Edit)
	g1.POST("/edit", "system:post:edit", post.EditSave)
	g1.POST("/export", "system:post:export", post.Export)
	g1.POST("/checkPostCodeUniqueAll", "system:post:list", post.CheckPostCodeUniqueAll)
	g1.POST("/checkPostCodeUnique", "system:post:list", post.CheckPostCodeUnique)
	g1.POST("/checkPostNameUniqueAll", "system:post:list", post.CheckPostNameUniqueAll)
	g1.POST("/checkPostNameUnique", "system:post:list", post.CheckPostNameUnique)
}
