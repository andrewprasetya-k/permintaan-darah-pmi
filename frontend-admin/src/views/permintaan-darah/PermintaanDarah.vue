<script setup lang="ts">
import { onMounted } from 'vue'
import { usePermintaanStore } from '@/stores/permintaan'
import { AlertCircle, Droplets } from '@lucide/vue'

const permintaanStore = usePermintaanStore()

onMounted(async () => await permintaanStore.fetchAll())

const formatDate = (date: string) =>
  new Date(date).toLocaleDateString('id-ID', { year: 'numeric', month: 'long', day: 'numeric' })

const statusStyle: Record<string, string> = {
  dibuat: 'bg-amber-50 text-amber-600 ',
  diproses: 'bg-blue-50 text-blue-600 ',
  bisa_diambil: 'bg-violet-50 text-violet-600 ',
  selesai: 'bg-green-50 text-green-600 ',
  ditolak: 'bg-red-50 text-red-600 ',
}
</script>

<template>
  <div>
    <!-- Loading -->
    <div
      v-if="permintaanStore.isLoading"
      class="flex items-center justify-center py-16 text-sm text-gray-400"
    >
      <span
        class="w-5 h-5 border-2 border-gray-200 border-t-blue-500 rounded-full animate-spin mr-3"
      />
      Memuat data...
    </div>

    <!-- Error -->
    <div
      v-else-if="permintaanStore.error"
      class="flex items-center gap-2 p-4 bg-red-50 border border-red-100 text-red-600 rounded-xl text-sm"
    >
      <AlertCircle :size="16" />
      {{ permintaanStore.error }}
    </div>

    <!-- Table -->
    <div v-else class="bg-white rounded-2xl border border-gray-100 overflow-hidden">
      <table class="w-full text-sm">
        <thead>
          <tr class="border-b border-gray-100">
            <th
              class="px-5 py-3.5 text-left text-xs font-semibold text-gray-400 uppercase tracking-wide"
            >
              Pasien
            </th>
            <th
              class="px-5 py-3.5 text-left text-xs font-semibold text-gray-400 uppercase tracking-wide"
            >
              Golongan Darah
            </th>
            <th
              class="px-5 py-3.5 text-left text-xs font-semibold text-gray-400 uppercase tracking-wide"
            >
              Status
            </th>
            <th
              class="px-5 py-3.5 text-left text-xs font-semibold text-gray-400 uppercase tracking-wide"
            >
              Tanggal
            </th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-50">
          <tr
            v-for="req in permintaanStore.requests"
            :key="req.permintaanDarahId"
            class="hover:bg-gray-50 transition-colors"
          >
            <td class="px-5 py-4 font-medium text-gray-800">{{ req.namaPasien }}</td>
            <td class="px-5 py-4">
              <span
                class="inline-flex items-center gap-1 px-2.5 py-1 bg-red-50 text-red-600 text-xs font-semibold rounded-lg"
              >
                {{ req.golonganDarah }} {{ req.rhesusDarah }}
              </span>
            </td>
            <td class="px-5 py-4">
              <span
                class="inline-flex items-center px-2.5 py-1 text-xs font-medium rounded-lg capitalize"
                :class="statusStyle[req.status] ?? 'bg-gray-50 text-gray-500'"
              >
                {{ req.status }}
              </span>
            </td>
            <td class="px-5 py-4 text-gray-500">{{ formatDate(req.createdAt) }}</td>
          </tr>
        </tbody>
      </table>

      <div
        v-if="permintaanStore.requests.length === 0"
        class="flex flex-col items-center justify-center py-16 text-gray-300"
      >
        <Droplets :size="40" class="mb-3" />
        <p class="text-sm">Belum ada permintaan darah</p>
      </div>
    </div>
  </div>
</template>
