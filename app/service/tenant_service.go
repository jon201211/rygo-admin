package service

import (
	"errors"
	"rygo/app/dao"
	"rygo/app/model"

	"rygo/app/utils/convert"
	"rygo/app/utils/page"
	"time"

	"github.com/gin-gonic/gin"
)

const Layout = "2006-01-02 15:04:05" //时间常量

var TenantService = newTenantService()

func newTenantService() *tenantService {
	return &tenantService{}
}

type tenantService struct {
}

//根据主键查询数据
func (s *tenantService) SelectRecordById(id int64) (*model.SysTenant, error) {
	entity := &model.SysTenant{Id: id}
	_, err := dao.TenantDao.FindOne(entity)
	return entity, err
}

//根据主键删除数据
func (s *tenantService) DeleteRecordById(id int64) bool {
	entity := &model.SysTenant{Id: id}
	rs, err := dao.TenantDao.Delete(entity)
	if err == nil {
		if rs > 0 {
			return true
		}
	}
	return false
}

//批量删除数据记录
func (s *tenantService) DeleteRecordByIds(ids string) int64 {
	ida := convert.ToInt64Array(ids, ",")
	result, err := dao.TenantDao.DeleteBatch(ida...)
	if err != nil {
		return 0
	}
	return result
}

//添加数据
func (s *tenantService) AddSave(req *model.TenantAddReq, ctx *gin.Context) (int64, error) {
	var entity model.SysTenant

	loc, _ := time.LoadLocation("Asia/Shanghai")
	st, _ := time.ParseInLocation(Layout, req.StartTime, loc)
	e, _ := time.ParseInLocation(Layout, req.EndTime, loc)

	entity.Name = req.Name
	entity.Address = req.Address
	entity.Manager = req.Manager
	entity.Phone = req.Phone
	entity.Remark = req.Remark
	entity.StartTime = st
	entity.EndTime = e
	entity.Email = req.Email
	entity.CreateTime = time.Now()
	entity.CreateBy = ""

	user := UserService.GetProfile(ctx)

	if user != nil {
		entity.CreateBy = user.LoginName
	}

	_, err := dao.TenantDao.Insert(&entity)
	return entity.Id, err
}

//修改数据
func (s *tenantService) EditSave(req *model.TenantEditReq, ctx *gin.Context) (int64, error) {
	entity := &model.SysTenant{Id: req.Id}
	ok, err := dao.TenantDao.FindOne(entity)

	if err != nil {
		return 0, err
	}

	if !ok {
		return 0, errors.New("数据不存在")
	}

	loc, _ := time.LoadLocation("Asia/Shanghai")
	st, _ := time.ParseInLocation(Layout, req.StartTime, loc)
	e, _ := time.ParseInLocation(Layout, req.EndTime, loc)

	entity.Name = req.Name
	entity.Address = req.Address
	entity.Manager = req.Manager
	entity.Phone = req.Phone
	entity.Remark = req.Remark
	entity.StartTime = st
	entity.EndTime = e
	entity.Email = req.Email
	entity.UpdateTime = time.Now()
	entity.UpdateBy = ""

	user := UserService.GetProfile(ctx)

	if user == nil {
		entity.UpdateBy = user.LoginName
	}

	return dao.TenantDao.Update(entity)
}

//根据条件查询数据
func (s *tenantService) SelectListAll(params *model.TenantSelectPageReq) ([]model.SysTenant, error) {
	return dao.TenantDao.SelectListAll(params)
}

//根据条件分页查询数据
func (s *tenantService) SelectListByPage(params *model.TenantSelectPageReq) ([]model.SysTenant, *page.Paging, error) {
	return dao.TenantDao.SelectListByPage(params)
}

// 导出excel
func (s *tenantService) Export(param *model.TenantSelectPageReq) (string, error) {
	head := []string{"ID", "商户名称", "联系地址", "负责人", "联系电话", "备注信息", "起租时间", "结束时间", "安全邮箱"}
	col := []string{"id", "name", "address", "manager", "phone", "remark", "start_time", "end_time", "email"}
	return dao.TenantDao.SelectListExport(param, head, col)
}
