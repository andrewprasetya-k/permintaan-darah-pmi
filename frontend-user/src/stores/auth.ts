import { computed, ref } from 'vue'
import { defineStore } from 'pinia'
import { AUTH_TOKEN_KEY, AUTH_USER_KEY } from '@/api/client'
import { authAPI } from '@/api/auth'
import type { User } from '@/types/models'

export const useAuthStore = defineStore('auth', () => {
  const user = ref<User | null>(null)
  const token = ref<string | null>(localStorage.getItem(AUTH_TOKEN_KEY))
  const isLoading = ref(false)
  const error = ref<string | null>(null)

  const isAuthenticated = computed(() => Boolean(token.value && user.value))
  const isRumahSakit = computed(() => user.value?.role === 'rumah_sakit')

  const initializeFromStorage = () => {
    const savedToken = localStorage.getItem(AUTH_TOKEN_KEY)
    const savedUser = localStorage.getItem(AUTH_USER_KEY)

    token.value = savedToken
    if (!savedUser || savedUser === 'undefined') {
      user.value = null
      return
    }

    try {
      user.value = JSON.parse(savedUser) as User
    } catch {
      localStorage.removeItem(AUTH_USER_KEY)
      user.value = null
    }
  }

  const login = async (username: string, password: string) => {
    isLoading.value = true
    error.value = null

    try {
      const response = await authAPI.loginRumahSakit({ username, password })
      const loginUser: User = {
        id: response.data.id,
        username: response.data.username,
        role: response.data.role,
      }

      if (loginUser.role !== 'rumah_sakit') {
        throw new Error('Akun ini bukan akun rumah sakit')
      }

      token.value = response.data.token
      user.value = loginUser
      localStorage.setItem(AUTH_TOKEN_KEY, response.data.token)
      localStorage.setItem(AUTH_USER_KEY, JSON.stringify(loginUser))
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Login gagal'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  const logout = () => {
    token.value = null
    user.value = null
    localStorage.removeItem(AUTH_TOKEN_KEY)
    localStorage.removeItem(AUTH_USER_KEY)
  }

  return {
    user,
    token,
    isLoading,
    error,
    isAuthenticated,
    isRumahSakit,
    initializeFromStorage,
    login,
    logout,
  }
})
