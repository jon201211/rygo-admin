package dao

import (
	"errors"
	"rygo/app/db"
	"rygo/app/model"
)

//数据库字符串类型
var COLUMNTYPE_STR = []string{"char", "varchar", "narchar", "varchar2", "tinytext", "text", "mediumtext", "longtext"}

//数据库时间类型
var COLUMNTYPE_TIME = []string{"datetime", "time", "date", "timestamp"}

//数据库数字类型
var COLUMNTYPE_NUMBER = []string{"tinyint", "smallint", "mediumint", "int", "number", "integer", "bigint", "float", "float", "double", "decimal"}

//页面不需要编辑字段
var COLUMNNAME_NOT_EDIT = []string{"id", "create_by", "create_time", "del_flag", "update_by", "update_time"}

//页面不需要显示的列表字段
var COLUMNNAME_NOT_LIST = []string{"id", "create_by", "create_time", "del_flag", "update_by", "update_time"}

//页面不需要查询字段
var COLUMNNAME_NOT_QUERY = []string{"id", "create_by", "create_time", "del_flag", "update_by", "update_time", "remark"}

var TableColumnDao = newTableColumnDao()

func newTableColumnDao() *tableColumnDao {
	return &tableColumnDao{}
}

type tableColumnDao struct {
}

//映射数据表
func (d *tableColumnDao) TableName() string {
	return "gen_table_column"
}

// 插入数据
func (d *tableColumnDao) Insert(r *model.TableColumnEntity) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).Insert(r)
}

// 更新数据
func (d *tableColumnDao) Update(r *model.TableColumnEntity) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).ID(r.ColumnId).Update(r)
}

// 删除
func (d *tableColumnDao) Delete(r *model.TableColumnEntity) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).ID(r.ColumnId).Delete(r)
}

//批量删除
func (d *tableColumnDao) DeleteBatch(ids ...int64) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).In("column_id", ids).Delete(new(model.TableColumnEntity))
}

// 根据结构体中已有的非空数据来获得单条数据
func (d *tableColumnDao) FindOne(r *model.TableColumnEntity) (bool, error) {
	return db.Instance().Engine().Table(d.TableName()).Get(r)
}

// 根据条件查询
func (d *tableColumnDao) Find(where, order string) ([]model.TableColumnEntity, error) {
	var list []model.TableColumnEntity
	err := db.Instance().Engine().Table(d.TableName()).Where(where).OrderBy(order).Find(&list)
	return list, err
}

//指定字段集合查询
func (d *tableColumnDao) FindIn(column string, args ...interface{}) ([]model.TableColumnEntity, error) {
	var list []model.TableColumnEntity
	err := db.Instance().Engine().Table(d.TableName()).In(column, args).Find(&list)
	return list, err
}

//排除指定字段集合查询
func (d *tableColumnDao) FindNotIn(column string, args ...interface{}) ([]model.TableColumnEntity, error) {
	var list []model.TableColumnEntity
	err := db.Instance().Engine().Table(d.TableName()).NotIn(column, args).Find(&list)
	return list, err
}

//查询业务字段列表
func (d *tableColumnDao) SelectGenTableColumnListByTableId(tableId int64) ([]model.TableColumnEntity, error) {
	db := db.Instance().Engine()

	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}

	var result []model.TableColumnEntity

	session := db.Table(d.TableName()).Alias("t").Where("table_id=?", tableId).OrderBy("sort")
	session.Find(&result)
	return result, nil
}

//根据表名称查询列信息
func (d *tableColumnDao) SelectDbTableColumnsByName(tableName string) ([]model.TableColumnEntity, error) {
	db := db.Instance().Engine()

	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}

	var result []model.TableColumnEntity

	session := db.Table("information_schema.columns")
	session.Where("table_schema = (select database())")
	session.Where("table_name=?", tableName).OrderBy("ordinal_position")
	session.Select("column_name, (case when (is_nullable = 'no' && column_key != 'PRI') then '1' else null end) as is_required, (case when column_key = 'PRI' then '1' else '0' end) as is_pk, ordinal_position as sort, column_comment, (case when extra = 'auto_increment' then '1' else '0' end) as is_increment, column_type")
	session.Find(&result)
	return result, nil
}

//判断string 是否存在在数组中
func (d *tableColumnDao) IsExistInArray(value string, array []string) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}

//判断是否是数据库字符串类型
func (d *tableColumnDao) IsStringObject(value string) bool {
	return d.IsExistInArray(value, COLUMNTYPE_STR)
}

//判断是否是数据库时间类型
func (d *tableColumnDao) IsTimeObject(value string) bool {
	return d.IsExistInArray(value, COLUMNTYPE_TIME)
}

//判断是否是数据库数字类型
func (d *tableColumnDao) IsNumberObject(value string) bool {
	return d.IsExistInArray(value, COLUMNTYPE_NUMBER)
}

//页面不需要编辑字段
func (d *tableColumnDao) IsNotEdit(value string) bool {
	return !d.IsExistInArray(value, COLUMNNAME_NOT_EDIT)
}

//页面不需要显示的列表字段
func (d *tableColumnDao) IsNotList(value string) bool {
	return !d.IsExistInArray(value, COLUMNNAME_NOT_LIST)
}

//页面不需要查询字段
func (d *tableColumnDao) IsNotQuery(value string) bool {
	return !d.IsExistInArray(value, COLUMNNAME_NOT_QUERY)
}
