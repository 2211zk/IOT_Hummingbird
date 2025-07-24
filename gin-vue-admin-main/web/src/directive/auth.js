// 权限按钮展示指令
import { useUserStore } from '@/pinia/modules/user'
import { useRoute } from 'vue-router'
export default {
  install: (app) => {
    const userStore = useUserStore()
    app.directive('auth', {
      // 当被绑定的元素插入到 DOM 中时……
      mounted: function (el, binding) {
        const userInfo = userStore.userInfo
        const route = useRoute()
        
        // 检查是否是按钮权限（btnAuth.xxx格式）
        if (typeof binding.value === 'string' && binding.value.startsWith('btnAuth.')) {
          const btnName = binding.value.replace('btnAuth.', '')
          const btnAuth = route.meta.btns || {}
          
          // 如果按钮权限存在，则显示元素
          if (btnAuth[btnName]) {
            return
          } else {
            el.parentNode.removeChild(el)
            return
          }
        }
        
        // 原有的角色权限检查逻辑
        let type = ''
        switch (Object.prototype.toString.call(binding.value)) {
          case '[object Array]':
            type = 'Array'
            break
          case '[object String]':
            type = 'String'
            break
          case '[object Number]':
            type = 'Number'
            break
          default:
            type = ''
            break
        }
        if (type === '') {
          el.parentNode.removeChild(el)
          return
        }
        const waitUse = binding.value.toString().split(',')
        let flag = waitUse.some((item) => Number(item) === userInfo.authorityId)
        if (binding.modifiers.not) {
          flag = !flag
        }
        if (!flag) {
          el.parentNode.removeChild(el)
        }
      }
    })
  }
}
