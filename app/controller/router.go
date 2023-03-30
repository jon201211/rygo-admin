package controller

import (
	"rygo/app/controller/admin"
	//_ "rygo/app/controller/api"
	_ "rygo/app/controller/base"

	"rygo/app/middleware/auth"
	"rygo/app/middleware/jwt"
	"rygo/app/router"
)

func init() {
	// 加载登陆路由
	index := &admin.IndexController{}
	login := &admin.LoginController{}
	err := &admin.ErrorController{}
	g0 := router.New("admin", "/")
	g0.GET("/", "", index.Index)
	g0.GET("/login", "", login.Login)
	g0.POST("/checklogin", "", login.CheckLogin)
	g0.GET("/captchaImage", "", login.CaptchaImage)
	g0.GET("/500", "", err.Error)
	g0.GET("/404", "", err.NotFound)
	g0.GET("/403", "", err.Unauth)
	//下在要检测是否登陆
	g1 := router.New("admin", "/", jwt.JWTAuthMiddleware())
	g1.GET("/index", "", index.Index)
	g1.GET("/index_left", "", index.Index)
	g1.GET("/logout", "", index.Logout)

	// 加载框架路由
	g2 := router.New("admin", "/system", jwt.JWTAuthMiddleware(), auth.Auth)
	g2.GET("/main", "", index.Main)
	g2.GET("/switchSkin", "", index.SwitchSkin)
	g2.GET("/download", "", index.Download)
}
