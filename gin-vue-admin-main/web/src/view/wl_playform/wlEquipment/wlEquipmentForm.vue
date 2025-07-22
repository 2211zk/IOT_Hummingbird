
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="设备名称:" prop="eqName">
    <el-input v-model="formData.eqName" :clearable="true" placeholder="请输入设备名称" />
</el-form-item>
        <el-form-item label="设备唯一标识:" prop="eqLogotype">
    <el-input v-model="formData.eqLogotype" :clearable="true" placeholder="请输入设备唯一标识" />
</el-form-item>
        <el-form-item label="所属产品:" prop="productsId">
    <el-input v-model.number="formData.productsId" :clearable="true" placeholder="请输入所属产品" />
</el-form-item>
        <el-form-item label="驱动id:" prop="driveId">
    <el-input v-model.number="formData.driveId" :clearable="true" placeholder="请输入驱动id" />
</el-form-item>
        <el-form-item label="设备坐标:" prop="eqCoordinate">
    <el-input v-model="formData.eqCoordinate" :clearable="true" placeholder="请输入设备坐标" />
</el-form-item>
        <el-form-item label="设备详细地址:" prop="eqAddress">
    <el-input v-model="formData.eqAddress" :clearable="true" placeholder="请输入设备详细地址" />
</el-form-item>
        <el-form-item label="设备描述:" prop="eqInfo">
    <el-input v-model="formData.eqInfo" :clearable="true" placeholder="请输入设备描述" />
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
  createWlEquipment,
  updateWlEquipment,
  findWlEquipment
} from '@/api/wl_playform/wlEquipment'

defineOptions({
    name: 'WlEquipmentForm'
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
const formData = ref({
            eqName: '',
            eqLogotype: '',
            productsId: undefined,
            driveId: undefined,
            eqCoordinate: '',
            eqAddress: '',
            eqInfo: '',
        })
// 验证规则
const rule = reactive({
               eqName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               productsId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findWlEquipment({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
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
               res = await createWlEquipment(formData.value)
               break
             case 'update':
               res = await updateWlEquipment(formData.value)
               break
             default:
               res = await createWlEquipment(formData.value)
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
