
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="设备ID:" prop="deviceId">
    <el-input v-model.number="formData.deviceId" :clearable="true" placeholder="请输入设备ID" />
</el-form-item>
        <el-form-item label="告警类型:" prop="alarmType">
    <el-input v-model="formData.alarmType" :clearable="true" placeholder="请输入告警类型" />
</el-form-item>
        <el-form-item label="告警级别:" prop="alarmLevel">
    <el-select v-model="formData.alarmLevel" placeholder="请选择告警级别" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in alarm_levelOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="告警状态:" prop="alarmStatus">
    <el-select v-model="formData.alarmStatus" placeholder="请选择告警状态" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in alarm_statusOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="告警内容描述:" prop="alarmContent">
    <el-input v-model="formData.alarmContent" :clearable="true" placeholder="请输入告警内容描述" />
</el-form-item>
        <el-form-item label="告警相关数据:" prop="alarmData">
    <el-input v-model="formData.alarmData" :clearable="true" placeholder="请输入告警相关数据" />
</el-form-item>
        <el-form-item label="告警创建时间:" prop="createTime">
    <el-date-picker v-model="formData.createTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="告警更新时间:" prop="updateTime">
    <el-date-picker v-model="formData.updateTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="处理时间:" prop="handleTime">
    <el-date-picker v-model="formData.handleTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="处理人:" prop="handleUser">
    <el-input v-model="formData.handleUser" :clearable="true" placeholder="请输入处理人" />
</el-form-item>
        <el-form-item label="处理备注:" prop="handleRemark">
    <el-input v-model="formData.handleRemark" :clearable="true" placeholder="请输入处理备注" />
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
  createWlAlarm,
  updateWlAlarm,
  findWlAlarm
} from '@/api/wl_playform/wlAlarm'

defineOptions({
    name: 'WlAlarmForm'
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
const alarm_statusOptions = ref([])
const alarm_levelOptions = ref([])
const formData = ref({
            deviceId: undefined,
            alarmType: '',
            alarmLevel: '',
            alarmStatus: '',
            alarmContent: '',
            alarmData: '',
            createTime: new Date(),
            updateTime: new Date(),
            handleTime: new Date(),
            handleUser: '',
            handleRemark: '',
            createdBy: undefined,
            updatedBy: undefined,
            deletedBy: undefined,
        })
// 验证规则
const rule = reactive({
               deviceId : [{
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
      const res = await findWlAlarm({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    alarm_statusOptions.value = await getDictFunc('alarm_status')
    alarm_levelOptions.value = await getDictFunc('alarm_level')
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
               res = await createWlAlarm(formData.value)
               break
             case 'update':
               res = await updateWlAlarm(formData.value)
               break
             default:
               res = await createWlAlarm(formData.value)
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
