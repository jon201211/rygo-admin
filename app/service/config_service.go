package service

import (
	"errors"
	"rygo/app/cache"
	"rygo/app/dao"
	"rygo/app/model"

	"rygo/app/utils/convert"
	"rygo/app/utils/page"
	"time"

	"github.com/gin-gonic/gin"
)

var ConfigService = newConfigService()

func newConfigService() *configService {
	return &configService{}
}

type configService struct {
}

func (s *configService) GetOssUrl() string {
	return s.GetValueByKey("sys.resource.url")
}

//根据用户id和权限字符串判断是否有此权限
func (s *configService) AddInt(a, b int) int {
	return a + b
}

//根据键获取值
func (s *configService) GetValueByKey(key string) string {
	resultStr := ""
	//从缓存读取
	c := cache.Instance()
	result, ok := c.Get(key)

	if ok {
		return result.(string)
	}

	if result == nil {
		entity := &model.SysConfig{ConfigKey: key}
		ok, _ := dao.ConfigDao.FindOne(entity)
		if !ok {
			return ""
		}

		resultStr = entity.ConfigValue
		c.Set(key, resultStr, 0)
	} else {
		resultStr = result.(string)
	}

	return resultStr
}

//根据主键查询数据
func (s *configService) SelectRecordById(id int64) (*model.SysConfig, error) {
	entity := &model.SysConfig{ConfigId: id}
	_, err := dao.ConfigDao.FindOne(entity)
	return entity, err
}

//根据主键删除数据
func (s *configService) DeleteRecordById(id int64) bool {
	entity := &model.SysConfig{ConfigId: id}
	ok, _ := dao.ConfigDao.FindOne(entity)
	if ok {
		result, err := dao.ConfigDao.Delete(entity)
		if err == nil {
			if result > 0 {
				//从缓存删除
				c := cache.Instance()
				c.Delete(entity.ConfigKey)
				return true
			}
		}
	}
	return false
}

//批量删除数据记录
func (s *configService) DeleteRecordByIds(ids string) int64 {
	idarr := convert.ToInt64Array(ids, ",")
	list, _ := dao.ConfigDao.FindIn("config_id", idarr)
	rs, err := dao.ConfigDao.DeleteBatch(idarr...)
	if err != nil {
		return 0
	}

	if len(list) > 0 {
		for _, item := range list {
			//从缓存删除
			c := cache.Instance()
			c.Delete(item.ConfigKey)
		}
	}

	return rs
}

//添加数据
func (s *configService) AddSave(req *model.ConfigAddReq, ctx *gin.Context) (int64, error) {
	var entity model.SysConfig
	entity.ConfigName = req.ConfigName
	entity.ConfigKey = req.ConfigKey
	entity.ConfigType = req.ConfigType
	entity.ConfigValue = req.ConfigValue
	entity.Remark = req.Remark
	entity.CreateTime = time.Now()
	entity.CreateBy = ""

	user := UserService.GetProfile(ctx)

	if user != nil {
		entity.CreateBy = user.LoginName
	}

	_, err := dao.ConfigDao.Insert(&entity)
	return entity.ConfigId, err
}

//修改数据
func (s *configService) EditSave(req *model.ConfigEditReq, ctx *gin.Context) (int64, error) {
	entity := &model.SysConfig{ConfigId: req.ConfigId}
	ok, err := dao.ConfigDao.FindOne(entity)

	if err != nil {
		return 0, err
	}

	if !ok {
		return 0, errors.New("数据不存在")
	}

	entity.ConfigName = req.ConfigName
	entity.ConfigKey = req.ConfigKey
	entity.ConfigValue = req.ConfigValue
	entity.Remark = req.Remark
	entity.ConfigType = req.ConfigType
	entity.UpdateTime = time.Now()
	entity.UpdateBy = ""

	user := UserService.GetProfile(ctx)

	if user == nil {
		entity.UpdateBy = user.LoginName
	}

	rs, err := dao.ConfigDao.Update(entity)

	if err != nil {
		return 0, err
	}

	//保存到缓存
	cache := cache.Instance()
	cache.Set(entity.ConfigKey, entity.ConfigValue, 0)

	return rs, nil
}

//根据条件分页查询角色数据
func (s *configService) SelectListAll(params *model.SelectPageReq) ([]model.SysConfig, error) {
	return dao.ConfigDao.SelectListAll(params)
}

//根据条件分页查询角色数据
func (s *configService) SelectListByPage(params *model.SelectPageReq) ([]model.SysConfig, *page.Paging, error) {
	return dao.ConfigDao.SelectListByPage(params)
}

// 导出excel
func (s *configService) Export(param *model.SelectPageReq) (string, error) {
	head := []string{"参数主键", "参数名称", "参数键名", "参数键值", "系统内置（Y是 N否）", "状态"}
	col := []string{"config_id", "config_name", "config_key", "config_value", "config_type"}
	return dao.ConfigDao.SelectListExport(param, head, col)
}

//检查角色名是否唯一
func (s *configService) CheckConfigKeyUniqueAll(configKey string) string {
	entity, err := dao.ConfigDao.CheckPostCodeUniqueAll(configKey)
	if err != nil {
		return "1"
	}
	if entity != nil && entity.ConfigId > 0 {
		return "1"
	}
	return "0"
}

//检查岗位名称是否唯一
func (s *configService) CheckConfigKeyUnique(configKey string, configId int64) string {
	entity, err := dao.ConfigDao.CheckPostCodeUniqueAll(configKey)
	if err != nil {
		return "1"
	}
	if entity != nil && entity.ConfigId > 0 && entity.ConfigId != configId {
		return "1"
	}
	return "0"
}
