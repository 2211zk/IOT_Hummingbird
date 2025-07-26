
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="部门名称:" prop="departmentName">
    <el-input v-model="formData.departmentName" :clearable="true" placeholder="请输入部门名称" />
</el-form-item>
        <el-form-item label="负责人:" prop="leader">
    <el-input v-model="formData.leader" :clearable="true" placeholder="请输入负责人" />
</el-form-item>
        <el-form-item label="电话:" prop="phone">
    <el-input v-model="formData.phone" :clearable="true" placeholder="请输入电话" />
</el-form-item>
        <el-form-item label="邮箱:" prop="email">
    <el-input v-model="formData.email" :clearable="true" placeholder="请输入邮箱" />
</el-form-item>
        <el-form-item label="状态:" prop="status">
    <el-input v-model="formData.status" :clearable="true" placeholder="请输入状态" />
</el-form-item>
        <el-form-item label="排序:" prop="sort">
    <el-input v-model.number="formData.sort" :clearable="true" placeholder="请输入排序" />
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
        <el-form-item label="上级部门:" prop="parentId">
  <el-tree-select
    v-model="formData.parentId"
    :data="departmentTree"
    :props="{ label: 'departmentName', value: 'id', children: 'children' }"
    placeholder="请选择上级部门"
    clearable
    filterable
    check-strictly
    style="width: 100%"
  />
</el-form-item>
<el-form-item label="设备分配:" prop="deviceIds">
  <el-select
    v-model="formData.deviceIds"
    multiple
    filterable
    remote
    reserve-keyword
    placeholder="请选择设备"
    :remote-method="fetchDeviceList"
    :loading="deviceLoading"
    style="width: 100%"
  >
    <el-option
      v-for="item in deviceList"
      :key="item.id"
      :label="item.name + (item.productName ? '（' + item.productName + '）' : '')"
      :value="item.id"
    />
  </el-select>
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
  createWlDepartment,
  updateWlDepartment,
  findWlDepartment
} from '@/api/wl_playform/wlDepartment'
import { getWlDepartmentList } from '@/api/wl_playform/wlDepartment'
import { getDevicesByDepartment } from '@/api/wl_playform/wlDepartment'
import { ref, reactive } from 'vue'

defineOptions({
    name: 'WlDepartmentForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const formData = ref({
            departmentName: '',
            leader: '',
            phone: '',
            email: '',
            status: '',
            sort: undefined,
            createdBy: undefined,
            updatedBy: undefined,
            deletedBy: undefined,
            parentId: undefined,
            deviceIds: [],
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

const departmentTree = ref([])
const deviceList = ref([])
const deviceLoading = ref(false)

// 获取部门树
const fetchDepartmentTree = async () => {
  const res = await getWlDepartmentList({ page: 1, pageSize: 9999 })
  if (res.code === 0) {
    departmentTree.value = buildTree(res.data.list)
  }
}
// 构建树结构
function buildTree(list, parentId = null) {
  return list.filter(item => item.parentId === parentId).map(item => ({
    ...item,
    children: buildTree(list, item.id)
  }))
}
// 获取设备列表
const fetchDeviceList = async (query) => {
  deviceLoading.value = true
  // 可根据query做模糊搜索，这里简单获取全部
  // 你可根据实际接口调整
  const res = await getDevicesByDepartment({ departmentId: 0 })
  if (res.code === 0) {
    deviceList.value = res.data || []
  }
  deviceLoading.value = false
}

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findWlDepartment({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 初始化加载
fetchDepartmentTree()
fetchDeviceList()
// 保存按钮
const save = async() => {
      btnLoading.value = true
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return btnLoading.value = false
            let res
           switch (type.value) {
             case 'create':
               res = await createWlDepartment(formData.value)
               break
             case 'update':
               res = await updateWlDepartment(formData.value)
               break
             default:
               res = await createWlDepartment(formData.value)
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
