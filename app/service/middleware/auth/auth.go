package auth

import (
	"net/http"
	"rygo/app/ginframe/router"
	"rygo/app/model"
	menuService "rygo/app/service/system/menu"
	userService "rygo/app/service/system/user"
	"strings"

	"github.com/gin-gonic/gin"
)

// 鉴权中间件，只有登录成功之后才能通过
func Auth(c *gin.Context) {

	//判断是否登陆
	if userService.IsSignedIn(c) {
		//根据url判断是否有权限
		url := c.Request.URL.Path
		strEnd := url[len(url)-1 : len(url)]
		if strings.EqualFold(strEnd, "/") {
			url = strings.TrimRight(url, "/")
		}
		//获取权限标识
		permission := router.FindPermission(url)
		if len(permission) > 0 {
			//获取用户信息
			user := userService.GetProfile(c)
			//获取用户菜单列表
			menus, err := menuService.SelectMenuNormalByUser(user.UserId)
			if err != nil {

				c.Redirect(http.StatusFound, "/500")
				c.Abort()
				return
			}

			if menus == nil {
				c.Redirect(http.StatusFound, "/500")
				c.Abort()
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
				ajaxString := c.Request.Header.Get("X-Requested-With")
				if strings.EqualFold(ajaxString, "XMLHttpRequest") {
					c.JSON(http.StatusOK, model.CommonRes{
						Code: 403,
						Msg:  "您没有操作权限",
					})
					c.Abort()
				} else {
					c.Redirect(http.StatusFound, "/403")
					c.Abort()
				}
			}
		}

		c.Next()
	} else {
		c.Redirect(http.StatusFound, "/demo")
		c.Abort()
	}
}
