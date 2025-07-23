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
<<<<<<< HEAD
=======
        <!-- 使用处理后的菜单项，实现自定义的菜单层级结构 -->
>>>>>>> ae3240f93462583fbaf4769f0c2d15372eb41e0e
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
    
    // 查找需要重新组织的菜单项
    const productMenu = items.find(item => item.name === 'wlProducts')
    const equipmentMenu = items.find(item => item.name === 'wlEquipment')
    
    console.log('找到的菜单:', { productMenu, equipmentMenu })
    
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
      
      // 从原始菜单中移除产品管理和设备管理，避免重复显示
      const filteredItems = items.filter(item => 
        item.name !== 'wlProducts' && item.name !== 'wlEquipment'
      )
      
      // 将设备接入菜单插入到合适的位置（在dashboard之后）
      const dashboardIndex = filteredItems.findIndex(item => item.name === 'dashboard')
      if (dashboardIndex !== -1) {
        filteredItems.splice(dashboardIndex + 1, 0, deviceAccessMenu)
      } else {
        // 如果找不到dashboard，则将设备接入菜单放在最前面
        filteredItems.unshift(deviceAccessMenu)
      }
      
      console.log('处理后的菜单:', filteredItems)
      return filteredItems
    }
    
    // 如果没有找到需要重组的菜单，则返回原始菜单
    // 但是需要确保设备分布菜单被包含
    const deviceDistributionMenu = items.find(item => item.name === 'deviceDistribution')
    if (deviceDistributionMenu) {
      console.log('找到设备分布菜单:', deviceDistributionMenu)
      return items
    }
    
    return items
  })
  const layoutSideWidth = computed(() => {
    if (!isCollapse.value) {
      return config.value.layout_side_width
    } else {
      return config.value.layout_side_collapsed_width
    }
  })

  // 处理菜单项，将特定菜单项移动到高级能力下
  const processedMenuItems = computed(() => {
    const originalItems = routerStore.asyncRouters[0]?.children || []
    const processedItems = []
    
    // 调试：打印原始菜单数据
    console.log('原始菜单数据:', originalItems.map(item => ({ name: item.name, title: item.meta?.title })))
    
    // 需要隐藏的菜单项名称
    const hiddenMenuNames = ['wlScenes', 'wlEngineRules', 'wlResources']
    
    // 需要添加到高级能力下的子菜单
    const advancedCapabilitiesSubMenus = []
    
    originalItems.forEach(item => {
      if (hiddenMenuNames.includes(item.name)) {
        // 隐藏这些菜单项，但保存它们作为高级能力的子菜单
        advancedCapabilitiesSubMenus.push({
          ...item,
          hidden: false // 确保在子菜单中显示
        })
      } else {
        // 其他菜单项正常显示
        processedItems.push(item)
      }
    })
    
    // 如果找到了子菜单，创建高级能力菜单
    if (advancedCapabilitiesSubMenus.length > 0) {
      const advancedCapabilitiesMenu = {
        name: 'advancedCapabilities',
        path: 'advancedCapabilities',
        component: 'view/routerHolder.vue',
        meta: {
          title: '高级能力',
          icon: 'cloud'
        },
        hidden: false,
        children: advancedCapabilitiesSubMenus
      }
      
      // 将高级能力菜单插入到合适的位置
      const insertIndex = Math.min(2, processedItems.length) // 插入到前3个位置
      processedItems.splice(insertIndex, 0, advancedCapabilitiesMenu)
    }
    
    // 调试：打印处理后的菜单数据
    console.log('处理后的菜单数据:', processedItems.map(item => ({ 
      name: item.name, 
      title: item.meta?.title,
      hasChildren: item.children && item.children.length > 0
    })))
    
    return processedItems
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
