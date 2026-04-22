<template>
  <div class="hosts-view">
    <div class="toolbar">
      <h2>系统 Hosts</h2>
      <button data-test="refresh" @click="loadHosts">刷新</button>
    </div>
    <div class="hosts-body">
      <pre v-if="content">{{ content }}</pre>
      <p v-else class="loading-text">加载中...</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import { GetSystemHosts } from "@wailsjs/go/app/App";

const content = ref("");

async function loadHosts() {
  try {
    content.value = await GetSystemHosts();
  } catch (e) {
    console.error("Failed to load hosts:", e);
  }
}

onMounted(() => {
  loadHosts();
});
</script>

<style scoped>
.hosts-view {
  height: 100%;
  display: flex;
  flex-direction: column;
  padding: 0 16px 16px;
}

.toolbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
  flex-shrink: 0;
}

.toolbar h2 {
  margin: 0;
  font-size: 16px;
  color: #e2e8f0;
}

button {
  padding: 6px 14px;
  border: 1px solid #4a9eff;
  background: transparent;
  color: #4a9eff;
  border-radius: 4px;
  cursor: pointer;
  font-size: 13px;
  transition: all 0.2s;
}

button:hover {
  background: #4a9eff;
  color: white;
}

.hosts-body {
  flex: 1;
  overflow: hidden;
  min-height: 0;
}

pre {
  height: 100%;
  overflow: auto;
  background: #1e293b;
  border: 1px solid #334155;
  border-radius: 6px;
  padding: 12px;
  margin: 0;
  color: #e2e8f0;
  font-family: "Fira Code", "Consolas", "Monaco", "Courier New", monospace;
  font-size: 14px;
  line-height: 1.6;
  text-align: left;
  box-sizing: border-box;
}

.loading-text {
  text-align: center;
  color: #94a3b8;
  padding: 24px;
}
</style>
