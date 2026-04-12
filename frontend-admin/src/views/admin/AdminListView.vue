<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useAdminStore } from '@/stores/admin'
import { Plus, Pencil, Trash2, X, AlertCircle, Users, Shield } from '@lucide/vue'

const adminStore = useAdminStore()
const showForm = ref(false)
const isEditing = ref(false)
const formData = ref({
  adminUsername: '',
  adminPassword: '',
  adminName: '',
  adminEmail: '',
  adminRole: 'admin' as 'admin' | 'superadmin',
})

onMounted(async () => await adminStore.fetchAll())

const resetForm = () => {
  formData.value = {
    adminUsername: '',
    adminPassword: '',
    adminName: '',
    adminEmail: '',
    adminRole: 'admin',
  }
  isEditing.value = false
  showForm.value = false
}

const handleSubmit = async () => {
  if (isEditing.value && adminStore.selectedAdmin) {
    await adminStore.update(adminStore.selectedAdmin.adminId, { ...formData.value })
  } else {
    await adminStore.create({ ...formData.value })
  }
  resetForm()
}

const editAdmin = (admin: any) => {
  adminStore.selectedAdmin = admin
  formData.value = {
    adminUsername: admin.adminUsername,
    adminPassword: '',
    adminName: admin.adminName,
    adminEmail: admin.adminEmail,
    adminRole: admin.adminRole,
  }
  isEditing.value = true
  showForm.value = true
}

const deleteAdmin = async (id: string) => {
  if (confirm('Yakin ingin menghapus admin ini?')) {
    await adminStore.deleteAdmin(id)
  }
}
</script>

<template>
  <div>
    <!-- Header -->
    <div class="flex items-center justify-between mb-6">
      <button
        @click="showForm = true"
        class="flex items-center gap-2 px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white text-sm font-medium rounded-xl transition-colors"
      >
        <Plus :size="16" />
        Tambah Admin
      </button>
    </div>

    <!-- Modal -->
    <Teleport to="body">
      <div
        v-if="showForm"
        class="fixed inset-0 bg-black/50 z-50 flex items-center justify-center p-4"
        @click.self="resetForm"
      >
        <div class="bg-white rounded-2xl shadow-xl w-full max-w-md">
          <div class="flex items-center justify-between px-6 py-4 border-b border-gray-100">
            <h2 class="text-base font-semibold text-gray-900">
              {{ isEditing ? 'Edit Admin' : 'Tambah Admin' }}
            </h2>
            <button
              @click="resetForm"
              class="p-1.5 rounded-lg hover:bg-gray-100 text-gray-400 transition-colors"
            >
              <X :size="16" />
            </button>
          </div>

          <form @submit.prevent="handleSubmit" class="px-6 py-5 space-y-4">
            <div
              v-for="field in [
                { key: 'adminUsername', label: 'Username', type: 'text' },
                { key: 'adminName', label: 'Nama', type: 'text' },
                { key: 'adminEmail', label: 'Email', type: 'email' },
              ]"
              :key="field.key"
            >
              <label
                class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-1.5"
                >{{ field.label }}</label
              >
              <input
                v-model="(formData as any)[field.key]"
                :type="field.type"
                required
                class="w-full px-3.5 py-2.5 text-sm text-gray-900 bg-gray-50 border border-gray-200 rounded-xl outline-none transition-all focus:bg-white focus:border-blue-400 focus:ring-2 focus:ring-blue-100"
              />
            </div>

            <div>
              <label class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-1.5">
                Password
                <span v-if="isEditing" class="normal-case text-gray-300 ml-1"
                  >(kosongkan jika tidak diubah)</span
                >
              </label>
              <input
                v-model="formData.adminPassword"
                type="password"
                :required="!isEditing"
                class="w-full px-3.5 py-2.5 text-sm text-gray-900 bg-gray-50 border border-gray-200 rounded-xl outline-none transition-all focus:bg-white focus:border-blue-400 focus:ring-2 focus:ring-blue-100"
              />
            </div>

            <div>
              <label class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-1.5"
                >Role</label
              >
              <select
                v-model="formData.adminRole"
                class="w-full px-3.5 py-2.5 text-sm text-gray-900 bg-gray-50 border border-gray-200 rounded-xl outline-none transition-all focus:bg-white focus:border-blue-400 focus:ring-2 focus:ring-blue-100"
              >
                <option value="admin">Admin</option>
                <option value="superadmin">Superadmin</option>
              </select>
            </div>

            <div
              v-if="adminStore.error"
              class="flex items-center gap-2 p-3 bg-red-50 border border-red-100 text-red-600 rounded-xl text-xs"
            >
              <AlertCircle :size="14" class="shrink-0" />
              {{ adminStore.error }}
            </div>

            <div class="flex gap-3 pt-1">
              <button
                type="submit"
                :disabled="adminStore.isLoading"
                class="flex-1 py-2.5 bg-blue-600 hover:bg-blue-700 disabled:opacity-50 text-white text-sm font-medium rounded-xl transition-colors"
              >
                {{ adminStore.isLoading ? 'Menyimpan...' : 'Simpan' }}
              </button>
              <button
                type="button"
                @click="resetForm"
                class="flex-1 py-2.5 bg-gray-100 hover:bg-gray-200 text-gray-600 text-sm font-medium rounded-xl transition-colors"
              >
                Batal
              </button>
            </div>
          </form>
        </div>
      </div>
    </Teleport>

    <!-- Loading -->
    <div
      v-if="adminStore.isLoading"
      class="flex items-center justify-center py-16 text-sm text-gray-400"
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
    <div v-else class="bg-white rounded-2xl border border-gray-100 overflow-hidden">
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
                  @click="editAdmin(admin)"
                  class="flex items-center gap-1.5 px-3 py-1.5 bg-blue-50 hover:bg-blue-100 text-blue-600 text-xs font-medium rounded-lg transition-colors"
                >
                  <Pencil :size="12" /> Edit
                </button>
                <button
                  @click="deleteAdmin(admin.adminId)"
                  class="flex items-center gap-1.5 px-3 py-1.5 bg-red-50 hover:bg-red-100 text-red-600 text-xs font-medium rounded-lg transition-colors"
                >
                  <Trash2 :size="12" /> Hapus
                </button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>

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
