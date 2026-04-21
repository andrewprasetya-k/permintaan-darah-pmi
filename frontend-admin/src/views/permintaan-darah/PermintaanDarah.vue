<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { usePermintaanStore } from '@/stores/permintaan'
import {
  Plus,
  AlertCircle,
  Search,
  Activity,
  CircleSlash,
  CheckCircle2,
  Clock3,
  Trash2,
  Droplets,
} from '@lucide/vue'
import PermintaanCreateDrawer from './PermintaanCreateDrawer.vue'
import PermintaanEditDrawer from './PermintaanEditDrawer.vue'
import PermintaanDetailDrawer from './PermintaanDetailDrawer.vue'
import type { PermintaanDarah } from '@/types/models'

const permintaanStore = usePermintaanStore()
const showCreateDrawer = ref(false)
const showEditDrawer = ref(false)
const showDetailDrawer = ref(false)
const searchTerm = ref('')
const statusFilter = ref<'all' | PermintaanDarah['status']>('all')
const currentPage = ref(1)
const itemsPerPage = 8
const deleteTarget = ref<PermintaanDarah | null>(null)
const deleting = ref(false)

onMounted(async () => await permintaanStore.fetchAll())

const formatDate = (date: string) =>
  new Date(date).toLocaleDateString('id-ID', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
  })

const formatGender = (gender: PermintaanDarah['jenisKelamin']) =>
  gender === 'L' ? 'Laki-laki' : 'Perempuan'

const formatBloodType = (request: PermintaanDarah) => {
  const type = request.golonganDarah ?? '-'
  const rhesus = request.rhesusDarah ?? ''
  return `${type}${rhesus}`
}

const statusMeta: Record<
  PermintaanDarah['status'],
  { label: string; color: string; icon: typeof Clock3 }
> = {
  dibuat: { label: 'Baru', color: 'bg-amber-50 text-amber-700', icon: Clock3 },
  diproses: { label: 'Diproses', color: 'bg-blue-50 text-blue-700', icon: Activity },
  bisa_diambil: { label: 'Siap Diambil', color: 'bg-violet-50 text-violet-700', icon: Droplets },
  selesai: { label: 'Selesai', color: 'bg-emerald-50 text-emerald-700', icon: CheckCircle2 },
  dibatalkan: { label: 'Dibatalkan', color: 'bg-red-50 text-red-700', icon: CircleSlash },
}

const filteredRequests = computed(() => {
  const keyword = searchTerm.value.trim().toLowerCase()

  return permintaanStore.requests.filter((item) => {
    const matchesStatus = statusFilter.value === 'all' || item.status === statusFilter.value
    const matchesSearch =
      keyword === '' ||
      item.namaPasien.toLowerCase().includes(keyword) ||
      item.permintaanDarahId.toLowerCase().includes(keyword) ||
      formatBloodType(item).toLowerCase().includes(keyword)

    return matchesStatus && matchesSearch
  })
})

const totalPages = computed(() =>
  Math.max(1, Math.ceil(filteredRequests.value.length / itemsPerPage)),
)

const paginatedRequests = computed(() => {
  const startIndex = (currentPage.value - 1) * itemsPerPage
  return filteredRequests.value.slice(startIndex, startIndex + itemsPerPage)
})

const pageRange = computed(() => {
  const startIndex = (currentPage.value - 1) * itemsPerPage
  const endIndex = Math.min(startIndex + itemsPerPage, filteredRequests.value.length)
  return { startIndex, endIndex }
})

const resetPage = () => {
  currentPage.value = 1
}

const openCreateDrawer = () => {
  showCreateDrawer.value = true
}

const openEditDrawer = (request: PermintaanDarah) => {
  permintaanStore.selectedRequest = request
  showEditDrawer.value = true
}

const openDetailDrawer = async (request: PermintaanDarah) => {
  await permintaanStore.fetchById(request.permintaanDarahId)
  showDetailDrawer.value = true
}

const askDelete = (request: PermintaanDarah) => {
  deleteTarget.value = request
}

const cancelDelete = () => {
  if (deleting.value) return
  deleteTarget.value = null
}

const confirmDelete = async () => {
  if (!deleteTarget.value) return

  deleting.value = true
  try {
    await permintaanStore.deleteRequest(deleteTarget.value.permintaanDarahId)
    deleteTarget.value = null
    if (currentPage.value > totalPages.value) {
      currentPage.value = totalPages.value
    }
  } finally {
    deleting.value = false
  }
}

const handleSubmit = async () => {
  await permintaanStore.fetchAll()
  resetPage()
}
</script>

<template>
  <div class="flex h-full min-h-0 flex-col">
    <PermintaanCreateDrawer
      :is-open="showCreateDrawer"
      @close="showCreateDrawer = false"
      @submit="handleSubmit"
    />
    <PermintaanEditDrawer
      :is-open="showEditDrawer"
      @close="showEditDrawer = false"
      @submit="handleSubmit"
    />
    <PermintaanDetailDrawer :is-open="showDetailDrawer" @close="showDetailDrawer = false" />

    <Teleport to="body">
      <Transition name="backdrop">
        <div
          v-if="deleteTarget"
          class="fixed inset-0 z-40 bg-black/30 backdrop-blur-sm"
          @click="cancelDelete"
        />
      </Transition>
      <Transition name="drawer">
        <div
          v-if="deleteTarget"
          class="fixed left-1/2 top-1/2 z-50 w-full max-w-md -translate-x-1/2 -translate-y-1/2 rounded-3xl bg-white p-6 shadow-2xl"
        >
          <div
            class="mb-4 flex h-12 w-12 items-center justify-center rounded-2xl bg-red-50 text-red-600"
          >
            <Trash2 :size="20" />
          </div>
          <h3 class="text-lg font-semibold text-gray-900">Hapus Permintaan Darah</h3>
          <p class="mt-2 text-sm leading-6 text-gray-500">
            Apakah Anda yakin ingin menghapus permintaan untuk
            <span class="font-semibold text-gray-900">{{ deleteTarget.namaPasien }}</span
            >? Data yang dihapus tidak dapat dikembalikan dari halaman ini.
          </p>
          <div class="mt-6 flex gap-3">
            <button
              type="button"
              @click="cancelDelete"
              class="flex-1 rounded-xl bg-gray-100 px-4 py-2.5 text-sm font-medium text-gray-700 transition-colors hover:bg-gray-200"
            >
              Batal
            </button>
            <button
              type="button"
              :disabled="deleting"
              @click="confirmDelete"
              class="flex-1 rounded-xl bg-red-600 px-4 py-2.5 text-sm font-medium text-white transition-colors hover:bg-red-700 disabled:opacity-60"
            >
              {{ deleting ? 'Menghapus...' : 'Hapus' }}
            </button>
          </div>
        </div>
      </Transition>
    </Teleport>

    <div class="mb-6 flex items-center justify-between">
      <button
        @click="openCreateDrawer"
        class="flex items-center gap-2 rounded-xl bg-blue-600 px-4 py-2 text-sm font-medium text-white transition-colors hover:bg-blue-700"
      >
        <Plus :size="16" />
        Tambah Permintaan
      </button>
    </div>

    <div
      class="flex min-h-0 flex-1 flex-col rounded-2xl border border-gray-100 bg-white overflow-hidden"
    >
      <div class="border-b border-gray-100 px-5 py-4">
        <div class="flex flex-col gap-4 lg:flex-row lg:items-center lg:justify-between">
          <div>
            <h2 class="text-lg font-semibold text-gray-900">Data Permintaan Darah</h2>
            <p class="mt-1 text-sm text-gray-500">
              Kelola data permintaan darah yang tersimpan di sistem.
            </p>
          </div>
          <div class="flex flex-col gap-3 sm:flex-row">
            <div class="relative min-w-[260px]">
              <Search
                :size="16"
                class="pointer-events-none absolute left-3 top-1/2 -translate-y-1/2 text-gray-400"
              />
              <input
                v-model="searchTerm"
                type="text"
                placeholder="Cari ID, nama pasien, golongan darah"
                class="w-full rounded-xl border border-gray-200 bg-white py-2.5 pl-9 pr-3 text-sm text-gray-700 outline-none transition-all focus:border-blue-400 focus:ring-2 focus:ring-blue-100"
                @input="resetPage"
              />
            </div>
            <select
              v-model="statusFilter"
              class="rounded-xl border border-gray-200 bg-white px-3.5 py-2.5 text-sm text-gray-700 outline-none transition-all focus:border-blue-400 focus:ring-2 focus:ring-blue-100"
              @change="resetPage"
            >
              <option value="all">Semua Status</option>
              <option value="dibuat">Baru</option>
              <option value="diproses">Diproses</option>
              <option value="bisa_diambil">Siap Diambil</option>
              <option value="selesai">Selesai</option>
              <option value="dibatalkan">Dibatalkan</option>
            </select>
          </div>
        </div>
        <p class="text-sm text-gray-600">
          <span class="mt-4 block">
            Menampilkan
            <span class="font-semibold text-gray-900">{{ filteredRequests.length }}</span>
            data
            <span v-if="searchTerm || statusFilter !== 'all'">
              dari
              <span class="font-semibold text-gray-900">{{ permintaanStore.requests.length }}</span>
              total permintaan darah
            </span>
          </span>
        </p>
      </div>

      <div
        v-if="permintaanStore.isLoading"
        class="flex flex-1 items-center justify-center px-6 py-16 text-sm text-gray-400"
      >
        <span
          class="mr-3 h-5 w-5 animate-spin rounded-full border-2 border-gray-200 border-t-blue-500"
        />
        Memuat data permintaan darah...
      </div>

      <div
        v-else-if="permintaanStore.error"
        class="mx-6 mb-6 flex items-center gap-2 rounded-2xl border border-red-100 bg-red-50 p-4 text-sm text-red-600"
      >
        <AlertCircle :size="16" />
        {{ permintaanStore.error }}
      </div>

      <div v-else class="flex min-h-0 flex-1 flex-col">
        <div class="min-h-0 flex-1 overflow-auto">
          <table class="w-full min-w-[980px] text-sm">
            <thead class="sticky top-0 z-10 bg-white shadow-sm">
              <tr class="border-b border-gray-100">
                <th
                  class="bg-white px-5 py-3.5 text-left text-xs font-semibold uppercase tracking-wide text-gray-400"
                >
                  ID Permintaan
                </th>
                <th
                  class="bg-white px-5 py-3.5 text-left text-xs font-semibold uppercase tracking-wide text-gray-400"
                >
                  Nama Pasien
                </th>
                <th
                  class="bg-white px-5 py-3.5 text-left text-xs font-semibold uppercase tracking-wide text-gray-400"
                >
                  Jenis Kelamin
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
              <tr v-if="paginatedRequests.length === 0">
                <td colspan="7" class="px-6 py-12 text-center text-gray-500">
                  Tidak ada data permintaan darah yang ditemukan
                </td>
              </tr>
              <tr
                v-for="item in paginatedRequests"
                :key="item.permintaanDarahId"
                class="transition-colors hover:bg-gray-50"
              >
                <td class="px-5 py-4 font-medium text-gray-800">
                  {{ item.permintaanDarahId }}
                </td>
                <td class="px-5 py-4 text-gray-800">{{ item.namaPasien }}</td>
                <td class="px-5 py-4 text-gray-500">{{ formatGender(item.jenisKelamin) }}</td>
                <td class="px-5 py-4 text-gray-500">{{ formatBloodType(item) }}</td>
                <td class="px-5 py-4">
                  <span
                    class="inline-flex items-center gap-1 rounded-lg px-2.5 py-1 text-xs font-medium"
                    :class="statusMeta[item.status]?.color || 'bg-gray-100 text-gray-700'"
                  >
                    <component :is="statusMeta[item.status]?.icon" :size="12" />
                    {{ statusMeta[item.status]?.label || item.status }}
                  </span>
                </td>
                <td class="px-5 py-4 text-gray-500">
                  {{ formatDate(item.tanggalPermintaan) }}
                </td>
                <td class="px-5 py-4">
                  <div class="flex items-center justify-center gap-2">
                    <button
                      @click="openDetailDrawer(item)"
                      class="rounded-lg px-3 py-1.5 text-xs font-medium text-green-600 transition-colors hover:bg-green-100"
                    >
                      Detail
                    </button>
                    <button
                      @click="openEditDrawer(item)"
                      class="rounded-lg px-3 py-1.5 text-xs font-medium text-blue-600 transition-colors hover:bg-blue-100"
                    >
                      Edit
                    </button>
                    <button
                      @click="askDelete(item)"
                      class="rounded-lg px-3 py-1.5 text-xs font-medium text-red-600 transition-colors hover:bg-red-100"
                    >
                      Hapus
                    </button>
                  </div>
                </td>
              </tr>
            </tbody>
          </table>
        </div>

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
  </div>
</template>

<style scoped>
.backdrop-enter-active,
.backdrop-leave-active {
  transition: opacity 0.2s ease;
}

.backdrop-enter-from,
.backdrop-leave-to {
  opacity: 0;
}

.drawer-enter-active,
.drawer-leave-active {
  transition:
    opacity 0.2s ease,
    transform 0.2s ease;
}

.drawer-enter-from,
.drawer-leave-to {
  opacity: 0;
  transform: translate(-50%, -48%);
}
</style>
