<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useKomponenStore } from '@/stores/komponen'
import { Plus, Pencil, Trash2, Eye, AlertCircle, Droplets } from '@lucide/vue'
import KomponenCreateDrawer from './KomponenCreateDrawer.vue'
import KomponenEditDrawer from './KomponenEditDrawer.vue'
import KomponenDetailDrawer from './KomponenDetailDrawer.vue'

const komponenStore = useKomponenStore()
const showCreateDrawer = ref(false)
const showEditDrawer = ref(false)
const showDetailDrawer = ref(false)

onMounted(async () => await komponenStore.fetchAll())

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
  }
}

const handleSubmit = () => {
  komponenStore.fetchAll()
}
</script>

<template>
  <div class="flex h-full min-h-0 flex-col">
    <!-- Header -->
    <div class="flex items-center justify-between mb-6">
      <button
        @click="openCreateDrawer"
        class="flex items-center gap-2 px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white text-sm font-medium rounded-xl transition-colors"
      >
        <Plus :size="16" />
        Tambah Komponen
      </button>
    </div>

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
    </div>
  </div>
</template>
