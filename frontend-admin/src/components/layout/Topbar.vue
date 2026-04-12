<script setup lang="ts">
import { computed } from 'vue'
import { useAuthStore } from '@/stores/auth'
import { Shield, Menu } from '@lucide/vue'

defineProps<{
  title?: string
  subtitle?: string
  onMenuClick?: () => void
}>()

const authStore = useAuthStore()

const userName = computed(() => authStore.user?.username || 'Admin')
const userRole = computed(() =>
  authStore.user?.role === 'superadmin' ? 'Superadmin' : 'Administrator',
)
const isSuperAdmin = computed(() => authStore.user?.role === 'superadmin')

const getInitials = (name: string) => {
  if (!name) return 'AD'
  const parts = name.trim().split(' ')
  if (parts.length === 1) return parts?.[0]?.substring(0, 2).toUpperCase()
  return ((parts?.[0]?.[0] ?? '') + (parts?.[parts.length - 1]?.[0] ?? '')).toUpperCase()
}
</script>

<template>
  <header
    class="bg-white px-4 lg:px-6 py-4 lg:py-6 flex items-center justify-between sticky top-0 z-30 shadow-sm lg:shadow-none"
  >
    <!-- Mobile menu button -->
    <button @click="onMenuClick" class="lg:hidden p-2 rounded-lg hover:bg-gray-100 text-gray-700">
      <Menu :size="24" />
    </button>

    <!-- Title -->
    <div class="flex items-center flex-1 lg:flex-none">
      <div class="lg:pl-0">
        <h1 class="text-xl lg:text-2xl font-semibold text-gray-900 truncate">
          {{ title || 'Dashboard' }}
        </h1>
        <p v-if="subtitle" class="text-xs lg:text-sm text-gray-500 mt-1 hidden sm:block">
          {{ subtitle }}
        </p>
      </div>
    </div>

    <!-- User -->
    <div class="flex items-center space-x-2 lg:space-x-6">
      <div class="flex items-center space-x-2 lg:space-x-4">
        <div class="text-right hidden md:block">
          <div class="text-sm font-medium text-gray-900">{{ userName }}</div>
          <div class="text-xs text-gray-500 flex items-center gap-1 justify-end">
            <Shield v-if="isSuperAdmin" :size="12" class="text-purple-600" />
            {{ userRole }}
          </div>
        </div>

        <div
          class="w-8 h-8 lg:w-10 lg:h-10 rounded-xl flex items-center justify-center text-sm lg:text-base font-semibold"
          :class="isSuperAdmin ? 'bg-purple-100 text-purple-600' : 'bg-blue-100 text-blue-600'"
        >
          {{ getInitials(userName) }}
        </div>
      </div>
    </div>
  </header>
</template>
