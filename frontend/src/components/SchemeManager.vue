<template>
  <div class="flex flex-col h-full">
    <div class="flex-1 overflow-y-auto p-8px">
      <div
        v-for="scheme in store.schemes"
        :key="scheme.id"
        class="group p-10px-12px rounded-6px mb-6px cursor-pointer bg-neutral-800 b b-transparent transition-all duration-200 hover:bg-neutral-700"
        :class="{
          'b-primary bg-primary-light': scheme.id === store.selectedSchemeId,
          'border-l-3 border-l-success': scheme.id === store.activeSchemeId,
        }"
        @click="store.selectScheme(scheme.id)"
      >
        <div class="scheme-info">
          <div class="flex items-center justify-between gap-8px mb-6px">
            <span
              v-if="editingId !== scheme.id"
              class="font-500 text-neutral-100 word-break-all text-14px"
              >{{ scheme.name }}</span
            >
            <n-input
              v-else
              v-model:value="editingName"
              size="small"
              placeholder="方案名称"
              @blur="finishRename"
              @keyup.enter="finishRename"
            />
            <span
              v-if="scheme.id === store.activeSchemeId"
              class="text-11px px-6px py-1px rounded-4px bg-success/20 text-success whitespace-nowrap flex-shrink-0"
              >激活</span
            >
          </div>
          <div class="flex items-center justify-between gap-8px">
            <span class="text-12px text-neutral-400 uppercase">{{
              scheme.type
            }}</span>
            <n-switch
              :value="scheme.enabled"
              size="small"
              @update:value="handleToggle(scheme)"
            />
          </div>
        </div>
        <div
          class="flex gap-8px mt-8px pt-8px b-t b-neutral-700 opacity-0 transition-opacity duration-200 group-hover:opacity-100"
        >
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
    </div>

    <div class="flex gap-8px p-12px-16px b-t b-neutral-700 flex-shrink-0">
      <n-button size="small" @click="handleExport" class="flex-1"
        >导出配置</n-button
      >
      <n-button size="small" @click="handleImport" class="flex-1"
        >导入配置</n-button
      >
      <input
        ref="fileInputRef"
        type="file"
        accept=".json"
        class="hidden"
        @change="onFileSelected"
      />
    </div>

    <!-- 添加方案对话框 -->
    <n-modal
      v-model:show="showAddModal"
      title="新建方案"
      preset="card"
      class="w-420px"
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
        <div class="flex justify-end gap-8px">
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
/* 所有样式已转换为 UnoCSS 原子类 */
</style>
