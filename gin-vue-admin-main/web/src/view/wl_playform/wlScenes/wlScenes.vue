
<template>
  <div class="scene-management-page">
    <!-- 页面头部 -->
    <div class="page-header">
      <h2>场景联动</h2>
      <p class="description">
        平台提供多种场景联动类型，支持不同业务场景的自动化流程配置，通过创建场景快速实现设备联动和业务自动化。
      </p>
    </div>

    <!-- 表格区域 -->
    <div class="scene-table">
      <div class="gva-btn-list">
        <el-button v-auth="btnAuth.add" type="primary" icon="plus" @click="openDialog()">新增</el-button>
        <el-button v-auth="btnAuth.batchDelete" icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
        <ExportTemplate v-auth="btnAuth.exportTemplate" template-id="wl_playform_WlScenes" />
        <ExportExcel v-auth="btnAuth.exportExcel" template-id="wl_playform_WlScenes" filterDeleted/>
        <ImportExcel v-auth="btnAuth.importExcel" template-id="wl_playform_WlScenes" @on-success="getTableData" />
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        
        <el-table-column sortable align="left" label="日期" prop="CreatedAt" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
        <el-table-column align="left" label="场景名称" prop="sceneName" width="120" />

        <el-table-column align="left" label="场景描述" prop="scenesDescription" width="120" />

        <el-table-column align="left" label="启动状态" prop="scenesStatus" width="120">
          <template #default="scope">
            <el-switch
              v-model="scope.row.scenesStatus"
              :active-value="'1'"
              :inactive-value="'0'"
              active-text="启用"
              inactive-text="禁用"
              @change="handleStatusChange(scope.row)"
              :loading="scope.row.statusLoading"
            />
          </template>
        </el-table-column>
        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
          <template #default="scope">
            <el-button v-auth="btnAuth.info" type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button v-auth="btnAuth.edit" type="primary" link icon="edit" class="table-button" @click="updateWlScenesFunc(scope.row)">编辑</el-button>
            <el-button v-auth="btnAuth.delete" type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
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
        <el-form-item label="场景名称:" prop="sceneName">
          <el-input v-model="formData.sceneName" :clearable="true" placeholder="请输入场景名称" />
        </el-form-item>
        <el-form-item label="场景描述:" prop="scenesDescription">
          <el-input 
            v-model="formData.scenesDescription" 
            type="textarea" 
            :rows="4"
            :clearable="true" 
            placeholder="请输入场景描述" 
            resize="vertical"
            maxlength="500"
            show-word-limit
          />
        </el-form-item>
      </el-form>
    </el-drawer>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
      <el-descriptions :column="1" border>
        <el-descriptions-item label="场景名称">
          {{ detailFrom.sceneName }}
        </el-descriptions-item>
        <el-descriptions-item label="场景描述">
          {{ detailFrom.scenesDescription }}
        </el-descriptions-item>
        <el-descriptions-item label="启动状态">
          {{ detailFrom.scenesStatus }}
        </el-descriptions-item>
      </el-descriptions>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createWlScenes,
  deleteWlScenes,
  deleteWlScenesByIds,
  updateWlScenes,
  findWlScenes,
  getWlScenesList
} from '@/api/wl_playform/wlScenes'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
// 引入按钮权限标识
import { useBtnAuth } from '@/utils/btnAuth'
import { useAppStore } from "@/pinia"

// 导出组件
import ExportExcel from '@/components/exportExcel/exportExcel.vue'
// 导入组件
import ImportExcel from '@/components/exportExcel/importExcel.vue'
// 导出模板组件
import ExportTemplate from '@/components/exportExcel/exportTemplate.vue'


defineOptions({
    name: 'WlScenes'
})
// 按钮权限实例化
    const btnAuth = useBtnAuth()

// 提交按钮loading
const btnLoading = ref(false)
const appStore = useAppStore()

// 自动化生成的字典（可能为空）以及字段
const statusOptions = ref([])
const formData = ref({
            sceneName: '',
            scenesDescription: '',
            scenesStatus: '',
        })



// 验证规则
const rule = reactive({
               sceneName : [{
                   required: true,
                   message: '请填写场景名称',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
})

const elFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])

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
  const table = await getWlScenesList({ page: page.value, pageSize: pageSize.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
    
    // 数据加载完成后，设置初始化标志为false
    setTimeout(() => {
      isInitializing.value = false
    }, 100)
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
    statusOptions.value = await getDictFunc('status')
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
            deleteWlScenesFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          IDs.push(item.ID)
        })
      const res = await deleteWlScenesByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateWlScenesFunc = async(row) => {
    const res = await findWlScenes({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteWlScenesFunc = async (row) => {
    const res = await deleteWlScenes({ ID: row.ID })
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
        sceneName: '',
        scenesDescription: '',
        scenesStatus: '',
        }
}
// 初始化标志
const isInitializing = ref(true)

// 状态切换处理
const handleStatusChange = async (row) => {
  // 如果是初始化阶段，不显示提示
  if (isInitializing.value) {
    return
  }
  
  // 设置loading状态
  row.statusLoading = true
  
  try {
    const res = await updateWlScenes({
      ID: row.ID,
      sceneName: row.sceneName,
      scenesDescription: row.scenesDescription,
      scenesStatus: row.scenesStatus
    })
    
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: row.scenesStatus === '1' ? '启用成功' : '禁用成功'
      })
    } else {
      // 如果更新失败，恢复原状态
      row.scenesStatus = row.scenesStatus === '1' ? '0' : '1'
      ElMessage({
        type: 'error',
        message: '状态更新失败'
      })
    }
  } catch (error) {
    // 如果出错，恢复原状态
    row.scenesStatus = row.scenesStatus === '1' ? '0' : '1'
    ElMessage({
      type: 'error',
      message: '状态更新失败'
    })
  } finally {
    row.statusLoading = false
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
                  res = await createWlScenes(formData.value)
                  break
                case 'update':
                  res = await updateWlScenes(formData.value)
                  break
                default:
                  res = await createWlScenes(formData.value)
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
  const res = await findWlScenes({ ID: row.ID })
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

<style scoped>
.scene-management-page {
  background: #0f172a;
  min-height: 100vh;
  padding: 20px;
  color: #fff;
}

.page-header {
  margin-bottom: 20px;
  background: #0f172a;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.4);
  border: 1px solid #334155;
}

.page-header h2 {
  margin: 0 0 10px 0;
  font-size: 24px;
  font-weight: 600;
  color: #fff;
}

.description {
  margin: 0;
  color: #cbd5e1;
  line-height: 1.6;
}

.scene-table {
  background: #0f172a;
  border-radius: 8px;
  padding: 20px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.4);
  border: 1px solid #334155;
}

/* 表格样式 */
.scene-table :deep(.el-table__header-wrapper) {
  background: #0f172a;
}

.scene-table :deep(.el-table__header) {
  background: #0f172a;
}

.scene-table :deep(.el-table__header th) {
  background: #0f172a !important;
  color: #fff !important;
  border-bottom: 1px solid #334155;
}

.scene-table :deep(.el-table__header th .cell) {
  color: #fff !important;
  font-weight: 600;
}

.scene-table :deep(.el-table__body tr) {
  background: #0f172a;
  color: #fff;
}

.scene-table :deep(.el-table__body tr:hover) {
  background: #1e293b;
}

.scene-table :deep(.el-table__body td) {
  background: #0f172a;
  color: #fff;
  border-bottom: 1px solid #334155;
}

.scene-table :deep(.el-table) {
  border: 1px solid #334155;
  border-radius: 8px;
  overflow: hidden;
  background: #0f172a;
}

.scene-table :deep(.el-table__border-line) {
  background: #334155;
}

.scene-table :deep(.el-table__body-wrapper) {
  background: #0f172a;
}

.scene-table :deep(.el-table__fixed-header-wrapper) {
  background: #0f172a;
}

.scene-table :deep(.el-table__fixed-body-wrapper) {
  background: #0f172a;
}

.scene-table :deep(.el-table__empty-block) {
  background: #0f172a;
}

.scene-table :deep(.el-table__empty-text) {
  color: #cbd5e1;
}

/* 分页样式 */
.gva-pagination {
  margin-top: 20px;
  text-align: right;
}

.gva-pagination :deep(.el-pagination) {
  color: #fff;
}

.gva-pagination :deep(.el-pagination .el-pager li) {
  background: #0f172a;
  color: #fff;
  border: 1px solid #334155;
}

.gva-pagination :deep(.el-pagination .el-pager li.is-active) {
  background: #409eff;
  color: #fff;
}

.gva-pagination :deep(.el-pagination .btn-prev),
.gva-pagination :deep(.el-pagination .btn-next) {
  background: #0f172a;
  color: #fff;
  border: 1px solid #334155;
}

/* 抽屉样式调整 */
:deep(.el-drawer) {
  background: #0f172a;
  color: #fff;
}

:deep(.el-drawer__header) {
  background: #0f172a;
  color: #fff;
  border-bottom: 1px solid #334155;
}

:deep(.el-drawer__body) {
  background: #0f172a;
  color: #fff;
}

/* 表单样式调整 */
:deep(.el-form-item__label) {
  color: #fff;
}

:deep(.el-input__wrapper) {
  background: #334155;
  border: 1px solid #475569;
}

:deep(.el-input__inner) {
  background: #334155;
  color: #fff;
  border: none;
}

:deep(.el-select .el-input__wrapper) {
  background: #334155;
  border: 1px solid #475569;
}

:deep(.el-select-dropdown) {
  background: #0f172a;
  border: 1px solid #334155;
}

:deep(.el-select-dropdown__item) {
  color: #fff;
}

:deep(.el-select-dropdown__item:hover) {
  background: #334155;
}

:deep(.el-select-dropdown__item.selected) {
  background: #409eff;
  color: #fff;
}

/* 开关样式 */
:deep(.el-switch) {
  --el-switch-on-color: #409eff;
  --el-switch-off-color: #606266;
}

:deep(.el-switch__core) {
  border-color: #475569;
}

:deep(.el-switch.is-checked .el-switch__core) {
  border-color: #409eff;
}

:deep(.el-switch__label) {
  color: #fff;
}

:deep(.el-switch__label.is-active) {
  color: #409eff;
}
</style>
