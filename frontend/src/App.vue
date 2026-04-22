<script setup lang="ts">
import { onMounted } from "vue";
import { useHostsStore } from "@/stores/hosts";
import SchemeManager from "@/components/SchemeManager.vue";
import ConfigManager from "@/components/ConfigManager.vue";
import NavBar from "@/components/NavBar.vue";
import { useAppStore } from "./stores/app.store";

const store = useHostsStore();
const appStore = useAppStore();

onMounted(() => {
  store.loadConfig();
  store.loadSystemHosts();
});
</script>

<template>
  <NavBar />
  <div class="flex h-screen w-screen overflow-hidden bg-page">
    <transition
      leave-to-class="duration-300 ease-out translate-x--80"
      enter-to-class="duration-300 ease-in"
    >
      <aside
        v-if="!appStore.isCollapsed"
        class="w-300px min-w-300px bg-neutral-800 b-r b-neutral-700 flex flex-col overflow-hidden"
      >
        <SchemeManager />
      </aside>
    </transition>
    <main class="flex-1 flex flex-col overflow-hidden min-w-0 bg-page">
      <ConfigManager />
    </main>
  </div>
</template>
