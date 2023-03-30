package admin

import (
	"html/template"
	"net/http"
	"os"
	"rygo/app/controller/base"
	"rygo/app/model"

	"rygo/app/response"
	"rygo/app/service"

	"rygo/app/utils/file"
	"rygo/app/utils/gconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type GenController struct {
	base.BaseController
}

func FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

//生成代码列表页面
func (c *GenController) Gen(ctx *gin.Context) {
	response.BuildTpl(ctx, "tool/gen/list").WriteTpl()
}

func (c *GenController) GenList(ctx *gin.Context) {
	var req *model.TableSelectPageReq
	//获取参数
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorResp(ctx).SetMsg(err.Error()).Log("生成代码", req).WriteJsonExit()
		return
	}
	rows := make([]model.TableEntity, 0)
	result, page, err := service.TableService.SelectListByPage(req)

	if err == nil && len(result) > 0 {
		rows = result
	}

	response.BuildTable(ctx, page.Total, rows).WriteJsonExit()
}

//导入数据表
func (c *GenController) ImportTable(ctx *gin.Context) {
	response.BuildTpl(ctx, "tool/gen/importTable").WriteTpl()
}

//删除数据
func (c *GenController) Remove(ctx *gin.Context) {
	var req *model.RemoveReq
	//获取参数
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Del).Log("生成代码", req).WriteJsonExit()
		return
	}

	rs := service.TableService.DeleteRecordByIds(req.Ids)

	if rs > 0 {
		response.SucessResp(ctx).SetBtype(model.Buniss_Del).Log("生成代码", req).WriteJsonExit()
	} else {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Del).Log("生成代码", req).WriteJsonExit()
	}
}

//修改数据
func (c *GenController) Edit(ctx *gin.Context) {
	id := gconv.Int64(ctx.Query("id"))
	if id <= 0 {
		response.BuildTpl(ctx, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "参数错误",
		})
		return
	}

	entity, err := service.TableService.SelectRecordById(id)

	if err != nil || entity == nil {
		response.BuildTpl(ctx, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "参数不存在",
		})
		return
	}

	goTypeTpl := service.TableService.GoTypeTpl()
	queryTypeTpl := service.TableService.QueryTypeTpl()
	htmlTypeTpl := service.TableService.HtmlTypeTpl()

	response.BuildTpl(ctx, "tool/gen/edit").WriteTpl(gin.H{
		"table":        entity,
		"goTypeTpl":    template.HTML(goTypeTpl),
		"queryTypeTpl": template.HTML(queryTypeTpl),
		"htmlTypeTpl":  template.HTML(htmlTypeTpl),
	})
}

//修改数据保存
func (c *GenController) EditSave(ctx *gin.Context) {
	var req model.TableEditReq
	//获取参数
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorResp(ctx).SetMsg(err.Error()).SetBtype(model.Buniss_Edit).Log("生成代码", gin.H{"tableName": req.TableName}).WriteJsonExit()
		return
	}
	_, err := service.TableService.SaveEdit(&req, ctx)
	if err != nil {
		response.ErrorResp(ctx).SetMsg(err.Error()).SetBtype(model.Buniss_Edit).Log("生成代码", gin.H{"tableName": req.TableName}).WriteJsonExit()
		return
	}
	response.SucessResp(ctx).SetBtype(model.Buniss_Edit).Log("生成代码", gin.H{"tableName": req.TableName}).WriteJsonExit()
}

//预览代码
func (c *GenController) Preview(ctx *gin.Context) {
	tableId := gconv.Int64(ctx.Query("tableId"))
	if tableId <= 0 {
		ctx.JSON(http.StatusOK, model.CommonRes{
			Code:  500,
			Btype: model.Buniss_Other,
			Msg:   "参数错误",
		})
	}

	entity, err := service.TableService.SelectRecordById(tableId)

	if err != nil || entity == nil {
		ctx.JSON(http.StatusOK, model.CommonRes{
			Code:  500,
			Btype: model.Buniss_Other,
			Msg:   "数据不存在",
		})
	}

	service.TableService.SetPkColumn(entity, entity.Columns)

	addKey := "vm/html/add.html.vm"
	addValue := ""
	editKey := "vm/html/edit.html.vm"
	editValue := ""

	listKey := "vm/html/list.html.vm"
	listValue := ""
	listTmp := "vm/html/list.txt"

	treeKey := "vm/html/tree.html.vm"
	treeValue := ""

	if entity.TplCategory == "tree" {
		listTmp = "vm/html/list-tree.txt"
	}

	sqlKey := "vm/sql/menu.sql.vm"
	sqlValue := ""
	entityKey := "vm/go/" + entity.BusinessName + "_entity.go.vm"
	entityValue := ""
	extendKey := "vm/go/" + entity.BusinessName + "_dao.go.vm"
	extendValue := ""
	serviceKey := "vm/go/" + entity.BusinessName + "_service.go.vm"
	serviceValue := ""
	routerKey := "vm/go/" + entity.BusinessName + "_router.go.vm"
	routerValue := ""
	controllerKey := "vm/go/" + entity.BusinessName + "_controller.go.vm"
	controllerValue := ""

	BigBusinessName := FirstUpper(entity.BusinessName)
	//新增页面模板
	addValue, _ = service.TableService.LoadTemplate("vm/html/add.txt", gin.H{"table": entity, "BigBusinessName": BigBusinessName})
	//修改页面模板
	editValue, _ = service.TableService.LoadTemplate("vm/html/edit.txt", gin.H{"table": entity, "BigBusinessName": BigBusinessName})

	//列表页面模板
	listValue, _ = service.TableService.LoadTemplate(listTmp, gin.H{"table": entity, "BigBusinessName": BigBusinessName})

	if entity.TplCategory == "tree" {
		//选择树页面模板
		treeValue, _ = service.TableService.LoadTemplate("vm/html/tree.txt", gin.H{"table": entity, "BigBusinessName": BigBusinessName})
	}

	//entity模板
	entityValue, _ = service.TableService.LoadTemplate("vm/go/entity.txt", gin.H{"table": entity, "BigBusinessName": BigBusinessName})

	//extend模板
	extendValue, _ = service.TableService.LoadTemplate("vm/go/dao.txt", gin.H{"table": entity, "BigBusinessName": BigBusinessName})

	//service模板
	serviceValue, _ = service.TableService.LoadTemplate("vm/go/service.txt", gin.H{"table": entity, "BigBusinessName": BigBusinessName})

	//router模板
	routerValue, _ = service.TableService.LoadTemplate("vm/go/router.txt", gin.H{"table": entity, "BigBusinessName": BigBusinessName})

	//controller模板
	controllerValue, _ = service.TableService.LoadTemplate("vm/go/controller.txt", gin.H{"table": entity, "BigBusinessName": BigBusinessName})

	//sql模板
	sqlValue, _ = service.TableService.LoadTemplate("vm/sql/sql.txt", gin.H{"table": entity, "BigBusinessName": BigBusinessName})

	if entity.TplCategory == "tree" {
		ctx.JSON(http.StatusOK, model.CommonRes{
			Code:  0,
			Btype: model.Buniss_Other,
			Data: gin.H{
				addKey:        addValue,
				editKey:       editValue,
				listKey:       listValue,
				treeKey:       treeValue,
				sqlKey:        sqlValue,
				entityKey:     entityValue,
				extendKey:     extendValue,
				serviceKey:    serviceValue,
				routerKey:     routerValue,
				controllerKey: controllerValue,
			},
		})
	} else {
		ctx.JSON(http.StatusOK, model.CommonRes{
			Code:  0,
			Btype: model.Buniss_Other,
			Data: gin.H{
				addKey:        addValue,
				editKey:       editValue,
				listKey:       listValue,
				sqlKey:        sqlValue,
				entityKey:     entityValue,
				extendKey:     extendValue,
				serviceKey:    serviceValue,
				routerKey:     routerValue,
				controllerKey: controllerValue,
			},
		})
	}

}

//生成代码
func (c *GenController) GenCode(ctx *gin.Context) {
	tableId := gconv.Int64(ctx.Query("tableId"))
	if tableId <= 0 {
		ctx.JSON(http.StatusOK, model.CommonRes{
			Code:  500,
			Btype: model.Buniss_Other,
			Msg:   "参数错误",
		})
	}

	entity, err := service.TableService.SelectRecordById(tableId)

	if err != nil || entity == nil {
		ctx.JSON(http.StatusOK, model.CommonRes{
			Code:  500,
			Btype: model.Buniss_Other,
			Msg:   "数据不存在",
		})
	}

	service.TableService.SetPkColumn(entity, entity.Columns)

	listTmp := "vm/html/list.txt"
	if entity.TplCategory == "tree" {
		listTmp = "vm/html/list-tree.txt"
	}

	//获取当前运行时目录
	curDir, err := os.Getwd()

	curDir = curDir + "/generated" //TEST

	if err != nil {
		response.ErrorResp(ctx).SetMsg(err.Error()).Log("生成代码", gin.H{"tableId": tableId}).WriteJsonExit()
	}

	//add模板
	BigBusinessName := FirstUpper(entity.BusinessName)
	if tmp, err := service.TableService.LoadTemplate("vm/html/add.txt", gin.H{"table": entity, "BigBusinessName": BigBusinessName}); err == nil {
		fileName := strings.Join([]string{curDir, "/template/", entity.ModuleName, "/", entity.BusinessName, "/add.html"}, "")

		//if !file.Exists(fileName) {
		f, err := file.Create(fileName)
		if err == nil {
			f.WriteString(tmp)
		}
		f.Close()
		//}
	}

	//edit模板
	if tmp, err := service.TableService.LoadTemplate("vm/html/edit.txt", gin.H{"table": entity, "BigBusinessName": BigBusinessName}); err == nil {
		fileName := strings.Join([]string{curDir, "/template/", entity.ModuleName, "/", entity.BusinessName, "/edit.html"}, "")

		//if !file.Exists(fileName) {
		f, err := file.Create(fileName)
		if err == nil {
			f.WriteString(tmp)
		}
		f.Close()
		//}
	}

	//list模板
	if tmp, err := service.TableService.LoadTemplate(listTmp, gin.H{"table": entity, "BigBusinessName": BigBusinessName}); err == nil {
		fileName := strings.Join([]string{curDir, "/template/", entity.ModuleName, "/", entity.BusinessName, "/list.html"}, "")

		//if !file.Exists(fileName) {
		f, err := file.Create(fileName)
		if err == nil {
			f.WriteString(tmp)
		}
		f.Close()
		//}
	}

	if entity.TplCategory == "tree" {
		//tree模板
		if tmp, err := service.TableService.LoadTemplate("vm/html/tree.txt", gin.H{"table": entity, "BigBusinessName": BigBusinessName}); err == nil {
			fileName := strings.Join([]string{curDir, "/template/", entity.ModuleName, "/", "tree.html"}, "")

			//if !file.Exists(fileName) {
			f, err := file.Create(fileName)
			if err == nil {
				f.WriteString(tmp)
			}
			f.Close()
			//}
		}
	}

	//entity模板
	if tmp, err := service.TableService.LoadTemplate("vm/go/entity.txt", gin.H{"table": entity, "BigBusinessName": BigBusinessName}); err == nil {
		fileName := strings.Join([]string{curDir, "/app/model/", entity.BusinessName, "_entity.go"}, "")
		if file.Exists(fileName) {
			os.RemoveAll(fileName)
		}

		f, err := file.Create(fileName)
		if err == nil {
			f.WriteString(tmp)
		}
		f.Close()
	}

	//extend模板
	if tmp, err := service.TableService.LoadTemplate("vm/go/dao.txt", gin.H{"table": entity, "BigBusinessName": BigBusinessName}); err == nil {
		fileName := strings.Join([]string{curDir, "/app/dao/", entity.BusinessName, "_dao.go"}, "")

		//if !file.Exists(fileName) {
		f, err := file.Create(fileName)
		if err == nil {
			f.WriteString(tmp)
		}
		f.Close()
		//}
	}

	//service模板
	if tmp, err := service.TableService.LoadTemplate("vm/go/service.txt", gin.H{"table": entity, "BigBusinessName": BigBusinessName}); err == nil {
		fileName := strings.Join([]string{curDir, "/app/service/", entity.BusinessName, "_service.go"}, "")

		//if !file.Exists(fileName) {
		f, err := file.Create(fileName)
		if err == nil {
			f.WriteString(tmp)
		}
		f.Close()
		//}
	}

	//router模板
	if tmp, err := service.TableService.LoadTemplate("vm/go/router.txt", gin.H{"table": entity, "BigBusinessName": BigBusinessName}); err == nil {
		fileName := strings.Join([]string{curDir, "/app/controller/", entity.BusinessName, "_router.go"}, "")

		//if !file.Exists(fileName) {
		f, err := file.Create(fileName)
		if err == nil {
			f.WriteString(tmp)
		}
		f.Close()
		//}
	}

	//controller模板
	if tmp, err := service.TableService.LoadTemplate("vm/go/controller.txt", gin.H{"table": entity, "BigBusinessName": BigBusinessName}); err == nil {
		fileName := strings.Join([]string{curDir, "/app/controller/", "admin", "/", entity.BusinessName, "_controller.go"}, "")

		//if !file.Exists(fileName) {
		f, err := file.Create(fileName)
		if err == nil {
			f.WriteString(tmp)
		}
		f.Close()
		//}
	}

	//sql模板
	if tmp, err := service.TableService.LoadTemplate("vm/sql/sql.txt", gin.H{"table": entity, "BigBusinessName": BigBusinessName}); err == nil {
		fileName := strings.Join([]string{curDir, "/data/sql/", entity.ModuleName, "/", entity.BusinessName, "_menu.sql"}, "")

		//if !file.Exists(fileName) {
		f, err := file.Create(fileName)
		if err == nil {
			f.WriteString(tmp)
		}
		f.Close()
		//}
	}
	response.SucessResp(ctx).Log("生成代码", gin.H{"tableId": tableId}).WriteJsonExit()
}

//查询数据库列表
func (c *GenController) DataList(ctx *gin.Context) {
	var req *model.TableSelectPageReq
	//获取参数
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorResp(ctx).SetMsg(err.Error()).Log("生成代码", req).WriteJsonExit()
	}
	rows := make([]model.TableEntity, 0)
	result, page, err := service.TableService.SelectDbTableList(req)

	if err == nil && len(result) > 0 {
		rows = result
	}

	ctx.JSON(http.StatusOK, model.TableDataInfo{
		Code:  0,
		Msg:   "操作成功",
		Total: page.Total,
		Rows:  rows,
	})
}

//导入表结构（保存）
func (c *GenController) ImportTableSave(ctx *gin.Context) {
	tables := ctx.PostForm("tables")
	if tables == "" {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Add).SetMsg("参数错误").Log("生成代码", gin.H{"tables": tables}).WriteJsonExit()
	}

	user := service.UserService.GetProfile(ctx)
	if user == nil {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Add).SetMsg("登陆超时").Log("生成代码", gin.H{"tables": tables}).WriteJsonExit()
	}

	operName := user.LoginName

	tableArr := strings.Split(tables, ",")
	tableList, err := service.TableService.SelectDbTableListByNames(tableArr)
	if err != nil {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Add).SetMsg(err.Error()).Log("生成代码", gin.H{"tables": tables}).WriteJsonExit()
		return
	}

	if tableList == nil {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Add).SetMsg("请选择需要导入的表").Log("生成代码", gin.H{"tables": tables}).WriteJsonExit()
		return
	}

	service.TableService.ImportGenTable(&tableList, operName)
	response.SucessResp(ctx).Log("导入表结构", gin.H{"tables": tables}).WriteJsonExit()
}

//根据table_id查询表列数据
func (c *GenController) ColumnList(ctx *gin.Context) {
	tableId := gconv.Int64(ctx.PostForm("tableId"))
	//获取参数
	if tableId <= 0 {
		response.ErrorResp(ctx).SetBtype(model.Buniss_Add).SetMsg("参数错误").Log("生成代码", gin.H{"tableId": tableId})
	}
	rows := make([]model.TableColumnEntity, 0)
	result, err := service.TableService.SelectGenTableColumnListByTableId(tableId)

	if err == nil && len(result) > 0 {
		rows = result
	}

	ctx.JSON(http.StatusOK, model.TableDataInfo{
		Code:  0,
		Msg:   "操作成功",
		Total: len(rows),
		Rows:  rows,
	})
}
