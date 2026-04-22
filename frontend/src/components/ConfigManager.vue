<template>
  <div class="h-full flex flex-col">
    <n-tabs v-model:value="activeTab" type="line" class="flex-1 flex flex-col">
      <n-tab-pane name="editor" tab="方案编辑">
        <div v-if="editingScheme" class="flex flex-col h-full px-16px pb-16px">
          <div class="flex-between gap-12px mb-12px flex-shrink-0">
            <n-input
              v-model:value="editingScheme.name"
              placeholder="方案名称"
              class="max-w-300px"
            />
            <div class="toolbar-actions">
              <n-button
                :loading="store.saving"
                type="primary"
                @click="handleSaveScheme"
              >
                保存
              </n-button>
            </div>
          </div>
          <div class="flex-1 overflow-hidden min-h-0">
            <textarea
              v-model="editingScheme.content"
              class="w-full h-full bg-neutral-800 text-neutral-100 b b-neutral-700 rounded-6px p-12px font-mono text-14px leading-1.6 resize-none outline-none box-border focus:b-primary"
              placeholder="# 在此输入 hosts 内容&#10;127.0.0.1 localhost"
              spellcheck="false"
            />
          </div>
          <div
            class="flex-between items-center mt-12px text-13px flex-shrink-0"
          >
            <span class="text-neutral-400">{{
              store.saving ? "保存中..." : "就绪"
            }}</span>
            <span v-if="store.error" class="text-error">{{ store.error }}</span>
          </div>
        </div>
        <div v-else class="flex-center h-full text-neutral-400 text-16px">
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
import { ref, watch } from "vue";
import { useHostsStore } from "@/stores/hosts";
import type { Scheme } from "@/types/schema";
import HostsView from "./HostsView.vue";
import { NTabs, NTabPane, NButton, NInput } from "naive-ui";

const store = useHostsStore();
const activeTab = ref("editor");
const editingScheme = ref<Scheme | null>(null);

watch(
  () => store.selectedScheme,
  (scheme) => {
    if (scheme) {
      editingScheme.value = JSON.parse(JSON.stringify(scheme));
    } else {
      editingScheme.value = null;
    }
  },
  { immediate: true },
);

async function handleSaveScheme() {
  if (!editingScheme.value) return;
  await store.saveScheme(editingScheme.value);
  const updated = store.schemes.find((s) => s.id === editingScheme.value?.id);
  if (updated) {
    editingScheme.value = JSON.parse(JSON.stringify(updated));
  }
}
</script>

<style scoped>
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
</style>
