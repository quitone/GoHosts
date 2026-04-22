<template>
  <div class="hosts-view">
    <div class="toolbar">
      <h2>系统 Hosts</h2>
      <button data-test="refresh" @click="loadHosts">刷新</button>
    </div>
    <pre v-if="content">{{ content }}</pre>
    <p v-else>加载中...</p>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import { GetSystemHosts } from "@wailsjs/go/main/App";

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
