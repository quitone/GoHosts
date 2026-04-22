import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import type { Config, Scheme } from '@/types/schema'
import {
  LoadConfig,
  SaveConfig,
  SaveScheme,
  DeleteScheme,
  GetSystemHosts,
} from '@wailsjs/go/app/App'

export const useHostsStore = defineStore('hosts', () => {
  // State
  const config = ref<Config | null>(null)
  const schemes = ref<Scheme[]>([])
  const selectedSchemeId = ref<string | null>(null)
  const systemHosts = ref<string>('')
  const loading = ref(false)
  const saving = ref(false)
  const error = ref('')

  // Getters
  const selectedScheme = computed(() =>
    schemes.value.find((s) => s.id === selectedSchemeId.value) || null
  )

  const activeSchemeId = computed(() => config.value?.activeScheme || '')

  // Actions
  async function loadConfig() {
    loading.value = true
    error.value = ''
    try {
      const data = await LoadConfig()
      config.value = data
      schemes.value = data.schemes || []
      // 默认选中激活方案或第一个方案
      if (data.activeScheme) {
        selectedSchemeId.value = data.activeScheme
      } else if (schemes.value.length > 0) {
        selectedSchemeId.value = schemes.value[0].id
      }
    } catch (e: any) {
      error.value = e.message || '加载配置失败'
      console.error('Failed to load config:', e)
    } finally {
      loading.value = false
    }
  }

  async function saveConfig() {
    if (!config.value) return
    saving.value = true
    error.value = ''
    try {
      await SaveConfig(config.value)
    } catch (e: any) {
      error.value = e.message || '保存配置失败'
      console.error('Failed to save config:', e)
      throw e
    } finally {
      saving.value = false
    }
  }

  async function saveScheme(scheme: Scheme) {
    saving.value = true
    error.value = ''
    try {
      await SaveScheme(scheme)
      const idx = schemes.value.findIndex((s) => s.id === scheme.id)
      if (idx >= 0) {
        schemes.value[idx] = { ...scheme }
      } else {
        schemes.value.push(scheme)
      }
      if (config.value) {
        config.value.schemes = [...schemes.value]
      }
    } catch (e: any) {
      error.value = e.message || '保存方案失败'
      console.error('Failed to save scheme:', e)
      throw e
    } finally {
      saving.value = false
    }
  }

  async function deleteScheme(id: string) {
    error.value = ''
    try {
      await DeleteScheme(id)
      schemes.value = schemes.value.filter((s) => s.id !== id)
      if (selectedSchemeId.value === id) {
        selectedSchemeId.value = schemes.value[0]?.id || null
      }
      if (config.value) {
        config.value.schemes = [...schemes.value]
        if (config.value.activeScheme === id) {
          config.value.activeScheme = ''
        }
      }
    } catch (e: any) {
      error.value = e.message || '删除方案失败'
      console.error('Failed to delete scheme:', e)
      throw e
    }
  }

  async function loadSystemHosts() {
    try {
      systemHosts.value = await GetSystemHosts()
    } catch (e: any) {
      console.error('Failed to load system hosts:', e)
    }
  }

  async function activateScheme(id: string) {
    if (!config.value) return
    config.value.activeScheme = id
    await saveConfig()
  }

  async function toggleSchemeEnabled(scheme: Scheme) {
    const updated = { ...scheme, enabled: !scheme.enabled }
    await saveScheme(updated)
  }

  function selectScheme(id: string) {
    selectedSchemeId.value = id
  }

  return {
    config,
    schemes,
    selectedSchemeId,
    selectedScheme,
    activeSchemeId,
    systemHosts,
    loading,
    saving,
    error,
    loadConfig,
    saveConfig,
    saveScheme,
    deleteScheme,
    loadSystemHosts,
    activateScheme,
    toggleSchemeEnabled,
    selectScheme,
  }
})
