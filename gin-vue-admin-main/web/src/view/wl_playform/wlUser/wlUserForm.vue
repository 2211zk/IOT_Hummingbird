
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="主键id:" prop="id">
    <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入主键id" />
</el-form-item>
        <el-form-item label="用户名:" prop="userName">
    <el-input v-model="formData.userName" :clearable="true" placeholder="请输入用户名" />
</el-form-item>
        <el-form-item label="用户昵称:" prop="userNickname">
    <el-input v-model="formData.userNickname" :clearable="true" placeholder="请输入用户昵称" />
</el-form-item>
        <el-form-item label="部门:" prop="department">
    <el-input v-model.number="formData.department" :clearable="true" placeholder="请输入部门" />
</el-form-item>
        <el-form-item label="手机号:" prop="mobile">
    <el-input v-model="formData.mobile" :clearable="true" placeholder="请输入手机号" />
</el-form-item>
        <el-form-item label="邮箱:" prop="email">
    <el-input v-model="formData.email" :clearable="true" placeholder="请输入邮箱" />
</el-form-item>
        <el-form-item label="密码:" prop="password">
    <el-input v-model="formData.password" :clearable="true" placeholder="请输入密码" />
</el-form-item>
        <el-form-item label="性别:" prop="gender">
    <el-input v-model="formData.gender" :clearable="true" placeholder="请输入性别" />
</el-form-item>
        <el-form-item label="角色:" prop="role">
    <el-input v-model.number="formData.role" :clearable="true" placeholder="请输入角色" />
</el-form-item>
        <el-form-item label="状态:" prop="userStatus">
    <el-input v-model="formData.userStatus" :clearable="true" placeholder="请输入状态" />
</el-form-item>
        <el-form-item label="备注:" prop="comment">
    <el-input v-model="formData.comment" :clearable="true" placeholder="请输入备注" />
</el-form-item>
        <el-form-item label="创建时间:" prop="creationTime">
    <el-date-picker v-model="formData.creationTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
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
  createWlUser,
  updateWlUser,
  findWlUser
} from '@/api/wl_playform/wlUser'

defineOptions({
    name: 'WlUserForm'
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
            id: undefined,
            userName: '',
            userNickname: '',
            department: undefined,
            mobile: '',
            email: '',
            password: '',
            gender: '',
            role: undefined,
            userStatus: '',
            comment: '',
            creationTime: new Date(),
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findWlUser({ ID: route.query.id })
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
               res = await createWlUser(formData.value)
               break
             case 'update':
               res = await updateWlUser(formData.value)
               break
             default:
               res = await createWlUser(formData.value)
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
