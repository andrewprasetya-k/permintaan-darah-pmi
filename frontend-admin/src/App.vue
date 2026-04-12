<script setup lang="ts">
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import MainLayout from './components/layout/MainLayout.vue'

const route = useRoute()

// Pages yang ga perlu MainLayout (seperti login)
const noLayoutPages = ['login']
const shouldShowLayout = computed(() => !noLayoutPages.includes(route.name as string))
</script>

<template>
  <MainLayout v-if="shouldShowLayout">
    <RouterView />
  </MainLayout>
  <RouterView v-else />
</template>
