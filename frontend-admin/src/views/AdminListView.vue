<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useAdminStore } from '@/stores/admin'

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

onMounted(async () => {
  await adminStore.fetchAll()
})

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
    await adminStore.update(adminStore.selectedAdmin.adminId, {
      adminUsername: formData.value.adminUsername,
      adminPassword: formData.value.adminPassword,
      adminName: formData.value.adminName,
      adminEmail: formData.value.adminEmail,
      adminRole: formData.value.adminRole,
    })
  } else {
    await adminStore.create({
      adminUsername: formData.value.adminUsername,
      adminPassword: formData.value.adminPassword,
      adminName: formData.value.adminName,
      adminEmail: formData.value.adminEmail,
      adminRole: formData.value.adminRole,
    })
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
  <div style="padding: 1.5rem;">
    <div style="display: flex; justify-content: space-between; align-items: center; margin-bottom: 2rem;">
      <h1 style="font-size: 1.875rem; font-weight: bold;">Manajemen Admin</h1>
      <button
        @click="showForm = true"
        style="padding: 0.5rem 1rem; background-color: #059669; color: white; border-radius: 0.375rem; border: none; cursor: pointer; font-weight: 500;"
      >
        + Tambah Admin
      </button>
    </div>

    <!-- Form Modal -->
    <div v-if="showForm" style="position: fixed; top: 0; left: 0; right: 0; bottom: 0; background-color: rgba(0,0,0,0.5); display: flex; align-items: center; justify-content: center;">
      <div style="background-color: white; padding: 2rem; border-radius: 0.5rem; width: 90%; max-width: 28rem;">
        <h2 style="font-size: 1.25rem; font-weight: bold; margin-bottom: 1rem;">
          {{ isEditing ? 'Edit Admin' : 'Tambah Admin' }}
        </h2>

        <form @submit.prevent="handleSubmit" style="display: flex; flex-direction: column; gap: 1rem;">
          <div>
            <label style="display: block; font-size: 0.875rem; font-weight: 500; margin-bottom: 0.25rem;">Username</label>
            <input
              v-model="formData.adminUsername"
              type="text"
              required
              style="width: 100%; padding: 0.5rem 0.75rem; border: 1px solid #d1d5db; border-radius: 0.375rem;"
            />
          </div>

          <div>
            <label style="display: block; font-size: 0.875rem; font-weight: 500; margin-bottom: 0.25rem;">Nama</label>
            <input
              v-model="formData.adminName"
              type="text"
              required
              style="width: 100%; padding: 0.5rem 0.75rem; border: 1px solid #d1d5db; border-radius: 0.375rem;"
            />
          </div>

          <div>
            <label style="display: block; font-size: 0.875rem; font-weight: 500; margin-bottom: 0.25rem;">Email</label>
            <input
              v-model="formData.adminEmail"
              type="email"
              required
              style="width: 100%; padding: 0.5rem 0.75rem; border: 1px solid #d1d5db; border-radius: 0.375rem;"
            />
          </div>

          <div>
            <label style="display: block; font-size: 0.875rem; font-weight: 500; margin-bottom: 0.25rem;">
              Password {{ isEditing ? '(kosongkan jika tidak ingin mengubah)' : '' }}
            </label>
            <input
              v-model="formData.adminPassword"
              type="password"
              :required="!isEditing"
              style="width: 100%; padding: 0.5rem 0.75rem; border: 1px solid #d1d5db; border-radius: 0.375rem;"
            />
          </div>

          <div>
            <label style="display: block; font-size: 0.875rem; font-weight: 500; margin-bottom: 0.25rem;">Role</label>
            <select
              v-model="formData.adminRole"
              style="width: 100%; padding: 0.5rem 0.75rem; border: 1px solid #d1d5db; border-radius: 0.375rem;"
            >
              <option value="admin">Admin</option>
              <option value="superadmin">Superadmin</option>
            </select>
          </div>

          <div v-if="adminStore.error" style="padding: 0.75rem; background-color: #fee2e2; color: #991b1b; border-radius: 0.375rem;">
            {{ adminStore.error }}
          </div>

          <div style="display: flex; gap: 1rem;">
            <button
              type="submit"
              :disabled="adminStore.isLoading"
              style="flex: 1; padding: 0.5rem; background-color: #2563eb; color: white; border-radius: 0.375rem; border: none; cursor: pointer; font-weight: 500;"
            >
              {{ adminStore.isLoading ? 'Menyimpan...' : 'Simpan' }}
            </button>
            <button
              type="button"
              @click="resetForm"
              style="flex: 1; padding: 0.5rem; background-color: #6b7280; color: white; border-radius: 0.375rem; border: none; cursor: pointer; font-weight: 500;"
            >
              Batal
            </button>
          </div>
        </form>
      </div>
    </div>

    <!-- List -->
    <div v-if="adminStore.isLoading" style="padding: 2rem; text-align: center;">Loading...</div>
    <div v-else-if="adminStore.error" style="padding: 1rem; background-color: #fee2e2; color: #991b1b; border-radius: 0.375rem;">
      {{ adminStore.error }}
    </div>
    <div v-else style="background-color: white; border-radius: 0.5rem; overflow: hidden; box-shadow: 0 1px 3px 0 rgba(0, 0, 0, 0.1);">
      <table style="width: 100%; border-collapse: collapse;">
        <thead style="background-color: #f3f4f6;">
          <tr>
            <th style="padding: 0.75rem; text-align: left; border-bottom: 1px solid #e5e7eb;">Nama</th>
            <th style="padding: 0.75rem; text-align: left; border-bottom: 1px solid #e5e7eb;">Username</th>
            <th style="padding: 0.75rem; text-align: left; border-bottom: 1px solid #e5e7eb;">Email</th>
            <th style="padding: 0.75rem; text-align: left; border-bottom: 1px solid #e5e7eb;">Role</th>
            <th style="padding: 0.75rem; text-align: center; border-bottom: 1px solid #e5e7eb;">Aksi</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="admin in adminStore.admins" :key="admin.adminId" style="border-bottom: 1px solid #e5e7eb;">
            <td style="padding: 0.75rem;">{{ admin.adminName }}</td>
            <td style="padding: 0.75rem;">{{ admin.adminUsername }}</td>
            <td style="padding: 0.75rem;">{{ admin.adminEmail }}</td>
            <td style="padding: 0.75rem;">{{ admin.adminRole }}</td>
            <td style="padding: 0.75rem; text-align: center;">
              <button
                @click="editAdmin(admin)"
                style="padding: 0.25rem 0.75rem; background-color: #3b82f6; color: white; border-radius: 0.25rem; border: none; cursor: pointer; font-size: 0.875rem; margin-right: 0.5rem;"
              >
                Edit
              </button>
              <button
                @click="deleteAdmin(admin.adminId)"
                style="padding: 0.25rem 0.75rem; background-color: #ef4444; color: white; border-radius: 0.25rem; border: none; cursor: pointer; font-size: 0.875rem;"
              >
                Hapus
              </button>
            </td>
          </tr>
        </tbody>
      </table>
      <div v-if="adminStore.admins.length === 0" style="padding: 2rem; text-align: center; color: #6b7280;">
        Belum ada data admin
      </div>
    </div>
  </div>
</template>
