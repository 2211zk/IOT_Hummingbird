# 设计文档

## 概述

本设计文档描述了动态表单引擎规则功能的技术实现方案。该功能将改进现有的引擎规则表单，通过渐进式表单显示提升用户体验。系统将基于Vue 3 + Element Plus技术栈，实现响应式的动态表单交互。

## 架构

### 技术栈
- **前端框架**: Vue 3 (Composition API)
- **UI组件库**: Element Plus
- **状态管理**: Pinia
- **HTTP客户端**: Axios
- **构建工具**: Vite

### 组件架构
```
WlEngineRules (主页面)
├── DynamicEngineRuleForm (动态表单组件)
│   ├── RuleNameInput (规则名称输入)
│   ├── FilterConditionGroup (过滤条件组)
│   ├── MessageSourceSelect (消息源选择)
│   ├── ConversionMethodRadio (转换方法单选)
│   └── ResourceSelect (资源选择)
└── ExistingTableComponents (现有表格组件)
```

## 组件和接口

### 1. DynamicEngineRuleForm 组件

**职责**: 管理动态表单的显示逻辑和数据流

**Props**:
```typescript
interface Props {
  modelValue: EngineRuleFormData
  mode: 'create' | 'edit'
}
```

**Events**:
```typescript
interface Events {
  'update:modelValue': (value: EngineRuleFormData) => void
  'submit': (data: EngineRuleFormData) => void
  'reset': () => void
}
```

**核心状态**:
```typescript
interface FormState {
  showFilterCondition: boolean
  showMessageSource: boolean  
  showConversionMethod: boolean
  showResourceSelect: boolean
  messageSourceOptions: SelectOption[]
  conversionMethodOptions: RadioOption[]
  resourceOptions: SelectOption[]
  loading: {
    messageSource: boolean
    resources: boolean
  }
}
```

### 2. 表单字段组件

#### RuleNameInput
- 监听输入变化，触发后续字段显示
- 实现防抖处理，避免频繁触发

#### FilterConditionGroup  
- 包含"查询字段"和"条件"两个输入框
- 支持动态验证规则

#### MessageSourceSelect
- 组合输入框：支持下拉选择 + 手动输入
- 实现搜索过滤功能
- 支持自定义选项添加

#### ConversionMethodRadio
- 单选框组件，提供预定义转换方法
- 选项：HTTP请求、数据库查询、文件处理、数据转换等

#### ResourceSelect
- 下拉选择框，支持搜索过滤
- 根据转换方法动态加载相关资源
- 支持分页加载大量数据

## 数据模型

### EngineRuleFormData
```typescript
interface EngineRuleFormData {
  ruleName: string
  ruleDescription: string
  queryField: string
  condition: string
  sqlStatement: string
  messageSource: string
  forwardingMethod: string
  resourceId: number | undefined
}
```

### SelectOption
```typescript
interface SelectOption {
  label: string
  value: string | number
  disabled?: boolean
  group?: string
}
```

### ConversionMethodOption
```typescript
interface ConversionMethodOption {
  label: string
  value: string
  description: string
  compatibleResources: string[]
}
```

## 错误处理

### 表单验证错误
- 实时验证：字段失焦时进行验证
- 提交验证：表单提交前进行完整验证
- 错误显示：使用Element Plus的表单验证机制

### API请求错误
- 网络错误：显示重试机制
- 数据加载失败：显示错误提示和刷新按钮
- 超时处理：设置合理的请求超时时间

### 数据一致性错误
- 依赖字段验证：上级字段变更时验证下级字段
- 兼容性检查：转换方法与资源的兼容性验证

## 测试策略

### 单元测试
- 组件渲染测试
- 表单验证逻辑测试
- 数据转换函数测试
- 事件处理函数测试

### 集成测试
- 表单交互流程测试
- API调用集成测试
- 组件间通信测试

### 端到端测试
- 完整的表单填写流程
- 错误场景处理
- 数据持久化验证

### 性能测试
- 大量选项数据的渲染性能
- 搜索过滤的响应时间
- 内存使用情况监控

## 实现细节

### 动态显示逻辑
```typescript
// 监听规则名称变化
watch(() => formData.ruleName, (newValue) => {
  showFilterCondition.value = !!newValue.trim()
  if (!newValue.trim()) {
    resetSubsequentFields()
  }
})

// 监听消息源变化
watch(() => formData.messageSource, (newValue) => {
  showConversionMethod.value = !!newValue
  if (!newValue) {
    resetConversionAndResource()
  }
})
```

### 数据加载策略
- 消息源选项：页面初始化时加载
- 资源选项：根据转换方法动态加载
- 搜索结果：实现防抖搜索，减少API调用

### 状态管理
- 使用Vue 3的响应式系统管理表单状态
- 利用computed属性处理派生状态
- 通过watch监听关键字段变化

### 动画效果
- 使用Vue的Transition组件实现平滑过渡
- CSS动画：淡入淡出效果
- 加载状态：骨架屏或加载指示器