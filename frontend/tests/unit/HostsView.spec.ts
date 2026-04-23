import { describe, it, expect, vi } from 'vitest'
import { mount } from '@vue/test-utils'
import HostsView from '@/components/HostsView.vue'
// import { GetSystemHosts } from '@/wailsjs/go/main/App'

vi.mock('@/wailsjs/go/main/App', () => ({
  GetSystemHosts: vi.fn()
}))

describe('HostsView', async () => {
  const { GetSystemHosts } = await import('@/wailsjs/go/main/App')
  it('displays system hosts content on mount', async () => {
    const mockContent = '127.0.0.1 localhost'
      ; (GetSystemHosts as any).mockResolvedValue(mockContent)

    const wrapper = mount(HostsView)
    await new Promise(resolve => setTimeout(resolve, 10))

    expect(wrapper.text()).toContain(mockContent)
    expect(GetSystemHosts).toHaveBeenCalled()
  })

  it('shows refresh button and reloads content on click', async () => {
    const mockContent = '127.0.0.1 localhost'
    const mockContent2 = '192.168.1.1 app.local'
      ; (GetSystemHosts as any)
        .mockResolvedValueOnce(mockContent)
        .mockResolvedValueOnce(mockContent2)

    const wrapper = mount(HostsView)
    await new Promise(resolve => setTimeout(resolve, 10))

    const button = wrapper.find('[data-test="refresh"]')
    await button.trigger('click')
    await new Promise(resolve => setTimeout(resolve, 10))

    expect(wrapper.text()).toContain(mockContent2)
  })
})