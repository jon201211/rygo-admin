package controller

import (
	"rygo/app/controller/admin"
	"rygo/app/middleware/auth"
	"rygo/app/middleware/jwt"
	"rygo/app/router"
)

//加载路由
func init() {
	// 分组路由注册方式
	dept := &admin.DeptController{}
	g1 := router.New("admin", "/system/dept", jwt.JWTAuthMiddleware(), auth.Auth)
	g1.GET("/", "system:dept:view", dept.List)
	g1.POST("/list", "system:dept:list", dept.ListAjax)
	g1.GET("/add", "system:dept:add", dept.Add)
	g1.POST("/add", "system:dept:add", dept.AddSave)
	g1.POST("/remove", "system:dept:remove", dept.Remove)
	g1.GET("/remove", "system:dept:remove", dept.Remove)
	g1.GET("/edit", "system:dept:edit", dept.Edit)
	g1.POST("/edit", "system:dept:edit", dept.EditSave)
	g1.POST("/checkDeptNameUnique", "system:dept:view", dept.CheckDeptNameUnique)
	g1.POST("/checkDeptNameUniqueAll", "system:dept:view", dept.CheckDeptNameUniqueAll)
	g1.GET("/treeData", "system:dept:view", dept.TreeData)
	g1.GET("/selectDeptTree", "system:dept:view", dept.SelectDeptTree)
	g1.GET("/roleDeptTreeData", "system:dept:view", dept.RoleDeptTreeData)
}
