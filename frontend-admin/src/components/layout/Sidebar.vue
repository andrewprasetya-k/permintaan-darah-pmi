<script setup lang="ts">
import { useRoute, useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { computed, ref, watch } from 'vue'
import {
  LayoutDashboard,
  Users,
  Droplets,
  MapPin,
  LogOut,
  Calendar,
  ClipboardList,
  ChevronLeft,
  ChevronRight,
  Building2,
  Shield,
  Menu,
  X,
} from '@lucide/vue'

const route = useRoute()
const router = useRouter()
const authStore = useAuthStore()

const isCollapsed = ref(JSON.parse(localStorage.getItem('sidebarCollapsed') || 'false'))
const isMobileOpen = ref(false)

watch(isCollapsed, (val) => localStorage.setItem('sidebarCollapsed', JSON.stringify(val)))
watch(
  () => route.path,
  () => {
    isMobileOpen.value = false
  },
)

const isSuperAdmin = computed(() => authStore.user?.role === 'superadmin')
const isActive = (path: string) => route.path === path

const menuGroups = computed(() => [
  {
    items: [{ name: 'Dashboard', to: '/dashboard', icon: LayoutDashboard }],
  },
  {
    title: 'PERMINTAAN PRODUK DARAH',
    items: [{ name: 'Permintaan Darah', to: '/permintaan-darah', icon: ClipboardList }],
  },
  {
    title: 'RUMAH SAKIT',
    items: [{ name: 'Daftar Rumah Sakit', to: '/rumah-sakit', icon: Building2 }],
  },
  ...(isSuperAdmin.value
    ? [
        {
          title: 'ADMIN',
          items: [{ name: 'Manajemen Admin', to: '/admin-management', icon: Shield }],
        },
      ]
    : []),
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
      isCollapsed ? 'lg:w-20' : 'lg:w-72',
      isMobileOpen ? 'translate-x-0 w-72' : '-translate-x-full lg:translate-x-0',
    ]"
  >
    <!-- Header -->
    <div class="transition-all duration-500" :class="isCollapsed ? 'px-3 py-6' : 'px-6 py-6'">
      <div class="flex items-center justify-between">
        <div v-if="!isCollapsed" class="flex items-center space-x-2">
          <div class="w-10 h-10 bg-red-100 rounded-lg flex items-center justify-center">
            <span class="text-red-600 font-bold text-sm">PMI</span>
          </div>
          <div>
            <h2 class="text-lg font-bold leading-tight">
              <span class="text-black">PMI </span>
              <span class="text-red-500">Salatiga</span>
            </h2>
            <p class="text-xs font-semibold text-gray-500">Portal Admin</p>
          </div>
        </div>

        <button
          @click="isCollapsed = !isCollapsed"
          class="hidden lg:block p-2.5 rounded-lg hover:bg-blue-100 text-blue-700 transition-colors"
          :class="{ 'mx-auto': isCollapsed }"
        >
          <ChevronLeft v-if="!isCollapsed" :size="20" />
          <ChevronRight v-else :size="20" />
        </button>
      </div>
    </div>

    <!-- Nav -->
    <nav
      class="flex-1 py-6 space-y-4 overflow-y-auto transition-all duration-500"
      :class="isCollapsed ? 'px-2' : 'px-4'"
    >
      <div v-for="(group, gi) in menuGroups" :key="gi" class="space-y-1">
        <h3
          v-if="group.title && !isCollapsed"
          class="px-4 py-3 text-xs font-semibold text-gray-400 uppercase tracking-wider"
        >
          {{ group.title }}
        </h3>

        <router-link
          v-for="item in group.items"
          :key="item.to"
          :to="item.to"
          class="flex items-center px-4 py-3.5 rounded-xl text-sm font-medium transition-all duration-300"
          :class="
            isActive(item.to)
              ? 'bg-blue-600 text-white shadow-md translate-x-1'
              : 'text-gray-700 hover:bg-blue-50 hover:text-blue-700'
          "
        >
          <component :is="item.icon" :size="20" class="shrink-0" />
          <span v-if="!isCollapsed" class="ml-4">{{ item.name }}</span>
        </router-link>
      </div>
    </nav>

    <!-- Logout -->
    <div class="py-6 transition-all duration-500" :class="isCollapsed ? 'px-2' : 'px-4'">
      <button
        @click="logout"
        class="flex items-center w-full px-4 py-3.5 text-sm font-medium text-gray-700 rounded-xl hover:bg-red-50 hover:text-red-700 transition-colors"
      >
        <LogOut :size="20" class="shrink-0" />
        <span v-if="!isCollapsed" class="ml-4">Logout</span>
      </button>
    </div>
  </div>
</template>
