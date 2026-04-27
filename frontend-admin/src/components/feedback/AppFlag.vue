<script setup lang="ts">
const props = withDefaults(
  defineProps<{
    title: string
    message?: string
    variant?: 'success' | 'error'
  }>(),
  {
    variant: 'success',
    message: '',
  },
)

defineEmits<{
  close: []
}>()

const variantClass =
  props.variant === 'success'
    ? 'border-emerald-100 bg-emerald-50 text-emerald-700'
    : 'border-red-100 bg-red-50 text-red-700'
const badgeClass = props.variant === 'success' ? 'bg-emerald-500' : 'bg-red-500'
</script>

<template>
  <div
    class="fixed right-6 top-6 z-100 w-full max-w-sm rounded-2xl border px-4 py-3 shadow-lg"
    :class="variantClass"
  >
    <div class="flex items-start gap-3">
      <span class="mt-1 h-2.5 w-2.5 rounded-full" :class="badgeClass" />
      <div class="min-w-0 flex-1">
        <p class="text-sm font-semibold">{{ title }}</p>
        <p v-if="message" class="mt-1 text-sm leading-5 opacity-90">{{ message }}</p>
      </div>
      <button
        type="button"
        class="text-xs font-medium opacity-70 transition-opacity hover:opacity-100"
        @click="$emit('close')"
      >
        x
      </button>
    </div>
  </div>
</template>
