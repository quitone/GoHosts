import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useAppStore = defineStore('app', () => {
  const isCollapsed = ref(false)
  const toggleCollapsed = () => {
    isCollapsed.value = !isCollapsed.value
  }

  const globalConfigVisible = ref(false)
  const toggleGlobalConfig = () => {
    globalConfigVisible.value = !globalConfigVisible.value
  }

  const historyVisible = ref(false)
  const toggleHistory = () => {
    historyVisible.value = !historyVisible.value
  }

  const editConfigVisible = ref(false)
  const openEditConfig = () => {
    editConfigVisible.value = true
  }

  return {
    isCollapsed,
    toggleCollapsed,

    configVisible: globalConfigVisible,
    toggleGlobalConfig,

    historyVisible,
    toggleHistory,

    editConfigVisible,
    openEditConfig,
  }
})
