<template>
  <router-view />
</template>

<script setup>
// 路由占位组件，用于有子路由的父级菜单
defineOptions({
  name: 'RouterHolder'
})
</script>