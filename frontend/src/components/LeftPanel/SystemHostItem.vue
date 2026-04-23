<template>
  <div class="flex items-center text-5` pl-4 hover:bg-neutral-700 cursor-pointer" @click="showSystemHosts">
    <i class="i-mdi-computer mr-2"></i>
    <span class="text-sm">系统Hosts</span>
  </div>
</template>

<script setup lang="ts">
import { useHostsStore } from '@/stores/hosts.store'

const hostsStore = useHostsStore()
const { showSystemHosts } = hostsStore
</script>
