import { vi } from 'vitest'

// Mock Wails 运行时（避免在测试环境中报错）
vi.mock('../wailsjs/runtime', () => ({
  default: {
    Log: {
      Print: vi.fn(),
      Debug: vi.fn(),
      Info: vi.fn(),
      Error: vi.fn()
    }
  }
}))