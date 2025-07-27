<template>
  <el-dialog
    v-model="visible"
    :title="dialogTitle"
    width="600px"
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
      <el-form-item label="设备名称" prop="eqName">
        <el-input
          v-model="formData.eqName"
          placeholder="请输入设备名称"
          :disabled="mode === 'view'"
          maxlength="100"
          show-word-limit
        />
      </el-form-item>
      
      <el-form-item label="产品ID" prop="productsId">
        <el-input
          v-model.number="formData.productsId"
          placeholder="请输入产品ID"
          :disabled="mode === 'view'"
          type="number"
        />
      </el-form-item>
      
      <el-form-item label="设备描述" prop="eqInfo">
        <el-input
          v-model="formData.eqInfo"
          placeholder="请输入设备描述"
          :disabled="mode === 'view'"
          type="textarea"
          :rows="3"
          maxlength="150"
          show-word-limit
        />
      </el-form-item>
      
      <el-form-item label="状态" prop="status">
        <el-radio-group v-model="formData.status" :disabled="mode === 'view'">
          <el-radio label="启用">启用</el-radio>
          <el-radio label="禁用">禁用</el-radio>
        </el-radio-group>
      </el-form-item>
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
  </el-dialog>
</template>

<script setup>
import { ref, reactive, computed, watch, nextTick } from 'vue'
import { ElMessage } from 'element-plus'
import { createDevice, updateDevice, getDeviceDetail } from '@/api/device'

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
  deviceId: {
    type: Number,
    default: null
  }
})

// Emits
const emit = defineEmits(['update:modelValue', 'success'])

// 响应式数据
const formRef = ref()
const submitLoading = ref(false)

// 计算属性
const visible = computed({
  get: () => props.modelValue,
  set: (value) => emit('update:modelValue', value)
})

const dialogTitle = computed(() => {
  switch (props.mode) {
    case 'create':
      return '新增设备'
    case 'edit':
      return '编辑设备'
    case 'view':
      return '查看设备'
    default:
      return '设备信息'
  }
})

// 表单数据
const formData = reactive({
  ID: null,
  eqName: '',
  productsId: '',
  eqInfo: '',
  status: '启用'
})

// 表单验证规则
const formRules = reactive({
  eqName: [
    { required: true, message: '请输入设备名称', trigger: 'blur' },
    { min: 1, max: 20, message: '设备名称长度在 1 到 20 个字符', trigger: 'blur' }
  ],
  productsId: [
    { required: true, message: '请输入产品ID', trigger: 'blur' },
    { type: 'number', message: '产品ID必须为数字', trigger: 'blur' }
  ],
  eqInfo: [
    { max: 150, message: '设备描述长度不能超过 150 个字符', trigger: 'blur' }
  ]
})

// 监听弹窗显示状态
watch(visible, async (newVal) => {
  if (newVal) {
    if (props.mode === 'edit' || props.mode === 'view') {
      await loadDeviceDetail()
    } else {
      resetForm()
    }
  }
})

// 加载设备详情
const loadDeviceDetail = async () => {
  if (!props.deviceId) return
  
  try {
    const response = await getDeviceDetail(props.deviceId)
    if (response.code === 0) {
      const data = response.data
      Object.assign(formData, {
        ID: data.ID,
        eqName: data.eqName || '',
        productsId: data.productsId || '',
        eqInfo: data.eqInfo || '',
        status: data.status || '启用'
      })
    } else {
      ElMessage.error(response.msg || '获取设备详情失败')
    }
  } catch (error) {
    console.error('获取设备详情失败:', error)
    ElMessage.error('获取设备详情失败')
  }
}

// 重置表单
const resetForm = () => {
  Object.assign(formData, {
    ID: null,
    eqName: '',
    productsId: '',
    eqInfo: '',
    status: '启用'
  })
  
  nextTick(() => {
    formRef.value?.clearValidate()
  })
}

// 提交表单
const handleSubmit = async () => {
  if (!formRef.value) return
  
  try {
    const valid = await formRef.value.validate()
    if (!valid) return
    
    submitLoading.value = true
    
    let response
    if (props.mode === 'create') {
      response = await createDevice(formData)
    } else if (props.mode === 'edit') {
      response = await updateDevice(formData)
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

:deep(.el-radio-group) {
  display: flex;
  gap: 16px;
}
</style>