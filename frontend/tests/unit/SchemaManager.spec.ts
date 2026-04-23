import { describe, beforeEach, it, expect, vi } from 'vitest'
import { mount, flushPromises } from '@vue/test-utils'
import { createPinia, setActivePinia } from 'pinia'
import { nextTick } from 'vue'
import SchemeManager from '@/components/SchemeManager.vue'
import { useHostsStore } from '@/stores/hosts.store'

vi.mock('@/wailsjs/go/main/App', () => ({
  LoadConfig: vi.fn(),
  SaveConfig: vi.fn(),
  SaveScheme: vi.fn(),
  DeleteScheme: vi.fn(),
  GetSystemHosts: vi.fn(),
}))

describe('SchemeManager', async () => {
  const { LoadConfig, SaveScheme } = await import('@/wailsjs/go/main/App')

  beforeEach(() => {
    vi.clearAllMocks()
    setActivePinia(createPinia())
  })

  it('displays list of schemes', async () => {
    const mockConfig = {
      schemes: [
        {
          id: '1',
          name: '开发',
          type: 'local',
          content: '127.0.0.1 dev',
          enabled: true,
          createdAt: '',
          updatedAt: '',
        },
      ],
      activeScheme: '1',
    }
    ;(LoadConfig as any).mockResolvedValue(mockConfig)

    const wrapper = mount(SchemeManager)
    await new Promise(resolve => setTimeout(resolve, 10))

    expect(wrapper.text()).toContain('开发')
  })

  it('adds a new scheme', async () => {
    ;(LoadConfig as any).mockResolvedValue({ schemes: [], activeScheme: '' })
    ;(SaveScheme as any).mockResolvedValue(undefined)

    const wrapper = mount(SchemeManager)
    await new Promise(resolve => setTimeout(resolve, 10))

    // 由于 naive-ui 组件在 happy-dom 中渲染限制，直接操作组件内部状态测试逻辑
    ;(wrapper.vm as any).showAddModal = true
    ;(wrapper.vm as any).newScheme.name = '测试方案'
    ;(wrapper.vm as any).newScheme.content = '1.1.1.1 test'
    await nextTick()
    ;(wrapper.vm as any).handleAddScheme()
    await new Promise(resolve => setTimeout(resolve, 10))

    expect(SaveScheme).toHaveBeenCalledWith(expect.objectContaining({ name: '测试方案', content: '1.1.1.1 test' }))
  })
})
