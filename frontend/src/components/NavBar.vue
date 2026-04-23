<template>
  <nav class="flex-between py-2 px-4 text-6 line-height-7">
    <div class="flex-center">
      <i
        :class="`i-mdi-hamburger-open cursor-pointer transition-base ${isCollapsed ? 'rotate-180' : 'rotate-0'}`"
        @click="appStore.toggleCollapsed" />
      <i class="i-mdi-plus cursor-pointer ml-4" @click="openNewSchema" />
    </div>

    <div>
      <!-- <i class="i-mdi-history cursor-pointer" /> -->
      <n-dropdown trigger="click" :options="menuOptions" @select="handleSelectMenu">
        <i class="i-mdi-settings cursor-pointer ml-4" @click="appStore.toggleGlobalConfig" />
      </n-dropdown>
      <i class="i-line-md-close cursor-pointer ml-4" @click="Quit" />
    </div>
  </nav>
</template>

<script lang="tsx" setup>
import { useAppStore } from '@/stores/app.store'
import { Quit } from '@wailsjs/go/main/App'

const appStore = useAppStore()

const { isCollapsed } = storeToRefs(appStore)

const renderIcon = (iconName: string) => {
  return () => h('i', { class: `${iconName} text-xl` })
}
const menuOptions = [
  {
    label: '关于',
    key: 'about',
    icon: renderIcon('i-mdi-information-outline'),
  },
  { key: 'divider-1', type: 'divider', show: true },
  {
    label: '选项',
    key: 'settings',
    icon: renderIcon('i-mdi-wrench-settings-outline'),
  },
  // {
  //   label: '退出',
  //   key: 'exit',
  //   icon: renderIcon('i-mdi-exit-to-app'),
  // },
]

function handleSelectMenu(key: string | number) {}
function openNewSchema() {
  appStore.openEditSchema()
}
</script>
