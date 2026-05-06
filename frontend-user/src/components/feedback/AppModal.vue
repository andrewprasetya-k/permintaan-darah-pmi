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
    <Transition
      enter-active-class="transition-opacity duration-200"
      enter-from-class="opacity-0"
      leave-active-class="transition-opacity duration-200"
      leave-to-class="opacity-0"
    >
      <div
        v-if="isOpen"
        class="fixed inset-0 z-[80] bg-black/30 backdrop-blur-sm"
        @click="$emit('close')"
      />
    </Transition>

    <Transition
      enter-active-class="transition duration-200"
      enter-from-class="-translate-x-1/2 -translate-y-[47%] opacity-0"
      leave-active-class="transition duration-200"
      leave-to-class="-translate-x-1/2 -translate-y-[47%] opacity-0"
    >
      <section
        v-if="isOpen"
        class="fixed left-1/2 top-1/2 z-[90] max-h-[calc(100vh-32px)] w-[calc(100vw-32px)] -translate-x-1/2 -translate-y-1/2 overflow-auto rounded-3xl bg-white p-6 shadow-2xl"
        :class="{
          'max-w-[420px]': width === 'sm',
          'max-w-[560px]': width === 'md',
          'max-w-[760px]': width === 'lg',
        }"
        role="dialog"
        aria-modal="true"
        :aria-label="title"
      >
        <div class="flex items-start gap-4">
          <slot name="icon" />
          <div class="min-w-0 flex-1">
            <h2 class="m-0 text-lg font-semibold text-gray-900">{{ title }}</h2>
            <p v-if="description" class="mt-2 text-sm leading-6 text-gray-500">
              {{ description }}
            </p>
          </div>
        </div>

        <div v-if="$slots.default" class="mt-4">
          <slot />
        </div>

        <footer v-if="$slots.footer" class="mt-6 flex gap-3 max-sm:flex-col-reverse">
          <slot name="footer" />
        </footer>
      </section>
    </Transition>
  </Teleport>
</template>
