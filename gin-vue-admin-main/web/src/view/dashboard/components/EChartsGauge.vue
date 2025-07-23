<template>
  <div class="echarts-gauge">
    <div ref="chartRef" class="chart-container"></div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted, watch } from 'vue'
import * as echarts from 'echarts'

const props = defineProps({
  value: {
    type: Number,
    default: 0
  },
  name: {
    type: String,
    default: ''
  },
  unit: {
    type: String,
    default: '%'
  },
  details: {
    type: String,
    default: ''
  },
  size: {
    type: String,
    default: '120px'
  }
})

const chartRef = ref(null)
let chartInstance = null

// 根据数值获取颜色
const getColor = (value) => {
  if (value < 30) return '#52c41a' // 绿色
  if (value < 70) return '#faad14' // 橙色
  return '#ff4d4f' // 红色
}

// 初始化图表
const initChart = () => {
  if (!chartRef.value) return

  chartInstance = echarts.init(chartRef.value)
  
  const color = getColor(props.value)
  
  const option = {
    series: [
      {
        type: 'gauge',
        startAngle: 180,
        endAngle: 0,
        min: 0,
        max: 100,
        splitNumber: 10,
        radius: '90%',
        center: ['50%', '60%'],
        axisLine: {
          lineStyle: {
            width: 8,
            color: [
              [0.3, '#52c41a'],
              [0.7, '#faad14'],
              [1, '#ff4d4f']
            ]
          }
        },
        pointer: {
          icon: 'path://M12.8,0.7l12,40.9H0.7L12.8,0.7z',
          length: '12%',
          width: 20,
          offsetCenter: [0, '-60%'],
          itemStyle: {
            color: 'auto'
          }
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
          fontSize: 12,
          distance: -60,
          formatter: function (value) {
            if (value === 0.875) {
              return '100'
            } else if (value === 0.625) {
              return '75'
            } else if (value === 0.375) {
              return '50'
            } else if (value === 0.125) {
              return '25'
            }
            return ''
          }
        },
        title: {
          offsetCenter: [0, '-20%'],
          fontSize: 14,
          color: '#333'
        },
        detail: {
          fontSize: 20,
          offsetCenter: [0, '10%'],
          valueAnimation: true,
          formatter: function (value) {
            return value.toFixed(2) + props.unit
          },
          color: 'auto'
        },
        data: [
          {
            value: props.value,
            name: props.name
          }
        ]
      }
    ]
  }
  
  chartInstance.setOption(option)
}

// 更新图表
const updateChart = () => {
  if (!chartInstance) return
  
  const color = getColor(props.value)
  
  const option = {
    series: [
      {
        data: [
          {
            value: props.value,
            name: props.name
          }
        ],
        detail: {
          formatter: function (value) {
            return value.toFixed(2) + props.unit
          }
        }
      }
    ]
  }
  
  chartInstance.setOption(option)
}

// 监听数据变化
watch(() => props.value, () => {
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
.echarts-gauge {
  width: 100%;
  height: 100%;
  
  .chart-container {
    width: v-bind(size);
    height: v-bind(size);
    margin: 0 auto;
  }
}
</style> 