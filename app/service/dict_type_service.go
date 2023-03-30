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

var DictTypeService = newDictTypeService()

func newDictTypeService() *dictTypeService {
	return &dictTypeService{}
}

type dictTypeService struct {
}

//根据主键查询数据
func (s *dictTypeService) SelectRecordById(id int64) (*model.SysDictType, error) {
	entity := &model.SysDictType{DictId: id}
	_, err := dao.DictTypeDao.FindOne(entity)
	return entity, err
}

//根据主键删除数据
func (s *dictTypeService) DeleteRecordById(id int64) bool {
	entity := &model.SysDictType{DictId: id}
	rs, err := dao.DictTypeDao.Delete(entity)
	if err == nil {
		if rs > 0 {
			return true
		}
	}
	return false
}

//批量删除数据记录
func (s *dictTypeService) DeleteRecordByIds(ids string) int64 {
	ida := convert.ToInt64Array(ids, ",")
	result, err := dao.DictTypeDao.DeleteBatch(ida...)
	if err != nil {
		return 0
	}
	return result
}

//添加数据
func (s *dictTypeService) AddSave(req *model.DictTypeAddReq, ctx *gin.Context) (int64, error) {
	var entity model.SysDictType
	entity.Status = req.Status
	entity.DictType = req.DictType
	entity.DictName = req.DictName
	entity.Remark = req.Remark
	entity.CreateTime = time.Now()
	entity.CreateBy = ""

	user := UserService.GetProfile(ctx)

	if user != nil {
		entity.CreateBy = user.LoginName
	}

	_, err := dao.DictTypeDao.Insert(&entity)

	return entity.DictId, err
}

//修改数据
func (s *dictTypeService) EditSave(req *model.DictTypeEditReq, ctx *gin.Context) (int64, error) {
	entity := &model.SysDictType{DictId: req.DictId}
	ok, err := dao.DictTypeDao.FindOne(entity)

	if err != nil || !ok {
		return 0, err
	}

	if entity == nil {
		return 0, errors.New("数据不存在")
	}
	entity.Status = req.Status
	entity.DictType = req.DictType
	entity.DictName = req.DictName
	entity.Remark = req.Remark
	entity.UpdateTime = time.Now()
	entity.UpdateBy = ""

	user := UserService.GetProfile(ctx)

	if user == nil {
		entity.UpdateBy = user.LoginName
	}

	return dao.DictTypeDao.Update(entity)
}

//根据条件分页查询角色数据
func (s *dictTypeService) SelectListAll(params *model.DictTypeSelectPageReq) ([]model.SysDictType, error) {
	return dao.DictTypeDao.SelectListAll(params)
}

//根据条件分页查询角色数据
func (s *dictTypeService) SelectListByPage(params *model.DictTypeSelectPageReq) ([]model.SysDictType, *page.Paging, error) {
	return dao.DictTypeDao.SelectListByPage(params)
}

//根据字典类型查询信息
func (s *dictTypeService) SelectDictTypeByType(dictType string) *model.SysDictType {
	entity := &model.SysDictType{DictType: dictType}
	ok, err := dao.DictTypeDao.FindOne(entity)
	if err != nil || !ok {
		return nil
	}
	return entity
}

// 导出excel
func (s *dictTypeService) Export(param *model.DictTypeSelectPageReq) (string, error) {
	head := []string{"字典主键", "字典名称", "字典类型", "状态", "创建者", "创建时间", "更新者", "更新时间", "备注"}
	col := []string{"dict_id", "dict_name", "dict_type", "status", "create_by", "create_time", "update_by", "update_time", "remark"}
	return dao.DictTypeDao.SelectListExport(param, head, col)
}

//检查字典类型是否唯一
func (s *dictTypeService) CheckDictTypeUniqueAll(configKey string) string {
	entity, err := dao.DictTypeDao.CheckDictTypeUniqueAll(configKey)
	if err != nil {
		return "1"
	}
	if entity != nil && entity.DictId > 0 {
		return "1"
	}
	return "0"
}

//检查字典类型是否唯一
func (s *dictTypeService) CheckDictTypeUnique(configKey string, dictId int64) string {
	entity, err := dao.DictTypeDao.CheckDictTypeUniqueAll(configKey)
	if err != nil {
		return "1"
	}
	if entity != nil && entity.DictId > 0 && entity.DictId != dictId {
		return "1"
	}
	return "0"
}

//查询字典类型树
func (s *dictTypeService) SelectDictTree(params *model.DictTypeSelectPageReq) *[]model.Ztree {
	var result []model.Ztree
	dictList, err := dao.DictTypeDao.SelectListAll(params)
	if err == nil && dictList != nil {
		for _, item := range dictList {
			var tmp model.Ztree
			tmp.Id = item.DictId
			tmp.Name = s.transDictName(item)
			tmp.Title = item.DictType
			result = append(result, tmp)
		}
	}
	return &result
}

func (s *dictTypeService) transDictName(entity model.SysDictType) string {
	return `(` + entity.DictName + `)&nbsp;&nbsp;&nbsp;` + entity.DictType
}
