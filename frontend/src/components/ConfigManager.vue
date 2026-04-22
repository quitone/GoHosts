<template>
  <div class="config-manager">
    <n-tabs v-model:value="activeTab" type="line" class="content-tabs">
      <n-tab-pane name="editor" tab="方案编辑">
        <div v-if="editingScheme" class="editor-panel">
          <div class="editor-toolbar">
            <n-input
              v-model:value="editingScheme.name"
              placeholder="方案名称"
              class="name-input"
            />
            <div class="toolbar-actions">
              <n-button :loading="store.saving" type="primary" @click="handleSaveScheme">
                保存
              </n-button>
            </div>
          </div>
          <div class="editor-body">
            <textarea
              v-model="editingScheme.content"
              class="hosts-textarea"
              placeholder="# 在此输入 hosts 内容&#10;127.0.0.1 localhost"
              spellcheck="false"
            />
          </div>
          <div class="editor-footer">
            <span class="status-text">{{ store.saving ? '保存中...' : '就绪' }}</span>
            <span v-if="store.error" class="error-text">{{ store.error }}</span>
          </div>
        </div>
        <div v-else class="empty-state">
          <p>请从左侧选择一个方案进行编辑</p>
        </div>
      </n-tab-pane>

      <n-tab-pane name="system" tab="系统 Hosts">
        <HostsView />
      </n-tab-pane>
    </n-tabs>
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import { useHostsStore } from '@/stores/hosts'
import type { Scheme } from '@/types/schema'
import HostsView from './HostsView.vue'
import { NTabs, NTabPane, NButton, NInput } from 'naive-ui'

const store = useHostsStore()
const activeTab = ref('editor')
const editingScheme = ref<Scheme | null>(null)

watch(
  () => store.selectedScheme,
  (scheme) => {
    if (scheme) {
      editingScheme.value = JSON.parse(JSON.stringify(scheme))
    } else {
      editingScheme.value = null
    }
  },
  { immediate: true }
)

async function handleSaveScheme() {
  if (!editingScheme.value) return
  await store.saveScheme(editingScheme.value)
  const updated = store.schemes.find((s) => s.id === editingScheme.value?.id)
  if (updated) {
    editingScheme.value = JSON.parse(JSON.stringify(updated))
  }
}
</script>

<style scoped>
.config-manager {
  height: 100%;
  display: flex;
  flex-direction: column;
}

.content-tabs {
  flex: 1;
  display: flex;
  flex-direction: column;
}

.content-tabs :deep(.n-tabs-nav) {
  padding: 0 16px;
}

.content-tabs :deep(.n-tabs-pane-wrapper) {
  flex: 1;
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.content-tabs :deep(.n-tab-pane) {
  height: 100%;
  display: flex;
  flex-direction: column;
}

.editor-panel {
  display: flex;
  flex-direction: column;
  height: 100%;
  padding: 0 16px 16px;
}

.editor-toolbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 12px;
  margin-bottom: 12px;
  flex-shrink: 0;
}

.name-input {
  max-width: 300px;
}

.editor-body {
  flex: 1;
  overflow: hidden;
  min-height: 0;
}

.hosts-textarea {
  width: 100%;
  height: 100%;
  background: #1e293b;
  color: #e2e8f0;
  border: 1px solid #334155;
  border-radius: 6px;
  padding: 12px;
  font-family: 'Fira Code', 'Consolas', 'Monaco', 'Courier New', monospace;
  font-size: 14px;
  line-height: 1.6;
  resize: none;
  outline: none;
  box-sizing: border-box;
}

.hosts-textarea:focus {
  border-color: #4a9eff;
}

.editor-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-top: 12px;
  font-size: 13px;
  flex-shrink: 0;
}

.status-text {
  color: #94a3b8;
}

.error-text {
  color: #ef4444;
}

.empty-state {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100%;
  color: #94a3b8;
  font-size: 16px;
}
</style>
