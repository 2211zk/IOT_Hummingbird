<template>
  <el-dialog
    v-model="visible"
    title="删除确认"
    width="600px"
    :close-on-click-modal="false"
    :close-on-press-escape="false"
    @close="handleClose"
  >
    <div class="delete-confirm">
      <!-- 删除警告 -->
      <div class="warning-section">
        <el-alert
          title="删除警告"
          type="warning"
          :closable="false"
          show-icon
        >
          <template #default>
            <p>您即将删除部门"<strong>{{ departmentInfo.name }}</strong>"，此操作不可撤销！</p>
          </template>
        </el-alert>
      </div>
      
      <!-- 部门信息 -->
      <div class="department-info">
        <div class="info-header">
          <span>部门信息</span>
        </div>
        <div class="info-content">
          <div class="info-item">
            <span class="label">部门名称：</span>
            <span class="value">{{ departmentInfo.name }}</span>
          </div>
          <div class="info-item">
            <span class="label">负责人：</span>
            <span class="value">{{ departmentInfo.leader || '-' }}</span>
          </div>
          <div class="info-item">
            <span class="label">创建时间：</span>
            <span class="value">{{ formatDateTime(departmentInfo.createdAt) }}</span>
          </div>
          <div class="info-item">
            <span class="label">当前状态：</span>
            <el-tag :type="departmentInfo.status === '启用' ? 'success' : 'danger'">
              {{ departmentInfo.status }}
            </el-tag>
          </div>
        </div>
      </div>
      
      <!-- 影响分析 -->
      <div class="impact-analysis">
        <div class="analysis-header">
          <span>影响分析</span>
        </div>
        <div class="analysis-content">
          <!-- 子部门检查 -->
          <div class="check-item" :class="{ error: checkResults.hasChildren }">
            <div class="check-icon">
              <el-icon v-if="checkResults.hasChildren">
                <CircleClose />
              </el-icon>
              <el-icon v-else>
                <CircleCheck />
              </el-icon>
            </div>
            <div class="check-content">
              <div class="check-title">子部门检查</div>
              <div class="check-desc">
                {{ checkResults.hasChildren ? 
                  `发现 ${checkResults.childrenCount} 个子部门，需要先删除子部门` : 
                  '没有子部门，可以安全删除' 
                }}
              </div>
            </div>
          </div>
          
          <!-- 设备关联检查 -->
          <div class="check-item" :class="{ warning: checkResults.hasDevices }">
            <div class="check-icon">
              <el-icon v-if="checkResults.hasDevices">
                <Warning />
              </el-icon>
              <el-icon v-else>
                <CircleCheck />
              </el-icon>
            </div>
            <div class="check-content">
              <div class="check-title">设备关联检查</div>
              <div class="check-desc">
                {{ checkResults.hasDevices ? 
                  `关联了 ${checkResults.devicesCount} 个设备，删除后将解除关联` : 
                  '没有关联设备' 
                }}
              </div>
            </div>
          </div>
          
          <!-- 权限检查 -->
          <div class="check-item" :class="{ error: !checkResults.hasPermission }">
            <div class="check-icon">
              <el-icon v-if="checkResults.hasPermission">
                <CircleCheck />
              </el-icon>
              <el-icon v-else>
                <CircleClose />
              </el-icon>
            </div>
            <div class="check-content">
              <div class="check-title">权限检查</div>
              <div class="check-desc">
                {{ checkResults.hasPermission ? 
                  '您有删除此部门的权限' : 
                  checkResults.permissionMessage 
                }}
              </div>
            </div>
          </div>
        </div>
      </div>
      
      <!-- 删除选项 -->
      <div class="delete-options" v-if="canDelete">
        <div class="options-header">
          <span>删除选项</span>
        </div>
        <div class="options-content">
          <el-checkbox v-model="deleteOptions.forceDelete" v-if="checkResults.hasDevices">
            强制删除（忽略设备关联）
          </el-checkbox>
          <el-checkbox v-model="deleteOptions.backupData">
            删除前备份数据
          </el-checkbox>
          <el-checkbox v-model="deleteOptions.notifyUsers">
            通知相关用户
          </el-checkbox>
        </div>
      </div>
      
      <!-- 确认输入 -->
      <div class="confirm-input" v-if="canDelete">
        <div class="input-header">
          <span>确认删除</span>
        </div>
        <div class="input-content">
          <p>请输入部门名称 "<strong>{{ departmentInfo.name }}</strong>" 以确认删除：</p>
          <el-input
            v-model="confirmText"
            placeholder="请输入部门名称"
            @keyup.enter="handleConfirm"
          />
        </div>
      </div>
    </div>
    
    <template #footer>
      <div class="dialog-footer">
        <el-button @click="handleClose">取消</el-button>
        <el-button
          type="danger"
          @click="handleConfirm"
          :disabled="!canConfirm"
          :loading="deleteLoading"
        >
          确认删除
        </el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script setup>
import { ref, reactive, computed, watch } from 'vue'
import { ElMessage } from 'element-plus'
import { CircleCheck, CircleClose, Warning } from '@element-plus/icons-vue'
import { getDepartmentList, getDepartmentDevices, deleteDepartment } from '@/api/wlDepartment'
import { formatDateTime } from '@/utils/format'

// Props
const props = defineProps({
  modelValue: {
    type: Boolean,
    default: false
  },
  departmentInfo: {
    type: Object,
    default: () => ({})
  }
})

// Emits
const emit = defineEmits(['update:modelValue', 'success'])

// 响应式数据
const deleteLoading = ref(false)
const confirmText = ref('')

// 检查结果
const checkResults = reactive({
  hasChildren: false,
  childrenCount: 0,
  hasDevices: false,
  devicesCount: 0,
  hasPermission: true,
  permissionMessage: ''
})

// 删除选项
const deleteOptions = reactive({
  forceDelete: false,
  backupData: true,
  notifyUsers: false
})

// 计算属性
const visible = computed({
  get: () => props.modelValue,
  set: (value) => emit('update:modelValue', value)
})

const canDelete = computed(() => {
  return !checkResults.hasChildren && checkResults.hasPermission
})

const canConfirm = computed(() => {
  return canDelete.value && confirmText.value === props.departmentInfo.name
})

// 监听弹窗显示状态
watch(visible, async (newVal) => {
  if (newVal && props.departmentInfo.id) {
    await performChecks()
  } else {
    resetData()
  }
})

// 执行检查
const performChecks = async () => {
  await Promise.all([
    checkChildren(),
    checkDevices(),
    checkPermission()
  ])
}

// 检查子部门
const checkChildren = async () => {
  try {
    const response = await getDepartmentList({
      page: 1,
      pageSize: 1,
      parentId: props.departmentInfo.id
    })
    
    if (response.code === 0) {
      checkResults.hasChildren = response.data.total > 0
      checkResults.childrenCount = response.data.total
    }
  } catch (error) {
    console.error('检查子部门失败:', error)
  }
}

// 检查设备关联
const checkDevices = async () => {
  try {
    const response = await getDepartmentDevices({
      departmentId: props.departmentInfo.id,
      page: 1,
      pageSize: 1
    })
    
    if (response.code === 0) {
      checkResults.hasDevices = response.data.total > 0
      checkResults.devicesCount = response.data.total
    }
  } catch (error) {
    console.error('检查设备关联失败:', error)
  }
}

// 检查权限
const checkPermission = () => {
  // 模拟权限检查逻辑
  const userRole = getUserRole()
  
  if (userRole === 'super_admin') {
    checkResults.hasPermission = true
  } else if (userRole === 'admin') {
    if (!props.departmentInfo.parentId) {
      checkResults.hasPermission = false
      checkResults.permissionMessage = '管理员不能删除顶级部门'
    } else {
      checkResults.hasPermission = true
    }
  } else {
    checkResults.hasPermission = false
    checkResults.permissionMessage = '您没有删除部门的权限'
  }
}

// 获取用户角色
const getUserRole = () => {
  // 实际应该从用户状态管理获取
  return 'admin'
}

// 确认删除
const handleConfirm = async () => {
  if (!canConfirm.value) {
    ElMessage.warning('请正确输入部门名称以确认删除')
    return
  }
  
  try {
    deleteLoading.value = true
    
    // 执行删除前的准备工作
    if (deleteOptions.backupData) {
      await backupDepartmentData()
    }
    
    if (deleteOptions.notifyUsers) {
      await notifyRelatedUsers()
    }
    
    // 执行删除
    const response = await deleteDepartment({ 
      id: props.departmentInfo.id,
      force: deleteOptions.forceDelete
    })
    
    if (response.code === 0) {
      ElMessage.success('部门删除成功')
      emit('success')
      handleClose()
    } else {
      ElMessage.error(response.msg || '删除失败')
    }
  } catch (error) {
    console.error('删除部门失败:', error)
    ElMessage.error('删除失败')
  } finally {
    deleteLoading.value = false
  }
}

// 备份部门数据
const backupDepartmentData = async () => {
  // 模拟备份操作
  return new Promise(resolve => {
    setTimeout(() => {
      console.log('部门数据已备份')
      resolve()
    }, 500)
  })
}

// 通知相关用户
const notifyRelatedUsers = async () => {
  // 模拟通知操作
  return new Promise(resolve => {
    setTimeout(() => {
      console.log('已通知相关用户')
      resolve()
    }, 300)
  })
}

// 重置数据
const resetData = () => {
  confirmText.value = ''
  Object.assign(checkResults, {
    hasChildren: false,
    childrenCount: 0,
    hasDevices: false,
    devicesCount: 0,
    hasPermission: true,
    permissionMessage: ''
  })
  Object.assign(deleteOptions, {
    forceDelete: false,
    backupData: true,
    notifyUsers: false
  })
}

// 关闭弹窗
const handleClose = () => {
  visible.value = false
  resetData()
}
</script>

<style scoped>
.delete-confirm {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.warning-section {
  margin-bottom: 4px;
}

.department-info,
.impact-analysis,
.delete-options,
.confirm-input {
  border: 1px solid #e4e7ed;
  border-radius: 4px;
  overflow: hidden;
}

.info-header,
.analysis-header,
.options-header,
.input-header {
  padding: 12px 16px;
  background: #fafafa;
  border-bottom: 1px solid #e4e7ed;
  font-size: 14px;
  font-weight: 500;
  color: #303133;
}

.info-content,
.analysis-content,
.options-content,
.input-content {
  padding: 16px;
}

.info-item {
  display: flex;
  align-items: center;
  margin-bottom: 8px;
}

.info-item:last-child {
  margin-bottom: 0;
}

.label {
  width: 80px;
  color: #606266;
  font-size: 14px;
}

.value {
  color: #303133;
  font-size: 14px;
}

.check-item {
  display: flex;
  align-items: flex-start;
  padding: 12px;
  border-radius: 4px;
  margin-bottom: 8px;
  gap: 12px;
}

.check-item:last-child {
  margin-bottom: 0;
}

.check-item.error {
  background: #fef0f0;
  border: 1px solid #fbc4c4;
}

.check-item.warning {
  background: #fdf6ec;
  border: 1px solid #f5dab1;
}

.check-item:not(.error):not(.warning) {
  background: #f0f9ff;
  border: 1px solid #b3d8ff;
}

.check-icon {
  width: 20px;
  height: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
  margin-top: 2px;
}

.check-item.error .check-icon {
  color: #f56c6c;
}

.check-item.warning .check-icon {
  color: #e6a23c;
}

.check-item:not(.error):not(.warning) .check-icon {
  color: #67c23a;
}

.check-content {
  flex: 1;
}

.check-title {
  font-size: 14px;
  font-weight: 500;
  color: #303133;
  margin-bottom: 4px;
}

.check-desc {
  font-size: 13px;
  color: #606266;
  line-height: 1.4;
}

.options-content {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.input-content p {
  margin: 0 0 12px 0;
  font-size: 14px;
  color: #606266;
}

.dialog-footer {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
}

:deep(.el-alert__content) {
  padding: 0;
}

:deep(.el-alert__content p) {
  margin: 0;
  font-size: 14px;
}
</style>