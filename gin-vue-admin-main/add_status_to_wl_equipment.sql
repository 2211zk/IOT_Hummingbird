-- 为wl_equipment表添加status字段
ALTER TABLE wl_equipment ADD COLUMN status VARCHAR(20) DEFAULT '启用' COMMENT '设备状态';

-- 更新现有记录的status字段
UPDATE wl_equipment SET status = '启用' WHERE status IS NULL; 