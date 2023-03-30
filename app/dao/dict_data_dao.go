package dao

import (
	"errors"
	"rygo/app/db"
	"rygo/app/model"
	"rygo/app/utils/excel"
	"rygo/app/utils/page"

	"xorm.io/builder"
)

var DictDataDao = newDictDataDao()

func newDictDataDao() *dictDataDao {
	return &dictDataDao{}
}

type dictDataDao struct {
}

//映射数据表
func (d *dictDataDao) TableName() string {
	return "sys_dict_data"
}

// 插入数据
func (d *dictDataDao) Insert(r *model.SysDictData) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).Insert(r)
}

// 更新数据
func (d *dictDataDao) Update(r *model.SysDictData) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).ID(r.DictCode).Update(r)
}

// 删除
func (d *dictDataDao) Delete(r *model.SysDictData) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).ID(r.DictCode).Delete(r)
}

//批量删除
func (d *dictDataDao) DeleteBatch(ids ...int64) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).In("dict_code", ids).Delete(new(model.SysDictData))
}

// 根据结构体中已有的非空数据来获得单条数据
func (d *dictDataDao) FindOne(r *model.SysDictData) (bool, error) {
	return db.Instance().Engine().Table(d.TableName()).Get(r)
}

// 根据条件查询
func (d *dictDataDao) Find(where, order string) ([]model.SysDictData, error) {
	var list []model.SysDictData
	err := db.Instance().Engine().Table(d.TableName()).Where(where).OrderBy(order).Find(&list)
	return list, err
}

//指定字段集合查询
func (d *dictDataDao) FindIn(column string, args ...interface{}) ([]model.SysDictData, error) {
	var list []model.SysDictData
	err := db.Instance().Engine().Table(d.TableName()).In(column, args).Find(&list)
	return list, err
}

//排除指定字段集合查询
func (d *dictDataDao) FindNotIn(column string, args ...interface{}) ([]model.SysDictData, error) {
	var list []model.SysDictData
	err := db.Instance().Engine().Table(d.TableName()).NotIn(column, args).Find(&list)
	return list, err
}

//根据条件分页查询数据
func (d *dictDataDao) SelectListByPage(param *model.DictDataSelectPageReq) (*[]model.SysDictData, *page.Paging, error) {
	db := db.Instance().Engine()
	p := new(page.Paging)
	if db == nil {
		return nil, p, errors.New("获取数据库连接失败")
	}

	m := db.Table(d.TableName()).Alias("t")

	if param != nil {
		if param.DictLabel != "" {
			m.Where("t.dict_label like ?", "%"+param.DictLabel+"%")
		}

		if param.Status != "" {
			m.Where("t.status = ", param.Status)
		}

		if param.DictType != "" {
			m.Where("t.dict_type like ?", "%"+param.DictType+"%")
		}

		if param.BeginTime != "" {
			m.Where("date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			m.Where("date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	tm := m.Clone()

	total, err := tm.Count()

	if err != nil {
		return nil, p, errors.New("读取行数失败")
	}

	p = page.CreatePaging(param.PageNum, param.PageSize, int(total))

	m.Limit(p.Pagesize, p.StartNum)

	var result []model.SysDictData
	m.Find(&result)
	return &result, p, nil
}

// 导出excel
func (d *dictDataDao) SelectListExport(param *model.DictDataSelectPageReq, head, col []string) (string, error) {
	db := db.Instance().Engine()

	if db == nil {
		return "", errors.New("获取数据库连接失败")
	}

	build := builder.Select(col...).From(d.TableName(), "t")

	if param != nil {
		if param.DictLabel != "" {
			build.Where(builder.Like{"t.dict_label", param.DictLabel})
		}

		if param.Status != "" {
			build.Where(builder.Eq{"t.status": param.Status})
		}

		if param.DictType != "" {
			build.Where(builder.Like{"t.dict_type", param.DictType})
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
func (d *dictDataDao) SelectListAll(param *model.DictDataSelectPageReq) ([]model.SysDictData, error) {
	db := db.Instance().Engine()

	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}

	session := db.Table(d.TableName()).Alias("t")

	if param != nil {
		if param.DictLabel != "" {
			session.Where("t.dict_label like ?", "%"+param.DictLabel+"%")
		}

		if param.Status != "" {
			session.Where("t.status = ", param.Status)
		}

		if param.DictType != "" {
			session.Where("t.dict_type like ?", "%"+param.DictType+"%")
		}

		if param.BeginTime != "" {
			session.Where("date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			session.Where("date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	var result []model.SysDictData
	session.Find(&result)
	return result, nil
}
