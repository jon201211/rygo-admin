package dao

import (
	"errors"
	"rygo/app/db"
	"rygo/app/model"
	"rygo/app/utils/excel"
	"rygo/app/utils/page"

	"xorm.io/builder"
)

var JobLogDao = newJobLogDao()

func newJobLogDao() *jobLogDao {
	return &jobLogDao{}
}

type jobLogDao struct {
}

//映射数据表
func (d *jobLogDao) TableName() string {
	return "sys_job_log"
}

// 插入数据
func (d *jobLogDao) Insert(r *model.JobLogEntity) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).Insert(r)
}

// 更新数据
func (d *jobLogDao) Update(r *model.JobLogEntity) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).ID(r.JobLogId).Update(r)
}

// 删除
func (d *jobLogDao) Delete(r *model.JobLogEntity) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).ID(r.JobLogId).Delete(r)
}

//批量删除
func (d *jobLogDao) DeleteBatch(ids ...int64) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).In("job_log_id", ids).Delete(new(model.JobLogEntity))
}

// 根据结构体中已有的非空数据来获得单条数据
func (d *jobLogDao) FindOne(r *model.JobLogEntity) (bool, error) {
	return db.Instance().Engine().Table(d.TableName()).Get(r)
}

// 根据条件查询
func (d *jobLogDao) Find(where, order string) ([]model.JobLogEntity, error) {
	var list []model.JobLogEntity
	err := db.Instance().Engine().Table(d.TableName()).Where(where).OrderBy(order).Find(&list)
	return list, err
}

//指定字段集合查询
func (d *jobLogDao) FindIn(column string, args ...interface{}) ([]model.JobLogEntity, error) {
	var list []model.JobLogEntity
	err := db.Instance().Engine().Table(d.TableName()).In(column, args).Find(&list)
	return list, err
}

//排除指定字段集合查询
func (d *jobLogDao) FindNotIn(column string, args ...interface{}) ([]model.JobLogEntity, error) {
	var list []model.JobLogEntity
	err := db.Instance().Engine().Table(d.TableName()).NotIn(column, args).Find(&list)
	return list, err
}

//根据条件分页查询数据
func (d *jobLogDao) SelectListByPage(param *model.JobLogSelectPageReq) (*[]model.JobLogEntity, *page.Paging, error) {
	db := db.Instance().Engine()
	p := new(page.Paging)
	if db == nil {
		return nil, p, errors.New("获取数据库连接失败")
	}

	session := db.Table(d.TableName()).Alias("t")

	if param != nil {

		if param.JobName != "" {
			session.Where("t.job_name like ?", "%"+param.JobName+"%")
		}

		if param.JobGroup != "" {
			session.Where("t.job_group = ?", param.JobGroup)
		}

		if param.InvokeTarget != "" {
			session.Where("t.invoke_target = ?", param.InvokeTarget)
		}

		if param.JobMessage != "" {
			session.Where("t.job_message = ?", param.JobMessage)
		}

		if param.Status != "" {
			session.Where("t.status = ?", param.Status)
		}

		if param.ExceptionInfo != "" {
			session.Where("t.exception_info = ?", param.ExceptionInfo)
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

	session.Limit(p.Pagesize, p.StartNum)
	var result []model.JobLogEntity

	err = session.Find(&result)

	if err != nil {
		return nil, p, errors.New("读取数据失败")
	}
	return &result, p, nil
}

// 导出excel
func (d *jobLogDao) SelectListExport(param *model.JobLogSelectPageReq, head, col []string) (string, error) {
	db := db.Instance().Engine()

	if db == nil {
		return "", errors.New("获取数据库连接失败")
	}

	build := builder.Select(col...).From(d.TableName(), "t")

	if param != nil {

		if param.JobName != "" {
			build.Where(builder.Like{"t.job_name", param.JobName})
		}

		if param.JobGroup != "" {
			build.Where(builder.Eq{"t.job_group": param.JobGroup})
		}

		if param.InvokeTarget != "" {
			build.Where(builder.Eq{"t.invoke_target": param.InvokeTarget})
		}

		if param.JobMessage != "" {
			build.Where(builder.Eq{"t.job_message": param.JobMessage})
		}

		if param.Status != "" {
			build.Where(builder.Eq{"t.status": param.Status})
		}

		if param.ExceptionInfo != "" {
			build.Where(builder.Eq{"t.exception_info": param.ExceptionInfo})
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
func (d *jobLogDao) SelectListAll(param *model.JobLogSelectPageReq) ([]model.JobLogEntity, error) {
	db := db.Instance().Engine()

	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}

	session := db.Table(d.TableName()).Alias("t")

	if param != nil {

		if param.JobName != "" {
			session.Where("t.job_name like ?", "%"+param.JobName+"%")
		}

		if param.JobGroup != "" {
			session.Where("t.job_group = ?", param.JobGroup)
		}

		if param.InvokeTarget != "" {
			session.Where("t.invoke_target = ?", param.InvokeTarget)
		}

		if param.JobMessage != "" {
			session.Where("t.job_message = ?", param.JobMessage)
		}

		if param.Status != "" {
			session.Where("t.status = ?", param.Status)
		}

		if param.ExceptionInfo != "" {
			session.Where("t.exception_info = ?", param.ExceptionInfo)
		}

		if param.BeginTime != "" {
			session.Where("date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			session.Where("date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}
	var result []model.JobLogEntity
	err := session.Find(&result)

	if err != nil {
		return nil, errors.New("读取数据失败")
	}
	return result, nil
}
