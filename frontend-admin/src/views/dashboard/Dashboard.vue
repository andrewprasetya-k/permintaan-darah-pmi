<script setup lang="ts">
import { computed, onMounted } from 'vue'
import { useAuthStore } from '@/stores/auth'
import { useDashboardStore } from '@/stores/dashboard'

const authStore = useAuthStore()
const dashboardStore = useDashboardStore()

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
    await dashboardStore.fetchStatusSummary('all')
  }
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
      <p v-if="dashboardStore.isLoading" style="color: #6b7280; text-align: center; padding: 2rem">
        Memuat ringkasan dashboard...
      </p>
      <p
        v-else-if="dashboardStore.error"
        style="color: #dc2626; text-align: center; padding: 2rem"
      >
        {{ dashboardStore.error }}
      </p>
      <p v-else style="color: #6b7280; text-align: center; padding: 2rem">
        Ringkasan status permintaan darah telah dimuat.
      </p>
    </div>
  </div>
</template>
