<template>
  <el-dialog
    v-model="dialogVisible"
    :title="formType === 'add' ? '添加驱动' : '编辑驱动'"
    width="600px"
    :before-close="handleClose"
  >
    <el-form
      ref="formRef"
      :model="formData"
      :rules="rules"
      label-width="100px"
    >
      <el-form-item label="驱动名称" prop="driverName">
        <el-input
          v-model="formData.driverName"
          placeholder="请输入驱动名称"
          clearable
        />
      </el-form-item>
      
      <el-form-item label="驱动编号" prop="driverId">
        <el-input
          v-model="formData.driverId"
          placeholder="请输入驱动编号"
          clearable
        />
      </el-form-item>
      
      <el-form-item label="版本" prop="version">
        <el-input
          v-model="formData.version"
          placeholder="请输入版本号"
          clearable
        />
      </el-form-item>
      
      <el-form-item label="驱动类型" prop="driverType">
        <el-select
          v-model="formData.driverType"
          placeholder="请选择驱动类型"
          clearable
          style="width: 100%"
        >
          <el-option
            v-for="item in driverTypeOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
      </el-form-item>
      
      <el-form-item label="驱动状态" prop="status">
        <el-input
          v-model="formData.status"
          placeholder="请输入驱动状态"
          clearable
        />
      </el-form-item>
      
      <el-form-item label="协议类型" prop="protocolType">
        <el-input
          v-model="formData.protocolType"
          placeholder="请输入协议类型"
          clearable
        />
      </el-form-item>
      
      <el-form-item label="设备类型" prop="deviceCategory">
        <el-input
          v-model="formData.deviceCategory"
          placeholder="请输入设备类型"
          clearable
        />
      </el-form-item>
    </el-form>
    
    <template #footer>
      <div class="dialog-footer">
        <el-button @click="handleClose">取消</el-button>
        <el-button type="primary" @click="handleSubmit" :loading="loading">
          {{ formType === 'add' ? '添加' : '更新' }}
        </el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script setup>
import { ref, reactive, watch, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { createWlDrivers, updateWlDrivers } from '@/api/wl_driver/wlDrivers'
import { getDict } from '@/utils/dictionary'

const props = defineProps({
  visible: {
    type: Boolean,
    default: false
  },
  formType: {
    type: String,
    default: 'add'
  },
  driverData: {
    type: Object,
    default: () => ({})
  }
})

const emit = defineEmits(['update:visible', 'success'])

const dialogVisible = ref(false)
const loading = ref(false)
const formRef = ref()

// 字典选项
const driverTypeOptions = ref([])

// 表单数据
const formData = reactive({
  driverName: '',
  driverId: '',
  version: '',
  driverType: '',
  status: '',
  protocolType: '',
  deviceCategory: ''
})

// 表单验证规则
const rules = {
  driverName: [
    { required: true, message: '请输入驱动名称', trigger: 'blur' }
  ],
  driverId: [
    { required: true, message: '请输入驱动编号', trigger: 'blur' }
  ],
  version: [
    { required: true, message: '请输入版本号', trigger: 'blur' }
  ],
  driverType: [
    { required: true, message: '请选择驱动类型', trigger: 'change' }
  ],
  status: [
    { required: true, message: '请输入驱动状态', trigger: 'blur' }
  ]
}

// 监听visible变化
watch(() => props.visible, (val) => {
  dialogVisible.value = val
  if (val) {
    loadDictionaries()
    if (props.formType === 'edit' && props.driverData) {
      Object.assign(formData, props.driverData)
    }
  }
})

// 监听dialogVisible变化
watch(dialogVisible, (val) => {
  emit('update:visible', val)
  if (!val) {
    resetForm()
  }
})

// 获取字典数据
const loadDictionaries = async () => {
  try {
    const driverTypeDict = await getDict('driver_type')
    driverTypeOptions.value = driverTypeDict || []
  } catch (error) {
    console.error('加载字典数据失败:', error)
  }
}

// 重置表单
const resetForm = () => {
  Object.assign(formData, {
    driverName: '',
    driverId: '',
    version: '',
    driverType: '',
    status: '',
    protocolType: '',
    deviceCategory: ''
  })
  formRef.value?.clearValidate()
}

// 关闭对话框
const handleClose = () => {
  dialogVisible.value = false
}

// 提交表单
const handleSubmit = async () => {
  if (!formRef.value) return
  
  try {
    await formRef.value.validate()
    loading.value = true
    
    let res
    if (props.formType === 'add') {
      res = await createWlDrivers(formData)
    } else {
      res = await updateWlDrivers(formData)
    }
    
    if (res.code === 0) {
      ElMessage.success(props.formType === 'add' ? '添加成功' : '更新成功')
      emit('success')
      handleClose()
    } else {
      ElMessage.error(res.msg || '操作失败')
    }
  } catch (error) {
    console.error('表单提交失败:', error)
    ElMessage.error('操作失败')
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  loadDictionaries()
})
</script>

<style scoped>
.dialog-footer {
  text-align: right;
}
</style> 