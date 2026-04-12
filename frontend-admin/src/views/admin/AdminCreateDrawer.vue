<script setup lang="ts">
import { ref } from 'vue'
import { useAdminStore } from '@/stores/admin'
import { X, AlertCircle } from '@lucide/vue'

interface AdminData {
  adminUsername: string
  adminPassword: string
  adminName: string
  adminEmail: string
  adminRole: 'admin' | 'superadmin'
}

const props = defineProps<{
  isOpen: boolean
}>()

const emit = defineEmits<{
  close: []
  submit: []
}>()

const adminStore = useAdminStore()
const formData = ref<AdminData>({
  adminUsername: '',
  adminPassword: '',
  adminName: '',
  adminEmail: '',
  adminRole: 'admin',
})

const isSubmitting = ref(false)

const resetForm = () => {
  formData.value = {
    adminUsername: '',
    adminPassword: '',
    adminName: '',
    adminEmail: '',
    adminRole: 'admin',
  }
}

const handleSubmit = async () => {
  isSubmitting.value = true
  try {
    await adminStore.create(formData.value)
    emit('submit')
    resetForm()
    emit('close')
  } finally {
    isSubmitting.value = false
  }
}

const handleClose = () => {
  resetForm()
  emit('close')
}
</script>

<template>
  <Teleport to="body">
    <div v-if="isOpen" class="fixed inset-0 z-50 overflow-hidden">
      <!-- Backdrop -->
      <div class="absolute inset-0 bg-black/50" @click="handleClose" />

      <!-- Drawer -->
      <div class="fixed inset-y-0 right-0 flex items-center justify-end pointer-events-none">
        <div class="pointer-events-auto w-full max-w-lg max-h-screen overflow-y-auto bg-white shadow-2xl">
          <!-- Header -->
          <div class="sticky top-0 z-10 flex items-center justify-between px-6 py-4 border-b border-gray-100 bg-white">
            <div>
              <h2 class="text-lg font-semibold text-gray-900">Tambah Admin</h2>
              <p class="text-sm text-gray-500 mt-1">Buat admin baru untuk sistem</p>
            </div>
            <button
              @click="handleClose"
              class="p-1.5 rounded-lg hover:bg-gray-100 text-gray-400 transition-colors"
            >
              <X :size="20" />
            </button>
          </div>

          <!-- Content -->
          <form @submit.prevent="handleSubmit" class="px-6 py-5 space-y-5">
            <div>
              <label class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-1.5">
                Username
              </label>
              <input
                v-model="formData.adminUsername"
                type="text"
                required
                class="w-full px-3.5 py-2.5 text-sm text-gray-900 bg-gray-50 border border-gray-200 rounded-xl outline-none transition-all focus:bg-white focus:border-blue-400 focus:ring-2 focus:ring-blue-100"
              />
            </div>

            <div>
              <label class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-1.5">
                Nama
              </label>
              <input
                v-model="formData.adminName"
                type="text"
                required
                class="w-full px-3.5 py-2.5 text-sm text-gray-900 bg-gray-50 border border-gray-200 rounded-xl outline-none transition-all focus:bg-white focus:border-blue-400 focus:ring-2 focus:ring-blue-100"
              />
            </div>

            <div>
              <label class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-1.5">
                Email
              </label>
              <input
                v-model="formData.adminEmail"
                type="email"
                required
                class="w-full px-3.5 py-2.5 text-sm text-gray-900 bg-gray-50 border border-gray-200 rounded-xl outline-none transition-all focus:bg-white focus:border-blue-400 focus:ring-2 focus:ring-blue-100"
              />
            </div>

            <div>
              <label class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-1.5">
                Password
              </label>
              <input
                v-model="formData.adminPassword"
                type="password"
                required
                class="w-full px-3.5 py-2.5 text-sm text-gray-900 bg-gray-50 border border-gray-200 rounded-xl outline-none transition-all focus:bg-white focus:border-blue-400 focus:ring-2 focus:ring-blue-100"
              />
            </div>

            <div>
              <label class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-1.5">
                Role
              </label>
              <select
                v-model="formData.adminRole"
                class="w-full px-3.5 py-2.5 text-sm text-gray-900 bg-gray-50 border border-gray-200 rounded-xl outline-none transition-all focus:bg-white focus:border-blue-400 focus:ring-2 focus:ring-blue-100"
              >
                <option value="admin">Admin</option>
                <option value="superadmin">Superadmin</option>
              </select>
            </div>

            <div v-if="adminStore.error" class="flex items-center gap-2 p-3 bg-red-50 border border-red-100 text-red-600 rounded-xl text-xs">
              <AlertCircle :size="14" class="shrink-0" />
              {{ adminStore.error }}
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
