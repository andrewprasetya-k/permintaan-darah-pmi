<script setup lang="ts">
import { computed } from 'vue'
import { useAdminStore } from '@/stores/admin'
import { X, Shield } from '@lucide/vue'

const props = defineProps<{
  isOpen: boolean
}>()

const emit = defineEmits<{
  close: []
}>()

const adminStore = useAdminStore()

const subtitle = computed(() => adminStore.selectedAdmin?.adminName)
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
            <h2 class="text-xl font-semibold text-gray-900">Detail Admin</h2>
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
        <div v-if="adminStore.selectedAdmin" class="flex-1 overflow-y-auto px-10 py-8">
          <div class="space-y-6 max-w-full">
            <div class="space-y-4">
              <div>
                <label class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-2">
                  Nama
                </label>
                <p class="text-base text-gray-900 font-medium">{{ adminStore.selectedAdmin.adminName }}</p>
              </div>

              <div>
                <label class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-2">
                  Username
                </label>
                <p class="text-base text-gray-900 font-medium">{{ adminStore.selectedAdmin.adminUsername }}</p>
              </div>

              <div>
                <label class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-2">
                  Email
                </label>
                <p class="text-base text-gray-900 font-medium">{{ adminStore.selectedAdmin.adminEmail }}</p>
              </div>

              <div>
                <label class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-2">
                  Role
                </label>
                <div class="flex items-center gap-2">
                  <span
                    class="inline-flex items-center gap-1.5 px-3 py-1.5 rounded-lg text-xs font-medium"
                    :class="
                      adminStore.selectedAdmin.adminRole === 'superadmin'
                        ? 'bg-purple-50 text-purple-600'
                        : 'bg-blue-50 text-blue-600'
                    "
                  >
                    <Shield :size="14" />
                    {{ adminStore.selectedAdmin.adminRole === 'superadmin' ? 'Superadmin' : 'Admin' }}
                  </span>
                </div>
              </div>

              <hr class="my-4 border-gray-100" />

              <div>
                <label class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-2">
                  Dibuat
                </label>
                <p class="text-sm text-gray-600">
                  {{ new Date(adminStore.selectedAdmin.createdAt).toLocaleDateString('id-ID', { 
                    weekday: 'short',
                    year: 'numeric', 
                    month: 'short', 
                    day: 'numeric',
                    hour: '2-digit',
                    minute: '2-digit'
                  }) }}
                </p>
              </div>

              <div>
                <label class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-2">
                  Diperbarui
                </label>
                <p class="text-sm text-gray-600">
                  {{ new Date(adminStore.selectedAdmin.updatedAt).toLocaleDateString('id-ID', { 
                    weekday: 'short',
                    year: 'numeric', 
                    month: 'short', 
                    day: 'numeric',
                    hour: '2-digit',
                    minute: '2-digit'
                  }) }}
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
