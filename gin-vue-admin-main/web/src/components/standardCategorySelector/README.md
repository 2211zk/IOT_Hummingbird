# 标准品类选择器 (StandardCategorySelector)

一个基于 Vue 3 和 Element Plus 的标准品类选择器组件，支持搜索、过滤、分页和多选功能。

## 特性

- 🔍 **搜索过滤**: 支持按名称、编码、类别进行搜索和过滤
- 📄 **分页支持**: 内置分页功能，支持大量数据展示
- ✅ **多选/单选**: 支持多选和单选模式
- 🎯 **数量限制**: 可设置最大选择数量
- 📱 **响应式设计**: 适配桌面和移动设备
- 🎨 **主题定制**: 支持自定义样式和主题
- ⚡ **性能优化**: 防抖搜索、错误重试、加载状态管理
- 🧪 **完整测试**: 包含单元测试和集成测试

## 安装

```bash
# 确保已安装依赖
npm install vue@^3.0.0 element-plus@^2.3.8 lodash-es
```

## 基础用法

```vue
<template>
  <StandardCategorySelector 
    v-model="selectedCategories"
    @change="handleChange"
  />
</template>

<script setup>
import { ref } from 'vue'
import StandardCategorySelector from '@/components/standardCategorySelector/index.vue'

const selectedCategories = ref([])

const handleChange = (categories) => {
  console.log('选择的品类:', categories)
}
</script>
```

## API

### Props

| 参数 | 说明 | 类型 | 默认值 |
|------|------|------|--------|
| modelValue (v-model) | 绑定值，已选择的品类数组 | Array | [] |
| multiple | 是否支持多选 | Boolean | true |
| placeholder | 按钮显示文本 | String | '选择标准品类' |
| disabled | 是否禁用 | Boolean | false |
| maxSelections | 最大选择数量，0表示无限制 | Number | 0 |
| maxDisplay | 最大显示标签数量 | Number | 3 |

### Events

| 事件名 | 说明 | 回调参数 |
|--------|------|----------|
| update:modelValue | 选择变化时触发 | (categories: Array) |
| change | 选择变化时触发 | (categories: Array) |
| confirm | 确认选择时触发 | (categories: Array) |

### Methods

通过 ref 可以调用以下方法：

| 方法名 | 说明 | 参数 |
|--------|------|------|
| openModal | 打开选择器弹框 | - |
| clearSelection | 清空所有选择 | - |
| getSelectedCategories | 获取当前选择的品类 | - |

### Slots

暂无插槽支持。

## 使用示例

### 单选模式

```vue
<template>
  <StandardCategorySelector 
    v-model="singleCategory"
    :multiple="false"
    placeholder="选择单个品类"
  />
</template>

<script setup>
import { ref } from 'vue'

const singleCategory = ref([])
</script>
```

### 限制选择数量

```vue
<template>
  <StandardCategorySelector 
    v-model="limitedCategories"
    :max-selections="3"
    placeholder="最多选择3个品类"
  />
</template>

<script setup>
import { ref } from 'vue'

const limitedCategories = ref([])
</script>
```

### 预设选择

```vue
<template>
  <StandardCategorySelector 
    v-model="presetCategories"
    placeholder="预设选择示例"
  />
</template>

<script setup>
import { ref } from 'vue'

const presetCategories = ref([
  {
    id: 1,
    name: '电子产品',
    code: 'ELEC001',
    category: '电子设备',
    description: '各类电子产品和设备',
    status: 1
  }
])
</script>
```

### 表单集成

```vue
<template>
  <el-form :model="form" label-width="120px">
    <el-form-item label="产品名称">
      <el-input v-model="form.name" />
    </el-form-item>
    <el-form-item label="标准品类">
      <StandardCategorySelector 
        v-model="form.categories"
        :max-selections="5"
      />
    </el-form-item>
  </el-form>
</template>

<script setup>
import { reactive } from 'vue'

const form = reactive({
  name: '',
  categories: []
})
</script>
```

### 方法调用

```vue
<template>
  <div>
    <el-button @click="openSelector">打开选择器</el-button>
    <el-button @click="clearAll">清空选择</el-button>
    <StandardCategorySelector 
      ref="selectorRef"
      v-model="categories"
    />
  </div>
</template>

<script setup>
import { ref } from 'vue'

const selectorRef = ref(null)
const categories = ref([])

const openSelector = () => {
  selectorRef.value?.openModal()
}

const clearAll = () => {
  selectorRef.value?.clearSelection()
}
</script>
```

## 数据格式

### 品类数据结构

```typescript
interface StandardCategory {
  id: number;           // 品类ID
  name: string;         // 品类名称
  code: string;         // 品类编码
  category: string;     // 所属类别
  description?: string; // 描述信息
  status: number;       // 状态 (1:启用, 0:禁用)
  createdAt: string;    // 创建时间
  updatedAt: string;    // 更新时间
}
```

### API 响应格式

```typescript
interface ApiResponse {
  code: number;
  data: {
    list: StandardCategory[];
    total: number;
    page: number;
    pageSize: number;
  };
  msg: string;
}
```

## 样式定制

### CSS 变量

组件使用以下 CSS 变量，可以通过覆盖这些变量来定制样式：

```css
:root {
  --category-selector-primary-color: #409eff;
  --category-selector-border-color: #dcdfe6;
  --category-selector-text-color: #303133;
  --category-selector-background-color: #ffffff;
}
```

### 自定义样式类

```css
/* 自定义触发按钮样式 */
.standard-category-selector .selector-trigger .el-button {
  border-radius: 6px;
  font-weight: 500;
}

/* 自定义标签样式 */
.standard-category-selector .selected-display .el-tag {
  margin-right: 8px;
  margin-bottom: 4px;
}

/* 自定义弹框样式 */
.category-modal .el-dialog {
  border-radius: 12px;
}
```

## 错误处理

组件内置了完善的错误处理机制：

- **网络错误**: 自动重试机制，支持断网检测
- **API错误**: 统一错误提示，支持自定义错误消息
- **数据验证**: 自动过滤无效数据，防止组件崩溃
- **用户操作**: 友好的操作提示和确认对话框

## 性能优化

- **防抖搜索**: 搜索输入使用 300ms 防抖
- **虚拟滚动**: 大数据量时自动启用虚拟滚动
- **缓存机制**: API 响应数据本地缓存
- **懒加载**: 弹框内容懒加载，提升首屏性能

## 浏览器兼容性

- Chrome >= 60
- Firefox >= 60
- Safari >= 12
- Edge >= 79

## 更新日志

### v1.0.0 (2024-01-20)

- 🎉 初始版本发布
- ✅ 支持基础的选择功能
- ✅ 支持搜索和过滤
- ✅ 支持分页
- ✅ 支持多选/单选模式
- ✅ 支持响应式设计
- ✅ 完整的测试覆盖

## 贡献指南

1. Fork 项目
2. 创建功能分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 打开 Pull Request

## 许可证

[MIT License](LICENSE)

## 支持

如果你觉得这个组件有用，请给项目一个 ⭐️！

如果遇到问题或有功能建议，请提交 [Issue](https://github.com/your-repo/issues)。