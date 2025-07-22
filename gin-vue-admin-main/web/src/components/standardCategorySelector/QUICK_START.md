# 快速开始指南

本指南将帮助你快速上手标准品类选择器组件。

## 1. 安装和导入

### 1.1 确保依赖已安装

```bash
npm install vue@^3.0.0 element-plus@^2.3.8 lodash-es
```

### 1.2 导入组件

```javascript
// 在需要使用的页面中导入
import StandardCategorySelector from '@/components/standardCategorySelector/index.vue'
```

### 1.3 全局注册（可选）

```javascript
// main.js
import { createApp } from 'vue'
import StandardCategorySelector from '@/components/standardCategorySelector/index.vue'

const app = createApp(App)
app.component('StandardCategorySelector', StandardCategorySelector)
```

## 2. 基础使用

### 2.1 最简单的用法

```vue
<template>
  <div>
    <StandardCategorySelector v-model="categories" />
    <p>已选择: {{ categories.length }} 个品类</p>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import StandardCategorySelector from '@/components/standardCategorySelector/index.vue'

const categories = ref([])
</script>
```

### 2.2 监听选择变化

```vue
<template>
  <StandardCategorySelector 
    v-model="categories"
    @change="handleChange"
  />
</template>

<script setup>
import { ref } from 'vue'

const categories = ref([])

const handleChange = (selectedCategories) => {
  console.log('选择变化:', selectedCategories)
  // 处理选择变化的逻辑
}
</script>
```

## 3. 常用配置

### 3.1 单选模式

```vue
<template>
  <StandardCategorySelector 
    v-model="singleCategory"
    :multiple="false"
    placeholder="选择一个品类"
  />
</template>

<script setup>
import { ref } from 'vue'

const singleCategory = ref([])
</script>
```

### 3.2 限制选择数量

```vue
<template>
  <StandardCategorySelector 
    v-model="categories"
    :max-selections="3"
    placeholder="最多选择3个品类"
  />
</template>

<script setup>
import { ref } from 'vue'

const categories = ref([])
</script>
```

### 3.3 自定义按钮文本

```vue
<template>
  <StandardCategorySelector 
    v-model="categories"
    placeholder="选择产品品类"
  />
</template>
```

## 4. 表单集成

### 4.1 在 Element Plus 表单中使用

```vue
<template>
  <el-form :model="form" :rules="rules" ref="formRef">
    <el-form-item label="产品名称" prop="name">
      <el-input v-model="form.name" />
    </el-form-item>
    
    <el-form-item label="产品品类" prop="categories">
      <StandardCategorySelector 
        v-model="form.categories"
        :max-selections="5"
      />
    </el-form-item>
    
    <el-form-item>
      <el-button type="primary" @click="submitForm">提交</el-button>
    </el-form-item>
  </el-form>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { ElMessage } from 'element-plus'

const formRef = ref()

const form = reactive({
  name: '',
  categories: []
})

const rules = {
  name: [
    { required: true, message: '请输入产品名称', trigger: 'blur' }
  ],
  categories: [
    { 
      required: true, 
      validator: (rule, value, callback) => {
        if (!value || value.length === 0) {
          callback(new Error('请选择至少一个品类'))
        } else {
          callback()
        }
      },
      trigger: 'change'
    }
  ]
}

const submitForm = () => {
  formRef.value.validate((valid) => {
    if (valid) {
      console.log('提交数据:', form)
      ElMessage.success('提交成功！')
    }
  })
}
</script>
```

## 5. 高级用法

### 5.1 使用 ref 调用方法

```vue
<template>
  <div>
    <el-button @click="openSelector">打开选择器</el-button>
    <el-button @click="clearAll">清空选择</el-button>
    <el-button @click="getCurrentSelection">获取当前选择</el-button>
    
    <StandardCategorySelector 
      ref="selectorRef"
      v-model="categories"
    />
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { ElMessage } from 'element-plus'

const selectorRef = ref()
const categories = ref([])

const openSelector = () => {
  selectorRef.value?.openModal()
}

const clearAll = () => {
  selectorRef.value?.clearSelection()
  ElMessage.success('已清空选择')
}

const getCurrentSelection = () => {
  const selection = selectorRef.value?.getSelectedCategories()
  ElMessage.info(`当前选择了 ${selection?.length || 0} 个品类`)
}
</script>
```

### 5.2 预设选择

```vue
<template>
  <div>
    <el-button @click="setPreset">设置预设</el-button>
    <StandardCategorySelector v-model="categories" />
  </div>
</template>

<script setup>
import { ref } from 'vue'

const categories = ref([])

const setPreset = () => {
  categories.value = [
    {
      id: 1,
      name: '电子产品',
      code: 'ELEC001',
      category: '电子设备',
      description: '各类电子产品和设备',
      status: 1
    }
  ]
}
</script>
```

## 6. 样式定制

### 6.1 自定义 CSS

```vue
<template>
  <StandardCategorySelector 
    v-model="categories"
    class="custom-selector"
  />
</template>

<style scoped>
.custom-selector :deep(.selector-trigger .el-button) {
  background: linear-gradient(45deg, #409eff, #67c23a);
  border: none;
  color: white;
  border-radius: 20px;
}

.custom-selector :deep(.selected-display .el-tag) {
  border-radius: 12px;
  margin-right: 8px;
}
</style>
```

### 6.2 使用 CSS 变量

```vue
<template>
  <div class="themed-container">
    <StandardCategorySelector v-model="categories" />
  </div>
</template>

<style scoped>
.themed-container {
  --category-selector-primary-color: #e74c3c;
  --category-selector-border-color: #bdc3c7;
}
</style>
```

## 7. 错误处理

### 7.1 监听错误

```vue
<template>
  <StandardCategorySelector 
    v-model="categories"
    @error="handleError"
  />
</template>

<script setup>
import { ref } from 'vue'
import { ElMessage } from 'element-plus'

const categories = ref([])

const handleError = (error) => {
  console.error('选择器错误:', error)
  ElMessage.error('操作失败，请重试')
}
</script>
```

## 8. 常见问题

### 8.1 数据不显示

**问题**: 打开选择器后没有数据显示

**解决方案**:
1. 检查后端 API 是否正常运行
2. 检查 API 路径是否正确
3. 查看浏览器控制台是否有错误信息

### 8.2 选择后数据格式不正确

**问题**: 选择的数据格式与预期不符

**解决方案**:
```javascript
// 确保数据格式正确
const categories = ref([
  {
    id: 1,                    // 必需
    name: '品类名称',          // 必需
    code: 'CODE001',          // 必需
    category: '类别',         // 必需
    description: '描述',      // 可选
    status: 1                 // 必需
  }
])
```

### 8.3 样式冲突

**问题**: 组件样式与项目样式冲突

**解决方案**:
```css
/* 使用更具体的选择器 */
.my-page .standard-category-selector {
  /* 自定义样式 */
}

/* 或使用 :deep() 穿透样式 */
.my-page :deep(.standard-category-selector) {
  /* 自定义样式 */
}
```

## 9. 性能优化建议

### 9.1 大数据量处理

```vue
<template>
  <!-- 对于大量数据，考虑设置合理的分页大小 -->
  <StandardCategorySelector 
    v-model="categories"
    :page-size="20"
  />
</template>
```

### 9.2 防抖优化

组件内置了防抖功能，但你也可以在外部进行额外优化：

```javascript
import { debounce } from 'lodash-es'

const debouncedHandler = debounce((categories) => {
  // 处理选择变化
}, 500)
```

## 10. 下一步

- 查看 [完整 API 文档](./README.md)
- 查看 [示例页面](../../view/example/standardCategorySelector/index.vue)
- 运行测试: `npm run test:unit`

如果你有任何问题，请查看 [常见问题](./FAQ.md) 或提交 [Issue](https://github.com/your-repo/issues)。