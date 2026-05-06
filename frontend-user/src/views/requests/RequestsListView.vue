<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { RouterLink, useRoute } from 'vue-router'
import { Plus, Search } from '@lucide/vue'
import AppModal from '@/components/feedback/AppModal.vue'
import StatusBadge from '@/components/rs/StatusBadge.vue'
import { useFeedbackStore } from '@/stores/feedback'
import { useMyRequestsStore } from '@/stores/my-requests'
import type { PermintaanDarahListItem, PermintaanStatus } from '@/types/models'
import { bloodLabel, formatDate, statusLabels } from '@/utils/format'

const route = useRoute()
const requestsStore = useMyRequestsStore()
const feedbackStore = useFeedbackStore()

const search = ref('')
const statusFilter = ref<PermintaanStatus | 'all'>('all')
const selectedRequest = ref<PermintaanDarahListItem | null>(null)
const cancelReason = ref('')
const isCancelOpen = ref(false)
const isPickupOpen = ref(false)

const statuses: Array<{ value: PermintaanStatus | 'all'; label: string }> = [
  { value: 'all', label: 'Semua Status' },
  { value: 'dibuat', label: statusLabels.dibuat },
  { value: 'diproses', label: statusLabels.diproses },
  { value: 'bisa_diambil', label: statusLabels.bisa_diambil },
  { value: 'selesai', label: statusLabels.selesai },
  { value: 'dibatalkan', label: statusLabels.dibatalkan },
]

const filteredRequests = computed(() => {
  const keyword = search.value.trim().toLowerCase()

  return requestsStore.requests.filter((request) => {
    const matchesStatus = statusFilter.value === 'all' || request.status === statusFilter.value
    const matchesKeyword =
      !keyword ||
      request.namaPasien.toLowerCase().includes(keyword) ||
      request.permintaanDarahId.toLowerCase().includes(keyword)

    return matchesStatus && matchesKeyword
  })
})

const openCancel = (request: PermintaanDarahListItem) => {
  selectedRequest.value = request
  cancelReason.value = ''
  isCancelOpen.value = true
}

const openPickup = (request: PermintaanDarahListItem) => {
  selectedRequest.value = request
  isPickupOpen.value = true
}

const confirmCancel = async () => {
  if (!selectedRequest.value || !cancelReason.value.trim()) {
    return
  }

  try {
    await requestsStore.cancelRequest(
      selectedRequest.value.permintaanDarahId,
      cancelReason.value.trim(),
    )
    isCancelOpen.value = false
    feedbackStore.showFlag({
      title: 'Permintaan dibatalkan',
      message: 'Status permintaan sudah diperbarui.',
      variant: 'success',
    })
  } catch (err) {
    feedbackStore.showFlag({
      title: 'Gagal membatalkan',
      message: err instanceof Error ? err.message : 'Coba lagi nanti.',
      variant: 'error',
    })
  }
}

const confirmPickup = async () => {
  if (!selectedRequest.value) {
    return
  }

  try {
    await requestsStore.confirmPickup(selectedRequest.value.permintaanDarahId)
    isPickupOpen.value = false
    feedbackStore.showFlag({
      title: 'Pengambilan dikonfirmasi',
      message: 'Permintaan sudah ditandai selesai.',
      variant: 'success',
    })
  } catch (err) {
    feedbackStore.showFlag({
      title: 'Gagal mengonfirmasi',
      message: err instanceof Error ? err.message : 'Coba lagi nanti.',
      variant: 'error',
    })
  }
}

onMounted(async () => {
  const status = route.query.status
  if (
    status === 'dibuat' ||
    status === 'diproses' ||
    status === 'bisa_diambil' ||
    status === 'selesai' ||
    status === 'dibatalkan'
  ) {
    statusFilter.value = status
  }

  await requestsStore.fetchAll().catch((err) => {
    feedbackStore.showFlag({
      title: 'Gagal memuat permintaan',
      message: err instanceof Error ? err.message : 'Coba muat ulang halaman.',
      variant: 'error',
    })
  })
})
</script>

<template>
  <section>
    <div class="page-header">
      <div>
        <p class="page-eyebrow">Permintaan</p>
        <h1 class="page-title">Permintaan Saya</h1>
        <p class="page-subtitle">Daftar permintaan darah milik rumah sakit yang sedang login.</p>
      </div>
      <RouterLink class="btn btn-primary" to="/requests/new">
        <Plus :size="16" />
        Buat Permintaan
      </RouterLink>
    </div>

    <section class="card list-panel">
      <div class="filters">
        <div class="form-field search-field">
          <label class="sr-only" for="search">Cari</label>
          <Search :size="16" />
          <input
            id="search"
            v-model="search"
            class="form-input"
            type="search"
            placeholder="Nama pasien atau ID permintaan"
          />
        </div>
        <div class="form-field">
          <label class="form-label" for="status">Status</label>
          <select id="status" v-model="statusFilter" class="form-select">
            <option v-for="status in statuses" :key="status.value" :value="status.value">
              {{ status.label }}
            </option>
          </select>
        </div>
      </div>

      <div v-if="requestsStore.isLoading" class="empty-state">
        <div>
          <h2>Memuat permintaan</h2>
          <p>Data sedang diambil dari backend.</p>
        </div>
      </div>

      <div v-else-if="filteredRequests.length === 0" class="empty-state">
        <div>
          <h2>Tidak ada data</h2>
          <p>Ubah filter atau buat permintaan baru.</p>
        </div>
      </div>

      <template v-else>
        <div class="table-wrap desktop-list">
          <table class="data-table">
            <thead>
              <tr>
                <th>Pasien</th>
                <th>Darah</th>
                <th>Tanggal</th>
                <th>Status</th>
                <th>Aksi</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="request in filteredRequests" :key="request.permintaanDarahId">
                <td>
                  <strong>{{ request.namaPasien }}</strong>
                  <span class="request-id">{{ request.permintaanDarahId }}</span>
                </td>
                <td>{{ bloodLabel(request.golonganDarah, request.rhesusDarah) }}</td>
                <td>{{ formatDate(request.tanggalPermintaan) }}</td>
                <td><StatusBadge :status="request.status" /></td>
                <td>
                  <div class="row-actions">
                    <RouterLink
                      class="btn btn-secondary"
                      :to="`/requests/${request.permintaanDarahId}`"
                    >
                      Detail
                    </RouterLink>
                    <RouterLink
                      v-if="requestsStore.canEdit(request)"
                      class="btn btn-secondary"
                      :to="`/requests/${request.permintaanDarahId}/edit`"
                    >
                      Edit
                    </RouterLink>
                    <button
                      v-if="requestsStore.canConfirmPickup(request)"
                      type="button"
                      class="btn btn-success"
                      @click="openPickup(request)"
                    >
                      Selesai
                    </button>
                    <button
                      v-if="requestsStore.canCancel(request)"
                      type="button"
                      class="btn btn-danger"
                      @click="openCancel(request)"
                    >
                      Batal
                    </button>
                  </div>
                </td>
              </tr>
            </tbody>
          </table>
        </div>

        <div class="mobile-list">
          <article
            v-for="request in filteredRequests"
            :key="request.permintaanDarahId"
            class="request-card"
          >
            <div class="request-card-head">
              <div>
                <h2>{{ request.namaPasien }}</h2>
                <p>{{ formatDate(request.tanggalPermintaan) }}</p>
              </div>
              <StatusBadge :status="request.status" />
            </div>
            <dl>
              <div>
                <dt>Darah</dt>
                <dd>{{ bloodLabel(request.golonganDarah, request.rhesusDarah) }}</dd>
              </div>
              <div>
                <dt>ID</dt>
                <dd>{{ request.permintaanDarahId }}</dd>
              </div>
            </dl>
            <div class="card-actions">
              <RouterLink class="btn btn-secondary" :to="`/requests/${request.permintaanDarahId}`">
                Detail
              </RouterLink>
              <RouterLink
                v-if="requestsStore.canEdit(request)"
                class="btn btn-secondary"
                :to="`/requests/${request.permintaanDarahId}/edit`"
              >
                Edit
              </RouterLink>
              <button
                v-if="requestsStore.canConfirmPickup(request)"
                type="button"
                class="btn btn-success"
                @click="openPickup(request)"
              >
                Selesai
              </button>
              <button
                v-if="requestsStore.canCancel(request)"
                type="button"
                class="btn btn-danger"
                @click="openCancel(request)"
              >
                Batal
              </button>
            </div>
          </article>
        </div>
      </template>
    </section>

    <AppModal
      :is-open="isCancelOpen"
      title="Batalkan permintaan?"
      description="Alasan pembatalan akan tersimpan pada permintaan."
      width="sm"
      @close="isCancelOpen = false"
    >
      <div class="form-field">
        <label class="form-label" for="cancelReason">Alasan pembatalan</label>
        <textarea id="cancelReason" v-model="cancelReason" class="form-textarea" required />
      </div>
      <template #footer>
        <button type="button" class="btn btn-secondary" @click="isCancelOpen = false">Batal</button>
        <button
          type="button"
          class="btn btn-danger"
          :disabled="requestsStore.isSubmitting || !cancelReason.trim()"
          @click="confirmCancel"
        >
          Batalkan Permintaan
        </button>
      </template>
    </AppModal>

    <AppModal
      :is-open="isPickupOpen"
      title="Konfirmasi pengambilan?"
      description="Permintaan akan ditandai selesai."
      width="sm"
      @close="isPickupOpen = false"
    >
      <template #footer>
        <button type="button" class="btn btn-secondary" @click="isPickupOpen = false">Batal</button>
        <button
          type="button"
          class="btn btn-success"
          :disabled="requestsStore.isSubmitting"
          @click="confirmPickup"
        >
          Konfirmasi
        </button>
      </template>
    </AppModal>
  </section>
</template>

<style scoped>
.list-panel {
  overflow: hidden;
}

.filters {
  display: grid;
  grid-template-columns: minmax(0, 1fr) 220px;
  gap: 14px;
  border-bottom: 1px solid #f3f4f6;
  padding: 20px;
}

.list-panel > .empty-state {
  margin: 20px;
}

.search-field {
  position: relative;
}

.search-field svg {
  position: absolute;
  left: 12px;
  top: 50%;
  z-index: 1;
  color: #9ca3af;
  transform: translateY(-50%);
}

.search-field .form-input {
  padding-left: 38px;
}

.request-id {
  display: block;
  margin-top: 4px;
  color: var(--text-muted);
  font-size: 12px;
}

.row-actions,
.card-actions {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}

.mobile-list {
  display: none;
  padding: 20px;
}

.request-card {
  border: 1px solid #f3f4f6;
  border-radius: 16px;
  padding: 16px;
  background: #ffffff;
}

.request-card + .request-card {
  margin-top: 12px;
}

.request-card-head {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 12px;
}

.request-card h2,
.request-card p,
.request-card dl,
.request-card dd {
  margin: 0;
}

.request-card h2 {
  font-size: 16px;
  font-weight: 900;
}

.request-card p {
  margin-top: 3px;
  color: var(--text-muted);
}

.request-card dl {
  display: grid;
  gap: 10px;
  margin: 14px 0;
}

.request-card dt {
  color: var(--text-muted);
  font-size: 12px;
  font-weight: 800;
}

.request-card dd {
  overflow-wrap: anywhere;
  font-weight: 750;
}

@media (max-width: 780px) {
  .filters {
    grid-template-columns: 1fr;
  }

  .desktop-list {
    display: none;
  }

  .mobile-list {
    display: block;
  }
}
</style>
