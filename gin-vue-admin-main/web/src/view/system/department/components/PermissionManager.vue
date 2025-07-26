<template>
  <div class="permission-manager">
    <div class="manager-header">
      <h4>权限管理</h4>
      <el-button
        type="text"
        size="small"
        @click="refreshData"
        :loading="loading"
        icon="Refresh"
      >
        刷新
      </el-button>
    </div>
    
    <div class="manager-content">
      <!-- 权限说明 -->
      <div class="permission-info">
        <el-alert
          title="权限说明"
          type="info"
          :closable="false"
          show-icon
        >
          <template #default>
            <div class="permission-rules">
              <div class="rule-item">
                <strong>超级管理员：</strong>拥有所有权限，可以管理任何部门
              </div>
              <div class="rule-item">
                <strong>管理员：</strong>可以管理非顶级部门，不能删除顶级部门
              </div>
              <div class="rule-item">
                <strong>普通用户：</strong>只能查看部门信息，不能进行增删改操作
              </div>
            </div>
          </template>
        </el-alert>
      </div>
      
      <!-- 当前用户权限 -->
      <div class="current-user">
        <div class="user-header">
          <span>当前用户权限</span>
        </div>
        <div class="user-info">
          <div class="user-role">
            <el-tag :type="getRoleType(currentUser.role)" size="large">
              {{ getRoleText(currentUser.role) }}
            </el-tag>
          </div>
          <div class="user-permissions">
            <div class="permission-grid">
              <div
                v-for="permission in permissions"
                :key="permission.key"
                class="permission-item"
                :class="{ granted: hasPermission(permission.key) }"
              >
                <div class="permission-icon">
                  <el-icon v-if="hasPermission(permission.key)">
                    <Check />
                  </el-icon>
                  <el-icon v-else>
                    <Close />
                  </el-icon>
                </div>
                <div class="permission-text">
                  <div class="permission-name">{{ permission.name }}</div>
                  <div class="permission-desc">{{ permission.description }}</div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      
      <!-- 操作日志 -->
      <div class="operation-log">
        <div class="log-header">
          <span>最近操作</span>
          <el-button
            type="text"
            size="small"
            @click="clearLog"
            v-if="operationLog.length > 0"
          >
            清空日志
          </el-button>
        </div>
        <div class="log-content">
          <div v-if="operationLog.length === 0" class="empty-log">
            <el-empty description="暂无操作记录" :image-size="60" />
          </div>
          <div v-else class="log-list">
            <div
              v-for="(log, index) in operationLog"
              :key="index"
              class="log-item"
              :class="log.type"
            >
              <div class="log-icon">
                <el-icon v-if="log.type === 'success'">
                  <CircleCheck />
                </el-icon>
                <el-icon v-else-if="log.type === 'error'">
                  <CircleClose />
                </el-icon>
                <el-icon v-else>
                  <Warning />
                </el-icon>
              </div>
              <div class="log-content-text">
                <div class="log-message">{{ log.message }}</div>
                <div class="log-time">{{ formatTime(log.time) }}</div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { Check, Close, CircleCheck, CircleClose, Warning } from '@element-plus/icons-vue'

// 响应式数据
const loading = ref(false)
const operationLog = ref([])

// 当前用户信息（模拟数据，实际应该从用户状态管理获取）
const currentUser = reactive({
  id: 1,
  name: '管理员',
  role: 'admin'
})

// 权限定义
const permissions = [
  {
    key: 'department:view',
    name: '查看部门',
    description: '查看部门列表和详情'
  },
  {
    key: 'department:create',
    name: '创建部门',
    description: '创建新的部门'
  },
  {
    key: 'department:edit',
    name: '编辑部门',
    description: '修改部门信息'
  },
  {
    key: 'department:delete',
    name: '删除部门',
    description: '删除部门（需满足删除条件）'
  },
  {
    key: 'department:status',
    name: '状态管理',
    description: '启用或禁用部门'
  },
  {
    key: 'department:sort',
    name: '排序管理',
    description: '调整部门排序'
  },
  {
    key: 'department:device',
    name: '设备关联',
    description: '管理部门设备关联'
  },
  {
    key: 'department:batch',
    name: '批量操作',
    description: '批量管理部门状态和排序'
  }
]

// 角色权限映射
const rolePermissions = {
  super_admin: [
    'department:view',
    'department:create',
    'department:edit',
    'department:delete',
    'department:status',
    'department:sort',
    'department:device',
    'department:batch'
  ],
  admin: [
    'department:view',
    'department:create',
    'department:edit',
    'department:delete',
    'department:status',
    'department:sort',
    'department:device'
  ],
  user: [
    'department:view'
  ]
}

// 检查权限
const hasPermission = (permissionKey) => {
  const userPermissions = rolePermissions[currentUser.role] || []
  return userPermissions.includes(permissionKey)
}

// 获取角色类型
const getRoleType = (role) => {
  switch (role) {
    case 'super_admin':
      return 'danger'
    case 'admin':
      return 'warning'
    case 'user':
      return 'info'
    default:
      return 'info'
  }
}

// 获取角色文本
const getRoleText = (role) => {
  switch (role) {
    case 'super_admin':
      return '超级管理员'
    case 'admin':
      return '管理员'
    case 'user':
      return '普通用户'
    default:
      return '未知角色'
  }
}

// 添加操作日志
const addLog = (message, type = 'info') => {
  operationLog.value.unshift({
    message,
    type,
    time: new Date()
  })
  
  // 限制日志数量
  if (operationLog.value.length > 50) {
    operationLog.value = operationLog.value.slice(0, 50)
  }
}

// 清空日志
const clearLog = () => {
  operationLog.value = []
  ElMessage.success('操作日志已清空')
}

// 格式化时间
const formatTime = (time) => {
  if (!time) return ''
  
  const now = new Date()
  const diff = now - time
  const minutes = Math.floor(diff / (1000 * 60))
  const hours = Math.floor(diff / (1000 * 60 * 60))
  
  if (minutes < 1) return '刚刚'
  if (minutes < 60) return `${minutes}分钟前`
  if (hours < 24) return `${hours}小时前`
  
  return time.toLocaleString()
}

// 刷新数据
const refreshData = () => {
  loading.value = true
  
  // 模拟刷新延迟
  setTimeout(() => {
    loading.value = false
    addLog('权限信息已刷新', 'success')
  }, 500)
}

// 权限检查函数（供外部调用）
const checkPermission = (permissionKey) => {
  const hasAuth = hasPermission(permissionKey)
  
  if (!hasAuth) {
    const permission = permissions.find(p => p.key === permissionKey)
    const message = `权限不足：${permission ? permission.name : permissionKey}`
    addLog(message, 'error')
    ElMessage.error(message)
  } else {
    const permission = permissions.find(p => p.key === permissionKey)
    addLog(`权限验证通过：${permission ? permission.name : permissionKey}`, 'success')
  }
  
  return hasAuth
}

// 暴露给父组件的方法
defineExpose({
  checkPermission,
  addLog,
  hasPermission
})

// 组件挂载时的初始化
onMounted(() => {
  addLog('权限管理器已初始化', 'info')
})
</script>

<style scoped>
.permission-manager {
  border: 1px solid #e4e7ed;
  border-radius: 4px;
  overflow: hidden;
}

.manager-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px 20px;
  background: #fafafa;
  border-bottom: 1px solid #e4e7ed;
}

.manager-header h4 {
  margin: 0;
  font-size: 16px;
  font-weight: 600;
  color: #303133;
}

.manager-content {
  padding: 20px;
}

.permission-info {
  margin-bottom: 24px;
}

.permission-rules {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.rule-item {
  font-size: 14px;
  line-height: 1.5;
}

.current-user {
  margin-bottom: 24px;
}

.user-header {
  font-size: 14px;
  color: #606266;
  font-weight: 500;
  margin-bottom: 12px;
}

.user-info {
  border: 1px solid #e4e7ed;
  border-radius: 4px;
  padding: 16px;
}

.user-role {
  margin-bottom: 16px;
}

.permission-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 12px;
}

.permission-item {
  display: flex;
  align-items: center;
  padding: 8px;
  border-radius: 4px;
  gap: 8px;
  transition: all 0.2s;
}

.permission-item.granted {
  background: #f0f9ff;
  border: 1px solid #b3d8ff;
}

.permission-item:not(.granted) {
  background: #fef0f0;
  border: 1px solid #fbc4c4;
}

.permission-icon {
  width: 20px;
  height: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
  font-size: 12px;
}

.permission-item.granted .permission-icon {
  background: #67c23a;
  color: white;
}

.permission-item:not(.granted) .permission-icon {
  background: #f56c6c;
  color: white;
}

.permission-text {
  flex: 1;
}

.permission-name {
  font-size: 14px;
  font-weight: 500;
  color: #303133;
  margin-bottom: 2px;
}

.permission-desc {
  font-size: 12px;
  color: #909399;
}

.operation-log {
  border-top: 1px solid #e4e7ed;
  padding-top: 20px;
}

.log-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-size: 14px;
  color: #606266;
  font-weight: 500;
  margin-bottom: 12px;
}

.log-content {
  max-height: 300px;
  overflow-y: auto;
}

.empty-log {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 120px;
}

.log-list {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.log-item {
  display: flex;
  align-items: flex-start;
  padding: 8px;
  border-radius: 4px;
  gap: 8px;
}

.log-item.success {
  background: #f0f9ff;
  border-left: 3px solid #67c23a;
}

.log-item.error {
  background: #fef0f0;
  border-left: 3px solid #f56c6c;
}

.log-item.info {
  background: #f4f4f5;
  border-left: 3px solid #909399;
}

.log-icon {
  width: 16px;
  height: 16px;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-top: 2px;
}

.log-item.success .log-icon {
  color: #67c23a;
}

.log-item.error .log-icon {
  color: #f56c6c;
}

.log-item.info .log-icon {
  color: #909399;
}

.log-content-text {
  flex: 1;
}

.log-message {
  font-size: 14px;
  color: #303133;
  margin-bottom: 2px;
}

.log-time {
  font-size: 12px;
  color: #909399;
}

:deep(.el-alert__content) {
  padding: 0;
}
</style>