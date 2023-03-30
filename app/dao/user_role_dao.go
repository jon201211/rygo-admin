package dao

import (
	"rygo/app/db"
	"rygo/app/model"

	"xorm.io/core"
)

var UserRoleDao = newUserRoleDao()

func newUserRoleDao() *uerRoleDao {
	return &uerRoleDao{}
}

type uerRoleDao struct {
}

//映射数据表
func (d *uerRoleDao) TableName() string {
	return "sys_user_role"
}

// 插入数据
func (d *uerRoleDao) Insert(r *model.UserRoleEntity) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).Insert(r)
}

// 更新数据
func (d *uerRoleDao) Update(r *model.UserRoleEntity) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).ID(core.PK{r.UserId, r.RoleId}).Update(r)
}

// 删除
func (d *uerRoleDao) Delete(r *model.UserRoleEntity) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).ID(core.PK{r.UserId, r.RoleId}).Delete(r)
}

// 根据结构体中已有的非空数据来获得单条数据
func (d *uerRoleDao) FindOne(r *model.UserRoleEntity) (bool, error) {
	return db.Instance().Engine().Table(d.TableName()).Get(r)
}

// 根据条件查询
func (d *uerRoleDao) Find(where, order string) ([]model.UserRoleEntity, error) {
	var list []model.UserRoleEntity
	err := db.Instance().Engine().Table(d.TableName()).Where(where).OrderBy(order).Find(&list)
	return list, err
}

//指定字段集合查询
func (d *uerRoleDao) FindIn(column string, args ...interface{}) ([]model.UserRoleEntity, error) {
	var list []model.UserRoleEntity
	err := db.Instance().Engine().Table(d.TableName()).In(column, args).Find(&list)
	return list, err
}

//排除指定字段集合查询
func (d *uerRoleDao) FindNotIn(column string, args ...interface{}) ([]model.UserRoleEntity, error) {
	var list []model.UserRoleEntity
	err := db.Instance().Engine().Table(d.TableName()).NotIn(column, args).Find(&list)
	return list, err
}

// Fill with you ideas below.
//批量删除
func (d *uerRoleDao) DeleteBatch(ids ...int64) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).In("user_id", ids).Delete(new(model.UserRoleEntity))
}
