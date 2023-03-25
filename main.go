package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"rygo/app/common/util"
	_ "rygo/app/controller"
	"rygo/app/ginframe/cfg"
	"rygo/app/ginframe/db"
	"rygo/app/ginframe/server"
	"rygo/app/model/module/tenant"
	"rygo/app/model/system/dept"
	"rygo/app/model/system/post"
	"rygo/app/model/system/user"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

// @title RYGO 自动生成API文档
// @version 1.0
// @description 生成文档请在调试模式下进行<a href="/tool/swagger?a=r">重新生成文档</a>

// @host localhost
// @BasePath /api
func main() {
	gin.SetMode("debug")
	config := cfg.Instance()
	if config == nil {
		fmt.Printf("参数错误")
		return
	}
	db.Instance().Engine().Sync2(dept.SysDept{}, tenant.SysTenant{}, user.SysUser{}, post.SysPost{})
	//后台服务状态
	admin := server.New("admin", config.Admin.Address)
	admin.Template("template").Static(config.Admin.ServerRoot)
	admin.Start(g)

	//打开浏览器
	if runtime.GOOS == "windows" {
		util.OpenWin("http://127.0.0.1:8080")
	}

	if err := g.Wait(); err != nil {
		fmt.Println(err.Error())
	}
	var state int32 = 1
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

EXIT:
	for {
		sig := <-sc
		fmt.Println("获取到信号[%s]", sig.String())
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			atomic.StoreInt32(&state, 0)
			break EXIT
		case syscall.SIGHUP:
		default:
			break EXIT
		}
	}

	fmt.Println("服务退出")
	time.Sleep(time.Second)
	os.Exit(int(atomic.LoadInt32(&state)))
}
