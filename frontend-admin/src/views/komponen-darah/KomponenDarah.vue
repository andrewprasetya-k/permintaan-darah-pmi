<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue'
import { useKomponenStore } from '@/stores/komponen'
import { Plus, AlertCircle, Droplets, Power, ChevronLeft, ChevronRight } from '@lucide/vue'
import KomponenCreateDrawer from './KomponenCreateDrawer.vue'
import KomponenEditDrawer from './KomponenEditDrawer.vue'
import KomponenDetailDrawer from './KomponenDetailDrawer.vue'

const komponenStore = useKomponenStore()
const showCreateDrawer = ref(false)
const showEditDrawer = ref(false)
const showDetailDrawer = ref(false)
const currentPage = ref(1)
const itemsPerPage = 10

const loadKomponen = async (page = currentPage.value) => {
  const offset = (page - 1) * itemsPerPage
  await komponenStore.fetchAll({ limit: itemsPerPage, offset })
}

onMounted(async () => await loadKomponen())

const openCreateDrawer = () => {
  showCreateDrawer.value = true
}

const openEditDrawer = (komponen: any) => {
  komponenStore.selectedKomponen = komponen
  showEditDrawer.value = true
}

const openDetailDrawer = (komponen: any) => {
  komponenStore.selectedKomponen = komponen
  showDetailDrawer.value = true
}

const deleteKomponen = async (id: number) => {
  if (confirm('Yakin ingin menghapus komponen darah ini?')) {
    await komponenStore.deleteKomponen(id)
    await loadKomponen(currentPage.value)
  }
}

const toggleKomponenStatus = async (id: number, isActive: boolean) => {
  if (isActive) {
    await komponenStore.deactivate(id)
  } else {
    await komponenStore.activate(id)
  }
}

const handleSubmit = async () => {
  await loadKomponen(currentPage.value)
}

const totalPages = computed(() =>
  Math.max(1, Math.ceil((komponenStore.pagination?.total ?? 0) / itemsPerPage)),
)

const pageRange = computed(() => {
  const startIndex = (currentPage.value - 1) * itemsPerPage
  const total = komponenStore.pagination?.total ?? komponenStore.komponens.length
  const endIndex = Math.min(startIndex + komponenStore.komponens.length, total)
  return { startIndex, endIndex }
})

watch(currentPage, async (page, previousPage) => {
  if (page === previousPage) return
  await loadKomponen(page)
})
</script>

<template>
  <div class="flex h-full min-h-0 flex-col">
    <!-- Drawers -->
    <KomponenCreateDrawer
      :is-open="showCreateDrawer"
      @close="showCreateDrawer = false"
      @submit="handleSubmit"
    />
    <KomponenEditDrawer
      :is-open="showEditDrawer"
      @close="showEditDrawer = false"
      @submit="handleSubmit"
    />
    <KomponenDetailDrawer :is-open="showDetailDrawer" @close="showDetailDrawer = false" />

    <!-- Loading -->
    <div
      v-if="komponenStore.isLoading"
      class="flex flex-1 items-center justify-center py-16 text-sm text-gray-400"
    >
      <span
        class="w-5 h-5 border-2 border-gray-200 border-t-blue-500 rounded-full animate-spin mr-3"
      />
      Memuat data...
    </div>

    <!-- Error -->
    <div
      v-else-if="komponenStore.error"
      class="flex items-center gap-2 p-4 bg-red-50 border border-red-100 text-red-600 rounded-xl text-sm"
    >
      <AlertCircle :size="16" />
      {{ komponenStore.error }}
    </div>

    <!-- Table -->
    <div
      v-else
      class="flex min-h-0 flex-1 flex-col rounded-2xl border border-gray-100 bg-white overflow-hidden"
    >
      <div class="border-b border-gray-100 px-5 py-4">
        <div class="flex justify-end">
          <button
            @click="openCreateDrawer"
            class="inline-flex items-center justify-center gap-2 rounded-xl bg-blue-600 px-4 py-2.5 text-sm font-medium text-white transition-colors hover:bg-blue-700 lg:px-5"
          >
            <Plus :size="16" />
            Tambah Komponen
          </button>
        </div>

        <div class="mt-4 border-t border-gray-100 pt-4">
          <p class="text-sm text-gray-600">
            Menampilkan
            <span class="font-semibold text-gray-900">{{ komponenStore.pagination?.total ?? komponenStore.komponens.length }}</span>
            data komponen darah
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
                Nama Komponen
              </th>
              <th
                class="bg-white px-5 py-3.5 text-left text-xs font-semibold text-gray-400 uppercase tracking-wide"
              >
                Kode
              </th>
              <th
                class="bg-white px-5 py-3.5 text-left text-xs font-semibold text-gray-400 uppercase tracking-wide"
              >
                Status
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
              v-for="kom in komponenStore.komponens"
              :key="kom.komponenId"
              class="hover:bg-gray-50 transition-colors"
            >
              <td class="px-5 py-4 font-medium text-gray-800">{{ kom.komponenDarah }}</td>
              <td class="px-5 py-4 text-gray-500">{{ kom.komponenKode }}</td>
              <td class="px-5 py-4">
                <span
                  class="inline-flex px-2.5 py-1 rounded-lg text-xs font-medium"
                  :class="kom.isActive ? 'bg-green-50 text-green-600' : 'bg-gray-50 text-gray-600'"
                >
                  {{ kom.isActive ? 'Aktif' : 'Nonaktif' }}
                </span>
              </td>
              <td class="px-5 py-4">
                <div class="flex items-center justify-center gap-2">
                  <button
                    @click="openDetailDrawer(kom)"
                    class="flex items-center gap-1.5 px-3 py-1.5 hover:bg-green-100 text-green-600 text-xs font-medium rounded-lg transition-colors"
                  >
                    Detail
                  </button>
                  <button
                    @click="openEditDrawer(kom)"
                    class="flex items-center gap-1.5 px-3 py-1.5 hover:bg-blue-100 text-blue-600 text-xs font-medium rounded-lg transition-colors"
                  >
                    Edit
                  </button>
                  <button
                    @click="toggleKomponenStatus(kom.komponenId, kom.isActive)"
                    class="flex items-center gap-1.5 px-3 py-1.5 text-xs font-medium rounded-lg transition-colors"
                    :class="
                      kom.isActive
                        ? 'text-amber-600 hover:bg-amber-100'
                        : 'text-emerald-600 hover:bg-emerald-100'
                    "
                  >
                    <Power :size="14" />
                    {{ kom.isActive ? 'Nonaktifkan' : 'Aktifkan' }}
                  </button>
                  <button
                    @click="deleteKomponen(kom.komponenId)"
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
        v-if="komponenStore.komponens.length === 0"
        class="flex flex-col items-center justify-center py-16 text-gray-300"
      >
        <Droplets :size="40" class="mb-3" />
        <p class="text-sm">Belum ada data komponen darah</p>
      </div>

      <div
        v-if="(komponenStore.pagination?.total ?? 0) > 0"
        class="flex items-center justify-between border-t border-gray-100 px-5 py-4"
      >
        <p class="text-sm text-gray-500">
          Menampilkan {{ pageRange.startIndex + 1 }} - {{ pageRange.endIndex }} dari
          {{ komponenStore.pagination?.total ?? komponenStore.komponens.length }} data
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
