<script setup lang="ts">
import { computed, ref, watch } from 'vue'
import { useKomponenStore } from '@/stores/komponen'
import { X, AlertCircle } from '@lucide/vue'

interface KomponenData {
  komponenDarah: string
  komponenKode: string
}

const props = defineProps<{
  isOpen: boolean
}>()

const emit = defineEmits<{
  close: []
  submit: []
}>()

const komponenStore = useKomponenStore()
const formData = ref<KomponenData>({
  komponenDarah: '',
  komponenKode: '',
})

const isSubmitting = ref(false)

const subtitle = computed(() => 'Perbarui informasi komponen darah')

watch(
  () => props.isOpen,
  (newVal) => {
    if (newVal && komponenStore.selectedKomponen) {
      formData.value = {
        komponenDarah: komponenStore.selectedKomponen.komponenDarah,
        komponenKode: komponenStore.selectedKomponen.komponenKode,
      }
    }
  }
)

const handleSubmit = async () => {
  isSubmitting.value = true
  try {
    if (komponenStore.selectedKomponen) {
      await komponenStore.update(komponenStore.selectedKomponen.komponenId, formData.value)
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
            <h2 class="text-xl font-semibold text-gray-900">Edit Komponen Darah</h2>
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
            <div>
              <label class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-1.5">
                Nama Komponen
              </label>
              <input
                v-model="formData.komponenDarah"
                type="text"
                required
                class="w-full px-3.5 py-2.5 text-sm text-gray-900 bg-gray-50 border border-gray-200 rounded-xl outline-none transition-all focus:bg-white focus:border-blue-400 focus:ring-2 focus:ring-blue-100"
              />
            </div>

            <div>
              <label class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-1.5">
                Kode Komponen
              </label>
              <input
                v-model="formData.komponenKode"
                type="text"
                required
                class="w-full px-3.5 py-2.5 text-sm text-gray-900 bg-gray-50 border border-gray-200 rounded-xl outline-none transition-all focus:bg-white focus:border-blue-400 focus:ring-2 focus:ring-blue-100"
              />
            </div>

            <div v-if="komponenStore.error" class="flex items-center gap-2 p-3 bg-red-50 border border-red-100 text-red-600 rounded-xl text-xs">
              <AlertCircle :size="14" class="shrink-0" />
              {{ komponenStore.error }}
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
