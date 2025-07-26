# 部门管理设计文档

## 概述

部门管理系统基于gin-vue-admin框架开发，采用前后端分离架构。后端使用Gin框架提供RESTful API，前端使用Vue3 + Element Plus构建用户界面。系统支持树形结构的部门管理，包括部门的增删改查、设备关联等功能。

## 架构设计

### 整体架构
```
前端 (Vue3 + Element Plus)
    ↓ HTTP API
后端 (Gin + GORM)
    ↓ SQL
数据库 (MySQL)
```

### 技术栈
- **后端**: Go + Gin + GORM + MySQL
- **前端**: Vue3 + Element Plus + Axios + Pinia
- **数据库**: MySQL 8.0+

## 组件和接口设计

### 数据模型

#### 1. 部门模型 (WlDepartment)
```go
type WlDepartment struct {
    ID             int       `json:"id" gorm:"primaryKey;autoIncrement"`
    ParentID       *int      `json:"parentId" gorm:"column:parent_id;index"`
    Name           string    `json:"name" gorm:"column:name;not null;size:100"`
    Leader         string    `json:"leader" gorm:"column:leader;size:32"`
    Phone          string    `json:"phone" gorm:"column:phone;size:20"`
    Email          string    `json:"email" gorm:"column:email;size:64"`
    Status         string    `json:"status" gorm:"column:status;size:8;default:启用"`
    Sort           int       `json:"sort" gorm:"column:sort;default:0"`
    CreatedAt      time.Time `json:"createdAt" gorm:"column:created_at"`
    UpdatedAt      time.Time `json:"updatedAt" gorm:"column:updated_at"`
    DeletedAt      gorm.DeletedAt `json:"deletedAt" gorm:"column:deleted_at;index"`
    DepartmentName string    `json:"departmentName" gorm:"column:department_name;size:64"` // 兼容字段
    CreatedBy      int       `json:"createdBy" gorm:"column:created_by"`
    UpdatedBy      int       `json:"updatedBy" gorm:"column:updated_by"`
    DeletedBy      int       `json:"deletedBy" gorm:"column:deleted_by"`
    
    // 关联字段
    Children []WlDepartment `json:"children" gorm:"-"`
    Devices  []WlDevice     `json:"devices" gorm:"many2many:wl_department_device;"`
}
```

#### 2. 部门设备关联模型 (WlDepartmentDevice)
```go
type WlDepartmentDevice struct {
    ID           int       `json:"id" gorm:"primaryKey;autoIncrement"`
    DepartmentID int       `json:"departmentId" gorm:"column:department_id;not null"`
    DeviceID     int       `json:"deviceId" gorm:"column:device_id;not null"`
    CreatedAt    time.Time `json:"createdAt" gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
}
```

#### 3. 设备模型 (WlDevice) - 引用现有
```go
type WlDevice struct {
    ID          int    `json:"id" gorm:"primaryKey"`
    DeviceName  string `json:"deviceName" gorm:"column:device_name"`
    ProductName string `json:"productName" gorm:"column:product_name"`
    Status      string `json:"status"`
    // 其他字段...
}
```

### API接口设计

#### 1. 部门管理API
```go
// GET /api/v1/department/list - 获取部门列表（支持树形和平铺）
type DepartmentListRequest struct {
    Page     int    `json:"page" form:"page"`
    PageSize int    `json:"pageSize" form:"pageSize"`
    Name     string `json:"name" form:"name"`
    Status   string `json:"status" form:"status"`
    TreeMode bool   `json:"treeMode" form:"treeMode"` // 是否返回树形结构
}

// POST /api/v1/department/create - 创建部门
type CreateDepartmentRequest struct {
    ParentID *int   `json:"parentId"`
    Name     string `json:"name" binding:"required"`
    Leader   string `json:"leader"`
    Phone    string `json:"phone"`
    Email    string `json:"email"`
    Status   string `json:"status"`
    Sort     int    `json:"sort"`
    DeviceIDs []int `json:"deviceIds"` // 关联的设备ID列表
}

// PUT /api/v1/department/update - 更新部门
type UpdateDepartmentRequest struct {
    ID       int    `json:"id" binding:"required"`
    ParentID *int   `json:"parentId"`
    Name     string `json:"name" binding:"required"`
    Leader   string `json:"leader"`
    Phone    string `json:"phone"`
    Email    string `json:"email"`
    Status   string `json:"status"`
    Sort     int    `json:"sort"`
    DeviceIDs []int `json:"deviceIds"`
}

// DELETE /api/v1/department/delete - 删除部门
type DeleteDepartmentRequest struct {
    ID int `json:"id" binding:"required"`
}

// GET /api/v1/department/tree - 获取部门树（用于选择上级部门）
type DepartmentTreeResponse struct {
    ID       int                      `json:"id"`
    Name     string                   `json:"name"`
    Children []DepartmentTreeResponse `json:"children"`
}
```

#### 2. 设备关联API
```go
// GET /api/v1/department/devices/available - 获取可关联的设备列表
type AvailableDevicesRequest struct {
    Page         int    `json:"page" form:"page"`
    PageSize     int    `json:"pageSize" form:"pageSize"`
    DeviceName   string `json:"deviceName" form:"deviceName"`
    ProductName  string `json:"productName" form:"productName"`
    DepartmentID int    `json:"departmentId" form:"departmentId"` // 排除已关联的设备
}

// GET /api/v1/department/{id}/devices - 获取部门已关联的设备
type DepartmentDevicesResponse struct {
    List  []WlDevice `json:"list"`
    Total int64      `json:"total"`
}
```

### 前端组件设计

#### 1. 主页面组件 (DepartmentManagement.vue)
```vue
<template>
  <div class="department-management">
    <!-- 搜索区域 -->
    <div class="search-area">
      <el-form :inline="true">
        <el-form-item label="部门名称">
          <el-input v-model="searchForm.name" placeholder="请输入部门名称" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    
    <!-- 操作区域 -->
    <div class="action-area">
      <el-button type="primary" @click="handleAdd">新增</el-button>
    </div>
    
    <!-- 表格区域 -->
    <el-table :data="tableData" row-key="id" :tree-props="{children: 'children'}">
      <el-table-column prop="name" label="部门名称" />
      <el-table-column prop="leader" label="负责人" />
      <el-table-column prop="phone" label="电话" />
      <el-table-column prop="email" label="邮箱" />
      <el-table-column prop="status" label="状态">
        <template #default="scope">
          <el-tag :type="scope.row.status === '启用' ? 'success' : 'danger'">
            {{ scope.row.status }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="sort" label="排序" />
      <el-table-column prop="createdAt" label="创建时间" />
      <el-table-column label="操作" width="200">
        <template #default="scope">
          <el-button size="small" @click="handleEdit(scope.row)">查看</el-button>
          <el-button size="small" @click="handleEdit(scope.row)">编辑</el-button>
          <el-button size="small" type="danger" @click="handleDelete(scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>
```

#### 2. 部门编辑弹窗组件 (DepartmentDialog.vue)
```vue
<template>
  <el-dialog v-model="visible" :title="isEdit ? '编辑' : '新增'" width="800px">
    <el-form :model="form" :rules="rules" ref="formRef" label-width="100px">
      <!-- 基本信息 -->
      <el-form-item label="上级部门" prop="parentId">
        <el-tree-select
          v-model="form.parentId"
          :data="departmentTree"
          :props="{ label: 'name', value: 'id' }"
          placeholder="请选择上级部门"
          clearable
        />
      </el-form-item>
      
      <el-form-item label="部门名称" prop="name">
        <el-input v-model="form.name" placeholder="请输入部门名称" />
      </el-form-item>
      
      <el-form-item label="负责人">
        <el-input v-model="form.leader" placeholder="请输入负责人" />
      </el-form-item>
      
      <el-form-item label="电话">
        <el-input v-model="form.phone" placeholder="请输入电话" />
      </el-form-item>
      
      <el-form-item label="邮箱">
        <el-input v-model="form.email" placeholder="请输入邮箱" />
      </el-form-item>
      
      <el-form-item label="状态">
        <el-radio-group v-model="form.status">
          <el-radio label="启用">启用</el-radio>
          <el-radio label="禁用">禁用</el-radio>
        </el-radio-group>
      </el-form-item>
      
      <el-form-item label="排序">
        <el-input-number v-model="form.sort" :min="0" />
      </el-form-item>
      
      <!-- 设备关联 -->
      <el-form-item label="关联设备">
        <el-button @click="showDeviceSelector = true">选择设备</el-button>
        <div v-if="selectedDevices.length > 0" class="selected-devices">
          <el-tag
            v-for="device in selectedDevices"
            :key="device.id"
            closable
            @close="removeDevice(device.id)"
          >
            {{ device.deviceName }}
          </el-tag>
        </div>
      </el-form-item>
    </el-form>
    
    <template #footer>
      <el-button @click="visible = false">取消</el-button>
      <el-button type="primary" @click="handleSubmit">确定</el-button>
    </template>
    
    <!-- 设备选择器 -->
    <DeviceSelector
      v-model="showDeviceSelector"
      :selected-devices="selectedDevices"
      @confirm="handleDeviceSelect"
    />
  </el-dialog>
</template>
```

#### 3. 设备选择器组件 (DeviceSelector.vue)
```vue
<template>
  <el-dialog v-model="visible" title="选择设备" width="1000px">
    <div class="device-selector">
      <!-- 搜索区域 -->
      <el-form :inline="true">
        <el-form-item label="设备名称">
          <el-input v-model="searchForm.deviceName" placeholder="请输入设备名称" />
        </el-form-item>
        <el-form-item label="产品名称">
          <el-input v-model="searchForm.productName" placeholder="请输入产品名称" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
        </el-form-item>
      </el-form>
      
      <!-- 设备列表 -->
      <div class="device-content">
        <div class="available-devices">
          <h4>待选设备</h4>
          <el-table
            :data="availableDevices"
            @selection-change="handleSelectionChange"
          >
            <el-table-column type="selection" width="55" />
            <el-table-column prop="deviceName" label="设备名称" />
            <el-table-column prop="productName" label="产品名称" />
          </el-table>
          <el-pagination
            v-model:current-page="pagination.page"
            v-model:page-size="pagination.pageSize"
            :total="pagination.total"
            @current-change="handlePageChange"
          />
        </div>
        
        <div class="selected-devices">
          <h4>已选设备</h4>
          <div class="device-list">
            <el-tag
              v-for="device in tempSelectedDevices"
              :key="device.id"
              closable
              @close="removeFromSelected(device.id)"
            >
              {{ device.deviceName }}
            </el-tag>
          </div>
        </div>
      </div>
    </div>
    
    <template #footer>
      <el-button @click="visible = false">取消</el-button>
      <el-button type="primary" @click="handleConfirm">确定</el-button>
    </template>
  </el-dialog>
</template>
```

## 数据模型

### 数据库表结构

#### 1. wl_department 表
- 主键: id (int, AUTO_INCREMENT)
- 上级部门: parent_id (bigint, 可为NULL)
- 部门名称: name (varchar(100), NOT NULL)
- 负责人: leader (varchar(32))
- 电话: phone (varchar(20))
- 邮箱: email (varchar(64))
- 状态: status (varchar(8), 默认'启用')
- 排序: sort (bigint, 默认0)
- 时间戳: created_at, updated_at, deleted_at
- 审计字段: created_by, updated_by, deleted_by

#### 2. wl_department_device 表
- 主键: id (int, AUTO_INCREMENT)
- 部门ID: department_id (int, NOT NULL)
- 设备ID: device_id (int, NOT NULL)
- 创建时间: created_at (datetime)
- 唯一约束: (department_id, device_id)
- 外键约束: 级联删除

### 数据流设计

#### 1. 部门列表查询流程
```
前端发起请求 → API接收参数 → Service层处理业务逻辑 → 
Repository层查询数据库 → 构建树形结构 → 返回结果
```

#### 2. 部门创建/更新流程
```
前端提交表单 → API验证参数 → Service层处理业务逻辑 → 
验证上级部门关系 → 保存部门信息 → 处理设备关联 → 返回结果
```

## 错误处理

### 业务规则验证
1. **部门名称唯一性**: 同级部门名称不能重复
2. **上级部门验证**: 不能选择自身或子部门作为上级
3. **删除限制**: 有子部门的部门不能删除
4. **状态继承**: 禁用部门时，其子部门也应被禁用

### 错误响应格式
```json
{
  "code": 7000,
  "data": null,
  "msg": "部门名称已存在"
}
```

### 常见错误码
- 7001: 部门名称已存在
- 7002: 上级部门不能是自身或子部门
- 7003: 该部门下还有子部门，无法删除
- 7004: 部门不存在

## 测试策略

### 单元测试
- Model层: 数据验证、关联关系测试
- Service层: 业务逻辑测试
- API层: 接口参数验证测试

### 集成测试
- 部门CRUD操作完整流程测试
- 设备关联功能测试
- 树形结构构建测试

### 前端测试
- 组件渲染测试
- 用户交互测试
- API调用测试

## 性能优化

### 数据库优化
1. 为parent_id字段添加索引
2. 使用软删除避免数据丢失
3. 分页查询减少数据传输量

### 前端优化
1. 树形数据懒加载
2. 设备选择器虚拟滚动
3. 防抖搜索减少API调用

### 缓存策略
1. 部门树结构缓存
2. 设备列表缓存
3. 用户权限缓存