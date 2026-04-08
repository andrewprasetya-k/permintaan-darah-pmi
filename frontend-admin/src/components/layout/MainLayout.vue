<script setup lang="ts">
import { ref } from 'vue'
import Sidebar from './Sidebar.vue'
import Topbar from './Topbar.vue'

defineProps<{ title?: string; subtitle?: string }>()

const isMobileOpen = ref(false)
</script>

<template>
  <div class="flex min-h-screen bg-gray-50">
    <Sidebar />
    <div class="flex flex-col flex-1 lg:ml-72 min-h-screen transition-all duration-500">
      <Topbar :title="title" :subtitle="subtitle" :on-menu-click="() => (isMobileOpen = true)" />
      <main class="flex-1 p-6 overflow-y-auto">
        <slot />
      </main>
    </div>
  </div>
</template>
