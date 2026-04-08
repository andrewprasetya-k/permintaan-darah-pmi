import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { authAPI, type LoginResponse } from '@/api/auth'
import type { User } from '@/types/models'

export const useAuthStore = defineStore('auth', () => {
  const user = ref<User | null>(null)
  const token = ref<string | null>(localStorage.getItem('authToken'))
  const isLoading = ref(false)
  const error = ref<string | null>(null)

  const isAuthenticated = computed(() => !!token.value && !!user.value)

  const loginAdmin = async (username: string, password: string) => {
    isLoading.value = true
    error.value = null
    try {
      const response = await authAPI.loginAdmin({ username, password })
      token.value = response.data.token
      user.value = response.data.user
      localStorage.setItem('authToken', response.data.token)
      localStorage.setItem('authUser', JSON.stringify(response.data.user))
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Login failed'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  const logout = () => {
    token.value = null
    user.value = null
    authAPI.logout()
  }

  const initializeFromStorage = () => {
    const savedToken = localStorage.getItem('authToken')
    const savedUser = localStorage.getItem('authUser')
    if (savedToken && savedUser) {
      token.value = savedToken
      user.value = JSON.parse(savedUser)
    }
  }

  return {
    user,
    token,
    isLoading,
    error,
    isAuthenticated,
    loginAdmin,
    logout,
    initializeFromStorage,
  }
})
