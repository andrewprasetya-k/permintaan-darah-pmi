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
  <div class="flex h-screen overflow-hidden bg-gray-50">
    <Sidebar />
    <div class="flex min-h-0 flex-1 flex-col transition-all duration-500">
      <Topbar :title="pageTitle" :subtitle="pageSubtitle" :on-menu-click="() => (isMobileOpen = true)" />
      <main class="flex-1 min-h-0 overflow-hidden p-6">
        <slot />
      </main>
    </div>
  </div>
</template>
