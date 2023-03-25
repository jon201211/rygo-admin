package index

import (
	"io/ioutil"
	"net/http"
	"os"
	"rygo/app/ginframe/response"
	"rygo/app/ginframe/utils/gconv"
	"rygo/app/model"
	"rygo/app/model/system/menu"
	configService "rygo/app/service/system/config"
	menuService "rygo/app/service/system/menu"
	userService "rygo/app/service/system/user"

	"github.com/gin-gonic/gin"
)

//后台框架首页
func Index(c *gin.Context) {
	if userService.IsSignedIn(c) {
		goIndex(c, "index")
	} else {
		c.Redirect(http.StatusFound, "/login")
	}
}

func goIndex(c *gin.Context, indexPageDefault string) {
	user := userService.GetProfile(c)
	loginname := user.LoginName
	username := user.UserName
	avatar := user.Avatar
	if avatar == "" {
		avatar = "/resource/img/profile.jpg"
	}

	var menus *[]menu.EntityExtend

	//获取菜单数据
	if userService.IsAdmin(user.UserId) {
		tmp, err := menuService.SelectMenuNormalAll()
		if err == nil {
			menus = tmp
		}

	} else {
		tmp, err := menuService.SelectMenusByUserId(gconv.String(user.UserId))
		if err == nil {
			menus = tmp
		}
	}

	//获取配置数据
	sideTheme := configService.GetValueByKey("sys.index.sideTheme")
	skinName := configService.GetValueByKey("sys.index.skinName")
	//设置首页风格

	menuStyle := c.Query("menuStyle")
	cookie, _ := c.Request.Cookie("menuStyle")
	if cookie == nil {
		cookie = &http.Cookie{
			Name:     "menuStyle",
			Value:    menuStyle,
			HttpOnly: true,
		}
		http.SetCookie(c.Writer, cookie)
	}
	if menuStyle == "" { //未指定则从cookie中取
		menuStyle = cookie.Value
	}
	var targetIndex string         //默认首页
	if menuStyle == "index_left" { //指定了左侧风格,
		targetIndex = "index_left"
	} else { //否则默认风格
		targetIndex = indexPageDefault
	}
	//"menuStyle", cookie.Value, 1000, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly
	c.SetCookie(cookie.Name, menuStyle, cookie.MaxAge, cookie.Path, cookie.Domain, cookie.SameSite, cookie.Secure, cookie.HttpOnly)
	response.BuildTpl(c, targetIndex).WriteTpl(gin.H{
		"avatar":    avatar,
		"loginname": loginname,
		"username":  username,
		"menus":     menus,
		"sideTheme": sideTheme,
		"skinName":  skinName,
	})
}

//后台框架欢迎页面
func Main(c *gin.Context) {
	response.BuildTpl(c, "main").WriteTpl()
}

//下载文件
func Download(c *gin.Context) {
	fileName := c.Query("fileName")
	delete := c.Query("delete")

	if fileName == "" {
		response.BuildTpl(c, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "参数错误",
		})
		return
	}

	// 创建路径
	curDir, err := os.Getwd()
	if err != nil {
		response.BuildTpl(c, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "获取目录失败",
		})
		return
	}

	filepath := curDir + "/public/upload/" + fileName
	file, err := os.Open(filepath)

	defer file.Close()

	if err != nil {
		response.BuildTpl(c, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "参数错误",
		})
		return
	}

	b, _ := ioutil.ReadAll(file)
	c.Writer.Header().Add("Content-Disposition", "attachment")
	c.Writer.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Writer.Write(b)

	if delete == "true" {
		os.Remove(filepath)
	}

}

//切换皮肤
func SwitchSkin(c *gin.Context) {
	response.BuildTpl(c, "skin").WriteTpl()
}

//注销
func Logout(c *gin.Context) {
	if userService.IsSignedIn(c) {
		userService.SignOut(c)
	}

	c.Redirect(http.StatusFound, "/login")
	c.Abort()
}
