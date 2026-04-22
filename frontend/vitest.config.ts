import { defineConfig } from 'vitest/config'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      '@': '/src',
      '@wailsjs': '/wailsjs'
    }
  },
  test: {
    environment: 'happy-dom',        // <-- add this
    globals: true,               // optional: enables describe/it globally
  },
})
