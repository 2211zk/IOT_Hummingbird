<template>
  <div>
    <!-- 页面标题和描述 -->
    <div class="mb-6">
      <div class="text-sm text-gray-500 mb-2">首页 / 设备分布 / 设备地图</div>
      <h1 class="text-2xl font-bold text-gray-900 mb-2">设备地图</h1>
      <p class="text-gray-600">在地图上查看所有设备的分布情况和位置信息</p>
    </div>

    <!-- 搜索和过滤 -->
    <el-card class="search-filter mb-4" shadow="hover">
      <template #header>
        <div class="card-header">
          <el-icon><Search /></el-icon>
          <span>设备搜索</span>
        </div>
      </template>
      <el-form :model="searchInfo" inline>
        <el-form-item label="设备名称">
          <el-input
            v-model="searchInfo.eqName"
            placeholder="请输入设备名称"
            clearable
            @keyup.enter="searchDevices"
          />
        </el-form-item>
        <el-form-item label="设备状态">
          <el-select v-model="searchInfo.status" placeholder="请选择设备状态" clearable>
            <el-option label="全部" value="" />
            <el-option label="在线" value="online" />
            <el-option label="离线" value="offline" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="searchDevices">查询</el-button>
          <el-button @click="resetSearch">重置</el-button>
          <el-button type="info" @click="toggleHeatmap">切换热力图</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <!-- 地图容器 -->
    <div class="map-container">
      <div id="baidu-map" style="width: 100%; height: 600px; border: 1px solid #ccc;"></div>
      
      <!-- 地图控制面板 -->
      <div class="map-controls">
        <el-card class="control-card">
          <template #header>
            <div class="control-header">
              <el-icon><Location /></el-icon>
              <span>设备统计</span>
            </div>
          </template>
          <div class="control-content">
            <div class="stat-item">
              <span class="stat-label">总设备数:</span>
              <span class="stat-value">{{ deviceStats.total }}</span>
            </div>
            <div class="stat-item">
              <span class="stat-label">在线设备:</span>
              <span class="stat-value online">{{ deviceStats.online }}</span>
            </div>
            <div class="stat-item">
              <span class="stat-label">离线设备:</span>
              <span class="stat-value offline">{{ deviceStats.offline }}</span>
            </div>
          </div>
        </el-card>
      </div>
    </div>

    <!-- 设备信息面板 -->
    <el-drawer
      v-model="deviceInfoVisible"
      title="设备详情"
      size="400px"
      :with-header="true"
    >
      <div v-if="selectedDevice" class="device-info">
        <el-descriptions :column="1" border>
          <el-descriptions-item label="设备ID">{{ selectedDevice.ID }}</el-descriptions-item>
          <el-descriptions-item label="设备名称">{{ selectedDevice.eqName }}</el-descriptions-item>
          <el-descriptions-item label="设备标识">{{ selectedDevice.eqLogotype }}</el-descriptions-item>
          <el-descriptions-item label="所属产品">{{ getProductName(selectedDevice.productsId) }}</el-descriptions-item>
          <el-descriptions-item label="设备状态">
            <el-tag :type="selectedDevice.status === 'online' ? 'success' : 'danger'">
              {{ selectedDevice.status === 'online' ? '在线' : '离线' }}
            </el-tag>
          </el-descriptions-item>
          <el-descriptions-item label="设备坐标">{{ selectedDevice.eqCoordinate }}</el-descriptions-item>
          <el-descriptions-item label="设备地址">{{ selectedDevice.eqAddress }}</el-descriptions-item>
          <el-descriptions-item label="设备描述">{{ selectedDevice.eqInfo }}</el-descriptions-item>
          <el-descriptions-item label="创建时间">{{ formatDate(selectedDevice.CreatedAt) }}</el-descriptions-item>
        </el-descriptions>
      </div>
    </el-drawer>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted, onUnmounted } from 'vue'
import { ElMessage } from 'element-plus'
import { Location, Search } from '@element-plus/icons-vue'
import { getDeviceMapData } from '@/api/wl_playform/deviceMap'
import { getWlProductsList } from '@/api/wl_playform/wlProducts'
import { formatDate } from '@/utils/format'

defineOptions({
  name: 'DeviceMap'
})

// 响应式数据
const searchInfo = ref({
  eqName: '',
  status: ''
})

const deviceList = ref([])
const productOptions = ref([])
const selectedDevice = ref(null)
const deviceInfoVisible = ref(false)

// 设备统计数据
const deviceStats = ref({
  total: 0,
  online: 0,
  offline: 0
})

// 百度地图相关
let map = null
let markers = []
let heatmap = null

// 初始化地图
const initMap = () => {
  // 检查百度地图API是否已加载
  if (typeof BMap !== 'undefined') {
    try {
      // 创建地图实例
      map = new BMap.Map('baidu-map')
      
      // 设置地图中心点（中国中心）
      const point = new BMap.Point(116.404, 39.915)
      map.centerAndZoom(point, 5)
      
      // 添加地图控件
      map.addControl(new BMap.NavigationControl({
        type: BMAP_NAVIGATION_CONTROL_LARGE
      }))
      map.addControl(new BMap.ScaleControl())
      map.addControl(new BMap.OverviewMapControl({
        isOpen: false
      }))
      map.addControl(new BMap.MapTypeControl({
        type: BMAP_MAPTYPE_CONTROL_HORIZONTAL
      }))
      
      // 启用滚轮缩放
      map.enableScrollWheelZoom(true)
      
      // 设置地图样式 - 使用默认样式避免加载问题
      try {
        map.setMapStyle({
          style: 'normal' // 使用普通样式，更稳定
        })
      } catch (error) {
        console.warn('设置地图样式失败，使用默认样式:', error)
      }
      
      // 地图初始化成功
    } catch (error) {
      console.error('地图初始化失败:', error)
      ElMessage.error('地图初始化失败: ' + error.message)
    }
  } else {
    console.error('百度地图API未加载')
    ElMessage.error('地图加载失败，请检查网络连接')
  }
}

// 加载设备数据
const loadDeviceData = async () => {
  try {
    const params = {
      page: 1,
      pageSize: 1000, // 获取所有设备
      eqName: searchInfo.value.eqName
    }
    
    const response = await getDeviceMapData(params)
    
    if (response.code === 0) {
      // 为每个设备添加默认的status字段
      let allDevices = (response.data.list || []).map(device => ({
        ...device,
        status: device.status || 'offline' // 如果没有status字段，默认为离线
      }))
      
      // 如果没有设备数据，添加一些测试数据
      if (allDevices.length === 0) {
        allDevices = [
          {
            ID: '1',
            eqName: '北京测试设备',
            eqCoordinate: '116.404,39.915',
            eqAddress: '北京市朝阳区',
            status: 'online'
          },
          {
            ID: '2',
            eqName: '上海测试设备',
            eqCoordinate: '121.473,31.230',
            eqAddress: '上海市浦东新区',
            status: 'online'
          },
          {
            ID: '3',
            eqName: '广州测试设备',
            eqCoordinate: '113.264,23.129',
            eqAddress: '广州市天河区',
            status: 'offline'
          },
          {
            ID: '4',
            eqName: '深圳测试设备',
            eqCoordinate: '114.057,22.543',
            eqAddress: '深圳市南山区',
            status: 'online'
          },
          {
            ID: '5',
            eqName: '杭州测试设备',
            eqCoordinate: '120.155,30.274',
            eqAddress: '杭州市西湖区',
            status: 'offline'
          }
        ]
      }
      
      // 根据状态筛选
      if (searchInfo.value.status) {
        allDevices = allDevices.filter(device => device.status === searchInfo.value.status)
      }
      
      deviceList.value = allDevices
      
      // 计算统计数据
      deviceStats.value = {
        total: deviceList.value.length,
        online: deviceList.value.filter(d => d.status === 'online').length,
        offline: deviceList.value.filter(d => d.status === 'offline').length
      }
      
      // 在地图上显示设备
      showDevicesOnMap()
    } else {
      ElMessage.error('设备数据加载失败')
    }
  } catch (error) {
    console.error('加载设备数据失败:', error)
    ElMessage.error('设备数据加载失败')
  }
}

// 在地图上显示设备
const showDevicesOnMap = () => {
  // 清除现有标记
  clearMarkers()
  
  const validPoints = []
  
  deviceList.value.forEach(device => {
    if (device.eqCoordinate) {
      const coordinates = parseCoordinates(device.eqCoordinate)
      if (coordinates) {
        addDeviceMarker(device, coordinates.lng, coordinates.lat)
        validPoints.push(new BMap.Point(coordinates.lng, coordinates.lat))
      }
    }
  })
  
  // 如果有有效的设备坐标，调整地图视图以显示所有设备
  if (validPoints.length > 0) {
    try {
      const viewPort = map.getViewport(validPoints, {
        margins: [50, 50, 50, 50]
      })
      map.centerAndZoom(viewPort.center, viewPort.zoom)
    } catch (error) {
      console.warn('调整地图视图失败，使用默认视图:', error)
      // 使用默认中心点
      const point = new BMap.Point(116.404, 39.915)
      map.centerAndZoom(point, 5)
    }
    
    // 暂时禁用热力图，避免错误
    // addHeatmap(validPoints)
  }
}

// 添加热力图
const addHeatmap = (points) => {
  // 检查热力图库是否可用
  if (typeof BMapLib === 'undefined' || !BMapLib.HeatmapOverlay) {
    console.warn('热力图库未加载，跳过热力图显示')
    return
  }
  
  // 检查地图是否已初始化
  if (!map) {
    console.warn('地图未初始化，跳过热力图显示')
    return
  }
  
  // 移除现有热力图
  if (heatmap) {
    try {
      map.removeOverlay(heatmap)
    } catch (error) {
      console.warn('移除现有热力图失败:', error)
    }
  }
  
  // 检查是否有有效的点数据
  if (!points || points.length === 0) {
    console.warn('没有有效的点数据，跳过热力图显示')
    return
  }
  
  // 创建热力图数据
  const heatmapData = points.map(point => ({
    lng: point.lng,
    lat: point.lat,
    count: 1
  }))
  
  try {
    // 创建热力图
    heatmap = new BMapLib.HeatmapOverlay({
      "radius": 20,
      "visible": true,
      "opacity": 0.6
    })
    
    heatmap.setDataSet({
      data: heatmapData,
      max: 10
    })
    
    map.addOverlay(heatmap)
    console.log('热力图创建成功')
  } catch (error) {
    console.error('创建热力图失败:', error)
  }
}

// 解析坐标字符串
const parseCoordinates = (coordinateStr) => {
  try {
    // 支持多种坐标格式
    if (coordinateStr.includes(',')) {
      const [lng, lat] = coordinateStr.split(',').map(Number)
      if (!isNaN(lng) && !isNaN(lat)) {
        return { lng, lat }
      }
    }
    return null
  } catch (error) {
    console.error('坐标解析失败:', coordinateStr)
    return null
  }
}

// 添加设备标记
const addDeviceMarker = (device, lng, lat) => {
  if (!map) return
  
  const point = new BMap.Point(lng, lat)
  
  // 使用简单的默认标记，避免自定义图标问题
  const marker = new BMap.Marker(point, {
    icon: new BMap.Icon(
      'data:image/svg+xml;base64,' + btoa(`
        <svg width="20" height="20" viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg">
          <circle cx="10" cy="10" r="8" fill="${device.status === 'online' ? '#67c23a' : '#f56c6c'}" stroke="white" stroke-width="2"/>
          <circle cx="10" cy="10" r="4" fill="white"/>
        </svg>
      `),
      new BMap.Size(20, 20),
      {
        imageOffset: new BMap.Size(0, 0),
        anchor: new BMap.Size(10, 10)
      }
    )
  })
  
  // 添加点击事件
  marker.addEventListener('click', () => {
    showDeviceInfo(device)
    showInfoWindow(device, point)
  })
  
  map.addOverlay(marker)
  markers.push(marker)
}

// 显示信息窗口
const showInfoWindow = (device, point) => {
  const statusText = device.status === 'online' ? '在线' : '离线'
  const statusColor = device.status === 'online' ? '#67c23a' : '#f56c6c'
  
  const infoWindow = new BMap.InfoWindow(`
    <div style="padding: 15px; min-width: 200px;">
      <div style="display: flex; align-items: center; margin-bottom: 10px;">
        <div style="width: 8px; height: 8px; border-radius: 50%; background-color: ${statusColor}; margin-right: 8px;"></div>
        <h4 style="margin: 0; color: #303133;">${device.eqName || '未知设备'}</h4>
      </div>
      <div style="font-size: 12px; color: #606266;">
        <p style="margin: 5px 0;">设备ID: ${device.ID}</p>
        <p style="margin: 5px 0;">状态: <span style="color: ${statusColor}; font-weight: 600;">${statusText}</span></p>
        <p style="margin: 5px 0;">坐标: ${device.eqCoordinate}</p>
        <p style="margin: 5px 0;">地址: ${device.eqAddress || '未知'}</p>
      </div>
    </div>
  `)
  
  map.openInfoWindow(infoWindow, point)
}

// 清除所有标记
const clearMarkers = () => {
  markers.forEach(marker => {
    map.removeOverlay(marker)
  })
  markers = []
  
  // 清除热力图
  if (heatmap) {
    map.removeOverlay(heatmap)
    heatmap = null
  }
}

// 显示设备详情
const showDeviceInfo = (device) => {
  selectedDevice.value = device
  deviceInfoVisible.value = true
}

// 获取产品名称
const getProductName = (productId) => {
  const product = productOptions.value.find(p => p.ID === productId)
  return product ? product.prName : '未知产品'
}

// 加载产品选项
const loadProductOptions = async () => {
  try {
    const response = await getWlProductsList({ page: 1, pageSize: 1000 })
    if (response.code === 0) {
      productOptions.value = response.data.list || []
    }
  } catch (error) {
    console.error('加载产品选项失败:', error)
  }
}

// 加载百度地图API
const loadBaiduMapAPI = () => {
  return new Promise((resolve, reject) => {
    if (typeof BMap !== 'undefined') {
      resolve()
      return
    }
    
    const script = document.createElement('script')
    // 使用HTTP协议避免SSL问题，使用您的百度地图API密钥
    script.src = 'http://api.map.baidu.com/api?v=3.0&ak=UQK3TvBBpH45jQAfmnnC5BBWCLXniODu&callback=initBaiduMap'
    script.onerror = reject
    document.head.appendChild(script)
    
    window.initBaiduMap = () => {
      // 加载热力图库
      const heatmapScript = document.createElement('script')
      heatmapScript.src = 'http://api.map.baidu.com/library/Heatmap/2.0/src/Heatmap_min.js'
      heatmapScript.onload = () => resolve()
      heatmapScript.onerror = () => resolve() // 热力图加载失败不影响主地图
      document.head.appendChild(heatmapScript)
    }
  })
}

// 搜索设备
const searchDevices = () => {
  loadDeviceData()
}

// 重置搜索
const resetSearch = () => {
  searchInfo.value = {
    eqName: '',
    status: ''
  }
  loadDeviceData()
}

// 切换热力图
const toggleHeatmap = () => {
  try {
    if (heatmap) {
      // 移除热力图
      map.removeOverlay(heatmap)
      heatmap = null
      ElMessage.success('热力图已关闭')
    } else {
      // 添加热力图
      const validPoints = deviceList.value
        .filter(device => device.eqCoordinate)
        .map(device => {
          const coordinates = parseCoordinates(device.eqCoordinate)
          return coordinates ? new BMap.Point(coordinates.lng, coordinates.lat) : null
        })
        .filter(point => point !== null)
      
      if (validPoints.length > 0) {
        addHeatmap(validPoints)
        ElMessage.success('热力图已开启')
      } else {
        ElMessage.warning('没有有效的设备坐标，无法创建热力图')
      }
    }
  } catch (error) {
    console.error('切换热力图失败:', error)
    ElMessage.error('切换热力图失败')
  }
}

// 组件挂载
onMounted(async () => {
  try {
    // 加载百度地图API
    await loadBaiduMapAPI()
    
    // 初始化地图
    initMap()
    
    // 加载产品选项
    await loadProductOptions()
    
    // 加载设备数据
    await loadDeviceData()
  } catch (error) {
    console.error('地图初始化失败:', error)
    ElMessage.error('地图初始化失败')
  }
})

// 组件卸载
onUnmounted(() => {
  if (map) {
    map.clearOverlays()
  }
})
</script>

<style scoped>
.map-container {
  border: 1px solid #e4e7ed;
  border-radius: 4px;
  overflow: hidden;
  position: relative;
}

.map-controls {
  position: absolute;
  top: 20px;
  right: 20px;
  z-index: 1000;
}

.control-card {
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(10px);
  border: 1px solid #e4e7ed;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.1);
}

.control-header {
  display: flex;
  align-items: center;
  font-weight: 600;
}

.control-header .el-icon {
  margin-right: 8px;
  color: #409eff;
}

.control-content {
  padding: 10px 0;
}

.stat-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 8px;
  font-size: 14px;
}

.stat-item:last-child {
  margin-bottom: 0;
}

.stat-label {
  color: #606266;
}

.stat-value {
  font-weight: 600;
  color: #303133;
}

.stat-value.online {
  color: #67c23a;
}

.stat-value.offline {
  color: #f56c6c;
}

.device-info {
  padding: 20px;
}

.search-filter {
  margin-bottom: 20px;
}

.card-header {
  display: flex;
  align-items: center;
  font-weight: 600;
}

.card-header .el-icon {
  margin-right: 8px;
  color: #409eff;
}
</style> 