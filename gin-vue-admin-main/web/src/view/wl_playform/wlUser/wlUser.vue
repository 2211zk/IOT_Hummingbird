
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button  type="primary" icon="plus" @click="openDialog()">新增</el-button>
            <el-button  icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
            
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="id"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        
            <el-table-column align="left" label="主键id" prop="id" width="120" />

            <el-table-column align="left" label="用户名" prop="userName" width="120" />

            <el-table-column align="left" label="用户昵称" prop="userNickname" width="120" />

            <el-table-column align="left" label="部门" prop="department" width="120" />

            <el-table-column align="left" label="手机号" prop="mobile" width="120" />

            <el-table-column align="left" label="邮箱" prop="email" width="120" />

            <el-table-column align="left" label="密码" prop="password" width="120" />

            <el-table-column align="left" label="性别" prop="gender" width="120" />

            <el-table-column align="left" label="角色" prop="role" width="120" />

            <el-table-column align="left" label="状态" prop="userStatus" width="120" />

            <el-table-column align="left" label="备注" prop="comment" width="120" />

            <el-table-column align="left" label="创建时间" prop="creationTime" width="180">
   <template #default="scope">{{ formatDate(scope.row.creationTime) }}</template>
</el-table-column>
        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateWlUserFunc(scope.row)">编辑</el-button>
            <el-button   type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'新增':'编辑'}}</span>
                <div>
                  <el-button :loading="btnLoading" type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
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
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="主键id">
    {{ detailFrom.id }}
</el-descriptions-item>
                    <el-descriptions-item label="用户名">
    {{ detailFrom.userName }}
</el-descriptions-item>
                    <el-descriptions-item label="用户昵称">
    {{ detailFrom.userNickname }}
</el-descriptions-item>
                    <el-descriptions-item label="部门">
    {{ detailFrom.department }}
</el-descriptions-item>
                    <el-descriptions-item label="手机号">
    {{ detailFrom.mobile }}
</el-descriptions-item>
                    <el-descriptions-item label="邮箱">
    {{ detailFrom.email }}
</el-descriptions-item>
                    <el-descriptions-item label="密码">
    {{ detailFrom.password }}
</el-descriptions-item>
                    <el-descriptions-item label="性别">
    {{ detailFrom.gender }}
</el-descriptions-item>
                    <el-descriptions-item label="角色">
    {{ detailFrom.role }}
</el-descriptions-item>
                    <el-descriptions-item label="状态">
    {{ detailFrom.userStatus }}
</el-descriptions-item>
                    <el-descriptions-item label="备注">
    {{ detailFrom.comment }}
</el-descriptions-item>
                    <el-descriptions-item label="创建时间">
    {{ detailFrom.creationTime }}
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createWlUser,
  deleteWlUser,
  deleteWlUserByIds,
  updateWlUser,
  findWlUser,
  getWlUserList
} from '@/api/wl_playform/wlUser'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { useAppStore } from "@/pinia"




defineOptions({
    name: 'WlUser'
})

// 提交按钮loading
const btnLoading = ref(false)
const appStore = useAppStore()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
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
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getWlUserList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteWlUserFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.id)
        })
      const res = await deleteWlUserByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateWlUserFunc = async(row) => {
    const res = await findWlUser({ id: row.id })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteWlUserFunc = async (row) => {
    const res = await deleteWlUser({ id: row.id })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
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
        }
}
// 弹窗确定
const enterDialog = async () => {
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
                closeDialog()
                getTableData()
              }
      })
}

const detailFrom = ref({})

// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findWlUser({ id: row.id })
  if (res.code === 0) {
    detailFrom.value = res.data
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailFrom.value = {}
}


</script>

<style>

</style>
