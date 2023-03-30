package auth

import (
	"net/http"
	"rygo/app/model"
	"rygo/app/router"
	"rygo/app/service"

	"strings"

	"github.com/gin-gonic/gin"
)

// 鉴权中间件，只有登录成功之后才能通过
func Auth(ctx *gin.Context) {

	//判断是否登陆
	if service.UserService.IsSignedIn(ctx) {
		//根据url判断是否有权限
		url := ctx.Request.URL.Path
		strEnd := url[len(url)-1 : len(url)]
		if strings.EqualFold(strEnd, "/") {
			url = strings.TrimRight(url, "/")
		}
		//获取权限标识
		permission := router.FindPermission(url)
		if len(permission) > 0 {
			//获取用户信息
			user := service.UserService.GetProfile(ctx)
			//获取用户菜单列表
			menus, err := service.MenuService.SelectMenuNormalByUser(user.UserId)
			if err != nil {

				ctx.Redirect(http.StatusFound, "/500")
				ctx.Abort()
				return
			}

			if menus == nil {
				ctx.Redirect(http.StatusFound, "/500")
				ctx.Abort()
				return
			}

			hasPermission := false

			for i := range *menus {
				if strings.EqualFold((*menus)[i].Perms, permission) {
					hasPermission = true
					break
				}
			}

			if !hasPermission {
				ajaxString := ctx.Request.Header.Get("X-Requested-With")
				if strings.EqualFold(ajaxString, "XMLHttpRequest") {
					ctx.JSON(http.StatusOK, model.CommonRes{
						Code: 403,
						Msg:  "您没有操作权限",
					})
					ctx.Abort()
				} else {
					ctx.Redirect(http.StatusFound, "/403")
					ctx.Abort()
				}
			}
		}

		ctx.Next()
	} else {
		ctx.Redirect(http.StatusFound, "/demo")
		ctx.Abort()
	}
}
