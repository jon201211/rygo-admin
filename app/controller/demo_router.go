package controller

import (
	"fmt"
	"rygo/app/controller/admin"

	"rygo/app/middleware/jwt"
	"rygo/app/router"
)

func init() {
	fmt.Println("##############111$####################################################")
	form := &admin.FormController{}
	g1 := router.New("admin", "/demo/form", jwt.JWTAuthMiddleware())
	g1.GET("/autocomplete", "", form.Autocomplete)
	g1.GET("/basic", "", form.Basic)
	g1.GET("/button", "", form.Button)
	g1.GET("/cards", "", form.Cards)
	g1.GET("/datetime", "", form.Datetime)
	g1.GET("/duallistbox", "", form.Duallistbox)
	g1.GET("/grid", "", form.Grid)

	g1.GET("/jasny", "", form.Jasny)
	g1.GET("/select", "", form.Select)
	g1.GET("/sortable", "", form.Sortable)
	g1.GET("/summernote", "", form.Summernote)
	g1.GET("/tabs_panels", "", form.Tabs_panels)

	g1.GET("/timeline", "", form.Timeline)
	g1.GET("/upload", "", form.Upload)
	g1.GET("/validate", "", form.Validate)
	g1.GET("/wizard", "", form.Wizard)

	icon := &admin.IconController{}
	g2 := router.New("admin", "/demo/icon", jwt.JWTAuthMiddleware())
	g2.GET("/fontawesome", "", icon.Fontawesome)
	g2.GET("/glyphicons", "", icon.Glyphicons)

	modal := &admin.ModalController{}
	g3 := router.New("admin", "/demo/modal", jwt.JWTAuthMiddleware())
	g3.GET("/dialog", "", modal.Dialog)
	g3.GET("/form", "", modal.Form)
	g3.GET("/layer", "", modal.Layer)
	g3.GET("/table", "", modal.Table)
	g3.GET("/check", "", modal.Check)
	g3.GET("/parent", "", modal.Parent)
	g3.GET("/radio", "", modal.Radio)

	operate := &admin.OperateController{}
	g4 := router.New("admin", "/demo/operate", jwt.JWTAuthMiddleware())
	g4.GET("/list", "", operate.List)
	g4.GET("/add", "", operate.Add)
	g4.GET("/detail", "", operate.Detail)
	g4.GET("/edit", "", operate.Edit)
	g4.POST("/edit", "", operate.EditSave)
	g4.GET("/other", "", operate.Other)
	g4.GET("/table", "", operate.Table)

	report := &admin.ReportController{}
	g5 := router.New("admin", "/demo/report", jwt.JWTAuthMiddleware())
	g5.GET("/echarts", "", report.Echarts)
	g5.GET("/metrics", "", report.Metrics)
	g5.GET("/peity", "", report.Peity)
	g5.GET("/sparkline", "", report.Sparkline)

	table := &admin.TableController{}
	g6 := router.New("admin", "/demo/table", jwt.JWTAuthMiddleware())
	g6.GET("/button", "", table.Button)
	g6.GET("/child", "", table.Child)
	g6.GET("/curd", "", table.Curd)
	g6.GET("/detail", "", table.Detail)
	g6.POST("list", "", table.List)

	g6.GET("/editable", "", table.Editable, jwt.JWTAuthMiddleware())
	g6.GET("/event", "", table.Event)
	g6.POST("/export", "", table.Export)
	g6.GET("/fixedColumns", "", table.FixedColumns)
	g6.GET("/footer", "", table.Footer)
	g6.GET("/groupHeader", "", table.GroupHeader)

	g6.GET("/image", "", table.Image)
	g6.GET("/multi", "", table.Multi)
	g6.GET("/other", "", table.Other)
	g6.GET("/pageGo", "", table.PageGo)

	g6.GET("/params", "", table.Params)
	g6.GET("/remember", "", table.Remember)
	g6.GET("/recorder", "", table.Recorder)
	g6.GET("/search", "", table.Search)
}
