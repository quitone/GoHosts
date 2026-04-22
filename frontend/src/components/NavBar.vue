<template>
  <nav class="flex-between py-2 px-4 text-xl">
    <div class="flex-center">
      <i
        :class="`i-line-md-menu-fold-left cursor-pointer transition-base ${isCollapsed ? 'rotate-180' : 'rotate-0'}`"
        @click="appStore.toggleCollapsed"
      />
      <i class="i-line-md-plus cursor-pointer ml-4" />
    </div>
    <div>
      <i class="i-mdi-history cursor-pointer" />
      <n-dropdown
        trigger="click"
        :options="menuOptions"
        @select="handleSelectMenu"
      >
        <i
          class="i-mdi-settings cursor-pointer ml-4"
          @click="appStore.toggleGlobalConfig"
        />
      </n-dropdown>

      <i class="i-line-md-close cursor-pointer ml-4" />
    </div>
  </nav>
</template>

<script lang="tsx" setup>
import { useAppStore } from "@/stores/app.store";
import { storeToRefs } from "pinia";
import { h, ref } from "vue";

const appStore = useAppStore();

const { isCollapsed } = storeToRefs(appStore);

const renderIcon = (iconName: string) => {
  return () => h("i", { class: `${iconName} text-xl` });
};
const menuOptions = [
  {
    label: "关于",
    key: "history",
    icon: renderIcon("i-mdi-information-outline"),
  },
  {
    label: "选项",
    key: "settings",
    icon: renderIcon("i-mdi-wrench-settings-outline"),
  },
  {
    label: "退出",
    key: "exit",
    icon: renderIcon("i-mdi-exit-to-app"),
  },
];

function handleSelectMenu(key: string | number) {}
</script>
