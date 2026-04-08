import { defineStore } from 'pinia'
import { ref } from 'vue'
import { dashboardAPI, type StatusSummary } from '@/api/dashboard'

export const useDashboardStore = defineStore('dashboard', () => {
  const statusSummary = ref<StatusSummary | null>(null)
  const isLoading = ref(false)
  const error = ref<string | null>(null)

  const fetchStatusSummary = async (rumahSakitID: string) => {
    isLoading.value = true
    error.value = null
    try {
      const response = await dashboardAPI.getStatusSummary(rumahSakitID)
      statusSummary.value = response.data
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to fetch status summary'
    } finally {
      isLoading.value = false
    }
  }

  return {
    statusSummary,
    isLoading,
    error,
    fetchStatusSummary,
  }
})
