<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRumahSakitStore } from '@/stores/rumah-sakit'
import { Plus, Pencil, Trash2, X, AlertCircle, Hospital } from '@lucide/vue'

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

onMounted(async () => await rumahSakitStore.fetchAll())

const resetForm = () => {
  formData.value = { nama: '', nomorTelepon: '', alamat: '', email: '', username: '', password: '' }
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
  <div>
    <!-- Header -->
    <div class="flex items-center justify-between mb-6">
      <button
        @click="showForm = true"
        class="flex items-center gap-2 px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white text-sm font-medium rounded-xl transition-colors"
      >
        <Plus :size="16" />
        Tambah Rumah Sakit
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
              {{ isEditing ? 'Edit Rumah Sakit' : 'Tambah Rumah Sakit' }}
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
                { key: 'nama', label: 'Nama', type: 'text', required: true },
                { key: 'nomorTelepon', label: 'Nomor Telepon', type: 'text', required: true },
                { key: 'email', label: 'Email', type: 'email', required: false },
              ]"
              :key="field.key"
            >
              <label class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-1.5">
                {{ field.label }}
              </label>
              <input
                v-model="(formData as any)[field.key]"
                :type="field.type"
                :required="field.required"
                class="w-full px-3.5 py-2.5 text-sm text-gray-900 bg-gray-50 border border-gray-200 rounded-xl outline-none transition-all focus:bg-white focus:border-blue-400 focus:ring-2 focus:ring-blue-100"
              />
            </div>

            <div>
              <label class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-1.5"
                >Alamat</label
              >
              <textarea
                v-model="formData.alamat"
                required
                rows="3"
                class="w-full px-3.5 py-2.5 text-sm text-gray-900 bg-gray-50 border border-gray-200 rounded-xl outline-none transition-all focus:bg-white focus:border-blue-400 focus:ring-2 focus:ring-blue-100 resize-none"
              />
            </div>

            <div v-if="!isEditing">
              <label class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-1.5"
                >Username</label
              >
              <input
                v-model="formData.username"
                type="text"
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
                v-model="formData.password"
                type="password"
                :required="!isEditing"
                class="w-full px-3.5 py-2.5 text-sm text-gray-900 bg-gray-50 border border-gray-200 rounded-xl outline-none transition-all focus:bg-white focus:border-blue-400 focus:ring-2 focus:ring-blue-100"
              />
            </div>

            <div
              v-if="rumahSakitStore.error"
              class="flex items-center gap-2 p-3 bg-red-50 border border-red-100 text-red-600 rounded-xl text-xs"
            >
              <AlertCircle :size="14" class="shrink-0" />
              {{ rumahSakitStore.error }}
            </div>

            <div class="flex gap-3 pt-1">
              <button
                type="submit"
                :disabled="rumahSakitStore.isLoading"
                class="flex-1 py-2.5 bg-blue-600 hover:bg-blue-700 disabled:opacity-50 text-white text-sm font-medium rounded-xl transition-colors"
              >
                {{ rumahSakitStore.isLoading ? 'Menyimpan...' : 'Simpan' }}
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
      v-if="rumahSakitStore.isLoading"
      class="flex items-center justify-center py-16 text-sm text-gray-400"
    >
      <span
        class="w-5 h-5 border-2 border-gray-200 border-t-blue-500 rounded-full animate-spin mr-3"
      />
      Memuat data...
    </div>

    <!-- Error -->
    <div
      v-else-if="rumahSakitStore.error"
      class="flex items-center gap-2 p-4 bg-red-50 border border-red-100 text-red-600 rounded-xl text-sm"
    >
      <AlertCircle :size="16" />
      {{ rumahSakitStore.error }}
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
                  @click="editHospital(rs)"
                  class="flex items-center gap-1.5 px-3 py-1.5 bg-blue-50 hover:bg-blue-100 text-blue-600 text-xs font-medium rounded-lg transition-colors"
                >
                  <Pencil :size="12" /> Edit
                </button>
                <button
                  @click="deleteHospital(rs.rumahSakitId)"
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
        v-if="rumahSakitStore.hospitals.length === 0"
        class="flex flex-col items-center justify-center py-16 text-gray-300"
      >
        <Hospital :size="40" class="mb-3" />
        <p class="text-sm">Belum ada data rumah sakit</p>
      </div>
    </div>
  </div>
</template>
