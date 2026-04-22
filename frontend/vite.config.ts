import { defineConfig } from 'vite'
import Vue from '@vitejs/plugin-vue'
import VueJsx from '@vitejs/plugin-vue-jsx'
import UnoCSS from 'unocss/vite'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [Vue(), VueJsx(), UnoCSS()],
  resolve: {
    alias: {
      '@': '/src',
      '@wailsjs': '/wailsjs'
    },
    // extensions: ['.js', '.ts', '.json']
  },

})
