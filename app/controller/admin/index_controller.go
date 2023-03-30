package admin

import (
	"io/ioutil"
	"net/http"
	"os"
	"rygo/app/controller/base"
	"rygo/app/model"

	"rygo/app/response"
	"rygo/app/service"

	"rygo/app/utils/gconv"

	"github.com/gin-gonic/gin"
)

type IndexController struct {
	base.BaseController
}

//后台框架首页
func (c *IndexController) Index(ctx *gin.Context) {
	if service.UserService.IsSignedIn(ctx) {
		c.goIndex(ctx, "index")
	} else {
		ctx.Redirect(http.StatusFound, "/login")
	}
}

func (c *IndexController) goIndex(ctx *gin.Context, indexPageDefault string) {
	user := service.UserService.GetProfile(ctx)
	loginname := user.LoginName
	username := user.UserName
	avatar := user.Avatar
	if avatar == "" {
		avatar = "/resource/img/profile.jpg"
	}

	var menus *[]model.MenuEntityExtend

	//获取菜单数据
	if service.UserService.IsAdmin(user.UserId) {
		tmp, err := service.MenuService.SelectMenuNormalAll()
		if err == nil {
			menus = tmp
		}

	} else {
		tmp, err := service.MenuService.SelectMenusByUserId(gconv.String(user.UserId))
		if err == nil {
			menus = tmp
		}
	}

	//获取配置数据
	sideTheme := service.ConfigService.GetValueByKey("sys.index.sideTheme")
	skinName := service.ConfigService.GetValueByKey("sys.index.skinName")
	//设置首页风格

	menuStyle := ctx.Query("menuStyle")
	cookie, _ := ctx.Request.Cookie("menuStyle")
	if cookie == nil {
		cookie = &http.Cookie{
			Name:     "menuStyle",
			Value:    menuStyle,
			HttpOnly: true,
		}
		http.SetCookie(ctx.Writer, cookie)
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
	ctx.SetCookie(cookie.Name, menuStyle, cookie.MaxAge, cookie.Path, cookie.Domain, cookie.SameSite, cookie.Secure, cookie.HttpOnly)
	response.BuildTpl(ctx, targetIndex).WriteTpl(gin.H{
		"avatar":    avatar,
		"loginname": loginname,
		"username":  username,
		"menus":     menus,
		"sideTheme": sideTheme,
		"skinName":  skinName,
	})
}

//后台框架欢迎页面
func (c *IndexController) Main(ctx *gin.Context) {
	response.BuildTpl(ctx, "main").WriteTpl()
}

//下载文件
func (c *IndexController) Download(ctx *gin.Context) {
	fileName := ctx.Query("fileName")
	delete := ctx.Query("delete")

	if fileName == "" {
		response.BuildTpl(ctx, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "参数错误",
		})
		return
	}

	// 创建路径
	curDir, err := os.Getwd()
	if err != nil {
		response.BuildTpl(ctx, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "获取目录失败",
		})
		return
	}

	filepath := curDir + "/public/upload/" + fileName
	file, err := os.Open(filepath)

	defer file.Close()

	if err != nil {
		response.BuildTpl(ctx, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "参数错误",
		})
		return
	}

	b, _ := ioutil.ReadAll(file)
	ctx.Writer.Header().Add("Content-Disposition", "attachment")
	ctx.Writer.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	ctx.Writer.Write(b)

	if delete == "true" {
		os.Remove(filepath)
	}

}

//切换皮肤
func (c *IndexController) SwitchSkin(ctx *gin.Context) {
	response.BuildTpl(ctx, "skin").WriteTpl()
}

//注销
func (c *IndexController) Logout(ctx *gin.Context) {
	if service.UserService.IsSignedIn(ctx) {
		service.UserService.SignOut(ctx)
	}

	ctx.Redirect(http.StatusFound, "/login")
	ctx.Abort()
}
