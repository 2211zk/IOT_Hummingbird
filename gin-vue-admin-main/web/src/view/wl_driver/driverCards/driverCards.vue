
<template>
  <div class="driver-image-page">
    <!-- 顶部介绍区块 -->
    <div class="top-desc">
      <h2>驱动镜像</h2>
      <div class="desc-text">
        驱动镜像是用户编写程序最终打包成可下载的docker镜像，它是实体设备与物联网平台通讯的桥梁，所有的设备数据由驱动上报到物联网平台，所有的设备控制指令由物联网平台下发给驱动。
      </div>
    </div>

    <!-- 查询与分类Tab -->
    <el-card class="filter-card" shadow="never">
      <div class="tab-search-row">
        <el-tabs v-model="activeTab" class="custom-tabs">
          <el-tab-pane label="驱动市场" name="market" />
          <el-tab-pane label="自定义" name="custom" />
        </el-tabs>
        <div class="filter-row">
          <el-input v-model="searchInfo.name" placeholder="名称" style="width: 200px; margin-right: 10px;" />
          <el-button type="primary" @click="onSubmit">查询</el-button>
          <el-button @click="onReset">重置</el-button>
        </div>
        <el-tabs v-model="categoryTab" class="category-tabs" style="margin-top: 10px;">
          <el-tab-pane label="官方协议" name="official" />
          <el-tab-pane label="网关" name="gateway" />
          <el-tab-pane label="摄像头" name="camera" />
          <el-tab-pane label="作业器" name="actuator" />
          <el-tab-pane label="开关" name="switch" />
          <el-tab-pane label="门禁" name="access" />
          <el-tab-pane label="探测器" name="detector" />
          <el-tab-pane label="水电表" name="meter" />
          <el-tab-pane label="检测仪" name="tester" />
        </el-tabs>
      </div>
    </el-card>

    <!-- 内容区 -->
    <div v-if="categoryTab === 'official'">
      <!-- 官方协议：图片卡片 -->
      <el-row :gutter="20" class="driver-card-list" style="margin-top: 20px;">
        <el-col :span="6" v-for="item in tableData" :key="item.ID">
          <el-card class="driver-card" shadow="hover">
            <div class="card-header">
              <img :src="item.img" class="driver-icon" v-if="item.img" />
              <div v-else class="driver-icon-placeholder">{{ item.name?.[0] }}</div>
            </div>
            <div class="driver-title">{{ item.name }}</div>
            <div class="driver-tags">
              <el-tag v-if="item.tags" size="small">{{ item.tags }}</el-tag>
            </div>
            <div class="driver-desc">{{ item.description }}</div>
            <div class="driver-actions">
              <el-button type="text" @click="getDetails(item)">手册</el-button>
              <el-button type="text">下载</el-button>
            </div>
          </el-card>
        </el-col>
      </el-row>
    </div>
    <div v-else>
      <!-- 其他tab：空数据提示 -->
      <el-empty description="暂无数据" />
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { getDriverCardsList, findDriverCards } from '@/api/wl_driver/driverCards'
import { ElMessage } from 'element-plus'

const activeTab = ref('market')
const categoryTab = ref('official')
const searchInfo = ref({})
const tableData = ref([])

const onSubmit = () => {
  getTableData()
}
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

const getTableData = async () => {
  const table = await getDriverCardsList({ ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
  }
}
getTableData()

const getDetails = async (item) => {
  const res = await findDriverCards({ ID: item.ID })
  if (res.code === 0) {
    ElMessage.info('手册功能待实现')
  }
}
</script>

<style scoped>
.driver-image-page { padding: 24px; }
.top-desc { margin-bottom: 16px; }
.driver-card-list { margin-top: 24px; }
.driver-card { min-height: 220px; display: flex; flex-direction: column; align-items: center; }
.card-header { margin-bottom: 8px; }
.driver-icon { width: 64px; height: 64px; }
.driver-icon-placeholder { width: 64px; height: 64px; background: #eee; display: flex; align-items: center; justify-content: center; font-size: 24px; }
.driver-title { font-weight: bold; margin-bottom: 8px; }
.driver-tags { margin-bottom: 8px; }
.driver-desc { color: #888; font-size: 13px; margin-bottom: 8px; text-align: center; }
.driver-actions { display: flex; gap: 8px; }
</style>
