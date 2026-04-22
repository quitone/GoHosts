<template>
  <div class="scheme-manager">
    <div class="toolbar">
      <h2>方案管理</h2>
      <button data-test="add-scheme" @click="showAddDialog = true">
        新建方案
      </button>
    </div>
    <ul>
      <li v-for="scheme in schemes" :key="scheme.id">
        <span>{{ scheme.name }}</span>
        <button @click="editScheme(scheme)">编辑</button>
        <button @click="deleteScheme(scheme.id)">删除</button>
      </li>
    </ul>
    <!-- 编辑对话框（略） -->
  </div>
</template>

<script setup lang="ts">
import { NDialog } from "naive-ui";
import { ref, onMounted } from "vue";
import {
  GetSchemes,
  SaveScheme,
  editScheme,
  DeleteScheme,
} from "@wailsjs/go/main/App";
import type { Scheme } from "@wailsjs/go/main/App";

const schemes = ref<Scheme[]>([]);
const showAddDialog = ref(false);

async function loadSchemes() {
  schemes.value = await GetSchemes();
}

async function deleteScheme(id: string) {
  await DeleteScheme(id);
  await loadSchemes();
}

// 其他方法略...
onMounted(loadSchemes);
</script>
