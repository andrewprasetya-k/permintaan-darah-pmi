<script setup lang="ts">
import { computed, onBeforeUnmount, onMounted, ref, watch } from 'vue'
import { RouterLink, useRoute } from 'vue-router'
import {
  Plus,
  Search,
  AlertCircle,
  Activity,
  CircleSlash,
  CheckCircle2,
  Clock3,
  Droplets,
} from '@lucide/vue'
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
const { clearActions, setActions } = useSetPageHeaderActions()

const search = ref('')
const statusFilter = ref<PermintaanStatus | 'all'>('all')
const selectedRequest = ref<PermintaanDarahListItem | null>(null)
const cancelReason = ref('')
const isCancelOpen = ref(false)
const isPickupOpen = ref(false)
const currentPage = ref(1)
const itemsPerPage = 8
let searchDebounceTimer: ReturnType<typeof setTimeout> | null = null

const statuses: Array<{ value: PermintaanStatus | 'all'; label: string }> = [
  { value: 'all', label: 'Semua Status' },
  { value: 'dibuat', label: statusLabels.dibuat },
  { value: 'diproses', label: statusLabels.diproses },
  { value: 'bisa_diambil', label: statusLabels.bisa_diambil },
  { value: 'selesai', label: statusLabels.selesai },
  { value: 'dibatalkan', label: statusLabels.dibatalkan },
]

const statusMeta: Record<PermintaanStatus, { label: string; color: string; icon: any }> = {
  dibuat: { label: 'Baru', color: 'bg-amber-50 text-amber-700', icon: Clock3 },
  diproses: { label: 'Diproses', color: 'bg-blue-50 text-blue-700', icon: Activity },
  bisa_diambil: { label: 'Siap Diambil', color: 'bg-violet-50 text-violet-700', icon: Droplets },
  selesai: { label: 'Selesai', color: 'bg-emerald-50 text-emerald-700', icon: CheckCircle2 },
  dibatalkan: { label: 'Dibatalkan', color: 'bg-red-50 text-red-700', icon: CircleSlash },
}

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

const totalPages = computed(() =>
  Math.max(1, Math.ceil(filteredRequests.value.length / itemsPerPage)),
)

const paginatedRequests = computed(() => {
  const startIndex = (currentPage.value - 1) * itemsPerPage
  const endIndex = startIndex + itemsPerPage
  return filteredRequests.value.slice(startIndex, endIndex)
})

const pageRange = computed(() => {
  const startIndex = (currentPage.value - 1) * itemsPerPage
  const endIndex = Math.min(startIndex + itemsPerPage, filteredRequests.value.length)
  return { startIndex, endIndex }
})

const handleSearchInput = () => {
  currentPage.value = 1
  if (searchDebounceTimer) {
    clearTimeout(searchDebounceTimer)
  }
  searchDebounceTimer = setTimeout(() => {
    // Search filter is applied in computed property
  }, 300)
}

const handleStatusChange = () => {
  currentPage.value = 1
}

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

onBeforeUnmount(clearActions)
</script>

<template>
  <div class="flex h-full min-h-0 flex-col">
    <div
      class="flex min-h-0 flex-1 flex-col rounded-2xl border border-gray-100 bg-white overflow-hidden"
    >
      <!-- Header Section -->
      <div class="border-b border-gray-100 px-5 py-4">
        <div class="flex flex-col gap-4">
          <div class="flex flex-col gap-3 lg:flex-row lg:items-center lg:justify-between">
            <div class="flex flex-1 flex-col gap-3 sm:flex-row sm:items-center">
              <div class="relative w-full sm:max-w-md">
                <Search
                  :size="16"
                  class="pointer-events-none absolute left-3 top-1/2 -translate-y-1/2 text-gray-400"
                />
                <input
                  v-model="search"
                  type="search"
                  placeholder="Cari nama pasien atau ID permintaan"
                  class="w-full rounded-xl border border-gray-200 bg-white py-2.5 pl-9 pr-3 text-sm text-gray-700 outline-none transition-all focus:border-blue-400 focus:ring-2 focus:ring-blue-100"
                  @input="handleSearchInput"
                />
              </div>
              <select
                v-model="statusFilter"
                class="w-full rounded-xl border border-gray-200 bg-white px-3.5 py-2.5 text-sm text-gray-700 outline-none transition-all focus:border-blue-400 focus:ring-2 focus:ring-blue-100 sm:w-[190px]"
                @change="handleStatusChange"
              >
                <option v-for="status in statuses" :key="status.value" :value="status.value">
                  {{ status.label }}
                </option>
              </select>
            </div>
            <RouterLink
              to="/requests/new"
              class="inline-flex items-center justify-center gap-2 rounded-xl bg-blue-600 px-4 py-2.5 text-sm font-medium text-white transition-colors hover:bg-blue-700 lg:px-5"
            >
              <Plus :size="16" />
              Buat Permintaan
            </RouterLink>
          </div>

          <div class="border-t border-gray-100 pt-4">
            <p class="text-sm text-gray-600">
              Menampilkan
              <span class="font-semibold text-gray-900">{{ paginatedRequests.length }}</span>
              data
              <span v-if="search || statusFilter !== 'all'">
                dari
                <span class="font-semibold text-gray-900">{{ filteredRequests.length }}</span>
                total permintaan darah
              </span>
            </p>
          </div>
        </div>
      </div>

      <!-- Loading State -->
      <div
        v-if="requestsStore.isLoading"
        class="flex flex-1 items-center justify-center px-6 py-16 text-sm text-gray-400"
      >
        <span
          class="mr-3 h-5 w-5 animate-spin rounded-full border-2 border-gray-200 border-t-blue-500"
        />
        Memuat data permintaan darah...
      </div>

      <!-- Empty State -->
      <div
        v-else-if="filteredRequests.length === 0"
        class="flex flex-1 items-center justify-center"
      >
        <div class="text-center">
          <h2 class="text-base font-semibold text-gray-900">Tidak ada data</h2>
          <p class="mt-1 text-sm text-gray-500">Ubah filter atau buat permintaan baru.</p>
        </div>
      </div>

      <!-- Table View -->
      <div v-else class="flex min-h-0 flex-1 flex-col">
        <!-- Desktop Table -->
        <div class="min-h-0 flex-1 overflow-auto max-md:hidden">
          <table class="w-full min-w-[1000px] text-sm">
            <thead class="sticky top-0 z-10 bg-white shadow-sm">
              <tr class="border-b border-gray-100">
                <th
                  class="bg-white px-5 py-3.5 text-left text-xs font-semibold uppercase tracking-wide text-gray-400"
                >
                  Nama Pasien
                </th>
                <th
                  class="bg-white px-5 py-3.5 text-left text-xs font-semibold uppercase tracking-wide text-gray-400"
                >
                  ID Permintaan
                </th>
                <th
                  class="bg-white px-5 py-3.5 text-left text-xs font-semibold uppercase tracking-wide text-gray-400"
                >
                  Golongan Darah
                </th>
                <th
                  class="bg-white px-5 py-3.5 text-left text-xs font-semibold uppercase tracking-wide text-gray-400"
                >
                  Status
                </th>
                <th
                  class="bg-white px-5 py-3.5 text-left text-xs font-semibold uppercase tracking-wide text-gray-400"
                >
                  Tanggal Permintaan
                </th>
                <th
                  class="bg-white px-5 py-3.5 text-center text-xs font-semibold uppercase tracking-wide text-gray-400"
                >
                  Aksi
                </th>
              </tr>
            </thead>
            <tbody class="divide-y divide-gray-50">
              <tr
                v-for="request in paginatedRequests"
                :key="request.permintaanDarahId"
                class="transition-colors hover:bg-gray-50"
              >
                <td class="px-5 py-4 font-medium text-gray-800">{{ request.namaPasien }}</td>
                <td class="px-5 py-4 text-xs font-mono text-gray-600">
                  {{ request.permintaanDarahId }}
                </td>
                <td class="px-5 py-4 text-gray-500">
                  {{ bloodLabel(request.golonganDarah, request.rhesusDarah) }}
                </td>
                <td class="px-5 py-4">
                  <span
                    class="inline-flex items-center gap-1 rounded-lg px-2.5 py-1 text-xs font-medium"
                    :class="statusMeta[request.status]?.color || 'bg-gray-100 text-gray-700'"
                  >
                    <component :is="statusMeta[request.status]?.icon" :size="12" />
                    {{ statusMeta[request.status]?.label || request.status }}
                  </span>
                </td>
                <td class="px-5 py-4 text-gray-500">{{ formatDate(request.tanggalPermintaan) }}</td>
                <td class="px-5 py-4">
                  <div class="flex items-center justify-center gap-2">
                    <RouterLink
                      :to="`/requests/${request.permintaanDarahId}`"
                      class="rounded-lg px-3 py-1.5 text-xs font-medium text-green-600 transition-colors hover:bg-green-100"
                    >
                      Detail
                    </RouterLink>
                    <RouterLink
                      v-if="requestsStore.canEdit(request)"
                      :to="`/requests/${request.permintaanDarahId}/edit`"
                      class="rounded-lg px-3 py-1.5 text-xs font-medium text-blue-600 transition-colors hover:bg-blue-100"
                    >
                      Edit
                    </RouterLink>
                    <button
                      v-if="requestsStore.canConfirmPickup(request)"
                      type="button"
                      class="rounded-lg px-3 py-1.5 text-xs font-medium text-emerald-600 transition-colors hover:bg-emerald-100"
                      @click="openPickup(request)"
                    >
                      Selesai
                    </button>
                    <button
                      v-if="requestsStore.canCancel(request)"
                      type="button"
                      class="rounded-lg px-3 py-1.5 text-xs font-medium text-red-600 transition-colors hover:bg-red-100"
                      @click="openCancel(request)"
                    >
                      Batalkan
                    </button>
                  </div>
                </td>
              </tr>
            </tbody>
          </table>
        </div>

        <!-- Mobile Card View -->
        <div class="hidden p-5 max-md:block">
          <article
            v-for="request in paginatedRequests"
            :key="request.permintaanDarahId"
            class="rounded-2xl border border-gray-100 bg-white p-4 [&+&]:mt-3"
          >
            <div class="flex items-start justify-between gap-3">
              <div>
                <h2 class="m-0 text-base font-semibold text-gray-900">{{ request.namaPasien }}</h2>
                <p class="mt-1 text-xs text-gray-500 font-mono">{{ request.permintaanDarahId }}</p>
              </div>
              <div
                class="inline-flex items-center gap-1 rounded-lg px-2.5 py-1 text-xs font-medium whitespace-nowrap"
                :class="statusMeta[request.status]?.color || 'bg-gray-100 text-gray-700'"
              >
                <component :is="statusMeta[request.status]?.icon" :size="12" />
                {{ statusMeta[request.status]?.label || request.status }}
              </div>
            </div>
            <dl class="my-3.5 grid gap-2.5">
              <div>
                <dt class="text-xs font-semibold text-gray-500">Golongan Darah</dt>
                <dd class="m-0 break-words font-semibold text-gray-900">
                  {{ bloodLabel(request.golonganDarah, request.rhesusDarah) }}
                </dd>
              </div>
              <div>
                <dt class="text-xs font-semibold text-gray-500">Tanggal Permintaan</dt>
                <dd class="m-0 break-words text-gray-600">
                  {{ formatDate(request.tanggalPermintaan) }}
                </dd>
              </div>
            </dl>
            <div class="flex flex-wrap gap-2">
              <RouterLink
                :to="`/requests/${request.permintaanDarahId}`"
                class="flex-1 rounded-lg px-3 py-1.5 text-center text-xs font-medium text-green-600 transition-colors hover:bg-green-100"
              >
                Detail
              </RouterLink>
              <RouterLink
                v-if="requestsStore.canEdit(request)"
                :to="`/requests/${request.permintaanDarahId}/edit`"
                class="flex-1 rounded-lg px-3 py-1.5 text-center text-xs font-medium text-blue-600 transition-colors hover:bg-blue-100"
              >
                Edit
              </RouterLink>
              <button
                v-if="requestsStore.canConfirmPickup(request)"
                type="button"
                class="flex-1 rounded-lg px-3 py-1.5 text-center text-xs font-medium text-emerald-600 transition-colors hover:bg-emerald-100"
                @click="openPickup(request)"
              >
                Selesai
              </button>
              <button
                v-if="requestsStore.canCancel(request)"
                type="button"
                class="flex-1 rounded-lg px-3 py-1.5 text-center text-xs font-medium text-red-600 transition-colors hover:bg-red-100"
                @click="openCancel(request)"
              >
                Batal
              </button>
            </div>
          </article>
        </div>

        <!-- Pagination -->
        <div
          v-if="filteredRequests.length > 0"
          class="flex items-center justify-between border-t border-gray-100 bg-gray-50 px-6 py-4"
        >
          <div class="text-sm text-gray-700">
            Menampilkan {{ pageRange.startIndex + 1 }} -
            {{ pageRange.endIndex }}
            dari {{ filteredRequests.length }} data
          </div>
          <div v-if="totalPages > 1" class="flex items-center space-x-2">
            <button
              @click="currentPage = Math.max(1, currentPage - 1)"
              :disabled="currentPage === 1"
              class="rounded-lg border border-gray-200 px-3 py-1 text-sm transition-colors hover:bg-gray-100 disabled:cursor-not-allowed disabled:opacity-50"
            >
              Sebelumnya
            </button>
            <button
              v-for="page in totalPages"
              :key="page"
              @click="currentPage = page"
              class="rounded-lg px-3 py-1 text-sm transition-colors"
              :class="
                currentPage === page
                  ? 'bg-blue-600 text-white'
                  : 'border border-gray-200 hover:bg-gray-100'
              "
            >
              {{ page }}
            </button>
            <button
              @click="currentPage = Math.min(totalPages, currentPage + 1)"
              :disabled="currentPage === totalPages"
              class="rounded-lg border border-gray-200 px-3 py-1 text-sm transition-colors hover:bg-gray-100 disabled:cursor-not-allowed disabled:opacity-50"
            >
              Selanjutnya
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Cancel Modal -->
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

    <!-- Pickup Confirmation Modal -->
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
  </div>
</template>
