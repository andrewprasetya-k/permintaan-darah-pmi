<script setup lang="ts">
import { computed, onBeforeUnmount, onMounted, ref, watch } from 'vue'
import { useRoute } from 'vue-router'
import { CheckCircle2, CircleSlash, Download, Pencil, Undo2 } from '@lucide/vue'
import AppModal from '@/components/feedback/AppModal.vue'
import StatusBadge from '@/components/rs/StatusBadge.vue'
import { useFeedbackStore } from '@/stores/feedback'
import { useMyRequestsStore } from '@/stores/my-requests'
import { useSetPageHeaderActions } from '@/composables/usePageHeader'
import { bloodLabel, formatDate, formatDateTime, statusDescriptions } from '@/utils/format'
import { btn, ui } from '@/utils/ui'

const route = useRoute()
const requestsStore = useMyRequestsStore()
const feedbackStore = useFeedbackStore()
const { clearActions, setActions } = useSetPageHeaderActions()

const cancelReason = ref('')
const isCancelOpen = ref(false)
const isPickupOpen = ref(false)

const requestId = computed(() => String(route.params.id || ''))
const request = computed(() => requestsStore.selectedRequest)

const loadRequest = async () => {
  if (!requestId.value) return

  await requestsStore.fetchById(requestId.value).catch((err) => {
    feedbackStore.showFlag({
      title: 'Gagal memuat detail',
      message: err instanceof Error ? err.message : 'Coba muat ulang halaman.',
      variant: 'error',
    })
  })
}

const syncHeaderActions = () => {
  setActions([{ label: 'Kembali', to: '/requests', icon: Undo2, variant: 'secondary' }])
}

const confirmCancel = async () => {
  if (!request.value || !cancelReason.value.trim()) return

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
  if (!request.value) return

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

const triggerDownload = (blob: Blob, filename: string) => {
  const url = URL.createObjectURL(blob)
  const link = document.createElement('a')
  link.href = url
  link.download = filename
  document.body.appendChild(link)
  link.click()
  link.remove()
  URL.revokeObjectURL(url)
}

const downloadBlanko = async () => {
  if (!request.value) return

  try {
    const response = await requestsStore.downloadBlanko(request.value.permintaanDarahId)
    triggerDownload(
      response.blob,
      response.filename || `blanko-permintaan-darah-${request.value.permintaanDarahId}.pdf`,
    )
    feedbackStore.showFlag({
      title: 'Blanko berhasil diunduh',
      message: 'Gunakan kertas ukuran F4 (210mm × 330mm) untuk mencetak blanko ini.',
      variant: 'success',
    })
  } catch (err) {
    feedbackStore.showFlag({
      title: 'Gagal mengunduh blanko',
      message: err instanceof Error ? err.message : 'Coba lagi nanti.',
      variant: 'error',
    })
  }
}

onMounted(async () => {
  syncHeaderActions()
  await loadRequest()
  syncHeaderActions()
})
watch(requestId, async () => {
  await loadRequest()
  syncHeaderActions()
})
watch(request, syncHeaderActions)
onBeforeUnmount(clearActions)
</script>

<template>
  <section>
    <div v-if="requestsStore.isLoading && !request" :class="ui.emptyState">
      <div>
        <h2 :class="ui.emptyTitle">Memuat detail</h2>
        <p :class="ui.emptyCopy">Data permintaan sedang diambil dari backend.</p>
      </div>
    </div>

    <div v-else-if="!request" :class="ui.emptyState">
      <div>
        <h2 :class="ui.emptyTitle">Detail tidak tersedia</h2>
        <p :class="ui.emptyCopy">Permintaan tidak ditemukan atau tidak dapat diakses.</p>
      </div>
    </div>

    <template v-else>
      <section :class="[ui.card, 'mb-5 p-6']">
        <div class="flex flex-wrap items-start justify-between gap-4">
          <div>
            <h2 :class="ui.sectionTitle">{{ request.namaPasien }}</h2>
            <p class="mt-1 text-sm text-gray-500">
              {{ request.permintaanDarahId }} - {{ statusDescriptions[request.status] }}
            </p>
          </div>
          <div class="flex flex-wrap justify-end gap-2.5 max-sm:w-full max-sm:flex-col">
            <button
              type="button"
              class="flex items-center justify-center gap-2 rounded-lg px-3 py-1.5 text-xs font-medium text-blue-600 transition-colors hover:bg-blue-100 max-sm:w-full bg-blue-50"
              :disabled="requestsStore.isDownloading"
              @click="downloadBlanko"
            >
              <Download :size="16" />
              {{ requestsStore.isDownloading ? 'Mengunduh...' : 'Download Blanko' }}
            </button>
            <RouterLink
              v-if="requestsStore.canEdit(request)"
              class="flex items-center justify-center gap-2 rounded-lg px-3 py-1.5 text-xs font-medium text-blue-600 transition-colors hover:bg-blue-100 max-sm:w-full"
              :to="`/requests/${request.permintaanDarahId}/edit`"
            >
              <Pencil :size="16" />
              Edit
            </RouterLink>
            <button
              v-if="requestsStore.canConfirmPickup(request)"
              type="button"
              class="flex items-center justify-center gap-2 rounded-lg px-3 py-1.5 text-xs font-medium text-emerald-600 transition-colors hover:bg-emerald-100 max-sm:w-full"
              @click="isPickupOpen = true"
            >
              <CheckCircle2 :size="16" />
              Selesai
            </button>
            <button
              v-if="requestsStore.canCancel(request)"
              type="button"
              class="flex items-center justify-center gap-2 rounded-lg px-3 py-1.5 text-xs font-medium text-red-600 transition-colors hover:bg-red-100 max-sm:w-full bg-red-50"
              @click="isCancelOpen = true"
            >
              <CircleSlash :size="16" />
              Batalkan permintaan
            </button>
          </div>
        </div>
        <div class="mt-4">
          <StatusBadge :status="request.status" />
        </div>

        <!-- F4 Paper Size Information -->
        <div class="mt-4 p-4 bg-blue-50 border border-blue-200 rounded-lg flex items-start gap-3">
          <div class="flex-shrink-0 pt-0.5">
            <svg class="w-5 h-5 text-blue-600" fill="currentColor" viewBox="0 0 20 20">
              <path
                fill-rule="evenodd"
                d="M18 5v8a2 2 0 01-2 2h-5l-5 4v-4H4a2 2 0 01-2-2V5a2 2 0 012-2h12a2 2 0 012 2zm-11-1a1 1 0 11-2 0 1 1 0 012 0z"
                clip-rule="evenodd"
              />
            </svg>
          </div>
          <div class="flex-1">
            <h3 class="text-sm font-semibold text-blue-900">Informasi Cetak</h3>
            <p class="mt-1 text-sm text-blue-800">
              Blanko ini memiliki ukuran <strong>kertas F4 (210mm × 330mm)</strong>. Pastikan cetak
              menggunakan ukuran kertas yang sesuai.
            </p>
          </div>
        </div>
      </section>

      <section class="grid grid-cols-[280px_minmax(0,1fr)_minmax(0,1fr)] gap-5 max-xl:grid-cols-1">
        <article :class="[ui.card, 'grid content-start gap-3.5 p-6']">
          <div>
            <span class="mb-2 block text-xs font-semibold uppercase text-gray-500">Status</span>
            <StatusBadge :status="request.status" />
          </div>
          <p v-if="request.cancelReason" class="m-0 text-sm text-gray-600">
            Alasan batal: {{ request.cancelReason }}
          </p>
          <p class="m-0 text-sm text-gray-600">Dibuat {{ formatDateTime(request.createdAt) }}</p>
        </article>

        <article :class="[ui.card, 'p-6']">
          <h2 :class="ui.sectionTitle">Data Pasien</h2>
          <dl class="mt-4 grid gap-3.5">
            <div class="grid gap-1">
              <dt class="text-xs font-semibold text-gray-500">Nama Pasien</dt>
              <dd class="m-0 break-words font-semibold text-gray-900">{{ request.namaPasien }}</dd>
            </div>
            <div class="grid gap-1">
              <dt class="text-xs font-semibold text-gray-500">Nama Dokter</dt>
              <dd class="m-0 break-words font-semibold text-gray-900">
                {{ request.namaDokter || '-' }}
              </dd>
            </div>
            <div class="grid gap-1">
              <dt class="text-xs font-semibold text-gray-500">No. RM</dt>
              <dd class="m-0 break-words font-semibold text-gray-900">{{ request.noRM || '-' }}</dd>
            </div>
            <div class="grid gap-1">
              <dt class="text-xs font-semibold text-gray-500">Tempat/Tanggal Lahir</dt>
              <dd class="m-0 break-words font-semibold text-gray-900">
                {{ request.tempatLahir }}, {{ formatDate(request.tanggalLahir) }}
              </dd>
            </div>
            <div class="grid gap-1">
              <dt class="text-xs font-semibold text-gray-500">Jenis Kelamin</dt>
              <dd class="m-0 break-words font-semibold text-gray-900">
                {{ request.jenisKelamin === 'L' ? 'Laki-laki' : 'Perempuan' }}
              </dd>
            </div>
            <div class="grid gap-1">
              <dt class="text-xs font-semibold text-gray-500">Golongan Darah</dt>
              <dd class="m-0 break-words font-semibold text-gray-900">
                {{ bloodLabel(request.golonganDarah, request.rhesusDarah) }}
              </dd>
            </div>
          </dl>
        </article>

        <article :class="[ui.card, 'p-6']">
          <h2 :class="ui.sectionTitle">Data Medis</h2>
          <dl class="mt-4 grid gap-3.5">
            <div class="grid gap-1">
              <dt class="text-xs font-semibold text-gray-500">Hemoglobin</dt>
              <dd class="m-0 break-words font-semibold text-gray-900">
                {{ request.hemoglobin ? `${request.hemoglobin} g/dL` : '-' }}
              </dd>
            </div>
            <div class="grid gap-1">
              <dt class="text-xs font-semibold text-gray-500">Ruang/Bagian/Kelas</dt>
              <dd class="m-0 break-words font-semibold text-gray-900">
                {{ request.ruangBagianKelas || '-' }}
              </dd>
            </div>
            <div class="grid gap-1">
              <dt class="text-xs font-semibold text-gray-500">Pernah Transfusi</dt>
              <dd class="m-0 break-words font-semibold text-gray-900">
                {{ request.pernahTransfusi ? 'Ya' : 'Tidak' }}
              </dd>
            </div>
            <div class="grid gap-1">
              <dt class="text-xs font-semibold text-gray-500">Indikasi Transfusi</dt>
              <dd class="m-0 break-words font-semibold text-gray-900">
                {{ request.indikasiTransfusi || '-' }}
              </dd>
            </div>
            <div class="grid gap-1">
              <dt class="text-xs font-semibold text-gray-500">Riwayat Hamil</dt>
              <dd class="m-0 break-words font-semibold text-gray-900">
                {{ request.pernahHamil || '-' }}
              </dd>
            </div>
            <div class="grid gap-1">
              <dt class="text-xs font-semibold text-gray-500">Tanggal Permintaan</dt>
              <dd class="m-0 break-words font-semibold text-gray-900">
                {{ formatDateTime(request.tanggalPermintaan) }}
              </dd>
            </div>
          </dl>
        </article>
      </section>

      <section :class="[ui.card, 'mt-5 p-6']">
        <h2 :class="ui.sectionTitle">Komponen Darah</h2>
        <div
          v-if="!request.detailPermintaanDarah?.length"
          class="mt-4 rounded-2xl border border-dashed border-gray-200 p-4 text-sm text-gray-600"
        >
          Tidak ada detail komponen pada permintaan ini.
        </div>
        <div v-else :class="ui.tableWrap">
          <table :class="ui.table">
            <thead>
              <tr>
                <th :class="ui.th">Komponen</th>
                <th :class="ui.th">Darah</th>
                <th :class="ui.th">Jumlah</th>
              </tr>
            </thead>
            <tbody>
              <tr
                v-for="detail in request.detailPermintaanDarah"
                :key="detail.detailId"
                class="transition-colors hover:bg-gray-50"
              >
                <td :class="ui.td">{{ detail.komponenNama }}</td>
                <td :class="ui.td">{{ bloodLabel(detail.golonganDarah, detail.rhesusDarah) }}</td>
                <td :class="ui.td">{{ detail.jumlahKantong }} kantong</td>
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
      <div :class="ui.formField">
        <label :class="ui.formLabel" for="cancelReason">Alasan pembatalan</label>
        <textarea id="cancelReason" v-model="cancelReason" :class="ui.formTextarea" required />
      </div>
      <template #footer>
        <button
          type="button"
          :class="[btn('btnSecondary'), 'flex-1']"
          @click="isCancelOpen = false"
        >
          Batal
        </button>
        <button
          type="button"
          :class="[btn('btnDanger'), 'flex-1']"
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
        <button
          type="button"
          :class="[btn('btnSecondary'), 'flex-1']"
          @click="isPickupOpen = false"
        >
          Batal
        </button>
        <button
          type="button"
          :class="[btn('btnSuccess'), 'flex-1']"
          :disabled="requestsStore.isSubmitting"
          @click="confirmPickup"
        >
          Konfirmasi
        </button>
      </template>
    </AppModal>
  </section>
</template>
