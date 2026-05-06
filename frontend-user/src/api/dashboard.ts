import { apiClient } from '@/api/client'
import type { StatusSummary } from '@/types/models'

export const dashboardAPI = {
  getStatusSummary(rumahSakitId: string) {
    return apiClient.get<StatusSummary>(`/dashboard/status-summary/${rumahSakitId}`)
  },
}
