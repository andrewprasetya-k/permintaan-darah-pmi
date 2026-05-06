<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue'
import { RouterLink, useRoute } from 'vue-router'
import AppModal from '@/components/feedback/AppModal.vue'
import StatusBadge from '@/components/rs/StatusBadge.vue'
import { useFeedbackStore } from '@/stores/feedback'
import { useMyRequestsStore } from '@/stores/my-requests'
import { bloodLabel, formatDate, formatDateTime, statusDescriptions } from '@/utils/format'

const route = useRoute()
const requestsStore = useMyRequestsStore()
const feedbackStore = useFeedbackStore()

const cancelReason = ref('')
const isCancelOpen = ref(false)
const isPickupOpen = ref(false)

const requestId = computed(() => String(route.params.id || ''))
const request = computed(() => requestsStore.selectedRequest)

const loadRequest = async () => {
  if (!requestId.value) {
    return
  }

  await requestsStore.fetchById(requestId.value).catch((err) => {
    feedbackStore.showFlag({
      title: 'Gagal memuat detail',
      message: err instanceof Error ? err.message : 'Coba muat ulang halaman.',
      variant: 'error',
    })
  })
}

const confirmCancel = async () => {
  if (!request.value || !cancelReason.value.trim()) {
    return
  }

  try {
    await requestsStore.cancelRequest(request.value.permintaanDarahId, cancelReason.value.trim())
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
  if (!request.value) {
    return
  }

  try {
    await requestsStore.confirmPickup(request.value.permintaanDarahId)
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

onMounted(loadRequest)
watch(requestId, loadRequest)
</script>

<template>
  <section>
    <div class="page-header">
      <div>
        <p class="page-eyebrow">Detail Permintaan</p>
        <h1 class="page-title">{{ request?.namaPasien || 'Memuat detail' }}</h1>
        <p v-if="request" class="page-subtitle">
          {{ request.permintaanDarahId }} - {{ statusDescriptions[request.status] }}
        </p>
      </div>
      <div v-if="request" class="detail-actions">
        <RouterLink class="btn btn-secondary" to="/requests">Kembali</RouterLink>
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
          @click="isPickupOpen = true"
        >
          Konfirmasi Selesai
        </button>
        <button
          v-if="requestsStore.canCancel(request)"
          type="button"
          class="btn btn-danger"
          @click="isCancelOpen = true"
        >
          Batalkan
        </button>
      </div>
    </div>

    <div v-if="requestsStore.isLoading && !request" class="empty-state">
      <div>
        <h2>Memuat detail</h2>
        <p>Data permintaan sedang diambil dari backend.</p>
      </div>
    </div>

    <div v-else-if="!request" class="empty-state">
      <div>
        <h2>Detail tidak tersedia</h2>
        <p>Permintaan tidak ditemukan atau tidak dapat diakses.</p>
      </div>
    </div>

    <template v-else>
      <section class="detail-grid">
        <article class="card detail-card status-card">
          <div>
            <span>Status</span>
            <StatusBadge :status="request.status" />
          </div>
          <p v-if="request.cancelReason">Alasan batal: {{ request.cancelReason }}</p>
          <p>Dibuat {{ formatDateTime(request.createdAt) }}</p>
        </article>

        <article class="card detail-card">
          <h2 class="section-title">Data Pasien</h2>
          <dl class="detail-list">
            <div>
              <dt>Nama Pasien</dt>
              <dd>{{ request.namaPasien }}</dd>
            </div>
            <div>
              <dt>No. RM</dt>
              <dd>{{ request.noRM || '-' }}</dd>
            </div>
            <div>
              <dt>Tempat/Tanggal Lahir</dt>
              <dd>{{ request.tempatLahir }}, {{ formatDate(request.tanggalLahir) }}</dd>
            </div>
            <div>
              <dt>Jenis Kelamin</dt>
              <dd>{{ request.jenisKelamin === 'L' ? 'Laki-laki' : 'Perempuan' }}</dd>
            </div>
            <div>
              <dt>Golongan Darah</dt>
              <dd>{{ bloodLabel(request.golonganDarah, request.rhesusDarah) }}</dd>
            </div>
          </dl>
        </article>

        <article class="card detail-card">
          <h2 class="section-title">Data Medis</h2>
          <dl class="detail-list">
            <div>
              <dt>Hemoglobin</dt>
              <dd>{{ request.hemoglobin ? `${request.hemoglobin} g/dL` : '-' }}</dd>
            </div>
            <div>
              <dt>Ruang/Bagian/Kelas</dt>
              <dd>{{ request.ruangBagianKelas || '-' }}</dd>
            </div>
            <div>
              <dt>Pernah Transfusi</dt>
              <dd>{{ request.pernahTransfusi ? 'Ya' : 'Tidak' }}</dd>
            </div>
            <div>
              <dt>Indikasi Transfusi</dt>
              <dd>{{ request.indikasiTransfusi || '-' }}</dd>
            </div>
            <div>
              <dt>Pernah Hamil</dt>
              <dd>{{ request.pernahHamil || '-' }}</dd>
            </div>
            <div>
              <dt>Tanggal Permintaan</dt>
              <dd>{{ formatDate(request.tanggalPermintaan) }}</dd>
            </div>
          </dl>
        </article>
      </section>

      <section class="card detail-card component-card">
        <h2 class="section-title">Komponen Darah</h2>
        <div v-if="!request.detailPermintaanDarah?.length" class="empty-row">
          Tidak ada detail komponen pada permintaan ini.
        </div>
        <div v-else class="table-wrap">
          <table class="data-table">
            <thead>
              <tr>
                <th>Komponen</th>
                <th>Darah</th>
                <th>Jumlah</th>
                <th>Diperlukan</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="detail in request.detailPermintaanDarah" :key="detail.detailId">
                <td>{{ detail.komponenNama }}</td>
                <td>{{ bloodLabel(detail.golonganDarah, detail.rhesusDarah) }}</td>
                <td>{{ detail.jumlahKantong }} kantong</td>
                <td>{{ formatDate(detail.tanggalDiperlukan) }}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </section>
    </template>

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
.detail-actions {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
}

.detail-grid {
  display: grid;
  grid-template-columns: 280px minmax(0, 1fr) minmax(0, 1fr);
  gap: 18px;
}

.detail-card {
  padding: 18px;
}

.status-card {
  display: grid;
  align-content: start;
  gap: 14px;
}

.status-card span {
  display: block;
  margin-bottom: 8px;
  color: var(--text-muted);
  font-size: 12px;
  font-weight: 900;
  text-transform: uppercase;
}

.status-card p {
  margin: 0;
  color: var(--text-soft);
}

.detail-list {
  display: grid;
  gap: 14px;
  margin: 16px 0 0;
}

.detail-list div {
  display: grid;
  gap: 4px;
}

.detail-list dt {
  color: var(--text-muted);
  font-size: 12px;
  font-weight: 850;
}

.detail-list dd {
  margin: 0;
  overflow-wrap: anywhere;
  font-weight: 750;
}

.component-card {
  margin-top: 18px;
}

.empty-row {
  margin-top: 16px;
  border: 1px dashed var(--line-strong);
  border-radius: 8px;
  padding: 18px;
  color: var(--text-soft);
}

@media (max-width: 1120px) {
  .detail-grid {
    grid-template-columns: 1fr;
  }
}

@media (max-width: 640px) {
  .detail-actions {
    width: 100%;
  }

  .detail-actions .btn {
    flex: 1 1 150px;
  }
}
</style>
