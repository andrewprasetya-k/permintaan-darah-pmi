<script setup lang="ts">
import { computed, ref, watch } from 'vue'
import { useRumahSakitStore } from '@/stores/rumah-sakit'
import { X, AlertCircle } from '@lucide/vue'

interface RumahSakitData {
  nama: string
  nomorTelepon: string
  alamat: string
  email: string
  username: string
  password: string
}

const props = defineProps<{
  isOpen: boolean
}>()

const emit = defineEmits<{
  close: []
  submit: []
}>()

const rumahSakitStore = useRumahSakitStore()
const formData = ref<RumahSakitData>({
  nama: '',
  nomorTelepon: '',
  alamat: '',
  email: '',
  username: '',
  password: '',
})

const isSubmitting = ref(false)

const subtitle = computed(() => 'Perbarui informasi rumah sakit')

watch(
  () => props.isOpen,
  (newVal) => {
    if (newVal && rumahSakitStore.selectedRumahSakit) {
      formData.value = {
        nama: rumahSakitStore.selectedRumahSakit.nama,
        nomorTelepon: rumahSakitStore.selectedRumahSakit.nomorTelepon,
        alamat: rumahSakitStore.selectedRumahSakit.alamat,
        email: rumahSakitStore.selectedRumahSakit.email || '',
        username: rumahSakitStore.selectedRumahSakit.username,
        password: '',
      }
    }
  }
)

const handleSubmit = async () => {
  isSubmitting.value = true
  try {
    if (rumahSakitStore.selectedRumahSakit) {
      await rumahSakitStore.update(rumahSakitStore.selectedRumahSakit.rumahSakitId, formData.value)
      emit('submit')
      emit('close')
    }
  } finally {
    isSubmitting.value = false
  }
}

const handleClose = () => {
  emit('close')
}
</script>

<template>
  <Teleport to="body">
    <div v-if="isOpen" class="fixed inset-0 z-50 overflow-hidden">
      <div class="absolute inset-0 bg-black/50" @click="handleClose" />

      <div class="fixed inset-y-0 right-0 flex items-center justify-end pointer-events-none">
        <div class="pointer-events-auto w-full max-w-lg max-h-screen overflow-y-auto bg-white shadow-2xl">
          <div class="sticky top-0 z-10 flex items-center justify-between px-6 py-4 border-b border-gray-100 bg-white">
            <div>
              <h2 class="text-lg font-semibold text-gray-900">Edit Rumah Sakit</h2>
              <p v-if="subtitle" class="text-sm text-gray-500 mt-1">{{ subtitle }}</p>
            </div>
            <button
              @click="handleClose"
              class="p-1.5 rounded-lg hover:bg-gray-100 text-gray-400 transition-colors"
            >
              <X :size="20" />
            </button>
          </div>

          <form @submit.prevent="handleSubmit" class="px-6 py-5 space-y-5">
            <div>
              <label class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-1.5">
                Nama Rumah Sakit
              </label>
              <input
                v-model="formData.nama"
                type="text"
                required
                class="w-full px-3.5 py-2.5 text-sm text-gray-900 bg-gray-50 border border-gray-200 rounded-xl outline-none transition-all focus:bg-white focus:border-blue-400 focus:ring-2 focus:ring-blue-100"
              />
            </div>

            <div>
              <label class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-1.5">
                Username
              </label>
              <input
                v-model="formData.username"
                type="text"
                required
                disabled
                class="w-full px-3.5 py-2.5 text-sm text-gray-900 bg-gray-50 border border-gray-200 rounded-xl outline-none transition-all focus:bg-white focus:border-blue-400 focus:ring-2 focus:ring-blue-100 disabled:opacity-50"
              />
            </div>

            <div>
              <label class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-1.5">
                Password
                <span class="normal-case text-gray-300 ml-1">(kosongkan jika tidak diubah)</span>
              </label>
              <input
                v-model="formData.password"
                type="password"
                class="w-full px-3.5 py-2.5 text-sm text-gray-900 bg-gray-50 border border-gray-200 rounded-xl outline-none transition-all focus:bg-white focus:border-blue-400 focus:ring-2 focus:ring-blue-100"
              />
            </div>

            <div>
              <label class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-1.5">
                Email
              </label>
              <input
                v-model="formData.email"
                type="email"
                required
                class="w-full px-3.5 py-2.5 text-sm text-gray-900 bg-gray-50 border border-gray-200 rounded-xl outline-none transition-all focus:bg-white focus:border-blue-400 focus:ring-2 focus:ring-blue-100"
              />
            </div>

            <div>
              <label class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-1.5">
                Nomor Telepon
              </label>
              <input
                v-model="formData.nomorTelepon"
                type="tel"
                required
                class="w-full px-3.5 py-2.5 text-sm text-gray-900 bg-gray-50 border border-gray-200 rounded-xl outline-none transition-all focus:bg-white focus:border-blue-400 focus:ring-2 focus:ring-blue-100"
              />
            </div>

            <div>
              <label class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-1.5">
                Alamat
              </label>
              <textarea
                v-model="formData.alamat"
                required
                rows="3"
                class="w-full px-3.5 py-2.5 text-sm text-gray-900 bg-gray-50 border border-gray-200 rounded-xl outline-none transition-all focus:bg-white focus:border-blue-400 focus:ring-2 focus:ring-blue-100"
              />
            </div>

            <div v-if="rumahSakitStore.error" class="flex items-center gap-2 p-3 bg-red-50 border border-red-100 text-red-600 rounded-xl text-xs">
              <AlertCircle :size="14" class="shrink-0" />
              {{ rumahSakitStore.error }}
            </div>

            <div class="flex gap-3 pt-1">
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
    </div>
  </Teleport>
</template>
