<script setup lang="ts">
import { onMounted } from 'vue'
import { usePermintaanStore } from '@/stores/permintaan'

const permintaanStore = usePermintaanStore()

onMounted(async () => {
  await permintaanStore.fetchAll()
})
</script>

<template>
  <div class="p-6">
    <h1 class="text-3xl font-bold">Daftar Permintaan Darah</h1>

    <div v-if="permintaanStore.isLoading" class="mt-4">Loading...</div>
    <div v-else-if="permintaanStore.error" class="mt-4 p-4 bg-red-100 text-red-700">
      {{ permintaanStore.error }}
    </div>
    <div v-else class="mt-4">
      <table class="w-full border">
        <thead class="bg-gray-100">
          <tr>
            <th class="p-2 border text-left">Pasien</th>
            <th class="p-2 border text-left">Tipe Darah</th>
            <th class="p-2 border text-left">Jumlah</th>
            <th class="p-2 border text-left">Status</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="req in permintaanStore.requests" :key="req.id">
            <td class="p-2 border">{{ req.patientName }}</td>
            <td class="p-2 border">{{ req.bloodType }}</td>
            <td class="p-2 border">{{ req.quantity }}</td>
            <td class="p-2 border">{{ req.status }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>
