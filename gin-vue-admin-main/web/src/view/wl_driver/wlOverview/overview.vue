<template>
  <div class="driver-overview">
    <!-- 顶部统计卡片区 -->
    <div class="top-cards">
      <el-card class="top-card" v-for="(item, idx) in topStats" :key="idx">
        <div class="top-card-title">{{ item.title }}</div>
        <div class="top-card-value" :style="{ color: item.color }">{{ item.value }}</div>
        <div class="top-card-desc">{{ item.desc }}</div>
      </el-card>
    </div>

    <!-- 设备驱动卡片区（轮播） -->
    <el-card class="card-area" shadow="never">
      <el-carousel :interval="5000" arrow="always" height="120px" indicator-position="outside">
        <el-carousel-item v-for="(item, idx) in cardList" :key="idx">
          <div class="card-item" :class="{ offline: item.status === '离线' }">
            <div class="card-title">
              <span>{{ item.title }}</span>
              <el-tag v-if="item.status === '在线'" type="success" size="small">在线</el-tag>
              <el-tag v-else type="info" size="small">离线</el-tag>
            </div>
            <div class="card-device">设备数：{{ item.deviceCount }}个</div>
            <div class="card-status">
              <span class="online">在线：{{ item.onlineCount }}个</span>
              <span class="offline">离线：{{ item.offlineCount }}个</span>
            </div>
          </div>
        </el-carousel-item>
      </el-carousel>
    </el-card>

    <!-- 消息趋势区（ECharts 折线图） -->
    <el-card class="trend-card" shadow="never">
      <div class="trend-title">消息趋势</div>
      <v-chart :option="trendOption" autoresize style="height: 180px;" />
    </el-card>

    <!-- 折线图区 -->
    <div class="chart-row">
      <el-card class="chart-card" shadow="never">
        <div class="chart-title">CPU使用率</div>
        <v-chart :option="cpuOption" autoresize style="height:220px;" />
      </el-card>
      <el-card class="chart-card" shadow="never">
        <div class="chart-title">内存占用</div>
        <v-chart :option="memOption" autoresize style="height:220px;" />
      </el-card>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { ElCard, ElTag, ElCarousel, ElCarouselItem } from 'element-plus'
import VChart from "vue-echarts"
import { use } from "echarts/core"
import { LineChart, BarChart } from "echarts/charts"
import { GridComponent, TooltipComponent, TitleComponent, LegendComponent } from "echarts/components"
import { CanvasRenderer } from "echarts/renderers"

use([LineChart, BarChart, GridComponent, TooltipComponent, TitleComponent, LegendComponent, CanvasRenderer])

// 顶部统计卡片数据
const topStats = ref([
  { title: '设备总数', value: 18, desc: '所有驱动下设备总数', color: '#409EFF' },
  { title: '在线设备', value: 12, desc: '当前在线设备数', color: '#67C23A' },
  { title: '离线设备', value: 6, desc: '当前离线设备数', color: '#909399' },
  { title: '报警数', value: 2, desc: '今日报警次数', color: '#F56C6C' },
  { title: '驱动总数', value: 5, desc: '已注册驱动数量', color: '#E6A23C' }
])

// 设备驱动卡片
const cardList = ref([
  { title: 'mqtt测试驱动2.7版本-30593142', status: '在线', deviceCount: 0, onlineCount: 0, offlineCount: 0 },
  { title: 'GB28181协议驱动-31747985', status: '离线', deviceCount: 1, onlineCount: 0, offlineCount: 1 },
  { title: 'mqtt-ca-48654311', status: '在线', deviceCount: 1, onlineCount: 0, offlineCount: 1 },
  { title: 'rtu驱动-48849713', status: '在线', deviceCount: 3, onlineCount: 2, offlineCount: 1 },
  { title: 'TCP协议驱动-53703706', status: '在线', deviceCount: 3, onlineCount: 0, offlineCount: 3 },
])

// 消息趋势折线图（数据更丰富）
const timeLabels = Array.from({ length: 20 }, (_, i) => `21:${(i * 3).toString().padStart(2, '0')}`)
const upData = [0, 1, 2, 2, 3, 4, 3, 2, 1, 0, 1, 2, 2, 3, 4, 3, 2, 1, 0, 1]
const downData = [0, 0, 1, 1, 2, 2, 1, 0, 1, 2, 2, 3, 2, 1, 0, 1, 2, 2, 1, 0]

const trendOption = {
  tooltip: { trigger: 'axis' },
  legend: { data: ['上行', '下行'], top: 0 },
  grid: { left: 40, right: 20, top: 30, bottom: 20 },
  xAxis: { type: 'category', data: timeLabels },
  yAxis: { type: 'value', name: '条/秒', minInterval: 1 },
  series: [
    { name: '上行', type: 'line', data: upData, smooth: true, symbol: 'circle', symbolSize: 6, lineStyle: { width: 2, color: '#42a5f5' }, areaStyle: { color: 'rgba(66,165,245,0.08)' } },
    { name: '下行', type: 'line', data: downData, smooth: true, symbol: 'circle', symbolSize: 6, lineStyle: { width: 2, color: '#ab47bc' }, areaStyle: { color: 'rgba(171,71,188,0.08)' } }
  ]
}

// CPU折线图
const cpuOption = {
  tooltip: { trigger: 'axis' },
  grid: { left: 40, right: 20, top: 30, bottom: 30 },
  xAxis: { type: 'category', data: timeLabels },
  yAxis: { type: 'value', name: 'CPU使用率(%)', min: 0, max: 1 },
  series: [{ data: upData.map(v => (v / 10 + Math.random() * 0.1).toFixed(2)), type: 'line', smooth: true, symbol: 'circle', symbolSize: 6, lineStyle: { width: 2, color: '#42a5f5' }, areaStyle: { color: 'rgba(66,165,245,0.08)' } }]
}

// 内存折线图
const memOption = {
  tooltip: { trigger: 'axis' },
  grid: { left: 40, right: 20, top: 30, bottom: 30 },
  xAxis: { type: 'category', data: timeLabels },
  yAxis: { type: 'value', name: '内存占用(MB)' },
  series: [{ data: downData.map(v => (15 + v * 2 + Math.random() * 2).toFixed(1)), type: 'line', smooth: true, symbol: 'circle', symbolSize: 6, lineStyle: { width: 2, color: '#ab47bc' }, areaStyle: { color: 'rgba(171,71,188,0.08)' } }]
}
</script>

<style scoped>
.driver-overview {
  padding: 24px;
  background: #f5f6fa;
}
.top-cards {
  display: flex;
  gap: 20px;
  margin-bottom: 24px;
}
.top-card {
  flex: 1;
  min-width: 160px;
  border-radius: 12px;
  box-shadow: 0 2px 8px #e3e6f0;
  text-align: center;
  padding: 18px 0 12px 0;
  border: none;
}
.top-card-title {
  font-size: 15px;
  color: #888;
  margin-bottom: 6px;
}
.top-card-value {
  font-size: 28px;
  font-weight: bold;
  margin-bottom: 4px;
}
.top-card-desc {
  font-size: 12px;
  color: #bbb;
}
.card-area {
  margin-bottom: 24px;
  background: transparent;
  box-shadow: none;
  border: none;
}
.card-item {
  min-width: 260px;
  max-width: 320px;
  background: #fff;
  border-radius: 12px;
  box-shadow: 0 4px 16px #e3e6f0;
  padding: 18px 20px 14px 20px;
  transition: box-shadow 0.2s, transform 0.2s;
  cursor: pointer;
  border: 1px solid #f0f1f2;
}
.card-item:hover {
  box-shadow: 0 8px 24px #d1d9e6;
  transform: translateY(-2px) scale(1.02);
}
.card-title {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-weight: bold;
  font-size: 15px;
  margin-bottom: 8px;
}
.card-device {
  font-size: 18px;
  font-weight: bold;
  margin-bottom: 8px;
}
.card-status {
  display: flex;
  justify-content: flex-start;
  gap: 16px;
  font-size: 13px;
}
.card-status .online {
  color: #67c23a;
}
.card-status .offline {
  color: #909399;
}
.card-item.offline {
  opacity: 0.7;
}
.trend-card {
  margin-bottom: 24px;
}
.trend-title {
  font-weight: bold;
  margin-bottom: 12px;
}
.chart-row {
  display: flex;
  gap: 16px;
}
.chart-card {
  flex: 1;
}
.chart-title {
  font-weight: bold;
  margin-bottom: 12px;
}
</style>