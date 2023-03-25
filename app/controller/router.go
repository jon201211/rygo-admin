package controller

import (
	_ "rygo/app/controller/demo"
	_ "rygo/app/controller/module"
	_ "rygo/app/controller/monitor"
	_ "rygo/app/controller/system"
	errorc "rygo/app/controller/system/error"
	"rygo/app/controller/system/index"
	_ "rygo/app/controller/tool"
	"rygo/app/ginframe/router"
	"rygo/app/service/middleware/auth"
	"rygo/app/service/middleware/jwt"
)

func init() {
	// 加载登陆路由
	g0 := router.New("admin", "/")
	g0.GET("/", "", index.Index)
	g0.GET("/login", "", index.Login)
	g0.POST("/checklogin", "", index.CheckLogin)
	g0.GET("/captchaImage", "", index.CaptchaImage)
	g0.GET("/500", "", errorc.Error)
	g0.GET("/404", "", errorc.NotFound)
	g0.GET("/403", "", errorc.Unauth)
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
