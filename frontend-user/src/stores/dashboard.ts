import { ref } from 'vue'
import { defineStore } from 'pinia'
import { dashboardAPI } from '@/api/dashboard'
import type { StatusSummary } from '@/types/models'

const emptySummary = (): StatusSummary => ({
  dibuat: 0,
  diproses: 0,
  bisaDiambil: 0,
  selesai: 0,
  dibatalkan: 0,
  total: 0,
})

export const useDashboardStore = defineStore('dashboard', () => {
  const summary = ref<StatusSummary>(emptySummary())
  const isLoading = ref(false)
  const error = ref<string | null>(null)

  const fetchSummary = async (rumahSakitId: string) => {
    isLoading.value = true
    error.value = null

    try {
      const response = await dashboardAPI.getStatusSummary(rumahSakitId)
      summary.value = response.data
      return response.data
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Gagal memuat ringkasan'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  return {
    summary,
    isLoading,
    error,
    fetchSummary,
  }
})
