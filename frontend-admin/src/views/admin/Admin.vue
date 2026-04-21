<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useAdminStore } from '@/stores/admin'
import { Plus, Pencil, Trash2, Eye, AlertCircle, Users, Shield } from '@lucide/vue'
import AdminCreateDrawer from './AdminCreateDrawer.vue'
import AdminEditDrawer from './AdminEditDrawer.vue'
import AdminDetailDrawer from './AdminDetailDrawer.vue'

const adminStore = useAdminStore()
const showCreateDrawer = ref(false)
const showEditDrawer = ref(false)
const showDetailDrawer = ref(false)

onMounted(async () => await adminStore.fetchAll())

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
  }
}

const handleSubmit = () => {
  adminStore.fetchAll()
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
        Tambah Admin
      </button>
    </div>

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
    <div v-else class="flex min-h-0 flex-1 flex-col rounded-2xl border border-gray-100 bg-white overflow-hidden">
      <div class="min-h-0 flex-1 overflow-auto">
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
              Username
            </th>
            <th
              class="px-5 py-3.5 text-left text-xs font-semibold text-gray-400 uppercase tracking-wide"
            >
              Email
            </th>
            <th
              class="px-5 py-3.5 text-left text-xs font-semibold text-gray-400 uppercase tracking-wide"
            >
              Role
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
              <div class="flex items-center justify-center gap-2">
                <button
                  @click="openDetailDrawer(admin)"
                  class="flex items-center gap-1.5 px-3 py-1.5 hover:bg-green-100 text-green-600 text-xs font-medium rounded-lg transition-colors"
                >
                  Detail
                </button>
                <button
                  @click="openEditDrawer(admin)"
                  class="flex items-center gap-1.5 px-3 py-1.5 hover:bg-blue-100 text-blue-600 text-xs font-medium rounded-lg transition-colors"
                >
                  Edit
                </button>
                <button
                  @click="deleteAdmin(admin.adminId)"
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
        v-if="adminStore.admins.length === 0"
        class="flex flex-col items-center justify-center py-16 text-gray-300"
      >
        <Users :size="40" class="mb-3" />
        <p class="text-sm">Belum ada data admin</p>
      </div>
    </div>
  </div>
</template>
