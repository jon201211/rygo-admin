package dao

import (
	"rygo/app/db"
	"rygo/app/model"

	"xorm.io/core"
)

// Fill with you ideas below.
var RoleMenuDao = newRoleMenuDao()

func newRoleMenuDao() *roleMenuDao {
	return &roleMenuDao{}
}

type roleMenuDao struct {
}

//映射数据表
func (d *roleMenuDao) TableName() string {
	return "sys_role_menu"
}

// 插入数据
func (d *roleMenuDao) Insert(r *model.RoleMenuEntity) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).Insert(r)
}

// 更新数据
func (d *roleMenuDao) Update(r *model.RoleMenuEntity) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).ID(core.PK{r.RoleId, r.MenuId}).Update(r)
}

// 删除
func (d *roleMenuDao) Delete(r *model.RoleMenuEntity) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).ID(core.PK{r.RoleId, r.MenuId}).Delete(r)
}

// 根据结构体中已有的非空数据来获得单条数据
func (d *roleMenuDao) FindOne(r *model.RoleMenuEntity) (bool, error) {
	return db.Instance().Engine().Table(d.TableName()).Get(r)
}

// 根据条件查询
func (d *roleMenuDao) Find(where, order string) ([]model.RoleMenuEntity, error) {
	var list []model.RoleMenuEntity
	err := db.Instance().Engine().Table(d.TableName()).Where(where).OrderBy(order).Find(&list)
	return list, err
}

//指定字段集合查询
func (d *roleMenuDao) FindIn(column string, args ...interface{}) ([]model.RoleMenuEntity, error) {
	var list []model.RoleMenuEntity
	err := db.Instance().Engine().Table(d.TableName()).In(column, args).Find(&list)
	return list, err
}

//排除指定字段集合查询
func (d *roleMenuDao) FindNotIn(column string, args ...interface{}) ([]model.RoleMenuEntity, error) {
	var list []model.RoleMenuEntity
	err := db.Instance().Engine().Table(d.TableName()).NotIn(column, args).Find(&list)
	return list, err
}
