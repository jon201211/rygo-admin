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

var DictDataService = newDictDataService()

func newDictDataService() *dictDataService {
	return &dictDataService{}
}

type dictDataService struct {
}

//根据主键查询数据
func (s *dictDataService) SelectRecordById(id int64) (*model.SysDictData, error) {
	entity := &model.SysDictData{DictCode: id}
	_, err := dao.DictDataDao.FindOne(entity)
	return entity, err
}

//根据主键删除数据
func (s *dictDataService) DeleteRecordById(id int64) bool {
	entity := &model.SysDictData{DictCode: id}
	rs, _ := dao.DictDataDao.Delete(entity)
	if rs > 0 {
		return true
	}
	return false
}

//批量删除数据记录
func (s *dictDataService) DeleteRecordByIds(ids string) int64 {
	ida := convert.ToInt64Array(ids, ",")
	result, err := dao.DictDataDao.DeleteBatch(ida...)
	if err != nil {
		return 0
	}
	return result
}

//添加数据
func (s *dictDataService) AddSave(req *model.DictDataAddReq, ctx *gin.Context) (int64, error) {
	var entity model.SysDictData
	entity.DictType = req.DictType
	entity.Status = req.Status
	entity.DictLabel = req.DictLabel
	entity.CssClass = req.CssClass
	entity.DictSort = req.DictSort
	entity.DictValue = req.DictValue
	entity.IsDefault = req.IsDefault
	entity.ListClass = req.ListClass
	entity.Remark = req.Remark
	entity.CreateTime = time.Now()
	entity.CreateBy = ""

	user := UserService.GetProfile(ctx)

	if user != nil {
		entity.CreateBy = user.LoginName
	}

	_, err := dao.DictDataDao.Insert(&entity)

	return entity.DictCode, err
}

//修改数据
func (s *dictDataService) EditSave(req *model.DictDataEditReq, ctx *gin.Context) (int64, error) {
	entity := &model.SysDictData{DictCode: req.DictCode}
	ok, err := dao.DictDataDao.FindOne(entity)

	if err != nil || !ok {
		return 0, err
	}

	if entity == nil {
		return 0, errors.New("数据不存在")
	}

	entity.DictType = req.DictType
	entity.Status = req.Status
	entity.DictLabel = req.DictLabel
	entity.CssClass = req.CssClass
	entity.DictSort = req.DictSort
	entity.DictValue = req.DictValue
	entity.IsDefault = req.IsDefault
	entity.ListClass = req.ListClass
	entity.Remark = req.Remark
	entity.UpdateTime = time.Now()
	entity.UpdateBy = ""

	user := UserService.GetProfile(ctx)

	if user == nil {
		entity.UpdateBy = user.LoginName
	}

	return dao.DictDataDao.Update(entity)
}

//根据条件分页查询角色数据
func (s *dictDataService) SelectListAll(params *model.DictDataSelectPageReq) ([]model.SysDictData, error) {
	return dao.DictDataDao.SelectListAll(params)
}

//根据条件分页查询角色数据
func (s *dictDataService) SelectListByPage(params *model.DictDataSelectPageReq) (*[]model.SysDictData, *page.Paging, error) {
	return dao.DictDataDao.SelectListByPage(params)
}

// 导出excel
func (s *dictDataService) Export(param *model.DictDataSelectPageReq) (string, error) {
	head := []string{"字典编码", "字典排序", "字典标签", "字典键值", "字典类型", "样式属性", "表格回显样式", "是否默认", "状态", "创建者", "创建时间", "更新者", "更新时间", "备注"}
	col := []string{"dict_code", "dict_sort", "dict_label", "dict_value", "dict_type", "css_class", "list_class", "is_default", "status", "create_by", "create_time", "update_by", "update_time", "remark"}
	return dao.DictDataDao.SelectListExport(param, head, col)
}
