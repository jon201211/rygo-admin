package dao

import (
	"errors"
	"rygo/app/db"
	"rygo/app/model"
	"rygo/app/utils/excel"
	"rygo/app/utils/page"

	"xorm.io/builder"
)

var RoleDao = newRoleDao()

func newRoleDao() *roleDao {
	return &roleDao{}
}

type roleDao struct {
}

//映射数据表
func (d *roleDao) TableName() string {
	return "sys_role"
}

// 插入数据
func (d *roleDao) Insert(r *model.RoleEntity) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).Insert(r)
}

// 更新数据
func (d *roleDao) Update(r *model.RoleEntity) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).ID(r.RoleId).Update(r)
}

// 删除
func (d *roleDao) Delete(r *model.RoleEntity) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).ID(r.RoleId).Delete(r)
}

//批量删除
func (d *roleDao) DeleteBatch(ids ...int64) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).In("role_id", ids).Delete(new(model.RoleEntity))
}

// 根据结构体中已有的非空数据来获得单条数据
func (d *roleDao) FindOne(r *model.RoleEntity) (bool, error) {
	return db.Instance().Engine().Table(d.TableName()).Get(r)
}

// 根据条件查询
func (d *roleDao) Find(where, order string) ([]model.RoleEntity, error) {
	var list []model.RoleEntity
	err := db.Instance().Engine().Table(d.TableName()).Where(where).OrderBy(order).Find(&list)
	return list, err
}

//指定字段集合查询
func (d *roleDao) FindIn(column string, args ...interface{}) ([]model.RoleEntity, error) {
	var list []model.RoleEntity
	err := db.Instance().Engine().Table(d.TableName()).In(column, args).Find(&list)
	return list, err
}

//排除指定字段集合查询
func (d *roleDao) FindNotIn(column string, args ...interface{}) ([]model.RoleEntity, error) {
	var list []model.RoleEntity
	err := db.Instance().Engine().Table(d.TableName()).NotIn(column, args).Find(&list)
	return list, err
}

//根据条件分页查询角色数据
func (d *roleDao) SelectListPage(param *model.RoleSelectPageReq) ([]model.RoleEntity, *page.Paging, error) {
	db := db.Instance().Engine()
	p := new(page.Paging)
	if db == nil {
		return nil, p, errors.New("获取数据库连接失败")
	}

	session := db.Table(d.TableName()).Alias("r").Where("r.del_flag = '0'")

	if param.RoleName != "" {
		session.Where("r.role_name like ?", "%"+param.RoleName+"%")
	}

	if param.Status != "" {
		session.Where("r.status = ?", param.Status)
	}

	if param.RoleKey != "" {
		session.Where("r.role_key like ?", "%"+param.RoleKey+"%")
	}

	if param.DataScope != "" {
		session.Where("r.data_scope = ?", param.DataScope)
	}

	if param.BeginTime != "" {
		session.Where("date_format(r.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
	}

	if param.EndTime != "" {
		session.Where("date_format(r.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
	}

	tm := session.Clone()

	total, err := tm.Count()

	if err != nil {
		return nil, p, errors.New("读取行数失败")
	}

	p = page.CreatePaging(param.PageNum, param.PageSize, int(total))

	if param.OrderByColumn != "" {
		session.OrderBy(param.OrderByColumn + " " + param.IsAsc + " ")
	}

	session.Limit(p.Pagesize, p.StartNum)

	var result []model.RoleEntity

	err = session.Find(&result)
	return result, p, err
}

// 导出excel
func (d *roleDao) SelectListExport(param *model.RoleSelectPageReq, head, col []string) (string, error) {
	db := db.Instance().Engine()
	if db == nil {
		return "", errors.New("获取数据库连接失败")
	}

	build := builder.Select(col...).From(d.TableName(), "t")

	if param != nil {
		if param.RoleName != "" {
			build.Where(builder.Like{"t.role_name", param.RoleName})
		}

		if param.Status != "" {
			build.Where(builder.Eq{"t.status": param.Status})
		}

		if param.RoleKey != "" {
			build.Where(builder.Like{"t.role_key", param.RoleKey})
		}

		if param.DataScope != "" {
			build.Where(builder.Eq{"t.data_scope": param.DataScope})
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

//获取所有角色数据
func (d *roleDao) SelectListAll(param *model.RoleSelectPageReq) ([]model.RoleEntityFlag, error) {
	db := db.Instance().Engine()

	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}

	session := db.Table(d.TableName()).Alias("r").Select("r.*,false as flag").Where("r.del_flag = '0'")
	if param != nil {
		if param.RoleName != "" {
			session.Where("r.role_name like ?", "%"+param.RoleName+"%")
		}

		if param.Status != "" {
			session.Where("r.status = ", param.Status)
		}

		if param.RoleKey != "" {
			session.Where("r.role_key like ?", "%"+param.RoleKey+"%")
		}

		if param.DataScope != "" {
			session.Where("r.data_scope = ", param.DataScope)
		}

		if param.BeginTime != "" {
			session.Where("date_format(r.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			session.Where("date_format(r.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	var result []model.RoleEntityFlag

	err := session.Find(&result)
	return result, err
}

//根据用户ID查询角色
func (d *roleDao) SelectRoleContactVo(userId int64) ([]model.RoleEntity, error) {
	db := db.Instance().Engine()

	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}

	session := db.Table(d.TableName()).Alias("r")
	session.Join("LEFT", []string{"sys_user_role", "ur"}, "ur.role_id = r.role_id")
	session.Join("LEFT", []string{"sys_user", "u"}, "u.user_id = ur.user_id")
	session.Join("LEFT", []string{"sys_dept", "d"}, "u.dept_id = d.dept_id")
	session.Where("r.del_flag = '0'")
	session.Where("ur.user_id = ?", userId)
	session.Select("distinct r.role_id, r.role_name, r.role_key, r.role_sort, r.data_scope,r.status, r.del_flag, r.create_time, r.remark")

	var result []model.RoleEntity

	err := session.Find(&result)
	return result, err
}

//检查角色键是否唯一
func (d *roleDao) CheckRoleNameUniqueAll(roleName string) (*model.RoleEntity, error) {
	var entity model.RoleEntity
	entity.RoleName = roleName
	_, err := d.FindOne(&entity)
	ok, err := d.FindOne(&entity)
	if ok {
		return &entity, err
	} else {
		return nil, err
	}
}

//检查角色键是否唯一
func (d *roleDao) CheckRoleKeyUniqueAll(roleKey string) (*model.RoleEntity, error) {
	var entity model.RoleEntity
	entity.RoleKey = roleKey
	ok, err := d.FindOne(&entity)
	if ok {
		return &entity, err
	} else {
		return nil, err
	}
}
