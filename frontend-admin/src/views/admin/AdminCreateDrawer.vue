<script setup lang="ts">
import { ref } from 'vue'
import { useAdminStore } from '@/stores/admin'
import { X, AlertCircle, Eye, EyeOff } from '@lucide/vue'
import AppModal from '@/components/feedback/AppModal.vue'
import AppFlag from '@/components/feedback/AppFlag.vue'

interface AdminData {
  adminUsername: string
  adminPassword: string
  adminName: string
  adminEmail: string
  adminRole: 'admin' | 'superadmin'
}

defineProps<{ isOpen: boolean }>()

const emit = defineEmits<{
  close: []
  submit: []
}>()

const adminStore = useAdminStore()
const showPassword = ref(false)
const isSubmitting = ref(false)
const showSubmitDialog = ref(false)
const flag = ref<{ variant: 'success' | 'error'; title: string; message?: string } | null>(null)
const formData = ref<AdminData>({
  adminUsername: '',
  adminPassword: '',
  adminName: '',
  adminEmail: '',
  adminRole: 'admin',
})

const resetForm = () => {
  formData.value = {
    adminUsername: '',
    adminPassword: '',
    adminName: '',
    adminEmail: '',
    adminRole: 'admin',
  }
  showPassword.value = false
}

const handleClose = () => {
  if (isSubmitting.value) return
  showSubmitDialog.value = false
  resetForm()
  emit('close')
}

const openSubmitDialog = () => {
  showSubmitDialog.value = true
}

const closeSubmitDialog = () => {
  if (isSubmitting.value) return
  showSubmitDialog.value = false
}

const handleSubmit = async () => {
  isSubmitting.value = true
  try {
    await adminStore.create(formData.value)
    flag.value = {
      variant: 'success',
      title: 'Admin dibuat',
      message: `${formData.value.adminName} berhasil ditambahkan ke sistem.`,
    }
    emit('submit')
    resetForm()
    emit('close')
  } catch (error) {
    flag.value = {
      variant: 'error',
      title: 'Operasi gagal',
      message: error instanceof Error ? error.message : 'Gagal membuat admin baru.',
    }
  } finally {
    isSubmitting.value = false
    showSubmitDialog.value = false
  }
}
</script>

<template>
  <Teleport to="body">
    <AppFlag
      v-if="flag"
      :variant="flag.variant"
      :title="flag.title"
      :message="flag.message"
      @close="flag = null"
    />

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
            <h2 class="text-xl font-semibold text-gray-900">Tambah Admin</h2>
            <p class="text-sm text-gray-400 mt-0.5">Buat admin baru untuk sistem</p>
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
          <form @submit.prevent="openSubmitDialog" class="space-y-5 max-w-full">
            <div
              v-for="field in [
                { key: 'adminUsername', label: 'Username', type: 'text' },
                { key: 'adminName', label: 'Nama', type: 'text' },
                { key: 'adminEmail', label: 'Email', type: 'email' },
              ]"
              :key="field.key"
            >
              <label class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-1.5">
                {{ field.label }}
              </label>
              <input
                v-model="(formData as any)[field.key]"
                :type="field.type"
                required
                class="w-full px-3.5 py-2.5 text-sm text-gray-900 bg-gray-50 border border-gray-200 rounded-xl outline-none transition-all focus:bg-white focus:border-blue-400 focus:ring-2 focus:ring-blue-100"
              />
            </div>

            <!-- Password -->
            <div>
              <label class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-1.5">
                Password
              </label>
              <div class="relative">
                <input
                  v-model="formData.adminPassword"
                  :type="showPassword ? 'text' : 'password'"
                  required
                  class="w-full px-3.5 py-2.5 pr-10 text-sm text-gray-900 bg-gray-50 border border-gray-200 rounded-xl outline-none transition-all focus:bg-white focus:border-blue-400 focus:ring-2 focus:ring-blue-100"
                />
                <button
                  type="button"
                  @click="showPassword = !showPassword"
                  class="absolute right-3 top-1/2 -translate-y-1/2 text-gray-400 hover:text-gray-600 transition-colors"
                >
                  <Eye v-if="!showPassword" :size="16" />
                  <EyeOff v-else :size="16" />
                </button>
              </div>
            </div>

            <!-- Role -->
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

            <!-- Error -->
            <div
              v-if="adminStore.error"
              class="flex items-center gap-2 p-3 bg-red-50 border border-red-100 text-red-600 rounded-xl text-xs"
            >
              <AlertCircle :size="14" class="shrink-0" />
              {{ adminStore.error }}
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

    <AppModal
      :is-open="showSubmitDialog"
      title="Simpan Admin Baru"
      :description="`Admin ${formData.adminName || formData.adminUsername || 'baru'} akan ditambahkan ke sistem.`"
      @close="closeSubmitDialog"
    >
      <template #footer>
        <button
          type="button"
          class="flex-1 rounded-xl bg-gray-100 px-4 py-2.5 text-sm font-medium text-gray-700 transition-colors hover:bg-gray-200"
          @click="closeSubmitDialog"
        >
          Batal
        </button>
        <button
          type="button"
          :disabled="isSubmitting"
          class="flex-1 rounded-xl bg-blue-600 px-4 py-2.5 text-sm font-medium text-white transition-colors hover:bg-blue-700 disabled:opacity-50"
          @click="handleSubmit"
        >
          {{ isSubmitting ? 'Menyimpan...' : 'Simpan' }}
        </button>
      </template>
    </AppModal>
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
