package service

import (
	"rygo/app/dao"
	"rygo/app/model"
	"rygo/app/utils/convert"
)

var TableColumnService = newTableColumnService()

func newTableColumnService() *tableColumnService {
	return &tableColumnService{}
}

type tableColumnService struct {
}

//新增业务字段
func (s *tableColumnService) Insert(entity *model.TableColumnEntity) (int64, error) {
	_, err := dao.TableColumnDao.Insert(entity)
	if err != nil {
		return 0, err
	}
	return entity.ColumnId, err
}

//修改业务字段
func (s *tableColumnService) Update(entity *model.TableColumnEntity) (int64, error) {
	return dao.TableColumnDao.Update(entity)
}

//根据主键查询数据
func (s *tableColumnService) SelectRecordById(id int64) (*model.TableColumnEntity, error) {
	entity := &model.TableColumnEntity{ColumnId: id}
	_, err := dao.TableColumnDao.FindOne(entity)
	return entity, err
}

//根据主键删除数据
func (s *tableColumnService) DeleteRecordById(id int64) bool {
	entity := &model.TableColumnEntity{ColumnId: id}
	rs, err := dao.TableColumnDao.Delete(entity)
	if err == nil && rs > 0 {
		return true
	}
	return false
}

//批量删除数据记录
func (s *tableColumnService) DeleteRecordByIds(ids string) int64 {
	idarr := convert.ToInt64Array(ids, ",")
	result, err := dao.TableColumnDao.DeleteBatch(idarr...)
	if err != nil {
		return 0
	}
	return result
}

//查询业务字段列表
func (s *tableColumnService) SelectGenTableColumnListByTableId(tableId int64) ([]model.TableColumnEntity, error) {
	return dao.TableColumnDao.SelectGenTableColumnListByTableId(tableId)
}

//根据表名称查询列信息
func (s *tableColumnService) SelectDbTableColumnsByName(tableName string) ([]model.TableColumnEntity, error) {
	return dao.TableColumnDao.SelectDbTableColumnsByName(tableName)
}
