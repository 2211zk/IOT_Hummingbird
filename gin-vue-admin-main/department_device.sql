-- 部门表
CREATE TABLE department (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
    parent_id INT DEFAULT NULL COMMENT '上级部门ID',
    name VARCHAR(100) NOT NULL COMMENT '部门名称',
    leader VARCHAR(50) DEFAULT NULL COMMENT '负责人',
    phone VARCHAR(20) DEFAULT NULL COMMENT '电话',
    email VARCHAR(100) DEFAULT NULL COMMENT '邮箱',
    status VARCHAR(20) DEFAULT '启用' COMMENT '状态（如：启用/禁用）',
    sort INT DEFAULT 0 COMMENT '排序',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    INDEX idx_parent_id (parent_id)
) COMMENT='部门表';

-- 设备表
CREATE TABLE device (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
    name VARCHAR(100) NOT NULL COMMENT '设备名称',
    product_name VARCHAR(100) DEFAULT NULL COMMENT '产品名称',
    -- 可根据实际需求添加更多字段
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) COMMENT='设备表';

-- 部门-设备关联表
CREATE TABLE department_device (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
    department_id INT NOT NULL COMMENT '部门ID',
    device_id INT NOT NULL COMMENT '设备ID',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    UNIQUE KEY uk_dept_device (department_id, device_id),
    FOREIGN KEY (department_id) REFERENCES department(id) ON DELETE CASCADE,
    FOREIGN KEY (device_id) REFERENCES device(id) ON DELETE CASCADE
) COMMENT='部门与设备关联表'; 