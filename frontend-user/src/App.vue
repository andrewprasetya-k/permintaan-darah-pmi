<script setup lang="ts">
import { RouterView } from 'vue-router'
import AppFlag from '@/components/feedback/AppFlag.vue'
import { useAuthStore } from '@/stores/auth'
import { useFeedbackStore } from '@/stores/feedback'

const authStore = useAuthStore()
const feedbackStore = useFeedbackStore()

authStore.initializeFromStorage()
</script>

<template>
  <RouterView />
  <AppFlag
    v-if="feedbackStore.flag"
    :title="feedbackStore.flag.title"
    :message="feedbackStore.flag.message"
    :variant="feedbackStore.flag.variant"
    @close="feedbackStore.clearFlag"
  />
</template>
