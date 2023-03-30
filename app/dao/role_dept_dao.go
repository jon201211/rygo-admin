package dao

import (
	"rygo/app/db"
	"rygo/app/model"

	"xorm.io/core"
)

// Fill with you ideas below.
var RoleDeptDao = newRoleDeptDao()

func newRoleDeptDao() *roleDeptDao {
	return &roleDeptDao{}
}

type roleDeptDao struct {
}

//映射数据表
func (d *roleDeptDao) TableName() string {
	return "sys_role_dept"
}

// 插入数据
func (d *roleDeptDao) Insert(r *model.RoleDeptEntity) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).Insert(r)
}

// 更新数据
func (d *roleDeptDao) Update(r *model.RoleDeptEntity) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).ID(core.PK{r.RoleId, r.DeptId}).Update(r)
}

// 删除
func (d *roleDeptDao) Delete(r *model.RoleDeptEntity) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).ID(core.PK{r.RoleId, r.DeptId}).Delete(r)
}

// 根据结构体中已有的非空数据来获得单条数据
func (d *roleDeptDao) FindOne(r *model.RoleDeptEntity) (bool, error) {
	return db.Instance().Engine().Table(d.TableName()).Get(r)
}

// 根据条件查询
func (d *roleDeptDao) Find(where, order string) ([]model.RoleDeptEntity, error) {
	var list []model.RoleDeptEntity
	err := db.Instance().Engine().Table(d.TableName()).Where(where).OrderBy(order).Find(&list)
	return list, err
}

//指定字段集合查询
func (d *roleDeptDao) FindIn(column string, args ...interface{}) ([]model.RoleDeptEntity, error) {
	var list []model.RoleDeptEntity
	err := db.Instance().Engine().Table(d.TableName()).In(column, args).Find(&list)
	return list, err
}

//排除指定字段集合查询
func (d *roleDeptDao) FindNotIn(column string, args ...interface{}) ([]model.RoleDeptEntity, error) {
	var list []model.RoleDeptEntity
	err := db.Instance().Engine().Table(d.TableName()).NotIn(column, args).Find(&list)
	return list, err
}
