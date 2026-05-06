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
      <div v-if="isOpen" class="modal-backdrop" @click="$emit('close')" />
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
        <div class="modal-copy">
          <slot name="icon" />
          <div class="modal-text">
            <h2>{{ title }}</h2>
            <p v-if="description">{{ description }}</p>
          </div>
        </div>

        <div v-if="$slots.default" class="modal-body">
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
  background: rgba(0, 0, 0, 0.3);
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
  border-radius: 24px;
  background: var(--surface);
  padding: 24px;
  box-shadow: var(--shadow-lg);
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

.modal-copy {
  display: flex;
  align-items: flex-start;
  gap: 16px;
}

.modal-text {
  min-width: 0;
  flex: 1;
}

.modal-text h2 {
  margin: 0;
  font-size: 18px;
  font-weight: 600;
  color: #111827;
}

.modal-text p {
  margin: 8px 0 0;
  color: #6b7280;
  font-size: 14px;
  line-height: 1.6;
}

.modal-body {
  margin-top: 18px;
}

.modal-footer {
  display: flex;
  justify-content: stretch;
  gap: 10px;
  margin-top: 24px;
}

.modal-footer :deep(.btn) {
  flex: 1;
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
