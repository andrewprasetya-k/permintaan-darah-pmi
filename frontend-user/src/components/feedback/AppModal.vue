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
      enter-active-class="transition duration-200"
      enter-from-class="opacity-0"
      leave-active-class="transition duration-200"
      leave-to-class="opacity-0"
    >
      <div
        v-if="isOpen"
        class="fixed inset-0 z-[90] grid min-h-dvh place-items-center bg-black/30 p-4 backdrop-blur-sm"
        @click="$emit('close')"
      >
        <section
          class="max-h-[calc(100dvh-32px)] w-full overflow-auto rounded-3xl bg-white p-6 shadow-2xl"
          :class="{
            'max-w-[420px]': width === 'sm',
            'max-w-[560px]': width === 'md',
            'max-w-[760px]': width === 'lg',
          }"
          role="dialog"
          aria-modal="true"
          :aria-label="title"
          @click.stop
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
      </div>
    </Transition>
  </Teleport>
</template>
