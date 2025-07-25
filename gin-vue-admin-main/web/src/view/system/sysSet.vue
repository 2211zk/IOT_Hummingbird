<template>
  <div class="sys-set-root">
    <div class="sys-set-container">
      <div class="sys-set-form">
        <el-form :model="form" label-width="90px" size="large">
          <el-form-item label="系统名称">
            <el-input v-model="form.name" placeholder="请输入系统名称" />
          </el-form-item>
          <el-form-item label="copyright">
            <el-input v-model="form.copyright" placeholder="请输入版权信息" />
          </el-form-item>
          <el-form-item label="系统logo">
            <el-upload
              class="avatar-uploader"
              action="#"
              :show-file-list="false"
              :on-change="handleLogoChange"
              :before-upload="beforeUpload"
            >
              <img v-if="form.logoUrl" :src="form.logoUrl" class="avatar" />
              <i v-else class="el-icon-plus avatar-uploader-icon"></i>
            </el-upload>
            <div class="tip">推荐尺寸: 200×200</div>
          </el-form-item>
          <el-form-item label="浏览器标签">
            <el-upload
              class="avatar-uploader"
              action="#"
              :show-file-list="false"
              :on-change="handleIconChange"
              :before-upload="beforeUpload"
            >
              <img v-if="form.iconUrl" :src="form.iconUrl" class="avatar" />
              <i v-else class="el-icon-plus avatar-uploader-icon"></i>
            </el-upload>
            <div class="tip">推荐尺寸: 64×64</div>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="onSubmit">保存</el-button>
          </el-form-item>
        </el-form>
      </div>
      <div class="sys-set-visual">
        <!-- 科技感SVG+蜂鸟logo -->
        <svg width="400" height="400" viewBox="0 0 400 400">
          <defs>
            <radialGradient id="bg" cx="50%" cy="50%" r="50%">
              <stop offset="0%" stop-color="#2d3a4b" />
              <stop offset="100%" stop-color="#1a2233" />
            </radialGradient>
            <filter id="glow" x="-50%" y="-50%" width="200%" height="200%">
              <feGaussianBlur stdDeviation="8" result="coloredBlur"/>
              <feMerge>
                <feMergeNode in="coloredBlur"/>
                <feMergeNode in="SourceGraphic"/>
              </feMerge>
            </filter>
          </defs>
          <circle cx="200" cy="200" r="160" fill="url(#bg)" filter="url(#glow)" />
          <circle cx="200" cy="200" r="120" fill="none" stroke="#3fa7ff" stroke-width="3" opacity="0.5" />
          <circle cx="200" cy="200" r="80" fill="none" stroke="#3fa7ff" stroke-width="2" opacity="0.3" />
          <circle cx="200" cy="200" r="40" fill="none" stroke="#3fa7ff" stroke-width="1.5" opacity="0.2" />
          <!-- 蜂鸟logo -->
          <image x="150" y="120" width="100" height="100" :href="hummingbirdLogo" />
          <!-- 协议标签 -->
          <g>
            <rect x="270" y="110" rx="12" width="60" height="32" fill="#3fa7ff" opacity="0.8" />
            <text x="300" y="132" fill="#fff" font-size="18" text-anchor="middle">MQTT</text>
            <rect x="290" y="200" rx="12" width="60" height="32" fill="#2ed0ff" opacity="0.8" />
            <text x="320" y="222" fill="#fff" font-size="18" text-anchor="middle">TCP</text>
            <rect x="70" y="220" rx="12" width="70" height="32" fill="#4be3a5" opacity="0.8" />
            <text x="105" y="242" fill="#fff" font-size="18" text-anchor="middle">Modbus</text>
            <rect x="180" y="300" rx="12" width="70" height="32" fill="#7a7fff" opacity="0.8" />
            <text x="215" y="322" fill="#fff" font-size="18" text-anchor="middle">HTTP</text>
          </g>
        </svg>
        <div class="visual-tip">推荐尺寸: 800×600</div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import hummingbirdLogo from '@/assets/logo_login.png'

const form = ref({
  name: '蜂鸟物联网平台',
  copyright: 'Copyright © 2021-2025',
  logoUrl: hummingbirdLogo,
  iconUrl: hummingbirdLogo
})

const beforeUpload = (file) => {
  const isImg = file.type.startsWith('image/')
  if (!isImg) {
    ElMessage.error('只能上传图片文件！')
  }
  return isImg
}
const handleLogoChange = (file) => {
  const reader = new FileReader()
  reader.onload = (e) => {
    form.value.logoUrl = e.target.result
  }
  reader.readAsDataURL(file.raw)
}
const handleIconChange = (file) => {
  const reader = new FileReader()
  reader.onload = (e) => {
    form.value.iconUrl = e.target.result
  }
  reader.readAsDataURL(file.raw)
}
const onSubmit = () => {
  ElMessage.success('保存成功（仅前端演示）')
}
</script>

<style scoped>
.sys-set-root {
  min-height: 100vh;
  background: linear-gradient(135deg, #232c3b 0%, #1a2233 100%);
  padding: 40px 0;
}
.sys-set-container {
  max-width: 1200px;
  margin: 0 auto;
  display: flex;
  gap: 48px;
  background: rgba(30,40,60,0.95);
  border-radius: 18px;
  box-shadow: 0 4px 32px 0 rgba(0,0,0,0.18);
  padding: 48px 32px;
}
.sys-set-form {
  flex: 1;
  min-width: 340px;
}
.sys-set-visual {
  flex: 1.2;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  min-width: 400px;
}
.avatar-uploader {
  display: inline-block;
  width: 100px;
  height: 100px;
  border: 1px dashed #3fa7ff;
  border-radius: 8px;
  overflow: hidden;
  background: #232c3b;
  cursor: pointer;
  position: relative;
}
.avatar-uploader-icon {
  font-size: 32px;
  color: #3fa7ff;
  line-height: 100px;
  text-align: center;
  width: 100%;
  height: 100%;
  display: block;
}
.avatar {
  width: 100px;
  height: 100px;
  display: block;
  object-fit: contain;
}
.tip {
  color: #aaa;
  font-size: 12px;
  margin-top: 4px;
}
.visual-tip {
  color: #aaa;
  font-size: 13px;
  margin-top: 12px;
  text-align: center;
}
.el-form {
  background: transparent;
}
.el-form-item__label {
  color: #e0e6f0;
}
.el-input__wrapper {
  background: #232c3b;
  border: 1px solid #3fa7ff;
  border-radius: 6px;
}
.el-input__inner {
  color: #fff;
}
.el-button--primary {
  background: linear-gradient(90deg, #3fa7ff 0%, #2ed0ff 100%);
  border: none;
}
</style> 