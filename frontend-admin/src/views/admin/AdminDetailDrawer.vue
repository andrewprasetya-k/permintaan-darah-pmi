<script setup lang="ts">
import { computed } from 'vue'
import { useAdminStore } from '@/stores/admin'
import { X, Shield, User, Mail, AtSign, Calendar, Clock } from '@lucide/vue'

defineProps<{ isOpen: boolean }>()
defineEmits<{ close: [] }>()

const adminStore = useAdminStore()
const admin = computed(() => adminStore.selectedAdmin)

const formatDate = (d: string) =>
  new Date(d).toLocaleDateString('id-ID', {
    weekday: 'short',
    year: 'numeric',
    month: 'short',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
  })
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
        <div class="flex items-center justify-between px-10 py-8 border-b border-gray-100">
          <div>
            <h2 class="text-xl font-semibold text-gray-900">Detail Admin</h2>
            <p class="text-sm text-gray-400 mt-0.5">{{ admin?.adminName }}</p>
          </div>
          <button
            @click="$emit('close')"
            class="p-2 rounded-lg hover:bg-gray-100 text-gray-400 transition-colors"
          >
            <X :size="20" />
          </button>
        </div>

        <div v-if="admin" class="flex-1 overflow-y-auto px-10 py-8 space-y-6">
          <!-- Info Fields -->
          <div class="bg-white border border-gray-100 rounded-2xl divide-y divide-gray-50">
            <div class="flex items-center gap-4 px-6 py-4">
              <div class="w-8 h-8 bg-gray-50 rounded-lg flex items-center justify-center shrink-0">
                <User :size="14" class="text-gray-400" />
              </div>
              <div>
                <p class="text-xs text-gray-400 uppercase tracking-wide font-medium">Nama</p>
                <p class="text-sm font-medium text-gray-900 mt-0.5">{{ admin.adminName }}</p>
              </div>
            </div>

            <div class="flex items-center gap-4 px-6 py-4">
              <div class="w-8 h-8 bg-gray-50 rounded-lg flex items-center justify-center shrink-0">
                <AtSign :size="14" class="text-gray-400" />
              </div>
              <div>
                <p class="text-xs text-gray-400 uppercase tracking-wide font-medium">Username</p>
                <p class="text-sm font-medium text-gray-900 mt-0.5">{{ admin.adminUsername }}</p>
              </div>
            </div>

            <div class="flex items-center gap-4 px-6 py-4">
              <div class="w-8 h-8 bg-gray-50 rounded-lg flex items-center justify-center shrink-0">
                <Mail :size="14" class="text-gray-400" />
              </div>
              <div>
                <p class="text-xs text-gray-400 uppercase tracking-wide font-medium">Email</p>
                <p class="text-sm font-medium text-gray-900 mt-0.5">{{ admin.adminEmail }}</p>
              </div>
            </div>
          </div>

          <!-- Timestamps -->
          <div class="bg-white border border-gray-100 rounded-2xl divide-y divide-gray-50">
            <div class="flex items-center gap-4 px-6 py-4">
              <div class="w-8 h-8 bg-gray-50 rounded-lg flex items-center justify-center shrink-0">
                <Calendar :size="14" class="text-gray-400" />
              </div>
              <div>
                <p class="text-xs text-gray-400 uppercase tracking-wide font-medium">Dibuat</p>
                <p class="text-sm text-gray-600 mt-0.5">{{ formatDate(admin.createdAt) }}</p>
              </div>
            </div>
            <div class="flex items-center gap-4 px-6 py-4">
              <div class="w-8 h-8 bg-gray-50 rounded-lg flex items-center justify-center shrink-0">
                <Clock :size="14" class="text-gray-400" />
              </div>
              <div>
                <p class="text-xs text-gray-400 uppercase tracking-wide font-medium">Diperbarui</p>
                <p class="text-sm text-gray-600 mt-0.5">{{ formatDate(admin.updatedAt) }}</p>
              </div>
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
