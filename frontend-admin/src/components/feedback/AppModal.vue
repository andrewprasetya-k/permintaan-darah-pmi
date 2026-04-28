<script setup lang="ts">
defineProps<{
  isOpen: boolean
  title: string
  description?: string
  widthClass?: string
}>()

defineEmits<{
  close: []
}>()
</script>

<template>
  <Teleport to="body">
    <Transition name="backdrop">
      <div
        v-if="isOpen"
        class="fixed inset-0 z-[80] bg-black/30 backdrop-blur-sm"
        @click="$emit('close')"
      />
    </Transition>

    <Transition name="modal">
      <div
        v-if="isOpen"
        class="fixed left-1/2 top-1/2 z-[90] w-full -translate-x-1/2 -translate-y-1/2 rounded-3xl bg-white p-6 shadow-2xl"
        :class="widthClass || 'max-w-lg'"
      >
        <div class="flex items-start gap-4">
          <slot name="icon" />
          <div class="min-w-0 flex-1">
            <h3 class="text-lg font-semibold text-gray-900">{{ title }}</h3>
            <p v-if="description" class="mt-2 text-sm leading-6 text-gray-500">
              {{ description }}
            </p>
          </div>
        </div>

        <div class="mt-4">
          <slot />
        </div>

        <div class="mt-6 flex gap-3">
          <slot name="footer" />
        </div>
      </div>
    </Transition>
  </Teleport>
</template>

<style scoped>
.backdrop-enter-active,
.backdrop-leave-active {
  transition: opacity 0.25s ease;
}

.backdrop-enter-from,
.backdrop-leave-to {
  opacity: 0;
}

.modal-enter-active,
.modal-leave-active {
  transition:
    transform 0.25s ease,
    opacity 0.25s ease;
}

.modal-enter-from,
.modal-leave-to {
  opacity: 0;
  transform: translate(-50%, -46%);
}
</style>
