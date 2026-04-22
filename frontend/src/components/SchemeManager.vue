<template>
  <div class="scheme-manager">
    <div class="sidebar-header">
      <h2>方案管理</h2>
      <n-button
        size="small"
        type="primary"
        @click="showAddModal = true"
        data-test="add-scheme"
      >
        + 新建
      </n-button>
    </div>

    <div class="scheme-list">
      <div
        v-for="scheme in store.schemes"
        :key="scheme.id"
        class="scheme-item"
        :class="{
          active: scheme.id === store.activeSchemeId,
          selected: scheme.id === store.selectedSchemeId,
        }"
        @click="store.selectScheme(scheme.id)"
      >
        <div class="scheme-info">
          <div class="scheme-name-row">
            <span v-if="editingId !== scheme.id" class="scheme-name">{{
              scheme.name
            }}</span>
            <n-input
              v-else
              v-model:value="editingName"
              size="small"
              placeholder="方案名称"
              @blur="finishRename"
              @keyup.enter="finishRename"
            />
            <span v-if="scheme.id === store.activeSchemeId" class="active-badge"
              >激活</span
            >
          </div>
          <div class="scheme-meta">
            <span class="scheme-type">{{ scheme.type }}</span>
            <n-switch
              :value="scheme.enabled"
              size="small"
              @update:value="handleToggle(scheme)"
            />
          </div>
        </div>
        <div class="scheme-actions">
          <n-button text size="tiny" @click.stop="startRename(scheme)"
            >重命名</n-button
          >
          <n-button
            text
            size="tiny"
            @click.stop="store.activateScheme(scheme.id)"
          >
            激活
          </n-button>
          <n-popconfirm @positive-click="store.deleteScheme(scheme.id)">
            <template #trigger>
              <n-button text size="tiny" type="error">删除</n-button>
            </template>
            确定要删除方案「{{ scheme.name }}」吗？
          </n-popconfirm>
        </div>
      </div>

      <div v-if="store.schemes.length === 0" class="empty-tip">
        暂无方案，点击「新建」添加
      </div>
    </div>

    <div class="sidebar-footer">
      <n-button size="small" @click="handleExport">导出配置</n-button>
      <n-button size="small" @click="handleImport">导入配置</n-button>
      <input
        ref="fileInputRef"
        type="file"
        accept=".json"
        style="display: none"
        @change="onFileSelected"
      />
    </div>

    <!-- 添加方案对话框 -->
    <n-modal
      v-model:show="showAddModal"
      title="新建方案"
      preset="card"
      style="width: 420px"
      :bordered="false"
    >
      <n-form label-placement="left" label-width="80">
        <n-form-item label="方案名称">
          <n-input
            v-model:value="newScheme.name"
            placeholder="请输入方案名称"
            data-test="scheme-name"
          />
        </n-form-item>
        <n-form-item label="方案类型">
          <n-input
            v-model:value="newScheme.type"
            placeholder="local 或 remote"
          />
        </n-form-item>
        <n-form-item label="Hosts 内容">
          <n-input
            v-model:value="newScheme.content"
            type="textarea"
            placeholder="# 输入 hosts 内容"
            :rows="6"
            data-test="scheme-content"
          />
        </n-form-item>
      </n-form>
      <template #footer>
        <div style="display: flex; justify-content: flex-end; gap: 8px">
          <n-button @click="showAddModal = false">取消</n-button>
          <n-button
            type="primary"
            @click="handleAddScheme"
            data-test="save-scheme"
          >
            保存
          </n-button>
        </div>
      </template>
    </n-modal>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useHostsStore } from "@/stores/hosts";
import type { Scheme } from "@/types/schema";
import {
  NButton,
  NSwitch,
  NModal,
  NForm,
  NFormItem,
  NInput,
  NPopconfirm,
} from "naive-ui";

const store = useHostsStore();

const showAddModal = ref(false);
const newScheme = ref<Partial<Scheme>>({
  name: "",
  type: "local",
  content: "",
});

const editingId = ref<string | null>(null);
const editingName = ref("");
const fileInputRef = ref<HTMLInputElement | null>(null);

onMounted(() => {
  if (store.schemes.length === 0) {
    store.loadConfig();
  }
});

function handleAddScheme() {
  if (!newScheme.value.name?.trim()) return;
  const scheme: Scheme = {
    id: crypto.randomUUID(),
    name: newScheme.value.name.trim(),
    type: newScheme.value.type?.trim() || "local",
    content: newScheme.value.content || "",
    enabled: true,
    createdAt: new Date().toISOString(),
    updatedAt: new Date().toISOString(),
  };
  store.saveScheme(scheme).then(() => {
    showAddModal.value = false;
    newScheme.value = { name: "", type: "local", content: "" };
    store.selectScheme(scheme.id);
  });
}

async function handleToggle(scheme: Scheme) {
  await store.toggleSchemeEnabled(scheme);
}

function startRename(scheme: Scheme) {
  editingId.value = scheme.id;
  editingName.value = scheme.name;
}

async function finishRename() {
  if (editingId.value && editingName.value.trim()) {
    const scheme = store.schemes.find((s) => s.id === editingId.value);
    if (scheme) {
      await store.saveScheme({ ...scheme, name: editingName.value.trim() });
    }
  }
  editingId.value = null;
  editingName.value = "";
}

function handleExport() {
  const data = {
    schemes: store.schemes,
    activeScheme: store.config?.activeScheme || "",
  };
  const blob = new Blob([JSON.stringify(data, null, 2)], {
    type: "application/json",
  });
  const url = URL.createObjectURL(blob);
  const a = document.createElement("a");
  a.href = url;
  a.download = `gohosts-config-${new Date().toISOString().slice(0, 10)}.json`;
  a.click();
  URL.revokeObjectURL(url);
}

function handleImport() {
  fileInputRef.value?.click();
}

function onFileSelected(e: Event) {
  const input = e.target as HTMLInputElement;
  const file = input.files?.[0];
  if (!file) return;
  const reader = new FileReader();
  reader.onload = (ev) => {
    try {
      const data = JSON.parse(ev.target?.result as string);
      if (confirm(`确定要导入 ${data.schemes?.length || 0} 个方案吗？`)) {
        const promises = (data.schemes || []).map((s: Scheme) =>
          store.saveScheme(s),
        );
        Promise.all(promises).then(() => {
          if (data.activeScheme && store.config) {
            store.config.activeScheme = data.activeScheme;
            store.saveConfig();
          }
        });
      }
    } catch {
      alert("导入失败：文件格式不正确");
    }
  };
  reader.readAsText(file);
  input.value = "";
}
</script>

<style scoped>
.scheme-manager {
  display: flex;
  flex-direction: column;
  height: 100%;
}

.sidebar-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 14px 16px;
  border-bottom: 1px solid #334155;
  flex-shrink: 0;
}

.sidebar-header h2 {
  margin: 0;
  font-size: 16px;
  color: #e2e8f0;
  font-weight: 600;
}

.scheme-list {
  flex: 1;
  overflow-y: auto;
  padding: 8px;
}

.scheme-item {
  padding: 10px 12px;
  border-radius: 6px;
  margin-bottom: 6px;
  cursor: pointer;
  background: #1e293b;
  border: 1px solid transparent;
  transition: all 0.2s;
}

.scheme-item:hover {
  background: #334155;
}

.scheme-item.selected {
  border-color: #4a9eff;
  background: #1e3a5f;
}

.scheme-item.active {
  border-left: 3px solid #22c55e;
}

.scheme-name-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 8px;
  margin-bottom: 6px;
}

.scheme-name {
  font-weight: 500;
  color: #e2e8f0;
  word-break: break-all;
  font-size: 14px;
}

.active-badge {
  font-size: 11px;
  padding: 1px 6px;
  border-radius: 4px;
  background: #22c55e30;
  color: #22c55e;
  white-space: nowrap;
  flex-shrink: 0;
}

.scheme-meta {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 8px;
}

.scheme-type {
  font-size: 12px;
  color: #94a3b8;
  text-transform: uppercase;
}

.scheme-actions {
  display: flex;
  gap: 8px;
  margin-top: 8px;
  padding-top: 8px;
  border-top: 1px solid #334155;
  opacity: 0;
  transition: opacity 0.2s;
}

.scheme-item:hover .scheme-actions {
  opacity: 1;
}

.empty-tip {
  text-align: center;
  color: #94a3b8;
  padding: 24px;
  font-size: 14px;
}

.sidebar-footer {
  display: flex;
  gap: 8px;
  padding: 12px 16px;
  border-top: 1px solid #334155;
  flex-shrink: 0;
}

.sidebar-footer .n-button {
  flex: 1;
}
</style>
