
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
        <div class="category-tabs-custom">
          <span
            v-for="cat in categories"
            :key="cat.value"
            :class="['category-tab-item', { active: categoryTab === cat.value }]"
            @click="categoryTab = cat.value"
          >
            {{ cat.label }}
          </span>
        </div>
      </div>
    </el-card>

    <!-- 卡片区 -->
    <el-row :gutter="20" class="driver-card-list" style="margin-top: 20px;">
      <el-col :span="6" v-for="item in tableData" :key="item.ID">
        <el-card class="driver-card" shadow="hover">
          <div class="card-header">
            <img :src="item.img" class="driver-icon" v-if="item.img" />
            <div v-else class="driver-icon-placeholder">{{ item.name?.[0] }}</div>
          </div>
          <div class="driver-title">{{ item.name }}</div>
          <div class="driver-tags">
            <el-tag v-if="item.downloaded === true" size="small" type="info">已下载</el-tag>
            <el-tag v-else size="small" type="info">未下载</el-tag>
            <el-tag v-if="item.pay === true" size="small" type="warning">付费</el-tag>
            <el-tag v-if="item.openSource === true" size="small" type="success">开源</el-tag>
            <el-tag v-if="item.version" size="small" type="success">{{ item.version }}版本</el-tag>
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
</template>

<script setup>
import { ref, watch } from 'vue'
import { getDriverCardsList, findDriverCards } from '@/api/wl_driver/driverCards'
import { ElMessage } from 'element-plus'

const categories = [
  { label: '官方协议', value: 'official' },
  { label: '网关', value: 'gateway' },
  { label: '摄像头', value: 'camera' },
  { label: '作业器', value: 'actuator' },
  { label: '开关', value: 'switch' },
  { label: '门禁', value: 'access' },
  { label: '探测器', value: 'detector' },
  { label: '水电表', value: 'meter' },
  { label: '检测仪', value: 'tester' }
]
const activeTab = ref('market')
const categoryTab = ref('official')
const searchInfo = ref({ name: '', tags: '', description: '' })
const tableData = ref([])

const onSubmit = () => {
  getTableData()
}
const onReset = () => {
  searchInfo.value = { name: '', tags: '', description: '' }
  getTableData()
}

const getTableData = async () => {
  const params = { ...searchInfo.value }
  // 联动类型筛选
  if (categoryTab.value) {
    params.driverType = categoryTab.value
  }
  const res = await getDriverCardsList(params)
  if (res.code === 0) {
    tableData.value = res.data.list
  }
}

// 监听类型切换自动筛选
watch(categoryTab, () => {
  getTableData()
})

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
.category-tabs-custom {
  display: flex;
  gap: 32px;
  border-bottom: 2px solid #f0f0f0;
  margin: 10px 0 0 0;
  padding-left: 4px;
}
.category-tab-item {
  font-size: 16px;
  color: #222;
  padding: 6px 8px 8px 8px;
  cursor: pointer;
  position: relative;
  transition: color 0.2s;
}
.category-tab-item.active {
  color: #409EFF;
  font-weight: bold;
}
.category-tab-item.active::after {
  content: '';
  display: block;
  position: absolute;
  left: 0;
  right: 0;
  bottom: -2px;
  height: 3px;
  background: #409EFF;
  border-radius: 2px 2px 0 0;
}
</style>
