
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="规则名称:" prop="ruleName">
    <el-input v-model="formData.ruleName" :clearable="true" placeholder="请输入规则名称" />
</el-form-item>
        <el-form-item label="规则描述:" prop="ruleDescription">
    <el-input v-model="formData.ruleDescription" :clearable="true" placeholder="请输入规则描述" />
</el-form-item>
        <el-form-item label="消息源:" prop="messageSource">
    <el-select v-model="formData.messageSource" placeholder="请选择消息源" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in informationOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="查询字段:" prop="queryField">
    <el-input v-model="formData.queryField" :clearable="true" placeholder="请输入查询字段" />
</el-form-item>
        <el-form-item label="条件:" prop="condition">
    <el-input v-model="formData.condition" :clearable="true" placeholder="请输入条件" />
</el-form-item>
        <el-form-item label="sql语句:" prop="sqlStatement">
    <el-input v-model="formData.sqlStatement" :clearable="true" placeholder="请输入sql语句" />
</el-form-item>
        <el-form-item label="转换方法:" prop="forwardingMethod">
    <el-input v-model="formData.forwardingMethod" :clearable="true" placeholder="请输入转换方法" />
</el-form-item>
        <el-form-item label="使用资源id:" prop="resourceId">
    <el-input v-model.number="formData.resourceId" :clearable="true" placeholder="请输入使用资源id" />
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
  createWlEngineRules,
  updateWlEngineRules,
  findWlEngineRules
} from '@/api/wl_playform/wlEngineRules'

defineOptions({
    name: 'WlEngineRulesForm'
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
const informationOptions = ref([])
const formData = ref({
            ruleName: '',
            ruleDescription: '',
            messageSource: '',
            queryField: '',
            condition: '',
            sqlStatement: '',
            forwardingMethod: '',
            resourceId: undefined,
        })
// 验证规则
const rule = reactive({
               ruleName : [{
                   required: true,
                   message: '请填写规则名称',
                   trigger: ['input','blur'],
               }],
               messageSource : [{
                   required: true,
                   message: '请填写消息源',
                   trigger: ['input','blur'],
               }],
               queryField : [{
                   required: true,
                   message: '请填写查询字段',
                   trigger: ['input','blur'],
               }],
               forwardingMethod : [{
                   required: true,
                   message: '请选择转换方法',
                   trigger: ['input','blur'],
               }],
               resourceId : [{
                   required: true,
                   message: '请选择资源',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findWlEngineRules({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    informationOptions.value = await getDictFunc('information')
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
               res = await createWlEngineRules(formData.value)
               break
             case 'update':
               res = await updateWlEngineRules(formData.value)
               break
             default:
               res = await createWlEngineRules(formData.value)
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
