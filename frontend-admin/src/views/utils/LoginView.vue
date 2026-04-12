<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { LogIn, AlertCircle } from '@lucide/vue'

const router = useRouter()
const authStore = useAuthStore()

const username = ref('')
const password = ref('')
const errorMsg = ref('')
const isLoading = ref(false)
const showPassword = ref(false)

const handleLogin = async () => {
  if (!username.value || !password.value) {
    errorMsg.value = 'Username dan password harus diisi'
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
  <div class="min-h-screen bg-gray-50 flex items-center justify-center px-4">
    <div class="w-full max-w-sm">
      <!-- Brand -->
      <div class="flex flex-col items-center mb-8">
        <div class="w-12 h-12 bg-red-100 rounded-xl flex items-center justify-center mb-4">
          <span class="text-red-600 font-bold text-sm">PMI</span>
        </div>
        <h1 class="text-xl font-bold text-gray-900">
          PMI <span class="text-red-500">Salatiga</span>
        </h1>
        <p class="text-sm text-gray-500 mt-1 font-bold">Portal Admin</p>
      </div>

      <!-- Card -->
      <div class="bg-white rounded-2xl border border-gray-200 shadow-md p-8">
        <h2 class="text-base font-semibold text-gray-800 mb-6">Masuk ke akun Anda</h2>

        <form @submit.prevent="handleLogin" class="space-y-4">
          <!-- Username -->
          <div>
            <label class="block text-xs font-medium text-gray-500 mb-1.5 uppercase tracking-wide">
              Username
            </label>
            <input
              v-model="username"
              type="text"
              autocomplete="username"
              placeholder="Masukkan username"
              class="w-full px-3.5 py-2.5 text-sm text-gray-900 bg-gray-50 border border-gray-200 rounded-xl outline-none transition-all focus:bg-white focus:border-blue-400 focus:ring-2 focus:ring-blue-100 placeholder:text-gray-300"
            />
          </div>

          <!-- Password -->
          <div>
            <label class="block text-xs font-medium text-gray-500 mb-1.5 uppercase tracking-wide">
              Password
            </label>
            <div class="relative">
              <input
                v-model="password"
                :type="showPassword ? 'text' : 'password'"
                autocomplete="current-password"
                placeholder="••••••••"
                class="w-full px-3.5 py-2.5 pr-10 text-sm text-gray-900 bg-gray-50 border border-gray-200 rounded-xl outline-none transition-all focus:bg-white focus:border-blue-400 focus:ring-2 focus:ring-blue-100 placeholder:text-gray-300"
              />
              <button
                type="button"
                @click="showPassword = !showPassword"
                class="absolute right-3 top-1/2 -translate-y-1/2 text-gray-900 hover:text-gray-600 transition-colors"
              >
                <Eye v-if="!showPassword" :size="16" />
                <EyeOff v-else :size="16" />
              </button>
            </div>
          </div>

          <!-- Error -->
          <div
            v-if="errorMsg"
            class="flex items-center gap-2 p-3 bg-red-50 border border-red-100 text-red-600 rounded-xl text-xs"
          >
            <AlertCircle :size="14" class="shrink-0" />
            {{ errorMsg }}
          </div>

          <!-- Submit -->
          <button
            type="submit"
            :disabled="isLoading"
            class="flex items-center justify-center gap-2 w-full py-2.5 px-4 bg-blue-600 hover:bg-blue-700 disabled:opacity-50 disabled:cursor-not-allowed text-white text-sm font-medium rounded-xl transition-colors duration-200 mt-2"
          >
            <span
              v-if="isLoading"
              class="w-4 h-4 border-2 border-white/30 border-t-white rounded-full animate-spin"
            />
            <LogIn v-else :size="16" />
            {{ isLoading ? 'Memproses...' : 'Masuk' }}
          </button>
        </form>
      </div>

      <p class="text-center text-xs text-gray-300 mt-6">
        PMI Salatiga {{ new Date().getFullYear() }}
      </p>
    </div>
  </div>
</template>
