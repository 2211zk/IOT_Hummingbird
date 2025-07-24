<template>
  <div class="echarts-line">
    <div ref="chartRef" class="chart-container"></div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted, watch } from 'vue'
import * as echarts from 'echarts'

const props = defineProps({
  data: {
    type: Array,
    default: () => []
  },
  height: {
    type: String,
    default: '200px'
  }
})

const chartRef = ref(null)
let chartInstance = null

// 初始化图表
const initChart = () => {
  if (!chartRef.value) return

  chartInstance = echarts.init(chartRef.value)
  
  const option = {
    grid: {
      left: '3%',
      right: '4%',
      bottom: '3%',
      top: '3%',
      containLabel: true
    },
    tooltip: {
      trigger: 'axis',
      axisPointer: {
        type: 'cross',
        label: {
          backgroundColor: '#6a7985'
        }
      }
    },
    xAxis: {
      type: 'category',
      boundaryGap: false,
      data: props.data.map(item => item.time),
      axisLine: {
        lineStyle: {
          color: '#e0e0e0'
        }
      },
      axisLabel: {
        color: '#999',
        fontSize: 12
      }
    },
    yAxis: {
      type: 'value',
      axisLine: {
        lineStyle: {
          color: '#e0e0e0'
        }
      },
      axisLabel: {
        color: '#999',
        fontSize: 12
      },
      splitLine: {
        lineStyle: {
          color: '#f5f5f5'
        }
      }
    },
    series: [
      {
        name: '消息数量',
        type: 'line',
        smooth: true,
        symbol: 'none',
        lineStyle: {
          width: 2,
          color: '#1890ff'
        },
        areaStyle: {
          color: {
            type: 'linear',
            x: 0,
            y: 0,
            x2: 0,
            y2: 1,
            colorStops: [
              { offset: 0, color: 'rgba(24, 144, 255, 0.2)' },
              { offset: 1, color: 'rgba(24, 144, 255, 0.05)' }
            ]
          }
        },
        data: props.data.map(item => item.value)
      }
    ]
  }
  
  chartInstance.setOption(option)
}

// 更新图表
const updateChart = () => {
  if (!chartInstance) return
  
  const option = {
    xAxis: {
      data: props.data.map(item => item.time)
    },
    series: [
      {
        data: props.data.map(item => item.value)
      }
    ]
  }
  
  chartInstance.setOption(option)
}

// 监听数据变化
watch(() => props.data, () => {
  updateChart()
}, { deep: true })

// 监听窗口大小变化
const handleResize = () => {
  if (chartInstance) {
    chartInstance.resize()
  }
}

onMounted(() => {
  initChart()
  window.addEventListener('resize', handleResize)
})

onUnmounted(() => {
  if (chartInstance) {
    chartInstance.dispose()
  }
  window.removeEventListener('resize', handleResize)
})
</script>

<style scoped lang="scss">
.echarts-line {
  width: 100%;
  height: 100%;
  
  .chart-container {
    width: 100%;
    height: v-bind(height);
  }
}
</style> 