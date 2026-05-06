<script setup lang="ts">
withDefaults(
  defineProps<{
    isOpen: boolean
    title: string
    description?: string
    width?: 'sm' | 'md' | 'lg'
  }>(),
  {
    description: '',
    width: 'md',
  },
)

defineEmits<{
  close: []
}>()
</script>

<template>
  <Teleport to="body">
    <Transition name="modal-backdrop">
      <button
        v-if="isOpen"
        type="button"
        class="modal-backdrop"
        aria-label="Tutup dialog"
        @click="$emit('close')"
      />
    </Transition>

    <Transition name="modal-panel">
      <section
        v-if="isOpen"
        class="modal-panel"
        :class="`modal-${width}`"
        role="dialog"
        aria-modal="true"
        :aria-label="title"
      >
        <header class="modal-header">
          <div>
            <h2>{{ title }}</h2>
            <p v-if="description">{{ description }}</p>
          </div>
          <button type="button" class="btn btn-ghost btn-icon" aria-label="Tutup dialog" @click="$emit('close')">
            x
          </button>
        </header>

        <div class="modal-body">
          <slot />
        </div>

        <footer v-if="$slots.footer" class="modal-footer">
          <slot name="footer" />
        </footer>
      </section>
    </Transition>
  </Teleport>
</template>

<style scoped>
.modal-backdrop {
  position: fixed;
  inset: 0;
  z-index: 80;
  border: 0;
  background: rgba(15, 23, 42, 0.42);
  backdrop-filter: blur(4px);
}

.modal-panel {
  position: fixed;
  top: 50%;
  left: 50%;
  z-index: 90;
  width: calc(100vw - 32px);
  max-height: calc(100vh - 32px);
  overflow: auto;
  transform: translate(-50%, -50%);
  border: 1px solid var(--line);
  border-radius: 8px;
  background: var(--surface);
  box-shadow: var(--shadow);
}

.modal-sm {
  max-width: 420px;
}

.modal-md {
  max-width: 560px;
}

.modal-lg {
  max-width: 760px;
}

.modal-header {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 16px;
  border-bottom: 1px solid var(--line);
  padding: 18px 20px;
}

.modal-header h2 {
  margin: 0;
  font-size: 18px;
  font-weight: 800;
}

.modal-header p {
  margin: 6px 0 0;
  color: var(--text-soft);
}

.modal-body {
  padding: 20px;
}

.modal-footer {
  display: flex;
  justify-content: flex-end;
  gap: 10px;
  border-top: 1px solid var(--line);
  padding: 16px 20px;
}

.modal-backdrop-enter-active,
.modal-backdrop-leave-active,
.modal-panel-enter-active,
.modal-panel-leave-active {
  transition:
    opacity 0.18s ease,
    transform 0.18s ease;
}

.modal-backdrop-enter-from,
.modal-backdrop-leave-to,
.modal-panel-enter-from,
.modal-panel-leave-to {
  opacity: 0;
}

.modal-panel-enter-from,
.modal-panel-leave-to {
  transform: translate(-50%, -47%);
}

@media (max-width: 560px) {
  .modal-footer {
    flex-direction: column-reverse;
  }

  .modal-footer :deep(.btn) {
    width: 100%;
  }
}
</style>
