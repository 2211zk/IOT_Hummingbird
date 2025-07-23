<template>
  <div class="overview-container">
    <!-- 面包屑导航 -->
    <div class="breadcrumb">
      <span>首页 / 驱动管理 / 驱动概览</span>
    </div>
    <!-- 数据概览部分 -->
    <div class="data-overview">
      <h3>数据概览</h3>
      <div class="driver-cards">
        <div class="card-container">
          <div
            v-for="(driver, index) in driverData"
            :key="index"
            :class="['driver-card', { active: index === activeCard }]"
          >
            <div class="card-header">
              <span class="driver-name">{{ driver.name }}</span>
              <span :class="['status-tag', driver.status]">{{ driver.status === 'online' ? '在线' : '离线' }}</span>
            </div>
            <div class="card-content">
              <div class="device-count">设备数: {{ driver.total }}个</div>
              <div class="status-details">
                <span>在线: {{ driver.online }}个</span>
                <span>离线: {{ driver.offline }}个</span>
              </div>
            </div>
          </div>
        </div>
        <div class="card-navigation">
          <div class="nav-dots">
            <span
              v-for="i in navDots"
              :key="i"
              :class="['dot', { active: i - 1 === activeCard }]"
            ></span>
          </div>
        </div>
      </div>
    </div>
    <!-- 消息趋势部分 -->
    <div class="message-trend">
      <h3>消息趋势</h3>
      <div class="trend-charts-vertical">
        <div class="chart-container">
          <div class="chart-title uplink-title">消息上行0条/秒</div>
          <div class="chart-area" ref="uplinkChart"></div>
        </div>
        <div class="chart-container">
          <div class="chart-title downlink-title">消息下行0条/秒</div>
          <div class="chart-area" ref="downlinkChart"></div>
        </div>
      </div>
    </div>
    <!-- 资源使用率部分 -->
    <div class="resource-usage">
      <div class="resource-card">
        <div class="resource-title">CPU使用率</div>
        <div class="resource-content" ref="cpuChart"></div>
      </div>
      <div class="resource-card">
        <div class="resource-title">内存占用</div>
        <div class="resource-content" ref="memoryChart"></div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import * as echarts from 'echarts'

const uplinkChart = ref(null)
const downlinkChart = ref(null)
const cpuChart = ref(null)
const memoryChart = ref(null)

const driverData = [
  { name: 'mqtt测试驱动2.7版本', total: 0, online: 0, offline: 0, status: 'offline' },
  { name: 'GB28181协议驱动', total: 0, online: 0, offline: 0, status: 'offline' },
  { name: 'mqtt-ca', total: 1, online: 1, offline: 0, status: 'online' },
  { name: 'rtu驱动', total: 3, online: 2, offline: 1, status: 'online' }
]
const activeCard = 2 // 第3个卡片高亮
const navDots = 4 // 导航点数量

onMounted(() => {
  // 生成60个时间点
  const xData = Array.from({ length: 60 }, (_, i) => i + 1)
  const yDataUplink = Array(60).fill(1)
  const yDataDownlink = Array(60).fill(1)

  // 消息上行图表
  const uplink = echarts.init(uplinkChart.value)
  uplink.setOption({
    grid: { top: 0, right: 0, bottom: 0, left: 0 },
    xAxis: {
      type: 'category',
      data: xData,
      show: false,
      axisLine: { show: false },
      axisTick: { show: false },
      splitLine: { show: false }
    },
    yAxis: {
      type: 'value',
      show: false,
      axisLine: { show: false },
      axisTick: { show: false },
      splitLine: { show: false },
      min: 0,
      max: 2
    },
    series: [{
      data: yDataUplink,
      type: 'bar',
      itemStyle: { color: '#e3f1fd' },
      barWidth: 6,
      barGap: '2%',
      emphasis: { disabled: true }
    }],
    animation: false
  })

  // 消息下行图表
  const downlink = echarts.init(downlinkChart.value)
  downlink.setOption({
    grid: { top: 0, right: 0, bottom: 0, left: 0 },
    xAxis: {
      type: 'category',
      data: xData,
      show: false,
      axisLine: { show: false },
      axisTick: { show: false },
      splitLine: { show: false }
    },
    yAxis: {
      type: 'value',
      show: false,
      axisLine: { show: false },
      axisTick: { show: false },
      splitLine: { show: false },
      min: 0,
      max: 2
    },
    series: [{
      data: yDataDownlink,
      type: 'bar',
      itemStyle: { color: '#f0e6f7' },
      barWidth: 6,
      barGap: '2%',
      emphasis: { disabled: true }
    }],
    animation: false
  })

  // CPU使用率图表
  const cpu = echarts.init(cpuChart.value)
  cpu.setOption({
    series: [{
      type: 'gauge',
      startAngle: 180,
      endAngle: 0,
      min: 0,
      max: 100,
      splitNumber: 8,
      axisLine: { lineStyle: { width: 6, color: [[1, '#E6EBF8']] } },
      pointer: { icon: 'path://M12.8,0.7l12,40.1H0.7L12.8,0.7z', length: '12%', width: 20, offsetCenter: [0, '-60%'], itemStyle: { color: 'auto' } },
      axisTick: { length: 12, lineStyle: { color: 'auto', width: 2 } },
      splitLine: { length: 20, lineStyle: { color: 'auto', width: 5 } },
      axisLabel: { color: '#464646', fontSize: 20, distance: -60 },
      detail: { valueAnimation: true, formatter: '{value}%', color: 'auto' },
      data: [{ value: 31 }]
    }]
  })

  // 内存占用图表
  const memory = echarts.init(memoryChart.value)
  memory.setOption({
    series: [{
      type: 'gauge',
      startAngle: 180,
      endAngle: 0,
      min: 0,
      max: 100,
      splitNumber: 8,
      axisLine: { lineStyle: { width: 6, color: [[1, '#E6EBF8']] } },
      pointer: { icon: 'path://M12.8,0.7l12,40.1H0.7L12.8,0.7z', length: '12%', width: 20, offsetCenter: [0, '-60%'], itemStyle: { color: 'auto' } },
      axisTick: { length: 12, lineStyle: { color: 'auto', width: 2 } },
      splitLine: { length: 20, lineStyle: { color: 'auto', width: 5 } },
      axisLabel: { color: '#464646', fontSize: 20, distance: -60 },
      detail: { valueAnimation: true, formatter: '{value}%', color: 'auto' },
      data: [{ value: 17 }]
    }]
  })
})
</script>

<style scoped>
.overview-container {
  padding: 20px;
  background-color: #f5f5f5;
  min-height: 100vh;
}

.breadcrumb {
  margin-bottom: 20px;
  color: #666;
  font-size: 14px;
}

.data-overview {
  background: white;
  border-radius: 8px;
  padding: 20px;
  margin-bottom: 20px;
}

.data-overview h3 {
  margin: 0 0 20px 0;
  color: #333;
}

.driver-cards {
  position: relative;
}

.card-container {
  display: flex;
  gap: 15px;
  overflow-x: auto;
  padding-bottom: 20px;
}

.driver-card {
  min-width: 200px;
  background: white;
  border: 2px solid #e0e0e0;
  border-radius: 8px;
  padding: 15px;
  position: relative;
  box-shadow: 0 2px 8px rgba(0,0,0,0.06);
}
.driver-card.active {
  border: 2px solid #409EFF;
  box-shadow: 0 4px 16px rgba(64,158,255,0.15);
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 10px;
}

.driver-name {
  font-size: 12px;
  color: #333;
  font-weight: 500;
}

.status-tag {
  padding: 2px 8px;
  border-radius: 4px;
  font-size: 10px;
  color: white;
}

.status-tag.online {
  background-color: #67C23A;
}
.status-tag.offline {
  background-color: #909399;
}

.card-content {
  text-align: center;
}

.device-count {
  font-size: 18px;
  font-weight: bold;
  color: #333;
  margin-bottom: 8px;
}

.status-details {
  display: flex;
  justify-content: space-between;
  font-size: 12px;
  color: #666;
}

.card-navigation {
  text-align: center;
  margin-top: 10px;
}

.nav-dots {
  display: flex;
  justify-content: center;
  gap: 8px;
}

.dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background-color: #ddd;
  cursor: pointer;
}

.dot.active {
  background-color: #409EFF;
}

.message-trend {
  background: white;
  border-radius: 8px;
  padding: 20px;
  margin-bottom: 20px;
}

.message-trend h3 {
  margin: 0 0 20px 0;
  color: #333;
}

.trend-charts {
  display: flex;
  gap: 20px;
}
.trend-charts-vertical {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.chart-container {
  flex: 1;
}

.chart-title {
  font-size: 14px;
  color: #333;
  margin-bottom: 10px;
}
.chart-title.uplink-title {
  color: #409EFF;
}
.chart-title.downlink-title {
  color: #a259ec;
}

.chart-area {
  height: 120px;
  background-color: #f8f9fa;
  border-radius: 4px;
}

.resource-usage {
  display: flex;
  gap: 20px;
}

.resource-card {
  flex: 1;
  background: white;
  border-radius: 8px;
  padding: 20px;
}

.resource-title {
  font-size: 14px;
  color: #333;
  margin-bottom: 15px;
  text-align: center;
}

.resource-content {
  height: 200px;
}
</style> 