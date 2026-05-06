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
  <div class="app-shell">
    <aside class="sidebar" :class="{ 'sidebar-open': isMobileNavOpen }" aria-label="Navigasi utama">
      <div class="brand">
        <div class="brand-mark">
          <Hospital :size="20" />
        </div>
        <div>
          <strong>PMI <span>Salatiga</span></strong>
          <small>Portal Rumah Sakit</small>
        </div>
      </div>

      <nav class="side-nav">
        <RouterLink
          v-for="item in navItems"
          :key="item.to"
          :to="item.to"
          class="side-link"
          :class="{ active: route.path === item.to }"
          @click="closeMobileNav"
        >
          <component :is="item.icon" :size="18" class="side-link-icon" />
          {{ item.label }}
        </RouterLink>
      </nav>

      <div class="side-footer">
        <div>
          <button
            type="button"
            class="logout-button"
            aria-label="Logout"
            @click="isLogoutOpen = true"
          >
            <LogOut :size="18" />
          </button>
        </div>
        <div>
          <Activity :size="16" />
        </div>
        <div>
          <span class="connection-dot" :class="`connection-${realtimeStore.status}`" />
          <span>{{ realtimeLabel }}</span>
        </div>
      </div>
    </aside>

    <button
      v-if="isMobileNavOpen"
      type="button"
      class="nav-backdrop"
      aria-label="Tutup navigasi"
      @click="closeMobileNav"
    />

    <div class="shell-main">
      <header class="topbar">
        <button
          type="button"
          class="mobile-menu"
          aria-label="Buka navigasi"
          @click="isMobileNavOpen = true"
        >
          <Menu :size="24" />
        </button>

        <div class="topbar-title">
          <h1>{{ pageTitle }}</h1>
          <p>{{ hospitalName }}</p>
        </div>

        <div class="topbar-actions">
          <div class="user-chip">
            <div>
              <strong>{{ hospitalName }}</strong>
              <span>Rumah Sakit</span>
            </div>
            <div class="user-avatar">{{ getInitials(hospitalName) }}</div>
          </div>
        </div>
      </header>

      <main class="content">
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
        <div class="modal-icon-danger">
          <LogOut :size="20" />
        </div>
      </template>
      <template #footer>
        <button type="button" class="btn btn-secondary" @click="isLogoutOpen = false">Batal</button>
        <button type="button" class="btn btn-danger" @click="confirmLogout">Logout</button>
      </template>
    </AppModal>
  </div>
</template>

<style scoped>
.app-shell {
  display: flex;
  min-height: 100vh;
  background: #f9fafb;
}

.sidebar {
  position: fixed;
  inset: 0 auto 0 0;
  z-index: 50;
  display: flex;
  width: 280px;
  flex-direction: column;
  background: #ffffff;
  padding: 24px 12px 16px;
  transition: transform 0.24s ease;
}

.brand {
  display: flex;
  align-items: center;
  gap: 12px;
  min-height: 48px;
  padding: 0 12px;
}

.brand-mark {
  display: flex;
  width: 36px;
  height: 36px;
  align-items: center;
  justify-content: center;
  border-radius: 12px;
  background: #eff6ff;
  color: #2563eb;
}

.brand strong,
.brand small {
  display: block;
}

.brand strong {
  color: #111827;
  font-size: 14px;
  font-weight: 700;
}

.brand strong span {
  color: #ef4444;
}

.brand small {
  color: var(--text-muted);
  font-size: 12px;
  font-weight: 400;
}

.side-nav {
  display: flex;
  flex: 1;
  flex-direction: column;
  gap: 4px;
  margin-top: 28px;
  overflow-y: auto;
}

.side-link {
  display: flex;
  min-height: 42px;
  align-items: center;
  gap: 10px;
  border-radius: 12px;
  padding: 10px 12px;
  color: #4b5563;
  font-size: 14px;
  font-weight: 500;
  transition:
    background-color 0.18s ease,
    color 0.18s ease,
    transform 0.18s ease;
}

.side-link:hover,
.side-link.active {
  background: #eff6ff;
  color: #1d4ed8;
}

.side-link.active {
  background: #2563eb;
  color: #ffffff;
  box-shadow: 0 1px 2px rgba(37, 99, 235, 0.18);
  transform: translateX(2px);
}

.side-link-icon {
  flex: 0 0 auto;
}

.side-footer {
  display: flex;
  align-items: flex-start;
  gap: 10px;
  margin-top: auto;
  border-top: 1px solid #f3f4f6;
  padding: 16px 12px 0;
  color: #4b5563;
  font-size: 12px;
}

.side-footer > div {
  display: flex;
  align-items: center;
  gap: 8px;
}

.connection-dot {
  width: 9px;
  height: 9px;
  border-radius: 999px;
  background: #9ca3af;
}

.connection-connected {
  background: #10b981;
}

.connection-connecting {
  background: #f59e0b;
}

.connection-error,
.connection-disconnected {
  background: #ef4444;
}

.shell-main {
  min-height: 100vh;
  flex: 1;
  padding-left: 280px;
}

.topbar {
  position: sticky;
  top: 0;
  z-index: 40;
  display: flex;
  min-height: 88px;
  align-items: center;
  justify-content: space-between;
  gap: 16px;
  background: #ffffff;
  padding: 20px 32px;
  box-shadow: 0 1px 2px rgba(15, 23, 42, 0.05);
}

.topbar-title {
  min-width: 0;
}

.topbar-title h1,
.topbar-title p {
  margin: 0;
}

.topbar-title h1 {
  overflow: hidden;
  max-width: 520px;
  color: #111827;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 24px;
  font-weight: 600;
}

.topbar-title p {
  overflow: hidden;
  max-width: 520px;
  text-overflow: ellipsis;
  white-space: nowrap;
  color: #6b7280;
  font-size: 14px;
  margin-top: 4px;
}

.topbar-actions {
  display: flex;
  align-items: center;
  gap: 12px;
}

.user-chip {
  display: flex;
  align-items: center;
  gap: 12px;
}

.user-chip div:first-child {
  display: none;
  text-align: right;
}

.user-chip strong,
.user-chip span {
  display: block;
}

.user-chip strong {
  max-width: 190px;
  overflow: hidden;
  color: #111827;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 14px;
  font-weight: 500;
}

.user-chip span {
  color: #6b7280;
  font-size: 12px;
}

.user-avatar {
  display: flex;
  width: 40px;
  height: 40px;
  align-items: center;
  justify-content: center;
  border-radius: 12px;
  background: #dbeafe;
  color: #2563eb;
  font-weight: 600;
}

.logout-button,
.mobile-menu {
  display: inline-flex;
  width: 40px;
  height: 40px;
  align-items: center;
  justify-content: center;
  border: 0;
  border-radius: 12px;
  background: transparent;
  color: #4b5563;
  transition:
    background-color 0.18s ease,
    color 0.18s ease;
}

.logout-button:hover,
.mobile-menu:hover {
  background: #f3f4f6;
  color: #111827;
}

.mobile-menu,
.nav-backdrop {
  display: none;
}

.content {
  width: min(1200px, 100%);
  margin: 0 auto;
  padding: 24px 32px;
}

.modal-icon-danger {
  display: flex;
  height: 48px;
  width: 48px;
  align-items: center;
  justify-content: center;
  border-radius: 16px;
  background: #fef2f2;
  color: #dc2626;
}

@media (max-width: 920px) {
  .sidebar {
    transform: translateX(-100%);
    width: min(280px, 88vw);
  }

  .sidebar-open {
    transform: translateX(0);
  }

  .nav-backdrop {
    position: fixed;
    inset: 0;
    z-index: 45;
    display: block;
    border: 0;
    background: rgba(15, 23, 42, 0.36);
  }

  .shell-main {
    padding-left: 0;
  }

  .mobile-menu {
    display: inline-flex;
  }

  .topbar {
    min-height: 76px;
    padding: 14px 16px;
  }

  .new-request-link,
  .user-chip div:first-child {
    display: none;
  }

  .content {
    padding: 20px 16px;
  }
}

@media (max-width: 560px) {
  .topbar-title h1,
  .topbar-title p {
    max-width: 180px;
  }
}

@media (min-width: 1180px) {
  .user-chip div:first-child {
    display: block;
  }
}
</style>
