
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="产品名称:" prop="prName">
    <el-input v-model="formData.prName" :clearable="true" placeholder="请输入产品名称" />
</el-form-item>
        <el-form-item label="所属品类:" prop="prCategory">
    <el-select v-model="formData.prCategory" placeholder="请选择所属品类" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in CategoryOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="标准品类:" prop="standardQuality">
    <el-input v-model.number="formData.standardQuality" :clearable="true" placeholder="请输入标准品类" />
</el-form-item>
        <el-form-item label="节点类型:" prop="nodeType">
    <el-select v-model="formData.nodeType" placeholder="请选择节点类型" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in Node_typeOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="接入协议:" prop="accessProtocol">
    <el-select v-model="formData.accessProtocol" placeholder="请选择接入协议" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in access_protocolOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="数据格式:" prop="dataFormat">
    <el-select v-model="formData.dataFormat" placeholder="请选择数据格式" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in data_formatOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="网络类型:" prop="networkType">
    <el-select v-model="formData.networkType" placeholder="请选择网络类型" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in Network_typeOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="工厂:" prop="factory">
    <el-input v-model="formData.factory" :clearable="true" placeholder="请输入工厂" />
</el-form-item>
        <el-form-item label="产品描述:" prop="prInfo">
    <el-input v-model="formData.prInfo" :clearable="true" placeholder="请输入产品描述" />
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
  createWlProducts,
  updateWlProducts,
  findWlProducts
} from '@/api/wl_playform/wlProducts'

defineOptions({
    name: 'WlProductsForm'
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
const Node_typeOptions = ref([])
const access_protocolOptions = ref([])
const Network_typeOptions = ref([])
const CategoryOptions = ref([])
const data_formatOptions = ref([])
const formData = ref({
            prName: '',
            prCategory: '',
            standardQuality: undefined,
            nodeType: '',
            accessProtocol: '',
            dataFormat: '',
            networkType: '',
            factory: '',
            prInfo: '',
        })
// 验证规则
const rule = reactive({
               prName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               prCategory : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               standardQuality : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               nodeType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               accessProtocol : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               dataFormat : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               networkType : [{
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
      const res = await findWlProducts({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    Node_typeOptions.value = await getDictFunc('Node_type')
    access_protocolOptions.value = await getDictFunc('access_protocol')
    Network_typeOptions.value = await getDictFunc('Network_type')
    CategoryOptions.value = await getDictFunc('Category')
    data_formatOptions.value = await getDictFunc('data_format')
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
               res = await createWlProducts(formData.value)
               break
             case 'update':
               res = await updateWlProducts(formData.value)
               break
             default:
               res = await createWlProducts(formData.value)
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
