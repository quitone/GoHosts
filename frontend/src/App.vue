<script setup lang="ts">
import { onMounted } from "vue";
import { useHostsStore } from "@/stores/hosts";
import SchemeManager from "@/components/SchemeManager.vue";
import ConfigManager from "@/components/ConfigManager.vue";

const store = useHostsStore();

onMounted(() => {
  store.loadConfig();
  store.loadSystemHosts();
});
</script>

<template>
  <div class="app-layout">
    <aside class="sidebar">
      <SchemeManager />
    </aside>
    <main class="main-content">
      <ConfigManager />
    </main>
  </div>
</template>

<style>
.app-layout {
  display: flex;
  height: 100vh;
  width: 100vw;
  overflow: hidden;
  background: #0f172a;
}

.sidebar {
  width: 300px;
  min-width: 300px;
  background: #1e293b;
  border-right: 1px solid #334155;
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.main-content {
  flex: 1;
  display: flex;
  flex-direction: column;
  overflow: hidden;
  min-width: 0;
  background: #0f172a;
}
</style>
