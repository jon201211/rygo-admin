package dao

import (
	"errors"
	"rygo/app/db"
	"rygo/app/model"
	"rygo/app/utils/excel"
	"rygo/app/utils/page"

	"xorm.io/builder"
)

var OperLogDao = newOperLogDao()

func newOperLogDao() *operLogDao {
	return &operLogDao{}
}

type operLogDao struct {
}

//映射数据表
func (d *operLogDao) TableName() string {
	return "sys_oper_log"
}

// 插入数据
func (d *operLogDao) Insert(r *model.OperLogEntity) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).Insert(r)
}

// 更新数据
func (d *operLogDao) Update(r *model.OperLogEntity) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).ID(r.OperId).Update(r)
}

// 删除
func (d *operLogDao) Delete(r *model.OperLogEntity) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).ID(r.OperId).Delete(r)
}

//批量删除
func (d *operLogDao) DeleteBatch(ids ...int64) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).In("oper_id", ids).Delete(new(model.OperLogEntity))
}

// 根据结构体中已有的非空数据来获得单条数据
func (d *operLogDao) FindOne(r *model.OperLogEntity) (bool, error) {
	return db.Instance().Engine().Table(d.TableName()).Get(r)
}

// 根据条件查询
func (d *operLogDao) Find(where, order string) ([]model.OperLogEntity, error) {
	var list []model.OperLogEntity
	err := db.Instance().Engine().Table(d.TableName()).Where(where).OrderBy(order).Find(&list)
	return list, err
}

//指定字段集合查询
func (d *operLogDao) FindIn(column string, args ...interface{}) ([]model.OperLogEntity, error) {
	var list []model.OperLogEntity
	err := db.Instance().Engine().Table(d.TableName()).In(column, args).Find(&list)
	return list, err
}

//排除指定字段集合查询
func (d *operLogDao) FindNotIn(column string, args ...interface{}) ([]model.OperLogEntity, error) {
	var list []model.OperLogEntity
	err := db.Instance().Engine().Table(d.TableName()).NotIn(column, args).Find(&list)
	return list, err
}

// 根据条件分页查询用户列表
func (d *operLogDao) SelectPageList(param *model.OperLogSelectPageReq) (*[]model.OperLogEntity, *page.Paging, error) {
	db := db.Instance().Engine()
	p := new(page.Paging)
	if db == nil {
		return nil, p, errors.New("获取数据库连接失败")
	}

	session := db.Table(d.TableName())

	if param != nil {
		if param.Title != "" {
			session.Where("title like ?", "%"+param.Title+"%")
		}

		if param.OperName != "" {
			session.Where("oper_name like ?", "%"+param.OperName+"%")
		}

		if param.Status != "" {
			session.Where("status = ?", param.Status)
		}

		if param.BusinessTypes >= 0 {
			session.Where("status = ?", param.BusinessTypes)
		}

		if param.BeginTime != "" {
			session.Where("date_format(oper_time,'%y%m%d') >= date_format(?,'%y%m%d')", param.BeginTime)
		}

		if param.EndTime != "" {
			session.Where("date_format(oper_time,'%y%m%d') <= date_format(?,'%y%m%d')", param.EndTime)
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

	var result []model.OperLogEntity

	err = session.Find(&result)
	return &result, p, nil
}

// 导出excel
func (d *operLogDao) SelectExportList(param *model.OperLogSelectPageReq, head, col []string) (string, error) {
	db := db.Instance().Engine()
	if db == nil {
		return "", errors.New("获取数据库连接失败")
	}

	build := builder.Select(col...).From(d.TableName())

	if param != nil {
		if param.Title != "" {
			build.Where(builder.Like{"title", param.Title})
		}

		if param.OperName != "" {
			build.Where(builder.Like{"oper_name", param.OperName})
		}

		if param.Status != "" {
			build.Where(builder.Eq{"status": param.Status})
		}

		if param.BusinessTypes >= 0 {
			build.Where(builder.Eq{"business_type": param.BusinessTypes})
		}

		if param.BeginTime != "" {
			build.Where(builder.Gte{"date_format(create_time,'%y%m%d')": "date_format('" + param.BeginTime + "','%y%m%d')"})
		}

		if param.EndTime != "" {
			build.Where(builder.Lte{"date_format(create_time,'%y%m%d')": "date_format('" + param.EndTime + "','%y%m%d')"})
		}
	}

	sqlStr, _, _ := build.ToSQL()
	arr, err := db.SQL(sqlStr).QuerySliceString()

	path, err := excel.DownlaodExcel(head, arr)

	return path, err
}

//清空记录
func (d *operLogDao) DeleteAll() (int64, error) {
	db := db.Instance().Engine()
	if db == nil {
		return 0, errors.New("获取数据库连接失败")
	}

	rs, _ := db.Exec("delete from sys_oper_log")

	return rs.RowsAffected()
}
