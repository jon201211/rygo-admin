package dao

import (
	"errors"
	"rygo/app/db"
	"rygo/app/model"
	"rygo/app/utils/excel"
	"rygo/app/utils/page"

	"xorm.io/builder"
)

var UserDao = newUserDao()

func newUserDao() *userDao {
	return &userDao{}
}

type userDao struct {
}

//映射数据表
func (d *userDao) TableName() string {
	return "sys_user"
}

// 插入数据
func (d *userDao) Insert(r *model.SysUser) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).Insert(r)
}

// 更新数据
func (d *userDao) Update(r *model.SysUser) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).ID(r.UserId).Update(r)
}

// 删除
func (d *userDao) Delete(r *model.SysUser) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).ID(r.UserId).Delete(r)
}

//批量删除
func (d *userDao) DeleteBatch(ids ...int64) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).In("user_id", ids).Delete(new(model.SysUser))
}

// 根据结构体中已有的非空数据来获得单条数据
func (d *userDao) FindOne(r *model.SysUser) (bool, error) {
	return db.Instance().Engine().Table(d.TableName()).Get(r)
}

// 根据条件查询
func (d *userDao) Find(where, order string) ([]model.SysUser, error) {
	var list []model.SysUser
	err := db.Instance().Engine().Table(d.TableName()).Where(where).OrderBy(order).Find(&list)
	return list, err
}

//指定字段集合查询
func (d *userDao) FindIn(column string, args ...interface{}) ([]model.SysUser, error) {
	var list []model.SysUser
	err := db.Instance().Engine().Table(d.TableName()).In(column, args).Find(&list)
	return list, err
}

//排除指定字段集合查询
func (d *userDao) FindNotIn(column string, args ...interface{}) ([]model.SysUser, error) {
	var list []model.SysUser
	err := db.Instance().Engine().Table(d.TableName()).NotIn(column, args).Find(&list)
	return list, err
}

// 根据条件分页查询用户列表
func (d *userDao) SelectPageList(param *model.UserSelectPageReq) ([]model.UserListEntity, *page.Paging, error) {
	db := db.Instance().Engine()
	p := new(page.Paging)
	if db == nil {
		return nil, p, errors.New("获取数据库连接失败")
	}

	session := db.Table(d.TableName()).Alias("u").Join("LEFT", []string{"sys_dept", "d"}, "u.dept_id = d.dept_id")
	session.Where(" u.del_flag = '0' ")

	if param != nil {
		if param.LoginName != "" {
			session.Where("u.login_name like ?", "%"+param.LoginName+"%")
		}

		if param.Phonenumber != "" {
			session.Where("u.phonenumber like ?", "%"+param.Phonenumber+"%")
		}

		if param.Status != "" {
			session.Where("u.status = ?", param.Status)
		}
		if param.TenantId != 0 {
			session.Where("u.tenant_id = ?", param.TenantId)
		}
		if param.BeginTime != "" {
			session.Where("date_format(u.create_time,'%y%m%d') >= date_format(?,'%y%m%d')", param.BeginTime)
		}

		if param.EndTime != "" {
			session.Where("date_format(u.create_time,'%y%m%d') <= date_format(?,'%y%m%d')", param.EndTime)
		}

		if param.DeptId != 0 {
			session.Where("(u.dept_id = ? OR u.dept_id IN ( SELECT t.dept_id FROM sys_dept t WHERE FIND_IN_SET (?,ancestors) ))", param.DeptId, param.DeptId)
		}
	}

	tm := session.Clone()

	total, err := tm.Count()

	if err != nil {
		return nil, p, errors.New("读取行数失败")
	}

	p = page.CreatePaging(param.PageNum, param.PageSize, int(total))

	session.Select("u.user_id, u.dept_id, u.login_name, u.user_name, u.email, u.avatar, u.phonenumber, u.password,u.sex, u.salt, u.status, u.del_flag, u.login_ip, u.login_date, u.create_by, u.create_time, u.remark,d.dept_name, d.leader")

	session.OrderBy("u." + param.SortName + " " + param.SortOrder + " ")

	session.Limit(p.Pagesize, p.StartNum)

	var result []model.UserListEntity
	err = session.Find(&result)
	return result, p, err
}

// 导出excel
func (d *userDao) SelectExportList(param *model.UserSelectPageReq, head, col []string) (string, error) {
	db := db.Instance().Engine()
	if db == nil {
		return "", errors.New("获取数据库连接失败")
	}

	build := builder.Select(col...).From(d.TableName(), "u").LeftJoin("sys_dept d", "u.dept_id = d.dept_id").Where(builder.Expr("u.del_flag = '0'"))

	if param != nil {
		if param.LoginName != "" {
			build.Where(builder.Like{"u.login_name", param.LoginName})
		}

		if param.Phonenumber != "" {
			build.Where(builder.Like{"u.phonenumber", param.Phonenumber})
		}

		if param.Status != "" {
			build.Where(builder.Eq{"u.status": param.Status})
		}

		if param.BeginTime != "" {
			build.Where(builder.Gte{"date_format(u.create_time,'%y%m%d')": "date_format('" + param.BeginTime + "','%y%m%d')"})
		}

		if param.EndTime != "" {
			build.Where(builder.Lte{"date_format(u.create_time,'%y%m%d')": "date_format('" + param.EndTime + "','%y%m%d')"})
		}

		if param.DeptId != 0 {
			build.Where(builder.Eq{"u.dept_id": param.DeptId}.Or(builder.In("u.dept_id", builder.Expr("SELECT t.dept_id FROM sys_dept t WHERE FIND_IN_SET (?,ancestors) )", param.DeptId))))
		}
	}

	sqlStr, _, _ := build.ToSQL()
	arr, err := db.SQL(sqlStr).QuerySliceString()

	path, err := excel.DownlaodExcel(head, arr)

	return path, err
}

// 根据条件分页查询已分配用户角色列表
func (d *userDao) SelectAllocatedList(roleId int64, loginName, phonenumber string) ([]model.SysUser, error) {
	db := db.Instance().Engine()

	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}

	session := db.Table(d.TableName()).Alias("u")
	session.Join("LEFT", []string{"sys_dept", "d"}, "u.dept_id = d.dept_id")
	session.Join("LEFT", []string{"sys_user_role", "ur"}, " u.user_id = ur.user_id")
	session.Join("LEFT", []string{"sys_role", "r"}, "r.role_id = ur.role_id")

	session.Where("u.del_flag =?", 0)
	session.Where("r.role_id = ?", roleId)

	if loginName != "" {
		session.Where("u.login_name like ?", "%"+loginName+"%")
	}

	if phonenumber != "" {
		session.Where("u.phonenumber like ?", "%"+phonenumber+"%")
	}

	var result []model.SysUser

	session.Select("distinct u.user_id, u.dept_id, u.login_name, u.user_name, u.email, u.avatar, u.phonenumber,u.status, u.create_time")

	err := session.Find(&result)
	return result, err
}

// 根据条件分页查询未分配用户角色列表
func (d *userDao) SelectUnallocatedList(roleId int64, loginName, phonenumber string) ([]model.SysUser, error) {
	db := db.Instance().Engine()
	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}

	session := db.Table(d.TableName()).Alias("u")
	session.Join("LEFT", []string{"sys_dept", "d"}, "u.dept_id = d.dept_id")
	session.Join("LEFT", []string{"sys_user_role", "ur"}, "u.user_id = ur.user_id")
	session.Join("LEFT", []string{"sys_role", "r"}, "r.role_id = ur.role_id")

	session.Where("u.user_id not in (select u.user_id from sys_user u inner join sys_user_role ur on u.user_id = ur.user_id and ur.role_id = ?)", roleId)

	if loginName != "" {
		session.Where("u.login_name like ?", "%"+loginName+"%")
	}

	if phonenumber != "" {
		session.Where("u.phonenumber like ?", "%"+phonenumber+"%")
	}

	session.Select("distinct u.user_id, u.dept_id, u.login_name, u.user_name, u.email, u.avatar, u.phonenumber, u.status, u.create_time")

	var result []model.SysUser
	err := session.Find(&result)
	return result, err
}

//检查邮箱是否已使用
func (d *userDao) CheckEmailUnique(userId int64, email string) bool {
	db := db.Instance().Engine()
	if db == nil {
		return false
	}
	rs, err := db.Table(d.TableName()).Where("email=? AND user_id<>?", email, userId).Count()
	if err != nil {
		return false
	}

	if rs > 0 {
		return true
	} else {
		return false
	}
}

//检查邮箱是否存在,存在返回true,否则false
func (d *userDao) CheckEmailUniqueAll(email string) bool {
	db := db.Instance().Engine()
	if db == nil {
		return false
	}
	rs, err := db.Table(d.TableName()).Where("email=?", email).Count()
	if err != nil {
		return false
	}

	if rs > 0 {
		return true
	} else {
		return false
	}
}

//检查手机号是否已使用,存在返回true,否则false
func (d *userDao) CheckPhoneUnique(userId int64, phone string) bool {
	db := db.Instance().Engine()
	if db == nil {
		return false
	}
	rs, err := db.Table(d.TableName()).Where("phonenumber = ? AND user_id<>?", phone, userId).Count()
	if err != nil {
		return false
	}

	if rs > 0 {
		return true
	} else {
		return false
	}
}

//检查手机号是否已使用 ,存在返回true,否则false
func (d *userDao) CheckPhoneUniqueAll(phone string) bool {
	db := db.Instance().Engine()
	if db == nil {
		return false
	}
	rs, err := db.Table(d.TableName()).Where("phonenumber = ?", phone).Count()
	if err != nil {
		return false
	}

	if rs > 0 {
		return true
	} else {
		return false
	}
}

//根据登陆名查询用户信息
func (d *userDao) SelectUserByLoginName(loginName string) (*model.SysUser, error) {
	var entity model.SysUser
	entity.LoginName = loginName
	ok, err := d.FindOne(&entity)
	if ok {
		return &entity, err
	} else {
		return nil, err
	}
}

//根据手机号查询用户信息
func (d *userDao) SelectUserByPhoneNumber(phonenumber string) (*model.SysUser, error) {
	var entity model.SysUser
	entity.Phonenumber = phonenumber
	ok, err := d.FindOne(&entity)
	if ok {
		return &entity, err
	} else {
		return nil, err
	}
}
