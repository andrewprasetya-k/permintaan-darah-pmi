<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { RouterLink, useRoute } from 'vue-router'
import { Plus, Search } from '@lucide/vue'
import AppModal from '@/components/feedback/AppModal.vue'
import StatusBadge from '@/components/rs/StatusBadge.vue'
import { useFeedbackStore } from '@/stores/feedback'
import { useMyRequestsStore } from '@/stores/my-requests'
import { useSetPageHeaderActions } from '@/composables/usePageHeader'
import type { PermintaanDarahListItem, PermintaanStatus } from '@/types/models'
import { bloodLabel, formatDate, statusLabels } from '@/utils/format'
import { btn, ui } from '@/utils/ui'

const route = useRoute()
const requestsStore = useMyRequestsStore()
const feedbackStore = useFeedbackStore()
const { setActions } = useSetPageHeaderActions()

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
  if (!selectedRequest.value || !cancelReason.value.trim()) return

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
  if (!selectedRequest.value) return

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

  setActions([
    {
      label: 'Buat Permintaan',
      to: '/requests/new',
      icon: Plus,
      variant: 'primary',
    },
  ])

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
    <div :class="[ui.card, 'overflow-hidden']">
      <div
        class="grid grid-cols-[minmax(0,1fr)_220px] gap-3.5 border-b border-gray-100 p-5 max-md:grid-cols-1"
      >
        <div class="relative">
          <label class="sr-only" for="search">Cari</label>
          <Search
            :size="16"
            class="pointer-events-none absolute left-3 top-1/2 z-10 -translate-y-1/2 text-gray-400"
          />
          <input
            id="search"
            v-model="search"
            :class="[ui.formControl, 'pl-9']"
            type="search"
            placeholder="Nama pasien atau ID permintaan"
          />
        </div>
        <div :class="ui.formField">
          <label :class="ui.formLabel" for="status">Status</label>
          <select id="status" v-model="statusFilter" :class="ui.formControl">
            <option v-for="status in statuses" :key="status.value" :value="status.value">
              {{ status.label }}
            </option>
          </select>
        </div>
      </div>

      <div v-if="requestsStore.isLoading" :class="[ui.emptyState, 'm-5']">
        <div>
          <h2 :class="ui.emptyTitle">Memuat permintaan</h2>
          <p :class="ui.emptyCopy">Data sedang diambil dari backend.</p>
        </div>
      </div>

      <div v-else-if="filteredRequests.length === 0" :class="[ui.emptyState, 'm-5']">
        <div>
          <h2 :class="ui.emptyTitle">Tidak ada data</h2>
          <p :class="ui.emptyCopy">Ubah filter atau buat permintaan baru.</p>
        </div>
      </div>

      <template v-else>
        <div :class="[ui.tableWrap, 'max-md:hidden']">
          <table :class="ui.table">
            <thead>
              <tr>
                <th :class="ui.th">Pasien</th>
                <th :class="ui.th">Darah</th>
                <th :class="ui.th">Tanggal</th>
                <th :class="ui.th">Status</th>
                <th :class="ui.th">Aksi</th>
              </tr>
            </thead>
            <tbody>
              <tr
                v-for="request in filteredRequests"
                :key="request.permintaanDarahId"
                class="transition-colors hover:bg-gray-50"
              >
                <td :class="ui.td">
                  <strong>{{ request.namaPasien }}</strong>
                  <span class="mt-1 block text-xs text-gray-500">{{
                    request.permintaanDarahId
                  }}</span>
                </td>
                <td :class="ui.td">{{ bloodLabel(request.golonganDarah, request.rhesusDarah) }}</td>
                <td :class="ui.td">{{ formatDate(request.tanggalPermintaan) }}</td>
                <td :class="ui.td"><StatusBadge :status="request.status" /></td>
                <td :class="ui.td">
                  <div class="flex flex-wrap gap-2">
                    <RouterLink
                      :class="btn('btnSecondary')"
                      :to="`/requests/${request.permintaanDarahId}`"
                    >
                      Detail
                    </RouterLink>
                    <RouterLink
                      v-if="requestsStore.canEdit(request)"
                      :class="btn('btnSecondary')"
                      :to="`/requests/${request.permintaanDarahId}/edit`"
                    >
                      Edit
                    </RouterLink>
                    <button
                      v-if="requestsStore.canConfirmPickup(request)"
                      type="button"
                      :class="btn('btnSuccess')"
                      @click="openPickup(request)"
                    >
                      Selesai
                    </button>
                    <button
                      v-if="requestsStore.canCancel(request)"
                      type="button"
                      :class="btn('btnDanger')"
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

        <div class="hidden p-5 max-md:block">
          <article
            v-for="request in filteredRequests"
            :key="request.permintaanDarahId"
            class="rounded-2xl border border-gray-100 bg-white p-4 [&+&]:mt-3"
          >
            <div class="flex items-start justify-between gap-3">
              <div>
                <h2 class="m-0 text-base font-semibold text-gray-900">{{ request.namaPasien }}</h2>
                <p class="mt-1 text-sm text-gray-500">
                  {{ formatDate(request.tanggalPermintaan) }}
                </p>
              </div>
              <StatusBadge :status="request.status" />
            </div>
            <dl class="my-3.5 grid gap-2.5">
              <div>
                <dt class="text-xs font-semibold text-gray-500">Darah</dt>
                <dd class="m-0 break-words font-semibold text-gray-900">
                  {{ bloodLabel(request.golonganDarah, request.rhesusDarah) }}
                </dd>
              </div>
              <div>
                <dt class="text-xs font-semibold text-gray-500">ID</dt>
                <dd class="m-0 break-words font-semibold text-gray-900">
                  {{ request.permintaanDarahId }}
                </dd>
              </div>
            </dl>
            <div class="flex flex-wrap gap-2">
              <RouterLink
                :class="btn('btnSecondary')"
                :to="`/requests/${request.permintaanDarahId}`"
              >
                Detail
              </RouterLink>
              <RouterLink
                v-if="requestsStore.canEdit(request)"
                :class="btn('btnSecondary')"
                :to="`/requests/${request.permintaanDarahId}/edit`"
              >
                Edit
              </RouterLink>
              <button
                v-if="requestsStore.canConfirmPickup(request)"
                type="button"
                :class="btn('btnSuccess')"
                @click="openPickup(request)"
              >
                Selesai
              </button>
              <button
                v-if="requestsStore.canCancel(request)"
                type="button"
                :class="btn('btnDanger')"
                @click="openCancel(request)"
              >
                Batal
              </button>
            </div>
          </article>
        </div>
      </template>
    </div>

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
