<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { RouterLink, RouterView, useRoute, useRouter } from 'vue-router'
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

const navItems = [
  { label: 'Dashboard', to: '/', marker: 'D' },
  { label: 'Permintaan', to: '/requests', marker: 'P' },
  { label: 'Buat Baru', to: '/requests/new', marker: '+' },
  { label: 'Profil', to: '/profile', marker: 'R' },
]

const hospitalName = computed(
  () => profileStore.profile?.nama || authStore.user?.username || 'Rumah Sakit',
)

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
        <div class="brand-mark">PMI</div>
        <div>
          <strong>Permintaan Darah</strong>
          <span>Portal Rumah Sakit</span>
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
          <span>{{ item.marker }}</span>
          {{ item.label }}
        </RouterLink>
      </nav>

      <div class="side-footer">
        <span class="connection-dot" :class="`connection-${realtimeStore.status}`" />
        <span>{{ realtimeLabel }}</span>
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
          class="btn btn-secondary btn-icon mobile-menu"
          aria-label="Buka navigasi"
          @click="isMobileNavOpen = true"
        >
          =
        </button>

        <div class="topbar-title">
          <span>Rumah Sakit</span>
          <strong>{{ hospitalName }}</strong>
        </div>

        <div class="topbar-actions">
          <RouterLink class="btn btn-secondary" to="/requests/new">+ Permintaan</RouterLink>
          <button type="button" class="btn btn-ghost" @click="isLogoutOpen = true">Logout</button>
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
      <template #footer>
        <button type="button" class="btn btn-secondary" @click="isLogoutOpen = false">Batal</button>
        <button type="button" class="btn btn-danger" @click="confirmLogout">Logout</button>
      </template>
    </AppModal>
  </div>
</template>

<style scoped>
.app-shell {
  min-height: 100vh;
}

.sidebar {
  position: fixed;
  inset: 0 auto 0 0;
  z-index: 50;
  display: flex;
  width: 270px;
  flex-direction: column;
  border-right: 1px solid var(--line);
  background: #ffffff;
  padding: 18px;
}

.brand {
  display: flex;
  align-items: center;
  gap: 12px;
  min-height: 54px;
}

.brand-mark {
  display: grid;
  width: 44px;
  height: 44px;
  place-items: center;
  border-radius: 8px;
  background: var(--red);
  color: #ffffff;
  font-size: 13px;
  font-weight: 900;
}

.brand strong,
.brand span {
  display: block;
}

.brand strong {
  font-size: 15px;
  font-weight: 900;
}

.brand span {
  color: var(--text-muted);
  font-size: 12px;
  font-weight: 700;
}

.side-nav {
  display: grid;
  gap: 6px;
  margin-top: 28px;
}

.side-link {
  display: flex;
  min-height: 44px;
  align-items: center;
  gap: 10px;
  border-radius: 8px;
  padding: 10px 12px;
  color: var(--text-soft);
  font-weight: 800;
}

.side-link span {
  display: grid;
  width: 26px;
  height: 26px;
  place-items: center;
  border: 1px solid var(--line);
  border-radius: 8px;
  background: var(--surface-muted);
  color: var(--text-muted);
  font-size: 12px;
  font-weight: 900;
}

.side-link:hover,
.side-link.active {
  background: var(--red-soft);
  color: #b91c1c;
}

.side-link.active span {
  border-color: #fecaca;
  background: #ffffff;
  color: #b91c1c;
}

.side-footer {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-top: auto;
  border: 1px solid var(--line);
  border-radius: 8px;
  padding: 11px 12px;
  color: var(--text-soft);
  font-size: 13px;
  font-weight: 700;
}

.connection-dot {
  width: 9px;
  height: 9px;
  border-radius: 999px;
  background: #9ca3af;
}

.connection-connected {
  background: var(--green);
}

.connection-connecting {
  background: var(--amber);
}

.connection-error,
.connection-disconnected {
  background: var(--red);
}

.shell-main {
  min-height: 100vh;
  padding-left: 270px;
}

.topbar {
  position: sticky;
  top: 0;
  z-index: 40;
  display: flex;
  min-height: 72px;
  align-items: center;
  justify-content: space-between;
  gap: 16px;
  border-bottom: 1px solid var(--line);
  background: rgba(246, 247, 251, 0.92);
  padding: 14px 28px;
  backdrop-filter: blur(12px);
}

.topbar-title {
  min-width: 0;
}

.topbar-title span,
.topbar-title strong {
  display: block;
}

.topbar-title span {
  color: var(--text-muted);
  font-size: 12px;
  font-weight: 800;
  text-transform: uppercase;
}

.topbar-title strong {
  overflow: hidden;
  max-width: 520px;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 17px;
  font-weight: 900;
}

.topbar-actions {
  display: flex;
  align-items: center;
  gap: 10px;
}

.mobile-menu,
.nav-backdrop {
  display: none;
}

.content {
  width: min(1180px, 100%);
  margin: 0 auto;
  padding: 28px;
}

@media (max-width: 920px) {
  .sidebar {
    transform: translateX(-100%);
    transition: transform 0.2s ease;
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
    padding: 12px 16px;
  }

  .topbar-actions .btn-secondary {
    display: none;
  }

  .content {
    padding: 20px 16px;
  }
}

@media (max-width: 560px) {
  .topbar-title strong {
    max-width: 180px;
  }
}
</style>
