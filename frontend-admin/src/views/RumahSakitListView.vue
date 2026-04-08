<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRumahSakitStore } from '@/stores/rumah-sakit'

const rumahSakitStore = useRumahSakitStore()
const showForm = ref(false)
const isEditing = ref(false)
const formData = ref({
  nama: '',
  nomorTelepon: '',
  alamat: '',
  email: '',
  username: '',
  password: '',
})

onMounted(async () => {
  await rumahSakitStore.fetchAll()
})

const resetForm = () => {
  formData.value = {
    nama: '',
    nomorTelepon: '',
    alamat: '',
    email: '',
    username: '',
    password: '',
  }
  isEditing.value = false
  showForm.value = false
}

const handleSubmit = async () => {
  if (isEditing.value && rumahSakitStore.selectedHospital) {
    await rumahSakitStore.update(rumahSakitStore.selectedHospital.rumahSakitId, {
      nama: formData.value.nama,
      nomorTelepon: formData.value.nomorTelepon,
      alamat: formData.value.alamat,
      email: formData.value.email || undefined,
      password: formData.value.password || undefined,
    })
  } else {
    await rumahSakitStore.create({
      nama: formData.value.nama,
      nomorTelepon: formData.value.nomorTelepon,
      alamat: formData.value.alamat,
      email: formData.value.email || undefined,
      username: formData.value.username,
      password: formData.value.password,
    })
  }
  resetForm()
}

const editHospital = (hospital: any) => {
  rumahSakitStore.selectedHospital = hospital
  formData.value = {
    nama: hospital.nama,
    nomorTelepon: hospital.nomorTelepon,
    alamat: hospital.alamat,
    email: hospital.email || '',
    username: hospital.username,
    password: '',
  }
  isEditing.value = true
  showForm.value = true
}

const deleteHospital = async (id: string) => {
  if (confirm('Yakin ingin menghapus rumah sakit ini?')) {
    await rumahSakitStore.deleteHospital(id)
  }
}
</script>

<template>
  <div style="padding: 1.5rem;">
    <div style="display: flex; justify-content: space-between; align-items: center; margin-bottom: 2rem;">
      <h1 style="font-size: 1.875rem; font-weight: bold;">Manajemen Rumah Sakit</h1>
      <button
        @click="showForm = true"
        style="padding: 0.5rem 1rem; background-color: #059669; color: white; border-radius: 0.375rem; border: none; cursor: pointer; font-weight: 500;"
      >
        + Tambah Rumah Sakit
      </button>
    </div>

    <!-- Form Modal -->
    <div v-if="showForm" style="position: fixed; top: 0; left: 0; right: 0; bottom: 0; background-color: rgba(0,0,0,0.5); display: flex; align-items: center; justify-content: center;">
      <div style="background-color: white; padding: 2rem; border-radius: 0.5rem; width: 90%; max-width: 28rem;">
        <h2 style="font-size: 1.25rem; font-weight: bold; margin-bottom: 1rem;">
          {{ isEditing ? 'Edit Rumah Sakit' : 'Tambah Rumah Sakit' }}
        </h2>

        <form @submit.prevent="handleSubmit" style="display: flex; flex-direction: column; gap: 1rem;">
          <div>
            <label style="display: block; font-size: 0.875rem; font-weight: 500; margin-bottom: 0.25rem;">Nama</label>
            <input
              v-model="formData.nama"
              type="text"
              required
              style="width: 100%; padding: 0.5rem 0.75rem; border: 1px solid #d1d5db; border-radius: 0.375rem;"
            />
          </div>

          <div>
            <label style="display: block; font-size: 0.875rem; font-weight: 500; margin-bottom: 0.25rem;">Nomor Telepon</label>
            <input
              v-model="formData.nomorTelepon"
              type="text"
              required
              style="width: 100%; padding: 0.5rem 0.75rem; border: 1px solid #d1d5db; border-radius: 0.375rem;"
            />
          </div>

          <div>
            <label style="display: block; font-size: 0.875rem; font-weight: 500; margin-bottom: 0.25rem;">Alamat</label>
            <textarea
              v-model="formData.alamat"
              required
              style="width: 100%; padding: 0.5rem 0.75rem; border: 1px solid #d1d5db; border-radius: 0.375rem; min-height: 80px;"
            />
          </div>

          <div>
            <label style="display: block; font-size: 0.875rem; font-weight: 500; margin-bottom: 0.25rem;">Email</label>
            <input
              v-model="formData.email"
              type="email"
              style="width: 100%; padding: 0.5rem 0.75rem; border: 1px solid #d1d5db; border-radius: 0.375rem;"
            />
          </div>

          <div v-if="!isEditing">
            <label style="display: block; font-size: 0.875rem; font-weight: 500; margin-bottom: 0.25rem;">Username</label>
            <input
              v-model="formData.username"
              type="text"
              required
              style="width: 100%; padding: 0.5rem 0.75rem; border: 1px solid #d1d5db; border-radius: 0.375rem;"
            />
          </div>

          <div>
            <label style="display: block; font-size: 0.875rem; font-weight: 500; margin-bottom: 0.25rem;">
              Password {{ isEditing ? '(kosongkan jika tidak ingin mengubah)' : '' }}
            </label>
            <input
              v-model="formData.password"
              type="password"
              :required="!isEditing"
              style="width: 100%; padding: 0.5rem 0.75rem; border: 1px solid #d1d5db; border-radius: 0.375rem;"
            />
          </div>

          <div v-if="rumahSakitStore.error" style="padding: 0.75rem; background-color: #fee2e2; color: #991b1b; border-radius: 0.375rem;">
            {{ rumahSakitStore.error }}
          </div>

          <div style="display: flex; gap: 1rem;">
            <button
              type="submit"
              :disabled="rumahSakitStore.isLoading"
              style="flex: 1; padding: 0.5rem; background-color: #2563eb; color: white; border-radius: 0.375rem; border: none; cursor: pointer; font-weight: 500;"
            >
              {{ rumahSakitStore.isLoading ? 'Menyimpan...' : 'Simpan' }}
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
    <div v-if="rumahSakitStore.isLoading" style="padding: 2rem; text-align: center;">Loading...</div>
    <div v-else-if="rumahSakitStore.error" style="padding: 1rem; background-color: #fee2e2; color: #991b1b; border-radius: 0.375rem;">
      {{ rumahSakitStore.error }}
    </div>
    <div v-else style="background-color: white; border-radius: 0.5rem; overflow: hidden; box-shadow: 0 1px 3px 0 rgba(0, 0, 0, 0.1);">
      <table style="width: 100%; border-collapse: collapse;">
        <thead style="background-color: #f3f4f6;">
          <tr>
            <th style="padding: 0.75rem; text-align: left; border-bottom: 1px solid #e5e7eb;">Nama</th>
            <th style="padding: 0.75rem; text-align: left; border-bottom: 1px solid #e5e7eb;">Telepon</th>
            <th style="padding: 0.75rem; text-align: left; border-bottom: 1px solid #e5e7eb;">Email</th>
            <th style="padding: 0.75rem; text-align: left; border-bottom: 1px solid #e5e7eb;">Username</th>
            <th style="padding: 0.75rem; text-align: center; border-bottom: 1px solid #e5e7eb;">Aksi</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="rs in rumahSakitStore.hospitals" :key="rs.rumahSakitId" style="border-bottom: 1px solid #e5e7eb;">
            <td style="padding: 0.75rem;">{{ rs.nama }}</td>
            <td style="padding: 0.75rem;">{{ rs.nomorTelepon }}</td>
            <td style="padding: 0.75rem;">{{ rs.email || '-' }}</td>
            <td style="padding: 0.75rem;">{{ rs.username }}</td>
            <td style="padding: 0.75rem; text-align: center;">
              <button
                @click="editHospital(rs)"
                style="padding: 0.25rem 0.75rem; background-color: #3b82f6; color: white; border-radius: 0.25rem; border: none; cursor: pointer; font-size: 0.875rem; margin-right: 0.5rem;"
              >
                Edit
              </button>
              <button
                @click="deleteHospital(rs.rumahSakitId)"
                style="padding: 0.25rem 0.75rem; background-color: #ef4444; color: white; border-radius: 0.25rem; border: none; cursor: pointer; font-size: 0.875rem;"
              >
                Hapus
              </button>
            </td>
          </tr>
        </tbody>
      </table>
      <div v-if="rumahSakitStore.hospitals.length === 0" style="padding: 2rem; text-align: center; color: #6b7280;">
        Belum ada data rumah sakit
      </div>
    </div>
  </div>
</template>
