<script setup lang="ts">
import { computed, onMounted } from 'vue'
import { RouterLink } from 'vue-router'
import StatusBadge from '@/components/rs/StatusBadge.vue'
import { useAuthStore } from '@/stores/auth'
import { useDashboardStore } from '@/stores/dashboard'
import { useMyRequestsStore } from '@/stores/my-requests'
import type { PermintaanStatus } from '@/types/models'
import { bloodLabel, formatDate, statusLabels } from '@/utils/format'

const authStore = useAuthStore()
const dashboardStore = useDashboardStore()
const requestsStore = useMyRequestsStore()

const summaryCards = computed<Array<{ key: PermintaanStatus; value: number }>>(() => [
  { key: 'dibuat', value: dashboardStore.summary.dibuat },
  { key: 'diproses', value: dashboardStore.summary.diproses },
  { key: 'bisa_diambil', value: dashboardStore.summary.bisaDiambil },
  { key: 'selesai', value: dashboardStore.summary.selesai },
  { key: 'dibatalkan', value: dashboardStore.summary.dibatalkan },
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
      <RouterLink class="btn btn-primary" to="/requests/new">+ Buat Permintaan</RouterLink>
    </div>

    <div class="summary-grid">
      <article class="summary-card card total-card">
        <span>Total</span>
        <strong>{{ dashboardStore.summary.total }}</strong>
        <p>Semua permintaan milik rumah sakit ini</p>
      </article>

      <article
        v-for="item in summaryCards"
        :key="item.key"
        class="summary-card card"
        :class="`summary-${item.key}`"
      >
        <span>{{ statusLabels[item.key] }}</span>
        <strong>{{ item.value }}</strong>
        <p>{{ item.key === 'bisa_diambil' ? 'Siap ditindaklanjuti' : 'Permintaan' }}</p>
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
                  <RouterLink class="btn btn-secondary" :to="`/requests/${request.permintaanDarahId}`">
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
  grid-template-columns: repeat(6, minmax(0, 1fr));
  gap: 12px;
  margin-bottom: 20px;
}

.summary-card {
  min-height: 132px;
  padding: 16px;
}

.summary-card span {
  color: var(--text-muted);
  font-size: 12px;
  font-weight: 900;
  text-transform: uppercase;
}

.summary-card strong {
  display: block;
  margin-top: 12px;
  font-size: 34px;
  font-weight: 900;
  line-height: 1;
}

.summary-card p {
  margin: 10px 0 0;
  color: var(--text-soft);
  font-size: 13px;
}

.total-card {
  grid-column: span 2;
  background: #111827;
  color: #ffffff;
}

.total-card span,
.total-card p {
  color: rgba(255, 255, 255, 0.72);
}

.summary-bisa_diambil {
  border-color: #bbf7d0;
}

.summary-dibatalkan {
  border-color: #fecaca;
}

.dashboard-grid {
  display: grid;
  grid-template-columns: minmax(0, 1fr) 320px;
  gap: 20px;
}

.dashboard-panel {
  padding: 18px;
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
  border: 1px solid var(--line);
  border-radius: 8px;
  padding: 14px;
}

.action-link:hover {
  border-color: var(--line-strong);
  background: var(--surface-muted);
}

.action-link strong {
  display: grid;
  width: 44px;
  height: 44px;
  place-items: center;
  border-radius: 8px;
  background: var(--red-soft);
  color: #b91c1c;
  font-size: 18px;
  font-weight: 900;
}

.action-link span {
  color: var(--text-soft);
  font-weight: 800;
}

@media (max-width: 1120px) {
  .summary-grid {
    grid-template-columns: repeat(3, minmax(0, 1fr));
  }

  .total-card {
    grid-column: span 1;
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
