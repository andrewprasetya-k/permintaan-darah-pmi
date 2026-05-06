<script setup lang="ts">
import { computed, onMounted, ref, type Component } from 'vue'
import { RouterLink, RouterView, useRoute, useRouter } from 'vue-router'
import {
  Activity,
  Droplets,
  Hospital,
  LayoutDashboard,
  LogOut,
  Menu,
  Plus,
  UserRound,
} from '@lucide/vue'
import AppModal from '@/components/feedback/AppModal.vue'
import { useAuthStore } from '@/stores/auth'
import { useFeedbackStore } from '@/stores/feedback'
import { useMyProfileStore } from '@/stores/my-profile'
import { useRealtimeStore } from '@/stores/realtime'

const authStore = useAuthStore()
const profileStore = useMyProfileStore()
const feedbackStore = useFeedbackStore()
const realtimeStore = useRealtimeStore()
const router = useRouter()
const route = useRoute()

const isMobileNavOpen = ref(false)
const isLogoutOpen = ref(false)

const navItems: Array<{ label: string; to: string; icon: Component }> = [
  { label: 'Dashboard', to: '/', icon: LayoutDashboard },
  { label: 'Permintaan', to: '/requests', icon: Droplets },
  { label: 'Buat Baru', to: '/requests/new', icon: Plus },
  { label: 'Profil', to: '/profile', icon: UserRound },
]

const hospitalName = computed(
  () => profileStore.profile?.nama || authStore.user?.username || 'Rumah Sakit',
)

const pageTitle = computed(() => {
  if (route.name === 'requests') return 'Permintaan Saya'
  if (route.name === 'request-create') return 'Buat Permintaan'
  if (route.name === 'request-edit') return 'Edit Permintaan'
  if (route.name === 'request-detail') return 'Detail Permintaan'
  if (route.name === 'profile') return 'Profil Rumah Sakit'
  return 'Dashboard'
})

const getInitials = (name: string) => {
  if (!name) return 'RS'
  const parts = name.trim().split(' ')
  if (parts.length === 1) return parts[0]?.substring(0, 2).toUpperCase()
  return ((parts[0]?.[0] ?? '') + (parts[parts.length - 1]?.[0] ?? '')).toUpperCase()
}

const realtimeLabel = computed(() => {
  if (realtimeStore.status === 'connected') {
    return 'Realtime aktif'
  }
  if (realtimeStore.status === 'connecting') {
    return 'Menghubungkan'
  }
  return 'Realtime offline'
})

const closeMobileNav = () => {
  isMobileNavOpen.value = false
}

const confirmLogout = async () => {
  realtimeStore.disconnect()
  authStore.logout()
  isLogoutOpen.value = false
  feedbackStore.showFlag({
    title: 'Logout berhasil',
    message: 'Sesi rumah sakit sudah ditutup.',
    variant: 'success',
  })
  await router.push('/login')
}

onMounted(async () => {
  if (authStore.token) {
    realtimeStore.connect(authStore.token)
  }

  if (!profileStore.profile) {
    await profileStore.fetchMe().catch(() => undefined)
  }
})
</script>

<template>
  <div class="flex min-h-screen bg-gray-50">
    <aside
      class="fixed inset-y-0 left-0 z-50 flex w-[280px] flex-col bg-white px-3 pb-4 pt-6 transition-transform duration-200 max-lg:w-[min(280px,88vw)] max-lg:-translate-x-full"
      :class="{ 'max-lg:translate-x-0': isMobileNavOpen }"
      aria-label="Navigasi utama"
    >
      <div class="flex min-h-12 items-center gap-3 px-3">
        <div class="flex h-9 w-9 items-center justify-center rounded-xl bg-blue-50 text-blue-600">
          <Hospital :size="20" />
        </div>
        <div>
          <strong class="block text-sm font-bold text-gray-900"
            >PMI <span class="text-red-500">Salatiga</span></strong
          >
          <small class="block text-xs text-gray-500">Portal Rumah Sakit</small>
        </div>
      </div>

      <nav class="mt-7 flex flex-1 flex-col gap-1 overflow-y-auto">
        <RouterLink
          v-for="item in navItems"
          :key="item.to"
          :to="item.to"
          class="flex min-h-[42px] items-center gap-2.5 rounded-xl px-3 py-2.5 text-sm font-medium text-gray-600 transition-all hover:bg-blue-50 hover:text-blue-700"
          :class="{
            'translate-x-0.5 bg-blue-600 text-white shadow-sm hover:bg-blue-600 hover:text-white':
              route.path === item.to,
          }"
          @click="closeMobileNav"
        >
          <component :is="item.icon" :size="18" class="shrink-0" />
          {{ item.label }}
        </RouterLink>
      </nav>

      <div
        class="mt-auto flex flex-col gap-3 border-t border-gray-100 px-3 pt-4 text-xs text-gray-600"
      >
        <div class="flex items-center gap-2">
          <Activity :size="16" />
          <span class="font-medium">Koneksi portal</span>
        </div>
        <div class="flex items-center gap-2">
          <span
            class="h-2.5 w-2.5 rounded-full"
            :class="{
              'bg-emerald-500': realtimeStore.status === 'connected',
              'bg-amber-500': realtimeStore.status === 'connecting',
              'bg-red-500':
                realtimeStore.status === 'error' || realtimeStore.status === 'disconnected',
              'bg-gray-400':
                realtimeStore.status !== 'connected' &&
                realtimeStore.status !== 'connecting' &&
                realtimeStore.status !== 'error' &&
                realtimeStore.status !== 'disconnected',
            }"
          />
          <span>{{ realtimeLabel }}</span>
        </div>
        <button
          type="button"
          class="flex min-h-10 w-full items-center gap-2 rounded-xl px-3 py-2 text-sm font-medium text-gray-600 transition hover:bg-red-50 hover:text-red-600"
          @click="isLogoutOpen = true"
        >
          <LogOut :size="18" />
          Logout
        </button>
      </div>
    </aside>

    <button
      v-if="isMobileNavOpen"
      type="button"
      class="fixed inset-0 z-40 border-0 bg-slate-900/40 lg:hidden"
      aria-label="Tutup navigasi"
      @click="closeMobileNav"
    />

    <div class="min-h-screen flex-1 pl-[280px] max-lg:pl-0">
      <header
        class="sticky top-0 z-30 flex min-h-[88px] items-center justify-between gap-4 bg-white px-8 py-5 max-lg:min-h-[76px] max-lg:px-4 max-lg:py-3"
      >
        <button
          type="button"
          class="hidden h-10 w-10 items-center justify-center rounded-xl text-gray-600 transition hover:bg-gray-100 hover:text-gray-900 max-lg:inline-flex"
          aria-label="Buka navigasi"
          @click="isMobileNavOpen = true"
        >
          <Menu :size="24" />
        </button>

        <div class="min-w-0 flex-1">
          <h1 class="truncate text-2xl font-semibold text-gray-900 max-sm:max-w-[180px]">
            {{ pageTitle }}
          </h1>
          <p class="mt-1 truncate text-sm text-gray-500 max-sm:max-w-[180px]">{{ hospitalName }}</p>
        </div>

        <div class="flex items-center gap-3">
          <div class="flex items-center gap-3">
            <div class="hidden text-right xl:block">
              <strong class="block max-w-48 truncate text-sm font-medium text-gray-900">{{
                hospitalName
              }}</strong>
              <span class="block text-xs text-gray-500">Rumah Sakit</span>
            </div>
            <div
              class="flex h-10 w-10 items-center justify-center rounded-xl bg-blue-100 font-semibold text-blue-600"
            >
              {{ getInitials(hospitalName) }}
            </div>
          </div>
        </div>
      </header>

      <main class="mx-auto w-full max-w-[1200px] px-8 py-6 max-lg:px-4 max-lg:py-5">
        <RouterView />
      </main>
    </div>

    <AppModal
      :is-open="isLogoutOpen"
      title="Keluar dari portal?"
      description="Sesi rumah sakit akan ditutup dari perangkat ini."
      width="sm"
      @close="isLogoutOpen = false"
    >
      <template #icon>
        <div class="flex h-12 w-12 items-center justify-center rounded-2xl bg-red-50 text-red-600">
          <LogOut :size="20" />
        </div>
      </template>
      <template #footer>
        <button
          type="button"
          class="inline-flex min-h-10 flex-1 items-center justify-center rounded-xl border border-gray-200 bg-white px-4 py-2.5 text-sm font-semibold text-gray-700 transition hover:bg-gray-50"
          @click="isLogoutOpen = false"
        >
          Batal
        </button>
        <button
          type="button"
          class="inline-flex min-h-10 flex-1 items-center justify-center rounded-xl border border-red-100 bg-red-50 px-4 py-2.5 text-sm font-semibold text-red-700 transition hover:border-red-200"
          @click="confirmLogout"
        >
          Logout
        </button>
      </template>
    </AppModal>
  </div>
</template>
