<script setup lang="ts">
import { useRouter } from 'vue-router'

const router = useRouter()
</script>

<template>
  <div class="min-h-screen flex flex-col items-center justify-center">
    <h1 class="text-6xl font-bold">404</h1>
    <p class="text-xl text-gray-600 mt-4">Halaman tidak ditemukan</p>
    <button
      @click="router.push('/')"
      class="mt-6 px-6 py-2 bg-blue-600 text-white rounded hover:bg-blue-700"
    >
      Kembali ke Home
    </button>
  </div>
</template>
