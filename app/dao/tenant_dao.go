package dao

import (
	"errors"
	"rygo/app/db"
	"rygo/app/model"
	"rygo/app/utils/excel"
	"rygo/app/utils/page"

	"xorm.io/builder"
)

var TenantDao = newTenantDao()

func newTenantDao() *tenantDao {
	return &tenantDao{}
}

type tenantDao struct {
}

//映射数据表
func (d *tenantDao) TableName() string {
	return "sys_tenant"
}

// 插入数据
func (d *tenantDao) Insert(e *model.SysTenant) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).Insert(e)
}

// 更新数据
func (d *tenantDao) Update(e *model.SysTenant) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).ID(e.Id).Update(e)
}

// 删除
func (d *tenantDao) Delete(e *model.SysTenant) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).ID(e.Id).Delete(e)
}

//批量删除
func (d *tenantDao) DeleteBatch(ids ...int64) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).In("id", ids).Delete(new(model.SysTenant))
}

// 根据结构体中已有的非空数据来获得单条数据
func (d *tenantDao) FindOne(e *model.SysTenant) (bool, error) {
	return db.Instance().Engine().Table(d.TableName()).Get(e)
}

// 根据条件查询
func (d *tenantDao) Find(where, order string) ([]model.SysTenant, error) {
	var list []model.SysTenant
	err := db.Instance().Engine().Table(d.TableName()).Where(where).OrderBy(order).Find(&list)
	return list, err
}

//指定字段集合查询
func (d *tenantDao) FindIn(column string, args ...interface{}) ([]model.SysTenant, error) {
	var list []model.SysTenant
	err := db.Instance().Engine().Table(d.TableName()).In(column, args).Find(&list)
	return list, err
}

//排除指定字段集合查询
func (d *tenantDao) FindNotIn(column string, args ...interface{}) ([]model.SysTenant, error) {
	var list []model.SysTenant
	err := db.Instance().Engine().Table(d.TableName()).NotIn(column, args).Find(&list)
	return list, err
}

//根据条件分页查询数据
func (d *tenantDao) SelectListByPage(param *model.TenantSelectPageReq) ([]model.SysTenant, *page.Paging, error) {
	db := db.Instance().Engine()
	paging := new(page.Paging)

	if db == nil {
		return nil, paging, errors.New("获取数据库连接失败")
	}

	session := db.Table("sys_tenant").Alias("t")

	if param != nil {

		if param.Name != "" {
			session.Where("t.name like ?", "%"+param.Name+"%")
		}

		if param.Address != "" {
			session.Where("t.address = ?", param.Address)
		}
		if param.BeginTime != "" {
			session.Where("t.create_time >= ?", param.BeginTime)
		}

		if param.EndTime != "" {
			session.Where("t.create_time<=?", param.EndTime)
		}
	}

	total, err := session.Clone().Count()
	if err != nil {
		return nil, paging, errors.New("读取行数失败")
	}

	paging = page.CreatePaging(param.PageNum, param.PageSize, int(total))
	session.Limit(paging.Pagesize, paging.StartNum)

	var result []model.SysTenant
	err = session.Find(&result)
	return result, paging, err
}

// 导出excel
func (d *tenantDao) SelectListExport(param *model.TenantSelectPageReq, head, col []string) (string, error) {
	db := db.Instance().Engine()

	if db == nil {
		return "", errors.New("获取数据库连接失败")
	}

	build := builder.Select(col...).From("sys_tenant", "t")

	if param != nil {

		if param.Name != "" {
			build.Where(builder.Like{"t.name", param.Name})
		}

		if param.Address != "" {
			build.Where(builder.Eq{"t.address": param.Address})
		}
		if param.BeginTime != "" {
			build.Where(builder.Gte{"date_format(t.create_time,'%y%m%d')": "date_format('" + param.BeginTime + "','%y%m%d')"})
		}

		if param.EndTime != "" {
			build.Where(builder.Lte{"date_format(t.create_time,'%y%m%d')": "date_format('" + param.EndTime + "','%y%m%d')"})
		}
	}

	sqlStr, _, _ := build.ToSQL()
	arr, err := db.SQL(sqlStr).QuerySliceString()

	path, err := excel.DownlaodExcel(head, arr)

	return path, err
}

//获取所有数据
func (d *tenantDao) SelectListAll(param *model.TenantSelectPageReq) ([]model.SysTenant, error) {
	db := db.Instance().Engine()

	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}

	session := db.Table("sys_tenant").Alias("t")

	if param != nil {

		if param.Name != "" {
			session.Where("t.name like ?", "%"+param.Name+"%")
		}

		if param.Address != "" {
			session.Where("t.address = ?", param.Address)
		}

		if param.BeginTime != "" {
			session.Where("date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			session.Where("date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	var result []model.SysTenant
	err := session.Find(&result)
	return result, err
}
