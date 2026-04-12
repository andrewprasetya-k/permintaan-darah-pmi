<script setup lang="ts">
import { computed } from 'vue'
import { useRumahSakitStore } from '@/stores/rumah-sakit'
import { X } from '@lucide/vue'

const props = defineProps<{
  isOpen: boolean
}>()

const emit = defineEmits<{
  close: []
}>()

const rumahSakitStore = useRumahSakitStore()

const subtitle = computed(() => rumahSakitStore.selectedRumahSakit?.nama || '')
</script>

<template>
  <Teleport to="body">
    <div v-if="isOpen" class="fixed inset-0 z-50 overflow-hidden">
      <div class="absolute inset-0 bg-black/50" @click="$emit('close')" />

      <div class="fixed inset-y-0 right-0 flex items-center justify-end pointer-events-none">
        <div
          class="pointer-events-auto w-full max-w-lg max-h-screen overflow-y-auto bg-white shadow-2xl"
        >
          <div
            class="sticky top-0 z-10 flex items-center justify-between px-6 py-4 border-b border-gray-100 bg-white"
          >
            <div>
              <h2 class="text-lg font-semibold text-gray-900">Detail Rumah Sakit</h2>
              <p v-if="subtitle" class="text-sm text-gray-500 mt-1">{{ subtitle }}</p>
            </div>
            <button
              @click="$emit('close')"
              class="p-1.5 rounded-lg hover:bg-gray-100 text-gray-400 transition-colors"
            >
              <X :size="20" />
            </button>
          </div>

          <div v-if="rumahSakitStore.selectedRumahSakit" class="px-6 py-5 space-y-6">
            <div class="space-y-4">
              <div>
                <label
                  class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-1.5"
                >
                  Nama Rumah Sakit
                </label>
                <p class="text-sm text-gray-900">{{ rumahSakitStore.selectedRumahSakit.nama }}</p>
              </div>

              <div>
                <label
                  class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-1.5"
                >
                  Username
                </label>
                <p class="text-sm text-gray-900">
                  {{ rumahSakitStore.selectedRumahSakit.username }}
                </p>
              </div>

              <div>
                <label
                  class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-1.5"
                >
                  Email
                </label>
                <p class="text-sm text-gray-900">
                  {{ rumahSakitStore.selectedRumahSakit.email || '—' }}
                </p>
              </div>

              <div>
                <label
                  class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-1.5"
                >
                  Nomor Telepon
                </label>
                <p class="text-sm text-gray-900">
                  {{ rumahSakitStore.selectedRumahSakit.nomorTelepon }}
                </p>
              </div>

              <div>
                <label
                  class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-1.5"
                >
                  Alamat
                </label>
                <p class="text-sm text-gray-900">
                  {{ rumahSakitStore.selectedRumahSakit.alamat }}
                </p>
              </div>

              <div>
                <label
                  class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-1.5"
                >
                  Dibuat
                </label>
                <p class="text-sm text-gray-900">
                  {{
                    new Date(rumahSakitStore.selectedRumahSakit.createdAt).toLocaleDateString(
                      'id-ID',
                    )
                  }}
                </p>
              </div>

              <div>
                <label
                  class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-1.5"
                >
                  Diperbarui
                </label>
                <p class="text-sm text-gray-900">
                  {{
                    new Date(rumahSakitStore.selectedRumahSakit.updatedAt).toLocaleDateString(
                      'id-ID',
                    )
                  }}
                </p>
              </div>
            </div>

            <button
              @click="$emit('close')"
              class="w-full py-2.5 bg-gray-100 hover:bg-gray-200 text-gray-600 text-sm font-medium rounded-xl transition-colors"
            >
              Tutup
            </button>
          </div>
        </div>
      </div>
    </div>
  </Teleport>
</template>
