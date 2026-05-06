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
import { btn, ui } from '@/utils/ui'

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
    <div :class="ui.pageHeader">
      <RouterLink :class="btn('btnPrimary')" to="/requests/new">
        <Plus :size="16" />
        Buat Permintaan
      </RouterLink>
    </div>

    <div
      class="mb-6 grid grid-cols-[repeat(auto-fit,minmax(180px,1fr))] gap-4 max-xl:grid-cols-3 max-md:grid-cols-2"
    >
      <article v-for="item in summaryCards" :key="item.key" :class="[ui.card, 'min-h-32 p-6']">
        <span class="text-sm text-gray-500">{{ item.label }}</span>
        <strong
          class="mt-2 block text-[40px] font-bold leading-none max-sm:text-[28px]"
          :style="{ color: item.color }"
        >
          {{ item.value }}
        </strong>
      </article>
    </div>

    <div class="grid grid-cols-[minmax(0,1fr)_320px] gap-5 max-xl:grid-cols-1">
      <section :class="[ui.card, 'p-6']">
        <div class="mb-2 flex items-start justify-between gap-4">
          <div>
            <h2 :class="ui.sectionTitle">Permintaan Terbaru</h2>
            <p class="mt-1 text-sm text-gray-600">Data dari daftar permintaan milik rumah sakit.</p>
          </div>
          <RouterLink :class="btn('btnSecondary')" to="/requests">Lihat Semua</RouterLink>
        </div>

        <div v-if="requestsStore.recentRequests.length === 0" :class="[ui.emptyState, 'min-h-44']">
          <div>
            <h3 :class="ui.emptyTitle">Belum ada permintaan</h3>
            <p :class="ui.emptyCopy">Buat permintaan darah baru saat data pasien sudah siap.</p>
          </div>
        </div>

        <div v-else :class="ui.tableWrap">
          <table :class="ui.table">
            <thead>
              <tr>
                <th :class="ui.th">Pasien</th>
                <th :class="ui.th">Darah</th>
                <th :class="ui.th">Tanggal</th>
                <th :class="ui.th">Status</th>
                <th :class="ui.th"></th>
              </tr>
            </thead>
            <tbody>
              <tr
                v-for="request in requestsStore.recentRequests"
                :key="request.permintaanDarahId"
                class="transition-colors hover:bg-gray-50"
              >
                <td :class="ui.td">
                  <strong>{{ request.namaPasien }}</strong>
                </td>
                <td :class="ui.td">{{ bloodLabel(request.golonganDarah, request.rhesusDarah) }}</td>
                <td :class="ui.td">{{ formatDate(request.tanggalPermintaan) }}</td>
                <td :class="ui.td"><StatusBadge :status="request.status" /></td>
                <td :class="ui.td">
                  <RouterLink
                    :class="btn('btnSecondary')"
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

      <section :class="[ui.card, 'grid content-start gap-3 p-6']">
        <h2 :class="ui.sectionTitle">Tindak Lanjut</h2>
        <RouterLink
          class="flex items-center gap-3 rounded-xl border border-gray-100 p-3.5 transition hover:bg-gray-50"
          to="/requests?status=bisa_diambil"
        >
          <strong
            class="flex h-11 w-11 items-center justify-center rounded-xl bg-blue-50 text-lg font-bold text-blue-600"
          >
            {{ dashboardStore.summary.bisaDiambil }}
          </strong>
          <span class="font-semibold text-gray-600">Permintaan bisa diambil</span>
        </RouterLink>
        <RouterLink
          class="flex items-center gap-3 rounded-xl border border-gray-100 p-3.5 transition hover:bg-gray-50"
          to="/requests?status=diproses"
        >
          <strong
            class="flex h-11 w-11 items-center justify-center rounded-xl bg-blue-50 text-lg font-bold text-blue-600"
          >
            {{ dashboardStore.summary.diproses }}
          </strong>
          <span class="font-semibold text-gray-600">Permintaan sedang diproses</span>
        </RouterLink>
        <RouterLink
          class="flex items-center gap-3 rounded-xl border border-gray-100 p-3.5 transition hover:bg-gray-50"
          to="/requests?status=dibuat"
        >
          <strong
            class="flex h-11 w-11 items-center justify-center rounded-xl bg-blue-50 text-lg font-bold text-blue-600"
          >
            {{ dashboardStore.summary.dibuat }}
          </strong>
          <span class="font-semibold text-gray-600">Permintaan baru dibuat</span>
        </RouterLink>
      </section>
    </div>
  </section>
</template>
