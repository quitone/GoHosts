import { createApp } from 'vue'
import { createPinia } from 'pinia'
import 'virtual:uno.css'
import naive from 'naive-ui'
import App from './App.vue'
import './style.css'

createApp(App).use(createPinia()).use(naive).mount('#app')
