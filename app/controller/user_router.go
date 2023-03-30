package controller

import (
	"rygo/app/controller/admin"
	"rygo/app/middleware/auth"
	"rygo/app/middleware/jwt"
	"rygo/app/router"
)

//加载路由
func init() {
	// 用户管理路由
	user := &admin.UserController{}
	g1 := router.New("admin", "/system/user", jwt.JWTAuthMiddleware(), auth.Auth)
	g1.GET("/", "system:user:view", user.List)
	g1.POST("/list", "system:user:list", user.ListAjax)
	g1.GET("/add", "system:user:add", user.Add)
	g1.POST("/add", "system:user:add", user.AddSave)
	g1.POST("/remove", "system:user:remove", user.Remove)
	g1.GET("/edit", "system:user:edit", user.Edit)
	g1.POST("/edit", "system:user:edit", user.EditSave)
	g1.POST("/export", "system:user:export", user.Export)
	g1.GET("/resetPwd", "system:user:resetPwd", user.ResetPwd)
	g1.POST("/resetPwd", "system:user:resetPwd", user.ResetPwdSave)

	// 个人中心路由
	profile := &admin.ProfileController{}
	g2 := router.New("admin", "/system/user/profile", jwt.JWTAuthMiddleware(), auth.Auth)
	g2.GET("/", "", profile.Profile)
	g2.GET("/avatar", "", profile.Avatar)
	g2.GET("/resetPwd", "", profile.EditPwd)
	g2.POST("/update", "", profile.Update)
	g2.POST("/resetSavePwd", "", profile.UpdatePassword)
	g2.POST("/checkEmailUnique", "", profile.CheckEmailUnique)
	g2.POST("/checkPhoneUnique", "", profile.CheckPhoneUnique)
	g2.POST("/checkLoginNameUnique", "", profile.CheckLoginNameUnique)
	g2.POST("/checkEmailUniqueAll", "", profile.CheckEmailUniqueAll)
	g2.POST("/checkPhoneUniqueAll", "", profile.CheckPhoneUniqueAll)
	g2.POST("/checkPassword", "", profile.CheckPassword)
	g2.POST("/updateAvatar", "", profile.UpdateAvatar)
}
