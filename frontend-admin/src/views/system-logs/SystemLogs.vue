<script setup lang="ts">
import { computed, onMounted } from 'vue'
import { useLogsStore } from '@/stores/logs'

const logsStore = useLogsStore()

const logs = computed(() => logsStore.systemLogs)

const formatDate = (date: string) =>
  new Date(date).toLocaleString('id-ID', {
    year: 'numeric',
    month: 'short',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
  })

onMounted(async () => {
  await logsStore.fetchSystemLogs({ limit: 50, offset: 0 })
  logsStore.connectRealtime()
})
</script>

<template>
  <div class="flex h-full min-h-0 flex-col space-y-6">
    <div class="flex items-center justify-between">
      <div
        class="inline-flex items-center gap-2 px-3 py-2 rounded-xl text-xs font-medium"
        :class="
          logsStore.isRealtimeConnected
            ? 'bg-emerald-50 text-emerald-700'
            : 'bg-gray-100 text-gray-500'
        "
      >
        <span
          class="w-2 h-2 rounded-full"
          :class="logsStore.isRealtimeConnected ? 'bg-emerald-500' : 'bg-gray-400'"
        />
        {{ logsStore.isRealtimeConnected ? 'Realtime online' : 'Realtime offline' }}
      </div>
    </div>

    <div
      v-if="logsStore.error"
      class="rounded-2xl border border-red-100 bg-red-50 px-4 py-3 text-sm text-red-700"
    >
      {{ logsStore.error }}
    </div>

    <div
      class="flex min-h-0 flex-1 flex-col overflow-hidden rounded-2xl border border-gray-100 bg-white"
    >
      <div class="min-h-0 flex-1 overflow-auto">
        <table class="w-full text-sm">
          <thead class="sticky top-0 z-10 bg-gray-50 text-left text-xs uppercase tracking-wide text-gray-500 shadow-sm">
            <tr>
              <th class="bg-gray-50 px-4 py-3">Waktu</th>
              <th class="bg-gray-50 px-4 py-3">User</th>
              <th class="bg-gray-50 px-4 py-3">Role</th>
              <th class="bg-gray-50 px-4 py-3">Action</th>
              <th class="bg-gray-50 px-4 py-3">Target</th>
              <th class="bg-gray-50 px-4 py-3">Catatan</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-100">
            <tr v-if="logsStore.isLoading">
              <td colspan="6" class="px-4 py-10 text-center text-gray-500">
                Memuat log aktivitas...
              </td>
            </tr>
            <tr v-else-if="logs.length === 0">
              <td colspan="6" class="px-4 py-10 text-center text-gray-400">
                Belum ada log aktivitas.
              </td>
            </tr>
            <tr v-for="log in logs" :key="log.sysLogId" class="align-top">
              <td class="px-4 py-3 text-gray-500">{{ formatDate(log.createdAt) }}</td>
              <td class="px-4 py-3 font-medium text-gray-900">{{ log.sysUserNama }}</td>
              <td class="px-4 py-3 text-gray-600">{{ log.sysUserRole }}</td>
              <td class="px-4 py-3">
                <span class="rounded-lg bg-blue-50 px-2.5 py-1 text-xs font-semibold text-blue-700">
                  {{ log.sysAction }}
                </span>
              </td>
              <td class="px-4 py-3 text-gray-600">
                {{ log.sysTargetTable || '-' }}
                <span v-if="log.sysTargetId" class="text-gray-400">#{{ log.sysTargetId }}</span>
              </td>
              <td class="px-4 py-3 text-gray-700">{{ log.sysNotes }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>
