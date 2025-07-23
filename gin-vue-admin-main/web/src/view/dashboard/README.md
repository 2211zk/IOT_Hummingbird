# Dashboard 仪表盘

这是一个现代化的IoT设备管理仪表盘，使用ECharts提供专业的图表可视化效果。

## 功能特性

### 1. 平台概述
- 产品数量统计（已发布/未发布）
- 设备数量统计（在线/离线）
- 驱动数量统计（运行中/已停止）
- 告警数量统计

### 2. 快捷入口
- 添加产品
- 添加设备
- 服务监控
- 规则引擎
- 告警中心
- 数据中心

### 3. 系统状态监控（ECharts仪表盘）
- CPU使用率（实时监控，动态颜色）
- 内存使用率（实时监控，动态颜色）
- 负载状态（实时监控，动态颜色）
- 磁盘使用率（实时监控，动态颜色）

### 4. 告警管理（ECharts环形图）
- 告警类型统计（提示/紧急/次要/重要）
- 环形图可视化
- 实时告警数据
- 交互式图例

### 5. 设备消息统计（ECharts折线图）
- 时间范围选择（最近一小时/24小时/一周）
- 消息数量趋势图
- 实时数据更新
- 平滑曲线动画

## 文件结构

```
dashboard/
├── index.vue                    # 主页面组件
├── data.js                     # 数据配置文件
├── api.js                      # API服务文件
├── components/                  # 子组件目录
│   ├── EChartsGauge.vue        # ECharts仪表盘组件
│   ├── EChartsDonut.vue        # ECharts环形图组件
│   ├── EChartsLine.vue         # ECharts折线图组件
│   ├── charts.vue              # 原有图表组件
│   ├── charts-people-numbers.vue
│   ├── charts-content-numbers.vue
│   └── ...
└── README.md                   # 说明文档
```

## ECharts图表特性

### EChartsGauge.vue - 仪表盘组件
- 支持动态颜色（根据数值变化）
- 平滑的动画效果
- 响应式设计
- 自定义大小和样式

```javascript
// 使用示例
<EChartsGauge 
  :value="23.38"
  :name="'CPU'"
  :unit="'%'"
  :color="'#52c41a'"
  size="120px"
/>
```

### EChartsDonut.vue - 环形图组件
- 支持多数据系列
- 交互式图例
- 悬停效果
- 自定义颜色

```javascript
// 使用示例
<EChartsDonut 
  :data="alarmChartData"
  size="200px"
/>
```

### EChartsLine.vue - 折线图组件
- 平滑曲线
- 渐变填充
- 响应式坐标轴
- 交互式提示

```javascript
// 使用示例
<EChartsLine 
  :data="deviceMessages.chartData"
  height="200px"
/>
```

## 数据配置

### data.js
包含所有静态数据配置和模拟数据生成函数：

```javascript
// 系统状态数据
systemStatus: {
  cpu: { 
    value: 23.38, 
    unit: '%', 
    label: 'CPU', 
    details: '(0 / 2)核',
    color: '#52c41a'  // 动态颜色
  },
  memory: { 
    value: 18.71, 
    unit: '%', 
    label: '内存', 
    details: '1.66 GB / 1.95 GB',
    color: '#52c41a'  // 动态颜色
  }
}

// 告警数据
alarmData: {
  total: 10,
  categories: [
    {
      name: '提示',
      value: 0,
      color: '#1890ff'
    },
    {
      name: '紧急',
      value: 1,
      color: '#ff4d4f'
    }
  ]
}
```

### api.js
提供模拟的API接口，可以轻松替换为真实的后端API：

```javascript
// 获取系统状态数据
export const getSystemStatus = async () => {
  await delay(300)
  return {
    cpu: {
      value: 23.38,
      unit: '%',
      label: 'CPU',
      details: '(0 / 2)核',
      color: '#52c41a'
    }
  }
}
```

## 使用方法

### 1. 基本使用
页面会自动加载所有数据，ECharts图表会自动渲染。

### 2. 自定义图表颜色
修改数据中的color属性：

```javascript
// 修改CPU颜色
dashboardData.systemStatus.cpu.color = '#ff4d4f'
```

### 3. 连接真实API
替换 `api.js` 中的模拟API为真实接口：

```javascript
export const getSystemStatus = async () => {
  const response = await fetch('/api/system/status')
  return response.json()
}
```

### 4. 添加新的图表类型
创建新的ECharts组件：

```vue
<!-- 新建 EChartsBar.vue -->
<template>
  <div class="echarts-bar">
    <div ref="chartRef" class="chart-container"></div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted, watch } from 'vue'
import * as echarts from 'echarts'

// 实现柱状图逻辑
</script>
```

## 响应式设计

页面支持多种屏幕尺寸：
- 桌面端：完整布局，图表自适应
- 平板端：自适应布局
- 移动端：垂直堆叠布局

## 实时更新

- 系统状态数据每30秒自动更新
- 告警数据实时刷新
- ECharts图表平滑动画
- 支持手动刷新功能

## 错误处理

- 网络错误自动重试
- 加载状态提示
- 错误信息显示
- 重试按钮功能

## 主题定制

可以通过修改CSS变量和ECharts配置来自定义主题：

```scss
// 主色调
--primary-color: #1890ff;
--success-color: #52c41a;
--warning-color: #faad14;
--error-color: #ff4d4f;
```

## ECharts配置

### 仪表盘配置
```javascript
const option = {
  series: [{
    type: 'gauge',
    startAngle: 90,
    endAngle: -270,
    progress: {
      show: true,
      roundCap: true,
      itemStyle: {
        color: {
          type: 'linear',
          colorStops: getColorStops(props.color)
        }
      }
    }
  }]
}
```

### 环形图配置
```javascript
const option = {
  series: [{
    type: 'pie',
    radius: ['40%', '70%'],
    center: ['40%', '50%'],
    itemStyle: {
      borderRadius: 10,
      borderColor: '#fff',
      borderWidth: 2
    }
  }]
}
```

### 折线图配置
```javascript
const option = {
  series: [{
    type: 'line',
    smooth: true,
    symbol: 'none',
    lineStyle: {
      width: 3,
      color: '#1890ff'
    },
    areaStyle: {
      color: {
        type: 'linear',
        colorStops: [
          { offset: 0, color: 'rgba(24, 144, 255, 0.3)' },
          { offset: 1, color: 'rgba(24, 144, 255, 0.1)' }
        ]
      }
    }
  }]
}
```

## 扩展功能

### 1. 添加新的图表类型
在 `components` 目录下创建新的ECharts组件。

### 2. 添加数据导出功能
实现数据导出为Excel或PDF功能。

### 3. 添加数据筛选功能
实现按时间、设备类型等条件筛选数据。

### 4. 添加通知功能
实现实时通知和消息推送功能。

## 注意事项

1. 确保所有API接口返回正确的数据格式
2. 实时更新功能会消耗一定的网络资源
3. 建议在生产环境中使用真实的API接口
4. 可以根据实际需求调整更新频率
5. ECharts图表会自动响应窗口大小变化

## 技术栈

- Vue 3 Composition API
- ECharts 5.5.1
- SCSS 样式预处理
- ES6+ JavaScript
- 响应式设计
- 模块化架构 