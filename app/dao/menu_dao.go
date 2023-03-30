package dao

import (
	"errors"
	"rygo/app/db"
	"rygo/app/model"
	"rygo/app/utils/page"
)

var MenuDao = newMenuDao()

func newMenuDao() *menuDao {
	return &menuDao{}
}

type menuDao struct {
}

//映射数据表
func (d *menuDao) TableName() string {
	return "sys_menu"
}

// 插入数据
func (d *menuDao) Insert(r *model.MenuEntity) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).Insert(r)
}

// 更新数据
func (d *menuDao) Update(r *model.MenuEntity) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).ID(r.MenuId).Update(r)
}

// 删除
func (d *menuDao) Delete(r *model.MenuEntity) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).ID(r.MenuId).Delete(r)
}

//批量删除
func (d *menuDao) DeleteBatch(ids ...int64) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).In("menu_id", ids).Delete(new(model.MenuEntity))
}

// 根据结构体中已有的非空数据来获得单条数据
func (d *menuDao) FindOne(r *model.MenuEntity) (bool, error) {
	return db.Instance().Engine().Table(d.TableName()).Get(r)
}

// 根据条件查询
func (d *menuDao) Find(where, order string) ([]model.MenuEntity, error) {
	var list []model.MenuEntity
	err := db.Instance().Engine().Table(d.TableName()).Where(where).OrderBy(order).Find(&list)
	return list, err
}

//指定字段集合查询
func (d *menuDao) FindIn(column string, args ...interface{}) ([]model.MenuEntity, error) {
	var list []model.MenuEntity
	err := db.Instance().Engine().Table(d.TableName()).In(column, args).Find(&list)
	return list, err
}

//排除指定字段集合查询
func (d *menuDao) FindNotIn(column string, args ...interface{}) ([]model.MenuEntity, error) {
	var list []model.MenuEntity
	err := db.Instance().Engine().Table(d.TableName()).NotIn(column, args).Find(&list)
	return list, err
}

//根据主键查询数据
func (d *menuDao) SelectRecordById(id int64) (*model.MenuEntityExtend, error) {
	db := db.Instance().Engine()

	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}
	var result model.MenuEntityExtend
	session := db.Table(d.TableName()).Alias("t")
	session.Select("t.menu_id, t.parent_id, t.menu_name, t.order_num, t.url, t.target, t.menu_type, t.visible, t.perms, t.icon, t.remark,(SELECT menu_name FROM sys_menu WHERE menu_id = t.parent_id) parent_name")
	session.Where("menu_id=?", id)
	_, err := session.Get(&result)
	if err != nil {
		return nil, errors.New("获取数据失败")
	}
	return &result, nil
}

//根据条件分页查询数据
func (d *menuDao) SelectListPage(param *model.MenuSelectPageReq) (*[]model.MenuEntity, *page.Paging, error) {
	db := db.Instance().Engine()
	p := new(page.Paging)
	if db == nil {
		return nil, p, errors.New("获取数据库连接失败")
	}

	session := db.Table(d.TableName()).Alias("t")

	if param != nil {
		if param.MenuName != "" {
			session.Where("m.menu_name like ?", "%"+param.MenuName+"%")
		}

		if param.Visible != "" {
			session.Where("m.visible = ", param.Visible)
		}

		if param.BeginTime != "" {
			session.Where("date_format(m.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			session.Where("date_format(m.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	tm := session.Clone()

	total, err := tm.Count()

	if err != nil {
		return nil, p, errors.New("读取行数失败")
	}

	p = page.CreatePaging(param.PageNum, param.PageSize, int(total))

	session.Limit(p.Pagesize, p.StartNum)

	var result []model.MenuEntity

	err = session.Find(&result)

	if err != nil {
		return nil, p, errors.New("读取数据失败")
	} else {
		return &result, p, nil
	}
}

//获取所有数据
func (d *menuDao) SelectListAll(param *model.MenuSelectPageReq) ([]model.MenuEntity, error) {
	db := db.Instance().Engine()

	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}

	session := db.Table(d.TableName()).Alias("t")

	if param != nil {

		if param.MenuName != "" {
			session.Where("t.menu_name like ?", "%"+param.MenuName+"%")
		}

		if param.Visible != "" {
			session.Where("t.visible = ?", param.Visible)
		}

		if param.BeginTime != "" {
			session.Where("date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			session.Where("date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}
	session.OrderBy("t.parent_id,t.order_num")
	var result []model.MenuEntity

	err := session.Find(&result)

	if err != nil {
		return nil, errors.New("读取数据失败")
	} else {
		return result, nil
	}
}

// 获取管理员菜单数据
func (d *menuDao) SelectMenuNormalAll() ([]model.MenuEntityExtend, error) {
	var result []model.MenuEntityExtend

	db := db.Instance().Engine()
	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}

	session := db.Table(d.TableName()).Alias("m")
	session.Select("distinct m.menu_id, m.parent_id, m.menu_name, m.url, m.visible, ifnull(m.perms,'') as perms, m.target, m.menu_type, m.icon, m.order_num, m.create_time")
	session.Where(" m.visible = 0")
	session.OrderBy("m.parent_id, m.order_num ")
	err := session.Find(&result)

	if err != nil {
		return nil, errors.New("读取数据失败")
	} else {
		return result, nil
	}
}

//根据用户ID读取菜单数据
func (d *menuDao) SelectMenusByUserId(userId string) ([]model.MenuEntityExtend, error) {
	var result []model.MenuEntityExtend

	db := db.Instance().Engine()
	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}

	session := db.Table(d.TableName()).Alias("m")
	session.Join("LEFT", []string{"sys_role_menu", "rm"}, "m.menu_id = rm.menu_id")
	session.Join("LEFT", []string{"sys_user_role", "ur"}, "rm.role_id = ur.role_id")
	session.Join("LEFT", []string{"sys_role", "ro"}, "ur.role_id = ro.role_id")
	session.Select("distinct m.menu_id, m.parent_id, m.menu_name, m.url, m.visible, ifnull(m.perms,'') as perms, m.target, m.menu_type, m.icon, m.order_num, m.create_time")
	session.Where("ur.user_id = ? and  m.visible = 0  AND ro.status = 0", userId)
	session.OrderBy("m.parent_id, m.order_num ")
	err := session.Find(&result)

	if err != nil {
		return nil, errors.New("读取数据失败")
	} else {
		return result, nil
	}
}

//根据角色ID查询菜单
func (d *menuDao) SelectMenuTree(roleId int64) ([]string, error) {
	db := db.Instance().Engine()

	var result []string

	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}
	session := db.Table(d.TableName()).Alias("m")
	session.Join("LEFT", []string{"sys_role_menu", "rm"}, "m.menu_id = rm.menu_id")
	session.Where("rm.role_id = ?", roleId)
	session.OrderBy("m.parent_id, m.order_num ")
	session.Select("concat(m.menu_id, ifnull(m.perms,'')) as perms")
	var list []model.MenuEntity
	err := session.Find(&list)
	if err != nil {
		return nil, errors.New("读取数据失败")
	}

	for _, record := range list {
		if record.Perms != "" {
			result = append(result, record.Perms)
		}
	}

	return result, nil
}

//校验菜单名称是否唯一
func (d *menuDao) CheckMenuNameUniqueAll(menuName string, parentId int64) (*model.MenuEntity, error) {
	var entity model.MenuEntity
	entity.MenuName = menuName
	entity.ParentId = parentId
	ok, err := d.FindOne(&entity)
	if ok {
		return &entity, err
	} else {
		return nil, err
	}
}

//校验菜单名称是否唯一
func (d *menuDao) CheckPermsUniqueAll(perms string) (*model.MenuEntity, error) {
	var entity model.MenuEntity
	entity.Perms = perms
	ok, err := d.FindOne(&entity)
	if ok {
		return &entity, err
	} else {
		return nil, err
	}
}
