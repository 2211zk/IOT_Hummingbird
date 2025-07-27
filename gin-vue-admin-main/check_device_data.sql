-- 检查设备数据
-- 1. 检查wl_equipment表是否存在
SHOW TABLES LIKE 'wl_equipment';

-- 2. 检查wl_equipment表结构
DESCRIBE wl_equipment;

-- 3. 检查是否有设备数据
SELECT COUNT(*) as device_count FROM wl_equipment;

-- 4. 查看设备数据
SELECT 
    ID,
    eq_name,
    products_id,
    status,
    created_at,
    updated_at
FROM wl_equipment 
LIMIT 10;

-- 5. 如果没有数据，插入一些测试设备
INSERT INTO wl_equipment (eq_name, products_id, eq_info, status, created_at, updated_at) VALUES
('测试设备001', 1, '这是一个测试设备', '启用', NOW(), NOW()),
('测试设备002', 2, '这是另一个测试设备', '启用', NOW(), NOW()),
('测试设备003', 3, '第三个测试设备', '启用', NOW(), NOW());

-- 6. 再次检查数据
SELECT COUNT(*) as device_count FROM wl_equipment; 