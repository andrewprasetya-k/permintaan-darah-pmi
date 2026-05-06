<script setup lang="ts">
const props = withDefaults(
  defineProps<{
    title: string
    message?: string
    variant?: 'success' | 'error' | 'info'
  }>(),
  {
    variant: 'success',
    message: '',
  },
)

defineEmits<{
  close: []
}>()
</script>

<template>
  <div
    class="fixed right-5 top-5 z-[100] flex w-[min(384px,calc(100vw-32px))] items-start gap-3 rounded-2xl border px-4 py-3 shadow-lg max-sm:left-3.5 max-sm:right-3.5 max-sm:top-3.5 max-sm:w-auto"
    :class="{
      'border-emerald-100 bg-emerald-50 text-emerald-700': props.variant === 'success',
      'border-red-100 bg-red-50 text-red-700': props.variant === 'error',
      'border-blue-100 bg-blue-50 text-blue-700': props.variant === 'info',
    }"
    role="status"
    aria-live="polite"
  >
    <span class="mt-1.5 h-2.5 w-2.5 shrink-0 rounded-full bg-current" />
    <div class="min-w-0 flex-1">
      <p class="m-0 text-sm font-semibold">{{ title }}</p>
      <span v-if="message" class="mt-1 block text-sm leading-5 opacity-90">{{ message }}</span>
    </div>
    <button
      type="button"
      class="h-6 w-6 shrink-0 rounded-lg text-xs font-semibold opacity-70 transition hover:bg-black/10 hover:opacity-100"
      aria-label="Tutup notifikasi"
      @click="$emit('close')"
    >
      x
    </button>
  </div>
</template>
