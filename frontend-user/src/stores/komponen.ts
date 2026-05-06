import { ref } from 'vue'
import { defineStore } from 'pinia'
import { komponenAPI } from '@/api/komponen'
import type { KomponenDarah } from '@/types/models'

export const useKomponenStore = defineStore('komponen', () => {
  const components = ref<KomponenDarah[]>([])
  const isLoading = ref(false)
  const error = ref<string | null>(null)

  const fetchAll = async () => {
    isLoading.value = true
    error.value = null

    try {
      const response = await komponenAPI.getAll()
      components.value = response.data
      return response.data
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Gagal memuat komponen darah'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  return {
    components,
    isLoading,
    error,
    fetchAll,
  }
})
