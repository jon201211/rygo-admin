package dao

import (
	"rygo/app/db"
	"rygo/app/model"

	"xorm.io/core"
)

var UserPostDao = newUserPostDao()

func newUserPostDao() *uerPostDao {
	return &uerPostDao{}
}

type uerPostDao struct {
}

//映射数据表
func (d *uerPostDao) TableName() string {
	return "sys_user_post"
}

// 插入数据
func (d *uerPostDao) Insert(r *model.UserPostEntity) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).Insert(r)
}

// 更新数据
func (d *uerPostDao) Update(r *model.UserPostEntity) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).ID(core.PK{r.UserId, r.PostId}).Update(r)
}

// 删除
func (d *uerPostDao) Delete(r *model.UserPostEntity) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).ID(core.PK{r.UserId, r.PostId}).Delete(r)
}

// 根据结构体中已有的非空数据来获得单条数据
func (d *uerPostDao) FindOne(r *model.UserPostEntity) (bool, error) {
	return db.Instance().Engine().Table(d.TableName()).Get(r)
}

// 根据条件查询
func (d *uerPostDao) Find(where, order string) ([]model.UserPostEntity, error) {
	var list []model.UserPostEntity
	err := db.Instance().Engine().Table(d.TableName()).Where(where).OrderBy(order).Find(&list)
	return list, err
}

//指定字段集合查询
func (d *uerPostDao) FindIn(column string, args ...interface{}) ([]model.UserPostEntity, error) {
	var list []model.UserPostEntity
	err := db.Instance().Engine().Table(d.TableName()).In(column, args).Find(&list)
	return list, err
}

//排除指定字段集合查询
func (d *uerPostDao) FindNotIn(column string, args ...interface{}) ([]model.UserPostEntity, error) {
	var list []model.UserPostEntity
	err := db.Instance().Engine().Table(d.TableName()).NotIn(column, args).Find(&list)
	return list, err
}

// Fill with you ideas below.

//批量删除
func (d *uerPostDao) DeleteBatch(ids ...int64) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).In("user_id", ids).Delete(new(model.UserPostEntity))
}
