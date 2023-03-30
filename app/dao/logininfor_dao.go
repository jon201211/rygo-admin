package dao

import (
	"errors"
	"rygo/app/db"
	"rygo/app/model"
	"rygo/app/utils/excel"
	"rygo/app/utils/page"

	"xorm.io/builder"
)

var LogininforDao = newLoginInfoDao()

func newLoginInfoDao() *logininforDao {
	return &logininforDao{}
}

type logininforDao struct {
}

//映射数据表
func (d *logininforDao) TableName() string {
	return "sys_logininfor"
}

// 插入数据
func (d *logininforDao) Insert(r *model.LogininforEntity) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).Insert(r)
}

// 更新数据
func (d *logininforDao) Update(r *model.LogininforEntity) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).ID(r.InfoId).Update(r)
}

// 删除
func (d *logininforDao) Delete(r *model.LogininforEntity) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).ID(r.InfoId).Delete(r)
}

//批量删除
func (d *logininforDao) DeleteBatch(ids ...int64) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).In("info_id", ids).Delete(new(model.LogininforEntity))
}

// 根据结构体中已有的非空数据来获得单条数据
func (d *logininforDao) FindOne(r *model.LogininforEntity) (bool, error) {
	return db.Instance().Engine().Table(d.TableName()).Get(r)
}

// 根据条件查询
func (d *logininforDao) Find(where, order string) ([]model.LogininforEntity, error) {
	var list []model.LogininforEntity
	err := db.Instance().Engine().Table(d.TableName()).Where(where).OrderBy(order).Find(&list)
	return list, err
}

//指定字段集合查询
func (d *logininforDao) FindIn(column string, args ...interface{}) ([]model.LogininforEntity, error) {
	var list []model.LogininforEntity
	err := db.Instance().Engine().Table(d.TableName()).In(column, args).Find(&list)
	return list, err
}

//排除指定字段集合查询
func (d *logininforDao) FindNotIn(column string, args ...interface{}) ([]model.LogininforEntity, error) {
	var list []model.LogininforEntity
	err := db.Instance().Engine().Table(d.TableName()).NotIn(column, args).Find(&list)
	return list, err
}

// 根据条件分页查询用户列表
func (d *logininforDao) SelectPageList(param *model.LogininforSelectPageReq) (*[]model.LogininforEntity, *page.Paging, error) {
	db := db.Instance().Engine()
	p := new(page.Paging)
	if db == nil {
		return nil, p, errors.New("获取数据库连接失败")
	}

	session := db.Table(d.TableName())

	if param != nil {
		if param.LoginName != "" {
			session.Where("login_name like ?", "%"+param.LoginName+"%")
		}

		if param.Ipaddr != "" {
			session.Where("ipaddr like ?", "%"+param.Ipaddr+"%")
		}

		if param.Status != "" {
			session.Where("status = ?", param.Status)
		}

		if param.BeginTime != "" {
			session.Where("date_format(login_time,'%y%m%d') >= date_format(?,'%y%m%d')", param.BeginTime)
		}

		if param.EndTime != "" {
			session.Where("date_format(login_time,'%y%m%d') <= date_format(?,'%y%m%d')", param.EndTime)
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

	var result []model.LogininforEntity

	err = session.Find(&result)
	return &result, p, nil
}

// 导出excel
func (d *logininforDao) SelectExportList(param *model.LogininforSelectPageReq, head, col []string) (string, error) {
	db := db.Instance().Engine()
	if db == nil {
		return "", errors.New("获取数据库连接失败")
	}

	build := builder.Select(col...).From(d.TableName(), "t")

	if param != nil {
		if param.LoginName != "" {
			build.Where(builder.Like{"t.login_name", param.LoginName})
		}

		if param.Ipaddr != "" {
			build.Where(builder.Like{"t.ipaddr", param.Ipaddr})
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

//清空记录
func (d *logininforDao) DeleteAll() (int64, error) {
	db := db.Instance().Engine()
	if db == nil {
		return 0, errors.New("获取数据库连接失败")
	}

	rs, _ := db.Exec("delete from sys_logininfor")

	return rs.RowsAffected()
}
