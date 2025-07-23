// Dashboard API 服务
import { generateMockData } from './data.js'

// 模拟API延迟
const delay = (ms) => new Promise(resolve => setTimeout(resolve, ms))

// 平台概述API
export const getPlatformOverview = async () => {
  await delay(500) // 模拟网络延迟
  return {
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
  }
}

// 系统状态API - 使用图一中的数据
export const getSystemStatus = async () => {
  await delay(300)
  return {
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
  }
}

// 告警数据API - 使用图一中的数据
export const getAlarmData = async () => {
  await delay(400)
  return {
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
  }
}

// 设备消息数据API
export const getDeviceMessages = async (timeRange = '1h') => {
  await delay(600)
  return generateMockData.generateDeviceMessages(timeRange)
}

// 快捷入口配置API
export const getQuickEntries = async () => {
  await delay(200)
  return [
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
  ]
}

// 获取实时数据API
export const getRealTimeData = async () => {
  await delay(1000)
  return {
    platformOverview: {
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
    },
    systemStatus: {
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
    },
    alarmData: {
      total: Math.floor(Math.random() * 20),
      categories: [
        {
          name: '提示',
          value: Math.floor(Math.random() * 5),
          color: '#1890ff'
        },
        {
          name: '紧急',
          value: Math.floor(Math.random() * 3),
          color: '#ff4d4f'
        },
        {
          name: '次要',
          value: Math.floor(Math.random() * 10),
          color: '#52c41a'
        },
        {
          name: '重要',
          value: Math.floor(Math.random() * 5),
          color: '#d9d9d9'
        }
      ],
      note: '数据统计截止昨日24时'
    },
    deviceMessages: generateMockData.generateDeviceMessages('1h')
  }
} 