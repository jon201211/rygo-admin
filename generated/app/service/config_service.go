// ==========================================================================
// RYGO auto gen code!
// datetime：2020-03-27 04:35:17 +0800 CST
// ==========================================================================
package config

import (
	"errors"
	"github.com/gin-gonic/gin"
	"time"
	yj-app/app/model
	yj-app/app/dao

	"yj-app/app/utils/convert"
	"yj-app/app/utils/page"
)


var ConfigService = newConfigService()

func newConfigService() *configService {
	return &configService{}
}

type configService struct {
}


//根据主键查询数据
func (s *configService) SelectRecordById(id int) (*model.configconfig, error) {
	entity := &model.configconfig{ ConfigId: id}
	_, err := dao.configDao.FindOne(entity)
	return entity, err
}

//根据主键删除数据
func (s *configService) DeleteRecordById(id int) bool {
	rs, err := (&model.configconfig{ ConfigId: id}).Delete()
	if err == nil {
		if rs > 0 {
			return true
		}
	}
	return false
}

//批量删除数据记录
func (s *configService) DeleteRecordByIds(ids string) int64 {
	ida := convert.ToInt64Array(ids, ",")
	result, err := dao.configDao.DeleteBatch(ida...)
	if err != nil {
		return 0
	}
	return result
}

//添加数据
func (s *configService) AddSave(req *model.configAddReq, ctx *gin.Context) (int, error) {
	var entity model.configconfig
	 
	entity.ConfigId = req.ConfigId  
	entity.ConfigName = req.ConfigName  
	entity.ConfigKey = req.ConfigKey  
	entity.ConfigValue = req.ConfigValue  
	entity.ConfigType = req.ConfigType          
	entity.Remark = req.Remark 
	entity.CreateTime = time.Now()
	entity.CreateBy = ""

	user := UserService.GetProfile(ctx)

	if user != nil {
		entity.CreateBy = user.LoginName
	}

	_, err := entity.Insert()
	return entity.ConfigId, err
}

//修改数据
func (s *configService) EditSave(req *model.configEditReq, ctx *gin.Context) (int64, error) {
	entity := &model.configconfig{ ConfigId: req.ConfigId }
	ok, err := dao.configDao.FindOne(entity)

	if err != nil {
		return 0, err
	}

	if !ok {
		return 0, errors.New("数据不存在")
	}

	   
	entity.ConfigName = req.ConfigName  
	entity.ConfigKey = req.ConfigKey  
	entity.ConfigValue = req.ConfigValue  
	entity.ConfigType = req.ConfigType          
	entity.Remark = req.Remark 
	entity.UpdateTime = time.Now()
	entity.UpdateBy = ""

	user := UserService.GetProfile(ctx)

	if user == nil {
		entity.UpdateBy = user.LoginName
	}

	return entity.Update()
}

//根据条件查询数据
func (s *configService) SelectListAll(params *model.configSelectPageReq) ([]model.configconfig, error) {
	return dao.configDao.SelectListAll(params)
}

//根据条件分页查询数据
func (s *configService) (s *configService) SelectListByPage(params *model.configSelectPageReq) ([]model.configconfig, *page.Paging, error) {
	return dao.configDao.SelectListByPage(params)
}

// 导出excel
func (s *configService) Export(param *model.configSelectPageReq) (string, error) {
	head := []string{  "参数主键11" ,"参数名称111" ,"参数键名111" ,"参数键值" ,"系统内置（Y是 N否）" ,"创建者" ,"创建时间" ,"更新者" ,"更新时间" ,"备注"}
	col := []string{  "config_id" ,"config_name" ,"config_key" ,"config_value" ,"config_type" ,"create_by" ,"create_time" ,"update_by" ,"update_time" ,"remark"}
	return dao.configDao.SelectListExport(param, head, col)
}