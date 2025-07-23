// Dashboard 数据配置文件
export const dashboardData = {
  // 平台概述数据
  platformOverview: {
    products: {
      total: 95,
      published: 40,
      unpublished: 55
    },
    devices: {
      total: 57,
      online: 5,
      offline: 52
    },
    drivers: {
      total: 8,
      running: 7,
      stopped: 1
    },
    alarms: {
      total: 0,
      note: '数据统计截止昨日24时'
    }
  },

  // 快捷入口配置
  quickEntries: [
    {
      id: 'add-product',
      icon: '📦',
      label: '添加产品',
      action: () => console.log('添加产品')
    },
    {
      id: 'add-device',
      icon: '📱',
      label: '添加设备',
      action: () => console.log('添加设备')
    },
    {
      id: 'service-monitor',
      icon: '📊',
      label: '服务监控',
      action: () => console.log('服务监控')
    },
    {
      id: 'rule-engine',
      icon: '⚙️',
      label: '规则引擎',
      action: () => console.log('规则引擎')
    },
    {
      id: 'alarm-center',
      icon: '🚨',
      label: '告警中心',
      action: () => console.log('告警中心')
    },
    {
      id: 'data-center',
      icon: '💾',
      label: '数据中心',
      action: () => console.log('数据中心')
    }
  ],

  // 系统状态数据 - 使用图一中的数据
  systemStatus: {
    cpu: {
      value: 0.00,
      unit: '%',
      label: 'CPU',
      details: '(0/2)核'
    },
    memory: {
      value: 84.95,
      unit: '%',
      label: '内存',
      details: '1.66 GB / 1.95 GB'
    },
    load: {
      value: 3.33,
      unit: '%',
      label: '负载',
      details: '运行流畅'
    },
    disk: {
      value: 36.63,
      unit: '%',
      label: '磁盘',
      details: '17.21 GB / 49.09 GB'
    }
  },

  // 告警相关数据 - 使用图一中的数据
  alarmData: {
    total: 0,
    categories: [
      {
        name: '提示',
        value: 0,
        color: '#1890ff'
      },
      {
        name: '紧急',
        value: 0,
        color: '#ff4d4f'
      },
      {
        name: '次要',
        value: 0,
        color: '#52c41a'
      },
      {
        name: '重要',
        value: 0,
        color: '#d9d9d9'
      }
    ],
    note: '数据统计截止昨日24时'
  },

  // 设备消息数据
  deviceMessages: {
    timeRanges: [
      { label: '最近一小时', value: '1h', active: true },
      { label: '最近24小时', value: '24h', active: false },
      { label: '近一周', value: '7d', active: false }
    ],
    dateRange: {
      start: '2025-07-18 20:17:54',
      end: '2025-07-18 21:17:54'
    },
    data: [], // 实际数据为空
    chartData: [] // 图表数据
  }
}

// 更新数据的函数
export const updateDashboardData = {
  // 更新平台概述
  updatePlatformOverview: (data) => {
    Object.assign(dashboardData.platformOverview, data)
  },

  // 更新系统状态
  updateSystemStatus: (data) => {
    Object.assign(dashboardData.systemStatus, data)
  },

  // 更新告警数据
  updateAlarmData: (data) => {
    Object.assign(dashboardData.alarmData, data)
  },

  // 更新设备消息数据
  updateDeviceMessages: (data) => {
    Object.assign(dashboardData.deviceMessages, data)
  },

  // 切换时间范围
  switchTimeRange: (range) => {
    dashboardData.deviceMessages.timeRanges.forEach(item => {
      item.active = item.value === range
    })
  }
}

// 模拟数据生成函数
export const generateMockData = {
  // 生成平台概述数据
  generatePlatformOverview: () => {
    return {
      products: {
        total: Math.floor(Math.random() * 100) + 50,
        published: Math.floor(Math.random() * 50) + 20,
        unpublished: Math.floor(Math.random() * 50) + 30
      },
      devices: {
        total: Math.floor(Math.random() * 100) + 30,
        online: Math.floor(Math.random() * 20) + 5,
        offline: Math.floor(Math.random() * 80) + 20
      },
      drivers: {
        total: Math.floor(Math.random() * 20) + 5,
        running: Math.floor(Math.random() * 15) + 3,
        stopped: Math.floor(Math.random() * 5) + 1
      },
      alarms: {
        total: Math.floor(Math.random() * 20),
        note: '数据统计截止昨日24时'
      }
    }
  },

  // 生成系统状态数据
  generateSystemStatus: () => {
    return {
      cpu: {
        value: Math.random() * 100,
        unit: '%',
        label: 'CPU',
        details: `(${Math.floor(Math.random() * 8)} / ${Math.floor(Math.random() * 8 + 4)})核`
      },
      memory: {
        value: Math.random() * 100,
        unit: '%',
        label: '内存',
        details: `${(Math.random() * 4).toFixed(2)} GB / ${(Math.random() * 4 + 2).toFixed(2)} GB`
      },
      load: {
        value: Math.random() * 100,
        unit: '%',
        label: '负载',
        details: '运行流畅'
      },
      disk: {
        value: Math.random() * 100,
        unit: '%',
        label: '磁盘',
        details: `${(Math.random() * 50).toFixed(2)} GB / ${(Math.random() * 100 + 50).toFixed(2)} GB`
      }
    }
  },

  // 生成告警数据
  generateAlarmData: () => {
    const categories = ['提示', '紧急', '次要', '重要']
    const colors = ['#1890ff', '#ff4d4f', '#52c41a', '#d9d9d9']
    
    const alarmCategories = categories.map((name, index) => ({
      name,
      value: Math.floor(Math.random() * 10),
      color: colors[index]
    }))

    return {
      total: alarmCategories.reduce((sum, item) => sum + item.value, 0),
      categories: alarmCategories,
      note: '数据统计截止昨日24时'
    }
  },

  // 生成设备消息数据
  generateDeviceMessages: (timeRange = '1h') => {
    const now = new Date()
    const data = []
    
    // 根据时间范围生成数据点
    let points = 60 // 默认1小时60个点
    if (timeRange === '24h') points = 24
    if (timeRange === '7d') points = 7

    for (let i = 0; i < points; i++) {
      const time = new Date(now.getTime() - (points - i) * (timeRange === '1h' ? 60000 : timeRange === '24h' ? 3600000 : 86400000))
      data.push({
        time: time.toLocaleTimeString('zh-CN', { hour: '2-digit', minute: '2-digit' }),
        value: Math.floor(Math.random() * 100)
      })
    }

    return {
      timeRanges: [
        { label: '最近一小时', value: '1h', active: timeRange === '1h' },
        { label: '最近24小时', value: '24h', active: timeRange === '24h' },
        { label: '近一周', value: '7d', active: timeRange === '7d' }
      ],
      dateRange: {
        start: data[0]?.time || now.toLocaleString('zh-CN'),
        end: data[data.length - 1]?.time || now.toLocaleString('zh-CN')
      },
      data,
      chartData: data
    }
  }
} 