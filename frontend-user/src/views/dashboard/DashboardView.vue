<script setup lang="ts">
import { computed, onMounted } from 'vue'
import { RouterLink } from 'vue-router'
import { Plus } from '@lucide/vue'
import StatusBadge from '@/components/rs/StatusBadge.vue'
import { useAuthStore } from '@/stores/auth'
import { useDashboardStore } from '@/stores/dashboard'
import { useMyRequestsStore } from '@/stores/my-requests'
import type { PermintaanStatus } from '@/types/models'
import { bloodLabel, formatDate, statusLabels } from '@/utils/format'

const authStore = useAuthStore()
const dashboardStore = useDashboardStore()
const requestsStore = useMyRequestsStore()

const summaryCards = computed<
  Array<{ key: PermintaanStatus | 'total'; label: string; value: number; color: string }>
>(() => [
  {
    key: 'dibuat',
    label: statusLabels.dibuat,
    value: dashboardStore.summary.dibuat,
    color: '#2563eb',
  },
  {
    key: 'diproses',
    label: statusLabels.diproses,
    value: dashboardStore.summary.diproses,
    color: '#f59e0b',
  },
  {
    key: 'bisa_diambil',
    label: statusLabels.bisa_diambil,
    value: dashboardStore.summary.bisaDiambil,
    color: '#10b981',
  },
  {
    key: 'selesai',
    label: statusLabels.selesai,
    value: dashboardStore.summary.selesai,
    color: '#8b5cf6',
  },
  {
    key: 'dibatalkan',
    label: statusLabels.dibatalkan,
    value: dashboardStore.summary.dibatalkan,
    color: '#ef4444',
  },
  { key: 'total', label: 'Total', value: dashboardStore.summary.total, color: '#111827' },
])

onMounted(async () => {
  const rumahSakitId = authStore.user?.id || 'me'
  await Promise.allSettled([dashboardStore.fetchSummary(rumahSakitId), requestsStore.fetchAll()])
})
</script>

<template>
  <section>
    <div class="page-header">
      <div>
        <p class="page-eyebrow">Ringkasan</p>
        <h1 class="page-title">Dashboard Permintaan Darah</h1>
        <p class="page-subtitle">
          Pantau status permintaan aktif dan lanjutkan permintaan yang siap diproses.
        </p>
      </div>
      <RouterLink class="btn btn-primary" to="/requests/new">
        <Plus :size="16" />
        Buat Permintaan
      </RouterLink>
    </div>

    <div class="summary-grid">
      <article v-for="item in summaryCards" :key="item.key" class="summary-card card">
        <span>{{ item.label }}</span>
        <strong :style="{ color: item.color }">{{ item.value }}</strong>
      </article>
    </div>

    <div class="dashboard-grid">
      <section class="card dashboard-panel">
        <div class="panel-header">
          <div>
            <h2 class="section-title">Permintaan Terbaru</h2>
            <p>Data dari daftar permintaan milik rumah sakit.</p>
          </div>
          <RouterLink class="btn btn-secondary" to="/requests">Lihat Semua</RouterLink>
        </div>

        <div v-if="requestsStore.recentRequests.length === 0" class="empty-state compact-empty">
          <div>
            <h3>Belum ada permintaan</h3>
            <p>Buat permintaan darah baru saat data pasien sudah siap.</p>
          </div>
        </div>

        <div v-else class="table-wrap">
          <table class="data-table">
            <thead>
              <tr>
                <th>Pasien</th>
                <th>Darah</th>
                <th>Tanggal</th>
                <th>Status</th>
                <th></th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="request in requestsStore.recentRequests" :key="request.permintaanDarahId">
                <td>
                  <strong>{{ request.namaPasien }}</strong>
                </td>
                <td>{{ bloodLabel(request.golonganDarah, request.rhesusDarah) }}</td>
                <td>{{ formatDate(request.tanggalPermintaan) }}</td>
                <td><StatusBadge :status="request.status" /></td>
                <td>
                  <RouterLink
                    class="btn btn-secondary"
                    :to="`/requests/${request.permintaanDarahId}`"
                  >
                    Detail
                  </RouterLink>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </section>

      <section class="card dashboard-panel action-panel">
        <h2 class="section-title">Tindak Lanjut</h2>
        <RouterLink class="action-link" to="/requests?status=bisa_diambil">
          <strong>{{ dashboardStore.summary.bisaDiambil }}</strong>
          <span>Permintaan bisa diambil</span>
        </RouterLink>
        <RouterLink class="action-link" to="/requests?status=diproses">
          <strong>{{ dashboardStore.summary.diproses }}</strong>
          <span>Permintaan sedang diproses</span>
        </RouterLink>
        <RouterLink class="action-link" to="/requests?status=dibuat">
          <strong>{{ dashboardStore.summary.dibuat }}</strong>
          <span>Permintaan baru dibuat</span>
        </RouterLink>
      </section>
    </div>
  </section>
</template>

<style scoped>
.summary-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(180px, 1fr));
  gap: 16px;
  margin-bottom: 24px;
}

.summary-card {
  min-height: 132px;
  padding: 24px;
}

.summary-card span {
  color: var(--text-muted);
  font-size: 14px;
  font-weight: 400;
}

.summary-card strong {
  display: block;
  margin-top: 8px;
  font-size: 40px;
  font-weight: 700;
  line-height: 1;
}

.dashboard-grid {
  display: grid;
  grid-template-columns: minmax(0, 1fr) 320px;
  gap: 20px;
}

.dashboard-panel {
  padding: 24px;
}

.panel-header {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 16px;
  margin-bottom: 8px;
}

.panel-header p {
  margin: 5px 0 0;
  color: var(--text-soft);
}

.compact-empty {
  min-height: 180px;
}

.action-panel {
  display: grid;
  align-content: start;
  gap: 12px;
}

.action-link {
  display: flex;
  align-items: center;
  gap: 12px;
  border: 1px solid #f3f4f6;
  border-radius: 12px;
  padding: 14px;
}

.action-link:hover {
  border-color: var(--line-strong);
  background: var(--surface-muted);
}

.action-link strong {
  display: flex;
  width: 44px;
  height: 44px;
  align-items: center;
  justify-content: center;
  border-radius: 12px;
  background: #eff6ff;
  color: #2563eb;
  font-size: 18px;
  font-weight: 700;
}

.action-link span {
  color: var(--text-soft);
  font-weight: 800;
}

@media (max-width: 1120px) {
  .summary-grid {
    grid-template-columns: repeat(3, minmax(0, 1fr));
  }

  .dashboard-grid {
    grid-template-columns: 1fr;
  }
}

@media (max-width: 680px) {
  .summary-grid {
    grid-template-columns: repeat(2, minmax(0, 1fr));
  }

  .summary-card {
    min-height: 118px;
  }

  .summary-card strong {
    font-size: 28px;
  }
}
</style>
