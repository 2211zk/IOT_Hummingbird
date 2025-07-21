<template>
  <div
    class="relative h-full bg-white text-slate-700 dark:text-slate-300 dark:bg-slate-900 border-r shadow dark:shadow-gray-700"
    :class="isCollapse ? '' : '  px-2'"
    :style="{
      width: layoutSideWidth + 'px'
    }"
  >
    <el-scrollbar>
      <el-menu
        :collapse="isCollapse"
        :collapse-transition="false"
        :default-active="active"
        class="border-r-0 w-full"
        unique-opened
        @select="selectMenuItem"
      >
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
