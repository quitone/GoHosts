import { describe, it, expect, vi } from 'vitest'
import { mount } from '@vue/test-utils'
import SchemeManager from '@/components/SchemeManager.vue'
// import { GetSchemes, SaveScheme, DeleteScheme } from '@wailsjs/go/main/App'

vi.mock('../../wailsjs/go/main/App', () => ({
  GetSchemes: vi.fn(),
  SaveScheme: vi.fn(),
  DeleteScheme: vi.fn()
}))

describe('SchemeManager', async () => {
  const { GetSchemes, SaveScheme, DeleteScheme } = await import('../../wailsjs/go/main/App')

  it('displays list of schemes', async () => {
    const mockSchemes = [
      { id: '1', name: '开发', type: 'local', content: '127.0.0.1 dev', enabled: true }
    ]
      ; (GetSchemes as any).mockResolvedValue(mockSchemes)

    const wrapper = mount(SchemeManager)
    await new Promise(resolve => setTimeout(resolve, 10))

    expect(wrapper.text()).toContain('开发')
  })

  it('adds a new scheme', async () => {
    ; (GetSchemes as any).mockResolvedValue([])
      ; (SaveScheme as any).mockResolvedValue({})

    const wrapper = mount(SchemeManager)
    await wrapper.find('[data-test="add-scheme"]').trigger('click')
    await wrapper.find('[data-test="scheme-name"]').setValue('测试方案')
    await wrapper.find('[data-test="scheme-content"]').setValue('1.1.1.1 test')
    await wrapper.find('[data-test="save-scheme"]').trigger('click')

    expect(SaveScheme).toHaveBeenCalledWith(
      expect.objectContaining({ name: '测试方案', content: '1.1.1.1 test' })
    )
  })
})