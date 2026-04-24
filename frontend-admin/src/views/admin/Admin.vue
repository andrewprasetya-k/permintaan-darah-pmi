<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue'
import { useAdminStore } from '@/stores/admin'
import { Plus, AlertCircle, Users, Shield, RotateCcw, ChevronLeft, ChevronRight } from '@lucide/vue'
import AdminCreateDrawer from './AdminCreateDrawer.vue'
import AdminEditDrawer from './AdminEditDrawer.vue'
import AdminDetailDrawer from './AdminDetailDrawer.vue'

const adminStore = useAdminStore()
const showCreateDrawer = ref(false)
const showEditDrawer = ref(false)
const showDetailDrawer = ref(false)
const currentPage = ref(1)
const itemsPerPage = 10

const loadAdmins = async (page = currentPage.value) => {
  const offset = (page - 1) * itemsPerPage
  await adminStore.fetchAll({
    status: adminStore.currentFilter,
    limit: itemsPerPage,
    offset,
  })
}

onMounted(async () => await loadAdmins())

const handleFilterChange = async (event: Event) => {
  const target = event.target as HTMLSelectElement
  currentPage.value = 1
  await adminStore.fetchAll({ status: target.value, limit: itemsPerPage, offset: 0 })
}

const openCreateDrawer = () => {
  showCreateDrawer.value = true
}

const openEditDrawer = (admin: any) => {
  adminStore.selectedAdmin = admin
  showEditDrawer.value = true
}

const openDetailDrawer = (admin: any) => {
  adminStore.selectedAdmin = admin
  showDetailDrawer.value = true
}

const deleteAdmin = async (id: string) => {
  if (confirm('Yakin ingin menghapus admin ini?')) {
    await adminStore.deleteAdmin(id)
    await loadAdmins(currentPage.value)
  }
}

const restoreAdmin = async (id: string) => {
  if (confirm('Yakin ingin memulihkan admin ini?')) {
    await adminStore.restore(id)
    await loadAdmins(currentPage.value)
  }
}

const handleSubmit = async () => {
  await loadAdmins(currentPage.value)
}

const totalPages = computed(() =>
  Math.max(1, Math.ceil((adminStore.pagination?.total ?? 0) / itemsPerPage)),
)

const pageRange = computed(() => {
  const startIndex = (currentPage.value - 1) * itemsPerPage
  const total = adminStore.pagination?.total ?? adminStore.admins.length
  const endIndex = Math.min(startIndex + adminStore.admins.length, total)
  return { startIndex, endIndex }
})

watch(currentPage, async (page, previousPage) => {
  if (page === previousPage) return
  await loadAdmins(page)
})
</script>

<template>
  <div class="flex h-full min-h-0 flex-col">
    <!-- Drawers -->
    <AdminCreateDrawer
      :is-open="showCreateDrawer"
      @close="showCreateDrawer = false"
      @submit="handleSubmit"
    />
    <AdminEditDrawer
      :is-open="showEditDrawer"
      @close="showEditDrawer = false"
      @submit="handleSubmit"
    />
    <AdminDetailDrawer :is-open="showDetailDrawer" @close="showDetailDrawer = false" />

    <!-- Loading -->
    <div
      v-if="adminStore.isLoading"
      class="flex flex-1 items-center justify-center py-16 text-sm text-gray-400"
    >
      <span
        class="w-5 h-5 border-2 border-gray-200 border-t-blue-500 rounded-full animate-spin mr-3"
      />
      Memuat data...
    </div>

    <!-- Error -->
    <div
      v-else-if="adminStore.error"
      class="flex items-center gap-2 p-4 bg-red-50 border border-red-100 text-red-600 rounded-xl text-sm"
    >
      <AlertCircle :size="16" />
      {{ adminStore.error }}
    </div>

    <!-- Table -->
    <div
      v-else
      class="flex min-h-0 flex-1 flex-col rounded-2xl border border-gray-100 bg-white overflow-hidden"
    >
      <div class="border-b border-gray-100 px-5 py-4">
        <div class="flex flex-col gap-3 lg:flex-row lg:items-center lg:justify-between">
          <select
            :value="adminStore.currentFilter"
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
            Tambah Admin
          </button>
        </div>

        <div class="mt-4 border-t border-gray-100 pt-4">
          <p class="text-sm text-gray-600">
            Menampilkan
            <span class="font-semibold text-gray-900">{{ adminStore.pagination?.total ?? adminStore.admins.length }}</span>
            data admin
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
                Username
              </th>
              <th
                class="bg-white px-5 py-3.5 text-left text-xs font-semibold text-gray-400 uppercase tracking-wide"
              >
                Email
              </th>
              <th
                class="bg-white px-5 py-3.5 text-left text-xs font-semibold text-gray-400 uppercase tracking-wide"
              >
                Role
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
              v-for="admin in adminStore.admins"
              :key="admin.adminId"
              class="hover:bg-gray-50 transition-colors"
            >
              <td class="px-5 py-4 font-medium text-gray-800">{{ admin.adminName }}</td>
              <td class="px-5 py-4 text-gray-500">{{ admin.adminUsername }}</td>
              <td class="px-5 py-4 text-gray-500">{{ admin.adminEmail }}</td>
              <td class="px-5 py-4">
                <span
                  class="inline-flex items-center gap-1 px-2.5 py-1 rounded-lg text-xs font-medium"
                  :class="
                    admin.adminRole === 'superadmin'
                      ? 'bg-purple-50 text-purple-600'
                      : 'bg-blue-50 text-blue-600'
                  "
                >
                  <Shield v-if="admin.adminRole === 'superadmin'" :size="10" />
                  {{ admin.adminRole === 'superadmin' ? 'Superadmin' : 'Admin' }}
                </span>
              </td>
              <td class="px-5 py-4">
                <span
                  class="inline-flex items-center gap-1 rounded-lg px-2.5 py-1 text-xs font-medium"
                  :class="
                    admin.isDeleted
                      ? 'bg-red-50 text-red-600'
                      : 'bg-emerald-50 text-emerald-600'
                  "
                >
                  {{ admin.isDeleted ? 'Dihapus' : 'Aktif' }}
                </span>
              </td>
              <td class="px-5 py-4">
                <div class="flex items-center justify-center gap-2">
                  <button
                    @click="openDetailDrawer(admin)"
                    class="flex items-center gap-1.5 px-3 py-1.5 hover:bg-green-100 text-green-600 text-xs font-medium rounded-lg transition-colors"
                  >
                    Detail
                  </button>
                  <button
                    v-if="!admin.isDeleted"
                    @click="openEditDrawer(admin)"
                    class="flex items-center gap-1.5 px-3 py-1.5 hover:bg-blue-100 text-blue-600 text-xs font-medium rounded-lg transition-colors"
                  >
                    Edit
                  </button>
                  <button
                    v-if="!admin.isDeleted"
                    @click="deleteAdmin(admin.adminId)"
                    class="flex items-center gap-1.5 px-3 py-1.5 hover:bg-red-100 text-red-600 text-xs font-medium rounded-lg transition-colors"
                  >
                    Hapus
                  </button>
                  <button
                    v-else
                    @click="restoreAdmin(admin.adminId)"
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
        v-if="adminStore.admins.length === 0"
        class="flex flex-col items-center justify-center py-16 text-gray-300"
      >
        <Users :size="40" class="mb-3" />
        <p class="text-sm">Belum ada data admin</p>
      </div>

      <div
        v-if="(adminStore.pagination?.total ?? 0) > 0"
        class="flex items-center justify-between border-t border-gray-100 px-5 py-4"
      >
        <p class="text-sm text-gray-500">
          Menampilkan {{ pageRange.startIndex + 1 }} - {{ pageRange.endIndex }} dari
          {{ adminStore.pagination?.total ?? adminStore.admins.length }} data
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
