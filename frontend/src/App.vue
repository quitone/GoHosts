<script setup lang="ts">
import { useHostsStore } from '@/stores/hosts.store'
import { useAppStore } from './stores/app.store'

const store = useHostsStore()
const appStore = useAppStore()

onMounted(() => {
  // store.loadConfig()
  store.loadSystemHosts()
})
</script>

<template>
  <NavBar />
  <div class="main flex relative w-screen overflow-hidden bg-page">
    <transition leave-to-class="duration-300 ease-out translate-x--80" enter-to-class="duration-300 ease-in">
      <aside v-if="!appStore.isCollapsed" class="w-300px min-w-300px bg-neutral-800 b-r b-neutral-700 flex flex-col overflow-hidden">
        <LeftPanel />
      </aside>
    </transition>
    <main class="flex-1 flex flex-col overflow-hidden min-w-0 bg-page">
      <EditHostInfo />
    </main>

    <EditHostInfo :show="appStore.editSchemaVisible" />
  </div>
</template>

<style>
:root {
  --editor-gutter-bg: #f5f5f5;
  --editor-line-number-color: #000;
  --nav-bar-height: 50px;
}
.main {
  height: calc(100vh - 50px);
}
</style>
