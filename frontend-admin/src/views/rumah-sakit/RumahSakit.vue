<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRumahSakitStore } from '@/stores/rumah-sakit'
import { Plus, Pencil, Trash2, Eye, AlertCircle, Hospital } from '@lucide/vue'
import RumahSakitCreateDrawer from './RumahSakitCreateDrawer.vue'
import RumahSakitEditDrawer from './RumahSakitEditDrawer.vue'
import RumahSakitDetailDrawer from './RumahSakitDetailDrawer.vue'

const rumahSakitStore = useRumahSakitStore()
const showCreateDrawer = ref(false)
const showEditDrawer = ref(false)
const showDetailDrawer = ref(false)

onMounted(async () => await rumahSakitStore.fetchAll())

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
  }
}

const handleSubmit = () => {
  rumahSakitStore.fetchAll()
}
</script>

<template>
  <div>
    <!-- Header -->
    <div class="flex items-center justify-between mb-6">
      <button
        @click="openCreateDrawer"
        class="flex items-center gap-2 px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white text-sm font-medium rounded-xl transition-colors"
      >
        <Plus :size="16" />
        Tambah Rumah Sakit
      </button>
    </div>

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
    <div class="bg-white rounded-2xl border border-gray-100 overflow-hidden">
      <table class="w-full text-sm">
        <thead>
          <tr class="border-b border-gray-100">
            <th
              class="px-5 py-3.5 text-left text-xs font-semibold text-gray-400 uppercase tracking-wide"
            >
              Nama
            </th>
            <th
              class="px-5 py-3.5 text-left text-xs font-semibold text-gray-400 uppercase tracking-wide"
            >
              Telepon
            </th>
            <th
              class="px-5 py-3.5 text-left text-xs font-semibold text-gray-400 uppercase tracking-wide"
            >
              Email
            </th>
            <th
              class="px-5 py-3.5 text-left text-xs font-semibold text-gray-400 uppercase tracking-wide"
            >
              Username
            </th>
            <th
              class="px-5 py-3.5 text-center text-xs font-semibold text-gray-400 uppercase tracking-wide"
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

      <div
        v-if="rumahSakitStore.hospitals.length === 0"
        class="flex flex-col items-center justify-center py-16 text-gray-300"
      >
        <Hospital :size="40" class="mb-3" />
        <p class="text-sm">Belum ada data rumah sakit</p>
      </div>
    </div>
  </div>
</template>
