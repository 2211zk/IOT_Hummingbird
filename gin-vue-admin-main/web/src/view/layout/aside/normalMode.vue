<template>
  <div
    class="relative h-full bg-white text-slate-700 dark:text-slate-300 dark:bg-slate-900 border-r shadow dark:shadow-gray-700"
    :class="isCollapse ? '' : '  px-2'"
    :style="{
      width: layoutSideWidth + 'px'
    }"
  >
    <!-- 调试信息 - 用于开发时查看菜单数据结构，生产环境应设置为 false -->
    <div v-if="false" style="background: red; color: white; padding: 10px; font-size: 12px;">
      <div>菜单数量: {{ routerStore.asyncRouters[0]?.children?.length || 0 }}</div>
      <div>菜单项: {{ routerStore.asyncRouters[0]?.children?.map(item => item.name).join(', ') }}</div>
    </div>
    
    <el-scrollbar>
      <el-menu
        :collapse="isCollapse"
        :collapse-transition="false"
        :default-active="active"
        class="border-r-0 w-full"
        unique-opened
        @select="selectMenuItem"
      >
        <!-- 使用处理后的菜单项，实现自定义的菜单层级结构 -->
        <template v-for="item in processedMenuItems">
          <aside-component
            v-if="!item.hidden"
            :key="item.name"
            :router-info="item"
          />
        </template>
      </el-menu>
    </el-scrollbar>
    <div
      class="absolute bottom-8 right-2 w-8 h-8 bg-gray-50 dark:bg-slate-800 flex items-center justify-center rounded cursor-pointer"
      :class="isCollapse ? 'right-0 left-0 mx-auto' : 'right-2'"
      @click="toggleCollapse"
    >
      <el-icon v-if="!isCollapse">
        <DArrowLeft />
      </el-icon>
      <el-icon v-else>
        <DArrowRight />
      </el-icon>
    </div>
  </div>
</template>

<script setup>
  import AsideComponent from '@/view/layout/aside/asideComponent/index.vue'
  import { ref, provide, watchEffect, computed } from 'vue'
  import { useRoute, useRouter } from 'vue-router'
  import { useRouterStore } from '@/pinia/modules/router'
  import { useAppStore } from '@/pinia'
  import { storeToRefs } from 'pinia'
  const appStore = useAppStore()
  const { device, config } = storeToRefs(appStore)

  defineOptions({
    name: 'GvaAside'
  })
  const route = useRoute()
  const router = useRouter()
  const routerStore = useRouterStore()
  const isCollapse = ref(false)
  const active = ref('')

  /**
   * 处理菜单项，创建自定义的层级结构
   * 目的：将"产品管理"和"设备管理"作为"设备接入"的子菜单显示
   * 实现方式：动态创建"设备接入"父菜单，并将原有的两个菜单作为其子菜单
   */
  const processedMenuItems = computed(() => {
    // 获取从后端接收到的原始菜单数据
    const items = routerStore.asyncRouters[0]?.children || []
    console.log('原始菜单项:', items)
    
    // 确保所有菜单都有正确的标题
    const processedItems = items.map(item => {
      // 确保每个菜单项都有正确的meta信息
      if (!item.meta) {
        item.meta = {}
      }
      
      // 为特定菜单设置正确的标题
      if (item.name === 'wlResources') {
        item.meta.title = '资源管理'
      } else if (item.name === 'wlProducts') {
        item.meta.title = '产品管理'
      } else if (item.name === 'wlEquipment') {
        item.meta.title = '设备管理'
      } else if (item.name === 'wlScenes') {
        item.meta.title = '场景联动'
      } else if (item.name === 'wlEngineRules') {
        item.meta.title = '引擎规则'
      }
      
      return item
    })
    
    // 查找需要重新组织的菜单项
    const productMenu = processedItems.find(item => item.name === 'wlProducts')
    const equipmentMenu = processedItems.find(item => item.name === 'wlEquipment')
    const resourcesMenu = processedItems.find(item => item.name === 'wlResources')
    const scenesMenu = processedItems.find(item => item.name === 'wlScenes')
    const engineRulesMenu = processedItems.find(item => item.name === 'wlEngineRules')
    const monitorMenu = processedItems.find(item => item.name === 'wlMonitor')
    const alarmMenu = processedItems.find(item => item.name === 'wlAlarm')
    
    console.log('找到的菜单:', { productMenu, equipmentMenu, resourcesMenu, scenesMenu, engineRulesMenu })
    
    // 创建系统管理父菜单 - 直接添加到现有系统管理菜单中
    const systemManagementMenu = processedItems.find(item => item.name === 'superAdmin')
    
    // 如果找到系统管理菜单，添加部门管理子菜单
    if (systemManagementMenu) {
      // 确保children数组存在
      if (!systemManagementMenu.children) {
        systemManagementMenu.children = []
      }
      
      // 检查是否已经存在部门管理菜单
      const existingDeptMenu = systemManagementMenu.children.find(child => child.name === 'DepartmentManagement')
      if (!existingDeptMenu) {
        // 添加部门管理子菜单
        systemManagementMenu.children.push({
          name: 'DepartmentManagement',
          path: 'system/department',
          component: 'view/system/department/index.vue',
          meta: {
            title: '部门管理',
            icon: 'office-building',
            keepAlive: true
          },
          hidden: false
        })
      }
      
      // 检查是否已经存在设备管理菜单
      const existingDeviceMenu = systemManagementMenu.children.find(child => child.name === 'DeviceManagement')
      if (!existingDeviceMenu) {
        // 添加设备管理子菜单
        systemManagementMenu.children.push({
          name: 'DeviceManagement',
          path: 'system/device',
          component: 'view/system/device/index.vue',
          meta: {
            title: '设备管理',
            icon: 'monitor',
            keepAlive: true
          },
          hidden: false
        })
      }
    }
    
    // 创建高级能力父菜单
    const advancedCapabilitiesMenu = {
      name: 'advancedCapabilities',
      path: 'advancedCapabilities',
      component: 'view/routerHolder.vue',
      meta: {
        title: '高级能力',
        icon: 'magic-stick'
      },
      children: []
    }
    
    // 将资源管理、场景联动、引擎规则添加到高级能力子菜单
    if (resourcesMenu) {
      advancedCapabilitiesMenu.children.push({
        ...resourcesMenu,
        hidden: false,
        meta: { ...resourcesMenu.meta, title: '资源管理', icon: 'link' }
      })
    }
    
    if (scenesMenu) {
      advancedCapabilitiesMenu.children.push({
        ...scenesMenu,
        hidden: false,
        meta: { ...scenesMenu.meta, title: '场景联动', icon: 'connection' }
      })
    }
    
    if (engineRulesMenu) {
      advancedCapabilitiesMenu.children.push({
        ...engineRulesMenu,
        hidden: false,
        meta: { ...engineRulesMenu.meta, title: '引擎规则', icon: 'document' }
      })
    }
    
    // 如果找到了产品管理和设备管理菜单，则进行重组
    if (productMenu && equipmentMenu) {
      console.log('开始处理菜单层级结构')
      
      // 动态创建设备接入父菜单
      const deviceAccessMenu = {
        name: 'wl_playform',
        path: 'wl_playform',
        component: 'view/wl_playform/deviceAccess/index.vue',
        meta: {
          title: '设备接入',
          icon: 'connection'
        },
        children: [
          // 将产品管理作为设备接入的子菜单
          {
            ...productMenu,
            hidden: false,
            meta: { ...productMenu.meta, title: '产品管理', icon: 'box' }
          },
          // 将设备管理作为设备接入的子菜单
          {
            ...equipmentMenu,
            hidden: false,
            meta: { ...equipmentMenu.meta, title: '设备管理', icon: 'monitor' }
          }
        ]
      }
      
      // 创建运维监控父菜单
      const opsMonitorMenu = {
        name: 'opsMonitor',
        path: 'opsMonitor',
        component: 'view/routerHolder.vue',
        meta: {
          title: '运维监控',
          icon: 'data-board'
        },
        children: []
      }
      if (monitorMenu) {
        opsMonitorMenu.children.push({
          ...monitorMenu,
          hidden: false,
          meta: { ...monitorMenu.meta, title: '系统监控', icon: 'monitor' }
        })
      }
      if (alarmMenu) {
        opsMonitorMenu.children.push({
          ...alarmMenu,
          hidden: false,
          meta: { ...alarmMenu.meta, title: '告警中心', icon: 'bell' }
        })
      }

      // 在filteredItems中移除wlMonitor和wlAlarm，插入opsMonitorMenu
      const filteredItems = processedItems.filter(item => 
        item.name !== 'wlProducts' && 
        item.name !== 'wlEquipment' && 
        item.name !== 'wlResources' && 
        item.name !== 'wlScenes' && 
        item.name !== 'wlEngineRules' &&
        item.name !== 'wlMonitor' &&
        item.name !== 'wlAlarm'
      )
      
      // 将设备接入菜单插入到合适的位置（在dashboard之后）
      const dashboardIndex = filteredItems.findIndex(item => item.name === 'dashboard')
      if (dashboardIndex !== -1) {
        filteredItems.splice(dashboardIndex + 1, 0, deviceAccessMenu)
        filteredItems.splice(dashboardIndex + 2, 0, opsMonitorMenu)
      } else {
        // 如果找不到dashboard，则将设备接入菜单放在最前面
        filteredItems.unshift(deviceAccessMenu)
        filteredItems.splice(1, 0, opsMonitorMenu)
      }
      
      // 将高级能力菜单插入到合适的位置
      const insertIndex = Math.min(2, filteredItems.length) // 插入到前3个位置
      filteredItems.splice(insertIndex, 0, advancedCapabilitiesMenu)
      
      console.log('处理后的菜单:', filteredItems)
      return filteredItems
    }
    
    // 如果没有找到需要重组的菜单，则返回处理后的菜单
    return processedItems
  })
  const layoutSideWidth = computed(() => {
    if (!isCollapse.value) {
      return config.value.layout_side_width
    } else {
      return config.value.layout_side_collapsed_width
    }
  })

  watchEffect(() => {
    if (route.name === 'Iframe') {
      active.value = decodeURIComponent(route.query.url)
      return
    }
    active.value = route.meta.activeName || route.name
  })

  watchEffect(() => {
    if (device.value === 'mobile') {
      isCollapse.value = true
    } else {
      isCollapse.value = false
    }
  })

  provide('isCollapse', isCollapse)

  const selectMenuItem = (index) => {
    const query = {}
    const params = {}
    routerStore.routeMap[index]?.parameters &&
      routerStore.routeMap[index]?.parameters.forEach((item) => {
        if (item.type === 'query') {
          query[item.key] = item.value
        } else {
          params[item.key] = item.value
        }
      })
    if (index === route.name) return
    if (index.indexOf('http://') > -1 || index.indexOf('https://') > -1) {
      if (index === 'Iframe') {
        query.url = decodeURIComponent(index)
        router.push({
          name: 'Iframe',
          query,
          params
        })
        return
      } else {
        window.open(index, '_blank')
        return
      }
    } else {
      router.push({ name: index, query, params })
    }
  }

  const toggleCollapse = () => {
    isCollapse.value = !isCollapse.value
  }
</script>

<style lang="scss"></style>
