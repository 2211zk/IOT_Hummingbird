
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="驱动编号:" prop="driverNum">
    <el-input v-model="formData.driverNum" :clearable="true" placeholder="请输入驱动编号" />
</el-form-item>
        <el-form-item label="驱动名称:" prop="driverName">
    <el-input v-model="formData.driverName" :clearable="true" placeholder="请输入驱动名称" />
</el-form-item>
        <el-form-item label="版本:" prop="version">
    <el-input v-model="formData.version" :clearable="true" placeholder="请输入版本" />
</el-form-item>
        <el-form-item label="驱动类型:" prop="driverType">
    <el-select v-model="formData.driverType" placeholder="请选择驱动类型" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in driver_typeOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="状态:" prop="status">
    <el-select v-model="formData.status" placeholder="请选择状态" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in statusOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="创建时间:" prop="createdTime">
    <el-date-picker v-model="formData.createdTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="更新时间:" prop="updateTime">
    <el-date-picker v-model="formData.updateTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="协议类型:" prop="protocolType">
    <el-input v-model="formData.protocolType" :clearable="true" placeholder="请输入协议类型" />
</el-form-item>
        <el-form-item label="设备类型:" prop="deviceCategory">
    <el-input v-model="formData.deviceCategory" :clearable="true" placeholder="请输入设备类型" />
</el-form-item>
        <el-form-item label="驱动编号:" prop="driverId">
    <el-input v-model="formData.driverId" :clearable="true" placeholder="请输入驱动编号" />
</el-form-item>
        <el-form-item label="创建者:" prop="createdBy">
    <el-input v-model.number="formData.createdBy" :clearable="true" placeholder="请输入创建者" />
</el-form-item>
        <el-form-item label="更新者:" prop="updatedBy">
    <el-input v-model.number="formData.updatedBy" :clearable="true" placeholder="请输入更新者" />
</el-form-item>
        <el-form-item label="删除者:" prop="deletedBy">
    <el-input v-model.number="formData.deletedBy" :clearable="true" placeholder="请输入删除者" />
</el-form-item>
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createWlDrivers,
  updateWlDrivers,
  findWlDrivers
} from '@/api/wl_driver/wlDrivers'

defineOptions({
    name: 'WlDriversForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const driver_typeOptions = ref([])
const statusOptions = ref([])
const formData = ref({
            driverNum: '',
            driverName: '',
            version: '',
            driverType: '',
            status: '',
            createdTime: new Date(),
            updateTime: new Date(),
            protocolType: '',
            deviceCategory: '',
            driverId: '',
            createdBy: undefined,
            updatedBy: undefined,
            deletedBy: undefined,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findWlDrivers({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    driver_typeOptions.value = await getDictFunc('driver_type')
    statusOptions.value = await getDictFunc('status')
}

init()
// 保存按钮
const save = async() => {
      btnLoading.value = true
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return btnLoading.value = false
            let res
           switch (type.value) {
             case 'create':
               res = await createWlDrivers(formData.value)
               break
             case 'update':
               res = await updateWlDrivers(formData.value)
               break
             default:
               res = await createWlDrivers(formData.value)
               break
           }
           btnLoading.value = false
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
