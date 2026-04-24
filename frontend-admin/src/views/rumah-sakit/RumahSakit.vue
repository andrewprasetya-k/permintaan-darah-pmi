<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue'
import { useRumahSakitStore } from '@/stores/rumah-sakit'
import { Plus, AlertCircle, Hospital, ChevronLeft, ChevronRight, RotateCcw, Trash2 } from '@lucide/vue'
import RumahSakitCreateDrawer from './RumahSakitCreateDrawer.vue'
import RumahSakitEditDrawer from './RumahSakitEditDrawer.vue'
import RumahSakitDetailDrawer from './RumahSakitDetailDrawer.vue'
import AppModal from '@/components/feedback/AppModal.vue'
import AppFlag from '@/components/feedback/AppFlag.vue'

const rumahSakitStore = useRumahSakitStore()
const showCreateDrawer = ref(false)
const showEditDrawer = ref(false)
const showDetailDrawer = ref(false)
const pendingAction = ref<{ type: 'delete' | 'restore'; id: string; name: string } | null>(null)
const flag = ref<{ variant: 'success' | 'error'; title: string; message?: string } | null>(null)
const currentPage = ref(1)
const itemsPerPage = 10

const loadHospitals = async (page = currentPage.value) => {
  const offset = (page - 1) * itemsPerPage
  await rumahSakitStore.fetchAll({
    status: rumahSakitStore.currentFilter,
    limit: itemsPerPage,
    offset,
  })
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

const handleFilterChange = async (event: Event) => {
  const target = event.target as HTMLSelectElement
  currentPage.value = 1
  await rumahSakitStore.fetchAll({ status: target.value, limit: itemsPerPage, offset: 0 })
}

const handleSubmit = async () => {
  await loadHospitals(currentPage.value)
}

const openDeleteDialog = (id: string, name: string) => {
  pendingAction.value = { type: 'delete', id, name }
}

const openRestoreDialog = (id: string, name: string) => {
  pendingAction.value = { type: 'restore', id, name }
}

const closeActionDialog = () => {
  pendingAction.value = null
}

const submitAction = async () => {
  if (!pendingAction.value) return
  try {
    if (pendingAction.value.type === 'delete') {
      await rumahSakitStore.deleteHospital(pendingAction.value.id)
      flag.value = {
        variant: 'success',
        title: 'Rumah sakit dinonaktifkan',
        message: `${pendingAction.value.name} dipindahkan ke status nonaktif.`,
      }
    } else {
      await rumahSakitStore.restore(pendingAction.value.id)
      flag.value = {
        variant: 'success',
        title: 'Rumah sakit dipulihkan',
        message: `${pendingAction.value.name} kembali aktif di sistem.`,
      }
    }
    await loadHospitals(currentPage.value)
  } catch (error) {
    flag.value = {
      variant: 'error',
      title: 'Operasi gagal',
      message: error instanceof Error ? error.message : 'Gagal memproses rumah sakit.',
    }
  } finally {
    closeActionDialog()
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
    <AppFlag
      v-if="flag"
      :variant="flag.variant"
      :title="flag.title"
      :message="flag.message"
      @close="flag = null"
    />

    <AppModal
      :is-open="!!pendingAction"
      :title="pendingAction?.type === 'delete' ? 'Nonaktifkan Rumah Sakit' : 'Pulihkan Rumah Sakit'"
      :description="
        pendingAction?.type === 'delete'
          ? `Rumah sakit ${pendingAction?.name} akan dipindahkan ke status nonaktif.`
          : `Rumah sakit ${pendingAction?.name} akan dipulihkan ke status aktif.`
      "
      @close="closeActionDialog"
    >
      <template #icon>
        <div
          class="flex h-12 w-12 items-center justify-center rounded-2xl"
          :class="
            pendingAction?.type === 'delete'
              ? 'bg-red-50 text-red-600'
              : 'bg-emerald-50 text-emerald-600'
          "
        >
          <Trash2 v-if="pendingAction?.type === 'delete'" :size="20" />
          <RotateCcw v-else :size="20" />
        </div>
      </template>

      <template #footer>
        <button
          type="button"
          class="flex-1 rounded-xl bg-gray-100 px-4 py-2.5 text-sm font-medium text-gray-700 transition-colors hover:bg-gray-200"
          @click="closeActionDialog"
        >
          Batal
        </button>
        <button
          type="button"
          class="flex-1 rounded-xl px-4 py-2.5 text-sm font-medium text-white transition-colors"
          :class="pendingAction?.type === 'delete' ? 'bg-red-600 hover:bg-red-700' : 'bg-emerald-600 hover:bg-emerald-700'"
          @click="submitAction"
        >
          {{ pendingAction?.type === 'delete' ? 'Nonaktifkan' : 'Pulihkan' }}
        </button>
      </template>
    </AppModal>

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

    <!-- Error -->
    <div
      v-if="rumahSakitStore.error"
      class="flex items-center gap-2 p-4 bg-red-50 border border-red-100 text-red-600 rounded-xl text-sm"
    >
      <AlertCircle :size="16" />
      {{ rumahSakitStore.error }}
    </div>

    <!-- Table -->
    <div
      v-else
      class="flex min-h-0 flex-1 flex-col rounded-2xl border border-gray-100 bg-white overflow-hidden"
    >
      <div class="border-b border-gray-100 px-5 py-4">
        <div class="flex flex-col gap-3 lg:flex-row lg:items-center lg:justify-between">
          <select
            :value="rumahSakitStore.currentFilter"
            @change="handleFilterChange"
            class="w-full rounded-xl border border-gray-200 bg-white px-3.5 py-2.5 text-sm text-gray-700 outline-none transition-all focus:border-blue-400 focus:ring-2 focus:ring-blue-100 sm:w-[190px]"
          >
            <option value="active">Aktif</option>
            <option value="deleted">Dihapus</option>
            <option value="all">Semua</option>
          </select>
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
                  v-if="!rs.isDeleted"
                  @click="openEditDrawer(rs)"
                  class="flex items-center gap-1.5 px-3 py-1.5 hover:bg-blue-100 text-blue-600 text-xs font-medium rounded-lg transition-colors"
                >
                  Edit
                </button>
                <button
                  v-if="!rs.isDeleted"
                  @click="openDeleteDialog(rs.rumahSakitId, rs.nama)"
                  class="flex items-center gap-1.5 px-3 py-1.5 hover:bg-red-100 text-red-600 text-xs font-medium rounded-lg transition-colors"
                >
                  Nonaktifkan
                </button>
                <button
                  v-else
                  @click="openRestoreDialog(rs.rumahSakitId, rs.nama)"
                  class="flex items-center gap-1.5 px-3 py-1.5 hover:bg-emerald-100 text-emerald-600 text-xs font-medium rounded-lg transition-colors"
                >
                  <RotateCcw :size="14" />
                  Restore
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
