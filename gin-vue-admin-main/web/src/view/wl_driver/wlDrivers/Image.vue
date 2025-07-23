<template>
  <div class="driver-image-page">
    <!-- 顶部介绍区块 -->
    <div class="top-desc">
      <h2>驱动镜像</h2>
      <div class="desc-text">
        驱动镜像是用户编写驱动程序经打包成可下载的docker镜像，它是实现设备与物联网平台通讯的桥梁，所有的设备数据由驱动上报到物联网平台，所有的设备控制指令由物联网平台下发给驱动。
      </div>
    </div>

    <!-- 分类Tab和查询区 -->
    <el-card class="filter-card" shadow="never">
      <div class="tab-search-row">
        <el-tabs v-model="activeTab" class="custom-tabs">
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
        <div class="filter-row">
          <el-input v-model="search" placeholder="名称" style="width: 200px; margin-right: 10px;" />
          <el-button type="primary" @click="onSearch">查询</el-button>
          <el-button @click="onReset">重置</el-button>
        </div>
      </div>
    </el-card>

    <!-- 镜像卡片列表 -->
    <el-row :gutter="32" class="card-list">
      <el-col :span="6" v-for="item in filteredList" :key="item.id" class="card-col">
        <el-card class="image-card" shadow="never">
          <div class="card-content">
            <div class="card-header">
              <img :src="item.icon" class="card-icon" v-if="item.icon" />
              <div v-else class="card-icon-placeholder">{{ item.name.split(' ')[0] }}</div>
            </div>
            <div class="card-labels">
              <el-tag v-if="item.downloaded" class="custom-tag info">已下载</el-tag>
              <el-tag v-else class="custom-tag warning">未下载</el-tag>
              <el-tag v-if="item.pay" class="custom-tag pay">付费</el-tag>
              <el-tag class="custom-tag version">{{ item.version }}版本</el-tag>
              <el-tag v-if="item.open" class="custom-tag open">开源</el-tag>
            </div>
            <div class="card-title">{{ item.name }}</div>
          </div>
          <div class="card-actions">
            <el-button type="text" size="small" class="manual-btn">手册</el-button>
            <el-button v-if="!item.downloaded" type="primary" size="small" class="download-btn">下载</el-button>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'

const activeTab = ref('official')
const search = ref('')

// 静态镜像数据
const imageList = ref([
  {
    id: 1,
    icon: 'https://edge-community-1257991367.cos.ap-shanghai.myqcloud.com/Ireland/driver-image/modbus.png',
    name: 'MODBUS RTU协议驱动',
    desc: '',
    downloaded: false,
    pay: true,
    version: '2.7',
    open: false
  },
  {
    id: 2,
    icon: 'https://edge-community-1257991367.cos.ap-shanghai.myqcloud.com/Ireland/driver-image/gb-28181.png',
    name: 'GB28181协议驱动',
    desc: '',
    downloaded: true,
    pay: true,
    version: '2.7',
    open: false
  },
  {
    id: 3,
    icon: 'https://edge-community-1257991367.cos.ap-shanghai.myqcloud.com/Ireland/driver-image/modbus.png',
    name: 'MODBUS TCP协议驱动',
    desc: '',
    downloaded: true,
    pay: true,
    version: '2.7',
    open: false
  },
  {
    id: 4,
    icon: 'https://edge-community-1257991367.cos.ap-shanghai.myqcloud.com/Ireland/driver-image/tcp.png',
    name: 'TCP协议驱动',
    desc: '',
    downloaded: true,
    pay: true,
    version: '2.7',
    open: true
  },
  {
    id: 5,
    icon: 'https://edge-community-1257991367.cos.ap-shanghai.myqcloud.com/Ireland/driver-image/http.png',
    name: 'HTTP协议驱动',
    desc: '',
    downloaded: true,
    pay: true,
    version: '2.7',
    open: true
  },
  {
    id: 6,
    icon: 'https://edge-community-1257991367.cos.ap-shanghai.myqcloud.com/Ireland/driver-image/mqtt.png',
    name: 'MQTT协议驱动',
    desc: '',
    downloaded: true,
    pay: true,
    version: '2.7',
    open: true
  }
])

const filteredList = computed(() => {
  if (!search.value) return imageList.value
  return imageList.value.filter(item => item.name.includes(search.value))
})

function onSearch() {
  // 这里只做本地过滤，后续可对接API
}
function onReset() {
  search.value = ''
}
</script>

<style scoped>
.driver-image-page {
  padding: 20px;
  background: #f7f8fa;
  min-height: 100vh;
}
.top-desc {
  margin-bottom: 20px;
}
.top-desc h2 {
  margin-bottom: 8px;
}
.desc-text {
  color: #666;
  font-size: 15px;
}
.filter-card {
  margin-bottom: 20px;
  border: 1px solid #f0f0f0 !important;
  box-shadow: none !important;
}
.tab-search-row {
  display: flex;
  align-items: flex-end;
  justify-content: space-between;
}
.custom-tabs {
  flex: 1;
}
.filter-row {
  margin-left: 20px;
  margin-bottom: 0;
  display: flex;
  align-items: center;
}
.card-list {
  margin-top: 10px;
}
.card-col {
  display: flex;
  justify-content: center;
}
.image-card {
  width: 100%;
  min-height: 260px;
  background: #fff;
  border: 1px solid #eaeaea;
  box-shadow: none;
  border-radius: 10px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: space-between;
  padding: 24px 0 12px 0;
  transition: box-shadow 0.2s;
}
.image-card:hover {
  box-shadow: 0 2px 8px 0 rgba(0,0,0,0.04);
  border-color: #d3d3d3;
}
.card-content {
  width: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
}
.card-header {
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 10px;
}
.card-icon {
  width: 64px;
  height: 64px;
  margin-bottom: 8px;
}
.card-icon-placeholder {
  width: 64px;
  height: 64px;
  background: #f0f0f0;
  color: #bbb;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 24px;
  border-radius: 50%;
  margin-bottom: 8px;
}
.card-labels {
  display: flex;
  gap: 6px;
  flex-wrap: wrap;
  justify-content: center;
  margin-bottom: 10px;
}
.custom-tag {
  border-radius: 12px !important;
  font-size: 12px !important;
  padding: 0 8px !important;
  height: 22px !important;
  line-height: 20px !important;
  background: #f6faff !important;
  color: #409eff !important;
  border: 1px solid #e0eaff !important;
}
.custom-tag.info {
  background: #f6faff !important;
  color: #409eff !important;
  border: 1px solid #e0eaff !important;
}
.custom-tag.warning {
  background: #fffbe6 !important;
  color: #faad14 !important;
  border: 1px solid #ffe58f !important;
}
.custom-tag.pay {
  background: #fff1f0 !important;
  color: #f56c6c !important;
  border: 1px solid #ffccc7 !important;
}
.custom-tag.version {
  background: #f6ffed !important;
  color: #52c41a !important;
  border: 1px solid #b7eb8f !important;
}
.custom-tag.open {
  background: #e6fffb !important;
  color: #13c2c2 !important;
  border: 1px solid #87e8de !important;
}
.card-title {
  font-weight: bold;
  font-size: 17px;
  margin-bottom: 6px;
  text-align: center;
  color: #222;
}
.card-actions {
  width: 100%;
  display: flex;
  justify-content: space-between;
  align-items: center;
  border-top: 1px solid #f0f0f0;
  padding-top: 10px;
  margin-top: 10px;
}
.manual-btn {
  color: #409eff !important;
  font-size: 14px !important;
  padding-left: 20px;
}
.download-btn {
  margin-right: 20px;
  font-size: 14px !important;
}
</style> 