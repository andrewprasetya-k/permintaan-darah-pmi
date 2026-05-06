import { ref } from 'vue'
import { defineStore } from 'pinia'
import { myProfileAPI } from '@/api/my-profile'
import type { RumahSakit, UpdateRumahSakitProfileRequest } from '@/types/models'

export const useMyProfileStore = defineStore('myProfile', () => {
  const profile = ref<RumahSakit | null>(null)
  const isLoading = ref(false)
  const error = ref<string | null>(null)

  const fetchMe = async () => {
    isLoading.value = true
    error.value = null

    try {
      const response = await myProfileAPI.getMe()
      profile.value = response.data
      return response.data
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Gagal memuat profil'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  const updateMe = async (payload: UpdateRumahSakitProfileRequest) => {
    isLoading.value = true
    error.value = null

    try {
      const response = await myProfileAPI.updateMe(payload)
      profile.value = response.data
      return response.data
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Gagal menyimpan profil'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  return {
    profile,
    isLoading,
    error,
    fetchMe,
    updateMe,
  }
})
