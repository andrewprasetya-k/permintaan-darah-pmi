<script setup lang="ts">
import { computed, ref } from 'vue'
import { useRoute } from 'vue-router'
import Sidebar from './Sidebar.vue'
import Topbar from './Topbar.vue'

defineProps<{ title?: string; subtitle?: string }>()

const route = useRoute()
const isMobileOpen = ref(false)

const pageTitle = computed(() => (route.meta.title as string) || 'Dashboard')
const pageSubtitle = computed(() => (route.meta.subtitle as string) || '')
</script>

<template>
  <div class="flex min-h-screen bg-gray-50">
    <Sidebar />
    <div class="flex flex-col flex-1 min-h-screen transition-all duration-500">
      <Topbar :title="pageTitle" :subtitle="pageSubtitle" :on-menu-click="() => (isMobileOpen = true)" />
      <main class="flex-1 p-6 overflow-y-auto">
        <slot />
      </main>
    </div>
  </div>
</template>
