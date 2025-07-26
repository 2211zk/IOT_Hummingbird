<template>
  <el-dialog
    v-model="visible"
    :title="dialogTitle"
    width="800px"
    :close-on-click-modal="false"
    :close-on-press-escape="false"
    @close="handleClose"
  >
    <el-form
      ref="formRef"
      :model="formData"
      :rules="formRules"
      label-width="100px"
      label-position="right"
    >
      <!-- 基本信息 -->
      <div class="form-section">
        <div class="section-title">基本信息</div>
        
        <el-form-item label="上级部门" prop="parentId">
          <el-tree-select
            v-model="formData.parentId"
            :data="departmentTree"
            :props="{ label: 'name', value: 'id', children: 'children' }"
            placeholder="请选择上级部门"
            clearable
            filterable
            check-strictly
            style="width: 100%"
            :disabled="mode === 'view'"
          />
        </el-form-item>
        
        <el-form-item label="部门名称" prop="name">
          <el-input
            v-model="formData.name"
            placeholder="请输入部门名称"
            :disabled="mode === 'view'"
            maxlength="100"
            show-word-limit
          />
        </el-form-item>
        
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="负责人" prop="leader">
              <el-input
                v-model="formData.leader"
                placeholder="请输入负责人"
                :disabled="mode === 'view'"
                maxlength="32"
              />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="电话" prop="phone">
              <el-input
                v-model="formData.phone"
                placeholder="请输入电话"
                :disabled="mode === 'view'"
                maxlength="20"
              />
            </el-form-item>
          </el-col>
        </el-row>
        
        <el-form-item label="邮箱" prop="email">
          <el-input
            v-model="formData.email"
            placeholder="请输入邮箱"
            :disabled="mode === 'view'"
            maxlength="64"
          />
        </el-form-item>
        
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="状态" prop="status">
              <el-radio-group v-model="formData.status" :disabled="mode === 'view'">
                <el-radio label="启用">启用</el-radio>
                <el-radio label="禁用">禁用</el-radio>
              </el-radio-group>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="排序" prop="sort">
              <el-input-number
                v-model="formData.sort"
                :min="0"
                :max="9999"
                placeholder="请输入排序"
                :disabled="mode === 'view'"
                style="width: 100%"
              />
            </el-form-item>
          </el-col>
        </el-row>
      </div>
      
      <!-- 设备关联 -->
      <div class="form-section">
        <div class="section-title">
          设备关联
          <el-button
            v-if="mode !== 'view'"
            type="primary"
            size="small"
            @click="showDeviceSelector = true"
          >
            选择设备
          </el-button>
        </div>
        
        <div class="device-list">
          <div v-if="selectedDevices.length === 0" class="empty-devices">
            <el-empty description="暂无关联设备" :image-size="60" />
          </div>
          <div v-else class="device-tags">
            <el-tag
              v-for="device in selectedDevices"
              :key="device.id"
              :closable="mode !== 'view'"
              @close="removeDevice(device.id)"
              class="device-tag"
            >
              {{ device.deviceName }}
              <span v-if="device.productName" class="product-name">
                ({{ device.productName }})
              </span>
            </el-tag>
          </div>
        </div>
      </div>
    </el-form>
    
    <template #footer>
      <div class="dialog-footer">
        <el-button @click="handleClose">
          {{ mode === 'view' ? '关闭' : '取消' }}
        </el-button>
        <el-button
          v-if="mode !== 'view'"
          type="primary"
          @click="handleSubmit"
          :loading="submitLoading"
        >
          确定
        </el-button>
      </div>
    </template>
    
    <!-- 设备选择器 -->
    <DeviceSelector
      v-model="showDeviceSelector"
      :selected-devices="selectedDevices"
      :department-id="formData.id"
      @confirm="handleDeviceSelect"
    />
  </el-dialog>
</template>

<script setup>
import { ref, reactive, computed, watch, nextTick } from 'vue'
import { ElMessage } from 'element-plus'
import { getDepartmentTree, createDepartment, updateDepartment, getDepartmentDetail } from '@/api/wlDepartment'
import DeviceSelector from './DeviceSelector.vue'

// Props
const props = defineProps({
  modelValue: {
    type: Boolean,
    default: false
  },
  mode: {
    type: String,
    default: 'create', // create, edit, view
    validator: (value) => ['create', 'edit', 'view'].includes(value)
  },
  departmentId: {
    type: Number,
    default: null
  }
})

// Emits
const emit = defineEmits(['update:modelValue', 'success'])

// 响应式数据
const formRef = ref()
const submitLoading = ref(false)
const departmentTree = ref([])
const selectedDevices = ref([])
const showDeviceSelector = ref(false)

// 计算属性
const visible = computed({
  get: () => props.modelValue,
  set: (value) => emit('update:modelValue', value)
})

const dialogTitle = computed(() => {
  switch (props.mode) {
    case 'create':
      return '新增部门'
    case 'edit':
      return '编辑部门'
    case 'view':
      return '查看部门'
    default:
      return '部门信息'
  }
})

// 表单数据
const formData = reactive({
  id: null,
  parentId: null,
  name: '',
  leader: '',
  phone: '',
  email: '',
  status: '启用',
  sort: 0,
  deviceIds: []
})

// 表单验证规则
const formRules = reactive({
  name: [
    { required: true, message: '请输入部门名称', trigger: 'blur' },
    { min: 1, max: 100, message: '部门名称长度在 1 到 100 个字符', trigger: 'blur' }
  ],
  leader: [
    { max: 32, message: '负责人长度不能超过 32 个字符', trigger: 'blur' }
  ],
  phone: [
    { max: 20, message: '电话长度不能超过 20 个字符', trigger: 'blur' },
    { pattern: /^[0-9-+\s()]*$/, message: '请输入有效的电话号码', trigger: 'blur' }
  ],
  email: [
    { max: 64, message: '邮箱长度不能超过 64 个字符', trigger: 'blur' },
    { type: 'email', message: '请输入有效的邮箱地址', trigger: 'blur' }
  ],
  sort: [
    { type: 'number', min: 0, max: 9999, message: '排序值应在 0 到 9999 之间', trigger: 'blur' }
  ]
})

// 监听弹窗显示状态
watch(visible, async (newVal) => {
  if (newVal) {
    await loadDepartmentTree()
    if (props.mode === 'edit' || props.mode === 'view') {
      await loadDepartmentDetail()
    } else {
      resetForm()
    }
  }
})

// 加载部门树
const loadDepartmentTree = async () => {
  try {
    const excludeId = props.mode === 'edit' ? props.departmentId : null
    const response = await getDepartmentTree({ excludeId })
    if (response.code === 0) {
      departmentTree.value = response.data || []
    } else {
      ElMessage.error(response.msg || '获取部门树失败')
    }
  } catch (error) {
    console.error('获取部门树失败:', error)
    ElMessage.error('获取部门树失败')
  }
}

// 加载部门详情
const loadDepartmentDetail = async () => {
  if (!props.departmentId) return
  
  try {
    const response = await getDepartmentDetail(props.departmentId)
    if (response.code === 0) {
      const data = response.data
      Object.assign(formData, {
        id: data.id,
        parentId: data.parentId,
        name: data.name,
        leader: data.leader || '',
        phone: data.phone || '',
        email: data.email || '',
        status: data.status || '启用',
        sort: data.sort || 0,
        deviceIds: data.devices ? data.devices.map(d => d.id) : []
      })
      
      // 设置已选择的设备
      selectedDevices.value = data.devices || []
    } else {
      ElMessage.error(response.msg || '获取部门详情失败')
    }
  } catch (error) {
    console.error('获取部门详情失败:', error)
    ElMessage.error('获取部门详情失败')
  }
}

// 重置表单
const resetForm = () => {
  Object.assign(formData, {
    id: null,
    parentId: null,
    name: '',
    leader: '',
    phone: '',
    email: '',
    status: '启用',
    sort: 0,
    deviceIds: []
  })
  selectedDevices.value = []
  
  nextTick(() => {
    formRef.value?.clearValidate()
  })
}

// 处理设备选择
const handleDeviceSelect = (devices) => {
  selectedDevices.value = devices
  formData.deviceIds = devices.map(d => d.id)
}

// 移除设备
const removeDevice = (deviceId) => {
  selectedDevices.value = selectedDevices.value.filter(d => d.id !== deviceId)
  formData.deviceIds = formData.deviceIds.filter(id => id !== deviceId)
}

// 提交表单
const handleSubmit = async () => {
  if (!formRef.value) return
  
  try {
    const valid = await formRef.value.validate()
    if (!valid) return
    
    submitLoading.value = true
    
    const submitData = {
      ...formData,
      deviceIds: formData.deviceIds
    }
    
    let response
    if (props.mode === 'create') {
      response = await createDepartment(submitData)
    } else if (props.mode === 'edit') {
      response = await updateDepartment(submitData)
    }
    
    if (response.code === 0) {
      ElMessage.success(props.mode === 'create' ? '创建成功' : '更新成功')
      emit('success')
      handleClose()
    } else {
      ElMessage.error(response.msg || '操作失败')
    }
  } catch (error) {
    console.error('提交失败:', error)
    ElMessage.error('操作失败')
  } finally {
    submitLoading.value = false
  }
}

// 关闭弹窗
const handleClose = () => {
  visible.value = false
  resetForm()
}
</script>

<style scoped>
.form-section {
  margin-bottom: 24px;
}

.section-title {
  font-size: 16px;
  font-weight: 600;
  color: #303133;
  margin-bottom: 16px;
  padding-bottom: 8px;
  border-bottom: 1px solid #e4e7ed;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.device-list {
  min-height: 80px;
}

.empty-devices {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 80px;
}

.device-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}

.device-tag {
  margin: 0;
  max-width: 200px;
}

.product-name {
  color: #909399;
  font-size: 12px;
}

.dialog-footer {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
}

:deep(.el-form-item__label) {
  font-weight: 500;
}

:deep(.el-input__wrapper) {
  border-radius: 4px;
}

:deep(.el-tree-select) {
  border-radius: 4px;
}

:deep(.el-radio-group) {
  display: flex;
  gap: 16px;
}

:deep(.el-input-number) {
  width: 100%;
}

:deep(.el-input-number .el-input__wrapper) {
  width: 100%;
}
</style>