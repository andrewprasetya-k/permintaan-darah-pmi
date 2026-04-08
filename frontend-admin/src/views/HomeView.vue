<script setup lang="ts">
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const authStore = useAuthStore()

const logout = () => {
  authStore.logout()
  router.push('/login')
}
</script>

<template>
  <div class="min-h-screen bg-gray-100">
    <nav class="bg-white shadow">
      <div class="max-w-7xl mx-auto px-4 py-4 flex justify-between items-center">
        <h1 class="text-2xl font-bold">Admin Dashboard</h1>
        <div class="flex items-center gap-4">
          <span class="text-gray-700">{{ authStore.user?.name }}</span>
          <button
            @click="logout"
            class="px-4 py-2 bg-red-600 text-white rounded hover:bg-red-700"
          >
            Logout
          </button>
        </div>
      </div>
    </nav>

    <div class="max-w-7xl mx-auto p-6">
      <div class="bg-white p-6 rounded shadow">
        <h2 class="text-2xl font-bold mb-4">Selamat Datang!</h2>
        <p class="text-gray-600">
          Halo {{ authStore.user?.name }}, selamat datang di Admin Dashboard Permintaan Darah PMI
        </p>

        <div class="mt-8 grid grid-cols-2 gap-4">
          <router-link
            to="/permintaan"
            class="p-4 bg-blue-100 rounded hover:bg-blue-200 transition"
          >
            <h3 class="font-semibold text-blue-900">Permintaan Darah</h3>
          </router-link>

          <router-link
            to="/dashboard"
            class="p-4 bg-green-100 rounded hover:bg-green-200 transition"
          >
            <h3 class="font-semibold text-green-900">Dashboard</h3>
          </router-link>

          <router-link
            v-if="authStore.user?.role === 'admin'"
            to="/admin"
            class="p-4 bg-purple-100 rounded hover:bg-purple-200 transition"
          >
            <h3 class="font-semibold text-purple-900">Manajemen Admin</h3>
          </router-link>

          <router-link
            v-if="authStore.user?.role === 'admin'"
            to="/rumah-sakit"
            class="p-4 bg-orange-100 rounded hover:bg-orange-200 transition"
          >
            <h3 class="font-semibold text-orange-900">Manajemen Rumah Sakit</h3>
          </router-link>
        </div>
      </div>
    </div>
  </div>
</template>
