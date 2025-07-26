-- 部门表
CREATE TABLE IF NOT EXISTS wl_department (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
    parent_id INT DEFAULT NULL COMMENT '上级部门ID',
    name VARCHAR(100) NOT NULL COMMENT '部门名称',
    department_name VARCHAR(64) DEFAULT NULL COMMENT '部门名称（兼容字段）',
    leader VARCHAR(32) DEFAULT NULL COMMENT '负责人',
    phone VARCHAR(20) DEFAULT NULL COMMENT '电话',
    email VARCHAR(64) DEFAULT NULL COMMENT '邮箱',
    status VARCHAR(8) DEFAULT '启用' COMMENT '状态（如：启用/禁用）',
    sort INT DEFAULT 0 COMMENT '排序',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at DATETIME DEFAULT NULL COMMENT '删除时间',
    created_by INT DEFAULT NULL COMMENT '创建者',
    updated_by INT DEFAULT NULL COMMENT '更新者',
    deleted_by INT DEFAULT NULL COMMENT '删除者',
    INDEX idx_parent_id (parent_id),
    INDEX idx_deleted_at (deleted_at)
) COMMENT='部门表';

-- 设备表
CREATE TABLE IF NOT EXISTS wl_device (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
    device_name VARCHAR(100) NOT NULL COMMENT '设备名称',
    product_name VARCHAR(100) DEFAULT NULL COMMENT '产品名称',
    status VARCHAR(20) DEFAULT '启用' COMMENT '状态',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) COMMENT='设备表';

-- 部门-设备关联表
CREATE TABLE IF NOT EXISTS wl_department_device (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
    department_id INT NOT NULL COMMENT '部门ID',
    device_id INT NOT NULL COMMENT '设备ID',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    UNIQUE KEY uk_dept_device (department_id, device_id),
    INDEX idx_department_id (department_id),
    INDEX idx_device_id (device_id)
) COMMENT='部门与设备关联表';

-- 插入一些测试数据
INSERT INTO wl_department (name, department_name, leader, phone, email, status, sort) VALUES 
('技术部', '技术部', '张三', '13800138001', 'tech@company.com', '启用', 1),
('市场部', '市场部', '李四', '13800138002', 'market@company.com', '启用', 2),
('人事部', '人事部', '王五', '13800138003', 'hr@company.com', '启用', 3);

-- 插入一些测试设备
INSERT INTO wl_device (device_name, product_name, status) VALUES 
('服务器001', 'Dell PowerEdge', '启用'),
('服务器002', 'HP ProLiant', '启用'),
('网络设备001', 'Cisco Switch', '启用'); 