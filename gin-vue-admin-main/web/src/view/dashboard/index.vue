<template>
  <div class="dashboard-container" style="background-color: #0f172a !important; min-height: 100vh; padding: 15px;">
    <!-- 平台概述 -->
    <div class="overview-section">
      <div class="overview-card" style="background-color: #0f172a !important; border-radius: 8px; padding: 15px; box-shadow: 0 2px 8px rgba(0, 0, 0, 0.4); border: 1px solid #334155; color: white;">
        <h3>平台概述</h3>
        <div class="overview-grid">
          <div class="overview-item">
            <div class="overview-number">{{ platformData.productCount }}</div>
            <div class="overview-label">产品数</div>
            <div class="overview-details">
              <span>已发布 {{ platformData.publishedProducts }}个</span>
              <span>未发布 {{ platformData.unpublishedProducts }}个</span>
            </div>
          </div>
          <div class="overview-item">
            <div class="overview-number">{{ platformData.deviceCount }}</div>
            <div class="overview-label">设备数</div>
            <div class="overview-details">
              <span>在线 {{ platformData.onlineDevices }}个</span>
              <span>离线 {{ platformData.offlineDevices }}个</span>
            </div>
          </div>
          <div class="overview-item">
            <div class="overview-number">{{ platformData.driverCount }}</div>
            <div class="overview-label">驱动数</div>
            <div class="overview-details">
              <span>运行中 {{ platformData.runningDrivers }}个</span>
              <span>已停止 {{ platformData.stoppedDrivers }}个</span>
            </div>
          </div>
          <div class="overview-item">
            <div class="overview-number">{{ platformData.alarmCount }}</div>
            <div class="overview-label">告警数</div>
            <div class="overview-details">
              <span>数据统计截止昨日24时</span>
            </div>
          </div>
        </div>
      </div>
      
      <div class="quick-entry-card" style="background-color: #0f172a !important; border-radius: 8px; padding: 15px; box-shadow: 0 2px 8px rgba(0, 0, 0, 0.4); border: 1px solid #334155; color: white;">
        <h3>快捷入口</h3>
        <div class="quick-entry-grid">
          <div 
            v-for="entryType in quickEntryTypes" 
            :key="entryType"
            class="quick-entry-item" 
            :class="{ 'loading': quickEntryLoading[entryType] }"
            @click="handleQuickEntry(entryType)" 
            :title="getQuickEntryTooltip(entryType)"
          >
            <div class="quick-entry-icon" v-if="!quickEntryLoading[entryType]">
              {{ getQuickEntryIcon(entryType) }}
            </div>
            <div class="quick-entry-loading" v-else>⏳</div>
            <div class="quick-entry-label">{{ getQuickEntryLabel(entryType) }}</div>
          </div>
        </div>
      </div>
    </div>

    <!-- 状态监控 -->
    <div class="status-section">
      <div class="status-card" style="background-color: #0f172a !important; border-radius: 8px; padding: 15px; box-shadow: 0 2px 8px rgba(0, 0, 0, 0.4); border: 1px solid #334155; color: white;">
        <h3>状态</h3>
        <div class="status-grid">
          <div class="status-item" @click="handleStatusItemClick('cpu')">
            <div ref="cpuChart" class="chart-container"></div>
            <div class="status-info">
              <div class="status-value">{{ systemStatus.cpu.usage }}%</div>
              <div class="status-label">CPU ({{ systemStatus.cpu.used }}/{{ systemStatus.cpu.total }})核</div>
              <div class="status-desc">{{ systemStatus.cpu.status }}</div>
            </div>
          </div>
          <div class="status-item" @click="handleStatusItemClick('memory')">
            <div ref="memoryChart" class="chart-container"></div>
            <div class="status-info">
              <div class="status-value">{{ systemStatus.memory.usage }}%</div>
              <div class="status-label">内存</div>
              <div class="status-desc">{{ systemStatus.memory.used }} GB / {{ systemStatus.memory.total }} GB</div>
            </div>
          </div>
          <div class="status-item" @click="handleStatusItemClick('load')">
            <div ref="loadChart" class="chart-container"></div>
            <div class="status-info">
              <div class="status-value">{{ systemStatus.load.usage }}%</div>
              <div class="status-label">负载</div>
              <div class="status-desc">{{ systemStatus.load.status }}</div>
            </div>
          </div>
          <div class="status-item" @click="handleStatusItemClick('disk')">
            <div ref="diskChart" class="chart-container"></div>
            <div class="status-info">
              <div class="status-value">{{ systemStatus.disk.usage }}%</div>
              <div class="status-label">磁盘</div>
              <div class="status-desc">{{ systemStatus.disk.used }} GB / {{ systemStatus.disk.total }} GB</div>
            </div>
          </div>
        </div>
      </div>
      
      <div class="alarm-card" style="background-color: #0f172a !important; border-radius: 8px; padding: 15px; box-shadow: 0 2px 8px rgba(0, 0, 0, 0.4); border: 1px solid #334155; color: white;">
        <h3>告警相关</h3>
        <div ref="alarmChart" class="alarm-chart-container"></div>
        <div class="alarm-legend">
          <div class="alarm-legend-item">
            <span class="legend-color hint"></span>
            <span>提示 {{ alarmData.hint }}</span>
          </div>
          <div class="alarm-legend-item">
            <span class="legend-color minor"></span>
            <span>次要 {{ alarmData.minor }}</span>
          </div>
          <div class="alarm-legend-item">
            <span class="legend-color important"></span>
            <span>重要 {{ alarmData.important }}</span>
          </div>
          <div class="alarm-legend-item">
            <span class="legend-color urgent"></span>
            <span>紧急 {{ alarmData.urgent }}</span>
          </div>
        </div>
        <div class="alarm-note">数据统计截止昨日24时</div>
      </div>
    </div>

    <!-- 设备消息总数 -->
    <div class="message-section">
      <div class="message-card" style="background-color: #0f172a !important; border-radius: 8px; padding: 15px; box-shadow: 0 2px 8px rgba(0, 0, 0, 0.4); border: 1px solid #334155; color: white;">
        <div class="message-header">
          <h3>设备消息总数</h3>
          <div class="time-selector">
            <button 
              v-for="timeRange in timeRanges" 
              :key="timeRange.key"
              :class="['time-btn', { active: selectedTimeRange === timeRange.key }]"
              @click="selectTimeRange(timeRange.key)"
            >
              {{ timeRange.label }}
            </button>
            <div class="date-picker">
              <span>{{ dateRange.start }}</span>
              <span>至</span>
              <span>{{ dateRange.end }}</span>
            </div>
          </div>
        </div>
        <div ref="messageChart" class="message-chart-container"></div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted, nextTick } from 'vue'
import { useRouter } from 'vue-router'
import { useRouterStore } from '@/pinia/modules/router'
import * as echarts from 'echarts'
import { getDashboardData } from '@/api/dashboard/dashboard'
import { ElMessage } from 'element-plus'

defineOptions({
  name: 'Dashboard'
})

// 平台概述数据
const platformData = reactive({
  productCount: 0,
  publishedProducts: 0,
  unpublishedProducts: 0,
  deviceCount: 0,
  onlineDevices: 0,
  offlineDevices: 0,
  driverCount: 0,
  runningDrivers: 0,
  stoppedDrivers: 0,
  alarmCount: 0
})

// 系统状态数据
const systemStatus = reactive({
  cpu: {
    usage: 0,
    used: 0,
    total: 2,
    status: '运行流畅'
  },
  memory: {
    usage: 0,
    used: 0,
    total: 1.95,
    status: '内存使用率较高'
  },
  load: {
    usage: 0,
    status: '运行流畅'
  },
  disk: {
    usage: 0,
    used: 0,
    total: 49.09,
    status: '磁盘空间充足'
  }
})

// 告警数据
const alarmData = reactive({
  hint: 0,
  minor: 0,
  important: 0,
  urgent: 0
})

// 时间范围选择
const timeRanges = [
  { key: '1h', label: '最近一小时' },
  { key: '24h', label: '最近24小时' },
  { key: '7d', label: '近一周' }
]
const selectedTimeRange = ref('1h')
const dateRange = reactive({
  start: '2025-07-20 20:41:48',
  end: '2025-07-20 21:41:48'
})

// 图表引用
const cpuChart = ref(null)
const memoryChart = ref(null)
const loadChart = ref(null)
const diskChart = ref(null)
const alarmChart = ref(null)
const messageChart = ref(null)

// 获取路由实例
const router = useRouter()
const routerStore = useRouterStore()

// 快捷入口加载状态
const quickEntryLoading = ref({})

// 快捷入口配置
const quickEntryConfig = {
  'addProduct': {
    name: 'wlProducts',
    label: '产品管理',
    icon: '📦',
    description: '管理物联网产品信息'
  },
  'addDevice': {
    name: 'wlEquipment', 
    label: '设备管理',
    icon: '📱',
    description: '管理物联网设备信息'
  },
  'serviceMonitor': {
    name: 'state',
    label: '系统监控',
    icon: '🖥️',
    description: '监控系统运行状态'
  },
  'ruleEngine': {
    name: 'wlEngineRules',
    label: '引擎规则',
    icon: '⚙️',
    description: '配置业务规则引擎'
  },
  'alarmCenter': {
    name: 'wlAlarm',
    label: '告警中心',
    icon: '🔔',
    description: '查看和处理系统告警'
  },
  'dataCenter': {
    name: 'state',
    label: '服务器状态',
    icon: '💾',
    description: '查看服务器运行状态'
  }
}

// 快捷入口类型列表
const quickEntryTypes = ['addProduct', 'addDevice', 'serviceMonitor', 'ruleEngine', 'alarmCenter', 'dataCenter']

// 获取快捷入口显示信息的方法
const getQuickEntryIcon = (entryType) => {
  return quickEntryConfig[entryType]?.icon || '📄'
}

const getQuickEntryLabel = (entryType) => {
  return quickEntryConfig[entryType]?.label || '未知功能'
}

const getQuickEntryTooltip = (entryType) => {
  const config = quickEntryConfig[entryType]
  return config?.description || `点击跳转到${config?.label || '未知功能'}页面`
}

// 图表实例
let cpuChartInstance = null
let memoryChartInstance = null
let loadChartInstance = null
let diskChartInstance = null
let alarmChartInstance = null
let messageChartInstance = null

// 快捷入口处理
const handleQuickEntry = async (type) => {
  console.log('快捷入口点击:', type)
  
  // 防止重复点击
  if (quickEntryLoading.value[type]) {
    return
  }

  const config = quickEntryConfig[type]
  if (!config) {
    ElMessage.error('功能暂未开放')
    return
  }

  try {
    // 设置加载状态
    quickEntryLoading.value[type] = true
    
    console.log(`正在跳转到${config.label}页面...`)
    console.log(`路由名称: ${config.name}`)
    console.log(`当前路由:`, router.currentRoute.value)
    
    // 检查路由是否存在
    if (routerStore && routerStore.routeMap) {
      console.log(`路由映射:`, Object.keys(routerStore.routeMap))
      console.log(`目标路由是否存在:`, !!routerStore.routeMap[config.name])
    }
    
    // 执行路由跳转
    await router.push({ name: config.name })
    
    console.log(`成功跳转到${config.label}页面`)
    
    // 显示成功提示
    ElMessage.success(`已跳转到${config.label}`)
    
  } catch (error) {
    console.error(`跳转失败:`, error)
    console.error(`错误详情:`, error.message)
    
    // 错误处理
    if (error.message && error.message.includes('No match')) {
      ElMessage.error(`${config.label}页面不存在或暂未配置`)
    } else {
      ElMessage.error(`跳转到${config.label}失败，请稍后重试`)
    }
    
  } finally {
    // 清除加载状态
    setTimeout(() => {
      quickEntryLoading.value[type] = false
    }, 500)
  }
}



// 获取仪表盘数据
const fetchDashboardData = async () => {
  try {
    const res = await getDashboardData()
    if (res.code === 0) {
      const data = res.data
      
      // 更新平台数据
      Object.assign(platformData, data.platformData)
      
      // 更新系统状态数据
      Object.assign(systemStatus, data.systemStatus)
      
      // 更新告警数据
      Object.assign(alarmData, data.alarmData)
      
      // 更新图表
      nextTick(() => {
        // 暂时注释掉，因为updateCharts函数未定义
        // updateCharts()
      })
    }
  } catch (error) {
    console.error('获取仪表盘数据失败:', error)
  }
}

// 时间范围选择
const selectTimeRange = (range) => {
  selectedTimeRange.value = range
  updateMessageChart()
}

// 初始化仪表盘图表
const initGaugeChart = (element, value, title, color) => {
  const chart = echarts.init(element)
  const option = {
    series: [{
      type: 'gauge',
      startAngle: 180,
      endAngle: 0,
      min: 0,
      max: 100,
      splitNumber: 8,
      axisLine: {
        lineStyle: {
          width: 6,
          color: [
            [0.3, '#67e0e3'],
            [0.7, '#37a2da'],
            [1, '#fd666d']
          ]
        }
      },
      pointer: {
        icon: 'path://M12.8,31.7l9.7,9.7c0.8,0.8,2.1,0.8,2.9,0l9.7-9.7c0.8-0.8,0.8-2.1,0-2.9l-9.7-9.7c-0.8-0.8-2.1-0.8-2.9,0l-9.7,9.7C12,29.6,12,30.9,12.8,31.7z',
        length: '12%',
        width: 20,
        offsetCenter: [0, '-60%']
      },
      axisTick: {
        length: 12,
        lineStyle: {
          color: 'auto',
          width: 2
        }
      },
      splitLine: {
        length: 20,
        lineStyle: {
          color: 'auto',
          width: 5
        }
      },
      axisLabel: {
        color: '#464646',
        fontSize: 16,
        distance: -60,
        formatter: function (value) {
          if (value === 0.875) {
            return '100'
          } else if (value === 0.625) {
            return '50'
          } else if (value === 0.375) {
            return '25'
          } else if (value === 0.125) {
            return '0'
          }
          return ''
        }
      },
      title: {
        offsetCenter: [0, '-20%'],
        fontSize: 16,
        color: '#333'
      },
      detail: {
        show: false // 隐藏中央的数值显示
      },
      data: [{
        value: value,
        name: title
      }]
    }]
  }
  chart.setOption(option)
  return chart
}

// 更新仪表盘数值
const updateGaugeValue = (chartInstance, newValue, title) => {
  if (chartInstance) {
    chartInstance.setOption({
      series: [{
        data: [{
          value: newValue,
          name: title
        }]
      }]
    })
  }
}

// 模拟实时数据更新
const simulateRealTimeData = () => {
  // CPU 使用率随机变化
  const cpuValue = Math.random() * 100
  updateGaugeValue(cpuChartInstance, cpuValue, 'CPU')
  systemStatus.cpu.usage = cpuValue.toFixed(2)
  systemStatus.cpu.used = Math.floor(cpuValue / 50 * systemStatus.cpu.total)
  
  // 内存使用率随机变化
  const memoryValue = Math.random() * 100
  updateGaugeValue(memoryChartInstance, memoryValue, '内存')
  systemStatus.memory.usage = memoryValue.toFixed(2)
  systemStatus.memory.used = (memoryValue / 100 * systemStatus.memory.total).toFixed(2)
  
  // 负载随机变化
  const loadValue = Math.random() * 100
  updateGaugeValue(loadChartInstance, loadValue, '负载')
  systemStatus.load.usage = loadValue.toFixed(2)
  
  // 磁盘使用率随机变化
  const diskValue = Math.random() * 100
  updateGaugeValue(diskChartInstance, diskValue, '磁盘')
  systemStatus.disk.usage = diskValue.toFixed(2)
  systemStatus.disk.used = (diskValue / 100 * systemStatus.disk.total).toFixed(2)
  
  // 更新状态描述
  updateStatusDescriptions()
}

// 更新状态描述
const updateStatusDescriptions = () => {
  // CPU 状态描述
  if (systemStatus.cpu.usage < 30) {
    systemStatus.cpu.status = '运行流畅'
  } else if (systemStatus.cpu.usage < 70) {
    systemStatus.cpu.status = '运行正常'
  } else {
    systemStatus.cpu.status = '运行繁忙'
  }
  
  // 内存状态描述
  if (systemStatus.memory.usage < 60) {
    systemStatus.memory.status = '内存充足'
  } else if (systemStatus.memory.usage < 85) {
    systemStatus.memory.status = '内存使用率较高'
  } else {
    systemStatus.memory.status = '内存使用率过高'
  }
  
  // 负载状态描述
  if (systemStatus.load.usage < 30) {
    systemStatus.load.status = '运行流畅'
  } else if (systemStatus.load.usage < 70) {
    systemStatus.load.status = '运行正常'
  } else {
    systemStatus.load.status = '负载较高'
  }
  
  // 磁盘状态描述
  if (systemStatus.disk.usage < 50) {
    systemStatus.disk.status = '磁盘空间充足'
  } else if (systemStatus.disk.usage < 80) {
    systemStatus.disk.status = '磁盘空间正常'
  } else {
    systemStatus.disk.status = '磁盘空间不足'
  }
}

// 添加点击事件来手动调整指针位置
const handleStatusItemClick = (type) => {
  let newValue
  switch (type) {
    case 'cpu':
      newValue = Math.random() * 100
      updateGaugeValue(cpuChartInstance, newValue, 'CPU')
      systemStatus.cpu.usage = newValue.toFixed(2)
      systemStatus.cpu.used = Math.floor(newValue / 50 * systemStatus.cpu.total)
      break
    case 'memory':
      newValue = Math.random() * 100
      updateGaugeValue(memoryChartInstance, newValue, '内存')
      systemStatus.memory.usage = newValue.toFixed(2)
      systemStatus.memory.used = (newValue / 100 * systemStatus.memory.total).toFixed(2)
      break
    case 'load':
      newValue = Math.random() * 100
      updateGaugeValue(loadChartInstance, newValue, '负载')
      systemStatus.load.usage = newValue.toFixed(2)
      break
    case 'disk':
      newValue = Math.random() * 100
      updateGaugeValue(diskChartInstance, newValue, '磁盘')
      systemStatus.disk.usage = newValue.toFixed(2)
      systemStatus.disk.used = (newValue / 100 * systemStatus.disk.total).toFixed(2)
      break
  }
  updateStatusDescriptions()
}

// 添加鼠标悬停效果
const addHoverEffect = () => {
  const statusItems = document.querySelectorAll('.status-item')
  statusItems.forEach(item => {
    item.addEventListener('mouseenter', () => {
      item.style.transform = 'scale(1.02)'
      item.style.transition = 'transform 0.2s ease'
    })
    
    item.addEventListener('mouseleave', () => {
      item.style.transform = 'scale(1)'
    })
  })
}

// 初始化告警饼图
const initAlarmChart = (element) => {
  const chart = echarts.init(element)
  const option = {
    tooltip: {
      trigger: 'item'
    },
    legend: {
      orient: 'vertical',
      left: 'left'
    },
    series: [{
      name: '告警类型',
      type: 'pie',
      radius: '50%',
      data: [
        { value: alarmData.hint, name: '提示', itemStyle: { color: '#409EFF' } },
        { value: alarmData.minor, name: '次要', itemStyle: { color: '#67C23A' } },
        { value: alarmData.important, name: '重要', itemStyle: { color: '#E6A23C' } },
        { value: alarmData.urgent, name: '紧急', itemStyle: { color: '#F56C6C' } }
      ],
      emphasis: {
        itemStyle: {
          shadowBlur: 10,
          shadowOffsetX: 0,
          shadowColor: 'rgba(0, 0, 0, 0.5)'
        }
      }
    }]
  }
  chart.setOption(option)
  return chart
}

// 初始化消息折线图
const initMessageChart = (element) => {
  const chart = echarts.init(element)
  const option = {
    tooltip: {
      trigger: 'axis'
    },
    xAxis: {
      type: 'category',
      data: ['20:45', '20:50', '20:55', '21:00', '21:05', '21:10', '21:15', '21:20', '21:25', '21:30', '21:35', '21:40', '21:41']
    },
    yAxis: {
      type: 'value'
    },
    series: [{
      data: [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
      type: 'line',
      smooth: true,
      symbol: 'circle',
      symbolSize: 6,
      lineStyle: {
        color: '#409EFF',
        width: 2
      },
      itemStyle: {
        color: '#409EFF'
      }
    }]
  }
  chart.setOption(option)
  return chart
}

// 更新消息图表
const updateMessageChart = () => {
  if (messageChartInstance) {
    // 这里可以根据选择的时间范围更新数据
    console.log('更新消息图表，时间范围:', selectedTimeRange.value)
  }
}

// 初始化所有图表
const initCharts = async () => {
  await nextTick()
  
  if (cpuChart.value) {
    cpuChartInstance = initGaugeChart(cpuChart.value, systemStatus.cpu.usage, 'CPU', '#67e0e3')
  }
  
  if (memoryChart.value) {
    memoryChartInstance = initGaugeChart(memoryChart.value, systemStatus.memory.usage, '内存', '#fd666d')
  }
  
  if (loadChart.value) {
    loadChartInstance = initGaugeChart(loadChart.value, systemStatus.load.usage, '负载', '#67e0e3')
  }
  
  if (diskChart.value) {
    diskChartInstance = initGaugeChart(diskChart.value, systemStatus.disk.usage, '磁盘', '#37a2da')
  }
  
  if (alarmChart.value) {
    alarmChartInstance = initAlarmChart(alarmChart.value)
  }
  
  if (messageChart.value) {
    messageChartInstance = initMessageChart(messageChart.value)
  }
}

// 预加载常用快捷入口组件（暂时禁用）
// const preloadQuickEntryComponents = async () => {
//   // 预加载逻辑暂时禁用
// }

// 组件挂载后初始化
onMounted(() => {
  // 初始化时获取数据
  fetchDashboardData()
  
  initCharts()
  
  // 添加悬停效果
  nextTick(() => {
    addHoverEffect()
  })
  
  // 调试路由信息
  nextTick(() => {
    console.log('=== 路由调试信息 ===')
    console.log('Router实例:', router)
    console.log('RouterStore实例:', routerStore)
    console.log('当前路由:', router.currentRoute.value)
    
    if (routerStore && routerStore.routeMap) {
      console.log('可用路由:', Object.keys(routerStore.routeMap))
    } else {
      console.warn('RouterStore或routeMap未初始化')
    }
  })
  
  // 监听窗口大小变化，重新调整图表大小
  window.addEventListener('resize', () => {
    cpuChartInstance?.resize()
    memoryChartInstance?.resize()
    loadChartInstance?.resize()
    diskChartInstance?.resize()
    alarmChartInstance?.resize()
    messageChartInstance?.resize()
  })
  
  // 启动实时数据更新
  setInterval(() => {
    simulateRealTimeData()
  }, 3000) // 每3秒更新一次
})

// 定期刷新数据
setInterval(() => {
  fetchDashboardData()
}, 30000) // 每30秒刷新一次
</script>

<style lang="scss">
/* 强制覆盖所有白色背景 */
.dashboard-container * {
  background-color: inherit !important;
}

.dashboard-container .overview-card,
.dashboard-container .quick-entry-card,
.dashboard-container .status-card,
.dashboard-container .alarm-card,
.dashboard-container .message-card {
  background-color: #0f172a !important;
  color: white !important;
  border-radius: 8px !important;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.4) !important;
  border: 1px solid #334155 !important;
}

.dashboard-container .quick-entry-item {
  background-color: #1890ff !important;
  color: white !important;
  border: 1px solid #1890ff !important;
  font-weight: bold !important;
  cursor: pointer !important;
  transition: all 0.3s ease !important;
  border-radius: 8px !important;
  padding: 16px 12px !important;
  display: flex !important;
  flex-direction: column !important;
  align-items: center !important;
  justify-content: center !important;
  min-height: 80px !important;
  position: relative !important;
  overflow: hidden !important;
}

.dashboard-container .quick-entry-item:hover {
  background-color: #40a9ff !important;
  border-color: #40a9ff !important;
  transform: translateY(-3px) !important;
  box-shadow: 0 6px 16px rgba(24, 144, 255, 0.4) !important;
}

.dashboard-container .quick-entry-item:active {
  transform: translateY(-1px) !important;
  box-shadow: 0 2px 8px rgba(24, 144, 255, 0.3) !important;
}

.dashboard-container .quick-entry-item::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.2), transparent);
  transition: left 0.5s;
}

.dashboard-container .quick-entry-item:hover::before {
  left: 100%;
}

.dashboard-container .quick-entry-item.loading {
  opacity: 0.7;
  cursor: not-allowed;
  transform: none !important;
}

.dashboard-container .quick-entry-item.loading:hover {
  transform: none !important;
  box-shadow: 0 2px 8px rgba(24, 144, 255, 0.3) !important;
}

.quick-entry-loading {
  font-size: 24px;
  margin-bottom: 8px;
  animation: spin 1s linear infinite;
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}

.status-item {
  display: flex;
  align-items: center;
  gap: 20px;
  padding: 10px;
  border-radius: 8px;
  transition: all 0.3s ease;
  cursor: pointer;
}

.status-item:hover {
  background-color: #f8f9fa;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.chart-container {
  width: 120px;
  height: 120px;
  position: relative;
  transition: transform 0.3s ease;
}

.chart-container:hover {
  transform: scale(1.05);
}

.status-info {
  flex: 1;
  min-width: 0;
  transition: all 0.3s ease;
}

.status-value {
  font-size: 28px;
  font-weight: bold;
  color: #409EFF;
  margin-bottom: 6px;
  line-height: 1.2;
  transition: color 0.3s ease;
}

.status-label {
  font-size: 12px;
  font-weight: 600;
  color: #333;
  margin-bottom: 4px;
  line-height: 1.2;
  transition: color 0.3s ease;
}

.status-desc {
  font-size: 10px;
  color: #666;
  line-height: 1.2;
  transition: color 0.3s ease;
}

.status-item:hover .status-value {
  color: #1890ff;
  transform: scale(1.05);
}

.status-item:hover .status-label {
  color: #409EFF;
}

.status-item:hover .status-desc {
  color: #409EFF;
}

.status-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 30px;
}

.overview-number {
  font-size: 28px;
  font-weight: bold;
  color: #409EFF;
  margin-bottom: 6px;
}

.overview-label {
  font-size: 12px;
  color: #ccc;
  margin-bottom: 6px;
}

.overview-details {
  font-size: 10px;
  color: #ccc;
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.quick-entry-label {
  font-size: 12px;
  color: white;
  font-weight: 500;
  text-align: center;
  line-height: 1.2;
}

.alarm-legend-item {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 10px;
  color: #ccc;
}

.legend-color {
  width: 12px;
  height: 12px;
  border-radius: 2px;
}

.legend-color.hint { background-color: #409EFF; }
.legend-color.minor { background-color: #67C23A; }
.legend-color.important { background-color: #E6A23C; }
.legend-color.urgent { background-color: #F56C6C; }

.alarm-note {
  font-size: 10px;
  color: #ccc;
  text-align: center;
}

.message-section {
  margin-bottom: 15px;
}

.message-card {
  background: #0f172a !important;
  border-radius: 8px;
  padding: 20px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.4);
  border: 1px solid #334155;
  color: white;
}

.message-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 15px;
}

.message-header h3 {
  margin: 0 0 15px 0;
  font-size: 16px;
  font-weight: 600;
  color: #333;
}

.time-selector {
  display: flex;
  align-items: center;
  gap: 12px;
}

.time-btn {
  padding: 4px 8px;
  border: 1px solid #555;
  background: #3d3d3d !important;
  border-radius: 4px;
  cursor: pointer;
  font-size: 10px;
  transition: all 0.3s;
  color: white;
}

.time-btn.active {
  background: #409EFF;
  color: white;
  border-color: #409EFF;
}

.date-picker {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 10px;
  color: #ccc;
}

.message-chart-container {
  height: 280px;
}

.overview-card, .quick-entry-card, .status-card, .alarm-card, .message-card {
  background: #0f172a !important;
  border-radius: 8px;
  padding: 15px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.4);
  border: 1px solid #334155;
  color: white;
}

.overview-section {
  display: grid;
  grid-template-columns: 2fr 1fr;
  gap: 15px;
  margin-bottom: 15px;
}

.quick-entry-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 12px;
}

.quick-entry-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 12px;
  border: 1px solid #555;
  border-radius: 6px;
  cursor: pointer;
  transition: all 0.3s;
  background-color: #3d3d3d !important;
  color: white;
}

.quick-entry-item:hover {
  border-color: #409EFF;
  background-color: #40a9ff !important;
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(24, 144, 255, 0.3);
}

.quick-entry-icon {
  font-size: 24px;
  margin-bottom: 8px;
  transition: transform 0.3s ease;
}

.dashboard-container .quick-entry-item:hover .quick-entry-icon {
  transform: scale(1.1);
}

.alarm-chart-container {
  height: 180px;
  margin-bottom: 12px;
}

.alarm-legend {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 8px;
  margin-bottom: 8px;
}

.dashboard-container {
  padding: 15px;
  background-color: #d0d0d0 !important;
  min-height: 100vh;
}

.overview-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 15px;
}

.overview-item {
  text-align: center;
}

.status-section {
  display: grid;
  grid-template-columns: 2fr 1fr;
  gap: 15px;
  margin-bottom: 15px;
}
</style>
