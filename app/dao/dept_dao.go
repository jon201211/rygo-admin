package dao

import (
	"errors"
	"fmt"
	"rygo/app/db"
	"rygo/app/model"
	"rygo/app/utils/gconv"
	"strings"
)

var DeptDao = newDeptDao()

func newDeptDao() *deptDao {
	return &deptDao{}
}

type deptDao struct {
}

//映射数据表
func (d *deptDao) TableName() string {
	return "sys_dept"
}

// 插入数据
func (d *deptDao) Insert(r *model.SysDept) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).Insert(r)
}

// 更新数据
func (d *deptDao) Update(r *model.SysDept) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).ID(r.DeptId).Update(r)
}

// 删除
func (d *deptDao) Delete(r *model.SysDept) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).ID(r.DeptId).Delete(r)
}

// 根据结构体中已有的非空数据来获得单条数据
func (d *deptDao) FindOne(r *model.SysDept) (bool, error) {
	return db.Instance().Engine().Table(d.TableName()).Get(r)
}

//根据部门ID查询信息
func (d *deptDao) SelectDeptById(id int64) (*model.SysDeptExtend, error) {
	db := db.Instance().Engine()

	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}
	var result model.SysDeptExtend
	session := db.Table(d.TableName()).Alias("d")
	session.Select("d.dept_id, d.parent_id, d.ancestors, d.dept_name, d.order_num, d.leader, d.phone, d.email, d.status,(select dept_name from sys_dept where dept_id = d.parent_id) parent_name")
	session.Where("d.dept_id = ?", id)
	_, err := session.Get(&result)
	return &result, err
}

//根据ID查询所有子部门
func (d *deptDao) SelectChildrenDeptById(deptId int64) []*model.SysDept {
	db := db.Instance().Engine()

	if db == nil {
		return nil
	}
	var rs []*model.SysDept
	db.Table(d.TableName()).Where("find_in_set(?, ancestors)", deptId).Find(&rs)
	return rs
}

//删除部门管理信息
func (d *deptDao) DeleteDeptById(deptId int64) int64 {
	var entity model.SysDept
	entity.DeptId = deptId
	entity.DelFlag = "2"
	rs, err := d.Update(&entity)
	if err != nil {
		return 0
	}
	return rs
}

//修改子元素关系
func (d *deptDao) UpdateDeptChildren(deptId int64, newAncestors, oldAncestors string) {
	deptList := d.SelectChildrenDeptById(deptId)

	if deptList == nil || len(deptList) <= 0 {
		return
	}

	for _, tmp := range deptList {
		tmp.Ancestors = strings.ReplaceAll(tmp.Ancestors, oldAncestors, newAncestors)
	}

	ancestors := " case dept_id"
	idStr := ""

	for _, dept := range deptList {
		ancestors += " when " + gconv.String(dept.DeptId) + " then " + dept.Ancestors
		if idStr == "" {
			idStr = gconv.String(dept.DeptId)
		} else {
			idStr += "," + gconv.String(dept.DeptId)
		}
	}

	ancestors += " end "
	db := db.Instance().Engine()

	if db == nil {
		return
	}

	rs, err := db.Table(d.TableName()).Where("dept_id in(?)", deptId).Update(map[string]interface{}{"ancestors": ancestors})
	fmt.Printf("修改了%v行 错误信息：%v", rs, err.Error())
}

//查询部门管理数据
func (d *deptDao) SelectDeptList(parentId int64, deptName, status string, tenantId int64) ([]model.SysDept, error) {
	var result []model.SysDept
	db := db.Instance().Engine()
	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}
	session := db.Table(d.TableName()).Alias("d").Where("d.del_flag = '0'")
	if parentId > 0 {
		session.Where("d.parent_id = ?", parentId)
	}
	if deptName != "" {
		session.Where("d.dept_name like ?", "%"+deptName+"%")
	}
	if status != "" {
		session.Where("d.status = ?", status)
	}
	if tenantId != 0 {
		session.Where("d.tenant_id = ?", tenantId)
	}

	session.OrderBy("d.parent_id, d.order_num")

	err := session.Find(&result)

	return result, err
}

//根据角色ID查询部门
func (d *deptDao) SelectRoleDeptTree(roleId int64) ([]string, error) {
	db := db.Instance().Engine()
	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}
	session := db.Table(d.TableName()).Alias("d").Join("LEFT", []string{"sys_role_dept", "rd"}, "d.dept_id = rd.dept_id")
	session.Where("d.del_flag = '0'").Where("rd.role_id = ?", roleId)
	session.OrderBy("d.parent_id, d.order_num ")
	session.Select("concat(d.dept_id, d.dept_name) as dept_name")

	var result []string
	var rs []model.SysDept
	err := session.Find(&result)
	if err == nil && rs != nil && len(rs) > 0 {
		for _, record := range rs {
			if record.DeptName != "" {
				result = append(result, record.DeptName)
			}
		}
	}
	return result, nil
}

//查询部门是否存在用户
func (d *deptDao) CheckDeptExistUser(deptId int64) bool {
	db := db.Instance().Engine()
	if db == nil {
		return false
	}

	num, _ := db.Table(d.TableName()).Where("dept_id = ? and del_flag = '0'", deptId).Count()

	if num > 0 {
		return true
	} else {
		return false
	}
}

//查询部门人数
func (d *deptDao) SelectDeptCount(deptId, parentId int64) int64 {
	db := db.Instance().Engine()
	if db == nil {
		return 0
	}

	result := int64(0)
	whereStr := "del_flag = '0'"
	if deptId > 0 {
		whereStr = whereStr + " and dept_id=" + gconv.String(deptId)
	}
	if parentId > 0 {
		whereStr = whereStr + " and parent_id=" + gconv.String(parentId)
	}

	rs, err := db.Table(d.TableName()).Where(whereStr).Count()
	if err != nil {
		result = rs
	}
	return result
}

//校验部门名称是否唯一
func (d *deptDao) CheckDeptNameUniqueAll(deptName string, parentId int64) (*model.SysDept, error) {
	var entity model.SysDept
	entity.DeptName = deptName
	entity.ParentId = parentId
	ok, err := d.FindOne(&entity)
	if ok {
		return &entity, err
	} else {
		return nil, err
	}
}

// 根据条件查询
func (d *deptDao) Find(where, order string) ([]model.SysDept, error) {
	var list []model.SysDept
	err := db.Instance().Engine().Table(d.TableName()).Where(where).OrderBy(order).Find(&list)
	return list, err
}

//指定字段集合查询
func (d *deptDao) FindIn(column string, args ...interface{}) ([]model.SysDept, error) {
	var list []model.SysDept
	err := db.Instance().Engine().Table(d.TableName()).In(column, args).Find(&list)
	return list, err
}

//排除指定字段集合查询
func (d *deptDao) FindNotIn(column string, args ...interface{}) ([]model.SysDept, error) {
	var list []model.SysDept
	err := db.Instance().Engine().Table(d.TableName()).NotIn(column, args).Find(&list)
	return list, err
}

//批量删除
func (d *deptDao) DeleteBatch(ids ...int64) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).In("dept_id", ids).Delete(new(model.SysDept))
}
