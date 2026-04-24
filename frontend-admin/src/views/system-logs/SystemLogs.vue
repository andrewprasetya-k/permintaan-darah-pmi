<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue'
import { useLogsStore } from '@/stores/logs'
import { Search, Radio, ChevronLeft, ChevronRight, FilterX } from '@lucide/vue'

const logsStore = useLogsStore()

const activeTab = ref<'system' | 'status'>('system')
const filterMode = ref<'all' | 'user' | 'action' | 'table' | 'target'>('all')
const filterValue = ref('')
const currentPage = ref(1)
const itemsPerPage = 20

const actionOptions = ['LOGIN', 'LOGIN_FAILED', 'CREATE', 'UPDATE', 'SOFT_DELETE', 'RESTORE']
const tableOptions = ['admins', 'rumah_sakit', 'komponen_darah', 'permintaan_darah']

const formatDate = (date: string) =>
  new Date(date).toLocaleString('id-ID', {
    year: 'numeric',
    month: 'short',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
  })

const displaySystemLogs = computed(() => logsStore.systemLogs)
const displayStatusLogs = computed(() => logsStore.statusLogs)

const totalPages = computed(() =>
  Math.max(1, Math.ceil((logsStore.pagination?.total ?? 0) / itemsPerPage)),
)

const pageRange = computed(() => {
  const items = activeTab.value === 'system' ? displaySystemLogs.value : displayStatusLogs.value
  const startIndex = (currentPage.value - 1) * itemsPerPage
  const total = logsStore.pagination?.total ?? items.length
  const endIndex = Math.min(startIndex + items.length, total)
  return { startIndex, endIndex }
})

const loadLogs = async (page = currentPage.value) => {
  const offset = (page - 1) * itemsPerPage

  if (activeTab.value === 'status') {
    await logsStore.fetchStatusLogs({ limit: itemsPerPage, offset })
    return
  }

  const params = { limit: itemsPerPage, offset }
  const value = filterValue.value.trim()

  switch (filterMode.value) {
    case 'user':
      if (!value) {
        await logsStore.fetchSystemLogs(params)
        return
      }
      await logsStore.fetchSystemLogsByUserId(value, params)
      return
    case 'action':
      if (!value) {
        await logsStore.fetchSystemLogs(params)
        return
      }
      await logsStore.fetchSystemLogsByAction(value, params)
      return
    case 'table':
      if (!value) {
        await logsStore.fetchSystemLogs(params)
        return
      }
      await logsStore.fetchSystemLogsByTable(value, params)
      return
    case 'target':
      if (!value) {
        await logsStore.fetchSystemLogs(params)
        return
      }
      await logsStore.fetchSystemLogsByTargetId(value, params)
      return
    default:
      await logsStore.fetchSystemLogs(params)
  }
}

const applyFilters = async () => {
  currentPage.value = 1
  await loadLogs(1)
}

const resetFilters = async () => {
  filterMode.value = 'all'
  filterValue.value = ''
  currentPage.value = 1
  await loadLogs(1)
}

watch(activeTab, async () => {
  currentPage.value = 1
  await loadLogs(1)
})

watch(currentPage, async (page, previousPage) => {
  if (page === previousPage) return
  await loadLogs(page)
})

onMounted(async () => {
  await loadLogs()
  logsStore.connectRealtime()
})
</script>

<template>
  <div class="flex h-full min-h-0 flex-col space-y-6">
    <div class="flex flex-col gap-4 rounded-2xl border border-gray-100 bg-white p-5">
      <div class="flex flex-col gap-3 lg:flex-row lg:items-center lg:justify-between">
        <div class="inline-flex w-full rounded-xl bg-gray-100 p-1 lg:w-auto">
          <button
            type="button"
            class="flex-1 rounded-lg px-4 py-2 text-sm font-medium transition-colors lg:flex-none"
            :class="
              activeTab === 'system'
                ? 'bg-white text-gray-900 shadow-sm'
                : 'text-gray-500 hover:text-gray-700'
            "
            @click="activeTab = 'system'"
          >
            System Logs
          </button>
          <button
            type="button"
            class="flex-1 rounded-lg px-4 py-2 text-sm font-medium transition-colors lg:flex-none"
            :class="
              activeTab === 'status'
                ? 'bg-white text-gray-900 shadow-sm'
                : 'text-gray-500 hover:text-gray-700'
            "
            @click="activeTab = 'status'"
          >
            Status Logs
          </button>
        </div>

        <div
          class="inline-flex items-center gap-2 rounded-xl px-3 py-2 text-xs font-medium"
          :class="
            logsStore.isRealtimeConnected
              ? 'bg-emerald-50 text-emerald-700'
              : 'bg-gray-100 text-gray-500'
          "
        >
          <Radio :size="14" />
          {{ logsStore.isRealtimeConnected ? 'Realtime online' : 'Realtime offline' }}
        </div>
      </div>

      <div v-if="activeTab === 'system'" class="grid gap-3 lg:grid-cols-[180px_minmax(0,1fr)_auto_auto]">
        <select
          v-model="filterMode"
          class="rounded-xl border border-gray-200 bg-white px-3.5 py-2.5 text-sm text-gray-700 outline-none transition-all focus:border-blue-400 focus:ring-2 focus:ring-blue-100"
        >
          <option value="all">Semua Log</option>
          <option value="user">User ID</option>
          <option value="action">Action</option>
          <option value="table">Table</option>
          <option value="target">Target ID</option>
        </select>

        <div class="relative">
          <Search
            v-if="filterMode === 'user' || filterMode === 'target'"
            :size="16"
            class="pointer-events-none absolute left-3 top-1/2 -translate-y-1/2 text-gray-400"
          />
          <input
            v-if="filterMode === 'user' || filterMode === 'target'"
            v-model="filterValue"
            type="text"
            :placeholder="filterMode === 'user' ? 'Masukkan user ID' : 'Masukkan target ID'"
            class="w-full rounded-xl border border-gray-200 bg-white py-2.5 pl-9 pr-3 text-sm text-gray-700 outline-none transition-all focus:border-blue-400 focus:ring-2 focus:ring-blue-100"
            @keyup.enter="applyFilters"
          />
          <select
            v-else-if="filterMode === 'action'"
            v-model="filterValue"
            class="w-full rounded-xl border border-gray-200 bg-white px-3.5 py-2.5 text-sm text-gray-700 outline-none transition-all focus:border-blue-400 focus:ring-2 focus:ring-blue-100"
          >
            <option value="">Pilih action</option>
            <option v-for="action in actionOptions" :key="action" :value="action">{{ action }}</option>
          </select>
          <select
            v-else-if="filterMode === 'table'"
            v-model="filterValue"
            class="w-full rounded-xl border border-gray-200 bg-white px-3.5 py-2.5 text-sm text-gray-700 outline-none transition-all focus:border-blue-400 focus:ring-2 focus:ring-blue-100"
          >
            <option value="">Pilih table</option>
            <option v-for="table in tableOptions" :key="table" :value="table">{{ table }}</option>
          </select>
          <div
            v-else
            class="flex h-full items-center rounded-xl border border-dashed border-gray-200 px-3.5 py-2.5 text-sm text-gray-400"
          >
            Menampilkan seluruh system logs
          </div>
        </div>

        <button
          type="button"
          class="inline-flex items-center justify-center rounded-xl bg-blue-600 px-4 py-2.5 text-sm font-medium text-white transition-colors hover:bg-blue-700"
          @click="applyFilters"
        >
          Terapkan
        </button>

        <button
          type="button"
          class="inline-flex items-center justify-center gap-2 rounded-xl border border-gray-200 bg-white px-4 py-2.5 text-sm font-medium text-gray-700 transition-colors hover:bg-gray-50"
          @click="resetFilters"
        >
          <FilterX :size="14" />
          Reset
        </button>
      </div>
    </div>

    <div
      v-if="logsStore.error"
      class="rounded-2xl border border-red-100 bg-red-50 px-4 py-3 text-sm text-red-700"
    >
      {{ logsStore.error }}
    </div>

    <div class="flex min-h-0 flex-1 flex-col overflow-hidden rounded-2xl border border-gray-100 bg-white">
      <div class="min-h-0 flex-1 overflow-auto">
        <table class="w-full text-sm">
          <thead class="sticky top-0 z-10 bg-gray-50 text-left text-xs uppercase tracking-wide text-gray-500 shadow-sm">
            <tr v-if="activeTab === 'system'">
              <th class="bg-gray-50 px-4 py-3">Waktu</th>
              <th class="bg-gray-50 px-4 py-3">User</th>
              <th class="bg-gray-50 px-4 py-3">Role</th>
              <th class="bg-gray-50 px-4 py-3">Action</th>
              <th class="bg-gray-50 px-4 py-3">Target</th>
              <th class="bg-gray-50 px-4 py-3">Catatan</th>
            </tr>
            <tr v-else>
              <th class="bg-gray-50 px-4 py-3">Waktu</th>
              <th class="bg-gray-50 px-4 py-3">Permintaan</th>
              <th class="bg-gray-50 px-4 py-3">Admin</th>
              <th class="bg-gray-50 px-4 py-3">Status</th>
              <th class="bg-gray-50 px-4 py-3">Catatan</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-100">
            <tr v-if="logsStore.isLoading">
              <td :colspan="activeTab === 'system' ? 6 : 5" class="px-4 py-10 text-center text-gray-500">
                Memuat data log...
              </td>
            </tr>
            <tr
              v-else-if="activeTab === 'system' && displaySystemLogs.length === 0"
            >
              <td colspan="6" class="px-4 py-10 text-center text-gray-400">
                Belum ada system logs.
              </td>
            </tr>
            <tr
              v-else-if="activeTab === 'status' && displayStatusLogs.length === 0"
            >
              <td colspan="5" class="px-4 py-10 text-center text-gray-400">
                Belum ada status logs.
              </td>
            </tr>

            <template v-if="activeTab === 'system'">
              <tr v-for="log in displaySystemLogs" :key="`sys-${log.sysLogId}`" class="align-top">
                <td class="px-4 py-3 text-gray-500">{{ formatDate(log.createdAt) }}</td>
                <td class="px-4 py-3 font-medium text-gray-900">
                  {{ log.sysUserNama }}
                  <div v-if="log.sysUserId" class="mt-1 text-xs font-normal text-gray-400">
                    {{ log.sysUserId }}
                  </div>
                </td>
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
            </template>

            <template v-else>
              <tr v-for="log in displayStatusLogs" :key="`status-${log.logId}`" class="align-top">
                <td class="px-4 py-3 text-gray-500">{{ formatDate(log.createdAt) }}</td>
                <td class="px-4 py-3 font-medium text-gray-900">{{ log.permintaanDarahId }}</td>
                <td class="px-4 py-3 text-gray-600">
                  {{ log.adminNama }}
                  <div v-if="log.adminId" class="mt-1 text-xs text-gray-400">{{ log.adminId }}</div>
                </td>
                <td class="px-4 py-3">
                  <div class="flex flex-wrap items-center gap-2 text-xs">
                    <span
                      v-if="log.statusFrom"
                      class="rounded-lg bg-gray-100 px-2.5 py-1 font-medium text-gray-600"
                    >
                      {{ log.statusFrom }}
                    </span>
                    <span v-if="log.statusFrom" class="text-gray-300">→</span>
                    <span class="rounded-lg bg-emerald-50 px-2.5 py-1 font-medium text-emerald-700">
                      {{ log.statusTo }}
                    </span>
                  </div>
                </td>
                <td class="px-4 py-3 text-gray-700">{{ log.notes || '-' }}</td>
              </tr>
            </template>
          </tbody>
        </table>
      </div>

      <div
        v-if="(logsStore.pagination?.total ?? 0) > 0"
        class="flex items-center justify-between border-t border-gray-100 px-4 py-4"
      >
        <p class="text-sm text-gray-500">
          Menampilkan {{ pageRange.startIndex + 1 }} - {{ pageRange.endIndex }} dari
          {{ logsStore.pagination?.total ?? 0 }} data
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
