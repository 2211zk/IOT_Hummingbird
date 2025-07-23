<template>
  <div class="selected-list">
    <div class="header">
      <h4>已选择品类 ({{ selected.length }})</h4>
      <el-button
        v-if="selected.length > 0"
        type="danger"
        size="small"
        text
        @click="handleClearAll"
      >
        清空所有
      </el-button>
    </div>

    <div class="list-container">
      <div v-if="selected.length === 0" class="empty-state">
        <el-empty description="暂未选择任何品类" :image-size="80" />
      </div>
      
      <div v-else class="selected-items">
        <div
          v-for="item in selected"
          :key="item.id"
          class="selected-item"
        >
          <div class="item-content">
            <div class="item-header">
              <span class="item-name">{{ item.name }}</span>
              <el-button
                type="danger"
                size="small"
                text
                @click="handleRemove(item)"
              >
                <el-icon><Close /></el-icon>
              </el-button>
            </div>
            <div class="item-details">
              <el-tag size="small" type="info">{{ item.code }}</el-tag>
              <span class="category">{{ item.category }}</span>
            </div>
            <div v-if="item.description" class="item-description">
              {{ item.description }}
            </div>
          </div>
        </div>
      </div>
    </div>

    <div v-if="maxSelections && selected.length >= maxSelections" class="limit-warning">
      <el-alert
        title="已达到最大选择数量限制"
        type="warning"
        :closable="false"
        show-icon
      />
    </div>
  </div>
</template>

<script setup>
import { Close } from '@element-plus/icons-vue'
import { ElMessageBox } from 'element-plus'

defineOptions({
  name: 'SelectedList'
})

const props = defineProps({
  selected: {
    type: Array,
    default: () => []
  },
  maxSelections: {
    type: Number,
    default: 0
  }
})

const emits = defineEmits(['remove', 'clear-all'])

// 移除单个品类
const handleRemove = (item) => {
  emits('remove', item)
}

// 清空所有选择
const handleClearAll = async () => {
  try {
    await ElMessageBox.confirm(
      '确定要清空所有已选择的品类吗？',
      '确认清空',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }
    )
    emits('clear-all')
  } catch {
    // 用户取消操作
  }
}
</script>

<style lang="scss" scoped>
.selected-list {
  height: 100%;
  display: flex;
  flex-direction: column;

  .header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding-bottom: 12px;
    border-bottom: 1px solid #ebeef5;
    margin-bottom: 16px;

    h4 {
      margin: 0;
      font-size: 14px;
      font-weight: 500;
      color: #303133;
    }
  }

  .list-container {
    flex: 1;
    overflow-y: auto;

    .empty-state {
      display: flex;
      justify-content: center;
      align-items: center;
      height: 200px;
    }

    .selected-items {
      .selected-item {
        margin-bottom: 12px;
        padding: 12px;
        background: #f8f9fa;
        border-radius: 6px;
        border: 1px solid #e4e7ed;
        transition: all 0.3s;

        &:hover {
          border-color: #409eff;
          box-shadow: 0 2px 4px rgba(64, 158, 255, 0.1);
        }

        .item-content {
          .item-header {
            display: flex;
            justify-content: space-between;
            align-items: center;
            margin-bottom: 8px;

            .item-name {
              font-weight: 500;
              color: #303133;
              font-size: 14px;
            }
          }

          .item-details {
            display: flex;
            align-items: center;
            gap: 8px;
            margin-bottom: 6px;

            .category {
              font-size: 12px;
              color: #909399;
            }
          }

          .item-description {
            font-size: 12px;
            color: #606266;
            line-height: 1.4;
            margin-top: 4px;
          }
        }
      }
    }
  }

  .limit-warning {
    margin-top: 12px;
  }
}

// 响应式设计
@media (max-width: 768px) {
  .selected-list {
    .selected-items {
      .selected-item {
        padding: 8px;
        margin-bottom: 8px;

        .item-content {
          .item-header {
            .item-name {
              font-size: 13px;
            }
          }

          .item-details {
            gap: 6px;
          }
        }
      }
    }
  }
}
</style>