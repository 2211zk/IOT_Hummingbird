
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="驱动名称:" prop="name">
    <el-input v-model="formData.name" :clearable="true" placeholder="请输入驱动名称" />
</el-form-item>
        <el-form-item label="图片链接:" prop="img">
    <el-input v-model="formData.img" :clearable="true" placeholder="请输入图片链接" />
</el-form-item>
        <el-form-item label="描述:" prop="description">
    <el-input v-model="formData.description" :clearable="true" placeholder="请输入描述" />
</el-form-item>
        <el-form-item label="标签:" prop="tags">
    <el-input v-model="formData.tags" :clearable="true" placeholder="请输入标签" />
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
  createDriverCards,
  updateDriverCards,
  findDriverCards
} from '@/api/wl_driver/driverCards'

defineOptions({
    name: 'DriverCardsForm'
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
            name: '',
            img: '',
            description: '',
            tags: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findDriverCards({ ID: route.query.id })
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
               res = await createDriverCards(formData.value)
               break
             case 'update':
               res = await updateDriverCards(formData.value)
               break
             default:
               res = await createDriverCards(formData.value)
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
