<script setup lang="ts">
import { computed, onMounted } from 'vue'
import { useAuthStore } from '@/stores/auth'
import { useDashboardStore } from '@/stores/dashboard'
import { useLogsStore } from '@/stores/logs'

const authStore = useAuthStore()
const dashboardStore = useDashboardStore()
const logsStore = useLogsStore()

const cards = computed(() => {
  const summary = dashboardStore.statusSummary
  return [
    { label: 'Permintaan Baru', value: summary?.dibuat ?? 0, color: '#2563eb' },
    { label: 'Dalam Proses', value: summary?.diproses ?? 0, color: '#f59e0b' },
    { label: 'Siap Diambil', value: summary?.bisaDiambil ?? 0, color: '#10b981' },
    { label: 'Selesai', value: summary?.selesai ?? 0, color: '#8b5cf6' },
    { label: 'Dibatalkan', value: summary?.dibatalkan ?? 0, color: '#ef4444' },
    { label: 'Total', value: summary?.total ?? 0, color: '#111827' },
  ]
})

onMounted(async () => {
  if (authStore.user?.role === 'admin' || authStore.user?.role === 'superadmin') {
    await Promise.all([
      dashboardStore.fetchStatusSummary('all'),
      logsStore.fetchSystemLogs({ limit: 10, offset: 0 }),
    ])
    logsStore.connectRealtime()
  }
})

const formatDate = (date: string) =>
  new Date(date).toLocaleString('id-ID', {
    year: 'numeric',
    month: 'short',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
  })
</script>

<template>
  <div style="max-width: 1200px; margin: 0 auto">
    <!-- Stats Grid -->
    <div
      style="
        display: grid;
        grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
        gap: 1rem;
        margin-bottom: 2rem;
      "
    >
      <div
        v-for="card in cards"
        :key="card.label"
        style="
          background-color: white;
          border-radius: 0.5rem;
          padding: 1.5rem;
          box-shadow: 0 1px 3px 0 rgba(0, 0, 0, 0.1);
        "
      >
        <p style="font-size: 0.875rem; color: #6b7280; margin-bottom: 0.5rem">{{ card.label }}</p>
        <p :style="{ fontSize: '2.5rem', fontWeight: '700', color: card.color }">{{ card.value }}</p>
      </div>
    </div>

    <!-- Recent Activity -->
    <div
      style="
        background-color: white;
        border-radius: 0.5rem;
        padding: 1.5rem;
        box-shadow: 0 1px 3px 0 rgba(0, 0, 0, 0.1);
      "
    >
      <h2 style="font-size: 1.25rem; font-weight: 600; color: #111827; margin-bottom: 1rem">
        Aktivitas Terbaru
      </h2>
      <p v-if="logsStore.isLoading" style="color: #6b7280; text-align: center; padding: 2rem">
        Memuat ringkasan dashboard...
      </p>
      <p
        v-else-if="logsStore.error"
        style="color: #dc2626; text-align: center; padding: 2rem"
      >
        {{ logsStore.error }}
      </p>
      <div v-else-if="logsStore.recentActivityItems.length > 0" class="space-y-3">
        <div
          v-for="log in logsStore.recentActivityItems"
          :key="log.sysLogId"
          class="flex items-start justify-between gap-4 rounded-xl border border-gray-100 px-4 py-3"
        >
          <div class="min-w-0">
            <p style="font-size: 0.875rem; font-weight: 600; color: #111827">
              {{ log.sysUserNama }}
              <span style="font-weight: 500; color: #2563eb">{{ log.sysAction }}</span>
              <span v-if="log.sysTargetTable" style="color: #6b7280">di {{ log.sysTargetTable }}</span>
            </p>
            <p style="font-size: 0.875rem; color: #6b7280; margin-top: 0.35rem">
              {{ log.sysNotes }}
            </p>
          </div>
          <div style="text-align: right; white-space: nowrap">
            <p style="font-size: 0.75rem; color: #9ca3af">{{ formatDate(log.createdAt) }}</p>
            <p style="font-size: 0.75rem; color: #6b7280; margin-top: 0.35rem">{{ log.sysUserRole }}</p>
          </div>
        </div>
        <div style="display: flex; justify-content: space-between; align-items: center; padding-top: 0.5rem">
          <p style="font-size: 0.75rem; color: #6b7280">
            {{ logsStore.isRealtimeConnected ? 'Realtime aktif' : 'Realtime belum terhubung' }}
          </p>
          <RouterLink to="/logs" style="font-size: 0.875rem; font-weight: 600; color: #2563eb">
            Lihat semua logs
          </RouterLink>
        </div>
      </div>
      <p v-else style="color: #6b7280; text-align: center; padding: 2rem">Belum ada aktivitas.</p>
    </div>
  </div>
</template>
