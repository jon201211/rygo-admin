package dao

import (
	"errors"
	"rygo/app/db"
	"rygo/app/model"
	"rygo/app/utils/excel"
	"rygo/app/utils/page"

	"xorm.io/builder"
)

var UserOnlineDao = newUserOnlineDao()

func newUserOnlineDao() *userOnlineDao {
	return &userOnlineDao{}
}

type userOnlineDao struct {
}

//映射数据表
func (d *userOnlineDao) TableName() string {
	return "sys_user_online"
}

// 插入数据
func (d *userOnlineDao) Insert(r *model.UserOnline) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).Insert(r)
}

// 更新数据
func (d *userOnlineDao) Update(r *model.UserOnline) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).ID(r.Sessionid).Update(r)
}

// 删除
func (d *userOnlineDao) Delete(r *model.UserOnline) (int64, error) {
	rs, err := db.Instance().Engine().Exec("delete from sys_user_online where sessionId = ?", r.Sessionid)
	if err != nil {
		return 0, err
	}
	return rs.RowsAffected()
}

//批量删除
func (d *userOnlineDao) DeleteBatch(ids ...string) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).In("sessionId", ids).Delete(new(model.UserOnline))
}

// 根据结构体中已有的非空数据来获得单条数据
func (d *userOnlineDao) FindOne(r *model.UserOnline) (bool, error) {
	return db.Instance().Engine().Table(d.TableName()).Get(r)
}

// 根据条件查询
func (d *userOnlineDao) Find(where, order string) ([]model.UserOnline, error) {
	var list []model.UserOnline
	err := db.Instance().Engine().Table(d.TableName()).Where(where).OrderBy(order).Find(&list)
	return list, err
}

//指定字段集合查询
func (d *userOnlineDao) FindIn(column string, args ...interface{}) ([]model.UserOnline, error) {
	var list []model.UserOnline
	err := db.Instance().Engine().Table(d.TableName()).In(column, args).Find(&list)
	return list, err
}

//排除指定字段集合查询
func (d *userOnlineDao) FindNotIn(column string, args ...interface{}) ([]model.UserOnline, error) {
	var list []model.UserOnline
	err := db.Instance().Engine().Table(d.TableName()).NotIn(column, args).Find(&list)
	return list, err
}

//根据条件分页查询数据
func (d *userOnlineDao) SelectListByPage(param *model.UserOnlineSelectPageReq) ([]model.UserOnline, *page.Paging, error) {
	db := db.Instance().Engine()
	p := new(page.Paging)
	if db == nil {
		return nil, p, errors.New("获取数据库连接失败")
	}

	session := db.Table(d.TableName()).Alias("t")

	if param != nil {

		if param.SessionId != "" {
			session.Where("t.sessionId = ?", param.SessionId)
		}

		if param.LoginName != "" {
			session.Where("t.login_name like ?", "%"+param.LoginName+"%")
		}

		if param.DeptName != "" {
			session.Where("t.dept_name like ?", "%"+param.DeptName+"%")
		}

		if param.Ipaddr != "" {
			session.Where("t.ipaddr = ?", param.Ipaddr)
		}

		if param.LoginLocation != "" {
			session.Where("t.login_location = ?", param.LoginLocation)
		}

		if param.Browser != "" {
			session.Where("t.browser = ?", param.Browser)
		}

		if param.Os != "" {
			session.Where("t.os = ?", param.Os)
		}

		if param.Status != "" {
			session.Where("t.status = ?", param.Status)
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

	if param.OrderByColumn != "" {
		session.OrderBy(param.OrderByColumn + " " + param.IsAsc + " ")
	}

	session.Limit(p.Pagesize, p.StartNum)

	var result []model.UserOnline
	err = session.Find(&result)

	return result, p, err
}

// 导出excel
func (d *userOnlineDao) SelectListExport(param *model.UserOnlineSelectPageReq, head, col []string) (string, error) {
	db := db.Instance().Engine()

	if db == nil {
		return "", errors.New("获取数据库连接失败")
	}

	build := builder.Select(col...).From(d.TableName(), "t")
	if param != nil {

		if param.SessionId != "" {
			build.Where(builder.Eq{"t.sessionId": param.SessionId})
		}

		if param.LoginName != "" {
			build.Where(builder.Like{"t.login_name", param.LoginName})
		}

		if param.DeptName != "" {
			build.Where(builder.Like{"t.dept_name", param.DeptName})
		}

		if param.Ipaddr != "" {
			build.Where(builder.Eq{"t.ipaddr": param.Ipaddr})
		}

		if param.LoginLocation != "" {
			build.Where(builder.Eq{"t.login_location": param.LoginLocation})
		}

		if param.Browser != "" {
			build.Where(builder.Eq{"t.browser": param.Browser})
		}

		if param.Os != "" {
			build.Where(builder.Eq{"t.os": param.Os})
		}

		if param.Status != "" {
			build.Where(builder.Eq{"t.status": param.Status})
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
func (d *userOnlineDao) SelectListAll(param *model.UserOnlineSelectPageReq) ([]model.UserOnline, error) {
	db := db.Instance().Engine()

	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}

	session := db.Table(d.TableName()).Alias("t")

	if param != nil {

		if param.SessionId != "" {
			session.Where("t.sessionId = ?", param.SessionId)
		}

		if param.LoginName != "" {
			session.Where("t.login_name like ?", "%"+param.LoginName+"%")
		}

		if param.DeptName != "" {
			session.Where("t.dept_name like ?", "%"+param.DeptName+"%")
		}

		if param.Ipaddr != "" {
			session.Where("t.ipaddr = ?", param.Ipaddr)
		}

		if param.LoginLocation != "" {
			session.Where("t.login_location = ?", param.LoginLocation)
		}

		if param.Browser != "" {
			session.Where("t.browser = ?", param.Browser)
		}

		if param.Os != "" {
			session.Where("t.os = ?", param.Os)
		}

		if param.Status != "" {
			session.Where("t.status = ?", param.Status)
		}

		if param.BeginTime != "" {
			session.Where("date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			session.Where("date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	var result []model.UserOnline
	err := session.Find(&result)
	return result, err
}

//批量删除除参数以外的数据
func (d *userOnlineDao) DeleteNotIn(ids ...string) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).NotIn("sessionId", ids).Delete(new(model.UserOnline))
}
