<template>
  <div class="h-full flex flex-col px-16px pb-16px">
    <div class="flex-between items-center mb-12px flex-shrink-0">
      <h2 class="m-0 text-16px text-neutral-100">系统 Hosts</h2>
      <button
        data-test="refresh"
        @click="loadHosts"
        class="px-6px py-14px b b-primary bg-transparent text-primary rounded-4px cursor-pointer text-13px transition-all duration-200 hover:bg-primary hover:text-white">
        刷新
      </button>
    </div>
    <div class="flex-1 overflow-hidden min-h-0">
      <pre
        v-if="content"
        class="h-full overflow-auto bg-neutral-800 b b-neutral-700 rounded-6px p-12px m-0 text-neutral-100 font-mono text-14px leading-1.6 text-left box-border"
        >{{ content }}</pre
      >
      <p v-else class="text-center text-neutral-400 p-24px">加载中...</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { GetSystemHosts } from '@wailsjs/go/main/App'

const content = ref('')

async function loadHosts() {
  try {
    content.value = await GetSystemHosts()
  } catch (e) {
    console.error('Failed to load hosts:', e)
  }
}

onMounted(() => {
  loadHosts()
})
</script>

<style scoped>
/* 所有样式已转换为 UnoCSS 原子类 */
</style>
