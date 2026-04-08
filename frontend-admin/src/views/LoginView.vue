<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const authStore = useAuthStore()

const username = ref('')
const password = ref('')
const errorMsg = ref('')
const isLoading = ref(false)

const handleLogin = async () => {
  if (!username.value || !password.value) {
    errorMsg.value = 'username dan password harus diisi'
    return
  }

  isLoading.value = true
  errorMsg.value = ''

  try {
    await authStore.loginAdmin(username.value, password.value)
    const redirect = router.currentRoute.value.query.redirect as string
    router.push(redirect || '/')
  } catch (err) {
    errorMsg.value = err instanceof Error ? err.message : 'Login gagal'
  } finally {
    isLoading.value = false
  }
}
</script>

<template>
  <div class="max-h-screen flex items-center justify-center bg-gray-50">
    <div class="max-w-md w-full space-y-8">
      <h1 class="text-center text-3xl font-bold">Admin Dashboard</h1>

      <form @submit.prevent="handleLogin" class="space-y-4">
        <div>
          <label class="block text-sm font-medium">Username</label>
          <input
            v-model="username"
            type="text"
            class="mt-1 w-full px-3 py-2 border border-gray-300 rounded-md"
            placeholder="Masukkan username"
          />
        </div>

        <div>
          <label class="block text-sm font-medium">Password</label>
          <input
            v-model="password"
            type="password"
            class="mt-1 w-full px-3 py-2 border border-gray-300 rounded-md"
            placeholder="••••••••"
          />
        </div>

        <div v-if="errorMsg" class="p-3 bg-red-100 text-red-700 rounded">
          {{ errorMsg }}
        </div>

        <button
          type="submit"
          :disabled="isLoading"
          class="w-full bg-blue-600 text-white py-2 rounded-md hover:bg-blue-700 disabled:opacity-50"
        >
          {{ isLoading ? 'Loading...' : 'Login' }}
        </button>
      </form>
    </div>
  </div>
</template>
