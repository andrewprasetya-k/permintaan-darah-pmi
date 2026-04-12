<script setup lang="ts">
import { computed, ref, watch } from 'vue'
import { usePermintaanStore } from '@/stores/permintaan'
import type { UpdatePermintaanRequest } from '@/api/permintaan'
import { X, AlertCircle } from '@lucide/vue'

const props = defineProps<{
  isOpen: boolean
}>()

const emit = defineEmits<{
  close: []
  submit: []
}>()

const permintaanStore = usePermintaanStore()
const isSubmitting = ref(false)
const formData = ref<UpdatePermintaanRequest>({
  rumahSakitId: '',
  namaPasien: '',
  tempatLahir: '',
  tanggalLahir: '',
  jenisKelamin: 'L',
  pernahTransfusi: false,
  status: 'dibuat',
  tanggalPermintaan: new Date().toISOString().split('T')[0],
} as UpdatePermintaanRequest)

const subtitle = computed(() => 'Perbarui informasi permintaan darah')

watch(
  () => props.isOpen,
  (newVal) => {
    if (newVal && permintaanStore.selectedRequest) {
      formData.value = {
        rumahSakitId: permintaanStore.selectedRequest.rumahSakitId,
        namaPasien: permintaanStore.selectedRequest.namaPasien,
        noRM: permintaanStore.selectedRequest.noRM || '',
        tempatLahir: permintaanStore.selectedRequest.tempatLahir,
        tanggalLahir: permintaanStore.selectedRequest.tanggalLahir,
        jenisKelamin: permintaanStore.selectedRequest.jenisKelamin,
        golonganDarah: permintaanStore.selectedRequest.golonganDarah || 'A',
        rhesusDarah: permintaanStore.selectedRequest.rhesusDarah || '+',
        hemoglobin: permintaanStore.selectedRequest.hemoglobin || 0,
        ruangBagianKelas: permintaanStore.selectedRequest.ruangBagianKelas || '',
        pernahTransfusi: permintaanStore.selectedRequest.pernahTransfusi,
        indikasiTransfusi: permintaanStore.selectedRequest.indikasiTransfusi || '',
        pernahHamil: permintaanStore.selectedRequest.pernahHamil || '',
        status: permintaanStore.selectedRequest.status as 'dibuat' | 'diproses' | 'bisa_diambil' | 'selesai' | 'dibatalkan',
        tanggalPermintaan: permintaanStore.selectedRequest.tanggalPermintaan,
        cancelReason: permintaanStore.selectedRequest.cancelReason || '',
      }
    }
  },
)

const handleClose = () => {
  emit('close')
}

const handleSubmit = async () => {
  isSubmitting.value = true
  try {
    if (permintaanStore.selectedRequest) {
      await permintaanStore.update(permintaanStore.selectedRequest.permintaanDarahId, formData.value)
      emit('submit')
      emit('close')
    }
  } finally {
    isSubmitting.value = false
  }
}
</script>

<template>
  <Teleport to="body">
    <Transition name="backdrop">
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
            <h2 class="text-xl font-semibold text-gray-900">Edit Permintaan Darah</h2>
            <p class="text-sm text-gray-400 mt-0.5">{{ subtitle }}</p>
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
                <label class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-1.5">
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
                <label class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-1.5">
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
                <label class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-1.5">
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
                <label class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-1.5">
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
                <label class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-1.5">
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
                <label class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-1.5">
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
                <label class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-1.5">
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
                <label class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-1.5">
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

            <!-- Pernah Hamil -->
            <div>
              <label class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-1.5">
                Pernah Hamil
              </label>
              <input
                v-model="formData.pernahHamil"
                type="text"
                class="w-full px-3.5 py-2.5 text-sm text-gray-900 bg-gray-50 border border-gray-200 rounded-xl outline-none transition-all focus:bg-white focus:border-blue-400 focus:ring-2 focus:ring-blue-100"
              />
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

            <!-- Status -->
            <div>
              <label class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-1.5">
                Status
              </label>
              <select
                v-model="formData.status"
                class="w-full px-3.5 py-2.5 text-sm text-gray-900 bg-gray-50 border border-gray-200 rounded-xl outline-none transition-all focus:bg-white focus:border-blue-400 focus:ring-2 focus:ring-blue-100"
              >
                <option value="dibuat">Dibuat</option>
                <option value="diproses">Diproses</option>
                <option value="bisa_diambil">Bisa Diambil</option>
                <option value="selesai">Selesai</option>
                <option value="dibatalkan">Dibatalkan</option>
              </select>
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
