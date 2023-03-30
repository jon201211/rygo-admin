/* ==========================================================================
 RYGO自动生成model扩展代码列表、增、删，改、查、导出，只生成一次，按需修改,再次生成不会覆盖.
 生成日期：2020-03-27 04:35:17 +0800 CST
 ==========================================================================*/
package config

import (
	"errors"
    "xorm.io/builder"
    "yj-app/app/db"
    "yj-app/app/utils/excel"
    "yj-app/app/utils/page"
)

var ConfigDao = newConfigDao()

func newConfigDao() *configDao {
	return &configDao{}
}

type configDao struct {
}

//映射数据表
func (d *configDao) TableName() string {
	return "sys_config"
}

// 插入数据
func (d *configDao)  Insert(e *config) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).Insert(e)
}

// 更新数据
func (d *configDao) Update(e *config) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).ID(e.ConfigId).Update(e)
}

// 删除
func (d *configDao) Delete(e *config) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).ID(e.ConfigId).Delete(e)
}

//批量删除
func (d *configDao) DeleteBatch(ids ...int64) (int64, error) {
	return db.Instance().Engine().Table(d.TableName()).In("config_id", ids).Delete(new(config))
}

// 根据结构体中已有的非空数据来获得单条数据
func (d *configDao)  FindOne(e *config) (bool, error) {
	return db.Instance().Engine().Table(d.TableName()).Get(e)
}

// 根据条件查询
func (d *configDao) Find(where, order string) ([]config, error) {
	var list []config
	err := db.Instance().Engine().Table(d.TableName()).Where(where).OrderBy(order).Find(&list)
	return list, err
}

//指定字段集合查询
func (d *configDao) FindIn(column string, args ...interface{}) ([]config, error) {
	var list []config
	err := db.Instance().Engine().Table(d.TableName()).In(column, args).Find(&list)
	return list, err
}

//排除指定字段集合查询
func (d *configDao) FindNotIn(column string, args ...interface{}) ([]config, error) {
	var list []config
	err := db.Instance().Engine().Table(d.TableName()).NotIn(column, args).Find(&list)
	return list, err
}

//根据条件分页查询数据
func (d *configDao) SelectListByPage(param *configSelectPageReq) ([]config, *page.Paging, error) {
	db := db.Instance().Engine()
    p := new(page.Paging)

	if db == nil {
		return nil, p, errors.New("获取数据库连接失败")
	}

	model := db.Table("sys_config").Alias("t")

	if param != nil {  
		 
		if param.ConfigId != 0 {
			model.Where("t.config_id = ?", param.ConfigId)
		}
		    
		
		if param.ConfigName != "" {
			model.Where("t.config_name like ?", "%"+param.ConfigName+"%")
		}    
		 
		if param.ConfigKey != "" {
			model.Where("t.config_key = ?", param.ConfigKey)
		}     
		 
		if param.ConfigValue != "" {
			model.Where("t.config_value = ?", param.ConfigValue)
		}     
		 
		if param.ConfigType != "" {
			model.Where("t.config_type = ?", param.ConfigType)
		}              
		if param.BeginTime != "" {
			model.Where("t.create_time >= ?", param.BeginTime)
		}

		if param.EndTime != "" {
			 model.Where("t.create_time<=?", param.EndTime)
		}
	}

	total, err := model.Clone().Count()
	if err != nil {
		return nil, p, errors.New("读取行数失败")
	}

	p = page.CreatePaging(param.PageNum, param.PageSize, int(total))
	model.Limit(p.Pagesize, p.StartNum)

	var result []config
    err = model.Find(&result)
    return result, p, err
}

// 导出excel
func (d *configDao) SelectListExport(param *configSelectPageReq, head, col []string) (string, error) {
	db := db.Instance().Engine()

	if db == nil {
		return "", errors.New("获取数据库连接失败")
	}

	build := builder.Select(col...).From("sys_config", "t")

	if param != nil {  
		 
		if param.ConfigId != 0 {
			build.Where(builder.Eq{"t.config_id": param.ConfigId})
		}
		    
		
		if param.ConfigName != "" {
			build.Where(builder.Like{"t.config_name", param.ConfigName})
		}    
		 
		if param.ConfigKey != "" {
			build.Where(builder.Eq{"t.config_key": param.ConfigKey})
		}     
		 
		if param.ConfigValue != "" {
			build.Where(builder.Eq{"t.config_value": param.ConfigValue})
		}     
		 
		if param.ConfigType != "" {
			build.Where(builder.Eq{"t.config_type": param.ConfigType})
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
func (d *configDao) SelectListAll(param *configSelectPageReq) ([]config, error) {
	db := db.Instance().Engine()

	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}

	model := db.Table("sys_config").Alias("t")

	if param != nil {  
		 
		if param.ConfigId != 0 {
			model.Where("t.config_id = ?", param.ConfigId)
		}
		   
		
		if param.ConfigName != "" {
			model.Where("t.config_name like ?", "%"+param.ConfigName+"%")
		}    
		 
		if param.ConfigKey != "" {
			model.Where("t.config_key = ?", param.ConfigKey)
		} 
		   
		 
		if param.ConfigValue != "" {
			model.Where("t.config_value = ?", param.ConfigValue)
		} 
		   
		 
		if param.ConfigType != "" {
			model.Where("t.config_type = ?", param.ConfigType)
		} 
		            
		if param.BeginTime != "" {
			model.Where("date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			model.Where("date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	var result []config
	err := model.Find(&result)
	return result, err
}