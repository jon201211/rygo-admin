/*
==========================================================================
RYGO自动生成菜单SQL,只生成一次,按需修改.
生成日期：2020-03-27 04:35:17 +0800 CST
==========================================================================
*/

-- 菜单 SQL
insert into sys_menu (menu_name, parent_id, order_num, url, menu_type, visible, perms, icon, create_by, create_time, update_by, update_time, remark)
values('参数配置', '4', '1', '/config', 'C', '0', 'config:view', '#', 'admin', '2020-01-01', 'admin', '2020-01-01', '参数配置菜单');

-- 按钮父菜单ID
SELECT @parentId := LAST_INSERT_ID();

-- 按钮 SQL
insert into sys_menu  (menu_name, parent_id, order_num, url, menu_type, visible, perms, icon, create_by, create_time, update_by, update_time, remark)
values('参数配置查询', @parentId, '1',  '#',  'F', '0', 'config:list',         '#', 'admin', '2020-01-01', 'admin', '2020-01-01', '');

insert into sys_menu  (menu_name, parent_id, order_num, url, menu_type, visible, perms, icon, create_by, create_time, update_by, update_time, remark)
values('参数配置新增', @parentId, '2',  '#',  'F', '0', 'config:add',          '#', 'admin', '2020-01-01', 'admin', '2020-01-01', '');

insert into sys_menu  (menu_name, parent_id, order_num, url, menu_type, visible, perms, icon, create_by, create_time, update_by, update_time, remark)
values('参数配置修改', @parentId, '3',  '#',  'F', '0', 'config:edit',         '#', 'admin', '2020-01-01', 'admin', '2020-01-01', '');

insert into sys_menu  (menu_name, parent_id, order_num, url, menu_type, visible, perms, icon, create_by, create_time, update_by, update_time, remark)
values('参数配置删除', @parentId, '4',  '#',  'F', '0', 'config:remove',       '#', 'admin', '2020-01-01', 'admin', '2020-01-01', '');

insert into sys_menu  (menu_name, parent_id, order_num, url, menu_type, visible, perms, icon, create_by, create_time, update_by, update_time, remark)
values('参数配置导出', @parentId, '5',  '#',  'F', '0', 'config:export',       '#', 'admin', '2020-01-01', 'admin', '2020-01-01', '');