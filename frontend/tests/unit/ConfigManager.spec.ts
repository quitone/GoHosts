import { describe, beforeEach, it, expect, vi } from 'vitest'
import { mount } from '@vue/test-utils'
import { createPinia, setActivePinia } from 'pinia'
import ConfigManager from '../../src/components/ConfigManager.vue'
import { useHostsStore } from '../../src/stores/hosts'

vi.mock('../../wailsjs/go/app/App', () => ({
  LoadConfig: vi.fn(),
  SaveConfig: vi.fn(),
  SaveScheme: vi.fn(),
  DeleteScheme: vi.fn(),
  GetSystemHosts: vi.fn(),
}))

describe('ConfigManager', async () => {
  const { SaveScheme } = await import('../../wailsjs/go/app/App')

  beforeEach(() => {
    vi.clearAllMocks()
    setActivePinia(createPinia())
  })

  it('shows empty state when no scheme is selected', () => {
    const wrapper = mount(ConfigManager)
    expect(wrapper.text()).toContain('请从左侧选择一个方案进行编辑')
  })

  it('shows editor when a scheme is selected via store', async () => {
    const store = useHostsStore()
    store.schemes = [
      {
        id: '1',
        name: '开发环境',
        type: 'local',
        content: '127.0.0.1 dev',
        enabled: true,
        createdAt: '',
        updatedAt: '',
      },
    ]
    store.selectScheme('1')

    const wrapper = mount(ConfigManager)
    await new Promise((resolve) => setTimeout(resolve, 10))

    expect(wrapper.find('.hosts-textarea').exists()).toBe(true)
    expect(wrapper.find('.name-input').exists()).toBe(true)
  })

  it('saves scheme on save button click', async () => {
    ;(SaveScheme as any).mockResolvedValue(undefined)

    const store = useHostsStore()
    store.schemes = [
      {
        id: '1',
        name: '开发环境',
        type: 'local',
        content: '127.0.0.1 dev',
        enabled: true,
        createdAt: '',
        updatedAt: '',
      },
    ]
    store.selectScheme('1')

    const wrapper = mount(ConfigManager)
    await new Promise((resolve) => setTimeout(resolve, 10))

    await wrapper.find('.hosts-textarea').setValue('127.0.0.1 new')
    await wrapper.find('button').trigger('click')
    await new Promise((resolve) => setTimeout(resolve, 10))

    expect(SaveScheme).toHaveBeenCalledWith(
      expect.objectContaining({ id: '1', content: '127.0.0.1 new' })
    )
  })
})
