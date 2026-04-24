<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue'
import { useRumahSakitStore } from '@/stores/rumah-sakit'
import { Plus, AlertCircle, Hospital, ChevronLeft, ChevronRight, RotateCcw } from '@lucide/vue'
import RumahSakitCreateDrawer from './RumahSakitCreateDrawer.vue'
import RumahSakitEditDrawer from './RumahSakitEditDrawer.vue'
import RumahSakitDetailDrawer from './RumahSakitDetailDrawer.vue'

const rumahSakitStore = useRumahSakitStore()
const showCreateDrawer = ref(false)
const showEditDrawer = ref(false)
const showDetailDrawer = ref(false)
const showRestoreModal = ref(false)
const restoreId = ref('')
const restoring = ref(false)
const currentPage = ref(1)
const itemsPerPage = 10

const loadHospitals = async (page = currentPage.value) => {
  const offset = (page - 1) * itemsPerPage
  await rumahSakitStore.fetchAll({ limit: itemsPerPage, offset })
}

onMounted(async () => await loadHospitals())

const openCreateDrawer = () => {
  showCreateDrawer.value = true
}

const openEditDrawer = (rumahSakit: any) => {
  rumahSakitStore.selectedRumahSakit = rumahSakit
  showEditDrawer.value = true
}

const openDetailDrawer = (rumahSakit: any) => {
  rumahSakitStore.selectedRumahSakit = rumahSakit
  showDetailDrawer.value = true
}

const deleteHospital = async (id: string) => {
  if (confirm('Yakin ingin menghapus rumah sakit ini?')) {
    await rumahSakitStore.deleteHospital(id)
    await loadHospitals(currentPage.value)
  }
}

const handleSubmit = async () => {
  await loadHospitals(currentPage.value)
}

const openRestoreModal = () => {
  restoreId.value = ''
  showRestoreModal.value = true
}

const submitRestore = async () => {
  if (!restoreId.value.trim()) return
  restoring.value = true
  try {
    await rumahSakitStore.restore(restoreId.value.trim())
    showRestoreModal.value = false
    await loadHospitals(currentPage.value)
  } finally {
    restoring.value = false
  }
}

const totalPages = computed(() =>
  Math.max(1, Math.ceil((rumahSakitStore.pagination?.total ?? 0) / itemsPerPage)),
)

const pageRange = computed(() => {
  const startIndex = (currentPage.value - 1) * itemsPerPage
  const total = rumahSakitStore.pagination?.total ?? rumahSakitStore.hospitals.length
  const endIndex = Math.min(startIndex + rumahSakitStore.hospitals.length, total)
  return { startIndex, endIndex }
})

watch(currentPage, async (page, previousPage) => {
  if (page === previousPage) return
  await loadHospitals(page)
})
</script>

<template>
  <div class="flex h-full min-h-0 flex-col">
    <Teleport to="body">
      <Transition name="backdrop">
        <div
          v-if="showRestoreModal"
          class="fixed inset-0 z-40 bg-black/30 backdrop-blur-sm"
          @click="showRestoreModal = false"
        />
      </Transition>
      <Transition name="drawer">
        <div
          v-if="showRestoreModal"
          class="fixed left-1/2 top-1/2 z-50 w-full max-w-md -translate-x-1/2 -translate-y-1/2 rounded-3xl bg-white p-6 shadow-2xl"
        >
          <div class="mb-4 flex h-12 w-12 items-center justify-center rounded-2xl bg-emerald-50 text-emerald-600">
            <RotateCcw :size="20" />
          </div>
          <h3 class="text-lg font-semibold text-gray-900">Pulihkan Rumah Sakit</h3>
          <p class="mt-2 text-sm leading-6 text-gray-500">
            Endpoint backend mendukung restore by ID. Masukkan `rumahSakitId` yang ingin dipulihkan.
          </p>
          <input
            v-model="restoreId"
            type="text"
            placeholder="Contoh: RS001"
            class="mt-4 w-full rounded-xl border border-gray-200 bg-gray-50 px-3.5 py-2.5 text-sm text-gray-900 outline-none transition-all focus:bg-white focus:border-blue-400 focus:ring-2 focus:ring-blue-100"
          />
          <div
            v-if="rumahSakitStore.error"
            class="mt-3 rounded-xl border border-red-100 bg-red-50 px-3 py-2 text-xs text-red-600"
          >
            {{ rumahSakitStore.error }}
          </div>
          <div class="mt-6 flex gap-3">
            <button
              type="button"
              class="flex-1 rounded-xl bg-gray-100 px-4 py-2.5 text-sm font-medium text-gray-700 transition-colors hover:bg-gray-200"
              @click="showRestoreModal = false"
            >
              Batal
            </button>
            <button
              type="button"
              :disabled="restoring || !restoreId.trim()"
              class="flex-1 rounded-xl bg-emerald-600 px-4 py-2.5 text-sm font-medium text-white transition-colors hover:bg-emerald-700 disabled:opacity-50"
              @click="submitRestore"
            >
              {{ restoring ? 'Memulihkan...' : 'Pulihkan' }}
            </button>
          </div>
        </div>
      </Transition>
    </Teleport>

    <!-- Drawers -->
    <RumahSakitCreateDrawer
      :is-open="showCreateDrawer"
      @close="showCreateDrawer = false"
      @submit="handleSubmit"
    />
    <RumahSakitEditDrawer
      :is-open="showEditDrawer"
      @close="showEditDrawer = false"
      @submit="handleSubmit"
    />
    <RumahSakitDetailDrawer :is-open="showDetailDrawer" @close="showDetailDrawer = false" />

    <!-- Loading -->
    <!-- <div
      v-if="rumahSakitStore.isLoading"
      class="flex items-center justify-center py-16 text-sm text-gray-400"
    >
      <span
        class="w-5 h-5 border-2 border-gray-200 border-t-blue-500 rounded-full animate-spin mr-3"
      />
      Memuat data...
    </div> -->

    <!-- Error -->
    <div
      v-if="rumahSakitStore.error"
      class="flex items-center gap-2 p-4 bg-red-50 border border-red-100 text-red-600 rounded-xl text-sm"
    >
      <AlertCircle :size="16" />
      {{ rumahSakitStore.error }}
    </div>

    <!-- Table -->
    <div class="flex min-h-0 flex-1 flex-col rounded-2xl border border-gray-100 bg-white overflow-hidden">
      <div class="border-b border-gray-100 px-5 py-4">
        <div class="flex flex-col gap-3 sm:flex-row sm:items-center sm:justify-end">
          <button
            @click="openRestoreModal"
            class="inline-flex items-center justify-center gap-2 rounded-xl border border-gray-200 bg-white px-4 py-2.5 text-sm font-medium text-gray-700 transition-colors hover:bg-gray-50 lg:px-5"
          >
            <RotateCcw :size="16" />
            Pulihkan ID
          </button>
          <button
            @click="openCreateDrawer"
            class="inline-flex items-center justify-center gap-2 rounded-xl bg-blue-600 px-4 py-2.5 text-sm font-medium text-white transition-colors hover:bg-blue-700 lg:px-5"
          >
            <Plus :size="16" />
            Tambah Rumah Sakit
          </button>
        </div>

        <div class="mt-4 border-t border-gray-100 pt-4">
          <p class="text-sm text-gray-600">
            Menampilkan
            <span class="font-semibold text-gray-900">{{ rumahSakitStore.pagination?.total ?? rumahSakitStore.hospitals.length }}</span>
            data rumah sakit
          </p>
        </div>
      </div>

      <div class="min-h-0 flex-1 overflow-auto">
      <table class="w-full text-sm">
        <thead class="sticky top-0 z-10 bg-white shadow-sm">
          <tr class="border-b border-gray-100">
            <th
              class="bg-white px-5 py-3.5 text-left text-xs font-semibold text-gray-400 uppercase tracking-wide"
            >
              Nama
            </th>
            <th
              class="bg-white px-5 py-3.5 text-left text-xs font-semibold text-gray-400 uppercase tracking-wide"
            >
              Telepon
            </th>
            <th
              class="bg-white px-5 py-3.5 text-left text-xs font-semibold text-gray-400 uppercase tracking-wide"
            >
              Email
            </th>
            <th
              class="bg-white px-5 py-3.5 text-left text-xs font-semibold text-gray-400 uppercase tracking-wide"
            >
              Username
            </th>
            <th
              class="bg-white px-5 py-3.5 text-center text-xs font-semibold text-gray-400 uppercase tracking-wide"
            >
              Aksi
            </th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-50">
          <tr
            v-for="rs in rumahSakitStore.hospitals"
            :key="rs.rumahSakitId"
            class="hover:bg-gray-50 transition-colors"
          >
            <td class="px-5 py-4 font-medium text-gray-800">{{ rs.nama }}</td>
            <td class="px-5 py-4 text-gray-500">{{ rs.nomorTelepon }}</td>
            <td class="px-5 py-4 text-gray-500">{{ rs.email || '—' }}</td>
            <td class="px-5 py-4 text-gray-500">{{ rs.username }}</td>
            <td class="px-5 py-4">
              <div class="flex items-center justify-center gap-2">
                <button
                  @click="openDetailDrawer(rs)"
                  class="flex items-center gap-1.5 px-3 py-1.5 hover:bg-green-100 text-green-600 text-xs font-medium rounded-lg transition-colors"
                >
                  Detail
                </button>
                <button
                  @click="openEditDrawer(rs)"
                  class="flex items-center gap-1.5 px-3 py-1.5 hover:bg-blue-100 text-blue-600 text-xs font-medium rounded-lg transition-colors"
                >
                  Edit
                </button>
                <button
                  @click="deleteHospital(rs.rumahSakitId)"
                  class="flex items-center gap-1.5 px-3 py-1.5 hover:bg-red-100 text-red-600 text-xs font-medium rounded-lg transition-colors"
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
        v-if="rumahSakitStore.hospitals.length === 0"
        class="flex flex-col items-center justify-center py-16 text-gray-300"
      >
        <Hospital :size="40" class="mb-3" />
        <p class="text-sm">Belum ada data rumah sakit</p>
      </div>

      <div
        v-if="(rumahSakitStore.pagination?.total ?? 0) > 0"
        class="flex items-center justify-between border-t border-gray-100 px-5 py-4"
      >
        <p class="text-sm text-gray-500">
          Menampilkan {{ pageRange.startIndex + 1 }} - {{ pageRange.endIndex }} dari
          {{ rumahSakitStore.pagination?.total ?? rumahSakitStore.hospitals.length }} data
        </p>
        <div class="flex items-center gap-2">
          <button
            type="button"
            :disabled="currentPage === 1"
            class="inline-flex items-center gap-2 rounded-lg border border-gray-200 px-3 py-2 text-sm text-gray-600 transition-colors hover:bg-gray-50 disabled:opacity-50"
            @click="currentPage -= 1"
          >
            <ChevronLeft :size="14" />
            Sebelumnya
          </button>
          <span class="min-w-20 text-center text-sm font-medium text-gray-700">
            {{ currentPage }} / {{ totalPages }}
          </span>
          <button
            type="button"
            :disabled="currentPage >= totalPages"
            class="inline-flex items-center gap-2 rounded-lg border border-gray-200 px-3 py-2 text-sm text-gray-600 transition-colors hover:bg-gray-50 disabled:opacity-50"
            @click="currentPage += 1"
          >
            Berikutnya
            <ChevronRight :size="14" />
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.backdrop-enter-active,
.backdrop-leave-active {
  transition: opacity 0.3s ease;
}

.backdrop-enter-from,
.backdrop-leave-to {
  opacity: 0;
}

.drawer-enter-active,
.drawer-leave-active {
  transition: transform 0.3s ease, opacity 0.3s ease;
}

.drawer-enter-from,
.drawer-leave-to {
  opacity: 0;
  transform: translate(-50%, -46%);
}
</style>
