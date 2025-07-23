import { asyncRouterHandle } from '@/utils/asyncRouter'
import { emitter } from '@/utils/bus.js'
import { asyncMenu } from '@/api/menu'
import { defineStore } from 'pinia'
import { ref, watchEffect } from 'vue'
import pathInfo from '@/pathInfo.json'
import {useRoute} from "vue-router";

const notLayoutRouterArr = []
const keepAliveRoutersArr = []
const nameMap = {}

/**
 * 处理菜单结构，创建层级关系
 * 注意：此函数目前未被使用，菜单处理逻辑已移至 normalMode.vue 中
 * 保留此函数以备将来需要统一处理菜单结构时使用
 */
const processMenuStructure = (menus) => {
  console.log('处理菜单结构:', menus)
  if (!menus || !Array.isArray(menus)) return menus
  
  // 查找设备接入菜单
  const deviceAccessMenu = menus.find(item => item.name === 'wl_playform')
  const productMenu = menus.find(item => item.name === 'wlProducts')
  const equipmentMenu = menus.find(item => item.name === 'wlEquipment')
  
  console.log('找到的菜单:', { deviceAccessMenu, productMenu, equipmentMenu })
  
  if (deviceAccessMenu && productMenu && equipmentMenu) {
    console.log('开始处理菜单层级结构')
    // 确保设备接入菜单有正确的图标和标题
    deviceAccessMenu.meta = {
      ...deviceAccessMenu.meta,
      title: '设备接入',
      icon: 'connection'
    }
    
    // 创建设备接入的子菜单
    deviceAccessMenu.children = [
      {
        ...productMenu,
        hidden: false,
        meta: { ...productMenu.meta, title: '产品管理', icon: 'box' }
      },
      {
        ...equipmentMenu,
        hidden: false,
        meta: { ...equipmentMenu.meta, title: '设备管理', icon: 'monitor' }
      },
      // 设备地图已移至独立的设备分布菜单中，此处移除
      // {
      //   name: 'deviceMap',
      //   path: 'deviceMap',
      //   component: 'wl_playform/deviceMap/index',
      //   meta: { 
      //     title: '设备地图', 
      //     icon: 'location',
      //     keepAlive: true
      //   },
      //   hidden: false
      // }
    ]
    
    // 隐藏原来的产品管理和设备管理菜单
    const filteredMenus = menus.filter(item => 
      item.name !== 'wlProducts' && item.name !== 'wlEquipment'
    )
    
    console.log('处理后的菜单:', filteredMenus)
    return filteredMenus
  }
  
  return menus
}

const formatRouter = (routes, routeMap, parent) => {
  routes &&
    routes.forEach((item) => {
      item.parent = parent
      item.meta.btns = item.btns
      item.meta.hidden = item.hidden
      if (item.meta.defaultMenu === true) {
        if (!parent) {
          item = { ...item, path: `/${item.path}` }
          notLayoutRouterArr.push(item)
        }
      }
      routeMap[item.name] = item
      if (item.children && item.children.length > 0) {
        formatRouter(item.children, routeMap, item)
      }
    })
}

const KeepAliveFilter = (routes) => {
  routes &&
    routes.forEach((item) => {
      // 子菜单中有 keep-alive 的，父菜单也必须 keep-alive，否则无效。这里将子菜单中有 keep-alive 的父菜单也加入。
      if (
        (item.children && item.children.some((ch) => ch.meta.keepAlive)) ||
        item.meta.keepAlive
      ) {
        const path = item.meta.path
        keepAliveRoutersArr.push(pathInfo[path])
        nameMap[item.name] = pathInfo[path]
      }
      if (item.children && item.children.length > 0) {
        KeepAliveFilter(item.children)
      }
    })
}

export const useRouterStore = defineStore('router', () => {
  const keepAliveRouters = ref([])
  const asyncRouterFlag = ref(0)
  const setKeepAliveRouters = (history) => {
    const keepArrTemp = []
    history.forEach((item) => {
      if (nameMap[item.name]) {
        keepArrTemp.push(nameMap[item.name])
      }
    })
    keepAliveRouters.value = Array.from(new Set(keepArrTemp))
  }

  const route = useRoute()

  emitter.on('setKeepAlive', setKeepAliveRouters)

  const asyncRouters = ref([])

  const topMenu = ref([])

  const leftMenu = ref([])

  const menuMap = {}

  const topActive = ref('')

  const setLeftMenu = (name) => {
    sessionStorage.setItem('topActive', name)
    topActive.value = name
    leftMenu.value = []
    if (menuMap[name]?.children) {
      leftMenu.value = menuMap[name].children
    }
    return menuMap[name]?.children
  }

  const findTopActive = (menuMap, routeName) => {
    for (let topName in menuMap) {
      const topItem = menuMap[topName];
      if (topItem.children?.some(item => item.name === routeName)) {
        return topName;
      }
      const foundName = findTopActive(topItem.children || {}, routeName);
      if (foundName) {
        return topName;
      }
    }
    return null;
  };

  watchEffect(() => {
    let topActive = sessionStorage.getItem('topActive')
    // 初始化菜单内容，防止重复添加
    topMenu.value = [];
    asyncRouters.value[0]?.children.forEach((item) => {
      if (item.hidden) return
      menuMap[item.name] = item
      topMenu.value.push({ ...item, children: [] })
    })
    if (!topActive || topActive === 'undefined' || topActive === 'null') {
      topActive = findTopActive(menuMap, route.name);
    }
    setLeftMenu(topActive)
  })

  const routeMap = ref({})
  // 从后台获取动态路由
  const SetAsyncRouter = async () => {
    asyncRouterFlag.value++
    const baseRouter = [
      {
        path: '/layout',
        name: 'layout',
        component: 'view/layout/index.vue',
        meta: {
          title: '底层layout'
        },
        children: []
      }
    ]
    const asyncRouterRes = await asyncMenu()
    const asyncRouter = asyncRouterRes.data.menus
    
    // 处理菜单结构，创建层级关系
    const processedAsyncRouter = processMenuStructure(asyncRouter)
    
    // 添加设备分布菜单到路由系统
    const deviceDistributionMenu = {
      name: 'deviceDistribution',
      path: 'deviceDistribution',
      component: 'view/wl_playform/deviceDistribution/index.vue',
      meta: { 
        title: '设备分布', 
        icon: 'location',
        keepAlive: true
      },
      hidden: false,
      children: [
        {
          name: 'deviceMap',
          path: 'deviceMap',
          component: 'view/wl_playform/deviceMap/index.vue',
          meta: { 
            title: '设备地图', 
            icon: 'location',
            keepAlive: true
          },
          hidden: false
        }
      ]
    }
    
    // 确保设备分布菜单被正确添加到路由系统
    console.log('添加设备分布菜单到路由系统:', deviceDistributionMenu)
    
    // 将设备分布菜单添加到路由系统
    processedAsyncRouter.push(deviceDistributionMenu)
    
    processedAsyncRouter &&
      processedAsyncRouter.push({
        path: 'reload',
        name: 'Reload',
        hidden: true,
        meta: {
          title: '',
          closeTab: true
        },
        component: 'view/error/reload.vue'
      })
    
    // 格式化路由并填充routeMap
    formatRouter(processedAsyncRouter, routeMap.value)
    baseRouter[0].children = processedAsyncRouter
    if (notLayoutRouterArr.length !== 0) {
      baseRouter.push(...notLayoutRouterArr)
    }
    asyncRouterHandle(baseRouter)
    KeepAliveFilter(processedAsyncRouter) // 使用processedAsyncRouter而不是asyncRouter
    asyncRouters.value = baseRouter
    console.log('路由配置完成:', baseRouter)
    return true
  }

  return {
    topActive,
    setLeftMenu,
    topMenu,
    leftMenu,
    asyncRouters,
    keepAliveRouters,
    asyncRouterFlag,
    SetAsyncRouter,
    routeMap
  }
})
