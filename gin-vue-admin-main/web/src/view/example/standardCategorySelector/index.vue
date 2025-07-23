<template>
  <div class="standard-category-selector-demo">
    <div class="demo-header">
      <h1>标准品类选择器示例</h1>
      <p>这个页面展示了标准品类选择器组件的各种使用方式和配置选项。</p>
    </div>

    <div class="demo-section">
      <h2>基础用法</h2>
      <p>最简单的用法，支持多选。</p>
      <div class="demo-block">
        <StandardCategorySelector 
          v-model="basicSelection"
          @change="handleBasicChange"
        />
        <div class="result-display">
          <h4>选择结果：</h4>
          <pre>{{ JSON.stringify(basicSelection, null, 2) }}</pre>
        </div>
      </div>
    </div>

    <div class="demo-section">
      <h2>单选模式</h2>
      <p>设置 multiple 为 false 启用单选模式。</p>
      <div class="demo-block">
        <StandardCategorySelector 
          v-model="singleSelection"
          :multiple="false"
          placeholder="选择单个品类"
          @change="handleSingleChange"
        />
        <div class="result-display">
          <h4>选择结果：</h4>
          <pre>{{ JSON.stringify(singleSelection, null, 2) }}</pre>
        </div>
      </div>
    </div>

    <div class="demo-section">
      <h2>限制选择数量</h2>
      <p>通过 maxSelections 属性限制最大选择数量。</p>
      <div class="demo-block">
        <StandardCategorySelector 
          v-model="limitedSelection"
          :max-selections="3"
          placeholder="最多选择3个品类"
          @change="handleLimitedChange"
        />
        <div class="result-display">
          <h4>选择结果（最多3个）：</h4>
          <pre>{{ JSON.stringify(limitedSelection, null, 2) }}</pre>
        </div>
      </div>
    </div>

    <div class="demo-section">
      <h2>自定义显示数量</h2>
      <p>通过 maxDisplay 属性控制显示的标签数量。</p>
      <div class="demo-block">
        <StandardCategorySelector 
          v-model="displaySelection"
          :max-display="2"
          placeholder="自定义显示数量"
          @change="handleDisplayChange"
        />
        <div class="result-display">
          <h4>选择结果：</h4>
          <pre>{{ JSON.stringify(displaySelection, null, 2) }}</pre>
        </div>
      </div>
    </div>

    <div class="demo-section">
      <h2>禁用状态</h2>
      <p>通过 disabled 属性禁用组件。</p>
      <div class="demo-block">
        <StandardCategorySelector 
          v-model="disabledSelection"
          :disabled="true"
          placeholder="禁用状态"
        />
        <el-button @click="toggleDisabled">
          {{ isDisabled ? '启用' : '禁用' }}
        </el-button>
        <StandardCategorySelector 
          v-model="disabledSelection"
          :disabled="isDisabled"
          placeholder="动态禁用"
        />
      </div>
    </div>

    <div class="demo-section">
      <h2>预设选择</h2>
      <p>通过 v-model 预设已选择的品类。</p>
      <div class="demo-block">
        <el-button @click="setPresetSelection">设置预设选择</el-button>
        <el-button @click="clearPresetSelection">清空选择</el-button>
        <StandardCategorySelector 
          v-model="presetSelection"
          placeholder="预设选择示例"
          @change="handlePresetChange"
        />
        <div class="result-display">
          <h4>选择结果：</h4>
          <pre>{{ JSON.stringify(presetSelection, null, 2) }}</pre>
        </div>
      </div>
    </div>

    <div class="demo-section">
      <h2>事件监听</h2>
      <p>监听组件的各种事件。</p>
      <div class="demo-block">
        <StandardCategorySelector 
          v-model="eventSelection"
          placeholder="事件监听示例"
          @change="handleEventChange"
          @confirm="handleEventConfirm"
        />
        <div class="event-log">
          <h4>事件日志：</h4>
          <div class="log-container">
            <div 
              v-for="(log, index) in eventLogs" 
              :key="index"
              class="log-item"
            >
              <span class="log-time">{{ log.time }}</span>
              <span class="log-event">{{ log.event }}</span>
              <span class="log-data">{{ log.data }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="demo-section">
      <h2>表单集成</h2>
      <p>在表单中使用标准品类选择器。</p>
      <div class="demo-block">
        <el-form :model="formData" label-width="120px">
          <el-form-item label="产品名称">
            <el-input v-model="formData.name" placeholder="请输入产品名称" />
          </el-form-item>
          <el-form-item label="产品描述">
            <el-input 
              v-model="formData.description" 
              type="textarea" 
              placeholder="请输入产品描述" 
            />
          </el-form-item>
          <el-form-item label="标准品类">
            <StandardCategorySelector 
              v-model="formData.categories"
              placeholder="选择产品品类"
              :max-selections="5"
            />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="submitForm">提交</el-button>
            <el-button @click="resetForm">重置</el-button>
          </el-form-item>
        </el-form>
        <div class="result-display">
          <h4>表单数据：</h4>
          <pre>{{ JSON.stringify(formData, null, 2) }}</pre>
        </div>
      </div>
    </div>

    <div class="demo-section">
      <h2>API 方法调用</h2>
      <p>通过 ref 调用组件的方法。</p>
      <div class="demo-block">
        <el-button @click="openModal">打开选择器</el-button>
        <el-button @click="clearSelection">清空选择</el-button>
        <el-button @click="getSelection">获取当前选择</el-button>
        <StandardCategorySelector 
          ref="selectorRef"
          v-model="apiSelection"
          placeholder="API方法示例"
        />
        <div class="result-display">
          <h4>当前选择：</h4>
          <pre>{{ JSON.stringify(apiSelection, null, 2) }}</pre>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { ElMessage } from 'element-plus'
import StandardCategorySelector from '@/components/standardCategorySelector/index.vue'

// 基础用法
const basicSelection = ref([])

// 单选模式
const singleSelection = ref([])

// 限制选择数量
const limitedSelection = ref([])

// 自定义显示数量
const displaySelection = ref([])

// 禁用状态
const disabledSelection = ref([])
const isDisabled = ref(false)

// 预设选择
const presetSelection = ref([])

// 事件监听
const eventSelection = ref([])
const eventLogs = ref([])

// 表单集成
const formData = reactive({
  name: '',
  description: '',
  categories: []
})

// API方法调用
const apiSelection = ref([])
const selectorRef = ref(null)

// 事件处理函数
const handleBasicChange = (categories) => {
  console.log('基础选择变化:', categories)
}

const handleSingleChange = (categories) => {
  console.log('单选变化:', categories)
}

const handleLimitedChange = (categories) => {
  console.log('限制选择变化:', categories)
}

const handleDisplayChange = (categories) => {
  console.log('显示选择变化:', categories)
}

const handlePresetChange = (categories) => {
  console.log('预设选择变化:', categories)
}

const handleEventChange = (categories) => {
  addEventLog('change', `选择了 ${categories.length} 个品类`)
}

const handleEventConfirm = (categories) => {
  addEventLog('confirm', `确认选择了 ${categories.length} 个品类`)
}

// 工具函数
const addEventLog = (event, data) => {
  eventLogs.value.unshift({
    time: new Date().toLocaleTimeString(),
    event,
    data
  })
  
  // 限制日志数量
  if (eventLogs.value.length > 10) {
    eventLogs.value = eventLogs.value.slice(0, 10)
  }
}

const toggleDisabled = () => {
  isDisabled.value = !isDisabled.value
}

const setPresetSelection = () => {
  // 模拟预设数据
  presetSelection.value = [
    {
      id: 1,
      name: '电子产品',
      code: 'ELEC001',
      category: '电子设备',
      description: '各类电子产品和设备',
      status: 1
    },
    {
      id: 2,
      name: '家用电器',
      code: 'HOME001',
      category: '家居用品',
      description: '家庭使用的各类电器产品',
      status: 1
    }
  ]
}

const clearPresetSelection = () => {
  presetSelection.value = []
}

const submitForm = () => {
  ElMessage.success('表单提交成功！')
  console.log('提交的表单数据:', formData)
}

const resetForm = () => {
  formData.name = ''
  formData.description = ''
  formData.categories = []
  ElMessage.info('表单已重置')
}

const openModal = () => {
  if (selectorRef.value) {
    selectorRef.value.openModal()
  }
}

const clearSelection = () => {
  if (selectorRef.value) {
    selectorRef.value.clearSelection()
    ElMessage.success('选择已清空')
  }
}

const getSelection = () => {
  if (selectorRef.value) {
    const selection = selectorRef.value.getSelectedCategories()
    ElMessage.info(`当前选择了 ${selection.length} 个品类`)
    console.log('当前选择:', selection)
  }
}
</script>

<style lang="scss" scoped>
.standard-category-selector-demo {
  padding: 20px;
  max-width: 1200px;
  margin: 0 auto;

  .demo-header {
    margin-bottom: 40px;
    text-align: center;

    h1 {
      color: #303133;
      margin-bottom: 16px;
    }

    p {
      color: #606266;
      font-size: 16px;
    }
  }

  .demo-section {
    margin-bottom: 40px;
    padding: 20px;
    border: 1px solid #ebeef5;
    border-radius: 8px;
    background: #fff;

    h2 {
      color: #303133;
      margin-bottom: 8px;
      font-size: 18px;
    }

    > p {
      color: #606266;
      margin-bottom: 20px;
      font-size: 14px;
    }

    .demo-block {
      .result-display {
        margin-top: 20px;
        padding: 16px;
        background: #f8f9fa;
        border-radius: 4px;
        border: 1px solid #e9ecef;

        h4 {
          margin: 0 0 12px 0;
          color: #495057;
          font-size: 14px;
        }

        pre {
          margin: 0;
          font-size: 12px;
          color: #495057;
          white-space: pre-wrap;
          word-break: break-all;
        }
      }

      .event-log {
        margin-top: 20px;

        h4 {
          margin: 0 0 12px 0;
          color: #495057;
          font-size: 14px;
        }

        .log-container {
          max-height: 200px;
          overflow-y: auto;
          border: 1px solid #e9ecef;
          border-radius: 4px;
          background: #f8f9fa;

          .log-item {
            display: flex;
            padding: 8px 12px;
            border-bottom: 1px solid #e9ecef;
            font-size: 12px;

            &:last-child {
              border-bottom: none;
            }

            .log-time {
              width: 80px;
              color: #6c757d;
              flex-shrink: 0;
            }

            .log-event {
              width: 80px;
              color: #007bff;
              font-weight: 500;
              flex-shrink: 0;
            }

            .log-data {
              color: #495057;
              flex: 1;
            }
          }
        }
      }
    }
  }
}

// 响应式设计
@media (max-width: 768px) {
  .standard-category-selector-demo {
    padding: 10px;

    .demo-section {
      padding: 15px;
      margin-bottom: 20px;

      .demo-block {
        .result-display {
          pre {
            font-size: 11px;
          }
        }

        .event-log {
          .log-container {
            .log-item {
              flex-direction: column;
              gap: 4px;

              .log-time,
              .log-event {
                width: auto;
              }
            }
          }
        }
      }
    }
  }
}
</style>