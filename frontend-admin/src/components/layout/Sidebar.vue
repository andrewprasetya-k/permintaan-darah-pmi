<script setup lang="ts">
import { useRoute, useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { computed, ref, watch } from 'vue'
import {
  LayoutDashboard,
  Droplets,
  Users,
  Hospital,
  FlaskConical,
  ScrollText,
  ChevronLeft,
  ChevronRight,
  LogOut,
  Menu,
  X,
} from '@lucide/vue'

const route = useRoute()
const router = useRouter()
const authStore = useAuthStore()
const userRole = computed(() => authStore.user?.role)

const isCollapsed = ref(JSON.parse(localStorage.getItem('sidebarCollapsed') || 'false'))
const isMobileOpen = ref(false)

watch(isCollapsed, (val) => localStorage.setItem('sidebarCollapsed', JSON.stringify(val)))
watch(
  () => route.path,
  () => {
    isMobileOpen.value = false
  },
)

const isAdmin = computed(
  () => authStore.user?.role === 'superadmin' || authStore.user?.role === 'admin',
)
const isActive = (path: string) => route.path === path

const menuGroups = computed(() => [
  {
    items: [
      { name: 'Dashboard', to: '/', icon: LayoutDashboard },
      { name: 'Permintaan Darah', to: '/permintaan', icon: Droplets },
      { name: 'Manajemen Rumah Sakit', to: '/rumah-sakit', icon: Hospital },
      { name: 'Komponen Darah', to: '/komponen', icon: FlaskConical },
      { name: 'System Logs', to: '/logs', icon: ScrollText },
      ...(userRole.value === 'superadmin'
        ? [{ name: 'Manajemen Admin', to: '/admin', icon: Users }]
        : []),
    ],
  },
])

const logout = () => {
  authStore.logout()
  router.push('/login')
}
</script>

<template>
  <!-- Mobile Toggle -->
  <button
    @click="isMobileOpen = !isMobileOpen"
    class="lg:hidden fixed top-4 left-4 z-50 p-2.5 bg-white rounded-xl shadow-lg text-blue-700 hover:bg-blue-50 transition-colors"
  >
    <X v-if="isMobileOpen" :size="24" />
    <Menu v-else :size="24" />
  </button>

  <!-- Overlay -->
  <div
    v-if="isMobileOpen"
    class="lg:hidden fixed inset-0 bg-black/50 z-40"
    @click="isMobileOpen = false"
  />

  <!-- Sidebar -->
  <div
    class="bg-white h-screen flex flex-col transition-all duration-500 ease-in-out fixed lg:relative z-40"
    :class="[
      isCollapsed ? 'lg:w-20' : 'lg:w-1/5',
      isMobileOpen ? 'translate-x-0 w-full' : '-translate-x-full lg:translate-x-0',
    ]"
  >
    <!-- Header -->
    <div class="transition-all duration-500" :class="isCollapsed ? 'px-3 py-6' : 'px-6 py-6'">
      <div class="flex items-center justify-between">
        <div v-if="!isCollapsed" class="flex items-center gap-2">
          <div class="w-9 h-9 bg-red-100 rounded-lg flex items-center justify-center shrink-0">
            <span class="text-red-600 font-bold text-xs">PMI</span>
          </div>
          <div>
            <h2 class="text-base font-bold leading-tight text-gray-900">
              PMI <span class="text-red-500">Salatiga</span>
            </h2>
            <p class="text-xs text-gray-400">Permintaan Darah</p>
          </div>
        </div>

        <button
          @click="isCollapsed = !isCollapsed"
          class="hidden lg:block p-2 rounded-lg hover:bg-blue-50 text-blue-700 transition-colors"
          :class="{ 'mx-auto': isCollapsed }"
        >
          <ChevronLeft v-if="!isCollapsed" :size="18" />
          <ChevronRight v-else :size="18" />
        </button>
      </div>
    </div>

    <!-- Nav -->
    <nav
      class="flex-1 py-4 space-y-4 overflow-y-auto transition-all duration-500"
      :class="isCollapsed ? 'px-2' : 'px-3'"
    >
      <div v-for="(group, gi) in menuGroups" :key="gi" class="space-y-1">
        <router-link
          v-for="item in group.items"
          :key="item.to"
          :to="item.to"
          class="flex items-center px-3 py-2.5 rounded-xl text-sm font-medium transition-all duration-200"
          :class="
            isActive(item.to)
              ? 'bg-blue-600 text-white shadow-sm translate-x-0.5'
              : 'text-gray-600 hover:bg-blue-50 hover:text-blue-700'
          "
          :title="isCollapsed ? item.name : undefined"
        >
          <component :is="item.icon" :size="18" class="shrink-0" />
          <span v-if="!isCollapsed" class="ml-3 truncate">{{ item.name }}</span>
        </router-link>
      </div>
    </nav>

    <!-- Logout -->
    <div
      class="py-4 transition-all duration-500 border-t border-gray-100"
      :class="isCollapsed ? 'px-2' : 'px-3'"
    >
      <button
        @click="logout"
        class="flex items-center w-full px-3 py-2.5 text-sm font-medium text-gray-600 rounded-xl hover:bg-red-50 hover:text-red-600 transition-colors"
        :title="isCollapsed ? 'Logout' : undefined"
      >
        <LogOut :size="18" class="shrink-0" />
        <span v-if="!isCollapsed" class="ml-3">Logout</span>
      </button>
    </div>
  </div>
</template>
