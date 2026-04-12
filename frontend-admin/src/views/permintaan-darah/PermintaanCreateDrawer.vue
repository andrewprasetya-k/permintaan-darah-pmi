<script setup lang="ts">
import { ref } from 'vue'
import { usePermintaanStore } from '@/stores/permintaan'
import type { CreatePermintaanRequest } from '@/api/permintaan'
import { X, AlertCircle } from '@lucide/vue'

defineProps<{ isOpen: boolean }>()

const emit = defineEmits<{
  close: []
  submit: []
}>()

const permintaanStore = usePermintaanStore()
const isSubmitting = ref(false)
const formData = ref<CreatePermintaanRequest>({
  namaPasien: '',
  tempatLahir: '',
  tanggalLahir: '',
  jenisKelamin: 'L',
  pernahTransfusi: false,
  status: 'dibuat',
  tanggalPermintaan: new Date().toISOString().split('T')[0],
} as CreatePermintaanRequest)

const resetForm = () => {
  formData.value = {
    namaPasien: '',
    tempatLahir: '',
    tanggalLahir: '',
    jenisKelamin: 'L',
    pernahTransfusi: false,
    status: 'dibuat',
    tanggalPermintaan: new Date().toISOString().split('T')[0],
  } as CreatePermintaanRequest
}

const handleClose = () => {
  resetForm()
  emit('close')
}

const handleSubmit = async () => {
  isSubmitting.value = true
  try {
    await permintaanStore.create(formData.value)
    emit('submit')
    resetForm()
    emit('close')
  } finally {
    isSubmitting.value = false
  }
}
</script>

<template>
  <Teleport to="body">
    <Transition name="backdrop"
      >{ "success": true, "message": "Data retrieved successfully", "data": { "permintaanDarahId":
      "PD0204260309002", "rumahSakitId": "RS-021", "namaPasien": "Budi awikwok", "noRM":
      "RM-992831", "tempatLahir": "Salatiga", "tanggalLahir": "1995-05-20T00:00:00Z",
      "jenisKelamin": "L", "golonganDarah": "A", "rhesusDarah": "+", "hemoglobin": 11.5,
      "ruangBagianKelas": "Bangsal Melati - Kelas 2", "pernahTransfusi": false, "indikasiTransfusi":
      "Anemia berat", "pernahHamil": "Tidak", "status": "dibuat", "tanggalPermintaan":
      "2026-06-02T10:00:00Z", "createdAt": "2026-04-02T03:09:47.447422Z", "updatedAt":
      "2026-04-08T07:49:07.738238Z", "detailPermintaanDarah": [ { "detailId": 3,
      "permintaanDarahId": "PD0204260309002", "komponenNama": "Packed Red Cell (PRC)",
      "golonganDarah": "O", "rhesusDarah": "+", "jumlahKantong": 2, "tanggalDiperlukan":
      "2026-04-03T08:00:00Z", "createdAt": "2026-04-02T08:55:15.337021Z" }, { "detailId": 4,
      "permintaanDarahId": "PD0204260309002", "komponenNama": "Fresh Frozen Plasma (FFP)",
      "golonganDarah": "O", "rhesusDarah": "+", "jumlahKantong": 2, "tanggalDiperlukan":
      "2026-04-03T08:00:00Z", "createdAt": "2026-04-02T09:00:56.538337Z" }, { "detailId": 5,
      "permintaanDarahId": "PD0204260309002", "komponenNama": "Fresh Frozen Plasma (FFP)",
      "golonganDarah": "O", "rhesusDarah": "+", "jumlahKantong": 2, "tanggalDiperlukan":
      "2026-04-03T08:00:00Z", "createdAt": "2026-04-02T09:01:41.694244Z" }, { "detailId": 6,
      "permintaanDarahId": "PD0204260309002", "komponenNama": "Whole Blood", "golonganDarah": "O",
      "rhesusDarah": "+", "jumlahKantong": 2, "tanggalDiperlukan": "2026-04-03T08:00:00Z",
      "createdAt": "2026-04-02T09:01:42.500653Z" } ] } }
      <div
        v-if="isOpen"
        class="fixed inset-0 bg-black/20 backdrop-blur-sm z-40"
        @click="handleClose"
      />
    </Transition>

    <Transition name="drawer">
      <div
        v-if="isOpen"
        class="fixed top-0 right-0 h-full bg-white shadow-2xl z-50 flex flex-col w-full lg:w-4/5 lg:rounded-tl-3xl lg:rounded-bl-3xl"
      >
        <!-- Header -->
        <div class="flex items-center justify-between px-10 py-8 border-b border-gray-200">
          <div>
            <h2 class="text-xl font-semibold text-gray-900">Tambah Permintaan Darah</h2>
            <p class="text-sm text-gray-400 mt-0.5">Buat permintaan darah baru ke sistem</p>
          </div>
          <button
            @click="handleClose"
            class="p-2 rounded-lg hover:bg-gray-100 text-gray-400 transition-colors"
          >
            <X :size="20" />
          </button>
        </div>

        <!-- Content -->
        <div class="flex-1 overflow-y-auto px-10 py-8">
          <form @submit.prevent="handleSubmit" class="space-y-5 max-w-full">
            <!-- Row 1: Nama Pasien & No. RM -->
            <div class="grid grid-cols-2 gap-4">
              <div>
                <label
                  class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-1.5"
                >
                  Nama Pasien
                </label>
                <input
                  v-model="formData.namaPasien"
                  type="text"
                  required
                  class="w-full px-3.5 py-2.5 text-sm text-gray-900 bg-gray-50 border border-gray-200 rounded-xl outline-none transition-all focus:bg-white focus:border-blue-400 focus:ring-2 focus:ring-blue-100"
                />
              </div>
              <div>
                <label
                  class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-1.5"
                >
                  No. Rekam Medis
                </label>
                <input
                  v-model="formData.noRM"
                  type="text"
                  required
                  class="w-full px-3.5 py-2.5 text-sm text-gray-900 bg-gray-50 border border-gray-200 rounded-xl outline-none transition-all focus:bg-white focus:border-blue-400 focus:ring-2 focus:ring-blue-100"
                />
              </div>
            </div>

            <!-- Row 2: Tempat Lahir & Tanggal Lahir -->
            <div class="grid grid-cols-2 gap-4">
              <div>
                <label
                  class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-1.5"
                >
                  Tempat Lahir
                </label>
                <input
                  v-model="formData.tempatLahir"
                  type="text"
                  required
                  class="w-full px-3.5 py-2.5 text-sm text-gray-900 bg-gray-50 border border-gray-200 rounded-xl outline-none transition-all focus:bg-white focus:border-blue-400 focus:ring-2 focus:ring-blue-100"
                />
              </div>
              <div>
                <label
                  class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-1.5"
                >
                  Tanggal Lahir
                </label>
                <input
                  v-model="formData.tanggalLahir"
                  type="date"
                  required
                  class="w-full px-3.5 py-2.5 text-sm text-gray-900 bg-gray-50 border border-gray-200 rounded-xl outline-none transition-all focus:bg-white focus:border-blue-400 focus:ring-2 focus:ring-blue-100"
                />
              </div>
            </div>

            <!-- Row 3: Jenis Kelamin & Golongan Darah -->
            <div class="grid grid-cols-2 gap-4">
              <div>
                <label
                  class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-1.5"
                >
                  Jenis Kelamin
                </label>
                <select
                  v-model="formData.jenisKelamin"
                  class="w-full px-3.5 py-2.5 text-sm text-gray-900 bg-gray-50 border border-gray-200 rounded-xl outline-none transition-all focus:bg-white focus:border-blue-400 focus:ring-2 focus:ring-blue-100"
                >
                  <option value="L">Laki-laki</option>
                  <option value="P">Perempuan</option>
                </select>
              </div>
              <div>
                <label
                  class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-1.5"
                >
                  Golongan Darah
                </label>
                <select
                  v-model="formData.golonganDarah"
                  class="w-full px-3.5 py-2.5 text-sm text-gray-900 bg-gray-50 border border-gray-200 rounded-xl outline-none transition-all focus:bg-white focus:border-blue-400 focus:ring-2 focus:ring-blue-100"
                >
                  <option value="A">A</option>
                  <option value="B">B</option>
                  <option value="AB">AB</option>
                  <option value="O">O</option>
                </select>
              </div>
            </div>

            <!-- Row 4: Rhesus & Hemoglobin -->
            <div class="grid grid-cols-2 gap-4">
              <div>
                <label
                  class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-1.5"
                >
                  Rhesus Darah
                </label>
                <select
                  v-model="formData.rhesusDarah"
                  class="w-full px-3.5 py-2.5 text-sm text-gray-900 bg-gray-50 border border-gray-200 rounded-xl outline-none transition-all focus:bg-white focus:border-blue-400 focus:ring-2 focus:ring-blue-100"
                >
                  <option value="+">Positif (+)</option>
                  <option value="-">Negatif (-)</option>
                </select>
              </div>
              <div>
                <label
                  class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-1.5"
                >
                  Hemoglobin
                </label>
                <input
                  v-model.number="formData.hemoglobin"
                  type="number"
                  step="0.1"
                  required
                  class="w-full px-3.5 py-2.5 text-sm text-gray-900 bg-gray-50 border border-gray-200 rounded-xl outline-none transition-all focus:bg-white focus:border-blue-400 focus:ring-2 focus:ring-blue-100"
                />
              </div>
            </div>

            <!-- Ruang/Bagian/Kelas -->
            <div>
              <label class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-1.5">
                Ruang/Bagian/Kelas
              </label>
              <input
                v-model="formData.ruangBagianKelas"
                type="text"
                class="w-full px-3.5 py-2.5 text-sm text-gray-900 bg-gray-50 border border-gray-200 rounded-xl outline-none transition-all focus:bg-white focus:border-blue-400 focus:ring-2 focus:ring-blue-100"
              />
            </div>

            <!-- Pernah Transfusi -->
            <div class="flex items-center gap-2">
              <input
                v-model="formData.pernahTransfusi"
                type="checkbox"
                id="pernah-transfusi"
                class="rounded border-gray-200 text-blue-600"
              />
              <label for="pernah-transfusi" class="text-sm text-gray-600 cursor-pointer">
                Pernah melakukan transfusi sebelumnya
              </label>
            </div>

            <!-- Indikasi Transfusi -->
            <div>
              <label class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-1.5">
                Indikasi Transfusi
              </label>
              <textarea
                v-model="formData.indikasiTransfusi"
                required
                rows="3"
                class="w-full px-3.5 py-2.5 text-sm text-gray-900 bg-gray-50 border border-gray-200 rounded-xl outline-none transition-all focus:bg-white focus:border-blue-400 focus:ring-2 focus:ring-blue-100"
              />
            </div>

            <!-- Error -->
            <div
              v-if="permintaanStore.error"
              class="flex items-center gap-2 p-3 bg-red-50 border border-red-100 text-red-600 rounded-xl text-xs"
            >
              <AlertCircle :size="14" class="shrink-0" />
              {{ permintaanStore.error }}
            </div>

            <!-- Actions -->
            <div class="flex gap-3 pt-2">
              <button
                type="submit"
                :disabled="isSubmitting"
                class="flex-1 py-2.5 bg-blue-600 hover:bg-blue-700 disabled:opacity-50 text-white text-sm font-medium rounded-xl transition-colors"
              >
                {{ isSubmitting ? 'Menyimpan...' : 'Simpan' }}
              </button>
              <button
                type="button"
                @click="handleClose"
                class="flex-1 py-2.5 bg-gray-100 hover:bg-gray-200 text-gray-600 text-sm font-medium rounded-xl transition-colors"
              >
                Batal
              </button>
            </div>
          </form>
        </div>
      </div>
    </Transition>
  </Teleport>
</template>

<style scoped>
.backdrop-enter-active,
.backdrop-leave-active {
  transition: opacity 0.3s ease;
}
.backdrop-enter-from,
.backdrop-leave-to {
  opacity: 0;
}

.drawer-enter-active,
.drawer-leave-active {
  transition: transform 0.3s ease;
}
.drawer-enter-from,
.drawer-leave-to {
  transform: translateX(100%);
}
</style>
