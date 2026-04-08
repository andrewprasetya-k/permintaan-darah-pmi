<script setup lang="ts">
import { useAuthStore } from '@/stores/auth'

const authStore = useAuthStore()
</script>

<template>
  <div class="p-6">
    <h1 class="text-3xl font-bold">Dashboard</h1>
    <p class="text-gray-600 mt-2">Welcome, {{ authStore.user?.username }}</p>
  </div>
</template>
