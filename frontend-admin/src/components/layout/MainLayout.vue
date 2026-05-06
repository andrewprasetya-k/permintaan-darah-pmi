<script setup lang="ts">
import { computed, ref, onMounted, onBeforeUnmount } from 'vue'
import { useRoute } from 'vue-router'
import { useWebsocketStore } from '@/stores/websocket'
import Sidebar from './Sidebar.vue'
import Topbar from './Topbar.vue'

defineProps<{ title?: string; subtitle?: string }>()

const route = useRoute()
const websocketStore = useWebsocketStore()
const isMobileOpen = ref(false)

const pageTitle = computed(() => (route.meta.title as string) || 'Dashboard')
const pageSubtitle = computed(() => (route.meta.subtitle as string) || '')

onMounted(() => {
  websocketStore.connect()
})

onBeforeUnmount(() => {
  websocketStore.disconnect()
})
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
