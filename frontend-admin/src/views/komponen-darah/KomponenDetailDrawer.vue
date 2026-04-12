<script setup lang="ts">
import { computed } from 'vue'
import { useKomponenStore } from '@/stores/komponen'
import { X } from '@lucide/vue'

const props = defineProps<{
  isOpen: boolean
}>()

const emit = defineEmits<{
  close: []
}>()

const komponenStore = useKomponenStore()

const subtitle = computed(() => komponenStore.selectedKomponen?.komponenDarah)
</script>

<template>
  <Teleport to="body">
    <Transition name="backdrop">
      <div
        v-if="isOpen"
        class="fixed inset-0 bg-black/20 backdrop-blur-sm z-40"
        @click="$emit('close')"
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
            <h2 class="text-xl font-semibold text-gray-900">Detail Komponen Darah</h2>
            <p class="text-sm text-gray-400 mt-0.5">{{ subtitle }}</p>
          </div>
          <button
            @click="$emit('close')"
            class="p-2 rounded-lg hover:bg-gray-100 text-gray-400 transition-colors"
          >
            <X :size="20" />
          </button>
        </div>

        <!-- Content -->
        <div v-if="komponenStore.selectedKomponen" class="flex-1 overflow-y-auto px-10 py-8">
          <div class="space-y-6 max-w-full">
            <div>
              <label class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-1.5">
                ID
              </label>
              <p class="text-sm text-gray-900">{{ komponenStore.selectedKomponen.komponenId }}</p>
            </div>

            <div>
              <label class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-1.5">
                Nama Komponen
              </label>
              <p class="text-sm text-gray-900">
                {{ komponenStore.selectedKomponen.komponenDarah }}
              </p>
            </div>

            <div>
              <label class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-1.5">
                Kode Komponen
              </label>
              <p class="text-sm text-gray-900">{{ komponenStore.selectedKomponen.komponenKode }}</p>
            </div>

            <div>
              <label class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-1.5">
                Status
              </label>
              <p class="text-sm text-gray-900">
                <span
                  class="inline-flex px-2.5 py-1 rounded-lg text-xs font-medium"
                  :class="
                    komponenStore.selectedKomponen.isActive
                      ? 'bg-green-50 text-green-600'
                      : 'bg-gray-50 text-gray-600'
                  "
                >
                  {{ komponenStore.selectedKomponen.isActive ? 'Aktif' : 'Nonaktif' }}
                </span>
              </p>
            </div>
          </div>

          <!-- Close Button -->
          <div class="flex gap-3 pt-4">
            <button
              @click="$emit('close')"
              class="flex-1 py-2.5 bg-gray-100 hover:bg-gray-200 text-gray-600 text-sm font-medium rounded-xl transition-colors"
            >
              Tutup
            </button>
          </div>
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
