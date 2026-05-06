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
  <div class="app-flag" :class="`app-flag-${props.variant}`" role="status" aria-live="polite">
    <span class="app-flag-dot" />
    <div class="app-flag-copy">
      <p>{{ title }}</p>
      <span v-if="message">{{ message }}</span>
    </div>
    <button type="button" class="app-flag-close" aria-label="Tutup notifikasi" @click="$emit('close')">
      x
    </button>
  </div>
</template>

<style scoped>
.app-flag {
  position: fixed;
  top: 20px;
  right: 20px;
  z-index: 100;
  display: flex;
  width: min(420px, calc(100vw - 32px));
  align-items: flex-start;
  gap: 12px;
  border: 1px solid;
  border-radius: 8px;
  padding: 14px;
  box-shadow: var(--shadow);
}

.app-flag-success {
  border-color: #bbf7d0;
  background: #f0fdf4;
  color: #166534;
}

.app-flag-error {
  border-color: #fecaca;
  background: #fef2f2;
  color: #991b1b;
}

.app-flag-info {
  border-color: #bfdbfe;
  background: #eff6ff;
  color: #1d4ed8;
}

.app-flag-dot {
  width: 10px;
  height: 10px;
  flex: 0 0 auto;
  border-radius: 999px;
  background: currentColor;
  margin-top: 6px;
}

.app-flag-copy {
  min-width: 0;
  flex: 1;
}

.app-flag-copy p {
  margin: 0;
  font-weight: 800;
}

.app-flag-copy span {
  display: block;
  margin-top: 3px;
  color: inherit;
  font-size: 13px;
  opacity: 0.84;
}

.app-flag-close {
  width: 28px;
  height: 28px;
  flex: 0 0 auto;
  border: 0;
  border-radius: 8px;
  background: transparent;
  color: currentColor;
  font-weight: 800;
}

.app-flag-close:hover {
  background: rgba(15, 23, 42, 0.08);
}

@media (max-width: 640px) {
  .app-flag {
    top: 14px;
    right: 14px;
    left: 14px;
    width: auto;
  }
}
</style>
