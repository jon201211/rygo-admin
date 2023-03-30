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
	dict_data := &admin.DictDataController{}
	g1 := router.New("admin", "/system/dict/data", jwt.JWTAuthMiddleware(), auth.Auth)
	g1.POST("/list", "system:dict:view", dict_data.ListAjax)
	g1.GET("/add", "system:dict:add", dict_data.Add)
	g1.POST("/add", "system:dict:add", dict_data.AddSave)
	g1.POST("/remove", "system:dict:remove", dict_data.Remove)
	g1.GET("/edit", "system:dict:edit", dict_data.Edit)
	g1.POST("/edit", "system:dict:edit", dict_data.EditSave)
	g1.POST("/export", "system:dict:export", dict_data.Export)
}
