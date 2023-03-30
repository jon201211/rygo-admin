/* ==========================================================================
 RYGO自动生成路由代码，只生成一次，按需修改,再次生成不会覆盖.
 生成日期：2020-03-27 04:35:17 +0800 CST
 ==========================================================================*/
package config

import (
	"yj-app/app/controller/config"
	"yj-app/app/service/middleware/auth"
	"yj-app/app/router"
)

//加载路由
func init() {
	// 参数路由
	config := &admin.ConfigController{}
	g1 := router.New("admin", "/config", auth.Auth)
	g1.GET("/", "config:view", config.List)
	g1.POST("/list", "config:list", config.ListAjax)
	g1.GET("/add", "config:add", config.Add)
	g1.POST("/add", "config:add", config.AddSave)
	g1.POST("/remove", "config:remove", config.Remove)
	g1.GET("/edit", "config:edit", config.Edit)
	g1.POST("/edit", "config:edit", config.EditSave)
	g1.POST("/export", "config:export", config.Export)
}
