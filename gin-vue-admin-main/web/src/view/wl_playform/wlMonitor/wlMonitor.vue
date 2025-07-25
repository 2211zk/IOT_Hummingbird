<template>
  <div class="monitor-page">
    <!-- 顶部统计卡片 -->
    <div class="monitor-cards">
      <div class="monitor-card">
        <div class="monitor-card-title">CPU使用率</div>
        <div class="monitor-card-value">{{ monitorStats.cpu }} <span class="unit">%</span></div>
      </div>
      <div class="monitor-card">
        <div class="monitor-card-title">内存占用</div>
        <div class="monitor-card-value">1.53 <span class="unit">GB</span></div>
        <div class="monitor-card-sub">内存总量 1.95 GB</div>
      </div>
      <div class="monitor-card">
        <div class="monitor-card-title">系统负载率</div>
        <div class="monitor-card-value">{{ monitorStats.load }} <span class="unit">%</span></div>
        <div class="monitor-card-sub">流畅程度 {{ monitorStats.loadDesc }}</div>
      </div>
      <div class="monitor-card">
        <div class="monitor-card-title">磁盘占用</div>
        <div class="monitor-card-value">17.11 <span class="unit">GB</span></div>
        <div class="monitor-card-sub">磁盘总量 49.09 GB</div>
      </div>
    </div>
    <!-- 网络流量折线图 -->
    <div class="monitor-section">
      <div class="monitor-section-header">
        <span>网络流量</span>
        <el-button-group>
          <el-button :type="netTab==='up'?'primary':'default'" @click="netTab='up'">上行</el-button>
          <el-button :type="netTab==='down'?'primary':'default'" @click="netTab='down'">下行</el-button>
        </el-button-group>
        <el-button-group style="margin-left:auto;">
          <el-button :type="netRange==='1h'?'primary':'default'" @click="netRange='1h'">最近一小时</el-button>
          <el-button :type="netRange==='24h'?'primary':'default'" @click="netRange='24h'">最近24小时</el-button>
          <el-button :type="netRange==='1w'?'primary':'default'" @click="netRange='1w'">近一周</el-button>
        </el-button-group>
      </div>
      <v-chart :option="netOption" autoresize style="height: 260px;" />
    </div>
    <!-- CPU/内存趋势 -->
    <div class="monitor-trend-row">
      <div class="monitor-section monitor-trend">
        <div class="monitor-section-header">
          <span>CPU使用率趋势</span>
          <el-button-group style="margin-left:auto;">
            <el-button :type="cpuRange==='1h'?'primary':'default'" @click="cpuRange='1h'">最近一小时</el-button>
            <el-button :type="cpuRange==='24h'?'primary':'default'" @click="cpuRange='24h'">最近24小时</el-button>
            <el-button :type="cpuRange==='1w'?'primary':'default'" @click="cpuRange='1w'">近一周</el-button>
          </el-button-group>
        </div>
        <v-chart :option="cpuOption" autoresize style="height: 180px;" />
      </div>
      <div class="monitor-section monitor-trend">
        <div class="monitor-section-header">
          <span>内存使用量趋势</span>
          <el-button-group style="margin-left:auto;">
            <el-button :type="memRange==='1h'?'primary':'default'" @click="memRange='1h'">最近一小时</el-button>
            <el-button :type="memRange==='24h'?'primary':'default'" @click="memRange='24h'">最近24小时</el-button>
            <el-button :type="memRange==='1w'?'primary':'default'" @click="memRange='1w'">近一周</el-button>
          </el-button-group>
        </div>
        <v-chart :option="memOption" autoresize style="height: 180px;" />
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted, watch } from 'vue'
import VChart from 'vue-echarts'
import { use } from 'echarts/core'
import { LineChart } from 'echarts/charts'
import { GridComponent, TooltipComponent, LegendComponent } from 'echarts/components'
import { CanvasRenderer } from 'echarts/renderers'
use([LineChart, GridComponent, TooltipComponent, LegendComponent, CanvasRenderer])

const netTab = ref('up')
const netRange = ref('1h')
const cpuRange = ref('1h')
const memRange = ref('1h')

// 顶部统计卡片数据
const monitorStats = ref({
  cpu: 0,
  memoryUsed: 1.53,
  memoryTotal: 1.95,
  load: 10,
  loadDesc: '运行流畅',
  diskUsed: 17.11,
  diskTotal: 49.09
})

let timer = null
function randomStats() {
  // CPU 0-100
  monitorStats.value.cpu = +(Math.random() * 100).toFixed(2)
  // 内存 0~total
  monitorStats.value.memoryUsed = +(Math.random() * monitorStats.value.memoryTotal).toFixed(2)
  // 负载 0-100
  monitorStats.value.load = +(Math.random() * 100).toFixed(2)
  // 负载描述
  if (monitorStats.value.load < 30) monitorStats.value.loadDesc = '运行流畅'
  else if (monitorStats.value.load < 70) monitorStats.value.loadDesc = '运行正常'
  else monitorStats.value.loadDesc = '负载较高'
  // 磁盘 0~total
  monitorStats.value.diskUsed = +(Math.random() * monitorStats.value.diskTotal).toFixed(2)
}
onMounted(() => {
  randomStats()
  timer = setInterval(randomStats, 3000)
})
onUnmounted(() => {
  if (timer) clearInterval(timer)
})

// 生成随机折线数据工具
function genRandomLineData(pointCount, min, max) {
  return Array(pointCount).fill(0).map(() => Math.floor(Math.random() * (max - min + 1)) + min)
}
function genRandomTimeLabels(pointCount, unit = 'minute') {
  const now = new Date()
  let arr = []
  for (let i = pointCount - 1; i >= 0; i--) {
    let d = new Date(now)
    if (unit === 'minute') d.setMinutes(now.getMinutes() - i * 5)
    if (unit === 'hour') d.setHours(now.getHours() - i)
    if (unit === 'day') d.setDate(now.getDate() - i)
    arr.push(d.toLocaleString('zh-CN', { hour12: false }))
  }
  return arr
}

// 网络流量数据结构
const netData = ref({
  up: {
    '1h': { x: genRandomTimeLabels(12, 'minute'), y: [genRandomLineData(12, 1000000, 2000000), genRandomLineData(12, 500000, 1500000)] },
    '24h': { x: genRandomTimeLabels(24, 'hour'), y: [genRandomLineData(24, 1000000, 2000000), genRandomLineData(24, 500000, 1500000)] },
    '1w': { x: genRandomTimeLabels(7, 'day'), y: [genRandomLineData(7, 1000000, 2000000), genRandomLineData(7, 500000, 1500000)] }
  },
  down: {
    '1h': { x: genRandomTimeLabels(12, 'minute'), y: [genRandomLineData(12, 800000, 1800000), genRandomLineData(12, 400000, 1200000)] },
    '24h': { x: genRandomTimeLabels(24, 'hour'), y: [genRandomLineData(24, 800000, 1800000), genRandomLineData(24, 400000, 1200000)] },
    '1w': { x: genRandomTimeLabels(7, 'day'), y: [genRandomLineData(7, 800000, 1800000), genRandomLineData(7, 400000, 1200000)] }
  }
})

// CPU趋势数据结构
const cpuData = ref({
  '1h': { x: genRandomTimeLabels(12, 'minute'), y: [genRandomLineData(12, 0, 8), genRandomLineData(12, 0, 8)] },
  '24h': { x: genRandomTimeLabels(24, 'hour'), y: [genRandomLineData(24, 0, 8), genRandomLineData(24, 0, 8)] },
  '1w': { x: genRandomTimeLabels(7, 'day'), y: [genRandomLineData(7, 0, 8), genRandomLineData(7, 0, 8)] }
})

// 内存趋势数据结构
const memData = ref({
  '1h': { x: genRandomTimeLabels(12, 'minute'), y: [genRandomLineData(12, 100, 600), genRandomLineData(12, 100, 600)] },
  '24h': { x: genRandomTimeLabels(24, 'hour'), y: [genRandomLineData(24, 100, 600), genRandomLineData(24, 100, 600)] },
  '1w': { x: genRandomTimeLabels(7, 'day'), y: [genRandomLineData(7, 100, 600), genRandomLineData(7, 100, 600)] }
})

// 定时刷新当前tab和range下的数据
let chartTimer = null
function randomizeChartData() {
  // 网络流量
  for (const tab of ['up', 'down']) {
    for (const range of ['1h', '24h', '1w']) {
      netData.value[tab][range].y = [genRandomLineData(netData.value[tab][range].x.length, 1000000, 2000000), genRandomLineData(netData.value[tab][range].x.length, 500000, 1500000)]
    }
  }
  // CPU
  for (const range of ['1h', '24h', '1w']) {
    cpuData.value[range].y = [genRandomLineData(cpuData.value[range].x.length, 0, 8), genRandomLineData(cpuData.value[range].x.length, 0, 8)]
  }
  // 内存
  for (const range of ['1h', '24h', '1w']) {
    memData.value[range].y = [genRandomLineData(memData.value[range].x.length, 100, 600), genRandomLineData(memData.value[range].x.length, 100, 600)]
  }
}
onMounted(() => {
  chartTimer = setInterval(randomizeChartData, 600000) // 10分钟刷新一次
})
onUnmounted(() => {
  if (chartTimer) clearInterval(chartTimer)
})

// 网络流量option
const netOption = computed(() => {
  const d = netData.value[netTab.value][netRange.value]
  return {
    tooltip: { trigger: 'axis' },
    legend: { data: ['tcp-official-driver', 'mysql'] },
    grid: { left: 60, right: 30, top: 40, bottom: 30 },
    xAxis: { type: 'category', data: d.x },
    yAxis: { type: 'value', name: '网络流量(KB)' },
    series: [
      { name: 'tcp-official-driver', type: 'line', data: d.y[0] },
      { name: 'mysql', type: 'line', data: d.y[1] }
    ]
  }
})

// CPU趋势option
const cpuOption = computed(() => {
  const d = cpuData.value[cpuRange.value]
  return {
    tooltip: { trigger: 'axis' },
    legend: { data: ['mysql', 'tcp-official-driver'] },
    grid: { left: 60, right: 30, top: 40, bottom: 30 },
    xAxis: { type: 'category', data: d.x },
    yAxis: { type: 'value', name: 'CPU使用率(%)' },
    series: [
      { name: 'mysql', type: 'line', data: d.y[0] },
      { name: 'tcp-official-driver', type: 'line', data: d.y[1] }
    ]
  }
})

// 内存趋势option
const memOption = computed(() => {
  const d = memData.value[memRange.value]
  return {
    tooltip: { trigger: 'axis' },
    legend: { data: ['http-official-driver', 'hummingbird-core'] },
    grid: { left: 60, right: 30, top: 40, bottom: 30 },
    xAxis: { type: 'category', data: d.x },
    yAxis: { type: 'value', name: '内存使用量(MB)' },
    series: [
      { name: 'http-official-driver', type: 'line', data: d.y[0] },
      { name: 'hummingbird-core', type: 'line', data: d.y[1] }
    ]
  }
})
</script>

<style scoped>
.monitor-page {
  padding: 24px;
  background: #f5f6fa;
}
.monitor-cards {
  display: flex;
  gap: 24px;
  margin-bottom: 24px;
}
.monitor-card {
  flex: 1;
  background: #fff;
  border-radius: 12px;
  box-shadow: 0 2px 8px #e3e6f0;
  padding: 24px 18px 18px 24px;
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  min-width: 200px;
}
.monitor-card-title {
  font-size: 15px;
  color: #888;
  margin-bottom: 8px;
}
.monitor-card-value {
  font-size: 32px;
  font-weight: bold;
  color: #222;
}
.unit {
  font-size: 16px;
  color: #bbb;
  margin-left: 2px;
}
.monitor-card-sub {
  font-size: 13px;
  color: #bbb;
  margin-top: 6px;
}
.monitor-section {
  background: #fff;
  border-radius: 12px;
  box-shadow: 0 2px 8px #e3e6f0;
  padding: 18px 24px 12px 24px;
  margin-bottom: 24px;
}
.monitor-section-header {
  display: flex;
  align-items: center;
  font-size: 16px;
  font-weight: 500;
  margin-bottom: 12px;
  gap: 18px;
}
.monitor-trend-row {
  display: flex;
  gap: 24px;
}
.monitor-trend {
  flex: 1;
}
</style> 