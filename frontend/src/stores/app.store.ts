import { defineStore } from 'pinia'

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

  const editSchemaVisible = ref(false)
  const openEditSchema = () => {
    editSchemaVisible.value = true
  }

  return {
    isCollapsed,
    toggleCollapsed,

    configVisible: globalConfigVisible,
    toggleGlobalConfig,

    historyVisible,
    toggleHistory,

    editSchemaVisible,
    openEditSchema,
  }
})
