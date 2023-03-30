package dao

import (
	"errors"
	"rygo/app/db"
	"rygo/app/model"

	"rygo/app/utils/page"
)

var TableDao = newTableDao()

func newTableDao() *tableDao {
	return &tableDao{}
}

type tableDao struct {
}

//映射数据表
func (d *tableDao) TableName() string {
	return "gen_table"
}

// 插入数据
func (d *tableDao) Insert(r *model.TableEntity) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).Insert(r)
}

// 更新数据
func (d *tableDao) Update(r *model.TableEntity) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).ID(r.TableId).Update(r)
}

// 删除
func (d *tableDao) Delete(r *model.TableEntity) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).ID(r.TableId).Delete(r)
}

//批量删除
func (d *tableDao) DeleteBatch(ids ...int64) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).In("table_id", ids).Delete(new(model.TableEntity))
}

// 根据结构体中已有的非空数据来获得单条数据
func (d *tableDao) FindOne(r *model.TableEntity) (bool, error) {
	return db.Instance().Engine().Table(d.TableName()).Get(r)
}

// 根据条件查询
func (d *tableDao) Find(where, order string) ([]model.TableEntity, error) {
	var list []model.TableEntity
	err := db.Instance().Engine().Table(d.TableName()).Where(where).OrderBy(order).Find(&list)
	return list, err
}

//指定字段集合查询
func (d *tableDao) FindIn(column string, args ...interface{}) ([]model.TableEntity, error) {
	var list []model.TableEntity
	err := db.Instance().Engine().Table(d.TableName()).In(column, args).Find(&list)
	return list, err
}

//排除指定字段集合查询
func (d *tableDao) FindNotIn(column string, args ...interface{}) ([]model.TableEntity, error) {
	var list []model.TableEntity
	err := db.Instance().Engine().Table(d.TableName()).NotIn(column, args).Find(&list)
	return list, err
}

//根据ID获取记录
func (d *tableDao) SelectRecordById(id int64) (*model.TableEntityExtend, error) {
	db := db.Instance().Engine()
	var result model.TableEntityExtend
	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}

	session := db.Table(d.TableName()).Where("table_id=?", id)
	ok, err := session.Get(&result)
	if !ok {
		return nil, err
	}

	//表数据列
	columModel := db.Table("gen_table_column").Where("table_id=?", id)

	var columList []model.TableColumnEntity
	err = columModel.Find(&columList)

	if err != nil {
		return nil, err
	}
	result.Columns = columList
	return &result, nil
}

//根据条件分页查询数据
func (d *tableDao) SelectListByPage(param *model.TableSelectPageReq) ([]model.TableEntity, *page.Paging, error) {
	db := db.Instance().Engine()
	p := new(page.Paging)
	if db == nil {
		return nil, p, errors.New("获取数据库连接失败")
	}

	session := db.Table(d.TableName()).Alias("t")

	if param != nil {
		if param.TableName != "" {
			session.Where("t.table_name like ?", "%"+param.TableName+"%")
		}

		if param.TableComment != "" {
			session.Where("t.table_comment like ?", "%"+param.TableComment+"%")
		}

		if param.BeginTime != "" {
			session.Where("date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			session.Where("date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	tm := session.Clone()

	total, err := tm.Count()

	if err != nil {
		return nil, p, errors.New("读取行数失败")
	}

	p = page.CreatePaging(param.PageNum, param.PageSize, int(total))

	session.Limit(p.Pagesize, p.StartNum)
	var result []model.TableEntity
	err = session.Find(&result)

	return result, p, err
}

//查询据库列表
func (d *tableDao) SelectDbTableList(param *model.TableSelectPageReq) ([]model.TableEntity, *page.Paging, error) {
	db := db.Instance().Engine()
	p := new(page.Paging)
	if db == nil {
		return nil, p, errors.New("获取数据库连接失败")
	}

	session := db.Table("information_schema.tables")
	session.Where("table_schema = (select database())")
	session.Where("table_name NOT LIKE 'qrtz_%' AND table_name NOT LIKE 'gen_%'")
	session.Where("table_name NOT IN (select table_name from gen_table)")
	if param != nil {
		if param.TableName != "" {
			session.Where("lower(table_name) like lower(?)", "%"+param.TableName+"%")
		}

		if param.TableComment != "" {
			session.Where("lower(table_comment) like lower(?)", "%"+param.TableComment+"%")
		}

		if param.BeginTime != "" {
			session.Where("date_format(create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			session.Where("date_format(create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	tm := session.Clone()

	total, err := tm.Count()

	if err != nil {
		return nil, p, errors.New("读取行数失败")
	}

	p = page.CreatePaging(param.PageNum, param.PageSize, int(total))

	session.Select("table_name, table_comment, create_time, update_time")
	session.Limit(p.Pagesize, p.StartNum)

	var result []model.TableEntity
	err = session.Find(&result)
	return result, p, err
}

//查询据库列表
func (d *tableDao) SelectDbTableListByNames(tableNames []string) ([]model.TableEntity, error) {
	db := db.Instance().Engine()

	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}

	session := db.Table("information_schema.tables")
	session.Select("0 as table_id, table_name, table_comment,'' as class_name,'' as tpl_category,'' as package_name,'' as module_name,'' as business_name,'' as function_name,'' as function_author,'' as options,'' as create_by, create_time,'' as update_by, update_time,'' as remark")
	session.Where("table_name NOT LIKE 'qrtz_%'")
	session.Where("table_name NOT LIKE 'gen_%'")
	session.Where("table_schema = (select database())")
	if len(tableNames) > 0 {
		session.In("table_name", tableNames)
	}

	var result []model.TableEntity
	err := session.Find(&result)
	return result, err
}

//查询据库列表
func (d *tableDao) SelectTableByName(tableName string) (*model.TableEntity, error) {
	db := db.Instance().Engine()

	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}

	session := db.Table("information_schema.tables")
	session.Select("0 as table_id, table_name, table_comment,'' as class_name,'' as tpl_category,'' as package_name,'' as module_name,'' as business_name,'' as function_name,'' as function_author,'' as options,'' as create_by, create_time,'' as update_by, update_time,'' as remark")
	session.Where("table_comment <> ''")
	session.Where("table_schema = (select database())")
	if tableName != "" {
		session.Where("table_name = ?", tableName)
	}

	var result model.TableEntity
	_, err := session.Get(&result)
	return &result, err
}

//查询表ID业务信息
func (d *tableDao) SelectGenTableById(tableId int64) (*model.TableEntity, error) {
	db := db.Instance().Engine()

	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}

	session := db.Table(d.TableName()).Alias("t")
	session.Join("LEFT", []string{"gen_table_column", "c"}, "t.table_id = c.table_id")
	session.Where("t.table_id = ?", tableId)
	session.Select("t.table_id, t.table_name, t.table_comment, t.class_name, t.tpl_category, t.package_name, t.module_name, t.business_name, t.function_name, t.function_author, t.options, t.remark,c.column_id, c.column_name, c.column_comment, c.column_type, c.java_type, c.java_field, c.is_pk, c.is_increment, c.is_required, c.is_insert, c.is_edit, c.is_list, c.is_query, c.query_type, c.html_type, c.dict_type, c.sort")

	var result model.TableEntity
	_, err := session.Get(&result)
	return &result, err
}

func (d *tableDao) SelectableExtendById(tableId int64) (*model.TableEntityExtend, error) {
	db := db.Instance().Engine()

	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}

	session := db.Table(d.TableName()).Alias("t")
	session.Where("t.table_id = ?", tableId)

	var result model.TableEntityExtend
	_, err := session.Get(&result)
	return &result, err
}

//查询表名称业务信息
func (d *tableDao) SelectGenTableByName(tableName string) (*model.TableEntity, error) {
	db := db.Instance().Engine()

	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}

	session := db.Table(d.TableName()).Alias("t")
	session.Join("LEFT", []string{"gen_table_column", "c"}, "t.table_id = c.table_id")
	session.Where("t.table_name = ?", tableName)
	session.Select("t.table_id, t.table_name, t.table_comment, t.class_name, t.tpl_category, t.package_name, t.module_name, t.business_name, t.function_name, t.function_author, t.options, t.remark,c.column_id, c.column_name, c.column_comment, c.column_type, c.java_type, c.java_field, c.is_pk, c.is_increment, c.is_required, c.is_insert, c.is_edit, c.is_list, c.is_query, c.query_type, c.html_type, c.dict_type, c.sort")

	var result model.TableEntity
	_, err := session.Get(&result)
	return &result, err
}
