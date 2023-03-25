// ==========================================================================
// RYGO自动生成业务逻辑层相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2021-06-28 22:20:45 +0800 CST
// 生成路径: app/service/module/tenant/tenant_service.go
// 生成人：rygo
// ==========================================================================
package tenant

import (
	"errors"
	"rygo/app/ginframe/utils/convert"
	"rygo/app/ginframe/utils/page"
	"rygo/app/model/module/tenant"
	userService "rygo/app/service/system/user"
	"time"

	"github.com/gin-gonic/gin"
)

const Layout = "2006-01-02 15:04:05" //时间常量

//根据主键查询数据
func SelectRecordById(id int64) (*tenant.SysTenant, error) {
	entity := &tenant.SysTenant{Id: id}
	_, err := entity.FindOne()
	return entity, err
}

//根据主键删除数据
func DeleteRecordById(id int64) bool {
	rs, err := (&tenant.SysTenant{Id: id}).Delete()
	if err == nil {
		if rs > 0 {
			return true
		}
	}
	return false
}

//批量删除数据记录
func DeleteRecordByIds(ids string) int64 {
	ida := convert.ToInt64Array(ids, ",")
	result, err := tenant.DeleteBatch(ida...)
	if err != nil {
		return 0
	}
	return result
}

//添加数据
func AddSave(req *tenant.AddReq, c *gin.Context) (int64, error) {
	var entity tenant.SysTenant

	loc, _ := time.LoadLocation("Asia/Shanghai")
	s, _ := time.ParseInLocation(Layout, req.StartTime, loc)
	e, _ := time.ParseInLocation(Layout, req.EndTime, loc)

	entity.Name = req.Name
	entity.Address = req.Address
	entity.Manager = req.Manager
	entity.Phone = req.Phone
	entity.Remark = req.Remark
	entity.StartTime = s
	entity.EndTime = e
	entity.Email = req.Email
	entity.CreateTime = time.Now()
	entity.CreateBy = ""

	user := userService.GetProfile(c)

	if user != nil {
		entity.CreateBy = user.LoginName
	}

	_, err := entity.Insert()
	return entity.Id, err
}

//修改数据
func EditSave(req *tenant.EditReq, c *gin.Context) (int64, error) {
	entity := &tenant.SysTenant{Id: req.Id}
	ok, err := entity.FindOne()

	if err != nil {
		return 0, err
	}

	if !ok {
		return 0, errors.New("数据不存在")
	}

	loc, _ := time.LoadLocation("Asia/Shanghai")
	s, _ := time.ParseInLocation(Layout, req.StartTime, loc)
	e, _ := time.ParseInLocation(Layout, req.EndTime, loc)

	entity.Name = req.Name
	entity.Address = req.Address
	entity.Manager = req.Manager
	entity.Phone = req.Phone
	entity.Remark = req.Remark
	entity.StartTime = s
	entity.EndTime = e
	entity.Email = req.Email
	entity.UpdateTime = time.Now()
	entity.UpdateBy = ""

	user := userService.GetProfile(c)

	if user == nil {
		entity.UpdateBy = user.LoginName
	}

	return entity.Update()
}

//根据条件查询数据
func SelectListAll(params *tenant.SelectPageReq) ([]tenant.SysTenant, error) {
	return tenant.SelectListAll(params)
}

//根据条件分页查询数据
func SelectListByPage(params *tenant.SelectPageReq) ([]tenant.SysTenant, *page.Paging, error) {
	return tenant.SelectListByPage(params)
}

// 导出excel
func Export(param *tenant.SelectPageReq) (string, error) {
	head := []string{"ID", "商户名称", "联系地址", "负责人", "联系电话", "备注信息", "起租时间", "结束时间", "安全邮箱"}
	col := []string{"id", "name", "address", "manager", "phone", "remark", "start_time", "end_time", "email"}
	return tenant.SelectListExport(param, head, col)
}
