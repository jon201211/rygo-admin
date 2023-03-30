package dao

import (
	"errors"
	"rygo/app/db"
	"rygo/app/model"
	"rygo/app/utils/excel"
	"rygo/app/utils/page"

	"xorm.io/builder"
)

var JobDao = newJobDao()

func newJobDao() *jobDao {
	return &jobDao{}
}

type jobDao struct {
}

//映射数据表
func (d *jobDao) TableName() string {
	return "sys_job"
}

// 插入数据
func (d *jobDao) Insert(r *model.SysJob) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).Insert(r)
}

// 更新数据
func (d *jobDao) Update(r *model.SysJob) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).ID(r.JobId).Update(r)
}

// 删除
func (d *jobDao) Delete(r *model.SysJob) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).ID(r.JobId).Delete(r)
}

//批量删除
func (d *jobDao) DeleteBatch(ids ...int64) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).In("job_id", ids).Delete(new(model.SysJob))
}

// 根据结构体中已有的非空数据来获得单条数据
func (d *jobDao) FindOne(r *model.SysJob) (bool, error) {
	return db.Instance().Engine().Table(d.TableName()).Get(r)
}

// 根据条件查询
func (d *jobDao) Find(where, order string) ([]model.SysJob, error) {
	var list []model.SysJob
	err := db.Instance().Engine().Table(d.TableName()).Where(where).OrderBy(order).Find(&list)
	return list, err
}

//指定字段集合查询
func (d *jobDao) FindIn(column string, args ...interface{}) ([]model.SysJob, error) {
	var list []model.SysJob
	err := db.Instance().Engine().Table(d.TableName()).In(column, args).Find(&list)
	return list, err
}

//排除指定字段集合查询
func (d *jobDao) FindNotIn(column string, args ...interface{}) ([]model.SysJob, error) {
	var list []model.SysJob
	err := db.Instance().Engine().Table(d.TableName()).NotIn(column, args).Find(&list)
	return list, err
}

//根据条件分页查询数据
func (d *jobDao) SelectListByPage(param *model.JobSelectPageReq) (*[]model.SysJob, *page.Paging, error) {
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

		if param.CronExpression != "" {
			session.Where("t.cron_expression = ?", param.CronExpression)
		}

		if param.MisfirePolicy != "" {
			session.Where("t.misfire_policy = ?", param.MisfirePolicy)
		}

		if param.Concurrent != "" {
			session.Where("t.concurrent = ?", param.Concurrent)
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

	var result []model.SysJob

	err = session.Find(&result)

	if err != nil {
		return nil, p, errors.New("读取数据失败")
	}
	return &result, p, nil
}

// 导出excel
func (d *jobDao) SelectListExport(param *model.JobSelectPageReq, head, col []string) (string, error) {
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

		if param.CronExpression != "" {
			build.Where(builder.Eq{"t.cron_expression": param.CronExpression})
		}

		if param.MisfirePolicy != "" {
			build.Where(builder.Eq{"t.misfire_policy": param.MisfirePolicy})
		}

		if param.Concurrent != "" {
			build.Where(builder.Eq{"t.concurrent": param.Concurrent})
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
func (d *jobDao) SelectListAll(param *model.JobSelectPageReq) ([]model.SysJob, error) {
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

		if param.CronExpression != "" {
			session.Where("t.cron_expression = ?", param.CronExpression)
		}

		if param.MisfirePolicy != "" {
			session.Where("t.misfire_policy = ?", param.MisfirePolicy)
		}

		if param.Concurrent != "" {
			session.Where("t.concurrent = ?", param.Concurrent)
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
	var result []model.SysJob
	err := session.Find(&result)

	if err != nil {
		return nil, errors.New("读取数据失败")
	}
	return result, nil
}

//批量修改状态
func (d *jobDao) UpdateState(ids, status string) (int64, error) {
	db := db.Instance().Engine()

	if db == nil {
		return 0, errors.New("获取数据库连接失败")
	}

	rs, err := db.Exec("update sys_job set status=? where job_id in (?)", status, ids)
	if err != nil {
		return 0, err
	}
	return rs.RowsAffected()
}
